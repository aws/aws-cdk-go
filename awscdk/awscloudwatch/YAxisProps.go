package awscloudwatch


// Properties for a Y-Axis.
//
// Example:
//   var dashboard dashboard
//   var errorAlarm alarm
//   var gaugeMetric metric
//
//
//   dashboard.AddWidgets(cloudwatch.NewGaugeWidget(&GaugeWidgetProps{
//   	Metrics: []iMetric{
//   		gaugeMetric,
//   	},
//   	LeftYAxis: &YAxisProps{
//   		Min: jsii.Number(0),
//   		Max: jsii.Number(1000),
//   	},
//   }))
//
type YAxisProps struct {
	// The label.
	// Default: - No label.
	//
	Label *string `field:"optional" json:"label" yaml:"label"`
	// The max value.
	// Default: - No maximum value.
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The min value.
	// Default: 0.
	//
	Min *float64 `field:"optional" json:"min" yaml:"min"`
	// Whether to show units.
	// Default: true.
	//
	ShowUnits *bool `field:"optional" json:"showUnits" yaml:"showUnits"`
}

