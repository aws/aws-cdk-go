package mixinsawsivs


// The `AWS::IVS::AutoParticipantRecordingConfiguration` property type describes a configuration for individual participant recording.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   autoParticipantRecordingConfigurationProperty := &AutoParticipantRecordingConfigurationProperty{
//   	HlsConfiguration: &HlsConfigurationProperty{
//   		ParticipantRecordingHlsConfiguration: &ParticipantRecordingHlsConfigurationProperty{
//   			TargetSegmentDurationSeconds: jsii.Number(123),
//   		},
//   	},
//   	MediaTypes: []*string{
//   		jsii.String("mediaTypes"),
//   	},
//   	RecordingReconnectWindowSeconds: jsii.Number(123),
//   	StorageConfigurationArn: jsii.String("storageConfigurationArn"),
//   	ThumbnailConfiguration: &ThumbnailConfigurationProperty{
//   		ParticipantThumbnailConfiguration: &ParticipantThumbnailConfigurationProperty{
//   			RecordingMode: jsii.String("recordingMode"),
//   			Storage: []*string{
//   				jsii.String("storage"),
//   			},
//   			TargetIntervalSeconds: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-stage-autoparticipantrecordingconfiguration.html
//
type CfnStagePropsMixin_AutoParticipantRecordingConfigurationProperty struct {
	// HLS configuration object for individual participant recording.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-stage-autoparticipantrecordingconfiguration.html#cfn-ivs-stage-autoparticipantrecordingconfiguration-hlsconfiguration
	//
	HlsConfiguration interface{} `field:"optional" json:"hlsConfiguration" yaml:"hlsConfiguration"`
	// Types of media to be recorded.
	//
	// Default: `AUDIO_VIDEO` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-stage-autoparticipantrecordingconfiguration.html#cfn-ivs-stage-autoparticipantrecordingconfiguration-mediatypes
	//
	MediaTypes *[]*string `field:"optional" json:"mediaTypes" yaml:"mediaTypes"`
	// If a stage publisher disconnects and then reconnects within the specified interval, the multiple recordings will be considered a single recording and merged together.
	//
	// The default value is 0, which disables merging.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-stage-autoparticipantrecordingconfiguration.html#cfn-ivs-stage-autoparticipantrecordingconfiguration-recordingreconnectwindowseconds
	//
	// Default: - 0.
	//
	RecordingReconnectWindowSeconds *float64 `field:"optional" json:"recordingReconnectWindowSeconds" yaml:"recordingReconnectWindowSeconds"`
	// ARN of the StorageConfiguration resource to use for individual participant recording.
	//
	// Default: "" (empty string, no storage configuration is specified). Individual participant recording cannot be started unless a storage configuration is specified, when a Stage is created or updated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-stage-autoparticipantrecordingconfiguration.html#cfn-ivs-stage-autoparticipantrecordingconfiguration-storageconfigurationarn
	//
	StorageConfigurationArn *string `field:"optional" json:"storageConfigurationArn" yaml:"storageConfigurationArn"`
	// A complex type that allows you to enable/disable the recording of thumbnails for individual participant recording and modify the interval at which thumbnails are generated for the live session.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-stage-autoparticipantrecordingconfiguration.html#cfn-ivs-stage-autoparticipantrecordingconfiguration-thumbnailconfiguration
	//
	ThumbnailConfiguration interface{} `field:"optional" json:"thumbnailConfiguration" yaml:"thumbnailConfiguration"`
}

