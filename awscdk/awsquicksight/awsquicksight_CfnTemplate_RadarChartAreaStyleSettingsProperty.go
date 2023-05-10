package awsquicksight


// The configured style settings of a radar chart.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   radarChartAreaStyleSettingsProperty := &RadarChartAreaStyleSettingsProperty{
//   	Visibility: jsii.String("visibility"),
//   }
//
type CfnTemplate_RadarChartAreaStyleSettingsProperty struct {
	// The visibility settings of a radar chart.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

