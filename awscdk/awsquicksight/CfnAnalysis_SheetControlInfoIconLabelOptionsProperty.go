package awsquicksight


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetcontrolinfoiconlabeloptions.html
//
type CfnAnalysis_SheetControlInfoIconLabelOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetcontrolinfoiconlabeloptions.html#cfn-quicksight-analysis-sheetcontrolinfoiconlabeloptions-infoicontext
	//
	InfoIconText *string `field:"optional" json:"infoIconText" yaml:"infoIconText"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetcontrolinfoiconlabeloptions.html#cfn-quicksight-analysis-sheetcontrolinfoiconlabeloptions-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

