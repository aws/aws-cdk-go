package awsmedialivealpha


// Options for an embedded (CEA-608/708) caption source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var convert608To708 Convert608To708
//   var scte20Detection Scte20Detection
//
//   embeddedCaptionSourceOptions := &EmbeddedCaptionSourceOptions{
//   	Convert608To708: convert608To708,
//   	Scte20Detection: scte20Detection,
//   	Source608ChannelNumber: jsii.Number(123),
//   }
//
// Experimental.
type EmbeddedCaptionSourceOptions struct {
	// Whether to upconvert 608 captions to 708.
	// Default: Convert608To708.DISABLED
	//
	// Experimental.
	Convert608To708 Convert608To708 `field:"optional" json:"convert608To708" yaml:"convert608To708"`
	// SCTE-20 detection mode.
	// Default: Scte20Detection.OFF
	//
	// Experimental.
	Scte20Detection Scte20Detection `field:"optional" json:"scte20Detection" yaml:"scte20Detection"`
	// The 608/708 channel number within the video track to extract captions from.
	// Default: - MediaLive service default.
	//
	// Experimental.
	Source608ChannelNumber *float64 `field:"optional" json:"source608ChannelNumber" yaml:"source608ChannelNumber"`
}

