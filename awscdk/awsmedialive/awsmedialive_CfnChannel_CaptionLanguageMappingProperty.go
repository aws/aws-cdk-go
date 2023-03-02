package awsmedialive


// Maps a captions channel to an ISO 693-2 language code (http://www.loc.gov/standards/iso639-2), with an optional description.
//
// The parent of this entity is HlsGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   captionLanguageMappingProperty := &captionLanguageMappingProperty{
//   	captionChannel: jsii.Number(123),
//   	languageCode: jsii.String("languageCode"),
//   	languageDescription: jsii.String("languageDescription"),
//   }
//
type CfnChannel_CaptionLanguageMappingProperty struct {
	// The closed caption channel being described by this CaptionLanguageMapping.
	//
	// Each channel mapping must have a unique channel number (maximum of 4).
	CaptionChannel *float64 `field:"optional" json:"captionChannel" yaml:"captionChannel"`
	// A three-character ISO 639-2 language code (see http://www.loc.gov/standards/iso639-2).
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// The textual description of language.
	LanguageDescription *string `field:"optional" json:"languageDescription" yaml:"languageDescription"`
}

