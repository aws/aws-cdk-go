package awsdevopsguru


// A reference to a NotificationChannel resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationChannelReference := &NotificationChannelReference{
//   	NotificationChannelId: jsii.String("notificationChannelId"),
//   }
//
type NotificationChannelReference struct {
	// The Id of the NotificationChannel resource.
	NotificationChannelId *string `field:"required" json:"notificationChannelId" yaml:"notificationChannelId"`
}

