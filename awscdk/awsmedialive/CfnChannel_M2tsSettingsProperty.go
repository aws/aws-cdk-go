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
//   m2tsSettingsProperty := &M2tsSettingsProperty{
//   	AbsentInputAudioBehavior: jsii.String("absentInputAudioBehavior"),
//   	Arib: jsii.String("arib"),
//   	AribCaptionsPid: jsii.String("aribCaptionsPid"),
//   	AribCaptionsPidControl: jsii.String("aribCaptionsPidControl"),
//   	AudioBufferModel: jsii.String("audioBufferModel"),
//   	AudioFramesPerPes: jsii.Number(123),
//   	AudioPids: jsii.String("audioPids"),
//   	AudioStreamType: jsii.String("audioStreamType"),
//   	Bitrate: jsii.Number(123),
//   	BufferModel: jsii.String("bufferModel"),
//   	CcDescriptor: jsii.String("ccDescriptor"),
//   	DvbNitSettings: &DvbNitSettingsProperty{
//   		NetworkId: jsii.Number(123),
//   		NetworkName: jsii.String("networkName"),
//   		RepInterval: jsii.Number(123),
//   	},
//   	DvbSdtSettings: &DvbSdtSettingsProperty{
//   		OutputSdt: jsii.String("outputSdt"),
//   		RepInterval: jsii.Number(123),
//   		ServiceName: jsii.String("serviceName"),
//   		ServiceProviderName: jsii.String("serviceProviderName"),
//   	},
//   	DvbSubPids: jsii.String("dvbSubPids"),
//   	DvbTdtSettings: &DvbTdtSettingsProperty{
//   		RepInterval: jsii.Number(123),
//   	},
//   	DvbTeletextPid: jsii.String("dvbTeletextPid"),
//   	Ebif: jsii.String("ebif"),
//   	EbpAudioInterval: jsii.String("ebpAudioInterval"),
//   	EbpLookaheadMs: jsii.Number(123),
//   	EbpPlacement: jsii.String("ebpPlacement"),
//   	EcmPid: jsii.String("ecmPid"),
//   	EsRateInPes: jsii.String("esRateInPes"),
//   	EtvPlatformPid: jsii.String("etvPlatformPid"),
//   	EtvSignalPid: jsii.String("etvSignalPid"),
//   	FragmentTime: jsii.Number(123),
//   	Klv: jsii.String("klv"),
//   	KlvDataPids: jsii.String("klvDataPids"),
//   	NielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   	NullPacketBitrate: jsii.Number(123),
//   	PatInterval: jsii.Number(123),
//   	PcrControl: jsii.String("pcrControl"),
//   	PcrPeriod: jsii.Number(123),
//   	PcrPid: jsii.String("pcrPid"),
//   	PmtInterval: jsii.Number(123),
//   	PmtPid: jsii.String("pmtPid"),
//   	ProgramNum: jsii.Number(123),
//   	RateMode: jsii.String("rateMode"),
//   	Scte27Pids: jsii.String("scte27Pids"),
//   	Scte35Control: jsii.String("scte35Control"),
//   	Scte35Pid: jsii.String("scte35Pid"),
//   	Scte35PrerollPullupMilliseconds: jsii.Number(123),
//   	SegmentationMarkers: jsii.String("segmentationMarkers"),
//   	SegmentationStyle: jsii.String("segmentationStyle"),
//   	SegmentationTime: jsii.Number(123),
//   	TimedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   	TimedMetadataPid: jsii.String("timedMetadataPid"),
//   	TransportStreamId: jsii.Number(123),
//   	VideoPid: jsii.String("videoPid"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html
//
type CfnChannel_M2tsSettingsProperty struct {
	// When set to drop, the output audio streams are removed from the program if the selected input audio stream is removed from the input.
	//
	// This allows the output audio configuration to dynamically change based on the input configuration. If this is set to encodeSilence, all output audio streams will output encoded silence when not connected to an active input stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-absentinputaudiobehavior
	//
	AbsentInputAudioBehavior *string `field:"optional" json:"absentInputAudioBehavior" yaml:"absentInputAudioBehavior"`
	// When set to enabled, uses ARIB-compliant field muxing and removes video descriptor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-arib
	//
	Arib *string `field:"optional" json:"arib" yaml:"arib"`
	// The PID for ARIB Captions in the transport stream.
	//
	// You can enter the value as a decimal or hexadecimal value. Valid values are 32 (or 0x20)..8182 (or 0x1ff6).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-aribcaptionspid
	//
	AribCaptionsPid *string `field:"optional" json:"aribCaptionsPid" yaml:"aribCaptionsPid"`
	// If set to auto, The PID number used for ARIB Captions will be auto-selected from unused PIDs.
	//
	// If set to useConfigured, ARIB captions will be on the configured PID number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-aribcaptionspidcontrol
	//
	AribCaptionsPidControl *string `field:"optional" json:"aribCaptionsPidControl" yaml:"aribCaptionsPidControl"`
	// When set to dvb, uses the DVB buffer model for Dolby Digital audio.
	//
	// When set to atsc, the ATSC model is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-audiobuffermodel
	//
	AudioBufferModel *string `field:"optional" json:"audioBufferModel" yaml:"audioBufferModel"`
	// The number of audio frames to insert for each PES packet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-audioframesperpes
	//
	AudioFramesPerPes *float64 `field:"optional" json:"audioFramesPerPes" yaml:"audioFramesPerPes"`
	// The PID of the elementary audio streams in the transport stream.
	//
	// Multiple values are accepted, and can be entered in ranges or by comma separation. You can enter the value as a decimal or hexadecimal value. Each PID specified must be in the range of 32 (or 0x20)..8182 (or 0x1ff6).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-audiopids
	//
	AudioPids *string `field:"optional" json:"audioPids" yaml:"audioPids"`
	// When set to atsc, uses stream type = 0x81 for AC3 and stream type = 0x87 for EAC3.
	//
	// When set to dvb, uses stream type = 0x06.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-audiostreamtype
	//
	AudioStreamType *string `field:"optional" json:"audioStreamType" yaml:"audioStreamType"`
	// The output bitrate of the transport stream in bits per second.
	//
	// Setting to 0 lets the muxer automatically determine the appropriate bitrate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-bitrate
	//
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// If set to multiplex, uses the multiplex buffer model for accurate interleaving.
	//
	// Setting to bufferModel to none can lead to lower latency, but low-memory devices might not be able to play back the stream without interruptions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-buffermodel
	//
	BufferModel *string `field:"optional" json:"bufferModel" yaml:"bufferModel"`
	// When set to enabled, generates captionServiceDescriptor in PMT.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-ccdescriptor
	//
	CcDescriptor *string `field:"optional" json:"ccDescriptor" yaml:"ccDescriptor"`
	// Inserts a DVB Network Information Table (NIT) at the specified table repetition interval.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-dvbnitsettings
	//
	DvbNitSettings interface{} `field:"optional" json:"dvbNitSettings" yaml:"dvbNitSettings"`
	// Inserts a DVB Service Description Table (SDT) at the specified table repetition interval.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-dvbsdtsettings
	//
	DvbSdtSettings interface{} `field:"optional" json:"dvbSdtSettings" yaml:"dvbSdtSettings"`
	// The PID for the input source DVB Subtitle data to this output.
	//
	// Multiple values are accepted, and can be entered in ranges and/or by comma separation. You can enter the value as a decimal or hexadecimal value. Each PID specified must be in the range of 32 (or 0x20)..8182 (or 0x1ff6).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-dvbsubpids
	//
	DvbSubPids *string `field:"optional" json:"dvbSubPids" yaml:"dvbSubPids"`
	// Inserts DVB Time and Date Table (TDT) at the specified table repetition interval.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-dvbtdtsettings
	//
	DvbTdtSettings interface{} `field:"optional" json:"dvbTdtSettings" yaml:"dvbTdtSettings"`
	// The PID for the input source DVB Teletext data to this output.
	//
	// You can enter the value as a decimal or hexadecimal value. Valid values are 32 (or 0x20)..8182 (or 0x1ff6).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-dvbteletextpid
	//
	DvbTeletextPid *string `field:"optional" json:"dvbTeletextPid" yaml:"dvbTeletextPid"`
	// If set to passthrough, passes any EBIF data from the input source to this output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-ebif
	//
	Ebif *string `field:"optional" json:"ebif" yaml:"ebif"`
	// When videoAndFixedIntervals is selected, audio EBP markers are added to partitions 3 and 4.
	//
	// The interval between these additional markers is fixed, and is slightly shorter than the video EBP marker interval. This is only available when EBP Cablelabs segmentation markers are selected. Partitions 1 and 2 always follow the video interval.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-ebpaudiointerval
	//
	EbpAudioInterval *string `field:"optional" json:"ebpAudioInterval" yaml:"ebpAudioInterval"`
	// When set, enforces that Encoder Boundary Points do not come within the specified time interval of each other by looking ahead at input video.
	//
	// If another EBP is going to come in within the specified time interval, the current EBP is not emitted, and the segment is "stretched" to the next marker. The lookahead value does not add latency to the system. The channel must be configured elsewhere to create sufficient latency to make the lookahead accurate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-ebplookaheadms
	//
	EbpLookaheadMs *float64 `field:"optional" json:"ebpLookaheadMs" yaml:"ebpLookaheadMs"`
	// Controls placement of EBP on audio PIDs.
	//
	// If set to videoAndAudioPids, EBP markers are placed on the video PID and all audio PIDs. If set to videoPid, EBP markers are placed on only the video PID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-ebpplacement
	//
	EbpPlacement *string `field:"optional" json:"ebpPlacement" yaml:"ebpPlacement"`
	// This field is unused and deprecated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-ecmpid
	//
	EcmPid *string `field:"optional" json:"ecmPid" yaml:"ecmPid"`
	// Includes or excludes the ES Rate field in the PES header.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-esrateinpes
	//
	EsRateInPes *string `field:"optional" json:"esRateInPes" yaml:"esRateInPes"`
	// The PID for the input source ETV Platform data to this output.
	//
	// You can enter it as a decimal or hexadecimal value. Valid values are 32 (or 0x20) to 8182 (or 0x1ff6).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-etvplatformpid
	//
	EtvPlatformPid *string `field:"optional" json:"etvPlatformPid" yaml:"etvPlatformPid"`
	// The PID for input source ETV Signal data to this output.
	//
	// You can enter the value as a decimal or hexadecimal value. Valid values are 32 (or 0x20)..8182 (or 0x1ff6).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-etvsignalpid
	//
	EtvSignalPid *string `field:"optional" json:"etvSignalPid" yaml:"etvSignalPid"`
	// The length in seconds of each fragment.
	//
	// This is used only with EBP markers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-fragmenttime
	//
	FragmentTime *float64 `field:"optional" json:"fragmentTime" yaml:"fragmentTime"`
	// If set to passthrough, passes any KLV data from the input source to this output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-klv
	//
	Klv *string `field:"optional" json:"klv" yaml:"klv"`
	// The PID for the input source KLV data to this output.
	//
	// Multiple values are accepted, and can be entered in ranges or by comma separation. You can enter the value as a decimal or hexadecimal value. Each PID specified must be in the range of 32 (or 0x20)..8182 (or 0x1ff6).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-klvdatapids
	//
	KlvDataPids *string `field:"optional" json:"klvDataPids" yaml:"klvDataPids"`
	// If set to passthrough, Nielsen inaudible tones for media tracking will be detected in the input audio and an equivalent ID3 tag will be inserted in the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-nielsenid3behavior
	//
	NielsenId3Behavior *string `field:"optional" json:"nielsenId3Behavior" yaml:"nielsenId3Behavior"`
	// The value, in bits per second, of extra null packets to insert into the transport stream.
	//
	// This can be used if a downstream encryption system requires periodic null packets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-nullpacketbitrate
	//
	NullPacketBitrate *float64 `field:"optional" json:"nullPacketBitrate" yaml:"nullPacketBitrate"`
	// The number of milliseconds between instances of this table in the output transport stream.
	//
	// Valid values are 0, 10..1000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-patinterval
	//
	PatInterval *float64 `field:"optional" json:"patInterval" yaml:"patInterval"`
	// When set to pcrEveryPesPacket, a Program Clock Reference value is inserted for every Packetized Elementary Stream (PES) header.
	//
	// This parameter is effective only when the PCR PID is the same as the video or audio elementary stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-pcrcontrol
	//
	PcrControl *string `field:"optional" json:"pcrControl" yaml:"pcrControl"`
	// The maximum time, in milliseconds, between Program Clock References (PCRs) inserted into the transport stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-pcrperiod
	//
	PcrPeriod *float64 `field:"optional" json:"pcrPeriod" yaml:"pcrPeriod"`
	// The PID of the Program Clock Reference (PCR) in the transport stream.
	//
	// When no value is given, MediaLive assigns the same value as the video PID. You can enter the value as a decimal or hexadecimal value. Valid values are 32 (or 0x20)..8182 (or 0x1ff6).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-pcrpid
	//
	PcrPid *string `field:"optional" json:"pcrPid" yaml:"pcrPid"`
	// The number of milliseconds between instances of this table in the output transport stream.
	//
	// Valid values are 0, 10..1000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-pmtinterval
	//
	PmtInterval *float64 `field:"optional" json:"pmtInterval" yaml:"pmtInterval"`
	// The PID for the Program Map Table (PMT) in the transport stream.
	//
	// You can enter the value as a decimal or hexadecimal value. Valid values are 32 (or 0x20)..8182 (or 0x1ff6).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-pmtpid
	//
	PmtPid *string `field:"optional" json:"pmtPid" yaml:"pmtPid"`
	// The value of the program number field in the Program Map Table (PMT).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-programnum
	//
	ProgramNum *float64 `field:"optional" json:"programNum" yaml:"programNum"`
	// When VBR, does not insert null packets into the transport stream to fill the specified bitrate.
	//
	// The bitrate setting acts as the maximum bitrate when VBR is set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-ratemode
	//
	RateMode *string `field:"optional" json:"rateMode" yaml:"rateMode"`
	// The PID for the input source SCTE-27 data to this output.
	//
	// Multiple values are accepted, and can be entered in ranges or by comma separation. You can enter the value as a decimal or hexadecimal value. Each PID specified must be in the range of 32 (or 0x20)..8182 (or 0x1ff6).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-scte27pids
	//
	Scte27Pids *string `field:"optional" json:"scte27Pids" yaml:"scte27Pids"`
	// Optionally passes SCTE-35 signals from the input source to this output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-scte35control
	//
	Scte35Control *string `field:"optional" json:"scte35Control" yaml:"scte35Control"`
	// The PID of the SCTE-35 stream in the transport stream.
	//
	// You can enter the value as a decimal or hexadecimal value. Valid values are 32 (or 0x20)..8182 (or 0x1ff6).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-scte35pid
	//
	Scte35Pid *string `field:"optional" json:"scte35Pid" yaml:"scte35Pid"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-scte35prerollpullupmilliseconds
	//
	Scte35PrerollPullupMilliseconds *float64 `field:"optional" json:"scte35PrerollPullupMilliseconds" yaml:"scte35PrerollPullupMilliseconds"`
	// Inserts segmentation markers at each segmentationTime period.
	//
	// raiSegstart sets the Random Access Indicator bit in the adaptation field. raiAdapt sets the RAI bit and adds the current timecode in the private data bytes. psiSegstart inserts PAT and PMT tables at the start of segments. ebp adds Encoder Boundary Point information to the adaptation field as per OpenCable specification OC-SP-EBP-I01-130118. ebpLegacy adds Encoder Boundary Point information to the adaptation field using a legacy proprietary format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-segmentationmarkers
	//
	SegmentationMarkers *string `field:"optional" json:"segmentationMarkers" yaml:"segmentationMarkers"`
	// The segmentation style parameter controls how segmentation markers are inserted into the transport stream.
	//
	// With avails, it is possible that segments might be truncated, which can influence where future segmentation markers are inserted. When a segmentation style of resetCadence is selected and a segment is truncated due to an avail, we will reset the segmentation cadence. This means the subsequent segment will have a duration of $segmentationTime seconds. When a segmentation style of maintainCadence is selected and a segment is truncated due to an avail, we will not reset the segmentation cadence. This means the subsequent segment will likely be truncated as well. However, all segments after that will have a duration of $segmentationTime seconds. Note that EBP lookahead is a slight exception to this rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-segmentationstyle
	//
	SegmentationStyle *string `field:"optional" json:"segmentationStyle" yaml:"segmentationStyle"`
	// The length, in seconds, of each segment.
	//
	// This is required unless markers is set to None_.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-segmentationtime
	//
	SegmentationTime *float64 `field:"optional" json:"segmentationTime" yaml:"segmentationTime"`
	// When set to passthrough, timed metadata is passed through from input to output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-timedmetadatabehavior
	//
	TimedMetadataBehavior *string `field:"optional" json:"timedMetadataBehavior" yaml:"timedMetadataBehavior"`
	// The PID of the timed metadata stream in the transport stream.
	//
	// You can enter the value as a decimal or hexadecimal value. Valid values are 32 (or 0x20)..8182 (or 0x1ff6).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-timedmetadatapid
	//
	TimedMetadataPid *string `field:"optional" json:"timedMetadataPid" yaml:"timedMetadataPid"`
	// The value of the transport stream ID field in the Program Map Table (PMT).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-transportstreamid
	//
	TransportStreamId *float64 `field:"optional" json:"transportStreamId" yaml:"transportStreamId"`
	// The PID of the elementary video stream in the transport stream.
	//
	// You can enter the value as a decimal or hexadecimal value. Valid values are 32 (or 0x20)..8182 (or 0x1ff6).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html#cfn-medialive-channel-m2tssettings-videopid
	//
	VideoPid *string `field:"optional" json:"videoPid" yaml:"videoPid"`
}

