package interfacesawsnotifications


// A reference to a OrganizationalUnitAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   organizationalUnitAssociationReference := &OrganizationalUnitAssociationReference{
//   	NotificationConfigurationArn: jsii.String("notificationConfigurationArn"),
//   	OrganizationalUnitId: jsii.String("organizationalUnitId"),
//   }
//
type OrganizationalUnitAssociationReference struct {
	// The NotificationConfigurationArn of the OrganizationalUnitAssociation resource.
	NotificationConfigurationArn *string `field:"required" json:"notificationConfigurationArn" yaml:"notificationConfigurationArn"`
	// The OrganizationalUnitId of the OrganizationalUnitAssociation resource.
	OrganizationalUnitId *string `field:"required" json:"organizationalUnitId" yaml:"organizationalUnitId"`
}

