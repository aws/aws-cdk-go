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
type StringConditions struct {
	// Match one or more values.
	// Default: - None.
	//
	Allowlist *[]*string `field:"optional" json:"allowlist" yaml:"allowlist"`
	// Match any value that doesn't include any of the specified values.
	// Default: - None.
	//
	Denylist *[]*string `field:"optional" json:"denylist" yaml:"denylist"`
	// Matches values that begins with the specified prefixes.
	// Default: - None.
	//
	MatchPrefixes *[]*string `field:"optional" json:"matchPrefixes" yaml:"matchPrefixes"`
	// Matches values that end with the specified suffixes.
	// Default: - None.
	//
	MatchSuffixes *[]*string `field:"optional" json:"matchSuffixes" yaml:"matchSuffixes"`
}

