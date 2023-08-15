package awscloudwatch


// Properties for a Text widget.
//
// Example:
//   var dashboard dashboard
//
//
//   dashboard.AddWidgets(cloudwatch.NewTextWidget(&TextWidgetProps{
//   	Markdown: jsii.String("# Key Performance Indicators"),
//   }))
//
type TextWidgetProps struct {
	// The text to display, in MarkDown format.
	Markdown *string `field:"required" json:"markdown" yaml:"markdown"`
	// Background for the widget.
	// Default: solid.
	//
	Background TextWidgetBackground `field:"optional" json:"background" yaml:"background"`
	// Height of the widget.
	// Default: 2.
	//
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// Width of the widget, in a grid of 24 units wide.
	// Default: 6.
	//
	Width *float64 `field:"optional" json:"width" yaml:"width"`
}

