package awsmedialive


// The output settings.
//
// The parent of this entity is OutputGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputProperty := &outputProperty{
//   	audioDescriptionNames: []*string{
//   		jsii.String("audioDescriptionNames"),
//   	},
//   	captionDescriptionNames: []*string{
//   		jsii.String("captionDescriptionNames"),
//   	},
//   	outputName: jsii.String("outputName"),
//   	outputSettings: &outputSettingsProperty{
//   		archiveOutputSettings: &archiveOutputSettingsProperty{
//   			containerSettings: &archiveContainerSettingsProperty{
//   				m2TsSettings: &m2tsSettingsProperty{
//   					absentInputAudioBehavior: jsii.String("absentInputAudioBehavior"),
//   					arib: jsii.String("arib"),
//   					aribCaptionsPid: jsii.String("aribCaptionsPid"),
//   					aribCaptionsPidControl: jsii.String("aribCaptionsPidControl"),
//   					audioBufferModel: jsii.String("audioBufferModel"),
//   					audioFramesPerPes: jsii.Number(123),
//   					audioPids: jsii.String("audioPids"),
//   					audioStreamType: jsii.String("audioStreamType"),
//   					bitrate: jsii.Number(123),
//   					bufferModel: jsii.String("bufferModel"),
//   					ccDescriptor: jsii.String("ccDescriptor"),
//   					dvbNitSettings: &dvbNitSettingsProperty{
//   						networkId: jsii.Number(123),
//   						networkName: jsii.String("networkName"),
//   						repInterval: jsii.Number(123),
//   					},
//   					dvbSdtSettings: &dvbSdtSettingsProperty{
//   						outputSdt: jsii.String("outputSdt"),
//   						repInterval: jsii.Number(123),
//   						serviceName: jsii.String("serviceName"),
//   						serviceProviderName: jsii.String("serviceProviderName"),
//   					},
//   					dvbSubPids: jsii.String("dvbSubPids"),
//   					dvbTdtSettings: &dvbTdtSettingsProperty{
//   						repInterval: jsii.Number(123),
//   					},
//   					dvbTeletextPid: jsii.String("dvbTeletextPid"),
//   					ebif: jsii.String("ebif"),
//   					ebpAudioInterval: jsii.String("ebpAudioInterval"),
//   					ebpLookaheadMs: jsii.Number(123),
//   					ebpPlacement: jsii.String("ebpPlacement"),
//   					ecmPid: jsii.String("ecmPid"),
//   					esRateInPes: jsii.String("esRateInPes"),
//   					etvPlatformPid: jsii.String("etvPlatformPid"),
//   					etvSignalPid: jsii.String("etvSignalPid"),
//   					fragmentTime: jsii.Number(123),
//   					klv: jsii.String("klv"),
//   					klvDataPids: jsii.String("klvDataPids"),
//   					nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   					nullPacketBitrate: jsii.Number(123),
//   					patInterval: jsii.Number(123),
//   					pcrControl: jsii.String("pcrControl"),
//   					pcrPeriod: jsii.Number(123),
//   					pcrPid: jsii.String("pcrPid"),
//   					pmtInterval: jsii.Number(123),
//   					pmtPid: jsii.String("pmtPid"),
//   					programNum: jsii.Number(123),
//   					rateMode: jsii.String("rateMode"),
//   					scte27Pids: jsii.String("scte27Pids"),
//   					scte35Control: jsii.String("scte35Control"),
//   					scte35Pid: jsii.String("scte35Pid"),
//   					segmentationMarkers: jsii.String("segmentationMarkers"),
//   					segmentationStyle: jsii.String("segmentationStyle"),
//   					segmentationTime: jsii.Number(123),
//   					timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   					timedMetadataPid: jsii.String("timedMetadataPid"),
//   					transportStreamId: jsii.Number(123),
//   					videoPid: jsii.String("videoPid"),
//   				},
//   				rawSettings: &rawSettingsProperty{
//   				},
//   			},
//   			extension: jsii.String("extension"),
//   			nameModifier: jsii.String("nameModifier"),
//   		},
//   		frameCaptureOutputSettings: &frameCaptureOutputSettingsProperty{
//   			nameModifier: jsii.String("nameModifier"),
//   		},
//   		hlsOutputSettings: &hlsOutputSettingsProperty{
//   			h265PackagingType: jsii.String("h265PackagingType"),
//   			hlsSettings: &hlsSettingsProperty{
//   				audioOnlyHlsSettings: &audioOnlyHlsSettingsProperty{
//   					audioGroupId: jsii.String("audioGroupId"),
//   					audioOnlyImage: &inputLocationProperty{
//   						passwordParam: jsii.String("passwordParam"),
//   						uri: jsii.String("uri"),
//   						username: jsii.String("username"),
//   					},
//   					audioTrackType: jsii.String("audioTrackType"),
//   					segmentType: jsii.String("segmentType"),
//   				},
//   				fmp4HlsSettings: &fmp4HlsSettingsProperty{
//   					audioRenditionSets: jsii.String("audioRenditionSets"),
//   					nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   					timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   				},
//   				frameCaptureHlsSettings: &frameCaptureHlsSettingsProperty{
//   				},
//   				standardHlsSettings: &standardHlsSettingsProperty{
//   					audioRenditionSets: jsii.String("audioRenditionSets"),
//   					m3U8Settings: &m3u8SettingsProperty{
//   						audioFramesPerPes: jsii.Number(123),
//   						audioPids: jsii.String("audioPids"),
//   						ecmPid: jsii.String("ecmPid"),
//   						nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   						patInterval: jsii.Number(123),
//   						pcrControl: jsii.String("pcrControl"),
//   						pcrPeriod: jsii.Number(123),
//   						pcrPid: jsii.String("pcrPid"),
//   						pmtInterval: jsii.Number(123),
//   						pmtPid: jsii.String("pmtPid"),
//   						programNum: jsii.Number(123),
//   						scte35Behavior: jsii.String("scte35Behavior"),
//   						scte35Pid: jsii.String("scte35Pid"),
//   						timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   						timedMetadataPid: jsii.String("timedMetadataPid"),
//   						transportStreamId: jsii.Number(123),
//   						videoPid: jsii.String("videoPid"),
//   					},
//   				},
//   			},
//   			nameModifier: jsii.String("nameModifier"),
//   			segmentModifier: jsii.String("segmentModifier"),
//   		},
//   		mediaPackageOutputSettings: &mediaPackageOutputSettingsProperty{
//   		},
//   		msSmoothOutputSettings: &msSmoothOutputSettingsProperty{
//   			h265PackagingType: jsii.String("h265PackagingType"),
//   			nameModifier: jsii.String("nameModifier"),
//   		},
//   		multiplexOutputSettings: &multiplexOutputSettingsProperty{
//   			destination: &outputLocationRefProperty{
//   				destinationRefId: jsii.String("destinationRefId"),
//   			},
//   		},
//   		rtmpOutputSettings: &rtmpOutputSettingsProperty{
//   			certificateMode: jsii.String("certificateMode"),
//   			connectionRetryInterval: jsii.Number(123),
//   			destination: &outputLocationRefProperty{
//   				destinationRefId: jsii.String("destinationRefId"),
//   			},
//   			numRetries: jsii.Number(123),
//   		},
//   		udpOutputSettings: &udpOutputSettingsProperty{
//   			bufferMsec: jsii.Number(123),
//   			containerSettings: &udpContainerSettingsProperty{
//   				m2TsSettings: &m2tsSettingsProperty{
//   					absentInputAudioBehavior: jsii.String("absentInputAudioBehavior"),
//   					arib: jsii.String("arib"),
//   					aribCaptionsPid: jsii.String("aribCaptionsPid"),
//   					aribCaptionsPidControl: jsii.String("aribCaptionsPidControl"),
//   					audioBufferModel: jsii.String("audioBufferModel"),
//   					audioFramesPerPes: jsii.Number(123),
//   					audioPids: jsii.String("audioPids"),
//   					audioStreamType: jsii.String("audioStreamType"),
//   					bitrate: jsii.Number(123),
//   					bufferModel: jsii.String("bufferModel"),
//   					ccDescriptor: jsii.String("ccDescriptor"),
//   					dvbNitSettings: &dvbNitSettingsProperty{
//   						networkId: jsii.Number(123),
//   						networkName: jsii.String("networkName"),
//   						repInterval: jsii.Number(123),
//   					},
//   					dvbSdtSettings: &dvbSdtSettingsProperty{
//   						outputSdt: jsii.String("outputSdt"),
//   						repInterval: jsii.Number(123),
//   						serviceName: jsii.String("serviceName"),
//   						serviceProviderName: jsii.String("serviceProviderName"),
//   					},
//   					dvbSubPids: jsii.String("dvbSubPids"),
//   					dvbTdtSettings: &dvbTdtSettingsProperty{
//   						repInterval: jsii.Number(123),
//   					},
//   					dvbTeletextPid: jsii.String("dvbTeletextPid"),
//   					ebif: jsii.String("ebif"),
//   					ebpAudioInterval: jsii.String("ebpAudioInterval"),
//   					ebpLookaheadMs: jsii.Number(123),
//   					ebpPlacement: jsii.String("ebpPlacement"),
//   					ecmPid: jsii.String("ecmPid"),
//   					esRateInPes: jsii.String("esRateInPes"),
//   					etvPlatformPid: jsii.String("etvPlatformPid"),
//   					etvSignalPid: jsii.String("etvSignalPid"),
//   					fragmentTime: jsii.Number(123),
//   					klv: jsii.String("klv"),
//   					klvDataPids: jsii.String("klvDataPids"),
//   					nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   					nullPacketBitrate: jsii.Number(123),
//   					patInterval: jsii.Number(123),
//   					pcrControl: jsii.String("pcrControl"),
//   					pcrPeriod: jsii.Number(123),
//   					pcrPid: jsii.String("pcrPid"),
//   					pmtInterval: jsii.Number(123),
//   					pmtPid: jsii.String("pmtPid"),
//   					programNum: jsii.Number(123),
//   					rateMode: jsii.String("rateMode"),
//   					scte27Pids: jsii.String("scte27Pids"),
//   					scte35Control: jsii.String("scte35Control"),
//   					scte35Pid: jsii.String("scte35Pid"),
//   					segmentationMarkers: jsii.String("segmentationMarkers"),
//   					segmentationStyle: jsii.String("segmentationStyle"),
//   					segmentationTime: jsii.Number(123),
//   					timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   					timedMetadataPid: jsii.String("timedMetadataPid"),
//   					transportStreamId: jsii.Number(123),
//   					videoPid: jsii.String("videoPid"),
//   				},
//   			},
//   			destination: &outputLocationRefProperty{
//   				destinationRefId: jsii.String("destinationRefId"),
//   			},
//   			fecOutputSettings: &fecOutputSettingsProperty{
//   				columnDepth: jsii.Number(123),
//   				includeFec: jsii.String("includeFec"),
//   				rowLength: jsii.Number(123),
//   			},
//   		},
//   	},
//   	videoDescriptionName: jsii.String("videoDescriptionName"),
//   }
//
type CfnChannel_OutputProperty struct {
	// The names of the audio descriptions that are used as audio sources for this output.
	AudioDescriptionNames *[]*string `field:"optional" json:"audioDescriptionNames" yaml:"audioDescriptionNames"`
	// The names of the caption descriptions that are used as captions sources for this output.
	CaptionDescriptionNames *[]*string `field:"optional" json:"captionDescriptionNames" yaml:"captionDescriptionNames"`
	// The name that is used to identify an output.
	OutputName *string `field:"optional" json:"outputName" yaml:"outputName"`
	// The output type-specific settings.
	OutputSettings interface{} `field:"optional" json:"outputSettings" yaml:"outputSettings"`
	// The name of the VideoDescription that is used as the source for this output.
	VideoDescriptionName *string `field:"optional" json:"videoDescriptionName" yaml:"videoDescriptionName"`
}

