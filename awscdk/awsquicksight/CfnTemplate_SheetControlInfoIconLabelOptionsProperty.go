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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetcontrolinfoiconlabeloptions.html
//
type CfnTemplate_SheetControlInfoIconLabelOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetcontrolinfoiconlabeloptions.html#cfn-quicksight-template-sheetcontrolinfoiconlabeloptions-infoicontext
	//
	InfoIconText *string `field:"optional" json:"infoIconText" yaml:"infoIconText"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetcontrolinfoiconlabeloptions.html#cfn-quicksight-template-sheetcontrolinfoiconlabeloptions-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

