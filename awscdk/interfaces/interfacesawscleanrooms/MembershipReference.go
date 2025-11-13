package interfacesawscleanrooms


// A reference to a Membership resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   membershipReference := &MembershipReference{
//   	MembershipArn: jsii.String("membershipArn"),
//   	MembershipIdentifier: jsii.String("membershipIdentifier"),
//   }
//
type MembershipReference struct {
	// The ARN of the Membership resource.
	MembershipArn *string `field:"required" json:"membershipArn" yaml:"membershipArn"`
	// The MembershipIdentifier of the Membership resource.
	MembershipIdentifier *string `field:"required" json:"membershipIdentifier" yaml:"membershipIdentifier"`
}

