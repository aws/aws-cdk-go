package previewawsidentitystoremixins


// Properties for CfnGroupMembershipPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGroupMembershipMixinProps := &CfnGroupMembershipMixinProps{
//   	GroupId: jsii.String("groupId"),
//   	IdentityStoreId: jsii.String("identityStoreId"),
//   	MemberId: &MemberIdProperty{
//   		UserId: jsii.String("userId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-groupmembership.html
//
type CfnGroupMembershipMixinProps struct {
	// The identifier for a group in the identity store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-groupmembership.html#cfn-identitystore-groupmembership-groupid
	//
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// The globally unique identifier for the identity store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-groupmembership.html#cfn-identitystore-groupmembership-identitystoreid
	//
	IdentityStoreId *string `field:"optional" json:"identityStoreId" yaml:"identityStoreId"`
	// An object containing the identifier of a group member.
	//
	// Setting the `MemberId` 's `UserId` field to a specific User's ID indicates that user is a member of the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-groupmembership.html#cfn-identitystore-groupmembership-memberid
	//
	MemberId interface{} `field:"optional" json:"memberId" yaml:"memberId"`
}

