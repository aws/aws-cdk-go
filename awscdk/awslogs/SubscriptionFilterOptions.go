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
//   subscriptionFilterOptions := &SubscriptionFilterOptions{
//   	Destination: logSubscriptionDestination,
//   	FilterPattern: filterPattern,
//
//   	// the properties below are optional
//   	Distribution: awscdk.Aws_logs.Distribution_BY_LOG_STREAM,
//   	FilterName: jsii.String("filterName"),
//   }
//
type SubscriptionFilterOptions struct {
	// The destination to send the filtered events to.
	//
	// For example, a Kinesis stream or a Lambda function.
	Destination ILogSubscriptionDestination `field:"required" json:"destination" yaml:"destination"`
	// Log events matching this pattern will be sent to the destination.
	FilterPattern IFilterPattern `field:"required" json:"filterPattern" yaml:"filterPattern"`
	// The method used to distribute log data to the destination.
	//
	// This property can only be used with KinesisDestination.
	// Default: Distribution.BY_LOG_STREAM
	//
	Distribution Distribution `field:"optional" json:"distribution" yaml:"distribution"`
	// The name of the subscription filter.
	// Default: Automatically generated.
	//
	FilterName *string `field:"optional" json:"filterName" yaml:"filterName"`
}

