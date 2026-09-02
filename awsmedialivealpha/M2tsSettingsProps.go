package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for MPEG-2 transport stream (M2TS) container settings.
//
// Used by the UDP, Archive, SRT, and MediaConnect Router output groups. All properties are
// optional; omit them to use MediaLive's service defaults.
//
// PID properties accept a decimal or hexadecimal value (and, where noted, ranges or comma-separated
// lists). Each PID must be in the range 32 (0x20)..8182 (0x1ff6).
//
// Example:
//   var video EncodeConfiguration
//   var audio EncodeConfiguration
//
//
//   medialive.OutputGroupConfiguration_Udp(&UdpOutputGroupProps{
//   	Name: jsii.String("udp_out"),
//   	Destinations: []UdpOutputDestination{
//   		medialive.UdpOutputDestination_Udp(&TransportOutputDestinationProps{
//   			Address: jsii.String("203.0.113.5"),
//   			Port: jsii.Number(5000),
//   		}),
//   	},
//   	Outputs: []UdpOutputDefinition{
//   		&UdpOutputDefinition{
//   			Encodes: []EncodeConfiguration{
//   				video,
//   				audio,
//   			},
//   			OutputName: jsii.String("ts"),
//   			M2tsSettings: medialive.M2tsSettings_Of(&M2tsSettingsProps{
//   				Bitrate: awscdk.Bitrate_Mbps(jsii.Number(8)),
//   				RateMode: medialive.M2tsRateMode_VBR(),
//   				ProgramNum: jsii.Number(1),
//   				PatInterval: awscdk.Duration_Millis(jsii.Number(100)),
//   				PmtInterval: awscdk.Duration_*Millis(jsii.Number(100)),
//   				Scte35Control: medialive.M2tsScte35Control_PASSTHROUGH(),
//   				DvbSdtSettings: &DvbSdtSettings{
//   					OutputSdt: medialive.DvbSdtOutputMode_SDT_MANUAL(),
//   					ServiceName: jsii.String("My Service"),
//   					RepInterval: awscdk.Duration_*Millis(jsii.Number(2000)),
//   				},
//   			}),
//   		},
//   	},
//   })
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html
//
// Experimental.
type M2tsSettingsProps struct {
	// Behavior when the selected input audio stream is removed.
	// Default: - service default.
	//
	// Experimental.
	AbsentInputAudioBehavior M2tsAbsentInputAudioBehavior `field:"optional" json:"absentInputAudioBehavior" yaml:"absentInputAudioBehavior"`
	// ARIB-compliant field muxing.
	// Default: - service default.
	//
	// Experimental.
	Arib M2tsArib `field:"optional" json:"arib" yaml:"arib"`
	// The PID for ARIB Captions.
	// Default: - service default.
	//
	// Experimental.
	AribCaptionsPid *string `field:"optional" json:"aribCaptionsPid" yaml:"aribCaptionsPid"`
	// How the ARIB Captions PID is selected.
	// Default: - service default.
	//
	// Experimental.
	AribCaptionsPidControl M2tsAribCaptionsPidControl `field:"optional" json:"aribCaptionsPidControl" yaml:"aribCaptionsPidControl"`
	// The buffer model for Dolby Digital audio.
	// Default: - service default.
	//
	// Experimental.
	AudioBufferModel M2tsAudioBufferModel `field:"optional" json:"audioBufferModel" yaml:"audioBufferModel"`
	// The number of audio frames to insert per PES packet.
	// Default: - service default.
	//
	// Experimental.
	AudioFramesPerPes *float64 `field:"optional" json:"audioFramesPerPes" yaml:"audioFramesPerPes"`
	// The PID(s) of the elementary audio streams (ranges/comma-separated allowed).
	// Default: - service default.
	//
	// Experimental.
	AudioPids *string `field:"optional" json:"audioPids" yaml:"audioPids"`
	// The stream type used for audio elementary streams.
	// Default: - service default.
	//
	// Experimental.
	AudioStreamType M2tsAudioStreamType `field:"optional" json:"audioStreamType" yaml:"audioStreamType"`
	// The output bitrate of the transport stream.
	//
	// Set to 0 bps to let the muxer choose.
	// Default: - muxer chooses.
	//
	// Experimental.
	Bitrate awscdk.Bitrate `field:"optional" json:"bitrate" yaml:"bitrate"`
	// The transport stream buffer model.
	// Default: - service default.
	//
	// Experimental.
	BufferModel M2tsBufferModel `field:"optional" json:"bufferModel" yaml:"bufferModel"`
	// Whether to generate the captionServiceDescriptor in the PMT.
	// Default: - service default.
	//
	// Experimental.
	CcDescriptor M2tsCcDescriptor `field:"optional" json:"ccDescriptor" yaml:"ccDescriptor"`
	// DVB Network Information Table (NIT) settings.
	// Default: - no NIT.
	//
	// Experimental.
	DvbNitSettings *DvbNitSettings `field:"optional" json:"dvbNitSettings" yaml:"dvbNitSettings"`
	// DVB Service Description Table (SDT) settings.
	// Default: - no SDT.
	//
	// Experimental.
	DvbSdtSettings *DvbSdtSettings `field:"optional" json:"dvbSdtSettings" yaml:"dvbSdtSettings"`
	// The PID(s) for input source DVB Subtitle data (ranges/comma-separated allowed).
	// Default: - service default.
	//
	// Experimental.
	DvbSubPids *string `field:"optional" json:"dvbSubPids" yaml:"dvbSubPids"`
	// DVB Time and Date Table (TDT) settings.
	// Default: - no TDT.
	//
	// Experimental.
	DvbTdtSettings *DvbTdtSettings `field:"optional" json:"dvbTdtSettings" yaml:"dvbTdtSettings"`
	// The PID for input source DVB Teletext data.
	// Default: - service default.
	//
	// Experimental.
	DvbTeletextPid *string `field:"optional" json:"dvbTeletextPid" yaml:"dvbTeletextPid"`
	// EBIF data passthrough behavior.
	// Default: - service default.
	//
	// Experimental.
	Ebif M2tsEbif `field:"optional" json:"ebif" yaml:"ebif"`
	// Placement of audio EBP markers.
	// Default: - service default.
	//
	// Experimental.
	EbpAudioInterval M2tsEbpAudioInterval `field:"optional" json:"ebpAudioInterval" yaml:"ebpAudioInterval"`
	// The EBP lookahead interval.
	// Default: - service default.
	//
	// Experimental.
	EbpLookahead awscdk.Duration `field:"optional" json:"ebpLookahead" yaml:"ebpLookahead"`
	// Placement of EBP markers on audio PIDs.
	// Default: - service default.
	//
	// Experimental.
	EbpPlacement M2tsEbpPlacement `field:"optional" json:"ebpPlacement" yaml:"ebpPlacement"`
	// Whether to include the ES Rate field in the PES header.
	// Default: - service default.
	//
	// Experimental.
	EsRateInPes M2tsEsRateInPes `field:"optional" json:"esRateInPes" yaml:"esRateInPes"`
	// The PID for input source ETV Platform data.
	// Default: - service default.
	//
	// Experimental.
	EtvPlatformPid *string `field:"optional" json:"etvPlatformPid" yaml:"etvPlatformPid"`
	// The PID for input source ETV Signal data.
	// Default: - service default.
	//
	// Experimental.
	EtvSignalPid *string `field:"optional" json:"etvSignalPid" yaml:"etvSignalPid"`
	// The length of each fragment (used only with EBP markers).
	// Default: - service default.
	//
	// Experimental.
	FragmentTime awscdk.Duration `field:"optional" json:"fragmentTime" yaml:"fragmentTime"`
	// KLV data passthrough behavior.
	// Default: - service default.
	//
	// Experimental.
	Klv M2tsKlv `field:"optional" json:"klv" yaml:"klv"`
	// The PID(s) for input source KLV data (ranges/comma-separated allowed).
	// Default: - service default.
	//
	// Experimental.
	KlvDataPids *string `field:"optional" json:"klvDataPids" yaml:"klvDataPids"`
	// Nielsen ID3 passthrough behavior.
	// Default: - service default.
	//
	// Experimental.
	NielsenId3Behavior M2tsNielsenId3Behavior `field:"optional" json:"nielsenId3Behavior" yaml:"nielsenId3Behavior"`
	// The bitrate of extra null packets to insert into the transport stream.
	// Default: - no null packets.
	//
	// Experimental.
	NullPacketBitrate awscdk.Bitrate `field:"optional" json:"nullPacketBitrate" yaml:"nullPacketBitrate"`
	// The interval between PAT instances (0, or 10ms..1000ms).
	// Default: - service default.
	//
	// Experimental.
	PatInterval awscdk.Duration `field:"optional" json:"patInterval" yaml:"patInterval"`
	// Controls insertion of the Program Clock Reference (PCR).
	// Default: - service default.
	//
	// Experimental.
	PcrControl M2tsPcrControl `field:"optional" json:"pcrControl" yaml:"pcrControl"`
	// The maximum interval between Program Clock References (PCRs).
	// Default: - service default.
	//
	// Experimental.
	PcrPeriod awscdk.Duration `field:"optional" json:"pcrPeriod" yaml:"pcrPeriod"`
	// The PID of the Program Clock Reference.
	// Default: - same as the video PID.
	//
	// Experimental.
	PcrPid *string `field:"optional" json:"pcrPid" yaml:"pcrPid"`
	// The interval between PMT instances (0, or 10ms..1000ms).
	// Default: - service default.
	//
	// Experimental.
	PmtInterval awscdk.Duration `field:"optional" json:"pmtInterval" yaml:"pmtInterval"`
	// The PID for the Program Map Table (PMT).
	// Default: - service default.
	//
	// Experimental.
	PmtPid *string `field:"optional" json:"pmtPid" yaml:"pmtPid"`
	// The value of the program number field in the PMT.
	// Default: - service default.
	//
	// Experimental.
	ProgramNum *float64 `field:"optional" json:"programNum" yaml:"programNum"`
	// The transport stream bitrate mode (CBR/VBR).
	// Default: - service default.
	//
	// Experimental.
	RateMode M2tsRateMode `field:"optional" json:"rateMode" yaml:"rateMode"`
	// The PID(s) for input source SCTE-27 data (ranges/comma-separated allowed).
	// Default: - service default.
	//
	// Experimental.
	Scte27Pids *string `field:"optional" json:"scte27Pids" yaml:"scte27Pids"`
	// SCTE-35 passthrough behavior.
	// Default: - service default.
	//
	// Experimental.
	Scte35Control M2tsScte35Control `field:"optional" json:"scte35Control" yaml:"scte35Control"`
	// The PID of the SCTE-35 stream.
	// Default: - service default.
	//
	// Experimental.
	Scte35Pid *string `field:"optional" json:"scte35Pid" yaml:"scte35Pid"`
	// The SCTE-35 preroll pullup interval.
	// Default: - service default.
	//
	// Experimental.
	Scte35PrerollPullup awscdk.Duration `field:"optional" json:"scte35PrerollPullup" yaml:"scte35PrerollPullup"`
	// The type of segmentation markers to insert.
	// Default: - service default.
	//
	// Experimental.
	SegmentationMarkers M2tsSegmentationMarkers `field:"optional" json:"segmentationMarkers" yaml:"segmentationMarkers"`
	// How segmentation markers respond to avails.
	// Default: - service default.
	//
	// Experimental.
	SegmentationStyle M2tsSegmentationStyle `field:"optional" json:"segmentationStyle" yaml:"segmentationStyle"`
	// The length of each segment (required unless `segmentationMarkers` is NONE).
	// Default: - service default.
	//
	// Experimental.
	SegmentationTime awscdk.Duration `field:"optional" json:"segmentationTime" yaml:"segmentationTime"`
	// Timed metadata passthrough behavior.
	// Default: - service default.
	//
	// Experimental.
	TimedMetadataBehavior M2tsTimedMetadataBehavior `field:"optional" json:"timedMetadataBehavior" yaml:"timedMetadataBehavior"`
	// The PID of the timed metadata stream.
	// Default: - service default.
	//
	// Experimental.
	TimedMetadataPid *string `field:"optional" json:"timedMetadataPid" yaml:"timedMetadataPid"`
	// The value of the transport stream ID field in the PMT.
	// Default: - service default.
	//
	// Experimental.
	TransportStreamId *float64 `field:"optional" json:"transportStreamId" yaml:"transportStreamId"`
	// The PID of the elementary video stream.
	// Default: - service default.
	//
	// Experimental.
	VideoPid *string `field:"optional" json:"videoPid" yaml:"videoPid"`
}

