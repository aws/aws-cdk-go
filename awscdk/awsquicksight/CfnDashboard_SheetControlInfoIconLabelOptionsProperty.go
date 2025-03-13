package awsquicksight


// A control to display info icons for filters and parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sheetControlInfoIconLabelOptionsProperty := &SheetControlInfoIconLabelOptionsProperty{
//   	InfoIconText: jsii.String("infoIconText"),
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetcontrolinfoiconlabeloptions.html
//
type CfnDashboard_SheetControlInfoIconLabelOptionsProperty struct {
	// The text content of info icon.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetcontrolinfoiconlabeloptions.html#cfn-quicksight-dashboard-sheetcontrolinfoiconlabeloptions-infoicontext
	//
	InfoIconText *string `field:"optional" json:"infoIconText" yaml:"infoIconText"`
	// The visibility configuration of info icon label options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetcontrolinfoiconlabeloptions.html#cfn-quicksight-dashboard-sheetcontrolinfoiconlabeloptions-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

