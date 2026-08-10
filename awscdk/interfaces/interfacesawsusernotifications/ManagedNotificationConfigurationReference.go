package interfacesawsusernotifications


// A reference to a ManagedNotificationConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedNotificationConfigurationReference := &ManagedNotificationConfigurationReference{
//   	ManagedNotificationConfigurationArn: jsii.String("managedNotificationConfigurationArn"),
//   }
//
type ManagedNotificationConfigurationReference struct {
	// The Arn of the ManagedNotificationConfiguration resource.
	ManagedNotificationConfigurationArn *string `field:"required" json:"managedNotificationConfigurationArn" yaml:"managedNotificationConfigurationArn"`
}

