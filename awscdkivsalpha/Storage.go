package awscdkivsalpha


// The format in which thumbnails are recorded for a stream.
//
// Example:
//   var recordingBucket Bucket
//
//
//   recordingConfiguration := ivs.NewRecordingConfiguration(this, jsii.String("RecordingConfiguration"), &RecordingConfigurationProps{
//   	Bucket: recordingBucket,
//
//   	// set thumbnail settings
//   	ThumbnailConfiguration: ivs.ThumbnailConfiguration_Interval(ivs.Resolution_HD, []Storage{
//   		ivs.Storage_LATEST,
//   		ivs.Storage_SEQUENTIAL,
//   	}, awscdk.Duration_Seconds(jsii.Number(30))),
//   })
//
// Experimental.
type Storage string

const (
	// SEQUENTIAL records all generated thumbnails in a serial manner, to the media/thumbnails directory.
	// Experimental.
	Storage_SEQUENTIAL Storage = "SEQUENTIAL"
	// LATEST saves the latest thumbnail in media/thumbnails/latest/thumb.jpg and overwrites it at the interval specified by thumbnailTargetInterval.
	// Experimental.
	Storage_LATEST Storage = "LATEST"
)

