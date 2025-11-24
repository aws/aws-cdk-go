package mixinsawsquicksight


// The text that appears in the sheet image tooltip.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sheetImageTooltipTextProperty := &SheetImageTooltipTextProperty{
//   	PlainText: jsii.String("plainText"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetimagetooltiptext.html
//
type CfnDashboardPropsMixin_SheetImageTooltipTextProperty struct {
	// The plain text format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetimagetooltiptext.html#cfn-quicksight-dashboard-sheetimagetooltiptext-plaintext
	//
	PlainText *string `field:"optional" json:"plainText" yaml:"plainText"`
}

