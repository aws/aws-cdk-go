package awscloudwatch


// Properties for an AlarmWidget.
//
// Example:
//   var dashboard dashboard
//   var errorAlarm alarm
//
//
//   dashboard.addWidgets(cloudwatch.NewAlarmWidget(&alarmWidgetProps{
//   	title: jsii.String("Errors"),
//   	alarm: errorAlarm,
//   }))
//
type AlarmWidgetProps struct {
	// Height of the widget.
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// The region the metrics of this graph should be taken from.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Title for the graph.
	Title *string `field:"optional" json:"title" yaml:"title"`
	// Width of the widget, in a grid of 24 units wide.
	Width *float64 `field:"optional" json:"width" yaml:"width"`
	// The alarm to show.
	Alarm IAlarm `field:"required" json:"alarm" yaml:"alarm"`
	// Left Y axis.
	LeftYAxis *YAxisProps `field:"optional" json:"leftYAxis" yaml:"leftYAxis"`
}

