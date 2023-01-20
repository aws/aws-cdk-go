package awsidentitystore


// Properties for defining a `CfnGroupMembership`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGroupMembershipProps := &cfnGroupMembershipProps{
//   	groupId: jsii.String("groupId"),
//   	identityStoreId: jsii.String("identityStoreId"),
//   	memberId: &memberIdProperty{
//   		userId: jsii.String("userId"),
//   	},
//   }
//
type CfnGroupMembershipProps struct {
	// `AWS::IdentityStore::GroupMembership.GroupId`.
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// `AWS::IdentityStore::GroupMembership.IdentityStoreId`.
	IdentityStoreId *string `field:"required" json:"identityStoreId" yaml:"identityStoreId"`
	// `AWS::IdentityStore::GroupMembership.MemberId`.
	MemberId interface{} `field:"required" json:"memberId" yaml:"memberId"`
}

