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
type CfnGroupMembership_MemberIdProperty struct {
	// `CfnGroupMembership.MemberIdProperty.UserId`.
	UserId *string `field:"required" json:"userId" yaml:"userId"`
}

