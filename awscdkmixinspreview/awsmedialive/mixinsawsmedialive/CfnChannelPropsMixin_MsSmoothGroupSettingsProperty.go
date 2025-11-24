package mixinsawsmedialive


// The settings for a Microsoft Smooth output group.
//
// The parent of this entity is OutputGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   msSmoothGroupSettingsProperty := &MsSmoothGroupSettingsProperty{
//   	AcquisitionPointId: jsii.String("acquisitionPointId"),
//   	AudioOnlyTimecodeControl: jsii.String("audioOnlyTimecodeControl"),
//   	CertificateMode: jsii.String("certificateMode"),
//   	ConnectionRetryInterval: jsii.Number(123),
//   	Destination: &OutputLocationRefProperty{
//   		DestinationRefId: jsii.String("destinationRefId"),
//   	},
//   	EventId: jsii.String("eventId"),
//   	EventIdMode: jsii.String("eventIdMode"),
//   	EventStopBehavior: jsii.String("eventStopBehavior"),
//   	FilecacheDuration: jsii.Number(123),
//   	FragmentLength: jsii.Number(123),
//   	InputLossAction: jsii.String("inputLossAction"),
//   	NumRetries: jsii.Number(123),
//   	RestartDelay: jsii.Number(123),
//   	SegmentationMode: jsii.String("segmentationMode"),
//   	SendDelayMs: jsii.Number(123),
//   	SparseTrackType: jsii.String("sparseTrackType"),
//   	StreamManifestBehavior: jsii.String("streamManifestBehavior"),
//   	TimestampOffset: jsii.String("timestampOffset"),
//   	TimestampOffsetMode: jsii.String("timestampOffsetMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html
//
type CfnChannelPropsMixin_MsSmoothGroupSettingsProperty struct {
	// The value of the Acquisition Point Identity element that is used in each message placed in the sparse track.
	//
	// Enabled only if sparseTrackType is not "none."
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-acquisitionpointid
	//
	AcquisitionPointId *string `field:"optional" json:"acquisitionPointId" yaml:"acquisitionPointId"`
	// If set to passthrough for an audio-only Microsoft Smooth output, the fragment absolute time is set to the current timecode.
	//
	// This option does not write timecodes to the audio elementary stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-audioonlytimecodecontrol
	//
	AudioOnlyTimecodeControl *string `field:"optional" json:"audioOnlyTimecodeControl" yaml:"audioOnlyTimecodeControl"`
	// If set to verifyAuthenticity, verifies the HTTPS certificate chain to a trusted certificate authority (CA).
	//
	// This causes HTTPS outputs to self-signed certificates to fail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-certificatemode
	//
	CertificateMode *string `field:"optional" json:"certificateMode" yaml:"certificateMode"`
	// The number of seconds to wait before retrying the connection to the IIS server if the connection is lost.
	//
	// Content is cached during this time, and the cache is delivered to the IIS server after the connection is re-established.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-connectionretryinterval
	//
	ConnectionRetryInterval *float64 `field:"optional" json:"connectionRetryInterval" yaml:"connectionRetryInterval"`
	// The Smooth Streaming publish point on an IIS server.
	//
	// MediaLive acts as a "Push" encoder to IIS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-destination
	//
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// The Microsoft Smooth channel ID that is sent to the IIS server.
	//
	// Specify the ID only if eventIdMode is set to useConfigured.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-eventid
	//
	EventId *string `field:"optional" json:"eventId" yaml:"eventId"`
	// Specifies whether to send a channel ID to the IIS server.
	//
	// If no channel ID is sent and the same channel is used without changing the publishing point, clients might see cached video from the previous run. Options: - "useConfigured" - use the value provided in eventId - "useTimestamp" - generate and send a channel ID based on the current timestamp - "noEventId" - do not send a channel ID to the IIS server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-eventidmode
	//
	EventIdMode *string `field:"optional" json:"eventIdMode" yaml:"eventIdMode"`
	// When set to sendEos, sends an EOS signal to an IIS server when stopping the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-eventstopbehavior
	//
	EventStopBehavior *string `field:"optional" json:"eventStopBehavior" yaml:"eventStopBehavior"`
	// The size, in seconds, of the file cache for streaming outputs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-filecacheduration
	//
	FilecacheDuration *float64 `field:"optional" json:"filecacheDuration" yaml:"filecacheDuration"`
	// The length, in seconds, of mp4 fragments to generate.
	//
	// The fragment length must be compatible with GOP size and frame rate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-fragmentlength
	//
	FragmentLength *float64 `field:"optional" json:"fragmentLength" yaml:"fragmentLength"`
	// A parameter that controls output group behavior on an input loss.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-inputlossaction
	//
	InputLossAction *string `field:"optional" json:"inputLossAction" yaml:"inputLossAction"`
	// The number of retry attempts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-numretries
	//
	NumRetries *float64 `field:"optional" json:"numRetries" yaml:"numRetries"`
	// The number of seconds before initiating a restart due to output failure, due to exhausting the numRetries on one segment, or exceeding filecacheDuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-restartdelay
	//
	RestartDelay *float64 `field:"optional" json:"restartDelay" yaml:"restartDelay"`
	// useInputSegmentation has been deprecated.
	//
	// The configured segment size is always used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-segmentationmode
	//
	SegmentationMode *string `field:"optional" json:"segmentationMode" yaml:"segmentationMode"`
	// The number of milliseconds to delay the output from the second pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-senddelayms
	//
	SendDelayMs *float64 `field:"optional" json:"sendDelayMs" yaml:"sendDelayMs"`
	// If set to scte35, uses incoming SCTE-35 messages to generate a sparse track in this group of Microsoft Smooth outputs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-sparsetracktype
	//
	SparseTrackType *string `field:"optional" json:"sparseTrackType" yaml:"sparseTrackType"`
	// When set to send, sends a stream manifest so that the publishing point doesn't start until all streams start.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-streammanifestbehavior
	//
	StreamManifestBehavior *string `field:"optional" json:"streamManifestBehavior" yaml:"streamManifestBehavior"`
	// The timestamp offset for the channel.
	//
	// Used only if timestampOffsetMode is set to useConfiguredOffset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-timestampoffset
	//
	TimestampOffset *string `field:"optional" json:"timestampOffset" yaml:"timestampOffset"`
	// The type of timestamp date offset to use.
	//
	// - useEventStartDate: Use the date the channel was started as the offset - useConfiguredOffset: Use an explicitly configured date as the offset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-timestampoffsetmode
	//
	TimestampOffsetMode *string `field:"optional" json:"timestampOffsetMode" yaml:"timestampOffsetMode"`
}

