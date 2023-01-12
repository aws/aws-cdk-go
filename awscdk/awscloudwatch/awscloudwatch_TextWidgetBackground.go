package awscloudwatch


// Background types available.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var dashboard dashboard
//
//
//   dashboard.addWidgets(cloudwatch.NewTextWidget(&textWidgetProps{
//   	markdown: jsii.String("# Key Performance Indicators"),
//   	background: textWidgetBackground_TRANSPARENT,
//   }))
//
type TextWidgetBackground string

const (
	// Solid background.
	TextWidgetBackground_SOLID TextWidgetBackground = "SOLID"
	// Transparent background.
	TextWidgetBackground_TRANSPARENT TextWidgetBackground = "TRANSPARENT"
)

