package awscdkivsalpha


// Resolution for rendition.
//
// Example:
//   var recordingBucket Bucket
//
//
//   recordingConfiguration := ivs.NewRecordingConfiguration(this, jsii.String("RecordingConfiguration"), &RecordingConfigurationProps{
//   	Bucket: recordingBucket,
//
//   	// set rendition configuration
//   	RenditionConfiguration: ivs.RenditionConfiguration_Custom([]Resolution{
//   		ivs.Resolution_HD,
//   		ivs.Resolution_SD,
//   	}),
//   })
//
// Experimental.
type Resolution string

const (
	// Full HD (1080p).
	// Experimental.
	Resolution_FULL_HD Resolution = "FULL_HD"
	// HD (720p).
	// Experimental.
	Resolution_HD Resolution = "HD"
	// SD (480p).
	// Experimental.
	Resolution_SD Resolution = "SD"
	// Lowest resolution.
	// Experimental.
	Resolution_LOWEST_RESOLUTION Resolution = "LOWEST_RESOLUTION"
)

