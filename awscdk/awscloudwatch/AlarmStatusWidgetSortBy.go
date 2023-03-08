package awscloudwatch


// The sort possibilities for AlarmStatusWidgets.
//
// Example:
//   var dashboard dashboard
//   var errorAlarm alarm
//
//
//   dashboard.AddWidgets(cloudwatch.NewAlarmStatusWidget(&AlarmStatusWidgetProps{
//   	Title: jsii.String("Errors"),
//   	Alarms: []iAlarm{
//   		errorAlarm,
//   	},
//   	SortBy: cloudwatch.AlarmStatusWidgetSortBy_STATE_UPDATED_TIMESTAMP,
//   	States: []alarmState{
//   		cloudwatch.*alarmState_ALARM,
//   	},
//   }))
//
type AlarmStatusWidgetSortBy string

const (
	// Choose DEFAULT to sort them in alphabetical order by alarm name.
	AlarmStatusWidgetSortBy_DEFAULT AlarmStatusWidgetSortBy = "DEFAULT"
	// Choose STATE_UPDATED_TIMESTAMP to sort them first by alarm state, with alarms in ALARM state first, INSUFFICIENT_DATA alarms next, and OK alarms last.
	//
	// Within each group, the alarms are sorted by when they last changed state, with more recent state changes listed first.
	AlarmStatusWidgetSortBy_STATE_UPDATED_TIMESTAMP AlarmStatusWidgetSortBy = "STATE_UPDATED_TIMESTAMP"
	// Choose TIMESTAMP to sort them by the time when the alarms most recently changed state, no matter the current alarm state.
	//
	// The alarm that changed state most recently is listed first.
	AlarmStatusWidgetSortBy_TIMESTAMP AlarmStatusWidgetSortBy = "TIMESTAMP"
)

