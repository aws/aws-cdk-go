package previewawsivsmixins


// Object specifying a configuration of participant HLS recordings for individual participant recording.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   participantRecordingHlsConfigurationProperty := &ParticipantRecordingHlsConfigurationProperty{
//   	TargetSegmentDurationSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-stage-participantrecordinghlsconfiguration.html
//
type CfnStagePropsMixin_ParticipantRecordingHlsConfigurationProperty struct {
	// Defines the target duration for recorded segments generated when recording a stage participant.
	//
	// Segments may have durations longer than the specified value when needed to ensure each segment begins with a keyframe. Default: 6.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-stage-participantrecordinghlsconfiguration.html#cfn-ivs-stage-participantrecordinghlsconfiguration-targetsegmentdurationseconds
	//
	// Default: - 6.
	//
	TargetSegmentDurationSeconds *float64 `field:"optional" json:"targetSegmentDurationSeconds" yaml:"targetSegmentDurationSeconds"`
}

