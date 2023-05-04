package awsmedialive


// The settings for one UDP output.
//
// The parent of this entity is OutputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   udpOutputSettingsProperty := &UdpOutputSettingsProperty{
//   	BufferMsec: jsii.Number(123),
//   	ContainerSettings: &UdpContainerSettingsProperty{
//   		M2TsSettings: &M2tsSettingsProperty{
//   			AbsentInputAudioBehavior: jsii.String("absentInputAudioBehavior"),
//   			Arib: jsii.String("arib"),
//   			AribCaptionsPid: jsii.String("aribCaptionsPid"),
//   			AribCaptionsPidControl: jsii.String("aribCaptionsPidControl"),
//   			AudioBufferModel: jsii.String("audioBufferModel"),
//   			AudioFramesPerPes: jsii.Number(123),
//   			AudioPids: jsii.String("audioPids"),
//   			AudioStreamType: jsii.String("audioStreamType"),
//   			Bitrate: jsii.Number(123),
//   			BufferModel: jsii.String("bufferModel"),
//   			CcDescriptor: jsii.String("ccDescriptor"),
//   			DvbNitSettings: &DvbNitSettingsProperty{
//   				NetworkId: jsii.Number(123),
//   				NetworkName: jsii.String("networkName"),
//   				RepInterval: jsii.Number(123),
//   			},
//   			DvbSdtSettings: &DvbSdtSettingsProperty{
//   				OutputSdt: jsii.String("outputSdt"),
//   				RepInterval: jsii.Number(123),
//   				ServiceName: jsii.String("serviceName"),
//   				ServiceProviderName: jsii.String("serviceProviderName"),
//   			},
//   			DvbSubPids: jsii.String("dvbSubPids"),
//   			DvbTdtSettings: &DvbTdtSettingsProperty{
//   				RepInterval: jsii.Number(123),
//   			},
//   			DvbTeletextPid: jsii.String("dvbTeletextPid"),
//   			Ebif: jsii.String("ebif"),
//   			EbpAudioInterval: jsii.String("ebpAudioInterval"),
//   			EbpLookaheadMs: jsii.Number(123),
//   			EbpPlacement: jsii.String("ebpPlacement"),
//   			EcmPid: jsii.String("ecmPid"),
//   			EsRateInPes: jsii.String("esRateInPes"),
//   			EtvPlatformPid: jsii.String("etvPlatformPid"),
//   			EtvSignalPid: jsii.String("etvSignalPid"),
//   			FragmentTime: jsii.Number(123),
//   			Klv: jsii.String("klv"),
//   			KlvDataPids: jsii.String("klvDataPids"),
//   			NielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   			NullPacketBitrate: jsii.Number(123),
//   			PatInterval: jsii.Number(123),
//   			PcrControl: jsii.String("pcrControl"),
//   			PcrPeriod: jsii.Number(123),
//   			PcrPid: jsii.String("pcrPid"),
//   			PmtInterval: jsii.Number(123),
//   			PmtPid: jsii.String("pmtPid"),
//   			ProgramNum: jsii.Number(123),
//   			RateMode: jsii.String("rateMode"),
//   			Scte27Pids: jsii.String("scte27Pids"),
//   			Scte35Control: jsii.String("scte35Control"),
//   			Scte35Pid: jsii.String("scte35Pid"),
//   			Scte35PrerollPullupMilliseconds: jsii.Number(123),
//   			SegmentationMarkers: jsii.String("segmentationMarkers"),
//   			SegmentationStyle: jsii.String("segmentationStyle"),
//   			SegmentationTime: jsii.Number(123),
//   			TimedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   			TimedMetadataPid: jsii.String("timedMetadataPid"),
//   			TransportStreamId: jsii.Number(123),
//   			VideoPid: jsii.String("videoPid"),
//   		},
//   	},
//   	Destination: &OutputLocationRefProperty{
//   		DestinationRefId: jsii.String("destinationRefId"),
//   	},
//   	FecOutputSettings: &FecOutputSettingsProperty{
//   		ColumnDepth: jsii.Number(123),
//   		IncludeFec: jsii.String("includeFec"),
//   		RowLength: jsii.Number(123),
//   	},
//   }
//
type CfnChannel_UdpOutputSettingsProperty struct {
	// The UDP output buffering in milliseconds.
	//
	// Larger values increase latency through the transcoder but simultaneously assist the transcoder in maintaining a constant, low-jitter UDP/RTP output while accommodating clock recovery, input switching, input disruptions, picture reordering, and so on.
	BufferMsec *float64 `field:"optional" json:"bufferMsec" yaml:"bufferMsec"`
	// The settings for the UDP output.
	ContainerSettings interface{} `field:"optional" json:"containerSettings" yaml:"containerSettings"`
	// The destination address and port number for RTP or UDP packets.
	//
	// These can be unicast or multicast RTP or UDP (for example, rtp://239.10.10.10:5001 or udp://10.100.100.100:5002).
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// The settings for enabling and adjusting Forward Error Correction on UDP outputs.
	FecOutputSettings interface{} `field:"optional" json:"fecOutputSettings" yaml:"fecOutputSettings"`
}

