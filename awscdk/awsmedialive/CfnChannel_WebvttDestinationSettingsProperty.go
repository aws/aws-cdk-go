package awsmedialive


// The configuration of Web VTT captions in the output.
//
// The parent of this entity is CaptionDestinationSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webvttDestinationSettingsProperty := &WebvttDestinationSettingsProperty{
//   	StyleControl: jsii.String("styleControl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-webvttdestinationsettings.html
//
type CfnChannel_WebvttDestinationSettingsProperty struct {
	// Controls whether the color and position of the source captions is passed through to the WebVTT output captions.
	//
	// PASSTHROUGH - Valid only if the source captions are EMBEDDED or TELETEXT. NO_STYLE_DATA - Don't pass through the style. The output captions will not contain any font styling information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-webvttdestinationsettings.html#cfn-medialive-channel-webvttdestinationsettings-stylecontrol
	//
	StyleControl *string `field:"optional" json:"styleControl" yaml:"styleControl"`
}

