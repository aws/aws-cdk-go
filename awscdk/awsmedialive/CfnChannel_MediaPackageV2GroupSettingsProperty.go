package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mediaPackageV2GroupSettingsProperty := &MediaPackageV2GroupSettingsProperty{
//   	AdditionalDestinations: []interface{}{
//   		&MediaPackageAdditionalDestinationsProperty{
//   			Destination: &OutputLocationRefProperty{
//   				DestinationRefId: jsii.String("destinationRefId"),
//   			},
//   		},
//   	},
//   	CaptionLanguageMappings: []interface{}{
//   		&CaptionLanguageMappingProperty{
//   			CaptionChannel: jsii.Number(123),
//   			LanguageCode: jsii.String("languageCode"),
//   			LanguageDescription: jsii.String("languageDescription"),
//   		},
//   	},
//   	Id3Behavior: jsii.String("id3Behavior"),
//   	KlvBehavior: jsii.String("klvBehavior"),
//   	NielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   	Scte35Type: jsii.String("scte35Type"),
//   	SegmentLength: jsii.Number(123),
//   	SegmentLengthUnits: jsii.String("segmentLengthUnits"),
//   	TimedMetadataId3Frame: jsii.String("timedMetadataId3Frame"),
//   	TimedMetadataId3Period: jsii.Number(123),
//   	TimedMetadataPassthrough: jsii.String("timedMetadataPassthrough"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackagev2groupsettings.html
//
type CfnChannel_MediaPackageV2GroupSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackagev2groupsettings.html#cfn-medialive-channel-mediapackagev2groupsettings-additionaldestinations
	//
	AdditionalDestinations interface{} `field:"optional" json:"additionalDestinations" yaml:"additionalDestinations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackagev2groupsettings.html#cfn-medialive-channel-mediapackagev2groupsettings-captionlanguagemappings
	//
	CaptionLanguageMappings interface{} `field:"optional" json:"captionLanguageMappings" yaml:"captionLanguageMappings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackagev2groupsettings.html#cfn-medialive-channel-mediapackagev2groupsettings-id3behavior
	//
	Id3Behavior *string `field:"optional" json:"id3Behavior" yaml:"id3Behavior"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackagev2groupsettings.html#cfn-medialive-channel-mediapackagev2groupsettings-klvbehavior
	//
	KlvBehavior *string `field:"optional" json:"klvBehavior" yaml:"klvBehavior"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackagev2groupsettings.html#cfn-medialive-channel-mediapackagev2groupsettings-nielsenid3behavior
	//
	NielsenId3Behavior *string `field:"optional" json:"nielsenId3Behavior" yaml:"nielsenId3Behavior"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackagev2groupsettings.html#cfn-medialive-channel-mediapackagev2groupsettings-scte35type
	//
	Scte35Type *string `field:"optional" json:"scte35Type" yaml:"scte35Type"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackagev2groupsettings.html#cfn-medialive-channel-mediapackagev2groupsettings-segmentlength
	//
	SegmentLength *float64 `field:"optional" json:"segmentLength" yaml:"segmentLength"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackagev2groupsettings.html#cfn-medialive-channel-mediapackagev2groupsettings-segmentlengthunits
	//
	SegmentLengthUnits *string `field:"optional" json:"segmentLengthUnits" yaml:"segmentLengthUnits"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackagev2groupsettings.html#cfn-medialive-channel-mediapackagev2groupsettings-timedmetadataid3frame
	//
	TimedMetadataId3Frame *string `field:"optional" json:"timedMetadataId3Frame" yaml:"timedMetadataId3Frame"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackagev2groupsettings.html#cfn-medialive-channel-mediapackagev2groupsettings-timedmetadataid3period
	//
	TimedMetadataId3Period *float64 `field:"optional" json:"timedMetadataId3Period" yaml:"timedMetadataId3Period"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackagev2groupsettings.html#cfn-medialive-channel-mediapackagev2groupsettings-timedmetadatapassthrough
	//
	TimedMetadataPassthrough *string `field:"optional" json:"timedMetadataPassthrough" yaml:"timedMetadataPassthrough"`
}

