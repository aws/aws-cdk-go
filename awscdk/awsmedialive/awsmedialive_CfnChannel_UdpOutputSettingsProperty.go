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
//   udpOutputSettingsProperty := &udpOutputSettingsProperty{
//   	bufferMsec: jsii.Number(123),
//   	containerSettings: &udpContainerSettingsProperty{
//   		m2TsSettings: &m2tsSettingsProperty{
//   			absentInputAudioBehavior: jsii.String("absentInputAudioBehavior"),
//   			arib: jsii.String("arib"),
//   			aribCaptionsPid: jsii.String("aribCaptionsPid"),
//   			aribCaptionsPidControl: jsii.String("aribCaptionsPidControl"),
//   			audioBufferModel: jsii.String("audioBufferModel"),
//   			audioFramesPerPes: jsii.Number(123),
//   			audioPids: jsii.String("audioPids"),
//   			audioStreamType: jsii.String("audioStreamType"),
//   			bitrate: jsii.Number(123),
//   			bufferModel: jsii.String("bufferModel"),
//   			ccDescriptor: jsii.String("ccDescriptor"),
//   			dvbNitSettings: &dvbNitSettingsProperty{
//   				networkId: jsii.Number(123),
//   				networkName: jsii.String("networkName"),
//   				repInterval: jsii.Number(123),
//   			},
//   			dvbSdtSettings: &dvbSdtSettingsProperty{
//   				outputSdt: jsii.String("outputSdt"),
//   				repInterval: jsii.Number(123),
//   				serviceName: jsii.String("serviceName"),
//   				serviceProviderName: jsii.String("serviceProviderName"),
//   			},
//   			dvbSubPids: jsii.String("dvbSubPids"),
//   			dvbTdtSettings: &dvbTdtSettingsProperty{
//   				repInterval: jsii.Number(123),
//   			},
//   			dvbTeletextPid: jsii.String("dvbTeletextPid"),
//   			ebif: jsii.String("ebif"),
//   			ebpAudioInterval: jsii.String("ebpAudioInterval"),
//   			ebpLookaheadMs: jsii.Number(123),
//   			ebpPlacement: jsii.String("ebpPlacement"),
//   			ecmPid: jsii.String("ecmPid"),
//   			esRateInPes: jsii.String("esRateInPes"),
//   			etvPlatformPid: jsii.String("etvPlatformPid"),
//   			etvSignalPid: jsii.String("etvSignalPid"),
//   			fragmentTime: jsii.Number(123),
//   			klv: jsii.String("klv"),
//   			klvDataPids: jsii.String("klvDataPids"),
//   			nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   			nullPacketBitrate: jsii.Number(123),
//   			patInterval: jsii.Number(123),
//   			pcrControl: jsii.String("pcrControl"),
//   			pcrPeriod: jsii.Number(123),
//   			pcrPid: jsii.String("pcrPid"),
//   			pmtInterval: jsii.Number(123),
//   			pmtPid: jsii.String("pmtPid"),
//   			programNum: jsii.Number(123),
//   			rateMode: jsii.String("rateMode"),
//   			scte27Pids: jsii.String("scte27Pids"),
//   			scte35Control: jsii.String("scte35Control"),
//   			scte35Pid: jsii.String("scte35Pid"),
//   			segmentationMarkers: jsii.String("segmentationMarkers"),
//   			segmentationStyle: jsii.String("segmentationStyle"),
//   			segmentationTime: jsii.Number(123),
//   			timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   			timedMetadataPid: jsii.String("timedMetadataPid"),
//   			transportStreamId: jsii.Number(123),
//   			videoPid: jsii.String("videoPid"),
//   		},
//   	},
//   	destination: &outputLocationRefProperty{
//   		destinationRefId: jsii.String("destinationRefId"),
//   	},
//   	fecOutputSettings: &fecOutputSettingsProperty{
//   		columnDepth: jsii.Number(123),
//   		includeFec: jsii.String("includeFec"),
//   		rowLength: jsii.Number(123),
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

