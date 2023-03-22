package awsbudgets


// The subscriber to a budget notification.
//
// The subscriber consists of a subscription type and either an Amazon SNS topic or an email address.
//
// For example, an email subscriber has the following parameters:
//
// - A `subscriptionType` of `EMAIL`
// - An `address` of `example@example.com`
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subscriberProperty := &subscriberProperty{
//   	address: jsii.String("address"),
//   	type: jsii.String("type"),
//   }
//
type CfnBudgetsAction_SubscriberProperty struct {
	// The address that AWS sends budget notifications to, either an SNS topic or an email.
	//
	// When you create a subscriber, the value of `Address` can't contain line breaks.
	Address *string `field:"required" json:"address" yaml:"address"`
	// The type of notification that AWS sends to a subscriber.
	Type *string `field:"required" json:"type" yaml:"type"`
}

