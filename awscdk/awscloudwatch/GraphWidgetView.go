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
type GraphWidgetView string

const (
	// Display as a line graph.
	GraphWidgetView_TIME_SERIES GraphWidgetView = "TIME_SERIES"
	// Display as a bar graph.
	GraphWidgetView_BAR GraphWidgetView = "BAR"
	// Display as a pie graph.
	GraphWidgetView_PIE GraphWidgetView = "PIE"
)

