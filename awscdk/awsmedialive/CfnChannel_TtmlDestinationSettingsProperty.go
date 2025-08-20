package awsmedialive


// The setup of TTML captions in the output.
//
// The parent of this entity is CaptionDestinationSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ttmlDestinationSettingsProperty := &TtmlDestinationSettingsProperty{
//   	StyleControl: jsii.String("styleControl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-ttmldestinationsettings.html
//
type CfnChannel_TtmlDestinationSettingsProperty struct {
	// When set to passthrough, passes through style and position information from a TTML-like input source (TTML, SMPTE-TT, CFF-TT) to the CFF-TT output or TTML output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-ttmldestinationsettings.html#cfn-medialive-channel-ttmldestinationsettings-stylecontrol
	//
	StyleControl *string `field:"optional" json:"styleControl" yaml:"styleControl"`
}

