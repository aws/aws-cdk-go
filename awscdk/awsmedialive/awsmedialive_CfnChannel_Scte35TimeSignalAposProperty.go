package awsmedialive


// The settings for the SCTE-35 time signal APOS mode.
//
// The parent of this entity is AvailSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scte35TimeSignalAposProperty := &scte35TimeSignalAposProperty{
//   	adAvailOffset: jsii.Number(123),
//   	noRegionalBlackoutFlag: jsii.String("noRegionalBlackoutFlag"),
//   	webDeliveryAllowedFlag: jsii.String("webDeliveryAllowedFlag"),
//   }
//
type CfnChannel_Scte35TimeSignalAposProperty struct {
	// When specified, this offset (in milliseconds) is added to the input ad avail PTS time.
	//
	// This applies only to embedded SCTE 104/35 messages. It doesn't apply to OOB messages.
	AdAvailOffset *float64 `field:"optional" json:"adAvailOffset" yaml:"adAvailOffset"`
	// When set to ignore, segment descriptors with noRegionalBlackoutFlag set to 0 no longer trigger blackouts or ad avail slates.
	NoRegionalBlackoutFlag *string `field:"optional" json:"noRegionalBlackoutFlag" yaml:"noRegionalBlackoutFlag"`
	// When set to ignore, segment descriptors with webDeliveryAllowedFlag set to 0 no longer trigger blackouts or ad avail slates.
	WebDeliveryAllowedFlag *string `field:"optional" json:"webDeliveryAllowedFlag" yaml:"webDeliveryAllowedFlag"`
}

