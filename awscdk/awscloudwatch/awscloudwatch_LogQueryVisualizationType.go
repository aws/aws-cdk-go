package awscloudwatch


// Types of view.
//
// Example:
//   var dashboard dashboard
//
//
//   dashboard.addWidgets(cloudwatch.NewLogQueryWidget(&logQueryWidgetProps{
//   	logGroupNames: []*string{
//   		jsii.String("my-log-group"),
//   	},
//   	view: cloudwatch.logQueryVisualizationType_TABLE,
//   	// The lines will be automatically combined using '\n|'.
//   	queryLines: []*string{
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

