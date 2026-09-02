package awsmedialivealpha


// SCTE-35 time signal APOS avail settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var scte35FlagBehavior Scte35FlagBehavior
//
//   scte35TimeSignalAposSettings := &Scte35TimeSignalAposSettings{
//   	AdAvailOffset: jsii.Number(123),
//   	NoRegionalBlackoutFlag: scte35FlagBehavior,
//   	WebDeliveryAllowedFlag: scte35FlagBehavior,
//   }
//
// Experimental.
type Scte35TimeSignalAposSettings struct {
	// Offset in milliseconds added to the input ad avail PTS time.
	//
	// Applies only to embedded SCTE 104/35 messages.
	// Default: - service default.
	//
	// Experimental.
	AdAvailOffset *float64 `field:"optional" json:"adAvailOffset" yaml:"adAvailOffset"`
	// When set to `IGNORE`, segment descriptors with `noRegionalBlackoutFlag` set to 0 no longer trigger blackouts or ad avail slates.
	// Default: - service default.
	//
	// Experimental.
	NoRegionalBlackoutFlag Scte35FlagBehavior `field:"optional" json:"noRegionalBlackoutFlag" yaml:"noRegionalBlackoutFlag"`
	// When set to `IGNORE`, segment descriptors with `webDeliveryAllowedFlag` set to 0 no longer trigger blackouts or ad avail slates.
	// Default: - service default.
	//
	// Experimental.
	WebDeliveryAllowedFlag Scte35FlagBehavior `field:"optional" json:"webDeliveryAllowedFlag" yaml:"webDeliveryAllowedFlag"`
}

