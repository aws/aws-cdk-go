package previewawsmedialivemixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cmafIngestGroupSettingsProperty := &CmafIngestGroupSettingsProperty{
//   	AdditionalDestinations: []interface{}{
//   		&AdditionalDestinationsProperty{
//   			Destination: &OutputLocationRefProperty{
//   				DestinationRefId: jsii.String("destinationRefId"),
//   			},
//   		},
//   	},
//   	CaptionLanguageMappings: []interface{}{
//   		&CmafIngestCaptionLanguageMappingProperty{
//   			CaptionChannel: jsii.Number(123),
//   			LanguageCode: jsii.String("languageCode"),
//   		},
//   	},
//   	Destination: &OutputLocationRefProperty{
//   		DestinationRefId: jsii.String("destinationRefId"),
//   	},
//   	Id3Behavior: jsii.String("id3Behavior"),
//   	Id3NameModifier: jsii.String("id3NameModifier"),
//   	KlvBehavior: jsii.String("klvBehavior"),
//   	KlvNameModifier: jsii.String("klvNameModifier"),
//   	NielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   	NielsenId3NameModifier: jsii.String("nielsenId3NameModifier"),
//   	Scte35NameModifier: jsii.String("scte35NameModifier"),
//   	Scte35Type: jsii.String("scte35Type"),
//   	SegmentLength: jsii.Number(123),
//   	SegmentLengthUnits: jsii.String("segmentLengthUnits"),
//   	SendDelayMs: jsii.Number(123),
//   	TimedMetadataId3Frame: jsii.String("timedMetadataId3Frame"),
//   	TimedMetadataId3Period: jsii.Number(123),
//   	TimedMetadataPassthrough: jsii.String("timedMetadataPassthrough"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestgroupsettings.html
//
type CfnChannelPropsMixin_CmafIngestGroupSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestgroupsettings.html#cfn-medialive-channel-cmafingestgroupsettings-additionaldestinations
	//
	AdditionalDestinations interface{} `field:"optional" json:"additionalDestinations" yaml:"additionalDestinations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestgroupsettings.html#cfn-medialive-channel-cmafingestgroupsettings-captionlanguagemappings
	//
	CaptionLanguageMappings interface{} `field:"optional" json:"captionLanguageMappings" yaml:"captionLanguageMappings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestgroupsettings.html#cfn-medialive-channel-cmafingestgroupsettings-destination
	//
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestgroupsettings.html#cfn-medialive-channel-cmafingestgroupsettings-id3behavior
	//
	Id3Behavior *string `field:"optional" json:"id3Behavior" yaml:"id3Behavior"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestgroupsettings.html#cfn-medialive-channel-cmafingestgroupsettings-id3namemodifier
	//
	Id3NameModifier *string `field:"optional" json:"id3NameModifier" yaml:"id3NameModifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestgroupsettings.html#cfn-medialive-channel-cmafingestgroupsettings-klvbehavior
	//
	KlvBehavior *string `field:"optional" json:"klvBehavior" yaml:"klvBehavior"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestgroupsettings.html#cfn-medialive-channel-cmafingestgroupsettings-klvnamemodifier
	//
	KlvNameModifier *string `field:"optional" json:"klvNameModifier" yaml:"klvNameModifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestgroupsettings.html#cfn-medialive-channel-cmafingestgroupsettings-nielsenid3behavior
	//
	NielsenId3Behavior *string `field:"optional" json:"nielsenId3Behavior" yaml:"nielsenId3Behavior"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestgroupsettings.html#cfn-medialive-channel-cmafingestgroupsettings-nielsenid3namemodifier
	//
	NielsenId3NameModifier *string `field:"optional" json:"nielsenId3NameModifier" yaml:"nielsenId3NameModifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestgroupsettings.html#cfn-medialive-channel-cmafingestgroupsettings-scte35namemodifier
	//
	Scte35NameModifier *string `field:"optional" json:"scte35NameModifier" yaml:"scte35NameModifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestgroupsettings.html#cfn-medialive-channel-cmafingestgroupsettings-scte35type
	//
	Scte35Type *string `field:"optional" json:"scte35Type" yaml:"scte35Type"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestgroupsettings.html#cfn-medialive-channel-cmafingestgroupsettings-segmentlength
	//
	SegmentLength *float64 `field:"optional" json:"segmentLength" yaml:"segmentLength"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestgroupsettings.html#cfn-medialive-channel-cmafingestgroupsettings-segmentlengthunits
	//
	SegmentLengthUnits *string `field:"optional" json:"segmentLengthUnits" yaml:"segmentLengthUnits"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestgroupsettings.html#cfn-medialive-channel-cmafingestgroupsettings-senddelayms
	//
	SendDelayMs *float64 `field:"optional" json:"sendDelayMs" yaml:"sendDelayMs"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestgroupsettings.html#cfn-medialive-channel-cmafingestgroupsettings-timedmetadataid3frame
	//
	TimedMetadataId3Frame *string `field:"optional" json:"timedMetadataId3Frame" yaml:"timedMetadataId3Frame"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestgroupsettings.html#cfn-medialive-channel-cmafingestgroupsettings-timedmetadataid3period
	//
	TimedMetadataId3Period *float64 `field:"optional" json:"timedMetadataId3Period" yaml:"timedMetadataId3Period"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestgroupsettings.html#cfn-medialive-channel-cmafingestgroupsettings-timedmetadatapassthrough
	//
	TimedMetadataPassthrough *string `field:"optional" json:"timedMetadataPassthrough" yaml:"timedMetadataPassthrough"`
}

