package awsevents


// Whether EventBridge include detailed event information in the records it generates.
//
// Detailed data can be useful for troubleshooting and debugging.
// This information includes details of the event itself, as well as target details.
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
type IncludeDetail string

const (
	// FULL: Include all details related to event itself and the request EventBridge sends to the target.
	//
	// Detailed data can be useful for troubleshooting and debugging.
	IncludeDetail_FULL IncludeDetail = "FULL"
	// NONE: Does not include any details.
	IncludeDetail_NONE IncludeDetail = "NONE"
)

