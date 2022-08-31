package awscloudwatch


// Types of view.
//
// Example:
//   var dashboard dashboard
//
//
//   dashboard.addWidgets(cloudwatch.NewGraphWidget(&graphWidgetProps{
//   	// ...
//
//   	view: cloudwatch.graphWidgetView_BAR,
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

