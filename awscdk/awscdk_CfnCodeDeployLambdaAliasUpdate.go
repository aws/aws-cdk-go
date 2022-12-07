// Version 2 of the AWS Cloud Development Kit library
package awscdk


// To perform an AWS CodeDeploy deployment when the version changes on an AWS::Lambda::Alias resource, use the CodeDeployLambdaAliasUpdate update policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCodeDeployLambdaAliasUpdate := &cfnCodeDeployLambdaAliasUpdate{
//   	applicationName: jsii.String("applicationName"),
//   	deploymentGroupName: jsii.String("deploymentGroupName"),
//
//   	// the properties below are optional
//   	afterAllowTrafficHook: jsii.String("afterAllowTrafficHook"),
//   	beforeAllowTrafficHook: jsii.String("beforeAllowTrafficHook"),
//   }
//
type CfnCodeDeployLambdaAliasUpdate struct {
	// The name of the AWS CodeDeploy application.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// The name of the AWS CodeDeploy deployment group.
	//
	// This is where the traffic-shifting policy is set.
	DeploymentGroupName *string `field:"required" json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// The name of the Lambda function to run after traffic routing completes.
	AfterAllowTrafficHook *string `field:"optional" json:"afterAllowTrafficHook" yaml:"afterAllowTrafficHook"`
	// The name of the Lambda function to run before traffic routing starts.
	BeforeAllowTrafficHook *string `field:"optional" json:"beforeAllowTrafficHook" yaml:"beforeAllowTrafficHook"`
}

