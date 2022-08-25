package awscloudwatch


// Basic properties for widgets that display metrics.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricWidgetProps := &metricWidgetProps{
//   	height: jsii.Number(123),
//   	region: jsii.String("region"),
//   	title: jsii.String("title"),
//   	width: jsii.Number(123),
//   }
//
type MetricWidgetProps struct {
	// Height of the widget.
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// The region the metrics of this graph should be taken from.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Title for the graph.
	Title *string `field:"optional" json:"title" yaml:"title"`
	// Width of the widget, in a grid of 24 units wide.
	Width *float64 `field:"optional" json:"width" yaml:"width"`
}

