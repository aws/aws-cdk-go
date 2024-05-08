package awsquicksight


// The color configuration of a waterfall visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   waterfallChartColorConfigurationProperty := &WaterfallChartColorConfigurationProperty{
//   	GroupColorConfiguration: &WaterfallChartGroupColorConfigurationProperty{
//   		NegativeBarColor: jsii.String("negativeBarColor"),
//   		PositiveBarColor: jsii.String("positiveBarColor"),
//   		TotalBarColor: jsii.String("totalBarColor"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-waterfallchartcolorconfiguration.html
//
type CfnTemplate_WaterfallChartColorConfigurationProperty struct {
	// The color configuration for individual groups within a waterfall visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-waterfallchartcolorconfiguration.html#cfn-quicksight-template-waterfallchartcolorconfiguration-groupcolorconfiguration
	//
	GroupColorConfiguration interface{} `field:"optional" json:"groupColorConfiguration" yaml:"groupColorConfiguration"`
}

