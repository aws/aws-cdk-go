package awslogs


// Properties for a SubscriptionFilter.
//
// Example:
//   import destinations "github.com/aws/aws-cdk-go/awscdk"
//
//   var fn Function
//   var logGroup LogGroup
//
//
//   logs.NewSubscriptionFilter(this, jsii.String("Subscription"), &SubscriptionFilterProps{
//   	LogGroup: LogGroup,
//   	Destination: destinations.NewLambdaDestination(fn),
//   	FilterPattern: logs.FilterPattern_AllTerms(jsii.String("ERROR"), jsii.String("MainThread")),
//   	FilterName: jsii.String("ErrorInMainThread"),
//   })
//
type SubscriptionFilterProps struct {
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
	// The log group to create the subscription on.
	LogGroup ILogGroup `field:"required" json:"logGroup" yaml:"logGroup"`
}

