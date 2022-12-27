package awsgreengrassv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentComponentUpdatePolicyProperty := &deploymentComponentUpdatePolicyProperty{
//   	action: jsii.String("action"),
//   	timeoutInSeconds: jsii.Number(123),
//   }
//
type CfnDeployment_DeploymentComponentUpdatePolicyProperty struct {
	// `CfnDeployment.DeploymentComponentUpdatePolicyProperty.Action`.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// `CfnDeployment.DeploymentComponentUpdatePolicyProperty.TimeoutInSeconds`.
	TimeoutInSeconds *float64 `field:"optional" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
}

