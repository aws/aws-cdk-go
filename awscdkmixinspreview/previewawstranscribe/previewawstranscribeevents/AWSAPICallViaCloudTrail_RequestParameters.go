package previewawstranscribeevents


// Type definition for RequestParameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   requestParameters := &RequestParameters{
//   	LanguageCode: []*string{
//   		jsii.String("languageCode"),
//   	},
//   	Media: &Media{
//   		MediaFileUri: []*string{
//   			jsii.String("mediaFileUri"),
//   		},
//   	},
//   	MediaFormat: []*string{
//   		jsii.String("mediaFormat"),
//   	},
//   	MediaSampleRateHertz: []*string{
//   		jsii.String("mediaSampleRateHertz"),
//   	},
//   	Settings: &Settings{
//   		ChannelIdentification: []*string{
//   			jsii.String("channelIdentification"),
//   		},
//   		VocabularyName: []*string{
//   			jsii.String("vocabularyName"),
//   		},
//   	},
//   	TranscriptionJobName: []*string{
//   		jsii.String("transcriptionJobName"),
//   	},
//   	VocabularyName: []*string{
//   		jsii.String("vocabularyName"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_RequestParameters struct {
	// languageCode property.
	//
	// Specify an array of string values to match this event if the actual value of languageCode is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LanguageCode *[]*string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// media property.
	//
	// Specify an array of string values to match this event if the actual value of media is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Media *AWSAPICallViaCloudTrail_Media `field:"optional" json:"media" yaml:"media"`
	// mediaFormat property.
	//
	// Specify an array of string values to match this event if the actual value of mediaFormat is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MediaFormat *[]*string `field:"optional" json:"mediaFormat" yaml:"mediaFormat"`
	// mediaSampleRateHertz property.
	//
	// Specify an array of string values to match this event if the actual value of mediaSampleRateHertz is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MediaSampleRateHertz *[]*string `field:"optional" json:"mediaSampleRateHertz" yaml:"mediaSampleRateHertz"`
	// settings property.
	//
	// Specify an array of string values to match this event if the actual value of settings is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Settings *AWSAPICallViaCloudTrail_Settings `field:"optional" json:"settings" yaml:"settings"`
	// transcriptionJobName property.
	//
	// Specify an array of string values to match this event if the actual value of transcriptionJobName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TranscriptionJobName *[]*string `field:"optional" json:"transcriptionJobName" yaml:"transcriptionJobName"`
	// vocabularyName property.
	//
	// Specify an array of string values to match this event if the actual value of vocabularyName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VocabularyName *[]*string `field:"optional" json:"vocabularyName" yaml:"vocabularyName"`
}

