package previewawsquicksightmixins


// The options that determine the visibility, color, type, and tooltip visibility of the sparkline of a KPI visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kPISparklineOptionsProperty := &KPISparklineOptionsProperty{
//   	Color: jsii.String("color"),
//   	TooltipVisibility: jsii.String("tooltipVisibility"),
//   	Type: jsii.String("type"),
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpisparklineoptions.html
//
type CfnTemplatePropsMixin_KPISparklineOptionsProperty struct {
	// The color of the sparkline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpisparklineoptions.html#cfn-quicksight-template-kpisparklineoptions-color
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
	// The tooltip visibility of the sparkline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpisparklineoptions.html#cfn-quicksight-template-kpisparklineoptions-tooltipvisibility
	//
	TooltipVisibility *string `field:"optional" json:"tooltipVisibility" yaml:"tooltipVisibility"`
	// The type of the sparkline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpisparklineoptions.html#cfn-quicksight-template-kpisparklineoptions-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The visibility of the sparkline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpisparklineoptions.html#cfn-quicksight-template-kpisparklineoptions-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

