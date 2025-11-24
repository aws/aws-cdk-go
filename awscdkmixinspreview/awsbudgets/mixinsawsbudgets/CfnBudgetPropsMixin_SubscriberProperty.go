package mixinsawsbudgets


// The `Subscriber` property type specifies who to notify for a Billing and Cost Management budget notification.
//
// The subscriber consists of a subscription type, and either an Amazon SNS topic or an email address.
//
// For example, an email subscriber would have the following parameters:
//
// - A `subscriptionType` of `EMAIL`
// - An `address` of `example@example.com`
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   subscriberProperty := &SubscriberProperty{
//   	Address: jsii.String("address"),
//   	SubscriptionType: jsii.String("subscriptionType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-subscriber.html
//
type CfnBudgetPropsMixin_SubscriberProperty struct {
	// The address that AWS sends budget notifications to, either an SNS topic or an email.
	//
	// When you create a subscriber, the value of `Address` can't contain line breaks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-subscriber.html#cfn-budgets-budget-subscriber-address
	//
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The type of notification that AWS sends to a subscriber.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-subscriber.html#cfn-budgets-budget-subscriber-subscriptiontype
	//
	SubscriptionType *string `field:"optional" json:"subscriptionType" yaml:"subscriptionType"`
}

