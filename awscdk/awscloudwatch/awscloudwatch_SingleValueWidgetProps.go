package awscloudwatch


// Properties for a SingleValueWidget.
//
// Example:
//   var dashboard dashboard
//   var visitorCount metric
//   var purchaseCount metric
//
//
//   dashboard.addWidgets(cloudwatch.NewSingleValueWidget(&singleValueWidgetProps{
//   	metrics: []iMetric{
//   		visitorCount,
//   		purchaseCount,
//   	},
//   }))
//
// Experimental.
type SingleValueWidgetProps struct {
	// Height of the widget.
	// Experimental.
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// The region the metrics of this graph should be taken from.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Title for the graph.
	// Experimental.
	Title *string `field:"optional" json:"title" yaml:"title"`
	// Width of the widget, in a grid of 24 units wide.
	// Experimental.
	Width *float64 `field:"optional" json:"width" yaml:"width"`
	// Metrics to display.
	// Experimental.
	Metrics *[]IMetric `field:"required" json:"metrics" yaml:"metrics"`
	// Whether to show as many digits as can fit, before rounding.
	// Experimental.
	FullPrecision *bool `field:"optional" json:"fullPrecision" yaml:"fullPrecision"`
	// Whether to show the value from the entire time range.
	// Experimental.
	SetPeriodToTimeRange *bool `field:"optional" json:"setPeriodToTimeRange" yaml:"setPeriodToTimeRange"`
}

