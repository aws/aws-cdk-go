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
	// Height of the widget.
	// Default: 6.
	//
	Height *float64 `field:"optional" json:"height" yaml:"height"`
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

