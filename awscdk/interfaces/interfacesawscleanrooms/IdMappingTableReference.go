package interfacesawscleanrooms


// A reference to a IdMappingTable resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   idMappingTableReference := &IdMappingTableReference{
//   	IdMappingTableArn: jsii.String("idMappingTableArn"),
//   	IdMappingTableIdentifier: jsii.String("idMappingTableIdentifier"),
//   	MembershipIdentifier: jsii.String("membershipIdentifier"),
//   }
//
type IdMappingTableReference struct {
	// The ARN of the IdMappingTable resource.
	IdMappingTableArn *string `field:"required" json:"idMappingTableArn" yaml:"idMappingTableArn"`
	// The IdMappingTableIdentifier of the IdMappingTable resource.
	IdMappingTableIdentifier *string `field:"required" json:"idMappingTableIdentifier" yaml:"idMappingTableIdentifier"`
	// The MembershipIdentifier of the IdMappingTable resource.
	MembershipIdentifier *string `field:"required" json:"membershipIdentifier" yaml:"membershipIdentifier"`
}

