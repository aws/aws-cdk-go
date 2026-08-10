package interfacesawscleanrooms


// A reference to a IntermediateTable resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   intermediateTableReference := &IntermediateTableReference{
//   	IntermediateTableArn: jsii.String("intermediateTableArn"),
//   	IntermediateTableIdentifier: jsii.String("intermediateTableIdentifier"),
//   	MembershipIdentifier: jsii.String("membershipIdentifier"),
//   }
//
type IntermediateTableReference struct {
	// The ARN of the IntermediateTable resource.
	IntermediateTableArn *string `field:"required" json:"intermediateTableArn" yaml:"intermediateTableArn"`
	// The IntermediateTableIdentifier of the IntermediateTable resource.
	IntermediateTableIdentifier *string `field:"required" json:"intermediateTableIdentifier" yaml:"intermediateTableIdentifier"`
	// The MembershipIdentifier of the IntermediateTable resource.
	MembershipIdentifier *string `field:"required" json:"membershipIdentifier" yaml:"membershipIdentifier"`
}

