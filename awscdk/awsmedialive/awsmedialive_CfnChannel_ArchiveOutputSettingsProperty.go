package awsmedialive


// The archive output settings.
//
// The parent of this entity is OutputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   archiveOutputSettingsProperty := &archiveOutputSettingsProperty{
//   	containerSettings: &archiveContainerSettingsProperty{
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
//   		rawSettings: &rawSettingsProperty{
//   		},
//   	},
//   	extension: jsii.String("extension"),
//   	nameModifier: jsii.String("nameModifier"),
//   }
//
type CfnChannel_ArchiveOutputSettingsProperty struct {
	// The settings that are specific to the container type of the file.
	ContainerSettings interface{} `field:"optional" json:"containerSettings" yaml:"containerSettings"`
	// The output file extension.
	//
	// If excluded, this is auto-selected from the container type.
	Extension *string `field:"optional" json:"extension" yaml:"extension"`
	// A string that is concatenated to the end of the destination file name.
	//
	// The string is required for multiple outputs of the same type.
	NameModifier *string `field:"optional" json:"nameModifier" yaml:"nameModifier"`
}

