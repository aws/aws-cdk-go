package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   anywhereSettingsProperty := &AnywhereSettingsProperty{
//   	ChannelPlacementGroupId: jsii.String("channelPlacementGroupId"),
//   	ClusterId: jsii.String("clusterId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-anywheresettings.html
//
type CfnChannel_AnywhereSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-anywheresettings.html#cfn-medialive-channel-anywheresettings-channelplacementgroupid
	//
	ChannelPlacementGroupId *string `field:"optional" json:"channelPlacementGroupId" yaml:"channelPlacementGroupId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-anywheresettings.html#cfn-medialive-channel-anywheresettings-clusterid
	//
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
}

