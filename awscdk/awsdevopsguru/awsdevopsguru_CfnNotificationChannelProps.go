package awsdevopsguru


// Properties for defining a `CfnNotificationChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNotificationChannelProps := &cfnNotificationChannelProps{
//   	config: &notificationChannelConfigProperty{
//   		filters: &notificationFilterConfigProperty{
//   			messageTypes: []*string{
//   				jsii.String("messageTypes"),
//   			},
//   			severities: []*string{
//   				jsii.String("severities"),
//   			},
//   		},
//   		sns: &snsChannelConfigProperty{
//   			topicArn: jsii.String("topicArn"),
//   		},
//   	},
//   }
//
type CfnNotificationChannelProps struct {
	// A `NotificationChannelConfig` object that contains information about configured notification channels.
	Config interface{} `field:"required" json:"config" yaml:"config"`
}

