package awscloudwatch


// Enumeration indicates state of Alarm used in building Alarm Rule.
//
// Example:
//   var dashboard dashboard
//   var errorAlarm alarm
//
//
//   dashboard.addWidgets(cloudwatch.NewAlarmStatusWidget(&alarmStatusWidgetProps{
//   	title: jsii.String("Errors"),
//   	alarms: []iAlarm{
//   		errorAlarm,
//   	},
//   	sortBy: cloudwatch.alarmStatusWidgetSortBy_STATE_UPDATED_TIMESTAMP,
//   	states: []alarmState{
//   		cloudwatch.*alarmState_ALARM,
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

