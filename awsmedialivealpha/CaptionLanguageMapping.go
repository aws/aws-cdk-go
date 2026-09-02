package awsmedialivealpha


// Maps a captions channel to an ISO 693-2 language code.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   captionLanguageMapping := &CaptionLanguageMapping{
//   	CaptionChannel: jsii.Number(123),
//   	LanguageCode: jsii.String("languageCode"),
//   	LanguageDescription: jsii.String("languageDescription"),
//   }
//
// Experimental.
type CaptionLanguageMapping struct {
	// The closed caption channel number (1-4).
	// Experimental.
	CaptionChannel *float64 `field:"required" json:"captionChannel" yaml:"captionChannel"`
	// A three-character ISO 639-2 language code.
	// Experimental.
	LanguageCode *string `field:"required" json:"languageCode" yaml:"languageCode"`
	// The textual description of the language.
	// Experimental.
	LanguageDescription *string `field:"required" json:"languageDescription" yaml:"languageDescription"`
}

