package awsmedialivealpha


// Maps a captions channel to an ISO 639-2 language code for a CMAF Ingest output group.
//
// Unlike `CaptionLanguageMapping`, the CMAF Ingest variant has no language description.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   cmafCaptionLanguageMapping := &CmafCaptionLanguageMapping{
//   	CaptionChannel: jsii.Number(123),
//   	LanguageCode: jsii.String("languageCode"),
//   }
//
// Experimental.
type CmafCaptionLanguageMapping struct {
	// The closed caption channel number (1-4).
	// Experimental.
	CaptionChannel *float64 `field:"required" json:"captionChannel" yaml:"captionChannel"`
	// A three-character ISO 639-2 language code.
	// Experimental.
	LanguageCode *string `field:"required" json:"languageCode" yaml:"languageCode"`
}

