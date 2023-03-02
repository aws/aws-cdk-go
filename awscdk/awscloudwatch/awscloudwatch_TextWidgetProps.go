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
// Experimental.
type TextWidgetProps struct {
	// The text to display, in MarkDown format.
	// Experimental.
	Markdown *string `field:"required" json:"markdown" yaml:"markdown"`
	// Height of the widget.
	// Experimental.
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// Width of the widget, in a grid of 24 units wide.
	// Experimental.
	Width *float64 `field:"optional" json:"width" yaml:"width"`
}

