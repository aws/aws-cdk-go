package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   abWatermarkingCustomProfileProperty := &AbWatermarkingCustomProfileProperty{
//   	EmbeddingFrequency: jsii.Number(123),
//   	SceneCut: jsii.Number(123),
//   	TargetPsnr: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-abwatermarkingcustomprofile.html
//
type CfnChannel_AbWatermarkingCustomProfileProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-abwatermarkingcustomprofile.html#cfn-medialive-channel-abwatermarkingcustomprofile-embeddingfrequency
	//
	EmbeddingFrequency *float64 `field:"optional" json:"embeddingFrequency" yaml:"embeddingFrequency"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-abwatermarkingcustomprofile.html#cfn-medialive-channel-abwatermarkingcustomprofile-scenecut
	//
	SceneCut *float64 `field:"optional" json:"sceneCut" yaml:"sceneCut"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-abwatermarkingcustomprofile.html#cfn-medialive-channel-abwatermarkingcustomprofile-targetpsnr
	//
	TargetPsnr *float64 `field:"optional" json:"targetPsnr" yaml:"targetPsnr"`
}

