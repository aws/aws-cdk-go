package awsquicksight


// A text box.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var interactions interface{}
//
//   sheetTextBoxProperty := &SheetTextBoxProperty{
//   	SheetTextBoxId: jsii.String("sheetTextBoxId"),
//
//   	// the properties below are optional
//   	Content: jsii.String("content"),
//   	Interactions: interactions,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheettextbox.html
//
type CfnTemplate_SheetTextBoxProperty struct {
	// The unique identifier for a text box.
	//
	// This identifier must be unique within the context of a dashboard, template, or analysis. Two dashboards, analyses, or templates can have text boxes that share identifiers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheettextbox.html#cfn-quicksight-template-sheettextbox-sheettextboxid
	//
	SheetTextBoxId *string `field:"required" json:"sheetTextBoxId" yaml:"sheetTextBoxId"`
	// The content that is displayed in the text box.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheettextbox.html#cfn-quicksight-template-sheettextbox-content
	//
	Content *string `field:"optional" json:"content" yaml:"content"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheettextbox.html#cfn-quicksight-template-sheettextbox-interactions
	//
	Interactions interface{} `field:"optional" json:"interactions" yaml:"interactions"`
}

