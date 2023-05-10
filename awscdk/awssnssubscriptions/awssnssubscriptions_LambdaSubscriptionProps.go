package awssnssubscriptions

import (
	"github.com/aws/aws-cdk-go/awscdk/awssns"
	"github.com/aws/aws-cdk-go/awscdk/awssqs"
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
// Experimental.
type LambdaSubscriptionProps struct {
	// Queue to be used as dead letter queue.
	//
	// If not passed no dead letter queue is enabled.
	// Experimental.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The filter policy.
	// Experimental.
	FilterPolicy *map[string]awssns.SubscriptionFilter `field:"optional" json:"filterPolicy" yaml:"filterPolicy"`
}

