package awspipes


// Filter events using an event pattern.
//
// For more information, see [Events and Event Patterns](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html) in the *Amazon EventBridge User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterProperty := &FilterProperty{
//   	Pattern: jsii.String("pattern"),
//   }
//
type CfnPipe_FilterProperty struct {
	// The event pattern.
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
}

