package awsnotifications


// Properties for defining a `CfnChannelAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnChannelAssociationProps := &CfnChannelAssociationProps{
//   	Arn: jsii.String("arn"),
//   	NotificationConfigurationArn: jsii.String("notificationConfigurationArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-channelassociation.html
//
type CfnChannelAssociationProps struct {
	// The Amazon Resource Name (ARN) of the `Channel` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-channelassociation.html#cfn-notifications-channelassociation-arn
	//
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The ARN of the `NotificationConfiguration` associated with the `Channel` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-channelassociation.html#cfn-notifications-channelassociation-notificationconfigurationarn
	//
	NotificationConfigurationArn *string `field:"required" json:"notificationConfigurationArn" yaml:"notificationConfigurationArn"`
}

