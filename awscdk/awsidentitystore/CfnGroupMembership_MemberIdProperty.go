package awsidentitystore


// An object containing the identifier of a group member.
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
	// An object containing the identifiers of resources that can be members.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-groupmembership-memberid.html#cfn-identitystore-groupmembership-memberid-userid
	//
	UserId *string `field:"required" json:"userId" yaml:"userId"`
}

