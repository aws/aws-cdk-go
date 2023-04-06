package awsce


// Properties for defining a `CfnAnomalySubscription`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAnomalySubscriptionProps := &CfnAnomalySubscriptionProps{
//   	Frequency: jsii.String("frequency"),
//   	MonitorArnList: []*string{
//   		jsii.String("monitorArnList"),
//   	},
//   	Subscribers: []interface{}{
//   		&SubscriberProperty{
//   			Address: jsii.String("address"),
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Status: jsii.String("status"),
//   		},
//   	},
//   	SubscriptionName: jsii.String("subscriptionName"),
//
//   	// the properties below are optional
//   	ResourceTags: []interface{}{
//   		&ResourceTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Threshold: jsii.Number(123),
//   	ThresholdExpression: jsii.String("thresholdExpression"),
//   }
//
type CfnAnomalySubscriptionProps struct {
	// The frequency that anomaly notifications are sent.
	//
	// Notifications are sent either over email (for DAILY and WEEKLY frequencies) or SNS (for IMMEDIATE frequency). For more information, see [Creating an Amazon SNS topic for anomaly notifications](https://docs.aws.amazon.com/cost-management/latest/userguide/ad-SNS.html) .
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
	// An absolute dollar value that must be exceeded by the anomaly's total impact (see [Impact](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Impact.html) for more details) for an anomaly notification to be generated.
	//
	// This field has been deprecated. To specify a threshold, use ThresholdExpression. Continued use of Threshold will be treated as shorthand syntax for a ThresholdExpression.
	//
	// One of Threshold or ThresholdExpression is required for `AWS::CE::AnomalySubscription` . You cannot specify both.
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
	// An [Expression](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html) object in JSON string format used to specify the anomalies that you want to generate alerts for. This supports dimensions and nested expressions. The supported dimensions are `ANOMALY_TOTAL_IMPACT_ABSOLUTE` and `ANOMALY_TOTAL_IMPACT_PERCENTAGE` , corresponding to an anomaly’s TotalImpact and TotalImpactPercentage, respectively (see [Impact](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Impact.html) for more details). The supported nested expression types are `AND` and `OR` . The match option `GREATER_THAN_OR_EQUAL` is required. Values must be numbers between 0 and 10,000,000,000 in string format.
	//
	// One of Threshold or ThresholdExpression is required for `AWS::CE::AnomalySubscription` . You cannot specify both.
	//
	// For further information, see the [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalysubscription.html#aws-resource-ce-anomalysubscription--examples) section of this page.
	ThresholdExpression *string `field:"optional" json:"thresholdExpression" yaml:"thresholdExpression"`
}

