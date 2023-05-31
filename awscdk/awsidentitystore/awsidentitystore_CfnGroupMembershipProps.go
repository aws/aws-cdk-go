package awsidentitystore


// Properties for defining a `CfnGroupMembership`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGroupMembershipProps := &CfnGroupMembershipProps{
//   	GroupId: jsii.String("groupId"),
//   	IdentityStoreId: jsii.String("identityStoreId"),
//   	MemberId: &MemberIdProperty{
//   		UserId: jsii.String("userId"),
//   	},
//   }
//
type CfnGroupMembershipProps struct {
	// `AWS::IdentityStore::GroupMembership.GroupId`.
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// `AWS::IdentityStore::GroupMembership.IdentityStoreId`.
	IdentityStoreId *string `field:"required" json:"identityStoreId" yaml:"identityStoreId"`
	// An object containing the identifier of a group member.
	//
	// Setting `MemberId` 's `UserId` field to a specific User's ID indicates we should consider that User as a group member.
	MemberId interface{} `field:"required" json:"memberId" yaml:"memberId"`
}

