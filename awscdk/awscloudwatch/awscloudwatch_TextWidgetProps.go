package awscloudwatch


// Properties for a Text widget.
//
// Example:
//   var dashboard dashboard
//
//
//   dashboard.addWidgets(cloudwatch.NewTextWidget(&textWidgetProps{
//   	markdown: jsii.String("# Key Performance Indicators"),
//   }))
//
type TextWidgetProps struct {
	// The text to display, in MarkDown format.
	Markdown *string `field:"required" json:"markdown" yaml:"markdown"`
	// Background for the widget.
	Background TextWidgetBackground `field:"optional" json:"background" yaml:"background"`
	// Height of the widget.
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// Width of the widget, in a grid of 24 units wide.
	Width *float64 `field:"optional" json:"width" yaml:"width"`
}

