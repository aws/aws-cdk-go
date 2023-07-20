package awsdevopsguru


// Properties for defining a `CfnNotificationChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNotificationChannelProps := &CfnNotificationChannelProps{
//   	Config: &NotificationChannelConfigProperty{
//   		Filters: &NotificationFilterConfigProperty{
//   			MessageTypes: []*string{
//   				jsii.String("messageTypes"),
//   			},
//   			Severities: []*string{
//   				jsii.String("severities"),
//   			},
//   		},
//   		Sns: &SnsChannelConfigProperty{
//   			TopicArn: jsii.String("topicArn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsguru-notificationchannel.html
//
type CfnNotificationChannelProps struct {
	// A `NotificationChannelConfig` object that contains information about configured notification channels.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsguru-notificationchannel.html#cfn-devopsguru-notificationchannel-config
	//
	Config interface{} `field:"required" json:"config" yaml:"config"`
}

