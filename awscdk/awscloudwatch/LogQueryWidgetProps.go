package awscloudwatch


// Properties for a Query widget.
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
type LogQueryWidgetProps struct {
	// Names of log groups to query.
	LogGroupNames *[]*string `field:"required" json:"logGroupNames" yaml:"logGroupNames"`
	// The AWS account ID where the log groups are located.
	//
	// This enables cross-account functionality for CloudWatch dashboards.
	// Before using this feature, ensure that proper cross-account sharing is configured
	// between the monitoring account and source account.
	//
	// For more information, see:
	// https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html
	// Default: - Current account.
	//
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Height of the widget.
	// Default: 6.
	//
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// The query language to use for the query.
	// Default: LogQueryLanguage.LOGS_INSIGHTS
	//
	QueryLanguage LogQueryLanguage `field:"optional" json:"queryLanguage" yaml:"queryLanguage"`
	// A sequence of lines to use to build the query.
	//
	// The query will be built by joining the lines together using `\n|`.
	// Default: - Exactly one of `queryString`, `queryLines` is required.
	//
	QueryLines *[]*string `field:"optional" json:"queryLines" yaml:"queryLines"`
	// Full query string for log insights.
	//
	// Be sure to prepend every new line with a newline and pipe character
	// (`\n|`).
	// Default: - Exactly one of `queryString`, `queryLines` is required.
	//
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
	// The region the metrics of this widget should be taken from.
	// Default: Current region.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Title for the widget.
	// Default: No title.
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The type of view to use.
	// Default: LogQueryVisualizationType.TABLE
	//
	View LogQueryVisualizationType `field:"optional" json:"view" yaml:"view"`
	// Width of the widget, in a grid of 24 units wide.
	// Default: 6.
	//
	Width *float64 `field:"optional" json:"width" yaml:"width"`
}

