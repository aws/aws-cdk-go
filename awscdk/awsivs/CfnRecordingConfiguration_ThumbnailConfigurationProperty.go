package awsivs


// The ThumbnailConfiguration property type describes a configuration of thumbnails for recorded video.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   thumbnailConfigurationProperty := &ThumbnailConfigurationProperty{
//   	RecordingMode: jsii.String("recordingMode"),
//   	Resolution: jsii.String("resolution"),
//   	Storage: []*string{
//   		jsii.String("storage"),
//   	},
//   	TargetIntervalSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-recordingconfiguration-thumbnailconfiguration.html
//
type CfnRecordingConfiguration_ThumbnailConfigurationProperty struct {
	// Thumbnail recording mode. Valid values:.
	//
	// - `DISABLED` : Use DISABLED to disable the generation of thumbnails for recorded video.
	// - `INTERVAL` : Use INTERVAL to enable the generation of thumbnails for recorded video at a time interval controlled by the [TargetIntervalSeconds](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-recordingconfiguration-thumbnailconfiguration.html#cfn-ivs-recordingconfiguration-thumbnailconfiguration-targetintervalseconds) property.
	//
	// *Default* : `INTERVAL`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-recordingconfiguration-thumbnailconfiguration.html#cfn-ivs-recordingconfiguration-thumbnailconfiguration-recordingmode
	//
	// Default: - "INTERVAL".
	//
	RecordingMode *string `field:"optional" json:"recordingMode" yaml:"recordingMode"`
	// The desired resolution of recorded thumbnails for a stream.
	//
	// Thumbnails are recorded at the selected resolution if the corresponding rendition is available during the stream; otherwise, they are recorded at source resolution. For more information about resolution values and their corresponding height and width dimensions, see [Auto-Record to Amazon S3](https://docs.aws.amazon.com//ivs/latest/LowLatencyUserGuide/record-to-s3.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-recordingconfiguration-thumbnailconfiguration.html#cfn-ivs-recordingconfiguration-thumbnailconfiguration-resolution
	//
	Resolution *string `field:"optional" json:"resolution" yaml:"resolution"`
	// The format in which thumbnails are recorded for a stream.
	//
	// `SEQUENTIAL` records all generated thumbnails in a serial manner, to the media/thumbnails directory. `LATEST` saves the latest thumbnail in media/thumbnails/latest/thumb.jpg and overwrites it at the interval specified by `targetIntervalSeconds` . You can enable both `SEQUENTIAL` and `LATEST` . Default: `SEQUENTIAL` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-recordingconfiguration-thumbnailconfiguration.html#cfn-ivs-recordingconfiguration-thumbnailconfiguration-storage
	//
	Storage *[]*string `field:"optional" json:"storage" yaml:"storage"`
	// The targeted thumbnail-generation interval in seconds. This is configurable (and required) only if [RecordingMode](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-recordingconfiguration-thumbnailconfiguration.html#cfn-ivs-recordingconfiguration-thumbnailconfiguration-recordingmode) is `INTERVAL` .
	//
	// > Setting a value for `TargetIntervalSeconds` does not guarantee that thumbnails are generated at the specified interval. For thumbnails to be generated at the `TargetIntervalSeconds` interval, the `IDR/Keyframe` value for the input video must be less than the `TargetIntervalSeconds` value. See [Amazon IVS Streaming Configuration](https://docs.aws.amazon.com/ivs/latest/LowLatencyUserGuide/streaming-config.html) for information on setting `IDR/Keyframe` to the recommended value in video-encoder settings.
	//
	// *Default* : 60.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-recordingconfiguration-thumbnailconfiguration.html#cfn-ivs-recordingconfiguration-thumbnailconfiguration-targetintervalseconds
	//
	// Default: - 60.
	//
	TargetIntervalSeconds *float64 `field:"optional" json:"targetIntervalSeconds" yaml:"targetIntervalSeconds"`
}

