package mixinsawsmedialive


// Destination settings for a MediaPackage output.
//
// The parent of this entity is OutputDestination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mediaPackageOutputDestinationSettingsProperty := &MediaPackageOutputDestinationSettingsProperty{
//   	ChannelGroup: jsii.String("channelGroup"),
//   	ChannelId: jsii.String("channelId"),
//   	ChannelName: jsii.String("channelName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackageoutputdestinationsettings.html
//
type CfnChannelPropsMixin_MediaPackageOutputDestinationSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackageoutputdestinationsettings.html#cfn-medialive-channel-mediapackageoutputdestinationsettings-channelgroup
	//
	ChannelGroup *string `field:"optional" json:"channelGroup" yaml:"channelGroup"`
	// The ID of the channel in MediaPackage that is the destination for this output group.
	//
	// You don't need to specify the individual inputs in MediaPackage; MediaLive handles the connection of the two MediaLive pipelines to the two MediaPackage inputs. The MediaPackage channel and MediaLive channel must be in the same Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackageoutputdestinationsettings.html#cfn-medialive-channel-mediapackageoutputdestinationsettings-channelid
	//
	ChannelId *string `field:"optional" json:"channelId" yaml:"channelId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackageoutputdestinationsettings.html#cfn-medialive-channel-mediapackageoutputdestinationsettings-channelname
	//
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
}

