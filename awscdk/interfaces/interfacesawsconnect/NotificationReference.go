package interfacesawsconnect


// A reference to a Notification resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationReference := &NotificationReference{
//   	NotificationArn: jsii.String("notificationArn"),
//   }
//
type NotificationReference struct {
	// The Arn of the Notification resource.
	NotificationArn *string `field:"required" json:"notificationArn" yaml:"notificationArn"`
}

