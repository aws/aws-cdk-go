package interfacesawsidentitystore


// A reference to a GroupMembership resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   groupMembershipReference := &GroupMembershipReference{
//   	IdentityStoreId: jsii.String("identityStoreId"),
//   	MembershipId: jsii.String("membershipId"),
//   }
//
type GroupMembershipReference struct {
	// The IdentityStoreId of the GroupMembership resource.
	IdentityStoreId *string `field:"required" json:"identityStoreId" yaml:"identityStoreId"`
	// The MembershipId of the GroupMembership resource.
	MembershipId *string `field:"required" json:"membershipId" yaml:"membershipId"`
}

