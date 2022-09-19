package awscloudwatch


// Properties for a SingleValueWidget.
//
// Example:
//   var dashboard dashboard
//
//
//   dashboard.addWidgets(cloudwatch.NewSingleValueWidget(&singleValueWidgetProps{
//   	metrics: []iMetric{
//   	},
//
//   	fullPrecision: jsii.Boolean(true),
//   }))
//
type SingleValueWidgetProps struct {
	// Height of the widget.
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// The region the metrics of this graph should be taken from.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Title for the graph.
	Title *string `field:"optional" json:"title" yaml:"title"`
	// Width of the widget, in a grid of 24 units wide.
	Width *float64 `field:"optional" json:"width" yaml:"width"`
	// Metrics to display.
	Metrics *[]IMetric `field:"required" json:"metrics" yaml:"metrics"`
	// Whether to show as many digits as can fit, before rounding.
	FullPrecision *bool `field:"optional" json:"fullPrecision" yaml:"fullPrecision"`
	// Whether to show the value from the entire time range.
	SetPeriodToTimeRange *bool `field:"optional" json:"setPeriodToTimeRange" yaml:"setPeriodToTimeRange"`
	// Whether to show a graph below the value illustrating the value for the whole time range.
	//
	// Cannot be used in combination with `setPeriodToTimeRange`.
	Sparkline *bool `field:"optional" json:"sparkline" yaml:"sparkline"`
}

