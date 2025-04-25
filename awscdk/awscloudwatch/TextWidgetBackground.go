package awscloudwatch


// Background types available.
//
// Example:
//   var dashboard dashboard
//
//
//   dashboard.AddWidgets(cloudwatch.NewTextWidget(&TextWidgetProps{
//   	Markdown: jsii.String("# Key Performance Indicators"),
//   	Background: cloudwatch.TextWidgetBackground_TRANSPARENT,
//   }))
//
type TextWidgetBackground string

const (
	// Solid background.
	TextWidgetBackground_SOLID TextWidgetBackground = "SOLID"
	// Transparent background.
	TextWidgetBackground_TRANSPARENT TextWidgetBackground = "TRANSPARENT"
)

