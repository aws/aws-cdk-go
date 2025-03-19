package awscdkivsalpha


// Thumbnail recording mode.
// Experimental.
type RecordingMode string

const (
	// Use INTERVAL to enable the generation of thumbnails for recorded video at a time interval controlled by the TargetIntervalSeconds property.
	// Experimental.
	RecordingMode_INTERVAL RecordingMode = "INTERVAL"
	// Use DISABLED to disable the generation of thumbnails for recorded video.
	// Experimental.
	RecordingMode_DISABLED RecordingMode = "DISABLED"
)

