package previewawsmedialivemixins


// Maps a captions channel to an ISO 693-2 language code (http://www.loc.gov/standards/iso639-2), with an optional description.
//
// The parent of this entity is HlsGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   captionLanguageMappingProperty := &CaptionLanguageMappingProperty{
//   	CaptionChannel: jsii.Number(123),
//   	LanguageCode: jsii.String("languageCode"),
//   	LanguageDescription: jsii.String("languageDescription"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captionlanguagemapping.html
//
type CfnChannelPropsMixin_CaptionLanguageMappingProperty struct {
	// The closed caption channel being described by this CaptionLanguageMapping.
	//
	// Each channel mapping must have a unique channel number (maximum of 4).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captionlanguagemapping.html#cfn-medialive-channel-captionlanguagemapping-captionchannel
	//
	CaptionChannel *float64 `field:"optional" json:"captionChannel" yaml:"captionChannel"`
	// A three-character ISO 639-2 language code (see http://www.loc.gov/standards/iso639-2).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captionlanguagemapping.html#cfn-medialive-channel-captionlanguagemapping-languagecode
	//
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// The textual description of language.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captionlanguagemapping.html#cfn-medialive-channel-captionlanguagemapping-languagedescription
	//
	LanguageDescription *string `field:"optional" json:"languageDescription" yaml:"languageDescription"`
}

