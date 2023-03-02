package awsconfig


// Provides options for how often AWS Config delivers configuration snapshots to the Amazon S3 bucket in your delivery channel.
//
// > If you want to create a rule that triggers evaluations for your resources when AWS Config delivers the configuration snapshot, see the following:
//
// The frequency for a rule that triggers evaluations for your resources when AWS Config delivers the configuration snapshot is set by one of two values, depending on which is less frequent:
//
// - The value for the `deliveryFrequency` parameter within the delivery channel configuration, which sets how often AWS Config delivers configuration snapshots. This value also sets how often AWS Config invokes evaluations for AWS Config rules.
// - The value for the `MaximumExecutionFrequency` parameter, which sets the maximum frequency with which AWS Config invokes evaluations for the rule. For more information, see [ConfigRule](https://docs.aws.amazon.com/config/latest/APIReference/API_ConfigRule.html) .
//
// If the `deliveryFrequency` value is less frequent than the `MaximumExecutionFrequency` value for a rule, AWS Config invokes the rule only as often as the `deliveryFrequency` value.
//
// - For example, you want your rule to run evaluations when AWS Config delivers the configuration snapshot.
// - You specify the `MaximumExecutionFrequency` value for `Six_Hours` .
// - You then specify the delivery channel `deliveryFrequency` value for `TwentyFour_Hours` .
// - Because the value for `deliveryFrequency` is less frequent than `MaximumExecutionFrequency` , AWS Config invokes evaluations for the rule every 24 hours.
//
// You should set the `MaximumExecutionFrequency` value to be at least as frequent as the `deliveryFrequency` value. You can view the `deliveryFrequency` value by using the `DescribeDeliveryChannnels` action.
//
// To update the `deliveryFrequency` with which AWS Config delivers your configuration snapshots, use the `PutDeliveryChannel` action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configSnapshotDeliveryPropertiesProperty := &configSnapshotDeliveryPropertiesProperty{
//   	deliveryFrequency: jsii.String("deliveryFrequency"),
//   }
//
type CfnDeliveryChannel_ConfigSnapshotDeliveryPropertiesProperty struct {
	// The frequency with which AWS Config delivers configuration snapshots.
	DeliveryFrequency *string `field:"optional" json:"deliveryFrequency" yaml:"deliveryFrequency"`
}

