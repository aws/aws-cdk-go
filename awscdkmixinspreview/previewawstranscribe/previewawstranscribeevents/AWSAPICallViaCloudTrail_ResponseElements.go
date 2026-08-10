package previewawstranscribeevents


// Type definition for ResponseElements.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   responseElements := &ResponseElements{
//   	LanguageCode: []*string{
//   		jsii.String("languageCode"),
//   	},
//   	TranscriptionJob: &TranscriptionJob{
//   		CreationTime: []*string{
//   			jsii.String("creationTime"),
//   		},
//   		LanguageCode: []*string{
//   			jsii.String("languageCode"),
//   		},
//   		Media: &Media1{
//   			MediaFileUri: []*string{
//   				jsii.String("mediaFileUri"),
//   			},
//   		},
//   		MediaFormat: []*string{
//   			jsii.String("mediaFormat"),
//   		},
//   		MediaSampleRateHertz: []*string{
//   			jsii.String("mediaSampleRateHertz"),
//   		},
//   		Settings: &Settings1{
//   			ChannelIdentification: []*string{
//   				jsii.String("channelIdentification"),
//   			},
//   			VocabularyName: []*string{
//   				jsii.String("vocabularyName"),
//   			},
//   		},
//   		TranscriptionJobName: []*string{
//   			jsii.String("transcriptionJobName"),
//   		},
//   		TranscriptionJobStatus: []*string{
//   			jsii.String("transcriptionJobStatus"),
//   		},
//   	},
//   	VocabularyName: []*string{
//   		jsii.String("vocabularyName"),
//   	},
//   	VocabularyState: []*string{
//   		jsii.String("vocabularyState"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_ResponseElements struct {
	// languageCode property.
	//
	// Specify an array of string values to match this event if the actual value of languageCode is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LanguageCode *[]*string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// transcriptionJob property.
	//
	// Specify an array of string values to match this event if the actual value of transcriptionJob is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TranscriptionJob *AWSAPICallViaCloudTrail_TranscriptionJob `field:"optional" json:"transcriptionJob" yaml:"transcriptionJob"`
	// vocabularyName property.
	//
	// Specify an array of string values to match this event if the actual value of vocabularyName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VocabularyName *[]*string `field:"optional" json:"vocabularyName" yaml:"vocabularyName"`
	// vocabularyState property.
	//
	// Specify an array of string values to match this event if the actual value of vocabularyState is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VocabularyState *[]*string `field:"optional" json:"vocabularyState" yaml:"vocabularyState"`
}

