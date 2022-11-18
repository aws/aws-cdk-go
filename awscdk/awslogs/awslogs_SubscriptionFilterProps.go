package awslogs


// Properties for a SubscriptionFilter.
//
// Example:
//   import destinations "github.com/aws/aws-cdk-go/awscdk"
//   var fn function
//   var logGroup logGroup
//
//
//   logs.NewSubscriptionFilter(this, jsii.String("Subscription"), &subscriptionFilterProps{
//   	logGroup: logGroup,
//   	destination: destinations.NewLambdaDestination(fn),
//   	filterPattern: logs.filterPattern.allTerms(jsii.String("ERROR"), jsii.String("MainThread")),
//   })
//
type SubscriptionFilterProps struct {
	// The destination to send the filtered events to.
	//
	// For example, a Kinesis stream or a Lambda function.
	Destination ILogSubscriptionDestination `field:"required" json:"destination" yaml:"destination"`
	// Log events matching this pattern will be sent to the destination.
	FilterPattern IFilterPattern `field:"required" json:"filterPattern" yaml:"filterPattern"`
	// The log group to create the subscription on.
	LogGroup ILogGroup `field:"required" json:"logGroup" yaml:"logGroup"`
}

