package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   waterfallChartGroupColorConfigurationProperty := &WaterfallChartGroupColorConfigurationProperty{
//   	NegativeBarColor: jsii.String("negativeBarColor"),
//   	PositiveBarColor: jsii.String("positiveBarColor"),
//   	TotalBarColor: jsii.String("totalBarColor"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-waterfallchartgroupcolorconfiguration.html
//
type CfnAnalysis_WaterfallChartGroupColorConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-waterfallchartgroupcolorconfiguration.html#cfn-quicksight-analysis-waterfallchartgroupcolorconfiguration-negativebarcolor
	//
	NegativeBarColor *string `field:"optional" json:"negativeBarColor" yaml:"negativeBarColor"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-waterfallchartgroupcolorconfiguration.html#cfn-quicksight-analysis-waterfallchartgroupcolorconfiguration-positivebarcolor
	//
	PositiveBarColor *string `field:"optional" json:"positiveBarColor" yaml:"positiveBarColor"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-waterfallchartgroupcolorconfiguration.html#cfn-quicksight-analysis-waterfallchartgroupcolorconfiguration-totalbarcolor
	//
	TotalBarColor *string `field:"optional" json:"totalBarColor" yaml:"totalBarColor"`
}

