package awstimestream


// Notification configuration for a scheduled query.
//
// A notification is sent by Timestream when a scheduled query is created, its state is updated or when it is deleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationConfigurationProperty := &NotificationConfigurationProperty{
//   	SnsConfiguration: &SnsConfigurationProperty{
//   		TopicArn: jsii.String("topicArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-scheduledquery-notificationconfiguration.html
//
type CfnScheduledQuery_NotificationConfigurationProperty struct {
	// Details on SNS configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-scheduledquery-notificationconfiguration.html#cfn-timestream-scheduledquery-notificationconfiguration-snsconfiguration
	//
	SnsConfiguration interface{} `field:"required" json:"snsConfiguration" yaml:"snsConfiguration"`
}

