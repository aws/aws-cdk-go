package pipelines


// Construction properties for a `ManualApprovalStep`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   manualApprovalStepProps := &manualApprovalStepProps{
//   	comment: jsii.String("comment"),
//   }
//
type ManualApprovalStepProps struct {
	// The comment to display with this manual approval.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

