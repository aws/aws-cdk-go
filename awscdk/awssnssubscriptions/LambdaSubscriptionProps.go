package awssnssubscriptions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Properties for a Lambda subscription.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//   var fn function
//
//
//   myTopic := sns.NewTopic(this, jsii.String("MyTopic"))
//
//   // Lambda should receive only message matching the following conditions on attributes:
//   // color: 'red' or 'orange' or begins with 'bl'
//   // size: anything but 'small' or 'medium'
//   // price: between 100 and 200 or greater than 300
//   // store: attribute must be present
//   myTopic.AddSubscription(subscriptions.NewLambdaSubscription(fn, &LambdaSubscriptionProps{
//   	FilterPolicy: map[string]subscriptionFilter{
//   		"color": sns.*subscriptionFilter_stringFilter(&StringConditions{
//   			"allowlist": []*string{
//   				jsii.String("red"),
//   				jsii.String("orange"),
//   			},
//   			"matchPrefixes": []*string{
//   				jsii.String("bl"),
//   			},
//   			"matchSuffixes": []*string{
//   				jsii.String("ue"),
//   			},
//   		}),
//   		"size": sns.*subscriptionFilter_stringFilter(&StringConditions{
//   			"denylist": []*string{
//   				jsii.String("small"),
//   				jsii.String("medium"),
//   			},
//   		}),
//   		"price": sns.*subscriptionFilter_numericFilter(&NumericConditions{
//   			"between": &BetweenCondition{
//   				"start": jsii.Number(100),
//   				"stop": jsii.Number(200),
//   			},
//   			"greaterThan": jsii.Number(300),
//   		}),
//   		"store": sns.*subscriptionFilter_existsFilter(),
//   	},
//   }))
//
type LambdaSubscriptionProps struct {
	// Queue to be used as dead letter queue.
	//
	// If not passed no dead letter queue is enabled.
	// Default: - No dead letter queue enabled.
	//
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The filter policy.
	// Default: - all messages are delivered.
	//
	FilterPolicy *map[string]awssns.SubscriptionFilter `field:"optional" json:"filterPolicy" yaml:"filterPolicy"`
	// The filter policy that is applied on the message body.
	//
	// To apply a filter policy to the message attributes, use `filterPolicy`. A maximum of one of `filterPolicyWithMessageBody` and `filterPolicy` may be used.
	// Default: - all messages are delivered.
	//
	FilterPolicyWithMessageBody *map[string]awssns.FilterOrPolicy `field:"optional" json:"filterPolicyWithMessageBody" yaml:"filterPolicyWithMessageBody"`
}

