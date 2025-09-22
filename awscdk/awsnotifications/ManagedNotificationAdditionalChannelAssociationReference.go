package awsnotifications


// A reference to a ManagedNotificationAdditionalChannelAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedNotificationAdditionalChannelAssociationReference := &ManagedNotificationAdditionalChannelAssociationReference{
//   	ChannelArn: jsii.String("channelArn"),
//   	ManagedNotificationConfigurationArn: jsii.String("managedNotificationConfigurationArn"),
//   }
//
type ManagedNotificationAdditionalChannelAssociationReference struct {
	// The ChannelArn of the ManagedNotificationAdditionalChannelAssociation resource.
	ChannelArn *string `field:"required" json:"channelArn" yaml:"channelArn"`
	// The ManagedNotificationConfigurationArn of the ManagedNotificationAdditionalChannelAssociation resource.
	ManagedNotificationConfigurationArn *string `field:"required" json:"managedNotificationConfigurationArn" yaml:"managedNotificationConfigurationArn"`
}

