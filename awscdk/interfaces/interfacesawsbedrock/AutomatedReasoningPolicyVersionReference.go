package interfacesawsbedrock


// A reference to a AutomatedReasoningPolicyVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   automatedReasoningPolicyVersionReference := &AutomatedReasoningPolicyVersionReference{
//   	PolicyArn: jsii.String("policyArn"),
//   	Version: jsii.String("version"),
//   }
//
type AutomatedReasoningPolicyVersionReference struct {
	// The PolicyArn of the AutomatedReasoningPolicyVersion resource.
	PolicyArn *string `field:"required" json:"policyArn" yaml:"policyArn"`
	// The Version of the AutomatedReasoningPolicyVersion resource.
	Version *string `field:"required" json:"version" yaml:"version"`
}

