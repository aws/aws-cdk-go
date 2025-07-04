package awscloudwatch


// Properties for an AlarmWidget.
//
// Example:
//   var dashboard dashboard
//   var errorAlarm alarm
//
//
//   dashboard.AddWidgets(cloudwatch.NewAlarmWidget(&AlarmWidgetProps{
//   	Title: jsii.String("Errors"),
//   	Alarm: errorAlarm,
//   }))
//
type AlarmWidgetProps struct {
	// The AWS account ID where the metrics are located.
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
	// Default: - 6 for Alarm and Graph widgets.
	// 3 for single value widgets where most recent value of a metric is displayed.
	//
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// The region the metrics of this graph should be taken from.
	// Default: - Current region.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Title for the graph.
	// Default: - None.
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
	// Width of the widget, in a grid of 24 units wide.
	// Default: 6.
	//
	Width *float64 `field:"optional" json:"width" yaml:"width"`
	// The alarm to show.
	Alarm IAlarm `field:"required" json:"alarm" yaml:"alarm"`
	// Left Y axis.
	// Default: - No minimum or maximum values for the left Y-axis.
	//
	LeftYAxis *YAxisProps `field:"optional" json:"leftYAxis" yaml:"leftYAxis"`
}

