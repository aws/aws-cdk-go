package previewawsquicksightmixins


// The color configuration of a `GaugeChartVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   gaugeChartColorConfigurationProperty := &GaugeChartColorConfigurationProperty{
//   	BackgroundColor: jsii.String("backgroundColor"),
//   	ForegroundColor: jsii.String("foregroundColor"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gaugechartcolorconfiguration.html
//
type CfnTemplatePropsMixin_GaugeChartColorConfigurationProperty struct {
	// The background color configuration of a `GaugeChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gaugechartcolorconfiguration.html#cfn-quicksight-template-gaugechartcolorconfiguration-backgroundcolor
	//
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// The foreground color configuration of a `GaugeChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gaugechartcolorconfiguration.html#cfn-quicksight-template-gaugechartcolorconfiguration-foregroundcolor
	//
	ForegroundColor *string `field:"optional" json:"foregroundColor" yaml:"foregroundColor"`
}

