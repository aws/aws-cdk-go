package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a GraphWidget.
//
// Example:
//   var dashboard dashboard
//
//
//   dashboard.AddWidgets(cloudwatch.NewGraphWidget(&GraphWidgetProps{
//   	// ...
//
//   	LegendPosition: cloudwatch.LegendPosition_RIGHT,
//   }))
//
type GraphWidgetProps struct {
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
	// The end of the time range to use for each widget independently from those of the dashboard.
	//
	// If you specify a value for end, you must also specify a value for start.
	// Specify an absolute time in the ISO 8601 format. For example, 2018-12-17T06:00:00.000Z.
	// Default: When the dashboard loads, the end date will be the current time.
	//
	End *string `field:"optional" json:"end" yaml:"end"`
	// Metrics to display on left Y axis.
	// Default: - No metrics.
	//
	Left *[]IMetric `field:"optional" json:"left" yaml:"left"`
	// Annotations for the left Y axis.
	// Default: - No annotations.
	//
	LeftAnnotations *[]*HorizontalAnnotation `field:"optional" json:"leftAnnotations" yaml:"leftAnnotations"`
	// Left Y axis.
	// Default: - None.
	//
	LeftYAxis *YAxisProps `field:"optional" json:"leftYAxis" yaml:"leftYAxis"`
	// Position of the legend.
	// Default: - bottom.
	//
	LegendPosition LegendPosition `field:"optional" json:"legendPosition" yaml:"legendPosition"`
	// Whether the graph should show live data.
	// Default: false.
	//
	LiveData *bool `field:"optional" json:"liveData" yaml:"liveData"`
	// The default period for all metrics in this widget.
	//
	// The period is the length of time represented by one data point on the graph.
	// This default can be overridden within each metric definition.
	// Default: cdk.Duration.seconds(300)
	//
	Period awscdk.Duration `field:"optional" json:"period" yaml:"period"`
	// Metrics to display on right Y axis.
	// Default: - No metrics.
	//
	Right *[]IMetric `field:"optional" json:"right" yaml:"right"`
	// Annotations for the right Y axis.
	// Default: - No annotations.
	//
	RightAnnotations *[]*HorizontalAnnotation `field:"optional" json:"rightAnnotations" yaml:"rightAnnotations"`
	// Right Y axis.
	// Default: - None.
	//
	RightYAxis *YAxisProps `field:"optional" json:"rightYAxis" yaml:"rightYAxis"`
	// Whether to show the value from the entire time range. Only applicable for Bar and Pie charts.
	//
	// If false, values will be from the most recent period of your chosen time range;
	// if true, shows the value from the entire time range.
	// Default: false.
	//
	SetPeriodToTimeRange *bool `field:"optional" json:"setPeriodToTimeRange" yaml:"setPeriodToTimeRange"`
	// Whether the graph should be shown as stacked lines.
	// Default: false.
	//
	Stacked *bool `field:"optional" json:"stacked" yaml:"stacked"`
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
	// The default statistic to be displayed for each metric.
	//
	// This default can be overridden within the definition of each individual metric.
	// Default: - The statistic for each metric is used.
	//
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
	// Annotations for the X axis.
	// Default: - No annotations.
	//
	VerticalAnnotations *[]*VerticalAnnotation `field:"optional" json:"verticalAnnotations" yaml:"verticalAnnotations"`
	// Display this metric.
	// Default: TimeSeries.
	//
	View GraphWidgetView `field:"optional" json:"view" yaml:"view"`
}

