package interfacesawssmsvoice


// A reference to a RcsAgent resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rcsAgentReference := &RcsAgentReference{
//   	RcsAgentArn: jsii.String("rcsAgentArn"),
//   }
//
type RcsAgentReference struct {
	// The RcsAgentArn of the RcsAgent resource.
	RcsAgentArn *string `field:"required" json:"rcsAgentArn" yaml:"rcsAgentArn"`
}

