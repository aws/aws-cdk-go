package interfacesawsbedrockagentcore


// A reference to a Policy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyReference := &PolicyReference{
//   	PolicyArn: jsii.String("policyArn"),
//   }
//
type PolicyReference struct {
	// The PolicyArn of the Policy resource.
	PolicyArn *string `field:"required" json:"policyArn" yaml:"policyArn"`
}

