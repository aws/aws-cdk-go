package awsivs


// The ThumbnailConfiguration property type describes a configuration of thumbnails for recorded video.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   thumbnailConfigurationProperty := &thumbnailConfigurationProperty{
//   	recordingMode: jsii.String("recordingMode"),
//
//   	// the properties below are optional
//   	targetIntervalSeconds: jsii.Number(123),
//   }
//
type CfnRecordingConfiguration_ThumbnailConfigurationProperty struct {
	// Thumbnail recording mode. Valid values:.
	//
	// - `DISABLED` : Use DISABLED to disable the generation of thumbnails for recorded video.
	// - `INTERVAL` : Use INTERVAL to enable the generation of thumbnails for recorded video at a time interval controlled by the [TargetIntervalSeconds](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-recordingconfiguration-thumbnailconfiguration.html#cfn-ivs-recordingconfiguration-thumbnailconfiguration-targetintervalseconds) property.
	//
	// *Default* : `INTERVAL`.
	RecordingMode *string `field:"required" json:"recordingMode" yaml:"recordingMode"`
	// The targeted thumbnail-generation interval in seconds. This is configurable (and required) only if [RecordingMode](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-recordingconfiguration-thumbnailconfiguration.html#cfn-ivs-recordingconfiguration-thumbnailconfiguration-recordingmode) is `INTERVAL` .
	//
	// > Setting a value for `TargetIntervalSeconds` does not guarantee that thumbnails are generated at the specified interval. For thumbnails to be generated at the `TargetIntervalSeconds` interval, the `IDR/Keyframe` value for the input video must be less than the `TargetIntervalSeconds` value. See [Amazon IVS Streaming Configuration](https://docs.aws.amazon.com/ivs/latest/userguide/streaming-config.html) for information on setting `IDR/Keyframe` to the recommended value in video-encoder settings.
	//
	// *Default* : 60
	//
	// *Valid Range* : Minumum value of 5. Maximum value of 60.
	TargetIntervalSeconds *float64 `field:"optional" json:"targetIntervalSeconds" yaml:"targetIntervalSeconds"`
}

