package awsdetective


// A reference to a MemberInvitation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memberInvitationReference := &MemberInvitationReference{
//   	GraphArn: jsii.String("graphArn"),
//   	MemberId: jsii.String("memberId"),
//   }
//
type MemberInvitationReference struct {
	// The GraphArn of the MemberInvitation resource.
	GraphArn *string `field:"required" json:"graphArn" yaml:"graphArn"`
	// The MemberId of the MemberInvitation resource.
	MemberId *string `field:"required" json:"memberId" yaml:"memberId"`
}

