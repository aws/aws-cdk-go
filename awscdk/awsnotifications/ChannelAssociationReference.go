package awsnotifications


// A reference to a ChannelAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   channelAssociationReference := &ChannelAssociationReference{
//   	ChannelAssociationArn: jsii.String("channelAssociationArn"),
//   	NotificationConfigurationArn: jsii.String("notificationConfigurationArn"),
//   }
//
type ChannelAssociationReference struct {
	// The Arn of the ChannelAssociation resource.
	ChannelAssociationArn *string `field:"required" json:"channelAssociationArn" yaml:"channelAssociationArn"`
	// The NotificationConfigurationArn of the ChannelAssociation resource.
	NotificationConfigurationArn *string `field:"required" json:"notificationConfigurationArn" yaml:"notificationConfigurationArn"`
}

