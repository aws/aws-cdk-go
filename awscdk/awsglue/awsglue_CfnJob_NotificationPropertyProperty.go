package awsglue


// Specifies configuration properties of a notification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationPropertyProperty := &notificationPropertyProperty{
//   	notifyDelayAfter: jsii.Number(123),
//   }
//
type CfnJob_NotificationPropertyProperty struct {
	// After a job run starts, the number of minutes to wait before sending a job run delay notification.
	NotifyDelayAfter *float64 `field:"optional" json:"notifyDelayAfter" yaml:"notifyDelayAfter"`
}

