package awsce


// Properties for defining a `CfnAnomalySubscription`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAnomalySubscriptionProps := &cfnAnomalySubscriptionProps{
//   	frequency: jsii.String("frequency"),
//   	monitorArnList: []*string{
//   		jsii.String("monitorArnList"),
//   	},
//   	subscribers: []interface{}{
//   		&subscriberProperty{
//   			address: jsii.String("address"),
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			status: jsii.String("status"),
//   		},
//   	},
//   	subscriptionName: jsii.String("subscriptionName"),
//
//   	// the properties below are optional
//   	resourceTags: []interface{}{
//   		&resourceTagProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	threshold: jsii.Number(123),
//   	thresholdExpression: jsii.String("thresholdExpression"),
//   }
//
type CfnAnomalySubscriptionProps struct {
	// The frequency that anomaly reports are sent over email.
	Frequency *string `field:"required" json:"frequency" yaml:"frequency"`
	// A list of cost anomaly monitors.
	MonitorArnList *[]*string `field:"required" json:"monitorArnList" yaml:"monitorArnList"`
	// A list of subscribers to notify.
	Subscribers interface{} `field:"required" json:"subscribers" yaml:"subscribers"`
	// The name for the subscription.
	SubscriptionName *string `field:"required" json:"subscriptionName" yaml:"subscriptionName"`
	// `AWS::CE::AnomalySubscription.ResourceTags`.
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// (deprecated).
	//
	// The dollar value that triggers a notification if the threshold is exceeded.
	//
	// This field has been deprecated. To specify a threshold, use ThresholdExpression. Continued use of Threshold will be treated as shorthand syntax for a ThresholdExpression.
	//
	// One of Threshold or ThresholdExpression is required for this resource.
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
	// An [Expression](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html) object used to specify the anomalies that you want to generate alerts for. This supports dimensions and nested expressions. The supported dimensions are `ANOMALY_TOTAL_IMPACT_ABSOLUTE` and `ANOMALY_TOTAL_IMPACT_PERCENTAGE` . The supported nested expression types are `AND` and `OR` . The match option `GREATER_THAN_OR_EQUAL` is required. Values must be numbers between 0 and 10,000,000,000.
	//
	// One of Threshold or ThresholdExpression is required for this resource.
	//
	// The following are examples of valid ThresholdExpressions:
	//
	// - Absolute threshold: `{ "Dimensions": { "Key": "ANOMALY_TOTAL_IMPACT_ABSOLUTE", "MatchOptions": [ "GREATER_THAN_OR_EQUAL" ], "Values": [ "100" ] } }`
	// - Percentage threshold: `{ "Dimensions": { "Key": "ANOMALY_TOTAL_IMPACT_PERCENTAGE", "MatchOptions": [ "GREATER_THAN_OR_EQUAL" ], "Values": [ "100" ] } }`
	// - `AND` two thresholds together: `{ "And": [ { "Dimensions": { "Key": "ANOMALY_TOTAL_IMPACT_ABSOLUTE", "MatchOptions": [ "GREATER_THAN_OR_EQUAL" ], "Values": [ "100" ] } }, { "Dimensions": { "Key": "ANOMALY_TOTAL_IMPACT_PERCENTAGE", "MatchOptions": [ "GREATER_THAN_OR_EQUAL" ], "Values": [ "100" ] } } ] }`
	// - `OR` two thresholds together: `{ "Or": [ { "Dimensions": { "Key": "ANOMALY_TOTAL_IMPACT_ABSOLUTE", "MatchOptions": [ "GREATER_THAN_OR_EQUAL" ], "Values": [ "100" ] } }, { "Dimensions": { "Key": "ANOMALY_TOTAL_IMPACT_PERCENTAGE", "MatchOptions": [ "GREATER_THAN_OR_EQUAL" ], "Values": [ "100" ] } } ] }`.
	ThresholdExpression *string `field:"optional" json:"thresholdExpression" yaml:"thresholdExpression"`
}

