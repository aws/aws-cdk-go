package awstranscribe

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnMedicalTranscriptionJobPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnMedicalTranscriptionJobMixinProps := &CfnMedicalTranscriptionJobMixinProps{
//   	LanguageCode: jsii.String("languageCode"),
//   	Media: &MediaProperty{
//   		MediaFileUri: jsii.String("mediaFileUri"),
//   	},
//   	MediaFormat: jsii.String("mediaFormat"),
//   	MediaSampleRateHertz: jsii.Number(123),
//   	MedicalTranscriptionJobName: jsii.String("medicalTranscriptionJobName"),
//   	Settings: &MedicalTranscriptionSettingProperty{
//   		ChannelIdentification: jsii.Boolean(false),
//   		ShowAlternatives: jsii.Boolean(false),
//   	},
//   	Specialty: jsii.String("specialty"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicaltranscriptionjob.html
//
type CfnMedicalTranscriptionJobMixinProps struct {
	// The language code for the language spoken in the input media file.
	//
	// Must be en-US.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicaltranscriptionjob.html#cfn-transcribe-medicaltranscriptionjob-languagecode
	//
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicaltranscriptionjob.html#cfn-transcribe-medicaltranscriptionjob-media
	//
	Media interface{} `field:"optional" json:"media" yaml:"media"`
	// The format of the input media file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicaltranscriptionjob.html#cfn-transcribe-medicaltranscriptionjob-mediaformat
	//
	MediaFormat *string `field:"optional" json:"mediaFormat" yaml:"mediaFormat"`
	// The sample rate of the audio in hertz.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicaltranscriptionjob.html#cfn-transcribe-medicaltranscriptionjob-mediasampleratehertz
	//
	MediaSampleRateHertz *float64 `field:"optional" json:"mediaSampleRateHertz" yaml:"mediaSampleRateHertz"`
	// A unique name for the medical transcription job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicaltranscriptionjob.html#cfn-transcribe-medicaltranscriptionjob-medicaltranscriptionjobname
	//
	MedicalTranscriptionJobName *string `field:"optional" json:"medicalTranscriptionJobName" yaml:"medicalTranscriptionJobName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicaltranscriptionjob.html#cfn-transcribe-medicaltranscriptionjob-settings
	//
	Settings interface{} `field:"optional" json:"settings" yaml:"settings"`
	// The medical specialty represented in the media.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicaltranscriptionjob.html#cfn-transcribe-medicaltranscriptionjob-specialty
	//
	Specialty *string `field:"optional" json:"specialty" yaml:"specialty"`
	// Tags associated with the medical transcription job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicaltranscriptionjob.html#cfn-transcribe-medicaltranscriptionjob-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Whether the input media is a dictation or conversation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicaltranscriptionjob.html#cfn-transcribe-medicaltranscriptionjob-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

