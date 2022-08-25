package awssns


// Conditions that can be applied to numeric attributes.
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
//   myTopic.addSubscription(subscriptions.NewLambdaSubscription(fn, &lambdaSubscriptionProps{
//   	filterPolicy: map[string]subscriptionFilter{
//   		"color": sns.*subscriptionFilter.stringFilter(&StringConditions{
//   			"allowlist": []*string{
//   				jsii.String("red"),
//   				jsii.String("orange"),
//   			},
//   			"matchPrefixes": []*string{
//   				jsii.String("bl"),
//   			},
//   		}),
//   		"size": sns.*subscriptionFilter.stringFilter(&StringConditions{
//   			"denylist": []*string{
//   				jsii.String("small"),
//   				jsii.String("medium"),
//   			},
//   		}),
//   		"price": sns.*subscriptionFilter.numericFilter(&NumericConditions{
//   			"between": &BetweenCondition{
//   				"start": jsii.Number(100),
//   				"stop": jsii.Number(200),
//   			},
//   			"greaterThan": jsii.Number(300),
//   		}),
//   		"store": sns.*subscriptionFilter.existsFilter(),
//   	},
//   }))
//
type NumericConditions struct {
	// Match one or more values.
	Allowlist *[]*float64 `field:"optional" json:"allowlist" yaml:"allowlist"`
	// Match values that are between the specified values.
	Between *BetweenCondition `field:"optional" json:"between" yaml:"between"`
	// Match values that are strictly between the specified values.
	BetweenStrict *BetweenCondition `field:"optional" json:"betweenStrict" yaml:"betweenStrict"`
	// Match values that are greater than the specified value.
	GreaterThan *float64 `field:"optional" json:"greaterThan" yaml:"greaterThan"`
	// Match values that are greater than or equal to the specified value.
	GreaterThanOrEqualTo *float64 `field:"optional" json:"greaterThanOrEqualTo" yaml:"greaterThanOrEqualTo"`
	// Match values that are less than the specified value.
	LessThan *float64 `field:"optional" json:"lessThan" yaml:"lessThan"`
	// Match values that are less than or equal to the specified value.
	LessThanOrEqualTo *float64 `field:"optional" json:"lessThanOrEqualTo" yaml:"lessThanOrEqualTo"`
}

