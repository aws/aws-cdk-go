package awsidentitystore


// Properties for CfnAllGroupMembershipsPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnAllGroupMembershipsMixinProps := &CfnAllGroupMembershipsMixinProps{
//   	GroupId: jsii.String("groupId"),
//   	IdentityStoreId: jsii.String("identityStoreId"),
//   	MemberId: &MemberIdProperty{
//   		UserId: jsii.String("userId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-allgroupmemberships.html
//
type CfnAllGroupMembershipsMixinProps struct {
	// The identifier for a group in the identity store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-allgroupmemberships.html#cfn-identitystore-allgroupmemberships-groupid
	//
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// The globally unique identifier for the identity store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-allgroupmemberships.html#cfn-identitystore-allgroupmemberships-identitystoreid
	//
	IdentityStoreId *string `field:"optional" json:"identityStoreId" yaml:"identityStoreId"`
	// An object containing the identifier of a group member.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-allgroupmemberships.html#cfn-identitystore-allgroupmemberships-memberid
	//
	MemberId interface{} `field:"optional" json:"memberId" yaml:"memberId"`
}

