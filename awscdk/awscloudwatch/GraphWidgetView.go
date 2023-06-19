package awscloudwatch


// Types of view.
//
// Example:
//   var dashboard dashboard
//
//
//   dashboard.AddWidgets(cloudwatch.NewGraphWidget(&GraphWidgetProps{
//   	// ...
//
//   	View: cloudwatch.GraphWidgetView_BAR,
//   }))
//
// Experimental.
type GraphWidgetView string

const (
	// Display as a line graph.
	// Experimental.
	GraphWidgetView_TIME_SERIES GraphWidgetView = "TIME_SERIES"
	// Display as a bar graph.
	// Experimental.
	GraphWidgetView_BAR GraphWidgetView = "BAR"
	// Display as a pie graph.
	// Experimental.
	GraphWidgetView_PIE GraphWidgetView = "PIE"
)

