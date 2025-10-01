package awsevents


// The level of logging detail to include.
//
// This applies to all log destinations for the event bus.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   bus := awscdk.NewEventBus(this, jsii.String("Bus"), &EventBusProps{
//   	LogConfig: &LogConfig{
//   		IncludeDetail: awscdk.IncludeDetail_FULL,
//   		Level: awscdk.Level_TRACE,
//   	},
//   })
//
type Level string

const (
	// INFO: EventBridge sends any logs related to errors, as well as major steps performed during event processing.
	Level_INFO Level = "INFO"
	// ERROR: EventBridge sends any logs related to errors generated during event processing and target delivery.
	Level_ERROR Level = "ERROR"
	// TRACE: EventBridge sends any logs generated during all steps in the event processing.
	Level_TRACE Level = "TRACE"
	// OFF: EventBridge does not send any logs.
	//
	// This is the default.
	Level_OFF Level = "OFF"
)

