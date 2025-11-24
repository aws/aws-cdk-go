package mixinsawssecurityhub


// The updated note.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   noteUpdateProperty := &NoteUpdateProperty{
//   	Text: jsii.String("text"),
//   	UpdatedBy: jsii.String("updatedBy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-noteupdate.html
//
type CfnAutomationRulePropsMixin_NoteUpdateProperty struct {
	// The updated note text.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-noteupdate.html#cfn-securityhub-automationrule-noteupdate-text
	//
	Text *string `field:"optional" json:"text" yaml:"text"`
	// The principal that updated the note.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-noteupdate.html#cfn-securityhub-automationrule-noteupdate-updatedby
	//
	UpdatedBy *string `field:"optional" json:"updatedBy" yaml:"updatedBy"`
}

