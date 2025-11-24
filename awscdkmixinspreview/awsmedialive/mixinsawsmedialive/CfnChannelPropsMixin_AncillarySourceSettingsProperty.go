package mixinsawsmedialive


// Information about the ancillary captions to extract from the input.
//
// The parent of this entity is CaptionSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ancillarySourceSettingsProperty := &AncillarySourceSettingsProperty{
//   	SourceAncillaryChannelNumber: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-ancillarysourcesettings.html
//
type CfnChannelPropsMixin_AncillarySourceSettingsProperty struct {
	// Specifies the number (1 to 4) of the captions channel you want to extract from the ancillary captions.
	//
	// If you plan to convert the ancillary captions to another format, complete this field. If you plan to choose Embedded as the captions destination in the output (to pass through all the channels in the ancillary captions), leave this field blank because MediaLive ignores the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-ancillarysourcesettings.html#cfn-medialive-channel-ancillarysourcesettings-sourceancillarychannelnumber
	//
	SourceAncillaryChannelNumber *float64 `field:"optional" json:"sourceAncillaryChannelNumber" yaml:"sourceAncillaryChannelNumber"`
}

