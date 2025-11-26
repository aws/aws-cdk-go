package previewawsquicksightmixins


// The tooltip configuration for a sheet image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var visibility interface{}
//
//   sheetImageTooltipConfigurationProperty := &SheetImageTooltipConfigurationProperty{
//   	TooltipText: &SheetImageTooltipTextProperty{
//   		PlainText: jsii.String("plainText"),
//   	},
//   	Visibility: visibility,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetimagetooltipconfiguration.html
//
type CfnTemplatePropsMixin_SheetImageTooltipConfigurationProperty struct {
	// The text that appears in the tooltip.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetimagetooltipconfiguration.html#cfn-quicksight-template-sheetimagetooltipconfiguration-tooltiptext
	//
	TooltipText interface{} `field:"optional" json:"tooltipText" yaml:"tooltipText"`
	// The visibility of the tooltip.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetimagetooltipconfiguration.html#cfn-quicksight-template-sheetimagetooltipconfiguration-visibility
	//
	Visibility interface{} `field:"optional" json:"visibility" yaml:"visibility"`
}

