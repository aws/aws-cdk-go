package awscloudwatchactions


// Types of OpsItem severity available.
//
// Example:
//   var alarm Alarm
//
//   // Create an OpsItem with specific severity and category when alarm triggers
//   alarm.AddAlarmAction(
//   actions.NewSsmAction(actions.OpsItemSeverity_CRITICAL, actions.OpsItemCategory_PERFORMANCE))
//
type OpsItemSeverity string

const (
	// Set the severity to critical.
	OpsItemSeverity_CRITICAL OpsItemSeverity = "CRITICAL"
	// Set the severity to high.
	OpsItemSeverity_HIGH OpsItemSeverity = "HIGH"
	// Set the severity to medium.
	OpsItemSeverity_MEDIUM OpsItemSeverity = "MEDIUM"
	// Set the severity to low.
	OpsItemSeverity_LOW OpsItemSeverity = "LOW"
)

