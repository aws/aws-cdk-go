package awsquicksight


// The series settings of a radar chart.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   radarChartSeriesSettingsProperty := &RadarChartSeriesSettingsProperty{
//   	AreaStyleSettings: &RadarChartAreaStyleSettingsProperty{
//   		Visibility: jsii.String("visibility"),
//   	},
//   }
//
type CfnDashboard_RadarChartSeriesSettingsProperty struct {
	// The area style settings of a radar chart.
	AreaStyleSettings interface{} `field:"optional" json:"areaStyleSettings" yaml:"areaStyleSettings"`
}

