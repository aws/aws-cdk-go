package awschime

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnChannelProps := &CfnChannelProps{
//   	AppInstanceArn: jsii.String("appInstanceArn"),
//   	ChimeBearer: jsii.String("chimeBearer"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	ChannelId: jsii.String("channelId"),
//   	ElasticChannelConfiguration: &ElasticChannelConfigurationProperty{
//   		MaximumSubChannels: jsii.Number(123),
//   		MinimumMembershipPercentage: jsii.Number(123),
//   		TargetMembershipsPerSubChannel: jsii.Number(123),
//   	},
//   	ExpirationSettings: &ExpirationSettingsProperty{
//   		ExpirationCriterion: jsii.String("expirationCriterion"),
//   		ExpirationDays: jsii.Number(123),
//   	},
//   	MemberArns: []*string{
//   		jsii.String("memberArns"),
//   	},
//   	Metadata: jsii.String("metadata"),
//   	Mode: jsii.String("mode"),
//   	ModeratorArns: []*string{
//   		jsii.String("moderatorArns"),
//   	},
//   	Privacy: jsii.String("privacy"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-channel.html
//
type CfnChannelProps struct {
	// The ARN of the AppInstance that contains the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-channel.html#cfn-chime-channel-appinstancearn
	//
	AppInstanceArn *string `field:"required" json:"appInstanceArn" yaml:"appInstanceArn"`
	// The ARN of the AppInstanceUser or AppInstanceBot that performs every operation on this channel.
	//
	// Whichever of the two creates a channel automatically becomes one of its moderators, so the same ARN can subsequently read, update and delete the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-channel.html#cfn-chime-channel-chimebearer
	//
	ChimeBearer *string `field:"required" json:"chimeBearer" yaml:"chimeBearer"`
	// The name of the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-channel.html#cfn-chime-channel-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the channel.
	//
	// When omitted, the service generates a UUID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-channel.html#cfn-chime-channel-channelid
	//
	ChannelId *string `field:"optional" json:"channelId" yaml:"channelId"`
	// The attributes required to configure and create an elastic channel.
	//
	// An elastic channel must use RESTRICTED mode, cannot be created with MemberArns, and is available only in some regions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-channel.html#cfn-chime-channel-elasticchannelconfiguration
	//
	ElasticChannelConfiguration interface{} `field:"optional" json:"elasticChannelConfiguration" yaml:"elasticChannelConfiguration"`
	// Settings that control the interval after which the channel is automatically deleted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-channel.html#cfn-chime-channel-expirationsettings
	//
	ExpirationSettings interface{} `field:"optional" json:"expirationSettings" yaml:"expirationSettings"`
	// The ARNs of the AppInstanceUsers to add to the channel as members when it is created.
	//
	// Cannot be combined with ElasticChannelConfiguration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-channel.html#cfn-chime-channel-memberarns
	//
	MemberArns *[]*string `field:"optional" json:"memberArns" yaml:"memberArns"`
	// The metadata of the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-channel.html#cfn-chime-channel-metadata
	//
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
	// The channel mode.
	//
	// In an UNRESTRICTED channel, members can add themselves and other members; in a RESTRICTED channel, only administrators and moderators can add members. An elastic channel must be RESTRICTED.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-channel.html#cfn-chime-channel-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The ARNs of the AppInstanceUsers to add to the channel as moderators when it is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-channel.html#cfn-chime-channel-moderatorarns
	//
	ModeratorArns *[]*string `field:"optional" json:"moderatorArns" yaml:"moderatorArns"`
	// The channel's privacy level.
	//
	// A PUBLIC channel is discoverable by anyone in the AppInstance; a PRIVATE channel is not. Privacy cannot be changed after creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-channel.html#cfn-chime-channel-privacy
	//
	Privacy *string `field:"optional" json:"privacy" yaml:"privacy"`
	// The tags for the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-channel.html#cfn-chime-channel-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

