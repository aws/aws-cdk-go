package awschime


// The attributes required to configure and create an elastic channel.
//
// An elastic channel must use RESTRICTED mode, cannot be created with MemberArns, and is available only in some regions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   elasticChannelConfigurationProperty := &ElasticChannelConfigurationProperty{
//   	MaximumSubChannels: jsii.Number(123),
//   	MinimumMembershipPercentage: jsii.Number(123),
//   	TargetMembershipsPerSubChannel: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-channel-elasticchannelconfiguration.html
//
type CfnChannel_ElasticChannelConfigurationProperty struct {
	// The maximum number of SubChannels allowed in the elastic channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-channel-elasticchannelconfiguration.html#cfn-chime-channel-elasticchannelconfiguration-maximumsubchannels
	//
	MaximumSubChannels *float64 `field:"required" json:"maximumSubChannels" yaml:"maximumSubChannels"`
	// The minimum allowed percentage of TargetMembershipsPerSubChannel users, used to balance members across SubChannels.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-channel-elasticchannelconfiguration.html#cfn-chime-channel-elasticchannelconfiguration-minimummembershippercentage
	//
	MinimumMembershipPercentage *float64 `field:"required" json:"minimumMembershipPercentage" yaml:"minimumMembershipPercentage"`
	// The maximum number of members allowed in a SubChannel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-channel-elasticchannelconfiguration.html#cfn-chime-channel-elasticchannelconfiguration-targetmembershipspersubchannel
	//
	TargetMembershipsPerSubChannel *float64 `field:"required" json:"targetMembershipsPerSubChannel" yaml:"targetMembershipsPerSubChannel"`
}

