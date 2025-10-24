package awssmsvoice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnConfigurationSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfigurationSetProps := &CfnConfigurationSetProps{
//   	ConfigurationSetName: jsii.String("configurationSetName"),
//   	DefaultSenderId: jsii.String("defaultSenderId"),
//   	EventDestinations: []interface{}{
//   		&EventDestinationProperty{
//   			Enabled: jsii.Boolean(false),
//   			EventDestinationName: jsii.String("eventDestinationName"),
//   			MatchingEventTypes: []*string{
//   				jsii.String("matchingEventTypes"),
//   			},
//
//   			// the properties below are optional
//   			CloudWatchLogsDestination: &CloudWatchLogsDestinationProperty{
//   				IamRoleArn: jsii.String("iamRoleArn"),
//   				LogGroupArn: jsii.String("logGroupArn"),
//   			},
//   			KinesisFirehoseDestination: &KinesisFirehoseDestinationProperty{
//   				DeliveryStreamArn: jsii.String("deliveryStreamArn"),
//   				IamRoleArn: jsii.String("iamRoleArn"),
//   			},
//   			SnsDestination: &SnsDestinationProperty{
//   				TopicArn: jsii.String("topicArn"),
//   			},
//   		},
//   	},
//   	MessageFeedbackEnabled: jsii.Boolean(false),
//   	ProtectConfigurationId: jsii.String("protectConfigurationId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-configurationset.html
//
type CfnConfigurationSetProps struct {
	// The name of the ConfigurationSet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-configurationset.html#cfn-smsvoice-configurationset-configurationsetname
	//
	ConfigurationSetName *string `field:"optional" json:"configurationSetName" yaml:"configurationSetName"`
	// The default sender ID used by the ConfigurationSet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-configurationset.html#cfn-smsvoice-configurationset-defaultsenderid
	//
	DefaultSenderId *string `field:"optional" json:"defaultSenderId" yaml:"defaultSenderId"`
	// An array of EventDestination objects that describe any events to log and where to log them.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-configurationset.html#cfn-smsvoice-configurationset-eventdestinations
	//
	EventDestinations interface{} `field:"optional" json:"eventDestinations" yaml:"eventDestinations"`
	// Set to true to enable feedback for the message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-configurationset.html#cfn-smsvoice-configurationset-messagefeedbackenabled
	//
	MessageFeedbackEnabled interface{} `field:"optional" json:"messageFeedbackEnabled" yaml:"messageFeedbackEnabled"`
	// The unique identifier for the protect configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-configurationset.html#cfn-smsvoice-configurationset-protectconfigurationid
	//
	ProtectConfigurationId *string `field:"optional" json:"protectConfigurationId" yaml:"protectConfigurationId"`
	// An array of key and value pair tags that's associated with the new configuration set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-configurationset.html#cfn-smsvoice-configurationset-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

