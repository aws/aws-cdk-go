package awscodepipeline


// The method that you want to configure for automatic stage retry on stage failure.
//
// You can specify to retry only failed action in the stage or all actions in the stage.
type RetryMode string

const (
	// Retry all actions under this stage.
	RetryMode_ALL_ACTIONS RetryMode = "ALL_ACTIONS"
	// Only retry failed actions.
	RetryMode_FAILED_ACTIONS RetryMode = "FAILED_ACTIONS"
)

