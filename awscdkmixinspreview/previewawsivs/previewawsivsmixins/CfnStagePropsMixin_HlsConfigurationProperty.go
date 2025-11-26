package previewawsivsmixins


// Object specifying an HLS configuration for individual participant recording.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   hlsConfigurationProperty := &HlsConfigurationProperty{
//   	ParticipantRecordingHlsConfiguration: &ParticipantRecordingHlsConfigurationProperty{
//   		TargetSegmentDurationSeconds: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-stage-hlsconfiguration.html
//
type CfnStagePropsMixin_HlsConfigurationProperty struct {
	// Object specifying a configuration of participant HLS recordings for individual participant recording.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-stage-hlsconfiguration.html#cfn-ivs-stage-hlsconfiguration-participantrecordinghlsconfiguration
	//
	ParticipantRecordingHlsConfiguration interface{} `field:"optional" json:"participantRecordingHlsConfiguration" yaml:"participantRecordingHlsConfiguration"`
}

