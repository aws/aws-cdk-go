package awsmediaconnectalpha


// Days of the week for scheduled maintenance windows.
//
// Used by both Flow and Router maintenance configurations.
// Experimental.
type MaintenanceDay string

const (
	// Monday.
	// Experimental.
	MaintenanceDay_MONDAY MaintenanceDay = "MONDAY"
	// Tuesday.
	// Experimental.
	MaintenanceDay_TUESDAY MaintenanceDay = "TUESDAY"
	// Wednesday.
	// Experimental.
	MaintenanceDay_WEDNESDAY MaintenanceDay = "WEDNESDAY"
	// Thursday.
	// Experimental.
	MaintenanceDay_THURSDAY MaintenanceDay = "THURSDAY"
	// Friday.
	// Experimental.
	MaintenanceDay_FRIDAY MaintenanceDay = "FRIDAY"
	// Saturday.
	// Experimental.
	MaintenanceDay_SATURDAY MaintenanceDay = "SATURDAY"
	// Sunday.
	// Experimental.
	MaintenanceDay_SUNDAY MaintenanceDay = "SUNDAY"
)

