package awsce


// The recipient of `AnomalySubscription` notifications.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subscriberProperty := &SubscriberProperty{
//   	Address: jsii.String("address"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ce-anomalysubscription-subscriber.html
//
type CfnAnomalySubscription_SubscriberProperty struct {
	// The email address or SNS Topic Amazon Resource Name (ARN), depending on the `Type` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ce-anomalysubscription-subscriber.html#cfn-ce-anomalysubscription-subscriber-address
	//
	Address *string `field:"required" json:"address" yaml:"address"`
	// The notification delivery channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ce-anomalysubscription-subscriber.html#cfn-ce-anomalysubscription-subscriber-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Indicates if the subscriber accepts the notifications.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ce-anomalysubscription-subscriber.html#cfn-ce-anomalysubscription-subscriber-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

