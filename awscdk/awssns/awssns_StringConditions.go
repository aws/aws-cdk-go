package awssns


// Conditions that can be applied to string attributes.
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
type StringConditions struct {
	// Match one or more values.
	// Experimental.
	Allowlist *[]*string `field:"optional" json:"allowlist" yaml:"allowlist"`
	// Match any value that doesn't include any of the specified values.
	// Deprecated: use `denylist`.
	Blacklist *[]*string `field:"optional" json:"blacklist" yaml:"blacklist"`
	// Match any value that doesn't include any of the specified values.
	// Experimental.
	Denylist *[]*string `field:"optional" json:"denylist" yaml:"denylist"`
	// Matches values that begins with the specified prefixes.
	// Experimental.
	MatchPrefixes *[]*string `field:"optional" json:"matchPrefixes" yaml:"matchPrefixes"`
	// Match one or more values.
	// Deprecated: use `allowlist`.
	Whitelist *[]*string `field:"optional" json:"whitelist" yaml:"whitelist"`
}

