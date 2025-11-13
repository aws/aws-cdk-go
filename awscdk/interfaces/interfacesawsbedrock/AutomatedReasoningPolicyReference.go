package interfacesawsbedrock


// A reference to a AutomatedReasoningPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   automatedReasoningPolicyReference := &AutomatedReasoningPolicyReference{
//   	PolicyArn: jsii.String("policyArn"),
//   }
//
type AutomatedReasoningPolicyReference struct {
	// The PolicyArn of the AutomatedReasoningPolicy resource.
	PolicyArn *string `field:"required" json:"policyArn" yaml:"policyArn"`
}

