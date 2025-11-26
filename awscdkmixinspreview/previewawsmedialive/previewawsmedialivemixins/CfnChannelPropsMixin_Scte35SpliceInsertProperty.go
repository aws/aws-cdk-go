package previewawsmedialivemixins


// The setup of SCTE-35 splice insert handling.
//
// The parent of this entity is AvailSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   scte35SpliceInsertProperty := &Scte35SpliceInsertProperty{
//   	AdAvailOffset: jsii.Number(123),
//   	NoRegionalBlackoutFlag: jsii.String("noRegionalBlackoutFlag"),
//   	WebDeliveryAllowedFlag: jsii.String("webDeliveryAllowedFlag"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-scte35spliceinsert.html
//
type CfnChannelPropsMixin_Scte35SpliceInsertProperty struct {
	// When specified, this offset (in milliseconds) is added to the input ad avail PTS time.
	//
	// This applies only to embedded SCTE 104/35 messages. It doesn't apply to OOB messages.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-scte35spliceinsert.html#cfn-medialive-channel-scte35spliceinsert-adavailoffset
	//
	AdAvailOffset *float64 `field:"optional" json:"adAvailOffset" yaml:"adAvailOffset"`
	// When set to ignore, segment descriptors with noRegionalBlackoutFlag set to 0 no longer trigger blackouts or ad avail slates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-scte35spliceinsert.html#cfn-medialive-channel-scte35spliceinsert-noregionalblackoutflag
	//
	NoRegionalBlackoutFlag *string `field:"optional" json:"noRegionalBlackoutFlag" yaml:"noRegionalBlackoutFlag"`
	// When set to ignore, segment descriptors with webDeliveryAllowedFlag set to 0 no longer trigger blackouts or ad avail slates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-scte35spliceinsert.html#cfn-medialive-channel-scte35spliceinsert-webdeliveryallowedflag
	//
	WebDeliveryAllowedFlag *string `field:"optional" json:"webDeliveryAllowedFlag" yaml:"webDeliveryAllowedFlag"`
}

