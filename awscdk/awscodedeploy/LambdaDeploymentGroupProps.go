package awscodedeploy

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudwatch"
)

// Construction properties for `LambdaDeploymentGroup`.
//
// Example:
//   var myApplication LambdaApplication
//   var func Function
//
//   version := func.currentVersion
//   version1Alias := lambda.NewAlias(this, jsii.String("alias"), &AliasProps{
//   	AliasName: jsii.String("prod"),
//   	Version: Version,
//   })
//
//   deploymentGroup := codedeploy.NewLambdaDeploymentGroup(this, jsii.String("BlueGreenDeployment"), &LambdaDeploymentGroupProps{
//   	Application: myApplication,
//   	 // optional property: one will be created for you if not provided
//   	Alias: version1Alias,
//   	DeploymentConfig: codedeploy.LambdaDeploymentConfig_LINEAR_10PERCENT_EVERY_1MINUTE(),
//   })
//
type LambdaDeploymentGroupProps struct {
	// Lambda Alias to shift traffic. Updating the version of the alias will trigger a CodeDeploy deployment.
	//
	// [disable-awslint:ref-via-interface] since we need to modify the alias CFN resource update policy.
	Alias awslambda.Alias `field:"required" json:"alias" yaml:"alias"`
	// The CloudWatch alarms associated with this Deployment Group.
	//
	// CodeDeploy will stop (and optionally roll back)
	// a deployment if during it any of the alarms trigger.
	//
	// Alarms can also be added after the Deployment Group is created using the `#addAlarm` method.
	// See: https://docs.aws.amazon.com/codedeploy/latest/userguide/monitoring-create-alarms.html
	//
	// Default: [].
	//
	Alarms *[]interfacesawscloudwatch.IAlarmRef `field:"optional" json:"alarms" yaml:"alarms"`
	// The reference to the CodeDeploy Lambda Application that this Deployment Group belongs to.
	// Default: - One will be created for you.
	//
	Application ILambdaApplication `field:"optional" json:"application" yaml:"application"`
	// The auto-rollback configuration for this Deployment Group.
	// Default: - default AutoRollbackConfig.
	//
	AutoRollback *AutoRollbackConfig `field:"optional" json:"autoRollback" yaml:"autoRollback"`
	// The Deployment Configuration this Deployment Group uses.
	// Default: LambdaDeploymentConfig.CANARY_10PERCENT_5MINUTES
	//
	DeploymentConfig ILambdaDeploymentConfig `field:"optional" json:"deploymentConfig" yaml:"deploymentConfig"`
	// The physical, human-readable name of the CodeDeploy Deployment Group.
	// Default: - An auto-generated name will be used.
	//
	DeploymentGroupName *string `field:"optional" json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// Whether to skip the step of checking CloudWatch alarms during the deployment process.
	// Default: - false.
	//
	IgnoreAlarmConfiguration *bool `field:"optional" json:"ignoreAlarmConfiguration" yaml:"ignoreAlarmConfiguration"`
	// Whether to continue a deployment even if fetching the alarm status from CloudWatch failed.
	// Default: false.
	//
	IgnorePollAlarmsFailure *bool `field:"optional" json:"ignorePollAlarmsFailure" yaml:"ignorePollAlarmsFailure"`
	// The Lambda function to run after traffic routing starts.
	// Default: - None.
	//
	PostHook awslambda.IFunction `field:"optional" json:"postHook" yaml:"postHook"`
	// The Lambda function to run before traffic routing starts.
	// Default: - None.
	//
	PreHook awslambda.IFunction `field:"optional" json:"preHook" yaml:"preHook"`
	// The service Role of this Deployment Group.
	// Default: - A new Role will be created.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

