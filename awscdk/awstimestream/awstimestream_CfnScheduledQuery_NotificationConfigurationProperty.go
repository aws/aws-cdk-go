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
//   notificationConfigurationProperty := &notificationConfigurationProperty{
//   	snsConfiguration: &snsConfigurationProperty{
//   		topicArn: jsii.String("topicArn"),
//   	},
//   }
//
type CfnScheduledQuery_NotificationConfigurationProperty struct {
	// Details on SNS configuration.
	SnsConfiguration interface{} `field:"required" json:"snsConfiguration" yaml:"snsConfiguration"`
}

