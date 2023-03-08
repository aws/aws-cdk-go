package awscloudwatchactions


// Types of OpsItem severity available.
//
// Example:
//   var alarm alarm
//
//   // Create an OpsItem with specific severity and category when alarm triggers
//   alarm.addAlarmAction(
//   actions.NewSsmAction(actions.opsItemSeverity_CRITICAL, actions.opsItemCategory_PERFORMANCE))
//
// Experimental.
type OpsItemSeverity string

const (
	// Set the severity to critical.
	// Experimental.
	OpsItemSeverity_CRITICAL OpsItemSeverity = "CRITICAL"
	// Set the severity to high.
	// Experimental.
	OpsItemSeverity_HIGH OpsItemSeverity = "HIGH"
	// Set the severity to medium.
	// Experimental.
	OpsItemSeverity_MEDIUM OpsItemSeverity = "MEDIUM"
	// Set the severity to low.
	// Experimental.
	OpsItemSeverity_LOW OpsItemSeverity = "LOW"
)

