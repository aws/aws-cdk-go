package interfacesawsnotifications


// A reference to a ManagedNotificationAccountContactAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedNotificationAccountContactAssociationReference := &ManagedNotificationAccountContactAssociationReference{
//   	ContactIdentifier: jsii.String("contactIdentifier"),
//   	ManagedNotificationConfigurationArn: jsii.String("managedNotificationConfigurationArn"),
//   }
//
type ManagedNotificationAccountContactAssociationReference struct {
	// The ContactIdentifier of the ManagedNotificationAccountContactAssociation resource.
	ContactIdentifier *string `field:"required" json:"contactIdentifier" yaml:"contactIdentifier"`
	// The ManagedNotificationConfigurationArn of the ManagedNotificationAccountContactAssociation resource.
	ManagedNotificationConfigurationArn *string `field:"required" json:"managedNotificationConfigurationArn" yaml:"managedNotificationConfigurationArn"`
}

