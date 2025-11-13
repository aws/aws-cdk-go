package interfacesawsnotifications


// A reference to a NotificationConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationConfigurationReference := &NotificationConfigurationReference{
//   	NotificationConfigurationArn: jsii.String("notificationConfigurationArn"),
//   }
//
type NotificationConfigurationReference struct {
	// The Arn of the NotificationConfiguration resource.
	NotificationConfigurationArn *string `field:"required" json:"notificationConfigurationArn" yaml:"notificationConfigurationArn"`
}

