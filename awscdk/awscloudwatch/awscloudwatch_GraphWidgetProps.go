package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk"
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
// Experimental.
type GraphWidgetProps struct {
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
	// Metrics to display on left Y axis.
	// Experimental.
	Left *[]IMetric `field:"optional" json:"left" yaml:"left"`
	// Annotations for the left Y axis.
	// Experimental.
	LeftAnnotations *[]*HorizontalAnnotation `field:"optional" json:"leftAnnotations" yaml:"leftAnnotations"`
	// Left Y axis.
	// Experimental.
	LeftYAxis *YAxisProps `field:"optional" json:"leftYAxis" yaml:"leftYAxis"`
	// Position of the legend.
	// Experimental.
	LegendPosition LegendPosition `field:"optional" json:"legendPosition" yaml:"legendPosition"`
	// Whether the graph should show live data.
	// Experimental.
	LiveData *bool `field:"optional" json:"liveData" yaml:"liveData"`
	// The default period for all metrics in this widget.
	//
	// The period is the length of time represented by one data point on the graph.
	// This default can be overridden within each metric definition.
	// Experimental.
	Period awscdk.Duration `field:"optional" json:"period" yaml:"period"`
	// Metrics to display on right Y axis.
	// Experimental.
	Right *[]IMetric `field:"optional" json:"right" yaml:"right"`
	// Annotations for the right Y axis.
	// Experimental.
	RightAnnotations *[]*HorizontalAnnotation `field:"optional" json:"rightAnnotations" yaml:"rightAnnotations"`
	// Right Y axis.
	// Experimental.
	RightYAxis *YAxisProps `field:"optional" json:"rightYAxis" yaml:"rightYAxis"`
	// Whether to show the value from the entire time range. Only applicable for Bar and Pie charts.
	//
	// If false, values will be from the most recent period of your chosen time range;
	// if true, shows the value from the entire time range.
	// Experimental.
	SetPeriodToTimeRange *bool `field:"optional" json:"setPeriodToTimeRange" yaml:"setPeriodToTimeRange"`
	// Whether the graph should be shown as stacked lines.
	// Experimental.
	Stacked *bool `field:"optional" json:"stacked" yaml:"stacked"`
	// The default statistic to be displayed for each metric.
	//
	// This default can be overridden within the definition of each individual metric.
	// Experimental.
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
	// Display this metric.
	// Experimental.
	View GraphWidgetView `field:"optional" json:"view" yaml:"view"`
}

