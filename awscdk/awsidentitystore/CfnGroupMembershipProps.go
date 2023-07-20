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
	// The unique identifier for a group in the identity store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-groupmembership.html#cfn-identitystore-groupmembership-groupid
	//
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// The globally unique identifier for the identity store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-groupmembership.html#cfn-identitystore-groupmembership-identitystoreid
	//
	IdentityStoreId *string `field:"required" json:"identityStoreId" yaml:"identityStoreId"`
	// An object containing the identifier of a group member.
	//
	// Setting `MemberId` 's `UserId` field to a specific User's ID indicates we should consider that User as a group member.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-groupmembership.html#cfn-identitystore-groupmembership-memberid
	//
	MemberId interface{} `field:"required" json:"memberId" yaml:"memberId"`
}

