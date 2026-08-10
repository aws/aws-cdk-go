package previewawstranscribeevents


// Type definition for TranscriptionJob.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   transcriptionJob := &TranscriptionJob{
//   	CreationTime: []*string{
//   		jsii.String("creationTime"),
//   	},
//   	LanguageCode: []*string{
//   		jsii.String("languageCode"),
//   	},
//   	Media: &Media1{
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
//   	Settings: &Settings1{
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
//   	TranscriptionJobStatus: []*string{
//   		jsii.String("transcriptionJobStatus"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_TranscriptionJob struct {
	// creationTime property.
	//
	// Specify an array of string values to match this event if the actual value of creationTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CreationTime *[]*string `field:"optional" json:"creationTime" yaml:"creationTime"`
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
	Media *AWSAPICallViaCloudTrail_Media1 `field:"optional" json:"media" yaml:"media"`
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
	Settings *AWSAPICallViaCloudTrail_Settings1 `field:"optional" json:"settings" yaml:"settings"`
	// transcriptionJobName property.
	//
	// Specify an array of string values to match this event if the actual value of transcriptionJobName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TranscriptionJobName *[]*string `field:"optional" json:"transcriptionJobName" yaml:"transcriptionJobName"`
	// transcriptionJobStatus property.
	//
	// Specify an array of string values to match this event if the actual value of transcriptionJobStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TranscriptionJobStatus *[]*string `field:"optional" json:"transcriptionJobStatus" yaml:"transcriptionJobStatus"`
}

