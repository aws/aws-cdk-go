package awscloudwatch


// Enumeration indicates state of Alarm used in building Alarm Rule.
//
// Example:
//   var dashboard Dashboard
//   var errorAlarm Alarm
//
//
//   dashboard.AddWidgets(cloudwatch.NewAlarmStatusWidget(&AlarmStatusWidgetProps{
//   	Title: jsii.String("Errors"),
//   	Alarms: []IAlarmRef{
//   		errorAlarm,
//   	},
//   	SortBy: cloudwatch.AlarmStatusWidgetSortBy_STATE_UPDATED_TIMESTAMP,
//   	States: []AlarmState{
//   		cloudwatch.AlarmState_ALARM,
//   	},
//   }))
//
type AlarmState string

const (
	// State indicates resource is in ALARM.
	AlarmState_ALARM AlarmState = "ALARM"
	// State indicates resource is not in ALARM.
	AlarmState_OK AlarmState = "OK"
	// State indicates there is not enough data to determine is resource is in ALARM.
	AlarmState_INSUFFICIENT_DATA AlarmState = "INSUFFICIENT_DATA"
)

