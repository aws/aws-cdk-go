package awsstepfunctions


// Three ways to call an integrated service: Request Response, Run a Job and Wait for a Callback with Task Token.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html
//
// Here, they are named as FIRE_AND_FORGET, SYNC and WAIT_FOR_TASK_TOKEN respectfully.
//
// Default: FIRE_AND_FORGET.
//
type ServiceIntegrationPattern string

const (
	// Call a service and progress to the next state immediately after the API call completes.
	ServiceIntegrationPattern_FIRE_AND_FORGET ServiceIntegrationPattern = "FIRE_AND_FORGET"
	// Call a service and wait for a job to complete.
	ServiceIntegrationPattern_SYNC ServiceIntegrationPattern = "SYNC"
	// Call a service with a task token and wait until that token is returned by SendTaskSuccess/SendTaskFailure with payload.
	ServiceIntegrationPattern_WAIT_FOR_TASK_TOKEN ServiceIntegrationPattern = "WAIT_FOR_TASK_TOKEN"
)

