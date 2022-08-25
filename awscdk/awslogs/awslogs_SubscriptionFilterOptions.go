package awslogs


// Properties for a new SubscriptionFilter created from a LogGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var filterPattern iFilterPattern
//   var logSubscriptionDestination iLogSubscriptionDestination
//
//   subscriptionFilterOptions := &subscriptionFilterOptions{
//   	destination: logSubscriptionDestination,
//   	filterPattern: filterPattern,
//   }
//
type SubscriptionFilterOptions struct {
	// The destination to send the filtered events to.
	//
	// For example, a Kinesis stream or a Lambda function.
	Destination ILogSubscriptionDestination `field:"required" json:"destination" yaml:"destination"`
	// Log events matching this pattern will be sent to the destination.
	FilterPattern IFilterPattern `field:"required" json:"filterPattern" yaml:"filterPattern"`
}

