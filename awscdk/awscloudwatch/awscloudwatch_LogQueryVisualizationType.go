package awscloudwatch


// Types of view.
//
// Example:
//   var dashboard dashboard
//
//
//   dashboard.AddWidgets(cloudwatch.NewLogQueryWidget(&LogQueryWidgetProps{
//   	LogGroupNames: []*string{
//   		jsii.String("my-log-group"),
//   	},
//   	View: cloudwatch.LogQueryVisualizationType_TABLE,
//   	// The lines will be automatically combined using '\n|'.
//   	QueryLines: []*string{
//   		jsii.String("fields @message"),
//   		jsii.String("filter @message like /Error/"),
//   	},
//   }))
//
// Experimental.
type LogQueryVisualizationType string

const (
	// Table view.
	// Experimental.
	LogQueryVisualizationType_TABLE LogQueryVisualizationType = "TABLE"
	// Line view.
	// Experimental.
	LogQueryVisualizationType_LINE LogQueryVisualizationType = "LINE"
	// Stacked area view.
	// Experimental.
	LogQueryVisualizationType_STACKEDAREA LogQueryVisualizationType = "STACKEDAREA"
	// Bar view.
	// Experimental.
	LogQueryVisualizationType_BAR LogQueryVisualizationType = "BAR"
	// Pie view.
	// Experimental.
	LogQueryVisualizationType_PIE LogQueryVisualizationType = "PIE"
)

