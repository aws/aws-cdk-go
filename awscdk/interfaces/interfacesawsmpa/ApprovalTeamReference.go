package interfacesawsmpa


// A reference to a ApprovalTeam resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   approvalTeamReference := &ApprovalTeamReference{
//   	ApprovalTeamArn: jsii.String("approvalTeamArn"),
//   }
//
type ApprovalTeamReference struct {
	// The Arn of the ApprovalTeam resource.
	ApprovalTeamArn *string `field:"required" json:"approvalTeamArn" yaml:"approvalTeamArn"`
}

