package previewawsivsmixins


// An object representing a configuration of thumbnails for recorded video.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   thumbnailConfigurationProperty := &ThumbnailConfigurationProperty{
//   	ParticipantThumbnailConfiguration: &ParticipantThumbnailConfigurationProperty{
//   		RecordingMode: jsii.String("recordingMode"),
//   		Storage: []*string{
//   			jsii.String("storage"),
//   		},
//   		TargetIntervalSeconds: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-stage-thumbnailconfiguration.html
//
type CfnStagePropsMixin_ThumbnailConfigurationProperty struct {
	// Object specifying a configuration of thumbnails for recorded video from an individual participant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-stage-thumbnailconfiguration.html#cfn-ivs-stage-thumbnailconfiguration-participantthumbnailconfiguration
	//
	ParticipantThumbnailConfiguration interface{} `field:"optional" json:"participantThumbnailConfiguration" yaml:"participantThumbnailConfiguration"`
}

