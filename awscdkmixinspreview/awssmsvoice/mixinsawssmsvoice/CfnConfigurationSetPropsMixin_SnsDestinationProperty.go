package mixinsawssmsvoice


// An object that defines an Amazon SNS destination for events.
//
// You can use Amazon SNS to send notification when certain events occur.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   snsDestinationProperty := &SnsDestinationProperty{
//   	TopicArn: jsii.String("topicArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-configurationset-snsdestination.html
//
type CfnConfigurationSetPropsMixin_SnsDestinationProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon SNS topic that you want to publish events to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-configurationset-snsdestination.html#cfn-smsvoice-configurationset-snsdestination-topicarn
	//
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}

