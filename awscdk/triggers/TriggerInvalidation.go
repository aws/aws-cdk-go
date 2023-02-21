package triggers


// Determines.
type TriggerInvalidation string

const (
	// The trigger will be executed every time the handler (or its configuration) changes.
	//
	// This is implemented by associated the trigger with the `currentVersion`
	// of the AWS Lambda function, which gets recreated every time the handler changes.
	TriggerInvalidation_HANDLER_CHANGE TriggerInvalidation = "HANDLER_CHANGE"
)

