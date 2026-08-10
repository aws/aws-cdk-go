package awstranscribe


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   medicalTranscriptionSettingProperty := &MedicalTranscriptionSettingProperty{
//   	ChannelIdentification: jsii.Boolean(false),
//   	ShowAlternatives: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-medicaltranscriptionjob-medicaltranscriptionsetting.html
//
type CfnMedicalTranscriptionJob_MedicalTranscriptionSettingProperty struct {
	// Enables channel identification in multi-channel audio.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-medicaltranscriptionjob-medicaltranscriptionsetting.html#cfn-transcribe-medicaltranscriptionjob-medicaltranscriptionsetting-channelidentification
	//
	ChannelIdentification interface{} `field:"optional" json:"channelIdentification" yaml:"channelIdentification"`
	// Include alternative transcriptions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-medicaltranscriptionjob-medicaltranscriptionsetting.html#cfn-transcribe-medicaltranscriptionjob-medicaltranscriptionsetting-showalternatives
	//
	ShowAlternatives interface{} `field:"optional" json:"showAlternatives" yaml:"showAlternatives"`
}

