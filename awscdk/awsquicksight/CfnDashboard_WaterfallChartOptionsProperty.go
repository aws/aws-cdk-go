package awsquicksight


// The options that determine the presentation of a waterfall visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   waterfallChartOptionsProperty := &WaterfallChartOptionsProperty{
//   	TotalBarLabel: jsii.String("totalBarLabel"),
//   }
//
type CfnDashboard_WaterfallChartOptionsProperty struct {
	// This option determines the total bar label of a waterfall visual.
	TotalBarLabel *string `field:"optional" json:"totalBarLabel" yaml:"totalBarLabel"`
}

