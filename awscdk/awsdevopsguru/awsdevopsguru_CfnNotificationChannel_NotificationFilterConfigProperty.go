package awsdevopsguru


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationFilterConfigProperty := &notificationFilterConfigProperty{
//   	messageTypes: []*string{
//   		jsii.String("messageTypes"),
//   	},
//   	severities: []*string{
//   		jsii.String("severities"),
//   	},
//   }
//
type CfnNotificationChannel_NotificationFilterConfigProperty struct {
	// `CfnNotificationChannel.NotificationFilterConfigProperty.MessageTypes`.
	MessageTypes *[]*string `field:"optional" json:"messageTypes" yaml:"messageTypes"`
	// `CfnNotificationChannel.NotificationFilterConfigProperty.Severities`.
	Severities *[]*string `field:"optional" json:"severities" yaml:"severities"`
}

