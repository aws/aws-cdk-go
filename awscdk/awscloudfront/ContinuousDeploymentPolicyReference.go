package awscloudfront


// A reference to a ContinuousDeploymentPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   continuousDeploymentPolicyReference := &ContinuousDeploymentPolicyReference{
//   	ContinuousDeploymentPolicyId: jsii.String("continuousDeploymentPolicyId"),
//   }
//
type ContinuousDeploymentPolicyReference struct {
	// The Id of the ContinuousDeploymentPolicy resource.
	ContinuousDeploymentPolicyId *string `field:"required" json:"continuousDeploymentPolicyId" yaml:"continuousDeploymentPolicyId"`
}

