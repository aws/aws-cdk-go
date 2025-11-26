package previewawscemixins


// Properties for CfnAnomalySubscriptionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAnomalySubscriptionMixinProps := &CfnAnomalySubscriptionMixinProps{
//   	Frequency: jsii.String("frequency"),
//   	MonitorArnList: []*string{
//   		jsii.String("monitorArnList"),
//   	},
//   	ResourceTags: []ResourceTagProperty{
//   		&ResourceTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Subscribers: []interface{}{
//   		&SubscriberProperty{
//   			Address: jsii.String("address"),
//   			Status: jsii.String("status"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	SubscriptionName: jsii.String("subscriptionName"),
//   	Threshold: jsii.Number(123),
//   	ThresholdExpression: jsii.String("thresholdExpression"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalysubscription.html
//
type CfnAnomalySubscriptionMixinProps struct {
	// The frequency that anomaly notifications are sent.
	//
	// Notifications are sent either over email (for DAILY and WEEKLY frequencies) or SNS (for IMMEDIATE frequency). For more information, see [Creating an Amazon SNS topic for anomaly notifications](https://docs.aws.amazon.com/cost-management/latest/userguide/ad-SNS.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalysubscription.html#cfn-ce-anomalysubscription-frequency
	//
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// A list of cost anomaly monitors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalysubscription.html#cfn-ce-anomalysubscription-monitorarnlist
	//
	MonitorArnList *[]*string `field:"optional" json:"monitorArnList" yaml:"monitorArnList"`
	// Tags to assign to subscription.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalysubscription.html#cfn-ce-anomalysubscription-resourcetags
	//
	ResourceTags *[]*CfnAnomalySubscriptionPropsMixin_ResourceTagProperty `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// A list of subscribers to notify.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalysubscription.html#cfn-ce-anomalysubscription-subscribers
	//
	Subscribers interface{} `field:"optional" json:"subscribers" yaml:"subscribers"`
	// The name for the subscription.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalysubscription.html#cfn-ce-anomalysubscription-subscriptionname
	//
	SubscriptionName *string `field:"optional" json:"subscriptionName" yaml:"subscriptionName"`
	// (deprecated).
	//
	// An absolute dollar value that must be exceeded by the anomaly's total impact (see [Impact](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Impact.html) for more details) for an anomaly notification to be generated.
	//
	// This field has been deprecated. To specify a threshold, use ThresholdExpression. Continued use of Threshold will be treated as shorthand syntax for a ThresholdExpression.
	//
	// One of Threshold or ThresholdExpression is required for `AWS::CE::AnomalySubscription` . You cannot specify both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalysubscription.html#cfn-ce-anomalysubscription-threshold
	//
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
	// An [Expression](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html) object in JSON string format used to specify the anomalies that you want to generate alerts for. This supports dimensions and nested expressions. The supported dimensions are `ANOMALY_TOTAL_IMPACT_ABSOLUTE` and `ANOMALY_TOTAL_IMPACT_PERCENTAGE` , corresponding to an anomaly’s TotalImpact and TotalImpactPercentage, respectively (see [Impact](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Impact.html) for more details). The supported nested expression types are `AND` and `OR` . The match option `GREATER_THAN_OR_EQUAL` is required. Values must be numbers between 0 and 10,000,000,000 in string format.
	//
	// One of Threshold or ThresholdExpression is required for `AWS::CE::AnomalySubscription` . You cannot specify both.
	//
	// For further information, see the [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalysubscription.html#aws-resource-ce-anomalysubscription--examples) section of this page.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalysubscription.html#cfn-ce-anomalysubscription-thresholdexpression
	//
	ThresholdExpression *string `field:"optional" json:"thresholdExpression" yaml:"thresholdExpression"`
}

