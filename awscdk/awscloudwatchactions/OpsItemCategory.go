package awscloudwatchactions


// Types of OpsItem category available.
//
// Example:
//   var alarm Alarm
//
//   // Create an OpsItem with specific severity and category when alarm triggers
//   alarm.AddAlarmAction(
//   actions.NewSsmAction(actions.OpsItemSeverity_CRITICAL, actions.OpsItemCategory_PERFORMANCE))
//
type OpsItemCategory string

const (
	// Set the category to availability.
	OpsItemCategory_AVAILABILITY OpsItemCategory = "AVAILABILITY"
	// Set the category to cost.
	OpsItemCategory_COST OpsItemCategory = "COST"
	// Set the category to performance.
	OpsItemCategory_PERFORMANCE OpsItemCategory = "PERFORMANCE"
	// Set the category to recovery.
	OpsItemCategory_RECOVERY OpsItemCategory = "RECOVERY"
	// Set the category to security.
	OpsItemCategory_SECURITY OpsItemCategory = "SECURITY"
)

