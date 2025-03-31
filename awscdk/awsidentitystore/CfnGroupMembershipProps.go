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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-groupmembership.html
//
type CfnGroupMembershipProps struct {
	// The identifier for a group in the identity store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-groupmembership.html#cfn-identitystore-groupmembership-groupid
	//
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// The globally unique identifier for the identity store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-groupmembership.html#cfn-identitystore-groupmembership-identitystoreid
	//
	IdentityStoreId *string `field:"required" json:"identityStoreId" yaml:"identityStoreId"`
	// An object that contains the identifier of a group member.
	//
	// Setting the `UserID` field to the specific identifier for a user indicates that the user is a member of the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-groupmembership.html#cfn-identitystore-groupmembership-memberid
	//
	MemberId interface{} `field:"required" json:"memberId" yaml:"memberId"`
}

