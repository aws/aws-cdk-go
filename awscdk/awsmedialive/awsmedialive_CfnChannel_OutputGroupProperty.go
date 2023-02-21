package awsmedialive


// The settings for one output group.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputGroupProperty := &outputGroupProperty{
//   	name: jsii.String("name"),
//   	outputGroupSettings: &outputGroupSettingsProperty{
//   		archiveGroupSettings: &archiveGroupSettingsProperty{
//   			archiveCdnSettings: &archiveCdnSettingsProperty{
//   				archiveS3Settings: &archiveS3SettingsProperty{
//   					cannedAcl: jsii.String("cannedAcl"),
//   				},
//   			},
//   			destination: &outputLocationRefProperty{
//   				destinationRefId: jsii.String("destinationRefId"),
//   			},
//   			rolloverInterval: jsii.Number(123),
//   		},
//   		frameCaptureGroupSettings: &frameCaptureGroupSettingsProperty{
//   			destination: &outputLocationRefProperty{
//   				destinationRefId: jsii.String("destinationRefId"),
//   			},
//   			frameCaptureCdnSettings: &frameCaptureCdnSettingsProperty{
//   				frameCaptureS3Settings: &frameCaptureS3SettingsProperty{
//   					cannedAcl: jsii.String("cannedAcl"),
//   				},
//   			},
//   		},
//   		hlsGroupSettings: &hlsGroupSettingsProperty{
//   			adMarkers: []*string{
//   				jsii.String("adMarkers"),
//   			},
//   			baseUrlContent: jsii.String("baseUrlContent"),
//   			baseUrlContent1: jsii.String("baseUrlContent1"),
//   			baseUrlManifest: jsii.String("baseUrlManifest"),
//   			baseUrlManifest1: jsii.String("baseUrlManifest1"),
//   			captionLanguageMappings: []interface{}{
//   				&captionLanguageMappingProperty{
//   					captionChannel: jsii.Number(123),
//   					languageCode: jsii.String("languageCode"),
//   					languageDescription: jsii.String("languageDescription"),
//   				},
//   			},
//   			captionLanguageSetting: jsii.String("captionLanguageSetting"),
//   			clientCache: jsii.String("clientCache"),
//   			codecSpecification: jsii.String("codecSpecification"),
//   			constantIv: jsii.String("constantIv"),
//   			destination: &outputLocationRefProperty{
//   				destinationRefId: jsii.String("destinationRefId"),
//   			},
//   			directoryStructure: jsii.String("directoryStructure"),
//   			discontinuityTags: jsii.String("discontinuityTags"),
//   			encryptionType: jsii.String("encryptionType"),
//   			hlsCdnSettings: &hlsCdnSettingsProperty{
//   				hlsAkamaiSettings: &hlsAkamaiSettingsProperty{
//   					connectionRetryInterval: jsii.Number(123),
//   					filecacheDuration: jsii.Number(123),
//   					httpTransferMode: jsii.String("httpTransferMode"),
//   					numRetries: jsii.Number(123),
//   					restartDelay: jsii.Number(123),
//   					salt: jsii.String("salt"),
//   					token: jsii.String("token"),
//   				},
//   				hlsBasicPutSettings: &hlsBasicPutSettingsProperty{
//   					connectionRetryInterval: jsii.Number(123),
//   					filecacheDuration: jsii.Number(123),
//   					numRetries: jsii.Number(123),
//   					restartDelay: jsii.Number(123),
//   				},
//   				hlsMediaStoreSettings: &hlsMediaStoreSettingsProperty{
//   					connectionRetryInterval: jsii.Number(123),
//   					filecacheDuration: jsii.Number(123),
//   					mediaStoreStorageClass: jsii.String("mediaStoreStorageClass"),
//   					numRetries: jsii.Number(123),
//   					restartDelay: jsii.Number(123),
//   				},
//   				hlsS3Settings: &hlsS3SettingsProperty{
//   					cannedAcl: jsii.String("cannedAcl"),
//   				},
//   				hlsWebdavSettings: &hlsWebdavSettingsProperty{
//   					connectionRetryInterval: jsii.Number(123),
//   					filecacheDuration: jsii.Number(123),
//   					httpTransferMode: jsii.String("httpTransferMode"),
//   					numRetries: jsii.Number(123),
//   					restartDelay: jsii.Number(123),
//   				},
//   			},
//   			hlsId3SegmentTagging: jsii.String("hlsId3SegmentTagging"),
//   			iFrameOnlyPlaylists: jsii.String("iFrameOnlyPlaylists"),
//   			incompleteSegmentBehavior: jsii.String("incompleteSegmentBehavior"),
//   			indexNSegments: jsii.Number(123),
//   			inputLossAction: jsii.String("inputLossAction"),
//   			ivInManifest: jsii.String("ivInManifest"),
//   			ivSource: jsii.String("ivSource"),
//   			keepSegments: jsii.Number(123),
//   			keyFormat: jsii.String("keyFormat"),
//   			keyFormatVersions: jsii.String("keyFormatVersions"),
//   			keyProviderSettings: &keyProviderSettingsProperty{
//   				staticKeySettings: &staticKeySettingsProperty{
//   					keyProviderServer: &inputLocationProperty{
//   						passwordParam: jsii.String("passwordParam"),
//   						uri: jsii.String("uri"),
//   						username: jsii.String("username"),
//   					},
//   					staticKeyValue: jsii.String("staticKeyValue"),
//   				},
//   			},
//   			manifestCompression: jsii.String("manifestCompression"),
//   			manifestDurationFormat: jsii.String("manifestDurationFormat"),
//   			minSegmentLength: jsii.Number(123),
//   			mode: jsii.String("mode"),
//   			outputSelection: jsii.String("outputSelection"),
//   			programDateTime: jsii.String("programDateTime"),
//   			programDateTimeClock: jsii.String("programDateTimeClock"),
//   			programDateTimePeriod: jsii.Number(123),
//   			redundantManifest: jsii.String("redundantManifest"),
//   			segmentationMode: jsii.String("segmentationMode"),
//   			segmentLength: jsii.Number(123),
//   			segmentsPerSubdirectory: jsii.Number(123),
//   			streamInfResolution: jsii.String("streamInfResolution"),
//   			timedMetadataId3Frame: jsii.String("timedMetadataId3Frame"),
//   			timedMetadataId3Period: jsii.Number(123),
//   			timestampDeltaMilliseconds: jsii.Number(123),
//   			tsFileMode: jsii.String("tsFileMode"),
//   		},
//   		mediaPackageGroupSettings: &mediaPackageGroupSettingsProperty{
//   			destination: &outputLocationRefProperty{
//   				destinationRefId: jsii.String("destinationRefId"),
//   			},
//   		},
//   		msSmoothGroupSettings: &msSmoothGroupSettingsProperty{
//   			acquisitionPointId: jsii.String("acquisitionPointId"),
//   			audioOnlyTimecodeControl: jsii.String("audioOnlyTimecodeControl"),
//   			certificateMode: jsii.String("certificateMode"),
//   			connectionRetryInterval: jsii.Number(123),
//   			destination: &outputLocationRefProperty{
//   				destinationRefId: jsii.String("destinationRefId"),
//   			},
//   			eventId: jsii.String("eventId"),
//   			eventIdMode: jsii.String("eventIdMode"),
//   			eventStopBehavior: jsii.String("eventStopBehavior"),
//   			filecacheDuration: jsii.Number(123),
//   			fragmentLength: jsii.Number(123),
//   			inputLossAction: jsii.String("inputLossAction"),
//   			numRetries: jsii.Number(123),
//   			restartDelay: jsii.Number(123),
//   			segmentationMode: jsii.String("segmentationMode"),
//   			sendDelayMs: jsii.Number(123),
//   			sparseTrackType: jsii.String("sparseTrackType"),
//   			streamManifestBehavior: jsii.String("streamManifestBehavior"),
//   			timestampOffset: jsii.String("timestampOffset"),
//   			timestampOffsetMode: jsii.String("timestampOffsetMode"),
//   		},
//   		multiplexGroupSettings: &multiplexGroupSettingsProperty{
//   		},
//   		rtmpGroupSettings: &rtmpGroupSettingsProperty{
//   			adMarkers: []*string{
//   				jsii.String("adMarkers"),
//   			},
//   			authenticationScheme: jsii.String("authenticationScheme"),
//   			cacheFullBehavior: jsii.String("cacheFullBehavior"),
//   			cacheLength: jsii.Number(123),
//   			captionData: jsii.String("captionData"),
//   			inputLossAction: jsii.String("inputLossAction"),
//   			restartDelay: jsii.Number(123),
//   		},
//   		udpGroupSettings: &udpGroupSettingsProperty{
//   			inputLossAction: jsii.String("inputLossAction"),
//   			timedMetadataId3Frame: jsii.String("timedMetadataId3Frame"),
//   			timedMetadataId3Period: jsii.Number(123),
//   		},
//   	},
//   	outputs: []interface{}{
//   		&outputProperty{
//   			audioDescriptionNames: []*string{
//   				jsii.String("audioDescriptionNames"),
//   			},
//   			captionDescriptionNames: []*string{
//   				jsii.String("captionDescriptionNames"),
//   			},
//   			outputName: jsii.String("outputName"),
//   			outputSettings: &outputSettingsProperty{
//   				archiveOutputSettings: &archiveOutputSettingsProperty{
//   					containerSettings: &archiveContainerSettingsProperty{
//   						m2TsSettings: &m2tsSettingsProperty{
//   							absentInputAudioBehavior: jsii.String("absentInputAudioBehavior"),
//   							arib: jsii.String("arib"),
//   							aribCaptionsPid: jsii.String("aribCaptionsPid"),
//   							aribCaptionsPidControl: jsii.String("aribCaptionsPidControl"),
//   							audioBufferModel: jsii.String("audioBufferModel"),
//   							audioFramesPerPes: jsii.Number(123),
//   							audioPids: jsii.String("audioPids"),
//   							audioStreamType: jsii.String("audioStreamType"),
//   							bitrate: jsii.Number(123),
//   							bufferModel: jsii.String("bufferModel"),
//   							ccDescriptor: jsii.String("ccDescriptor"),
//   							dvbNitSettings: &dvbNitSettingsProperty{
//   								networkId: jsii.Number(123),
//   								networkName: jsii.String("networkName"),
//   								repInterval: jsii.Number(123),
//   							},
//   							dvbSdtSettings: &dvbSdtSettingsProperty{
//   								outputSdt: jsii.String("outputSdt"),
//   								repInterval: jsii.Number(123),
//   								serviceName: jsii.String("serviceName"),
//   								serviceProviderName: jsii.String("serviceProviderName"),
//   							},
//   							dvbSubPids: jsii.String("dvbSubPids"),
//   							dvbTdtSettings: &dvbTdtSettingsProperty{
//   								repInterval: jsii.Number(123),
//   							},
//   							dvbTeletextPid: jsii.String("dvbTeletextPid"),
//   							ebif: jsii.String("ebif"),
//   							ebpAudioInterval: jsii.String("ebpAudioInterval"),
//   							ebpLookaheadMs: jsii.Number(123),
//   							ebpPlacement: jsii.String("ebpPlacement"),
//   							ecmPid: jsii.String("ecmPid"),
//   							esRateInPes: jsii.String("esRateInPes"),
//   							etvPlatformPid: jsii.String("etvPlatformPid"),
//   							etvSignalPid: jsii.String("etvSignalPid"),
//   							fragmentTime: jsii.Number(123),
//   							klv: jsii.String("klv"),
//   							klvDataPids: jsii.String("klvDataPids"),
//   							nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   							nullPacketBitrate: jsii.Number(123),
//   							patInterval: jsii.Number(123),
//   							pcrControl: jsii.String("pcrControl"),
//   							pcrPeriod: jsii.Number(123),
//   							pcrPid: jsii.String("pcrPid"),
//   							pmtInterval: jsii.Number(123),
//   							pmtPid: jsii.String("pmtPid"),
//   							programNum: jsii.Number(123),
//   							rateMode: jsii.String("rateMode"),
//   							scte27Pids: jsii.String("scte27Pids"),
//   							scte35Control: jsii.String("scte35Control"),
//   							scte35Pid: jsii.String("scte35Pid"),
//   							segmentationMarkers: jsii.String("segmentationMarkers"),
//   							segmentationStyle: jsii.String("segmentationStyle"),
//   							segmentationTime: jsii.Number(123),
//   							timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   							timedMetadataPid: jsii.String("timedMetadataPid"),
//   							transportStreamId: jsii.Number(123),
//   							videoPid: jsii.String("videoPid"),
//   						},
//   						rawSettings: &rawSettingsProperty{
//   						},
//   					},
//   					extension: jsii.String("extension"),
//   					nameModifier: jsii.String("nameModifier"),
//   				},
//   				frameCaptureOutputSettings: &frameCaptureOutputSettingsProperty{
//   					nameModifier: jsii.String("nameModifier"),
//   				},
//   				hlsOutputSettings: &hlsOutputSettingsProperty{
//   					h265PackagingType: jsii.String("h265PackagingType"),
//   					hlsSettings: &hlsSettingsProperty{
//   						audioOnlyHlsSettings: &audioOnlyHlsSettingsProperty{
//   							audioGroupId: jsii.String("audioGroupId"),
//   							audioOnlyImage: &inputLocationProperty{
//   								passwordParam: jsii.String("passwordParam"),
//   								uri: jsii.String("uri"),
//   								username: jsii.String("username"),
//   							},
//   							audioTrackType: jsii.String("audioTrackType"),
//   							segmentType: jsii.String("segmentType"),
//   						},
//   						fmp4HlsSettings: &fmp4HlsSettingsProperty{
//   							audioRenditionSets: jsii.String("audioRenditionSets"),
//   							nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   							timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   						},
//   						frameCaptureHlsSettings: &frameCaptureHlsSettingsProperty{
//   						},
//   						standardHlsSettings: &standardHlsSettingsProperty{
//   							audioRenditionSets: jsii.String("audioRenditionSets"),
//   							m3U8Settings: &m3u8SettingsProperty{
//   								audioFramesPerPes: jsii.Number(123),
//   								audioPids: jsii.String("audioPids"),
//   								ecmPid: jsii.String("ecmPid"),
//   								nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   								patInterval: jsii.Number(123),
//   								pcrControl: jsii.String("pcrControl"),
//   								pcrPeriod: jsii.Number(123),
//   								pcrPid: jsii.String("pcrPid"),
//   								pmtInterval: jsii.Number(123),
//   								pmtPid: jsii.String("pmtPid"),
//   								programNum: jsii.Number(123),
//   								scte35Behavior: jsii.String("scte35Behavior"),
//   								scte35Pid: jsii.String("scte35Pid"),
//   								timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   								timedMetadataPid: jsii.String("timedMetadataPid"),
//   								transportStreamId: jsii.Number(123),
//   								videoPid: jsii.String("videoPid"),
//   							},
//   						},
//   					},
//   					nameModifier: jsii.String("nameModifier"),
//   					segmentModifier: jsii.String("segmentModifier"),
//   				},
//   				mediaPackageOutputSettings: &mediaPackageOutputSettingsProperty{
//   				},
//   				msSmoothOutputSettings: &msSmoothOutputSettingsProperty{
//   					h265PackagingType: jsii.String("h265PackagingType"),
//   					nameModifier: jsii.String("nameModifier"),
//   				},
//   				multiplexOutputSettings: &multiplexOutputSettingsProperty{
//   					destination: &outputLocationRefProperty{
//   						destinationRefId: jsii.String("destinationRefId"),
//   					},
//   				},
//   				rtmpOutputSettings: &rtmpOutputSettingsProperty{
//   					certificateMode: jsii.String("certificateMode"),
//   					connectionRetryInterval: jsii.Number(123),
//   					destination: &outputLocationRefProperty{
//   						destinationRefId: jsii.String("destinationRefId"),
//   					},
//   					numRetries: jsii.Number(123),
//   				},
//   				udpOutputSettings: &udpOutputSettingsProperty{
//   					bufferMsec: jsii.Number(123),
//   					containerSettings: &udpContainerSettingsProperty{
//   						m2TsSettings: &m2tsSettingsProperty{
//   							absentInputAudioBehavior: jsii.String("absentInputAudioBehavior"),
//   							arib: jsii.String("arib"),
//   							aribCaptionsPid: jsii.String("aribCaptionsPid"),
//   							aribCaptionsPidControl: jsii.String("aribCaptionsPidControl"),
//   							audioBufferModel: jsii.String("audioBufferModel"),
//   							audioFramesPerPes: jsii.Number(123),
//   							audioPids: jsii.String("audioPids"),
//   							audioStreamType: jsii.String("audioStreamType"),
//   							bitrate: jsii.Number(123),
//   							bufferModel: jsii.String("bufferModel"),
//   							ccDescriptor: jsii.String("ccDescriptor"),
//   							dvbNitSettings: &dvbNitSettingsProperty{
//   								networkId: jsii.Number(123),
//   								networkName: jsii.String("networkName"),
//   								repInterval: jsii.Number(123),
//   							},
//   							dvbSdtSettings: &dvbSdtSettingsProperty{
//   								outputSdt: jsii.String("outputSdt"),
//   								repInterval: jsii.Number(123),
//   								serviceName: jsii.String("serviceName"),
//   								serviceProviderName: jsii.String("serviceProviderName"),
//   							},
//   							dvbSubPids: jsii.String("dvbSubPids"),
//   							dvbTdtSettings: &dvbTdtSettingsProperty{
//   								repInterval: jsii.Number(123),
//   							},
//   							dvbTeletextPid: jsii.String("dvbTeletextPid"),
//   							ebif: jsii.String("ebif"),
//   							ebpAudioInterval: jsii.String("ebpAudioInterval"),
//   							ebpLookaheadMs: jsii.Number(123),
//   							ebpPlacement: jsii.String("ebpPlacement"),
//   							ecmPid: jsii.String("ecmPid"),
//   							esRateInPes: jsii.String("esRateInPes"),
//   							etvPlatformPid: jsii.String("etvPlatformPid"),
//   							etvSignalPid: jsii.String("etvSignalPid"),
//   							fragmentTime: jsii.Number(123),
//   							klv: jsii.String("klv"),
//   							klvDataPids: jsii.String("klvDataPids"),
//   							nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   							nullPacketBitrate: jsii.Number(123),
//   							patInterval: jsii.Number(123),
//   							pcrControl: jsii.String("pcrControl"),
//   							pcrPeriod: jsii.Number(123),
//   							pcrPid: jsii.String("pcrPid"),
//   							pmtInterval: jsii.Number(123),
//   							pmtPid: jsii.String("pmtPid"),
//   							programNum: jsii.Number(123),
//   							rateMode: jsii.String("rateMode"),
//   							scte27Pids: jsii.String("scte27Pids"),
//   							scte35Control: jsii.String("scte35Control"),
//   							scte35Pid: jsii.String("scte35Pid"),
//   							segmentationMarkers: jsii.String("segmentationMarkers"),
//   							segmentationStyle: jsii.String("segmentationStyle"),
//   							segmentationTime: jsii.Number(123),
//   							timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   							timedMetadataPid: jsii.String("timedMetadataPid"),
//   							transportStreamId: jsii.Number(123),
//   							videoPid: jsii.String("videoPid"),
//   						},
//   					},
//   					destination: &outputLocationRefProperty{
//   						destinationRefId: jsii.String("destinationRefId"),
//   					},
//   					fecOutputSettings: &fecOutputSettingsProperty{
//   						columnDepth: jsii.Number(123),
//   						includeFec: jsii.String("includeFec"),
//   						rowLength: jsii.Number(123),
//   					},
//   				},
//   			},
//   			videoDescriptionName: jsii.String("videoDescriptionName"),
//   		},
//   	},
//   }
//
type CfnChannel_OutputGroupProperty struct {
	// A custom output group name that you can optionally define.
	//
	// Only letters, numbers, and the underscore character are allowed. The maximum length is 32 characters.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The settings associated with the output group.
	OutputGroupSettings interface{} `field:"optional" json:"outputGroupSettings" yaml:"outputGroupSettings"`
	// The settings for the outputs in the output group.
	Outputs interface{} `field:"optional" json:"outputs" yaml:"outputs"`
}

