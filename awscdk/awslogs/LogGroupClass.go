package awslogs


// Class of Log Group.
//
// Example:
//   logs.NewLogGroup(this, jsii.String("DeliveryLogGroup"), &LogGroupProps{
//   	LogGroupClass: logs.LogGroupClass_DELIVERY,
//   })
//
type LogGroupClass string

const (
	// Default class of logs services.
	LogGroupClass_STANDARD LogGroupClass = "STANDARD"
	// Class for reduced logs services.
	LogGroupClass_INFREQUENT_ACCESS LogGroupClass = "INFREQUENT_ACCESS"
	// Class for delivering logs to a destination such as Amazon S3 or Amazon Data Firehose (for example, Lambda vended logs).
	LogGroupClass_DELIVERY LogGroupClass = "DELIVERY"
)

