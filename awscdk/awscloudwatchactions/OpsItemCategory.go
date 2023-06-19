package awscloudwatchactions


// Types of OpsItem category available.
//
// Example:
//   var alarm alarm
//
//   // Create an OpsItem with specific severity and category when alarm triggers
//   alarm.AddAlarmAction(
//   actions.NewSsmAction(actions.OpsItemSeverity_CRITICAL, actions.OpsItemCategory_PERFORMANCE))
//
// Experimental.
type OpsItemCategory string

const (
	// Set the category to availability.
	// Experimental.
	OpsItemCategory_AVAILABILITY OpsItemCategory = "AVAILABILITY"
	// Set the category to cost.
	// Experimental.
	OpsItemCategory_COST OpsItemCategory = "COST"
	// Set the category to performance.
	// Experimental.
	OpsItemCategory_PERFORMANCE OpsItemCategory = "PERFORMANCE"
	// Set the category to recovery.
	// Experimental.
	OpsItemCategory_RECOVERY OpsItemCategory = "RECOVERY"
	// Set the category to security.
	// Experimental.
	OpsItemCategory_SECURITY OpsItemCategory = "SECURITY"
)

