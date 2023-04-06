package awsquicksight


// The options for configuring a donut chart or pie chart.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   donutOptionsProperty := &DonutOptionsProperty{
//   	ArcOptions: &ArcOptionsProperty{
//   		ArcThickness: jsii.String("arcThickness"),
//   	},
//   	DonutCenterOptions: &DonutCenterOptionsProperty{
//   		LabelVisibility: jsii.String("labelVisibility"),
//   	},
//   }
//
type CfnDashboard_DonutOptionsProperty struct {
	// The option for define the arc of the chart shape. Valid values are as follows:.
	//
	// - `WHOLE` - A pie chart
	// - `SMALL` - A small-sized donut chart
	// - `MEDIUM` - A medium-sized donut chart
	// - `LARGE` - A large-sized donut chart.
	ArcOptions interface{} `field:"optional" json:"arcOptions" yaml:"arcOptions"`
	// The label options of the label that is displayed in the center of a donut chart.
	//
	// This option isn't available for pie charts.
	DonutCenterOptions interface{} `field:"optional" json:"donutCenterOptions" yaml:"donutCenterOptions"`
}

