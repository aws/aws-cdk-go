package awsquicksight


// The options that determine the visibility, color, type, and tooltip visibility of the sparkline of a KPI visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kPISparklineOptionsProperty := &KPISparklineOptionsProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Color: jsii.String("color"),
//   	TooltipVisibility: jsii.String("tooltipVisibility"),
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-kpisparklineoptions.html
//
type CfnAnalysis_KPISparklineOptionsProperty struct {
	// The type of the sparkline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-kpisparklineoptions.html#cfn-quicksight-analysis-kpisparklineoptions-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The color of the sparkline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-kpisparklineoptions.html#cfn-quicksight-analysis-kpisparklineoptions-color
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
	// The tooltip visibility of the sparkline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-kpisparklineoptions.html#cfn-quicksight-analysis-kpisparklineoptions-tooltipvisibility
	//
	TooltipVisibility *string `field:"optional" json:"tooltipVisibility" yaml:"tooltipVisibility"`
	// The visibility of the sparkline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-kpisparklineoptions.html#cfn-quicksight-analysis-kpisparklineoptions-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

