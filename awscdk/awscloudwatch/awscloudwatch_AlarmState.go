package awscloudwatch


// Enumeration indicates state of Alarm used in building Alarm Rule.
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
// Experimental.
type AlarmState string

const (
	// State indicates resource is in ALARM.
	// Experimental.
	AlarmState_ALARM AlarmState = "ALARM"
	// State indicates resource is not in ALARM.
	// Experimental.
	AlarmState_OK AlarmState = "OK"
	// State indicates there is not enough data to determine is resource is in ALARM.
	// Experimental.
	AlarmState_INSUFFICIENT_DATA AlarmState = "INSUFFICIENT_DATA"
)

