package awslogs


// A reference to a SubscriptionFilter resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subscriptionFilterReference := &SubscriptionFilterReference{
//   	FilterName: jsii.String("filterName"),
//   	LogGroupName: jsii.String("logGroupName"),
//   }
//
type SubscriptionFilterReference struct {
	// The FilterName of the SubscriptionFilter resource.
	FilterName *string `field:"required" json:"filterName" yaml:"filterName"`
	// The LogGroupName of the SubscriptionFilter resource.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
}

