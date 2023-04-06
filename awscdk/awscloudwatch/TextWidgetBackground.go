package awscloudwatch


// Background types available.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var dashboard dashboard
//
//
//   dashboard.AddWidgets(cloudwatch.NewTextWidget(&TextWidgetProps{
//   	Markdown: jsii.String("# Key Performance Indicators"),
//   	Background: textWidgetBackground_TRANSPARENT,
//   }))
//
type TextWidgetBackground string

const (
	// Solid background.
	TextWidgetBackground_SOLID TextWidgetBackground = "SOLID"
	// Transparent background.
	TextWidgetBackground_TRANSPARENT TextWidgetBackground = "TRANSPARENT"
)

