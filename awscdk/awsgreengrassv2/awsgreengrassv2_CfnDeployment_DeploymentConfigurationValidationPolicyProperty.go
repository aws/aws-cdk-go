package awsgreengrassv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentConfigurationValidationPolicyProperty := &deploymentConfigurationValidationPolicyProperty{
//   	timeoutInSeconds: jsii.Number(123),
//   }
//
type CfnDeployment_DeploymentConfigurationValidationPolicyProperty struct {
	// `CfnDeployment.DeploymentConfigurationValidationPolicyProperty.TimeoutInSeconds`.
	TimeoutInSeconds *float64 `field:"optional" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
}

