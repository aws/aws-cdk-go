package previewawsquicksightmixins


// The tooltip configuration for a sheet image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sheetImageTooltipConfigurationProperty := &SheetImageTooltipConfigurationProperty{
//   	TooltipText: &SheetImageTooltipTextProperty{
//   		PlainText: jsii.String("plainText"),
//   	},
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetimagetooltipconfiguration.html
//
type CfnDashboardPropsMixin_SheetImageTooltipConfigurationProperty struct {
	// The text that appears in the tooltip.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetimagetooltipconfiguration.html#cfn-quicksight-dashboard-sheetimagetooltipconfiguration-tooltiptext
	//
	TooltipText interface{} `field:"optional" json:"tooltipText" yaml:"tooltipText"`
	// The visibility of the tooltip.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetimagetooltipconfiguration.html#cfn-quicksight-dashboard-sheetimagetooltipconfiguration-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

