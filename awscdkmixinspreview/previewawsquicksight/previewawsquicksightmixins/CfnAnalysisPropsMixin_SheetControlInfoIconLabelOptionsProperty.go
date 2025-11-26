package previewawsquicksightmixins


// A control to display info icons for filters and parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sheetControlInfoIconLabelOptionsProperty := &SheetControlInfoIconLabelOptionsProperty{
//   	InfoIconText: jsii.String("infoIconText"),
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetcontrolinfoiconlabeloptions.html
//
type CfnAnalysisPropsMixin_SheetControlInfoIconLabelOptionsProperty struct {
	// The text content of info icon.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetcontrolinfoiconlabeloptions.html#cfn-quicksight-analysis-sheetcontrolinfoiconlabeloptions-infoicontext
	//
	InfoIconText *string `field:"optional" json:"infoIconText" yaml:"infoIconText"`
	// The visibility configuration of info icon label options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetcontrolinfoiconlabeloptions.html#cfn-quicksight-analysis-sheetcontrolinfoiconlabeloptions-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

