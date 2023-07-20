package awsquicksight


// A text box.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sheetTextBoxProperty := &SheetTextBoxProperty{
//   	SheetTextBoxId: jsii.String("sheetTextBoxId"),
//
//   	// the properties below are optional
//   	Content: jsii.String("content"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheettextbox.html
//
type CfnDashboard_SheetTextBoxProperty struct {
	// The unique identifier for a text box.
	//
	// This identifier must be unique within the context of a dashboard, template, or analysis. Two dashboards, analyses, or templates can have text boxes that share identifiers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheettextbox.html#cfn-quicksight-dashboard-sheettextbox-sheettextboxid
	//
	SheetTextBoxId *string `field:"required" json:"sheetTextBoxId" yaml:"sheetTextBoxId"`
	// The content that is displayed in the text box.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheettextbox.html#cfn-quicksight-dashboard-sheettextbox-content
	//
	Content *string `field:"optional" json:"content" yaml:"content"`
}

