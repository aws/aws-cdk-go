package awsmedialive


// The configuration of a UDP output.
//
// The parent of this entity is UdpOutputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   udpContainerSettingsProperty := &udpContainerSettingsProperty{
//   	m2TsSettings: &m2tsSettingsProperty{
//   		absentInputAudioBehavior: jsii.String("absentInputAudioBehavior"),
//   		arib: jsii.String("arib"),
//   		aribCaptionsPid: jsii.String("aribCaptionsPid"),
//   		aribCaptionsPidControl: jsii.String("aribCaptionsPidControl"),
//   		audioBufferModel: jsii.String("audioBufferModel"),
//   		audioFramesPerPes: jsii.Number(123),
//   		audioPids: jsii.String("audioPids"),
//   		audioStreamType: jsii.String("audioStreamType"),
//   		bitrate: jsii.Number(123),
//   		bufferModel: jsii.String("bufferModel"),
//   		ccDescriptor: jsii.String("ccDescriptor"),
//   		dvbNitSettings: &dvbNitSettingsProperty{
//   			networkId: jsii.Number(123),
//   			networkName: jsii.String("networkName"),
//   			repInterval: jsii.Number(123),
//   		},
//   		dvbSdtSettings: &dvbSdtSettingsProperty{
//   			outputSdt: jsii.String("outputSdt"),
//   			repInterval: jsii.Number(123),
//   			serviceName: jsii.String("serviceName"),
//   			serviceProviderName: jsii.String("serviceProviderName"),
//   		},
//   		dvbSubPids: jsii.String("dvbSubPids"),
//   		dvbTdtSettings: &dvbTdtSettingsProperty{
//   			repInterval: jsii.Number(123),
//   		},
//   		dvbTeletextPid: jsii.String("dvbTeletextPid"),
//   		ebif: jsii.String("ebif"),
//   		ebpAudioInterval: jsii.String("ebpAudioInterval"),
//   		ebpLookaheadMs: jsii.Number(123),
//   		ebpPlacement: jsii.String("ebpPlacement"),
//   		ecmPid: jsii.String("ecmPid"),
//   		esRateInPes: jsii.String("esRateInPes"),
//   		etvPlatformPid: jsii.String("etvPlatformPid"),
//   		etvSignalPid: jsii.String("etvSignalPid"),
//   		fragmentTime: jsii.Number(123),
//   		klv: jsii.String("klv"),
//   		klvDataPids: jsii.String("klvDataPids"),
//   		nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   		nullPacketBitrate: jsii.Number(123),
//   		patInterval: jsii.Number(123),
//   		pcrControl: jsii.String("pcrControl"),
//   		pcrPeriod: jsii.Number(123),
//   		pcrPid: jsii.String("pcrPid"),
//   		pmtInterval: jsii.Number(123),
//   		pmtPid: jsii.String("pmtPid"),
//   		programNum: jsii.Number(123),
//   		rateMode: jsii.String("rateMode"),
//   		scte27Pids: jsii.String("scte27Pids"),
//   		scte35Control: jsii.String("scte35Control"),
//   		scte35Pid: jsii.String("scte35Pid"),
//   		segmentationMarkers: jsii.String("segmentationMarkers"),
//   		segmentationStyle: jsii.String("segmentationStyle"),
//   		segmentationTime: jsii.Number(123),
//   		timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   		timedMetadataPid: jsii.String("timedMetadataPid"),
//   		transportStreamId: jsii.Number(123),
//   		videoPid: jsii.String("videoPid"),
//   	},
//   }
//
type CfnChannel_UdpContainerSettingsProperty struct {
	// The M2TS configuration for this UDP output.
	M2TsSettings interface{} `field:"optional" json:"m2TsSettings" yaml:"m2TsSettings"`
}

