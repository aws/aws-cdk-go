package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for an MS Smooth output group.
//
// Example:
//   var video EncodeConfiguration
//   var audio EncodeConfiguration
//
//
//   medialive.OutputGroupConfiguration_MsSmooth(&MsSmoothOutputGroupProps{
//   	Name: jsii.String("smooth"),
//   	Destinations: []OutputDestination{
//   		medialive.OutputDestination_Url(jsii.String("https://smooth.example.com/live")),
//   	},
//   	Outputs: []MsSmoothOutputDefinition{
//   		&MsSmoothOutputDefinition{
//   			Encodes: []EncodeConfiguration{
//   				video,
//   				audio,
//   			},
//   			OutputName: jsii.String("smooth_out"),
//   		},
//   	},
//   })
//
// Experimental.
type MsSmoothOutputGroupProps struct {
	// The destinations for this output group — one per pipeline.
	//
	// Array position determines the pipeline mapping:
	// - `destinations[0]` → Pipeline 0
	// - `destinations[1]` → Pipeline 1 (STANDARD channels only)
	//
	// For a SINGLE_PIPELINE channel, provide exactly 1 destination.
	// For a STANDARD channel, provide exactly 2 destinations.
	// Experimental.
	Destinations *[]OutputDestination `field:"required" json:"destinations" yaml:"destinations"`
	// The name of this output group.
	//
	// Used as the destination reference ID. Underscores are normalised to hyphens internally.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value of the Acquisition Point Identity element used in each message placed in the sparse track.
	// Default: - service default.
	//
	// Experimental.
	AcquisitionPointId *string `field:"optional" json:"acquisitionPointId" yaml:"acquisitionPointId"`
	// If set to passthrough for an audio-only output, the fragment absolute time is set to the current timecode.
	// Default: - service default.
	//
	// Experimental.
	AudioOnlyTimecodeControl MsSmoothAudioOnlyTimecodeControl `field:"optional" json:"audioOnlyTimecodeControl" yaml:"audioOnlyTimecodeControl"`
	// If set to VERIFY_AUTHENTICITY, verifies the HTTPS certificate chain to a trusted CA.
	// Default: - service default.
	//
	// Experimental.
	CertificateMode MsSmoothCertificateMode `field:"optional" json:"certificateMode" yaml:"certificateMode"`
	// The number of seconds to wait before retrying the connection to the IIS server if the connection is lost.
	// Default: - service default.
	//
	// Experimental.
	ConnectionRetryInterval awscdk.Duration `field:"optional" json:"connectionRetryInterval" yaml:"connectionRetryInterval"`
	// The Microsoft Smooth channel ID that is sent to the IIS server.
	// Default: - service default.
	//
	// Experimental.
	EventId *string `field:"optional" json:"eventId" yaml:"eventId"`
	// Specifies whether to send a channel ID to the IIS server.
	// Default: - service default.
	//
	// Experimental.
	EventIdMode MsSmoothEventIdMode `field:"optional" json:"eventIdMode" yaml:"eventIdMode"`
	// When set to SEND_EOS, sends an EOS signal to an IIS server when stopping the channel.
	// Default: - service default.
	//
	// Experimental.
	EventStopBehavior MsSmoothEventStopBehavior `field:"optional" json:"eventStopBehavior" yaml:"eventStopBehavior"`
	// The size, in seconds, of the file cache for streaming outputs.
	// Default: - service default.
	//
	// Experimental.
	FilecacheDuration awscdk.Duration `field:"optional" json:"filecacheDuration" yaml:"filecacheDuration"`
	// The length, in seconds, of mp4 fragments to generate.
	// Default: - service default.
	//
	// Experimental.
	FragmentLength awscdk.Duration `field:"optional" json:"fragmentLength" yaml:"fragmentLength"`
	// A parameter that controls output group behavior on an input loss.
	// Default: - service default.
	//
	// Experimental.
	InputLossAction MsSmoothInputLossAction `field:"optional" json:"inputLossAction" yaml:"inputLossAction"`
	// The number of retry attempts.
	// Default: 10.
	//
	// Experimental.
	NumRetries *float64 `field:"optional" json:"numRetries" yaml:"numRetries"`
	// The outputs for this MS Smooth output group.
	// Default: - no initial outputs.
	//
	// Experimental.
	Outputs *[]*MsSmoothOutputDefinition `field:"optional" json:"outputs" yaml:"outputs"`
	// The number of seconds before initiating a restart due to output failure.
	// Default: - Duration.seconds(1)
	//
	// Experimental.
	RestartDelay awscdk.Duration `field:"optional" json:"restartDelay" yaml:"restartDelay"`
	// The segmentation mode.
	// Default: - service default.
	//
	// Experimental.
	SegmentationMode MsSmoothSegmentationMode `field:"optional" json:"segmentationMode" yaml:"segmentationMode"`
	// The number of milliseconds to delay the output from the second pipeline.
	// Default: - service default.
	//
	// Experimental.
	SendDelayMs *float64 `field:"optional" json:"sendDelayMs" yaml:"sendDelayMs"`
	// If set to SCTE_35, uses incoming SCTE-35 messages to generate a sparse track.
	// Default: - service default.
	//
	// Experimental.
	SparseTrackType MsSmoothSparseTrackType `field:"optional" json:"sparseTrackType" yaml:"sparseTrackType"`
	// When set to SEND, sends a stream manifest so that the publishing point doesn't start until all streams start.
	// Default: - service default.
	//
	// Experimental.
	StreamManifestBehavior MsSmoothStreamManifestBehavior `field:"optional" json:"streamManifestBehavior" yaml:"streamManifestBehavior"`
	// The timestamp offset for the channel.
	//
	// Used only if timestampOffsetMode is set to USE_CONFIGURED_OFFSET.
	// Default: - service default.
	//
	// Experimental.
	TimestampOffset *string `field:"optional" json:"timestampOffset" yaml:"timestampOffset"`
	// The type of timestamp date offset to use.
	// Default: - service default.
	//
	// Experimental.
	TimestampOffsetMode MsSmoothTimestampOffsetMode `field:"optional" json:"timestampOffsetMode" yaml:"timestampOffsetMode"`
}

