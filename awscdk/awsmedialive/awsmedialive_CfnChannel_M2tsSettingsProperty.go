package awsmedialive


// The configuration of the M2TS in the output.
//
// The parents of this entity are ArchiveContainerSettings and UdpContainerSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   m2tsSettingsProperty := &m2tsSettingsProperty{
//   	absentInputAudioBehavior: jsii.String("absentInputAudioBehavior"),
//   	arib: jsii.String("arib"),
//   	aribCaptionsPid: jsii.String("aribCaptionsPid"),
//   	aribCaptionsPidControl: jsii.String("aribCaptionsPidControl"),
//   	audioBufferModel: jsii.String("audioBufferModel"),
//   	audioFramesPerPes: jsii.Number(123),
//   	audioPids: jsii.String("audioPids"),
//   	audioStreamType: jsii.String("audioStreamType"),
//   	bitrate: jsii.Number(123),
//   	bufferModel: jsii.String("bufferModel"),
//   	ccDescriptor: jsii.String("ccDescriptor"),
//   	dvbNitSettings: &dvbNitSettingsProperty{
//   		networkId: jsii.Number(123),
//   		networkName: jsii.String("networkName"),
//   		repInterval: jsii.Number(123),
//   	},
//   	dvbSdtSettings: &dvbSdtSettingsProperty{
//   		outputSdt: jsii.String("outputSdt"),
//   		repInterval: jsii.Number(123),
//   		serviceName: jsii.String("serviceName"),
//   		serviceProviderName: jsii.String("serviceProviderName"),
//   	},
//   	dvbSubPids: jsii.String("dvbSubPids"),
//   	dvbTdtSettings: &dvbTdtSettingsProperty{
//   		repInterval: jsii.Number(123),
//   	},
//   	dvbTeletextPid: jsii.String("dvbTeletextPid"),
//   	ebif: jsii.String("ebif"),
//   	ebpAudioInterval: jsii.String("ebpAudioInterval"),
//   	ebpLookaheadMs: jsii.Number(123),
//   	ebpPlacement: jsii.String("ebpPlacement"),
//   	ecmPid: jsii.String("ecmPid"),
//   	esRateInPes: jsii.String("esRateInPes"),
//   	etvPlatformPid: jsii.String("etvPlatformPid"),
//   	etvSignalPid: jsii.String("etvSignalPid"),
//   	fragmentTime: jsii.Number(123),
//   	klv: jsii.String("klv"),
//   	klvDataPids: jsii.String("klvDataPids"),
//   	nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   	nullPacketBitrate: jsii.Number(123),
//   	patInterval: jsii.Number(123),
//   	pcrControl: jsii.String("pcrControl"),
//   	pcrPeriod: jsii.Number(123),
//   	pcrPid: jsii.String("pcrPid"),
//   	pmtInterval: jsii.Number(123),
//   	pmtPid: jsii.String("pmtPid"),
//   	programNum: jsii.Number(123),
//   	rateMode: jsii.String("rateMode"),
//   	scte27Pids: jsii.String("scte27Pids"),
//   	scte35Control: jsii.String("scte35Control"),
//   	scte35Pid: jsii.String("scte35Pid"),
//   	segmentationMarkers: jsii.String("segmentationMarkers"),
//   	segmentationStyle: jsii.String("segmentationStyle"),
//   	segmentationTime: jsii.Number(123),
//   	timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   	timedMetadataPid: jsii.String("timedMetadataPid"),
//   	transportStreamId: jsii.Number(123),
//   	videoPid: jsii.String("videoPid"),
//   }
//
type CfnChannel_M2tsSettingsProperty struct {
	// When set to drop, the output audio streams are removed from the program if the selected input audio stream is removed from the input.
	//
	// This allows the output audio configuration to dynamically change based on the input configuration. If this is set to encodeSilence, all output audio streams will output encoded silence when not connected to an active input stream.
	AbsentInputAudioBehavior *string `field:"optional" json:"absentInputAudioBehavior" yaml:"absentInputAudioBehavior"`
	// When set to enabled, uses ARIB-compliant field muxing and removes video descriptor.
	Arib *string `field:"optional" json:"arib" yaml:"arib"`
	// The PID for ARIB Captions in the transport stream.
	//
	// You can enter the value as a decimal or hexadecimal value. Valid values are 32 (or 0x20)..8182 (or 0x1ff6).
	AribCaptionsPid *string `field:"optional" json:"aribCaptionsPid" yaml:"aribCaptionsPid"`
	// If set to auto, The PID number used for ARIB Captions will be auto-selected from unused PIDs.
	//
	// If set to useConfigured, ARIB captions will be on the configured PID number.
	AribCaptionsPidControl *string `field:"optional" json:"aribCaptionsPidControl" yaml:"aribCaptionsPidControl"`
	// When set to dvb, uses the DVB buffer model for Dolby Digital audio.
	//
	// When set to atsc, the ATSC model is used.
	AudioBufferModel *string `field:"optional" json:"audioBufferModel" yaml:"audioBufferModel"`
	// The number of audio frames to insert for each PES packet.
	AudioFramesPerPes *float64 `field:"optional" json:"audioFramesPerPes" yaml:"audioFramesPerPes"`
	// The PID of the elementary audio streams in the transport stream.
	//
	// Multiple values are accepted, and can be entered in ranges or by comma separation. You can enter the value as a decimal or hexadecimal value. Each PID specified must be in the range of 32 (or 0x20)..8182 (or 0x1ff6).
	AudioPids *string `field:"optional" json:"audioPids" yaml:"audioPids"`
	// When set to atsc, uses stream type = 0x81 for AC3 and stream type = 0x87 for EAC3.
	//
	// When set to dvb, uses stream type = 0x06.
	AudioStreamType *string `field:"optional" json:"audioStreamType" yaml:"audioStreamType"`
	// The output bitrate of the transport stream in bits per second.
	//
	// Setting to 0 lets the muxer automatically determine the appropriate bitrate.
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// If set to multiplex, uses the multiplex buffer model for accurate interleaving.
	//
	// Setting to bufferModel to none can lead to lower latency, but low-memory devices might not be able to play back the stream without interruptions.
	BufferModel *string `field:"optional" json:"bufferModel" yaml:"bufferModel"`
	// When set to enabled, generates captionServiceDescriptor in PMT.
	CcDescriptor *string `field:"optional" json:"ccDescriptor" yaml:"ccDescriptor"`
	// Inserts a DVB Network Information Table (NIT) at the specified table repetition interval.
	DvbNitSettings interface{} `field:"optional" json:"dvbNitSettings" yaml:"dvbNitSettings"`
	// Inserts a DVB Service Description Table (SDT) at the specified table repetition interval.
	DvbSdtSettings interface{} `field:"optional" json:"dvbSdtSettings" yaml:"dvbSdtSettings"`
	// The PID for the input source DVB Subtitle data to this output.
	//
	// Multiple values are accepted, and can be entered in ranges and/or by comma separation. You can enter the value as a decimal or hexadecimal value. Each PID specified must be in the range of 32 (or 0x20)..8182 (or 0x1ff6).
	DvbSubPids *string `field:"optional" json:"dvbSubPids" yaml:"dvbSubPids"`
	// Inserts DVB Time and Date Table (TDT) at the specified table repetition interval.
	DvbTdtSettings interface{} `field:"optional" json:"dvbTdtSettings" yaml:"dvbTdtSettings"`
	// The PID for the input source DVB Teletext data to this output.
	//
	// You can enter the value as a decimal or hexadecimal value. Valid values are 32 (or 0x20)..8182 (or 0x1ff6).
	DvbTeletextPid *string `field:"optional" json:"dvbTeletextPid" yaml:"dvbTeletextPid"`
	// If set to passthrough, passes any EBIF data from the input source to this output.
	Ebif *string `field:"optional" json:"ebif" yaml:"ebif"`
	// When videoAndFixedIntervals is selected, audio EBP markers are added to partitions 3 and 4.
	//
	// The interval between these additional markers is fixed, and is slightly shorter than the video EBP marker interval. This is only available when EBP Cablelabs segmentation markers are selected. Partitions 1 and 2 always follow the video interval.
	EbpAudioInterval *string `field:"optional" json:"ebpAudioInterval" yaml:"ebpAudioInterval"`
	// When set, enforces that Encoder Boundary Points do not come within the specified time interval of each other by looking ahead at input video.
	//
	// If another EBP is going to come in within the specified time interval, the current EBP is not emitted, and the segment is "stretched" to the next marker. The lookahead value does not add latency to the system. The channel must be configured elsewhere to create sufficient latency to make the lookahead accurate.
	EbpLookaheadMs *float64 `field:"optional" json:"ebpLookaheadMs" yaml:"ebpLookaheadMs"`
	// Controls placement of EBP on audio PIDs.
	//
	// If set to videoAndAudioPids, EBP markers are placed on the video PID and all audio PIDs. If set to videoPid, EBP markers are placed on only the video PID.
	EbpPlacement *string `field:"optional" json:"ebpPlacement" yaml:"ebpPlacement"`
	// This field is unused and deprecated.
	EcmPid *string `field:"optional" json:"ecmPid" yaml:"ecmPid"`
	// Includes or excludes the ES Rate field in the PES header.
	EsRateInPes *string `field:"optional" json:"esRateInPes" yaml:"esRateInPes"`
	// The PID for the input source ETV Platform data to this output.
	//
	// You can enter it as a decimal or hexadecimal value. Valid values are 32 (or 0x20) to 8182 (or 0x1ff6).
	EtvPlatformPid *string `field:"optional" json:"etvPlatformPid" yaml:"etvPlatformPid"`
	// The PID for input source ETV Signal data to this output.
	//
	// You can enter the value as a decimal or hexadecimal value. Valid values are 32 (or 0x20)..8182 (or 0x1ff6).
	EtvSignalPid *string `field:"optional" json:"etvSignalPid" yaml:"etvSignalPid"`
	// The length in seconds of each fragment.
	//
	// This is used only with EBP markers.
	FragmentTime *float64 `field:"optional" json:"fragmentTime" yaml:"fragmentTime"`
	// If set to passthrough, passes any KLV data from the input source to this output.
	Klv *string `field:"optional" json:"klv" yaml:"klv"`
	// The PID for the input source KLV data to this output.
	//
	// Multiple values are accepted, and can be entered in ranges or by comma separation. You can enter the value as a decimal or hexadecimal value. Each PID specified must be in the range of 32 (or 0x20)..8182 (or 0x1ff6).
	KlvDataPids *string `field:"optional" json:"klvDataPids" yaml:"klvDataPids"`
	// If set to passthrough, Nielsen inaudible tones for media tracking will be detected in the input audio and an equivalent ID3 tag will be inserted in the output.
	NielsenId3Behavior *string `field:"optional" json:"nielsenId3Behavior" yaml:"nielsenId3Behavior"`
	// The value, in bits per second, of extra null packets to insert into the transport stream.
	//
	// This can be used if a downstream encryption system requires periodic null packets.
	NullPacketBitrate *float64 `field:"optional" json:"nullPacketBitrate" yaml:"nullPacketBitrate"`
	// The number of milliseconds between instances of this table in the output transport stream.
	//
	// Valid values are 0, 10..1000.
	PatInterval *float64 `field:"optional" json:"patInterval" yaml:"patInterval"`
	// When set to pcrEveryPesPacket, a Program Clock Reference value is inserted for every Packetized Elementary Stream (PES) header.
	//
	// This parameter is effective only when the PCR PID is the same as the video or audio elementary stream.
	PcrControl *string `field:"optional" json:"pcrControl" yaml:"pcrControl"`
	// The maximum time, in milliseconds, between Program Clock References (PCRs) inserted into the transport stream.
	PcrPeriod *float64 `field:"optional" json:"pcrPeriod" yaml:"pcrPeriod"`
	// The PID of the Program Clock Reference (PCR) in the transport stream.
	//
	// When no value is given, MediaLive assigns the same value as the video PID. You can enter the value as a decimal or hexadecimal value. Valid values are 32 (or 0x20)..8182 (or 0x1ff6).
	PcrPid *string `field:"optional" json:"pcrPid" yaml:"pcrPid"`
	// The number of milliseconds between instances of this table in the output transport stream.
	//
	// Valid values are 0, 10..1000.
	PmtInterval *float64 `field:"optional" json:"pmtInterval" yaml:"pmtInterval"`
	// The PID for the Program Map Table (PMT) in the transport stream.
	//
	// You can enter the value as a decimal or hexadecimal value. Valid values are 32 (or 0x20)..8182 (or 0x1ff6).
	PmtPid *string `field:"optional" json:"pmtPid" yaml:"pmtPid"`
	// The value of the program number field in the Program Map Table (PMT).
	ProgramNum *float64 `field:"optional" json:"programNum" yaml:"programNum"`
	// When VBR, does not insert null packets into the transport stream to fill the specified bitrate.
	//
	// The bitrate setting acts as the maximum bitrate when VBR is set.
	RateMode *string `field:"optional" json:"rateMode" yaml:"rateMode"`
	// The PID for the input source SCTE-27 data to this output.
	//
	// Multiple values are accepted, and can be entered in ranges or by comma separation. You can enter the value as a decimal or hexadecimal value. Each PID specified must be in the range of 32 (or 0x20)..8182 (or 0x1ff6).
	Scte27Pids *string `field:"optional" json:"scte27Pids" yaml:"scte27Pids"`
	// Optionally passes SCTE-35 signals from the input source to this output.
	Scte35Control *string `field:"optional" json:"scte35Control" yaml:"scte35Control"`
	// The PID of the SCTE-35 stream in the transport stream.
	//
	// You can enter the value as a decimal or hexadecimal value. Valid values are 32 (or 0x20)..8182 (or 0x1ff6).
	Scte35Pid *string `field:"optional" json:"scte35Pid" yaml:"scte35Pid"`
	// Inserts segmentation markers at each segmentationTime period.
	//
	// raiSegstart sets the Random Access Indicator bit in the adaptation field. raiAdapt sets the RAI bit and adds the current timecode in the private data bytes. psiSegstart inserts PAT and PMT tables at the start of segments. ebp adds Encoder Boundary Point information to the adaptation field as per OpenCable specification OC-SP-EBP-I01-130118. ebpLegacy adds Encoder Boundary Point information to the adaptation field using a legacy proprietary format.
	SegmentationMarkers *string `field:"optional" json:"segmentationMarkers" yaml:"segmentationMarkers"`
	// The segmentation style parameter controls how segmentation markers are inserted into the transport stream.
	//
	// With avails, it is possible that segments might be truncated, which can influence where future segmentation markers are inserted. When a segmentation style of resetCadence is selected and a segment is truncated due to an avail, we will reset the segmentation cadence. This means the subsequent segment will have a duration of $segmentationTime seconds. When a segmentation style of maintainCadence is selected and a segment is truncated due to an avail, we will not reset the segmentation cadence. This means the subsequent segment will likely be truncated as well. However, all segments after that will have a duration of $segmentationTime seconds. Note that EBP lookahead is a slight exception to this rule.
	SegmentationStyle *string `field:"optional" json:"segmentationStyle" yaml:"segmentationStyle"`
	// The length, in seconds, of each segment.
	//
	// This is required unless markers is set to None_.
	SegmentationTime *float64 `field:"optional" json:"segmentationTime" yaml:"segmentationTime"`
	// When set to passthrough, timed metadata is passed through from input to output.
	TimedMetadataBehavior *string `field:"optional" json:"timedMetadataBehavior" yaml:"timedMetadataBehavior"`
	// The PID of the timed metadata stream in the transport stream.
	//
	// You can enter the value as a decimal or hexadecimal value. Valid values are 32 (or 0x20)..8182 (or 0x1ff6).
	TimedMetadataPid *string `field:"optional" json:"timedMetadataPid" yaml:"timedMetadataPid"`
	// The value of the transport stream ID field in the Program Map Table (PMT).
	TransportStreamId *float64 `field:"optional" json:"transportStreamId" yaml:"transportStreamId"`
	// The PID of the elementary video stream in the transport stream.
	//
	// You can enter the value as a decimal or hexadecimal value. Valid values are 32 (or 0x20)..8182 (or 0x1ff6).
	VideoPid *string `field:"optional" json:"videoPid" yaml:"videoPid"`
}

