package mixinsawsce


// The recipient of `AnomalySubscription` notifications.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   subscriberProperty := &SubscriberProperty{
//   	Address: jsii.String("address"),
//   	Status: jsii.String("status"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ce-anomalysubscription-subscriber.html
//
type CfnAnomalySubscriptionPropsMixin_SubscriberProperty struct {
	// The email address or SNS Topic Amazon Resource Name (ARN), depending on the `Type` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ce-anomalysubscription-subscriber.html#cfn-ce-anomalysubscription-subscriber-address
	//
	Address *string `field:"optional" json:"address" yaml:"address"`
	// Indicates if the subscriber accepts the notifications.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ce-anomalysubscription-subscriber.html#cfn-ce-anomalysubscription-subscriber-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The notification delivery channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ce-anomalysubscription-subscriber.html#cfn-ce-anomalysubscription-subscriber-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

