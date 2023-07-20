package awsidentitystore


// An object that contains the identifier of a group member.
//
// Setting the `UserID` field to the specific identifier for a user indicates that the user is a member of the group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memberIdProperty := &MemberIdProperty{
//   	UserId: jsii.String("userId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-groupmembership-memberid.html
//
type CfnGroupMembership_MemberIdProperty struct {
	// The identifier for a user in the identity store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-groupmembership-memberid.html#cfn-identitystore-groupmembership-memberid-userid
	//
	UserId *string `field:"required" json:"userId" yaml:"userId"`
}

