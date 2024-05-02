package awsquicksight


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-waterfallchartcolorconfiguration.html
//
type CfnDashboard_WaterfallChartColorConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-waterfallchartcolorconfiguration.html#cfn-quicksight-dashboard-waterfallchartcolorconfiguration-groupcolorconfiguration
	//
	GroupColorConfiguration interface{} `field:"optional" json:"groupColorConfiguration" yaml:"groupColorConfiguration"`
}

