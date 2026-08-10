package awstranscribe


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   medicalTranscriptProperty := &MedicalTranscriptProperty{
//   	TranscriptFileUri: jsii.String("transcriptFileUri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-medicaltranscriptionjob-medicaltranscript.html
//
type CfnMedicalTranscriptionJobPropsMixin_MedicalTranscriptProperty struct {
	// The Amazon S3 location of the transcript.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-medicaltranscriptionjob-medicaltranscript.html#cfn-transcribe-medicaltranscriptionjob-medicaltranscript-transcriptfileuri
	//
	TranscriptFileUri *string `field:"optional" json:"transcriptFileUri" yaml:"transcriptFileUri"`
}

