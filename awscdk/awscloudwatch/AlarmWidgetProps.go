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
// Experimental.
type AlarmWidgetProps struct {
	// Height of the widget.
	// Experimental.
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// The region the metrics of this graph should be taken from.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Title for the graph.
	// Experimental.
	Title *string `field:"optional" json:"title" yaml:"title"`
	// Width of the widget, in a grid of 24 units wide.
	// Experimental.
	Width *float64 `field:"optional" json:"width" yaml:"width"`
	// The alarm to show.
	// Experimental.
	Alarm IAlarm `field:"required" json:"alarm" yaml:"alarm"`
	// Left Y axis.
	// Experimental.
	LeftYAxis *YAxisProps `field:"optional" json:"leftYAxis" yaml:"leftYAxis"`
}

