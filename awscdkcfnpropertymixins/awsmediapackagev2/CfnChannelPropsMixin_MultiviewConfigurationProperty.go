package awsmediapackagev2


// <p>The multiview configuration for a channel.
//
// A multiview channel composites video from several source channels into a single tiled output stream. Players receive one standard HLS or DASH stream instead of several separate streams. This setting is required when <code>InputType</code> is <code>MULTIVIEW</code>, and can't be set for any other input type.</p>
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   multiviewConfigurationProperty := &MultiviewConfigurationProperty{
//   	AvailableLayouts: []*string{
//   		jsii.String("availableLayouts"),
//   	},
//   	AvailableSources: []*string{
//   		jsii.String("availableSources"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-channel-multiviewconfiguration.html
//
type CfnChannelPropsMixin_MultiviewConfigurationProperty struct {
	// <p>The tile layouts that players can request from this multiview channel's origin endpoints.
	//
	// Only the layouts that you list here are available. Each layout must appear at most once.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-channel-multiviewconfiguration.html#cfn-mediapackagev2-channel-multiviewconfiguration-availablelayouts
	//
	AvailableLayouts *[]*string `field:"optional" json:"availableLayouts" yaml:"availableLayouts"`
	// <p>The channels that players can use as tiles in this multiview channel's output.
	//
	// Each source channel must be in the same channel group as the multiview channel, and must have an <code>InputType</code> of <code>CMAF</code>. Only the channels that you list here are available as tiles.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-channel-multiviewconfiguration.html#cfn-mediapackagev2-channel-multiviewconfiguration-availablesources
	//
	AvailableSources *[]*string `field:"optional" json:"availableSources" yaml:"availableSources"`
}

