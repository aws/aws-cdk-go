package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   followerChannelSettingsProperty := &FollowerChannelSettingsProperty{
//   	LinkedChannelType: jsii.String("linkedChannelType"),
//   	PrimaryChannelArn: jsii.String("primaryChannelArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-followerchannelsettings.html
//
type CfnChannel_FollowerChannelSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-followerchannelsettings.html#cfn-medialive-channel-followerchannelsettings-linkedchanneltype
	//
	LinkedChannelType *string `field:"optional" json:"linkedChannelType" yaml:"linkedChannelType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-followerchannelsettings.html#cfn-medialive-channel-followerchannelsettings-primarychannelarn
	//
	PrimaryChannelArn *string `field:"optional" json:"primaryChannelArn" yaml:"primaryChannelArn"`
}

