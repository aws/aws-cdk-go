package awsmedialive


// The settings for a Microsoft Smooth output group.
//
// The parent of this entity is OutputGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   msSmoothGroupSettingsProperty := &msSmoothGroupSettingsProperty{
//   	acquisitionPointId: jsii.String("acquisitionPointId"),
//   	audioOnlyTimecodeControl: jsii.String("audioOnlyTimecodeControl"),
//   	certificateMode: jsii.String("certificateMode"),
//   	connectionRetryInterval: jsii.Number(123),
//   	destination: &outputLocationRefProperty{
//   		destinationRefId: jsii.String("destinationRefId"),
//   	},
//   	eventId: jsii.String("eventId"),
//   	eventIdMode: jsii.String("eventIdMode"),
//   	eventStopBehavior: jsii.String("eventStopBehavior"),
//   	filecacheDuration: jsii.Number(123),
//   	fragmentLength: jsii.Number(123),
//   	inputLossAction: jsii.String("inputLossAction"),
//   	numRetries: jsii.Number(123),
//   	restartDelay: jsii.Number(123),
//   	segmentationMode: jsii.String("segmentationMode"),
//   	sendDelayMs: jsii.Number(123),
//   	sparseTrackType: jsii.String("sparseTrackType"),
//   	streamManifestBehavior: jsii.String("streamManifestBehavior"),
//   	timestampOffset: jsii.String("timestampOffset"),
//   	timestampOffsetMode: jsii.String("timestampOffsetMode"),
//   }
//
type CfnChannel_MsSmoothGroupSettingsProperty struct {
	// The value of the Acquisition Point Identity element that is used in each message placed in the sparse track.
	//
	// Enabled only if sparseTrackType is not "none."
	AcquisitionPointId *string `field:"optional" json:"acquisitionPointId" yaml:"acquisitionPointId"`
	// If set to passthrough for an audio-only Microsoft Smooth output, the fragment absolute time is set to the current timecode.
	//
	// This option does not write timecodes to the audio elementary stream.
	AudioOnlyTimecodeControl *string `field:"optional" json:"audioOnlyTimecodeControl" yaml:"audioOnlyTimecodeControl"`
	// If set to verifyAuthenticity, verifies the HTTPS certificate chain to a trusted certificate authority (CA).
	//
	// This causes HTTPS outputs to self-signed certificates to fail.
	CertificateMode *string `field:"optional" json:"certificateMode" yaml:"certificateMode"`
	// The number of seconds to wait before retrying the connection to the IIS server if the connection is lost.
	//
	// Content is cached during this time, and the cache is delivered to the IIS server after the connection is re-established.
	ConnectionRetryInterval *float64 `field:"optional" json:"connectionRetryInterval" yaml:"connectionRetryInterval"`
	// The Smooth Streaming publish point on an IIS server.
	//
	// MediaLive acts as a "Push" encoder to IIS.
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// The Microsoft Smooth channel ID that is sent to the IIS server.
	//
	// Specify the ID only if eventIdMode is set to useConfigured.
	EventId *string `field:"optional" json:"eventId" yaml:"eventId"`
	// Specifies whether to send a channel ID to the IIS server.
	//
	// If no channel ID is sent and the same channel is used without changing the publishing point, clients might see cached video from the previous run. Options: - "useConfigured" - use the value provided in eventId - "useTimestamp" - generate and send a channel ID based on the current timestamp - "noEventId" - do not send a channel ID to the IIS server.
	EventIdMode *string `field:"optional" json:"eventIdMode" yaml:"eventIdMode"`
	// When set to sendEos, sends an EOS signal to an IIS server when stopping the channel.
	EventStopBehavior *string `field:"optional" json:"eventStopBehavior" yaml:"eventStopBehavior"`
	// The size, in seconds, of the file cache for streaming outputs.
	FilecacheDuration *float64 `field:"optional" json:"filecacheDuration" yaml:"filecacheDuration"`
	// The length, in seconds, of mp4 fragments to generate.
	//
	// The fragment length must be compatible with GOP size and frame rate.
	FragmentLength *float64 `field:"optional" json:"fragmentLength" yaml:"fragmentLength"`
	// A parameter that controls output group behavior on an input loss.
	InputLossAction *string `field:"optional" json:"inputLossAction" yaml:"inputLossAction"`
	// The number of retry attempts.
	NumRetries *float64 `field:"optional" json:"numRetries" yaml:"numRetries"`
	// The number of seconds before initiating a restart due to output failure, due to exhausting the numRetries on one segment, or exceeding filecacheDuration.
	RestartDelay *float64 `field:"optional" json:"restartDelay" yaml:"restartDelay"`
	// useInputSegmentation has been deprecated.
	//
	// The configured segment size is always used.
	SegmentationMode *string `field:"optional" json:"segmentationMode" yaml:"segmentationMode"`
	// The number of milliseconds to delay the output from the second pipeline.
	SendDelayMs *float64 `field:"optional" json:"sendDelayMs" yaml:"sendDelayMs"`
	// If set to scte35, uses incoming SCTE-35 messages to generate a sparse track in this group of Microsoft Smooth outputs.
	SparseTrackType *string `field:"optional" json:"sparseTrackType" yaml:"sparseTrackType"`
	// When set to send, sends a stream manifest so that the publishing point doesn't start until all streams start.
	StreamManifestBehavior *string `field:"optional" json:"streamManifestBehavior" yaml:"streamManifestBehavior"`
	// The timestamp offset for the channel.
	//
	// Used only if timestampOffsetMode is set to useConfiguredOffset.
	TimestampOffset *string `field:"optional" json:"timestampOffset" yaml:"timestampOffset"`
	// The type of timestamp date offset to use.
	//
	// - useEventStartDate: Use the date the channel was started as the offset - useConfiguredOffset: Use an explicitly configured date as the offset.
	TimestampOffsetMode *string `field:"optional" json:"timestampOffsetMode" yaml:"timestampOffsetMode"`
}

