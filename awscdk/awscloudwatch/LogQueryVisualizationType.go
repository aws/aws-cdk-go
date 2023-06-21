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
type LogQueryVisualizationType string

const (
	// Table view.
	LogQueryVisualizationType_TABLE LogQueryVisualizationType = "TABLE"
	// Line view.
	LogQueryVisualizationType_LINE LogQueryVisualizationType = "LINE"
	// Stacked area view.
	LogQueryVisualizationType_STACKEDAREA LogQueryVisualizationType = "STACKEDAREA"
	// Bar view.
	LogQueryVisualizationType_BAR LogQueryVisualizationType = "BAR"
	// Pie view.
	LogQueryVisualizationType_PIE LogQueryVisualizationType = "PIE"
)

