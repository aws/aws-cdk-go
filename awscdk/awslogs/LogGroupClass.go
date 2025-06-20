package awslogs


// Class of Log Group.
type LogGroupClass string

const (
	// Default class of logs services.
	LogGroupClass_STANDARD LogGroupClass = "STANDARD"
	// Class for reduced logs services.
	LogGroupClass_INFREQUENT_ACCESS LogGroupClass = "INFREQUENT_ACCESS"
)

