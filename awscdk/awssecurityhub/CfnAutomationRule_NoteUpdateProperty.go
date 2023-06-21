package awssecurityhub


// The updated note.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var updatedBy interface{}
//
//   noteUpdateProperty := &NoteUpdateProperty{
//   	Text: jsii.String("text"),
//   	UpdatedBy: updatedBy,
//   }
//
type CfnAutomationRule_NoteUpdateProperty struct {
	// The updated note text.
	Text *string `field:"required" json:"text" yaml:"text"`
	// The principal that updated the note.
	UpdatedBy interface{} `field:"required" json:"updatedBy" yaml:"updatedBy"`
}

