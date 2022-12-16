package awsgreengrassv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentPoliciesProperty := &deploymentPoliciesProperty{
//   	componentUpdatePolicy: &deploymentComponentUpdatePolicyProperty{
//   		action: jsii.String("action"),
//   		timeoutInSeconds: jsii.Number(123),
//   	},
//   	configurationValidationPolicy: &deploymentConfigurationValidationPolicyProperty{
//   		timeoutInSeconds: jsii.Number(123),
//   	},
//   	failureHandlingPolicy: jsii.String("failureHandlingPolicy"),
//   }
//
type CfnDeployment_DeploymentPoliciesProperty struct {
	// `CfnDeployment.DeploymentPoliciesProperty.ComponentUpdatePolicy`.
	ComponentUpdatePolicy interface{} `field:"optional" json:"componentUpdatePolicy" yaml:"componentUpdatePolicy"`
	// `CfnDeployment.DeploymentPoliciesProperty.ConfigurationValidationPolicy`.
	ConfigurationValidationPolicy interface{} `field:"optional" json:"configurationValidationPolicy" yaml:"configurationValidationPolicy"`
	// `CfnDeployment.DeploymentPoliciesProperty.FailureHandlingPolicy`.
	FailureHandlingPolicy *string `field:"optional" json:"failureHandlingPolicy" yaml:"failureHandlingPolicy"`
}

