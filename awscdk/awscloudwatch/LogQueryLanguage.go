package awscloudwatch


// Logs Query Language.
//
// Example:
//   var dashboard Dashboard
//
//
//   dashboard.AddWidgets(cloudwatch.NewLogQueryWidget(&LogQueryWidgetProps{
//   	LogGroupNames: []*string{
//   		jsii.String("my-log-group"),
//   	},
//   	View: cloudwatch.LogQueryVisualizationType_TABLE,
//   	QueryString: jsii.String("SELECT count(*) as count FROM 'my-log-group'"),
//   	QueryLanguage: cloudwatch.LogQueryLanguage_SQL,
//   }))
//
type LogQueryLanguage string

const (
	// Logs Insights QL.
	LogQueryLanguage_LOGS_INSIGHTS LogQueryLanguage = "LOGS_INSIGHTS"
	// OpenSearch SQL.
	LogQueryLanguage_SQL LogQueryLanguage = "SQL"
	// OpenSearch PPL.
	LogQueryLanguage_PPL LogQueryLanguage = "PPL"
)

