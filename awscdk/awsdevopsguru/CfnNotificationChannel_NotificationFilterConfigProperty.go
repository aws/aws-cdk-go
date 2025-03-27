package awsdevopsguru


// The filter configurations for the Amazon SNS notification topic you use with DevOps Guru.
//
// You can choose to specify which events or message types to receive notifications for. You can also choose to specify which severity levels to receive notifications for.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationFilterConfigProperty := &NotificationFilterConfigProperty{
//   	MessageTypes: []*string{
//   		jsii.String("messageTypes"),
//   	},
//   	Severities: []*string{
//   		jsii.String("severities"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsguru-notificationchannel-notificationfilterconfig.html
//
type CfnNotificationChannel_NotificationFilterConfigProperty struct {
	// The events that you want to receive notifications for.
	//
	// For example, you can choose to receive notifications only when the severity level is upgraded or a new insight is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsguru-notificationchannel-notificationfilterconfig.html#cfn-devopsguru-notificationchannel-notificationfilterconfig-messagetypes
	//
	MessageTypes *[]*string `field:"optional" json:"messageTypes" yaml:"messageTypes"`
	// The severity levels that you want to receive notifications for.
	//
	// For example, you can choose to receive notifications only for insights with `HIGH` and `MEDIUM` severity levels. For more information, see [Understanding insight severities](https://docs.aws.amazon.com/devops-guru/latest/userguide/working-with-insights.html#understanding-insights-severities) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsguru-notificationchannel-notificationfilterconfig.html#cfn-devopsguru-notificationchannel-notificationfilterconfig-severities
	//
	Severities *[]*string `field:"optional" json:"severities" yaml:"severities"`
}

