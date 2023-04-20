// An experiment to bundle the entire CDK into a single module
package awscdk


// To perform an AWS CodeDeploy deployment when the version changes on an AWS::Lambda::Alias resource, use the CodeDeployLambdaAliasUpdate update policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCodeDeployLambdaAliasUpdate := &CfnCodeDeployLambdaAliasUpdate{
//   	ApplicationName: jsii.String("applicationName"),
//   	DeploymentGroupName: jsii.String("deploymentGroupName"),
//
//   	// the properties below are optional
//   	AfterAllowTrafficHook: jsii.String("afterAllowTrafficHook"),
//   	BeforeAllowTrafficHook: jsii.String("beforeAllowTrafficHook"),
//   }
//
// Experimental.
type CfnCodeDeployLambdaAliasUpdate struct {
	// The name of the AWS CodeDeploy application.
	// Experimental.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// The name of the AWS CodeDeploy deployment group.
	//
	// This is where the traffic-shifting policy is set.
	// Experimental.
	DeploymentGroupName *string `field:"required" json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// The name of the Lambda function to run after traffic routing completes.
	// Experimental.
	AfterAllowTrafficHook *string `field:"optional" json:"afterAllowTrafficHook" yaml:"afterAllowTrafficHook"`
	// The name of the Lambda function to run before traffic routing starts.
	// Experimental.
	BeforeAllowTrafficHook *string `field:"optional" json:"beforeAllowTrafficHook" yaml:"beforeAllowTrafficHook"`
}

