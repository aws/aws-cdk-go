package awstranscribe

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnMedicalTranscriptionJob`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMedicalTranscriptionJobProps := &CfnMedicalTranscriptionJobProps{
//   	LanguageCode: jsii.String("languageCode"),
//   	Media: &MediaProperty{
//   		MediaFileUri: jsii.String("mediaFileUri"),
//   	},
//   	MedicalTranscriptionJobName: jsii.String("medicalTranscriptionJobName"),
//   	Specialty: jsii.String("specialty"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	MediaFormat: jsii.String("mediaFormat"),
//   	MediaSampleRateHertz: jsii.Number(123),
//   	Settings: &MedicalTranscriptionSettingProperty{
//   		ChannelIdentification: jsii.Boolean(false),
//   		ShowAlternatives: jsii.Boolean(false),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicaltranscriptionjob.html
//
type CfnMedicalTranscriptionJobProps struct {
	// The language code for the language spoken in the input media file.
	//
	// Must be en-US.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicaltranscriptionjob.html#cfn-transcribe-medicaltranscriptionjob-languagecode
	//
	LanguageCode *string `field:"required" json:"languageCode" yaml:"languageCode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicaltranscriptionjob.html#cfn-transcribe-medicaltranscriptionjob-media
	//
	Media interface{} `field:"required" json:"media" yaml:"media"`
	// A unique name for the medical transcription job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicaltranscriptionjob.html#cfn-transcribe-medicaltranscriptionjob-medicaltranscriptionjobname
	//
	MedicalTranscriptionJobName *string `field:"required" json:"medicalTranscriptionJobName" yaml:"medicalTranscriptionJobName"`
	// The medical specialty represented in the media.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicaltranscriptionjob.html#cfn-transcribe-medicaltranscriptionjob-specialty
	//
	Specialty *string `field:"required" json:"specialty" yaml:"specialty"`
	// Whether the input media is a dictation or conversation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicaltranscriptionjob.html#cfn-transcribe-medicaltranscriptionjob-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The format of the input media file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicaltranscriptionjob.html#cfn-transcribe-medicaltranscriptionjob-mediaformat
	//
	MediaFormat *string `field:"optional" json:"mediaFormat" yaml:"mediaFormat"`
	// The sample rate of the audio in hertz.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicaltranscriptionjob.html#cfn-transcribe-medicaltranscriptionjob-mediasampleratehertz
	//
	MediaSampleRateHertz *float64 `field:"optional" json:"mediaSampleRateHertz" yaml:"mediaSampleRateHertz"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicaltranscriptionjob.html#cfn-transcribe-medicaltranscriptionjob-settings
	//
	Settings interface{} `field:"optional" json:"settings" yaml:"settings"`
	// Tags associated with the medical transcription job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicaltranscriptionjob.html#cfn-transcribe-medicaltranscriptionjob-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

