package awstranscribe


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mediaProperty := &MediaProperty{
//   	MediaFileUri: jsii.String("mediaFileUri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-medicaltranscriptionjob-media.html
//
type CfnMedicalTranscriptionJob_MediaProperty struct {
	// The Amazon S3 location of the media file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-medicaltranscriptionjob-media.html#cfn-transcribe-medicaltranscriptionjob-media-mediafileuri
	//
	MediaFileUri *string `field:"optional" json:"mediaFileUri" yaml:"mediaFileUri"`
}

