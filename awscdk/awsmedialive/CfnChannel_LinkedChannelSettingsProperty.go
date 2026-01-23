package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   linkedChannelSettingsProperty := &LinkedChannelSettingsProperty{
//   	FollowerChannelSettings: &FollowerChannelSettingsProperty{
//   		LinkedChannelType: jsii.String("linkedChannelType"),
//   		PrimaryChannelArn: jsii.String("primaryChannelArn"),
//   	},
//   	PrimaryChannelSettings: &PrimaryChannelSettingsProperty{
//   		LinkedChannelType: jsii.String("linkedChannelType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-linkedchannelsettings.html
//
type CfnChannel_LinkedChannelSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-linkedchannelsettings.html#cfn-medialive-channel-linkedchannelsettings-followerchannelsettings
	//
	FollowerChannelSettings interface{} `field:"optional" json:"followerChannelSettings" yaml:"followerChannelSettings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-linkedchannelsettings.html#cfn-medialive-channel-linkedchannelsettings-primarychannelsettings
	//
	PrimaryChannelSettings interface{} `field:"optional" json:"primaryChannelSettings" yaml:"primaryChannelSettings"`
}

