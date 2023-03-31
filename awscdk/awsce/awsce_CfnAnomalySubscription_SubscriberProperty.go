package awsce


// The recipient of `AnomalySubscription` notifications.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subscriberProperty := &subscriberProperty{
//   	address: jsii.String("address"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	status: jsii.String("status"),
//   }
//
type CfnAnomalySubscription_SubscriberProperty struct {
	// The email address or SNS Topic Amazon Resource Name (ARN), depending on the `Type` .
	Address *string `field:"required" json:"address" yaml:"address"`
	// The notification delivery channel.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Indicates if the subscriber accepts the notifications.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

