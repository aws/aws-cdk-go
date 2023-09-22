package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a SingleValueWidget.
//
// Example:
//   var dashboard dashboard
//
//
//   dashboard.AddWidgets(cloudwatch.NewSingleValueWidget(&SingleValueWidgetProps{
//   	Metrics: []iMetric{
//   	},
//
//   	Period: awscdk.Duration_Minutes(jsii.Number(15)),
//   }))
//
type SingleValueWidgetProps struct {
	// Height of the widget.
	// Default: - 6 for Alarm and Graph widgets.
	// 3 for single value widgets where most recent value of a metric is displayed.
	//
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// The region the metrics of this graph should be taken from.
	// Default: - Current region.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Title for the graph.
	// Default: - None.
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
	// Width of the widget, in a grid of 24 units wide.
	// Default: 6.
	//
	Width *float64 `field:"optional" json:"width" yaml:"width"`
	// Metrics to display.
	Metrics *[]IMetric `field:"required" json:"metrics" yaml:"metrics"`
	// The end of the time range to use for each widget independently from those of the dashboard.
	//
	// If you specify a value for end, you must also specify a value for start.
	// Specify an absolute time in the ISO 8601 format. For example, 2018-12-17T06:00:00.000Z.
	// Default: When the dashboard loads, the end date will be the current time.
	//
	End *string `field:"optional" json:"end" yaml:"end"`
	// Whether to show as many digits as can fit, before rounding.
	// Default: false.
	//
	FullPrecision *bool `field:"optional" json:"fullPrecision" yaml:"fullPrecision"`
	// The default period for all metrics in this widget.
	//
	// The period is the length of time represented by one data point on the graph.
	// This default can be overridden within each metric definition.
	// Default: cdk.Duration.seconds(300)
	//
	Period awscdk.Duration `field:"optional" json:"period" yaml:"period"`
	// Whether to show the value from the entire time range.
	// Default: false.
	//
	SetPeriodToTimeRange *bool `field:"optional" json:"setPeriodToTimeRange" yaml:"setPeriodToTimeRange"`
	// Whether to show a graph below the value illustrating the value for the whole time range.
	//
	// Cannot be used in combination with `setPeriodToTimeRange`.
	// Default: false.
	//
	Sparkline *bool `field:"optional" json:"sparkline" yaml:"sparkline"`
	// The start of the time range to use for each widget independently from those of the dashboard.
	//
	// You can specify start without specifying end to specify a relative time range that ends with the current time.
	// In this case, the value of start must begin with -P, and you can use M, H, D, W and M as abbreviations for
	// minutes, hours, days, weeks and months. For example, -PT8H shows the last 8 hours and -P3M shows the last three months.
	// You can also use start along with an end field, to specify an absolute time range.
	// When specifying an absolute time range, use the ISO 8601 format. For example, 2018-12-17T06:00:00.000Z.
	// Default: When the dashboard loads, the start time will be the default time range.
	//
	Start *string `field:"optional" json:"start" yaml:"start"`
}

