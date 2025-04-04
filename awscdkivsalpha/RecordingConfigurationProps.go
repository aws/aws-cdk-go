package awscdkivsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Properties of the IVS Recording configuration.
//
// Example:
//   var recordingBucket bucket
//
//
//   recordingConfiguration := ivs.NewRecordingConfiguration(this, jsii.String("RecordingConfiguration"), &RecordingConfigurationProps{
//   	Bucket: recordingBucket,
//
//   	// set rendition configuration
//   	RenditionConfiguration: ivs.RenditionConfiguration_Custom([]resolution{
//   		ivs.*resolution_HD,
//   		ivs.*resolution_SD,
//   	}),
//   })
//
// Experimental.
type RecordingConfigurationProps struct {
	// S3 bucket where recorded videos will be stored.
	// Experimental.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The name of the Recording configuration.
	//
	// The value does not need to be unique.
	// Default: - auto generate.
	//
	// Experimental.
	RecordingConfigurationName *string `field:"optional" json:"recordingConfigurationName" yaml:"recordingConfigurationName"`
	// If a broadcast disconnects and then reconnects within the specified interval, the multiple streams will be considered a single broadcast and merged together.
	//
	// `recordingReconnectWindow` must be between 0 and 300 seconds.
	// Default: - 0 seconds (means disabled).
	//
	// Experimental.
	RecordingReconnectWindow awscdk.Duration `field:"optional" json:"recordingReconnectWindow" yaml:"recordingReconnectWindow"`
	// A rendition configuration describes which renditions should be recorded for a stream.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-recordingconfiguration-renditionconfiguration.html
	//
	// Default: - no rendition configuration.
	//
	// Experimental.
	RenditionConfiguration RenditionConfiguration `field:"optional" json:"renditionConfiguration" yaml:"renditionConfiguration"`
	// A thumbnail configuration enables/disables the recording of thumbnails for a live session and controls the interval at which thumbnails are generated for the live session.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-recordingconfiguration-thumbnailconfiguration.html
	//
	// Default: - no thumbnail configuration.
	//
	// Experimental.
	ThumbnailConfiguration ThumbnailConfiguration `field:"optional" json:"thumbnailConfiguration" yaml:"thumbnailConfiguration"`
}

