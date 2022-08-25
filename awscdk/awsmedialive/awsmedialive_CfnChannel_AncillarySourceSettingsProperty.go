package awsmedialive


// Information about the ancillary captions to extract from the input.
//
// The parent of this entity is CaptionSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ancillarySourceSettingsProperty := &ancillarySourceSettingsProperty{
//   	sourceAncillaryChannelNumber: jsii.Number(123),
//   }
//
type CfnChannel_AncillarySourceSettingsProperty struct {
	// Specifies the number (1 to 4) of the captions channel you want to extract from the ancillary captions.
	//
	// If you plan to convert the ancillary captions to another format, complete this field. If you plan to choose Embedded as the captions destination in the output (to pass through all the channels in the ancillary captions), leave this field blank because MediaLive ignores the field.
	SourceAncillaryChannelNumber *float64 `field:"optional" json:"sourceAncillaryChannelNumber" yaml:"sourceAncillaryChannelNumber"`
}

