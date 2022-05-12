package awsmedialive

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::MediaLive::Channel`.
//
// The AWS::MediaLive::Channel resource is a MediaLive resource type that creates a channel.
//
// A MediaLive channel ingests and transcodes (decodes and encodes) source content from the inputs that are attached to that channel, and packages the new content into outputs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnChannel := awscdk.Aws_medialive.NewCfnChannel(this, jsii.String("MyCfnChannel"), &cfnChannelProps{
//   	cdiInputSpecification: &cdiInputSpecificationProperty{
//   		resolution: jsii.String("resolution"),
//   	},
//   	channelClass: jsii.String("channelClass"),
//   	destinations: []interface{}{
//   		&outputDestinationProperty{
//   			id: jsii.String("id"),
//   			mediaPackageSettings: []interface{}{
//   				&mediaPackageOutputDestinationSettingsProperty{
//   					channelId: jsii.String("channelId"),
//   				},
//   			},
//   			multiplexSettings: &multiplexProgramChannelDestinationSettingsProperty{
//   				multiplexId: jsii.String("multiplexId"),
//   				programName: jsii.String("programName"),
//   			},
//   			settings: []interface{}{
//   				&outputDestinationSettingsProperty{
//   					passwordParam: jsii.String("passwordParam"),
//   					streamName: jsii.String("streamName"),
//   					url: jsii.String("url"),
//   					username: jsii.String("username"),
//   				},
//   			},
//   		},
//   	},
//   	encoderSettings: &encoderSettingsProperty{
//   		audioDescriptions: []interface{}{
//   			&audioDescriptionProperty{
//   				audioNormalizationSettings: &audioNormalizationSettingsProperty{
//   					algorithm: jsii.String("algorithm"),
//   					algorithmControl: jsii.String("algorithmControl"),
//   					targetLkfs: jsii.Number(123),
//   				},
//   				audioSelectorName: jsii.String("audioSelectorName"),
//   				audioType: jsii.String("audioType"),
//   				audioTypeControl: jsii.String("audioTypeControl"),
//   				audioWatermarkingSettings: &audioWatermarkSettingsProperty{
//   					nielsenWatermarksSettings: &nielsenWatermarksSettingsProperty{
//   						nielsenCbetSettings: &nielsenCBETProperty{
//   							cbetCheckDigitString: jsii.String("cbetCheckDigitString"),
//   							cbetStepaside: jsii.String("cbetStepaside"),
//   							csid: jsii.String("csid"),
//   						},
//   						nielsenDistributionType: jsii.String("nielsenDistributionType"),
//   						nielsenNaesIiNwSettings: &nielsenNaesIiNwProperty{
//   							checkDigitString: jsii.String("checkDigitString"),
//   							sid: jsii.Number(123),
//   						},
//   					},
//   				},
//   				codecSettings: &audioCodecSettingsProperty{
//   					aacSettings: &aacSettingsProperty{
//   						bitrate: jsii.Number(123),
//   						codingMode: jsii.String("codingMode"),
//   						inputType: jsii.String("inputType"),
//   						profile: jsii.String("profile"),
//   						rateControlMode: jsii.String("rateControlMode"),
//   						rawFormat: jsii.String("rawFormat"),
//   						sampleRate: jsii.Number(123),
//   						spec: jsii.String("spec"),
//   						vbrQuality: jsii.String("vbrQuality"),
//   					},
//   					ac3Settings: &ac3SettingsProperty{
//   						bitrate: jsii.Number(123),
//   						bitstreamMode: jsii.String("bitstreamMode"),
//   						codingMode: jsii.String("codingMode"),
//   						dialnorm: jsii.Number(123),
//   						drcProfile: jsii.String("drcProfile"),
//   						lfeFilter: jsii.String("lfeFilter"),
//   						metadataControl: jsii.String("metadataControl"),
//   					},
//   					eac3Settings: &eac3SettingsProperty{
//   						attenuationControl: jsii.String("attenuationControl"),
//   						bitrate: jsii.Number(123),
//   						bitstreamMode: jsii.String("bitstreamMode"),
//   						codingMode: jsii.String("codingMode"),
//   						dcFilter: jsii.String("dcFilter"),
//   						dialnorm: jsii.Number(123),
//   						drcLine: jsii.String("drcLine"),
//   						drcRf: jsii.String("drcRf"),
//   						lfeControl: jsii.String("lfeControl"),
//   						lfeFilter: jsii.String("lfeFilter"),
//   						loRoCenterMixLevel: jsii.Number(123),
//   						loRoSurroundMixLevel: jsii.Number(123),
//   						ltRtCenterMixLevel: jsii.Number(123),
//   						ltRtSurroundMixLevel: jsii.Number(123),
//   						metadataControl: jsii.String("metadataControl"),
//   						passthroughControl: jsii.String("passthroughControl"),
//   						phaseControl: jsii.String("phaseControl"),
//   						stereoDownmix: jsii.String("stereoDownmix"),
//   						surroundExMode: jsii.String("surroundExMode"),
//   						surroundMode: jsii.String("surroundMode"),
//   					},
//   					mp2Settings: &mp2SettingsProperty{
//   						bitrate: jsii.Number(123),
//   						codingMode: jsii.String("codingMode"),
//   						sampleRate: jsii.Number(123),
//   					},
//   					passThroughSettings: &passThroughSettingsProperty{
//   					},
//   					wavSettings: &wavSettingsProperty{
//   						bitDepth: jsii.Number(123),
//   						codingMode: jsii.String("codingMode"),
//   						sampleRate: jsii.Number(123),
//   					},
//   				},
//   				languageCode: jsii.String("languageCode"),
//   				languageCodeControl: jsii.String("languageCodeControl"),
//   				name: jsii.String("name"),
//   				remixSettings: &remixSettingsProperty{
//   					channelMappings: []interface{}{
//   						&audioChannelMappingProperty{
//   							inputChannelLevels: []interface{}{
//   								&inputChannelLevelProperty{
//   									gain: jsii.Number(123),
//   									inputChannel: jsii.Number(123),
//   								},
//   							},
//   							outputChannel: jsii.Number(123),
//   						},
//   					},
//   					channelsIn: jsii.Number(123),
//   					channelsOut: jsii.Number(123),
//   				},
//   				streamName: jsii.String("streamName"),
//   			},
//   		},
//   		availBlanking: &availBlankingProperty{
//   			availBlankingImage: &inputLocationProperty{
//   				passwordParam: jsii.String("passwordParam"),
//   				uri: jsii.String("uri"),
//   				username: jsii.String("username"),
//   			},
//   			state: jsii.String("state"),
//   		},
//   		availConfiguration: &availConfigurationProperty{
//   			availSettings: &availSettingsProperty{
//   				scte35SpliceInsert: &scte35SpliceInsertProperty{
//   					adAvailOffset: jsii.Number(123),
//   					noRegionalBlackoutFlag: jsii.String("noRegionalBlackoutFlag"),
//   					webDeliveryAllowedFlag: jsii.String("webDeliveryAllowedFlag"),
//   				},
//   				scte35TimeSignalApos: &scte35TimeSignalAposProperty{
//   					adAvailOffset: jsii.Number(123),
//   					noRegionalBlackoutFlag: jsii.String("noRegionalBlackoutFlag"),
//   					webDeliveryAllowedFlag: jsii.String("webDeliveryAllowedFlag"),
//   				},
//   			},
//   		},
//   		blackoutSlate: &blackoutSlateProperty{
//   			blackoutSlateImage: &inputLocationProperty{
//   				passwordParam: jsii.String("passwordParam"),
//   				uri: jsii.String("uri"),
//   				username: jsii.String("username"),
//   			},
//   			networkEndBlackout: jsii.String("networkEndBlackout"),
//   			networkEndBlackoutImage: &inputLocationProperty{
//   				passwordParam: jsii.String("passwordParam"),
//   				uri: jsii.String("uri"),
//   				username: jsii.String("username"),
//   			},
//   			networkId: jsii.String("networkId"),
//   			state: jsii.String("state"),
//   		},
//   		captionDescriptions: []interface{}{
//   			&captionDescriptionProperty{
//   				captionSelectorName: jsii.String("captionSelectorName"),
//   				destinationSettings: &captionDestinationSettingsProperty{
//   					aribDestinationSettings: &aribDestinationSettingsProperty{
//   					},
//   					burnInDestinationSettings: &burnInDestinationSettingsProperty{
//   						alignment: jsii.String("alignment"),
//   						backgroundColor: jsii.String("backgroundColor"),
//   						backgroundOpacity: jsii.Number(123),
//   						font: &inputLocationProperty{
//   							passwordParam: jsii.String("passwordParam"),
//   							uri: jsii.String("uri"),
//   							username: jsii.String("username"),
//   						},
//   						fontColor: jsii.String("fontColor"),
//   						fontOpacity: jsii.Number(123),
//   						fontResolution: jsii.Number(123),
//   						fontSize: jsii.String("fontSize"),
//   						outlineColor: jsii.String("outlineColor"),
//   						outlineSize: jsii.Number(123),
//   						shadowColor: jsii.String("shadowColor"),
//   						shadowOpacity: jsii.Number(123),
//   						shadowXOffset: jsii.Number(123),
//   						shadowYOffset: jsii.Number(123),
//   						teletextGridControl: jsii.String("teletextGridControl"),
//   						xPosition: jsii.Number(123),
//   						yPosition: jsii.Number(123),
//   					},
//   					dvbSubDestinationSettings: &dvbSubDestinationSettingsProperty{
//   						alignment: jsii.String("alignment"),
//   						backgroundColor: jsii.String("backgroundColor"),
//   						backgroundOpacity: jsii.Number(123),
//   						font: &inputLocationProperty{
//   							passwordParam: jsii.String("passwordParam"),
//   							uri: jsii.String("uri"),
//   							username: jsii.String("username"),
//   						},
//   						fontColor: jsii.String("fontColor"),
//   						fontOpacity: jsii.Number(123),
//   						fontResolution: jsii.Number(123),
//   						fontSize: jsii.String("fontSize"),
//   						outlineColor: jsii.String("outlineColor"),
//   						outlineSize: jsii.Number(123),
//   						shadowColor: jsii.String("shadowColor"),
//   						shadowOpacity: jsii.Number(123),
//   						shadowXOffset: jsii.Number(123),
//   						shadowYOffset: jsii.Number(123),
//   						teletextGridControl: jsii.String("teletextGridControl"),
//   						xPosition: jsii.Number(123),
//   						yPosition: jsii.Number(123),
//   					},
//   					ebuTtDDestinationSettings: &ebuTtDDestinationSettingsProperty{
//   						copyrightHolder: jsii.String("copyrightHolder"),
//   						fillLineGap: jsii.String("fillLineGap"),
//   						fontFamily: jsii.String("fontFamily"),
//   						styleControl: jsii.String("styleControl"),
//   					},
//   					embeddedDestinationSettings: &embeddedDestinationSettingsProperty{
//   					},
//   					embeddedPlusScte20DestinationSettings: &embeddedPlusScte20DestinationSettingsProperty{
//   					},
//   					rtmpCaptionInfoDestinationSettings: &rtmpCaptionInfoDestinationSettingsProperty{
//   					},
//   					scte20PlusEmbeddedDestinationSettings: &scte20PlusEmbeddedDestinationSettingsProperty{
//   					},
//   					scte27DestinationSettings: &scte27DestinationSettingsProperty{
//   					},
//   					smpteTtDestinationSettings: &smpteTtDestinationSettingsProperty{
//   					},
//   					teletextDestinationSettings: &teletextDestinationSettingsProperty{
//   					},
//   					ttmlDestinationSettings: &ttmlDestinationSettingsProperty{
//   						styleControl: jsii.String("styleControl"),
//   					},
//   					webvttDestinationSettings: &webvttDestinationSettingsProperty{
//   						styleControl: jsii.String("styleControl"),
//   					},
//   				},
//   				languageCode: jsii.String("languageCode"),
//   				languageDescription: jsii.String("languageDescription"),
//   				name: jsii.String("name"),
//   			},
//   		},
//   		featureActivations: &featureActivationsProperty{
//   			inputPrepareScheduleActions: jsii.String("inputPrepareScheduleActions"),
//   		},
//   		globalConfiguration: &globalConfigurationProperty{
//   			initialAudioGain: jsii.Number(123),
//   			inputEndAction: jsii.String("inputEndAction"),
//   			inputLossBehavior: &inputLossBehaviorProperty{
//   				blackFrameMsec: jsii.Number(123),
//   				inputLossImageColor: jsii.String("inputLossImageColor"),
//   				inputLossImageSlate: &inputLocationProperty{
//   					passwordParam: jsii.String("passwordParam"),
//   					uri: jsii.String("uri"),
//   					username: jsii.String("username"),
//   				},
//   				inputLossImageType: jsii.String("inputLossImageType"),
//   				repeatFrameMsec: jsii.Number(123),
//   			},
//   			outputLockingMode: jsii.String("outputLockingMode"),
//   			outputTimingSource: jsii.String("outputTimingSource"),
//   			supportLowFramerateInputs: jsii.String("supportLowFramerateInputs"),
//   		},
//   		motionGraphicsConfiguration: &motionGraphicsConfigurationProperty{
//   			motionGraphicsInsertion: jsii.String("motionGraphicsInsertion"),
//   			motionGraphicsSettings: &motionGraphicsSettingsProperty{
//   				htmlMotionGraphicsSettings: &htmlMotionGraphicsSettingsProperty{
//   				},
//   			},
//   		},
//   		nielsenConfiguration: &nielsenConfigurationProperty{
//   			distributorId: jsii.String("distributorId"),
//   			nielsenPcmToId3Tagging: jsii.String("nielsenPcmToId3Tagging"),
//   		},
//   		outputGroups: []interface{}{
//   			&outputGroupProperty{
//   				name: jsii.String("name"),
//   				outputGroupSettings: &outputGroupSettingsProperty{
//   					archiveGroupSettings: &archiveGroupSettingsProperty{
//   						archiveCdnSettings: &archiveCdnSettingsProperty{
//   							archiveS3Settings: &archiveS3SettingsProperty{
//   								cannedAcl: jsii.String("cannedAcl"),
//   							},
//   						},
//   						destination: &outputLocationRefProperty{
//   							destinationRefId: jsii.String("destinationRefId"),
//   						},
//   						rolloverInterval: jsii.Number(123),
//   					},
//   					frameCaptureGroupSettings: &frameCaptureGroupSettingsProperty{
//   						destination: &outputLocationRefProperty{
//   							destinationRefId: jsii.String("destinationRefId"),
//   						},
//   						frameCaptureCdnSettings: &frameCaptureCdnSettingsProperty{
//   							frameCaptureS3Settings: &frameCaptureS3SettingsProperty{
//   								cannedAcl: jsii.String("cannedAcl"),
//   							},
//   						},
//   					},
//   					hlsGroupSettings: &hlsGroupSettingsProperty{
//   						adMarkers: []*string{
//   							jsii.String("adMarkers"),
//   						},
//   						baseUrlContent: jsii.String("baseUrlContent"),
//   						baseUrlContent1: jsii.String("baseUrlContent1"),
//   						baseUrlManifest: jsii.String("baseUrlManifest"),
//   						baseUrlManifest1: jsii.String("baseUrlManifest1"),
//   						captionLanguageMappings: []interface{}{
//   							&captionLanguageMappingProperty{
//   								captionChannel: jsii.Number(123),
//   								languageCode: jsii.String("languageCode"),
//   								languageDescription: jsii.String("languageDescription"),
//   							},
//   						},
//   						captionLanguageSetting: jsii.String("captionLanguageSetting"),
//   						clientCache: jsii.String("clientCache"),
//   						codecSpecification: jsii.String("codecSpecification"),
//   						constantIv: jsii.String("constantIv"),
//   						destination: &outputLocationRefProperty{
//   							destinationRefId: jsii.String("destinationRefId"),
//   						},
//   						directoryStructure: jsii.String("directoryStructure"),
//   						discontinuityTags: jsii.String("discontinuityTags"),
//   						encryptionType: jsii.String("encryptionType"),
//   						hlsCdnSettings: &hlsCdnSettingsProperty{
//   							hlsAkamaiSettings: &hlsAkamaiSettingsProperty{
//   								connectionRetryInterval: jsii.Number(123),
//   								filecacheDuration: jsii.Number(123),
//   								httpTransferMode: jsii.String("httpTransferMode"),
//   								numRetries: jsii.Number(123),
//   								restartDelay: jsii.Number(123),
//   								salt: jsii.String("salt"),
//   								token: jsii.String("token"),
//   							},
//   							hlsBasicPutSettings: &hlsBasicPutSettingsProperty{
//   								connectionRetryInterval: jsii.Number(123),
//   								filecacheDuration: jsii.Number(123),
//   								numRetries: jsii.Number(123),
//   								restartDelay: jsii.Number(123),
//   							},
//   							hlsMediaStoreSettings: &hlsMediaStoreSettingsProperty{
//   								connectionRetryInterval: jsii.Number(123),
//   								filecacheDuration: jsii.Number(123),
//   								mediaStoreStorageClass: jsii.String("mediaStoreStorageClass"),
//   								numRetries: jsii.Number(123),
//   								restartDelay: jsii.Number(123),
//   							},
//   							hlsS3Settings: &hlsS3SettingsProperty{
//   								cannedAcl: jsii.String("cannedAcl"),
//   							},
//   							hlsWebdavSettings: &hlsWebdavSettingsProperty{
//   								connectionRetryInterval: jsii.Number(123),
//   								filecacheDuration: jsii.Number(123),
//   								httpTransferMode: jsii.String("httpTransferMode"),
//   								numRetries: jsii.Number(123),
//   								restartDelay: jsii.Number(123),
//   							},
//   						},
//   						hlsId3SegmentTagging: jsii.String("hlsId3SegmentTagging"),
//   						iFrameOnlyPlaylists: jsii.String("iFrameOnlyPlaylists"),
//   						incompleteSegmentBehavior: jsii.String("incompleteSegmentBehavior"),
//   						indexNSegments: jsii.Number(123),
//   						inputLossAction: jsii.String("inputLossAction"),
//   						ivInManifest: jsii.String("ivInManifest"),
//   						ivSource: jsii.String("ivSource"),
//   						keepSegments: jsii.Number(123),
//   						keyFormat: jsii.String("keyFormat"),
//   						keyFormatVersions: jsii.String("keyFormatVersions"),
//   						keyProviderSettings: &keyProviderSettingsProperty{
//   							staticKeySettings: &staticKeySettingsProperty{
//   								keyProviderServer: &inputLocationProperty{
//   									passwordParam: jsii.String("passwordParam"),
//   									uri: jsii.String("uri"),
//   									username: jsii.String("username"),
//   								},
//   								staticKeyValue: jsii.String("staticKeyValue"),
//   							},
//   						},
//   						manifestCompression: jsii.String("manifestCompression"),
//   						manifestDurationFormat: jsii.String("manifestDurationFormat"),
//   						minSegmentLength: jsii.Number(123),
//   						mode: jsii.String("mode"),
//   						outputSelection: jsii.String("outputSelection"),
//   						programDateTime: jsii.String("programDateTime"),
//   						programDateTimeClock: jsii.String("programDateTimeClock"),
//   						programDateTimePeriod: jsii.Number(123),
//   						redundantManifest: jsii.String("redundantManifest"),
//   						segmentationMode: jsii.String("segmentationMode"),
//   						segmentLength: jsii.Number(123),
//   						segmentsPerSubdirectory: jsii.Number(123),
//   						streamInfResolution: jsii.String("streamInfResolution"),
//   						timedMetadataId3Frame: jsii.String("timedMetadataId3Frame"),
//   						timedMetadataId3Period: jsii.Number(123),
//   						timestampDeltaMilliseconds: jsii.Number(123),
//   						tsFileMode: jsii.String("tsFileMode"),
//   					},
//   					mediaPackageGroupSettings: &mediaPackageGroupSettingsProperty{
//   						destination: &outputLocationRefProperty{
//   							destinationRefId: jsii.String("destinationRefId"),
//   						},
//   					},
//   					msSmoothGroupSettings: &msSmoothGroupSettingsProperty{
//   						acquisitionPointId: jsii.String("acquisitionPointId"),
//   						audioOnlyTimecodeControl: jsii.String("audioOnlyTimecodeControl"),
//   						certificateMode: jsii.String("certificateMode"),
//   						connectionRetryInterval: jsii.Number(123),
//   						destination: &outputLocationRefProperty{
//   							destinationRefId: jsii.String("destinationRefId"),
//   						},
//   						eventId: jsii.String("eventId"),
//   						eventIdMode: jsii.String("eventIdMode"),
//   						eventStopBehavior: jsii.String("eventStopBehavior"),
//   						filecacheDuration: jsii.Number(123),
//   						fragmentLength: jsii.Number(123),
//   						inputLossAction: jsii.String("inputLossAction"),
//   						numRetries: jsii.Number(123),
//   						restartDelay: jsii.Number(123),
//   						segmentationMode: jsii.String("segmentationMode"),
//   						sendDelayMs: jsii.Number(123),
//   						sparseTrackType: jsii.String("sparseTrackType"),
//   						streamManifestBehavior: jsii.String("streamManifestBehavior"),
//   						timestampOffset: jsii.String("timestampOffset"),
//   						timestampOffsetMode: jsii.String("timestampOffsetMode"),
//   					},
//   					multiplexGroupSettings: &multiplexGroupSettingsProperty{
//   					},
//   					rtmpGroupSettings: &rtmpGroupSettingsProperty{
//   						adMarkers: []*string{
//   							jsii.String("adMarkers"),
//   						},
//   						authenticationScheme: jsii.String("authenticationScheme"),
//   						cacheFullBehavior: jsii.String("cacheFullBehavior"),
//   						cacheLength: jsii.Number(123),
//   						captionData: jsii.String("captionData"),
//   						inputLossAction: jsii.String("inputLossAction"),
//   						restartDelay: jsii.Number(123),
//   					},
//   					udpGroupSettings: &udpGroupSettingsProperty{
//   						inputLossAction: jsii.String("inputLossAction"),
//   						timedMetadataId3Frame: jsii.String("timedMetadataId3Frame"),
//   						timedMetadataId3Period: jsii.Number(123),
//   					},
//   				},
//   				outputs: []interface{}{
//   					&outputProperty{
//   						audioDescriptionNames: []*string{
//   							jsii.String("audioDescriptionNames"),
//   						},
//   						captionDescriptionNames: []*string{
//   							jsii.String("captionDescriptionNames"),
//   						},
//   						outputName: jsii.String("outputName"),
//   						outputSettings: &outputSettingsProperty{
//   							archiveOutputSettings: &archiveOutputSettingsProperty{
//   								containerSettings: &archiveContainerSettingsProperty{
//   									m2TsSettings: &m2tsSettingsProperty{
//   										absentInputAudioBehavior: jsii.String("absentInputAudioBehavior"),
//   										arib: jsii.String("arib"),
//   										aribCaptionsPid: jsii.String("aribCaptionsPid"),
//   										aribCaptionsPidControl: jsii.String("aribCaptionsPidControl"),
//   										audioBufferModel: jsii.String("audioBufferModel"),
//   										audioFramesPerPes: jsii.Number(123),
//   										audioPids: jsii.String("audioPids"),
//   										audioStreamType: jsii.String("audioStreamType"),
//   										bitrate: jsii.Number(123),
//   										bufferModel: jsii.String("bufferModel"),
//   										ccDescriptor: jsii.String("ccDescriptor"),
//   										dvbNitSettings: &dvbNitSettingsProperty{
//   											networkId: jsii.Number(123),
//   											networkName: jsii.String("networkName"),
//   											repInterval: jsii.Number(123),
//   										},
//   										dvbSdtSettings: &dvbSdtSettingsProperty{
//   											outputSdt: jsii.String("outputSdt"),
//   											repInterval: jsii.Number(123),
//   											serviceName: jsii.String("serviceName"),
//   											serviceProviderName: jsii.String("serviceProviderName"),
//   										},
//   										dvbSubPids: jsii.String("dvbSubPids"),
//   										dvbTdtSettings: &dvbTdtSettingsProperty{
//   											repInterval: jsii.Number(123),
//   										},
//   										dvbTeletextPid: jsii.String("dvbTeletextPid"),
//   										ebif: jsii.String("ebif"),
//   										ebpAudioInterval: jsii.String("ebpAudioInterval"),
//   										ebpLookaheadMs: jsii.Number(123),
//   										ebpPlacement: jsii.String("ebpPlacement"),
//   										ecmPid: jsii.String("ecmPid"),
//   										esRateInPes: jsii.String("esRateInPes"),
//   										etvPlatformPid: jsii.String("etvPlatformPid"),
//   										etvSignalPid: jsii.String("etvSignalPid"),
//   										fragmentTime: jsii.Number(123),
//   										klv: jsii.String("klv"),
//   										klvDataPids: jsii.String("klvDataPids"),
//   										nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   										nullPacketBitrate: jsii.Number(123),
//   										patInterval: jsii.Number(123),
//   										pcrControl: jsii.String("pcrControl"),
//   										pcrPeriod: jsii.Number(123),
//   										pcrPid: jsii.String("pcrPid"),
//   										pmtInterval: jsii.Number(123),
//   										pmtPid: jsii.String("pmtPid"),
//   										programNum: jsii.Number(123),
//   										rateMode: jsii.String("rateMode"),
//   										scte27Pids: jsii.String("scte27Pids"),
//   										scte35Control: jsii.String("scte35Control"),
//   										scte35Pid: jsii.String("scte35Pid"),
//   										segmentationMarkers: jsii.String("segmentationMarkers"),
//   										segmentationStyle: jsii.String("segmentationStyle"),
//   										segmentationTime: jsii.Number(123),
//   										timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   										timedMetadataPid: jsii.String("timedMetadataPid"),
//   										transportStreamId: jsii.Number(123),
//   										videoPid: jsii.String("videoPid"),
//   									},
//   									rawSettings: &rawSettingsProperty{
//   									},
//   								},
//   								extension: jsii.String("extension"),
//   								nameModifier: jsii.String("nameModifier"),
//   							},
//   							frameCaptureOutputSettings: &frameCaptureOutputSettingsProperty{
//   								nameModifier: jsii.String("nameModifier"),
//   							},
//   							hlsOutputSettings: &hlsOutputSettingsProperty{
//   								h265PackagingType: jsii.String("h265PackagingType"),
//   								hlsSettings: &hlsSettingsProperty{
//   									audioOnlyHlsSettings: &audioOnlyHlsSettingsProperty{
//   										audioGroupId: jsii.String("audioGroupId"),
//   										audioOnlyImage: &inputLocationProperty{
//   											passwordParam: jsii.String("passwordParam"),
//   											uri: jsii.String("uri"),
//   											username: jsii.String("username"),
//   										},
//   										audioTrackType: jsii.String("audioTrackType"),
//   										segmentType: jsii.String("segmentType"),
//   									},
//   									fmp4HlsSettings: &fmp4HlsSettingsProperty{
//   										audioRenditionSets: jsii.String("audioRenditionSets"),
//   										nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   										timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   									},
//   									frameCaptureHlsSettings: &frameCaptureHlsSettingsProperty{
//   									},
//   									standardHlsSettings: &standardHlsSettingsProperty{
//   										audioRenditionSets: jsii.String("audioRenditionSets"),
//   										m3U8Settings: &m3u8SettingsProperty{
//   											audioFramesPerPes: jsii.Number(123),
//   											audioPids: jsii.String("audioPids"),
//   											ecmPid: jsii.String("ecmPid"),
//   											nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   											patInterval: jsii.Number(123),
//   											pcrControl: jsii.String("pcrControl"),
//   											pcrPeriod: jsii.Number(123),
//   											pcrPid: jsii.String("pcrPid"),
//   											pmtInterval: jsii.Number(123),
//   											pmtPid: jsii.String("pmtPid"),
//   											programNum: jsii.Number(123),
//   											scte35Behavior: jsii.String("scte35Behavior"),
//   											scte35Pid: jsii.String("scte35Pid"),
//   											timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   											timedMetadataPid: jsii.String("timedMetadataPid"),
//   											transportStreamId: jsii.Number(123),
//   											videoPid: jsii.String("videoPid"),
//   										},
//   									},
//   								},
//   								nameModifier: jsii.String("nameModifier"),
//   								segmentModifier: jsii.String("segmentModifier"),
//   							},
//   							mediaPackageOutputSettings: &mediaPackageOutputSettingsProperty{
//   							},
//   							msSmoothOutputSettings: &msSmoothOutputSettingsProperty{
//   								h265PackagingType: jsii.String("h265PackagingType"),
//   								nameModifier: jsii.String("nameModifier"),
//   							},
//   							multiplexOutputSettings: &multiplexOutputSettingsProperty{
//   								destination: &outputLocationRefProperty{
//   									destinationRefId: jsii.String("destinationRefId"),
//   								},
//   							},
//   							rtmpOutputSettings: &rtmpOutputSettingsProperty{
//   								certificateMode: jsii.String("certificateMode"),
//   								connectionRetryInterval: jsii.Number(123),
//   								destination: &outputLocationRefProperty{
//   									destinationRefId: jsii.String("destinationRefId"),
//   								},
//   								numRetries: jsii.Number(123),
//   							},
//   							udpOutputSettings: &udpOutputSettingsProperty{
//   								bufferMsec: jsii.Number(123),
//   								containerSettings: &udpContainerSettingsProperty{
//   									m2TsSettings: &m2tsSettingsProperty{
//   										absentInputAudioBehavior: jsii.String("absentInputAudioBehavior"),
//   										arib: jsii.String("arib"),
//   										aribCaptionsPid: jsii.String("aribCaptionsPid"),
//   										aribCaptionsPidControl: jsii.String("aribCaptionsPidControl"),
//   										audioBufferModel: jsii.String("audioBufferModel"),
//   										audioFramesPerPes: jsii.Number(123),
//   										audioPids: jsii.String("audioPids"),
//   										audioStreamType: jsii.String("audioStreamType"),
//   										bitrate: jsii.Number(123),
//   										bufferModel: jsii.String("bufferModel"),
//   										ccDescriptor: jsii.String("ccDescriptor"),
//   										dvbNitSettings: &dvbNitSettingsProperty{
//   											networkId: jsii.Number(123),
//   											networkName: jsii.String("networkName"),
//   											repInterval: jsii.Number(123),
//   										},
//   										dvbSdtSettings: &dvbSdtSettingsProperty{
//   											outputSdt: jsii.String("outputSdt"),
//   											repInterval: jsii.Number(123),
//   											serviceName: jsii.String("serviceName"),
//   											serviceProviderName: jsii.String("serviceProviderName"),
//   										},
//   										dvbSubPids: jsii.String("dvbSubPids"),
//   										dvbTdtSettings: &dvbTdtSettingsProperty{
//   											repInterval: jsii.Number(123),
//   										},
//   										dvbTeletextPid: jsii.String("dvbTeletextPid"),
//   										ebif: jsii.String("ebif"),
//   										ebpAudioInterval: jsii.String("ebpAudioInterval"),
//   										ebpLookaheadMs: jsii.Number(123),
//   										ebpPlacement: jsii.String("ebpPlacement"),
//   										ecmPid: jsii.String("ecmPid"),
//   										esRateInPes: jsii.String("esRateInPes"),
//   										etvPlatformPid: jsii.String("etvPlatformPid"),
//   										etvSignalPid: jsii.String("etvSignalPid"),
//   										fragmentTime: jsii.Number(123),
//   										klv: jsii.String("klv"),
//   										klvDataPids: jsii.String("klvDataPids"),
//   										nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   										nullPacketBitrate: jsii.Number(123),
//   										patInterval: jsii.Number(123),
//   										pcrControl: jsii.String("pcrControl"),
//   										pcrPeriod: jsii.Number(123),
//   										pcrPid: jsii.String("pcrPid"),
//   										pmtInterval: jsii.Number(123),
//   										pmtPid: jsii.String("pmtPid"),
//   										programNum: jsii.Number(123),
//   										rateMode: jsii.String("rateMode"),
//   										scte27Pids: jsii.String("scte27Pids"),
//   										scte35Control: jsii.String("scte35Control"),
//   										scte35Pid: jsii.String("scte35Pid"),
//   										segmentationMarkers: jsii.String("segmentationMarkers"),
//   										segmentationStyle: jsii.String("segmentationStyle"),
//   										segmentationTime: jsii.Number(123),
//   										timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   										timedMetadataPid: jsii.String("timedMetadataPid"),
//   										transportStreamId: jsii.Number(123),
//   										videoPid: jsii.String("videoPid"),
//   									},
//   								},
//   								destination: &outputLocationRefProperty{
//   									destinationRefId: jsii.String("destinationRefId"),
//   								},
//   								fecOutputSettings: &fecOutputSettingsProperty{
//   									columnDepth: jsii.Number(123),
//   									includeFec: jsii.String("includeFec"),
//   									rowLength: jsii.Number(123),
//   								},
//   							},
//   						},
//   						videoDescriptionName: jsii.String("videoDescriptionName"),
//   					},
//   				},
//   			},
//   		},
//   		timecodeConfig: &timecodeConfigProperty{
//   			source: jsii.String("source"),
//   			syncThreshold: jsii.Number(123),
//   		},
//   		videoDescriptions: []interface{}{
//   			&videoDescriptionProperty{
//   				codecSettings: &videoCodecSettingsProperty{
//   					frameCaptureSettings: &frameCaptureSettingsProperty{
//   						captureInterval: jsii.Number(123),
//   						captureIntervalUnits: jsii.String("captureIntervalUnits"),
//   					},
//   					h264Settings: &h264SettingsProperty{
//   						adaptiveQuantization: jsii.String("adaptiveQuantization"),
//   						afdSignaling: jsii.String("afdSignaling"),
//   						bitrate: jsii.Number(123),
//   						bufFillPct: jsii.Number(123),
//   						bufSize: jsii.Number(123),
//   						colorMetadata: jsii.String("colorMetadata"),
//   						colorSpaceSettings: &h264ColorSpaceSettingsProperty{
//   							colorSpacePassthroughSettings: &colorSpacePassthroughSettingsProperty{
//   							},
//   							rec601Settings: &rec601SettingsProperty{
//   							},
//   							rec709Settings: &rec709SettingsProperty{
//   							},
//   						},
//   						entropyEncoding: jsii.String("entropyEncoding"),
//   						filterSettings: &h264FilterSettingsProperty{
//   							temporalFilterSettings: &temporalFilterSettingsProperty{
//   								postFilterSharpening: jsii.String("postFilterSharpening"),
//   								strength: jsii.String("strength"),
//   							},
//   						},
//   						fixedAfd: jsii.String("fixedAfd"),
//   						flickerAq: jsii.String("flickerAq"),
//   						forceFieldPictures: jsii.String("forceFieldPictures"),
//   						framerateControl: jsii.String("framerateControl"),
//   						framerateDenominator: jsii.Number(123),
//   						framerateNumerator: jsii.Number(123),
//   						gopBReference: jsii.String("gopBReference"),
//   						gopClosedCadence: jsii.Number(123),
//   						gopNumBFrames: jsii.Number(123),
//   						gopSize: jsii.Number(123),
//   						gopSizeUnits: jsii.String("gopSizeUnits"),
//   						level: jsii.String("level"),
//   						lookAheadRateControl: jsii.String("lookAheadRateControl"),
//   						maxBitrate: jsii.Number(123),
//   						minIInterval: jsii.Number(123),
//   						numRefFrames: jsii.Number(123),
//   						parControl: jsii.String("parControl"),
//   						parDenominator: jsii.Number(123),
//   						parNumerator: jsii.Number(123),
//   						profile: jsii.String("profile"),
//   						qualityLevel: jsii.String("qualityLevel"),
//   						qvbrQualityLevel: jsii.Number(123),
//   						rateControlMode: jsii.String("rateControlMode"),
//   						scanType: jsii.String("scanType"),
//   						sceneChangeDetect: jsii.String("sceneChangeDetect"),
//   						slices: jsii.Number(123),
//   						softness: jsii.Number(123),
//   						spatialAq: jsii.String("spatialAq"),
//   						subgopLength: jsii.String("subgopLength"),
//   						syntax: jsii.String("syntax"),
//   						temporalAq: jsii.String("temporalAq"),
//   						timecodeInsertion: jsii.String("timecodeInsertion"),
//   					},
//   					h265Settings: &h265SettingsProperty{
//   						adaptiveQuantization: jsii.String("adaptiveQuantization"),
//   						afdSignaling: jsii.String("afdSignaling"),
//   						alternativeTransferFunction: jsii.String("alternativeTransferFunction"),
//   						bitrate: jsii.Number(123),
//   						bufSize: jsii.Number(123),
//   						colorMetadata: jsii.String("colorMetadata"),
//   						colorSpaceSettings: &h265ColorSpaceSettingsProperty{
//   							colorSpacePassthroughSettings: &colorSpacePassthroughSettingsProperty{
//   							},
//   							hdr10Settings: &hdr10SettingsProperty{
//   								maxCll: jsii.Number(123),
//   								maxFall: jsii.Number(123),
//   							},
//   							rec601Settings: &rec601SettingsProperty{
//   							},
//   							rec709Settings: &rec709SettingsProperty{
//   							},
//   						},
//   						filterSettings: &h265FilterSettingsProperty{
//   							temporalFilterSettings: &temporalFilterSettingsProperty{
//   								postFilterSharpening: jsii.String("postFilterSharpening"),
//   								strength: jsii.String("strength"),
//   							},
//   						},
//   						fixedAfd: jsii.String("fixedAfd"),
//   						flickerAq: jsii.String("flickerAq"),
//   						framerateDenominator: jsii.Number(123),
//   						framerateNumerator: jsii.Number(123),
//   						gopClosedCadence: jsii.Number(123),
//   						gopSize: jsii.Number(123),
//   						gopSizeUnits: jsii.String("gopSizeUnits"),
//   						level: jsii.String("level"),
//   						lookAheadRateControl: jsii.String("lookAheadRateControl"),
//   						maxBitrate: jsii.Number(123),
//   						minIInterval: jsii.Number(123),
//   						parDenominator: jsii.Number(123),
//   						parNumerator: jsii.Number(123),
//   						profile: jsii.String("profile"),
//   						qvbrQualityLevel: jsii.Number(123),
//   						rateControlMode: jsii.String("rateControlMode"),
//   						scanType: jsii.String("scanType"),
//   						sceneChangeDetect: jsii.String("sceneChangeDetect"),
//   						slices: jsii.Number(123),
//   						tier: jsii.String("tier"),
//   						timecodeInsertion: jsii.String("timecodeInsertion"),
//   					},
//   					mpeg2Settings: &mpeg2SettingsProperty{
//   						adaptiveQuantization: jsii.String("adaptiveQuantization"),
//   						afdSignaling: jsii.String("afdSignaling"),
//   						colorMetadata: jsii.String("colorMetadata"),
//   						colorSpace: jsii.String("colorSpace"),
//   						displayAspectRatio: jsii.String("displayAspectRatio"),
//   						filterSettings: &mpeg2FilterSettingsProperty{
//   							temporalFilterSettings: &temporalFilterSettingsProperty{
//   								postFilterSharpening: jsii.String("postFilterSharpening"),
//   								strength: jsii.String("strength"),
//   							},
//   						},
//   						fixedAfd: jsii.String("fixedAfd"),
//   						framerateDenominator: jsii.Number(123),
//   						framerateNumerator: jsii.Number(123),
//   						gopClosedCadence: jsii.Number(123),
//   						gopNumBFrames: jsii.Number(123),
//   						gopSize: jsii.Number(123),
//   						gopSizeUnits: jsii.String("gopSizeUnits"),
//   						scanType: jsii.String("scanType"),
//   						subgopLength: jsii.String("subgopLength"),
//   						timecodeInsertion: jsii.String("timecodeInsertion"),
//   					},
//   				},
//   				height: jsii.Number(123),
//   				name: jsii.String("name"),
//   				respondToAfd: jsii.String("respondToAfd"),
//   				scalingBehavior: jsii.String("scalingBehavior"),
//   				sharpness: jsii.Number(123),
//   				width: jsii.Number(123),
//   			},
//   		},
//   	},
//   	inputAttachments: []interface{}{
//   		&inputAttachmentProperty{
//   			automaticInputFailoverSettings: &automaticInputFailoverSettingsProperty{
//   				errorClearTimeMsec: jsii.Number(123),
//   				failoverConditions: []interface{}{
//   					&failoverConditionProperty{
//   						failoverConditionSettings: &failoverConditionSettingsProperty{
//   							audioSilenceSettings: &audioSilenceFailoverSettingsProperty{
//   								audioSelectorName: jsii.String("audioSelectorName"),
//   								audioSilenceThresholdMsec: jsii.Number(123),
//   							},
//   							inputLossSettings: &inputLossFailoverSettingsProperty{
//   								inputLossThresholdMsec: jsii.Number(123),
//   							},
//   							videoBlackSettings: &videoBlackFailoverSettingsProperty{
//   								blackDetectThreshold: jsii.Number(123),
//   								videoBlackThresholdMsec: jsii.Number(123),
//   							},
//   						},
//   					},
//   				},
//   				inputPreference: jsii.String("inputPreference"),
//   				secondaryInputId: jsii.String("secondaryInputId"),
//   			},
//   			inputAttachmentName: jsii.String("inputAttachmentName"),
//   			inputId: jsii.String("inputId"),
//   			inputSettings: &inputSettingsProperty{
//   				audioSelectors: []interface{}{
//   					&audioSelectorProperty{
//   						name: jsii.String("name"),
//   						selectorSettings: &audioSelectorSettingsProperty{
//   							audioHlsRenditionSelection: &audioHlsRenditionSelectionProperty{
//   								groupId: jsii.String("groupId"),
//   								name: jsii.String("name"),
//   							},
//   							audioLanguageSelection: &audioLanguageSelectionProperty{
//   								languageCode: jsii.String("languageCode"),
//   								languageSelectionPolicy: jsii.String("languageSelectionPolicy"),
//   							},
//   							audioPidSelection: &audioPidSelectionProperty{
//   								pid: jsii.Number(123),
//   							},
//   							audioTrackSelection: &audioTrackSelectionProperty{
//   								tracks: []interface{}{
//   									&audioTrackProperty{
//   										track: jsii.Number(123),
//   									},
//   								},
//   							},
//   						},
//   					},
//   				},
//   				captionSelectors: []interface{}{
//   					&captionSelectorProperty{
//   						languageCode: jsii.String("languageCode"),
//   						name: jsii.String("name"),
//   						selectorSettings: &captionSelectorSettingsProperty{
//   							ancillarySourceSettings: &ancillarySourceSettingsProperty{
//   								sourceAncillaryChannelNumber: jsii.Number(123),
//   							},
//   							aribSourceSettings: &aribSourceSettingsProperty{
//   							},
//   							dvbSubSourceSettings: &dvbSubSourceSettingsProperty{
//   								ocrLanguage: jsii.String("ocrLanguage"),
//   								pid: jsii.Number(123),
//   							},
//   							embeddedSourceSettings: &embeddedSourceSettingsProperty{
//   								convert608To708: jsii.String("convert608To708"),
//   								scte20Detection: jsii.String("scte20Detection"),
//   								source608ChannelNumber: jsii.Number(123),
//   								source608TrackNumber: jsii.Number(123),
//   							},
//   							scte20SourceSettings: &scte20SourceSettingsProperty{
//   								convert608To708: jsii.String("convert608To708"),
//   								source608ChannelNumber: jsii.Number(123),
//   							},
//   							scte27SourceSettings: &scte27SourceSettingsProperty{
//   								ocrLanguage: jsii.String("ocrLanguage"),
//   								pid: jsii.Number(123),
//   							},
//   							teletextSourceSettings: &teletextSourceSettingsProperty{
//   								outputRectangle: &captionRectangleProperty{
//   									height: jsii.Number(123),
//   									leftOffset: jsii.Number(123),
//   									topOffset: jsii.Number(123),
//   									width: jsii.Number(123),
//   								},
//   								pageNumber: jsii.String("pageNumber"),
//   							},
//   						},
//   					},
//   				},
//   				deblockFilter: jsii.String("deblockFilter"),
//   				denoiseFilter: jsii.String("denoiseFilter"),
//   				filterStrength: jsii.Number(123),
//   				inputFilter: jsii.String("inputFilter"),
//   				networkInputSettings: &networkInputSettingsProperty{
//   					hlsInputSettings: &hlsInputSettingsProperty{
//   						bandwidth: jsii.Number(123),
//   						bufferSegments: jsii.Number(123),
//   						retries: jsii.Number(123),
//   						retryInterval: jsii.Number(123),
//   						scte35Source: jsii.String("scte35Source"),
//   					},
//   					serverValidation: jsii.String("serverValidation"),
//   				},
//   				scte35Pid: jsii.Number(123),
//   				smpte2038DataPreference: jsii.String("smpte2038DataPreference"),
//   				sourceEndBehavior: jsii.String("sourceEndBehavior"),
//   				videoSelector: &videoSelectorProperty{
//   					colorSpace: jsii.String("colorSpace"),
//   					colorSpaceSettings: &videoSelectorColorSpaceSettingsProperty{
//   						hdr10Settings: &hdr10SettingsProperty{
//   							maxCll: jsii.Number(123),
//   							maxFall: jsii.Number(123),
//   						},
//   					},
//   					colorSpaceUsage: jsii.String("colorSpaceUsage"),
//   					selectorSettings: &videoSelectorSettingsProperty{
//   						videoSelectorPid: &videoSelectorPidProperty{
//   							pid: jsii.Number(123),
//   						},
//   						videoSelectorProgramId: &videoSelectorProgramIdProperty{
//   							programId: jsii.Number(123),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	inputSpecification: &inputSpecificationProperty{
//   		codec: jsii.String("codec"),
//   		maximumBitrate: jsii.String("maximumBitrate"),
//   		resolution: jsii.String("resolution"),
//   	},
//   	logLevel: jsii.String("logLevel"),
//   	name: jsii.String("name"),
//   	roleArn: jsii.String("roleArn"),
//   	tags: tags,
//   	vpc: &vpcOutputSettingsProperty{
//   		publicAddressAllocationIds: []*string{
//   			jsii.String("publicAddressAllocationIds"),
//   		},
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   })
//
type CfnChannel interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ARN of the MediaLive channel.
	//
	// For example: arn:aws:medialive:us-west-1:111122223333:medialive:channel:1234567.
	AttrArn() *string
	// The inputs that are attached to this channel.
	//
	// The inputs are identified by their IDs (not by their names or their ARNs).
	AttrInputs() *[]*string
	// Specification of CDI inputs for this channel.
	CdiInputSpecification() interface{}
	SetCdiInputSpecification(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The class for this channel.
	//
	// For a channel with two pipelines, the class is STANDARD. For a channel with one pipeline, the class is SINGLE_PIPELINE.
	ChannelClass() *string
	SetChannelClass(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The settings that identify the destination for the outputs in this MediaLive output package.
	Destinations() interface{}
	SetDestinations(val interface{})
	// The encoding configuration for the output content.
	EncoderSettings() interface{}
	SetEncoderSettings(val interface{})
	// The list of input attachments for the channel.
	InputAttachments() interface{}
	SetInputAttachments(val interface{})
	// The input specification for this channel.
	//
	// It specifies the key characteristics of the inputs for this channel: the maximum bitrate, the resolution, and the codec.
	InputSpecification() interface{}
	SetInputSpecification(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The verbosity for logging activity for this channel.
	//
	// Charges for logging (which are generated through Amazon CloudWatch Logging) are higher for higher verbosities.
	LogLevel() *string
	SetLogLevel(val *string)
	// A name for this audio selector.
	//
	// The AudioDescription (in an output) references this name in order to identify a specific input audio to include in that output.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The IAM role for MediaLive to assume when running this channel.
	//
	// The role is identified by its ARN.
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// A collection of tags for this channel.
	//
	// Each tag is a key-value pair.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Settings to enable VPC mode in the channel, so that the endpoints for all outputs are in your VPC.
	Vpc() interface{}
	SetVpc(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnChannel
type jsiiProxy_CfnChannel struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnChannel) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) AttrInputs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrInputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) CdiInputSpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cdiInputSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) ChannelClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Destinations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) EncoderSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encoderSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) InputAttachments() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputAttachments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) InputSpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Vpc() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


// Create a new `AWS::MediaLive::Channel`.
func NewCfnChannel(scope constructs.Construct, id *string, props *CfnChannelProps) CfnChannel {
	_init_.Initialize()

	j := jsiiProxy_CfnChannel{}

	_jsii_.Create(
		"aws-cdk-lib.aws_medialive.CfnChannel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MediaLive::Channel`.
func NewCfnChannel_Override(c CfnChannel, scope constructs.Construct, id *string, props *CfnChannelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_medialive.CfnChannel",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnChannel) SetCdiInputSpecification(val interface{}) {
	_jsii_.Set(
		j,
		"cdiInputSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnChannel) SetChannelClass(val *string) {
	_jsii_.Set(
		j,
		"channelClass",
		val,
	)
}

func (j *jsiiProxy_CfnChannel) SetDestinations(val interface{}) {
	_jsii_.Set(
		j,
		"destinations",
		val,
	)
}

func (j *jsiiProxy_CfnChannel) SetEncoderSettings(val interface{}) {
	_jsii_.Set(
		j,
		"encoderSettings",
		val,
	)
}

func (j *jsiiProxy_CfnChannel) SetInputAttachments(val interface{}) {
	_jsii_.Set(
		j,
		"inputAttachments",
		val,
	)
}

func (j *jsiiProxy_CfnChannel) SetInputSpecification(val interface{}) {
	_jsii_.Set(
		j,
		"inputSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnChannel) SetLogLevel(val *string) {
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_CfnChannel) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnChannel) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnChannel) SetVpc(val interface{}) {
	_jsii_.Set(
		j,
		"vpc",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnChannel_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_medialive.CfnChannel",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnChannel_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_medialive.CfnChannel",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_medialive.CfnChannel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnChannel_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_medialive.CfnChannel",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnChannel) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnChannel) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnChannel) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnChannel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnChannel) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnChannel) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnChannel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnChannel) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnChannel) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnChannel) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnChannel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnChannel) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnChannel) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnChannel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnChannel) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The settings for an AAC audio encode in the output.
//
// The parent of this entity is AudioCodecSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aacSettingsProperty := &aacSettingsProperty{
//   	bitrate: jsii.Number(123),
//   	codingMode: jsii.String("codingMode"),
//   	inputType: jsii.String("inputType"),
//   	profile: jsii.String("profile"),
//   	rateControlMode: jsii.String("rateControlMode"),
//   	rawFormat: jsii.String("rawFormat"),
//   	sampleRate: jsii.Number(123),
//   	spec: jsii.String("spec"),
//   	vbrQuality: jsii.String("vbrQuality"),
//   }
//
type CfnChannel_AacSettingsProperty struct {
	// The average bitrate in bits/second.
	//
	// Valid values depend on the rate control mode and profile.
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// Mono, stereo, or 5.1 channel layout. Valid values depend on the rate control mode and profile. The adReceiverMix setting receives a stereo description plus control track, and emits a mono AAC encode of the description track, with control data emitted in the PES header as per ETSI TS 101 154 Annex E.
	CodingMode *string `field:"optional" json:"codingMode" yaml:"codingMode"`
	// Set to broadcasterMixedAd when the input contains pre-mixed main audio + AD (narration) as a stereo pair.
	//
	// The Audio Type field (audioType) will be set to 3, which signals to downstream systems that this stream contains broadcaster mixed AD. Note that the input received by the encoder must contain pre-mixed audio; MediaLive does not perform the mixing. The values in audioTypeControl and audioType (in AudioDescription) are ignored when set to broadcasterMixedAd. Leave this set to normal when the input does not contain pre-mixed audio + AD.
	InputType *string `field:"optional" json:"inputType" yaml:"inputType"`
	// The AAC profile.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// The rate control mode.
	RateControlMode *string `field:"optional" json:"rateControlMode" yaml:"rateControlMode"`
	// Sets the LATM/LOAS AAC output for raw containers.
	RawFormat *string `field:"optional" json:"rawFormat" yaml:"rawFormat"`
	// The sample rate in Hz.
	//
	// Valid values depend on the rate control mode and profile.
	SampleRate *float64 `field:"optional" json:"sampleRate" yaml:"sampleRate"`
	// Uses MPEG-2 AAC audio instead of MPEG-4 AAC audio for raw or MPEG-2 Transport Stream containers.
	Spec *string `field:"optional" json:"spec" yaml:"spec"`
	// The VBR quality level.
	//
	// This is used only if rateControlMode is VBR.
	VbrQuality *string `field:"optional" json:"vbrQuality" yaml:"vbrQuality"`
}

// The settings for an AC3 audio encode in the output.
//
// The parent of this entity is AudioCodecSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ac3SettingsProperty := &ac3SettingsProperty{
//   	bitrate: jsii.Number(123),
//   	bitstreamMode: jsii.String("bitstreamMode"),
//   	codingMode: jsii.String("codingMode"),
//   	dialnorm: jsii.Number(123),
//   	drcProfile: jsii.String("drcProfile"),
//   	lfeFilter: jsii.String("lfeFilter"),
//   	metadataControl: jsii.String("metadataControl"),
//   }
//
type CfnChannel_Ac3SettingsProperty struct {
	// The average bitrate in bits/second.
	//
	// Valid bitrates depend on the coding mode.
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// Specifies the bitstream mode (bsmod) for the emitted AC-3 stream.
	//
	// For more information about these values, see ATSC A/52-2012.
	BitstreamMode *string `field:"optional" json:"bitstreamMode" yaml:"bitstreamMode"`
	// The Dolby Digital coding mode.
	//
	// This determines the number of channels.
	CodingMode *string `field:"optional" json:"codingMode" yaml:"codingMode"`
	// Sets the dialnorm for the output.
	//
	// If excluded and the input audio is Dolby Digital, dialnorm is passed through.
	Dialnorm *float64 `field:"optional" json:"dialnorm" yaml:"dialnorm"`
	// If set to filmStandard, adds dynamic range compression signaling to the output bitstream as defined in the Dolby Digital specification.
	DrcProfile *string `field:"optional" json:"drcProfile" yaml:"drcProfile"`
	// When set to enabled, applies a 120Hz lowpass filter to the LFE channel prior to encoding.
	//
	// This is valid only in codingMode32Lfe mode.
	LfeFilter *string `field:"optional" json:"lfeFilter" yaml:"lfeFilter"`
	// When set to followInput, encoder metadata is sourced from the DD, DD+, or DolbyE decoder that supplies this audio data.
	//
	// If the audio is supplied from one of these streams, the static metadata settings are used.
	MetadataControl *string `field:"optional" json:"metadataControl" yaml:"metadataControl"`
}

// Information about the ancillary captions to extract from the input.
//
// The parent of this entity is CaptionSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ancillarySourceSettingsProperty := &ancillarySourceSettingsProperty{
//   	sourceAncillaryChannelNumber: jsii.Number(123),
//   }
//
type CfnChannel_AncillarySourceSettingsProperty struct {
	// Specifies the number (1 to 4) of the captions channel you want to extract from the ancillary captions.
	//
	// If you plan to convert the ancillary captions to another format, complete this field. If you plan to choose Embedded as the captions destination in the output (to pass through all the channels in the ancillary captions), leave this field blank because MediaLive ignores the field.
	SourceAncillaryChannelNumber *float64 `field:"optional" json:"sourceAncillaryChannelNumber" yaml:"sourceAncillaryChannelNumber"`
}

// Settings to configure the destination of an Archive output.
//
// The parent of this entity is ArchiveGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   archiveCdnSettingsProperty := &archiveCdnSettingsProperty{
//   	archiveS3Settings: &archiveS3SettingsProperty{
//   		cannedAcl: jsii.String("cannedAcl"),
//   	},
//   }
//
type CfnChannel_ArchiveCdnSettingsProperty struct {
	// Sets up Amazon S3 as the destination for this Archive output.
	ArchiveS3Settings interface{} `field:"optional" json:"archiveS3Settings" yaml:"archiveS3Settings"`
}

// The archive container settings.
//
// The parent of this entity is ArchiveOutputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   archiveContainerSettingsProperty := &archiveContainerSettingsProperty{
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
//   	rawSettings: &rawSettingsProperty{
//   	},
//   }
//
type CfnChannel_ArchiveContainerSettingsProperty struct {
	// The settings for the M2TS in the archive output.
	M2TsSettings interface{} `field:"optional" json:"m2TsSettings" yaml:"m2TsSettings"`
	// The settings for Raw archive output type.
	RawSettings interface{} `field:"optional" json:"rawSettings" yaml:"rawSettings"`
}

// The settings for an archive output group.
//
// The parent of this entity is OutputGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   archiveGroupSettingsProperty := &archiveGroupSettingsProperty{
//   	archiveCdnSettings: &archiveCdnSettingsProperty{
//   		archiveS3Settings: &archiveS3SettingsProperty{
//   			cannedAcl: jsii.String("cannedAcl"),
//   		},
//   	},
//   	destination: &outputLocationRefProperty{
//   		destinationRefId: jsii.String("destinationRefId"),
//   	},
//   	rolloverInterval: jsii.Number(123),
//   }
//
type CfnChannel_ArchiveGroupSettingsProperty struct {
	// Settings to configure the destination of an Archive output.
	ArchiveCdnSettings interface{} `field:"optional" json:"archiveCdnSettings" yaml:"archiveCdnSettings"`
	// A directory and base file name where archive files should be written.
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// The number of seconds to write to an archive file before closing and starting a new one.
	RolloverInterval *float64 `field:"optional" json:"rolloverInterval" yaml:"rolloverInterval"`
}

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

// Sets up Amazon S3 as the destination for this Archive output.
//
// The parent of this entity is ArchiveCdnSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   archiveS3SettingsProperty := &archiveS3SettingsProperty{
//   	cannedAcl: jsii.String("cannedAcl"),
//   }
//
type CfnChannel_ArchiveS3SettingsProperty struct {
	// Specify the canned ACL to apply to each S3 request.
	//
	// Defaults to none.
	CannedAcl *string `field:"optional" json:"cannedAcl" yaml:"cannedAcl"`
}

// The configuration of ARIB captions in the output.
//
// The parent of this entity is CaptionDestinationSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aribDestinationSettingsProperty := &aribDestinationSettingsProperty{
//   }
//
type CfnChannel_AribDestinationSettingsProperty struct {
}

// Information about the ARIB captions to extract from the input.
//
// The parent of this entity is CaptionSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aribSourceSettingsProperty := &aribSourceSettingsProperty{
//   }
//
type CfnChannel_AribSourceSettingsProperty struct {
}

// The settings for remixing audio.
//
// The parent of this entity is RemixSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioChannelMappingProperty := &audioChannelMappingProperty{
//   	inputChannelLevels: []interface{}{
//   		&inputChannelLevelProperty{
//   			gain: jsii.Number(123),
//   			inputChannel: jsii.Number(123),
//   		},
//   	},
//   	outputChannel: jsii.Number(123),
//   }
//
type CfnChannel_AudioChannelMappingProperty struct {
	// The indices and gain values for each input channel that should be remixed into this output channel.
	InputChannelLevels interface{} `field:"optional" json:"inputChannelLevels" yaml:"inputChannelLevels"`
	// The index of the output channel that is being produced.
	OutputChannel *float64 `field:"optional" json:"outputChannel" yaml:"outputChannel"`
}

// The configuration of the audio codec in the audio output.
//
// The parent of this entity is AudioDescription.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioCodecSettingsProperty := &audioCodecSettingsProperty{
//   	aacSettings: &aacSettingsProperty{
//   		bitrate: jsii.Number(123),
//   		codingMode: jsii.String("codingMode"),
//   		inputType: jsii.String("inputType"),
//   		profile: jsii.String("profile"),
//   		rateControlMode: jsii.String("rateControlMode"),
//   		rawFormat: jsii.String("rawFormat"),
//   		sampleRate: jsii.Number(123),
//   		spec: jsii.String("spec"),
//   		vbrQuality: jsii.String("vbrQuality"),
//   	},
//   	ac3Settings: &ac3SettingsProperty{
//   		bitrate: jsii.Number(123),
//   		bitstreamMode: jsii.String("bitstreamMode"),
//   		codingMode: jsii.String("codingMode"),
//   		dialnorm: jsii.Number(123),
//   		drcProfile: jsii.String("drcProfile"),
//   		lfeFilter: jsii.String("lfeFilter"),
//   		metadataControl: jsii.String("metadataControl"),
//   	},
//   	eac3Settings: &eac3SettingsProperty{
//   		attenuationControl: jsii.String("attenuationControl"),
//   		bitrate: jsii.Number(123),
//   		bitstreamMode: jsii.String("bitstreamMode"),
//   		codingMode: jsii.String("codingMode"),
//   		dcFilter: jsii.String("dcFilter"),
//   		dialnorm: jsii.Number(123),
//   		drcLine: jsii.String("drcLine"),
//   		drcRf: jsii.String("drcRf"),
//   		lfeControl: jsii.String("lfeControl"),
//   		lfeFilter: jsii.String("lfeFilter"),
//   		loRoCenterMixLevel: jsii.Number(123),
//   		loRoSurroundMixLevel: jsii.Number(123),
//   		ltRtCenterMixLevel: jsii.Number(123),
//   		ltRtSurroundMixLevel: jsii.Number(123),
//   		metadataControl: jsii.String("metadataControl"),
//   		passthroughControl: jsii.String("passthroughControl"),
//   		phaseControl: jsii.String("phaseControl"),
//   		stereoDownmix: jsii.String("stereoDownmix"),
//   		surroundExMode: jsii.String("surroundExMode"),
//   		surroundMode: jsii.String("surroundMode"),
//   	},
//   	mp2Settings: &mp2SettingsProperty{
//   		bitrate: jsii.Number(123),
//   		codingMode: jsii.String("codingMode"),
//   		sampleRate: jsii.Number(123),
//   	},
//   	passThroughSettings: &passThroughSettingsProperty{
//   	},
//   	wavSettings: &wavSettingsProperty{
//   		bitDepth: jsii.Number(123),
//   		codingMode: jsii.String("codingMode"),
//   		sampleRate: jsii.Number(123),
//   	},
//   }
//
type CfnChannel_AudioCodecSettingsProperty struct {
	// The setup of the AAC audio codec in the output.
	AacSettings interface{} `field:"optional" json:"aacSettings" yaml:"aacSettings"`
	// The setup of an AC3 audio codec in the output.
	Ac3Settings interface{} `field:"optional" json:"ac3Settings" yaml:"ac3Settings"`
	// The setup of an EAC3 audio codec in the output.
	Eac3Settings interface{} `field:"optional" json:"eac3Settings" yaml:"eac3Settings"`
	// The setup of an MP2 audio codec in the output.
	Mp2Settings interface{} `field:"optional" json:"mp2Settings" yaml:"mp2Settings"`
	// The setup to pass through the Dolby audio codec to the output.
	PassThroughSettings interface{} `field:"optional" json:"passThroughSettings" yaml:"passThroughSettings"`
	// Settings for audio encoded with the WAV codec.
	WavSettings interface{} `field:"optional" json:"wavSettings" yaml:"wavSettings"`
}

// The encoding information for one output audio.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioDescriptionProperty := &audioDescriptionProperty{
//   	audioNormalizationSettings: &audioNormalizationSettingsProperty{
//   		algorithm: jsii.String("algorithm"),
//   		algorithmControl: jsii.String("algorithmControl"),
//   		targetLkfs: jsii.Number(123),
//   	},
//   	audioSelectorName: jsii.String("audioSelectorName"),
//   	audioType: jsii.String("audioType"),
//   	audioTypeControl: jsii.String("audioTypeControl"),
//   	audioWatermarkingSettings: &audioWatermarkSettingsProperty{
//   		nielsenWatermarksSettings: &nielsenWatermarksSettingsProperty{
//   			nielsenCbetSettings: &nielsenCBETProperty{
//   				cbetCheckDigitString: jsii.String("cbetCheckDigitString"),
//   				cbetStepaside: jsii.String("cbetStepaside"),
//   				csid: jsii.String("csid"),
//   			},
//   			nielsenDistributionType: jsii.String("nielsenDistributionType"),
//   			nielsenNaesIiNwSettings: &nielsenNaesIiNwProperty{
//   				checkDigitString: jsii.String("checkDigitString"),
//   				sid: jsii.Number(123),
//   			},
//   		},
//   	},
//   	codecSettings: &audioCodecSettingsProperty{
//   		aacSettings: &aacSettingsProperty{
//   			bitrate: jsii.Number(123),
//   			codingMode: jsii.String("codingMode"),
//   			inputType: jsii.String("inputType"),
//   			profile: jsii.String("profile"),
//   			rateControlMode: jsii.String("rateControlMode"),
//   			rawFormat: jsii.String("rawFormat"),
//   			sampleRate: jsii.Number(123),
//   			spec: jsii.String("spec"),
//   			vbrQuality: jsii.String("vbrQuality"),
//   		},
//   		ac3Settings: &ac3SettingsProperty{
//   			bitrate: jsii.Number(123),
//   			bitstreamMode: jsii.String("bitstreamMode"),
//   			codingMode: jsii.String("codingMode"),
//   			dialnorm: jsii.Number(123),
//   			drcProfile: jsii.String("drcProfile"),
//   			lfeFilter: jsii.String("lfeFilter"),
//   			metadataControl: jsii.String("metadataControl"),
//   		},
//   		eac3Settings: &eac3SettingsProperty{
//   			attenuationControl: jsii.String("attenuationControl"),
//   			bitrate: jsii.Number(123),
//   			bitstreamMode: jsii.String("bitstreamMode"),
//   			codingMode: jsii.String("codingMode"),
//   			dcFilter: jsii.String("dcFilter"),
//   			dialnorm: jsii.Number(123),
//   			drcLine: jsii.String("drcLine"),
//   			drcRf: jsii.String("drcRf"),
//   			lfeControl: jsii.String("lfeControl"),
//   			lfeFilter: jsii.String("lfeFilter"),
//   			loRoCenterMixLevel: jsii.Number(123),
//   			loRoSurroundMixLevel: jsii.Number(123),
//   			ltRtCenterMixLevel: jsii.Number(123),
//   			ltRtSurroundMixLevel: jsii.Number(123),
//   			metadataControl: jsii.String("metadataControl"),
//   			passthroughControl: jsii.String("passthroughControl"),
//   			phaseControl: jsii.String("phaseControl"),
//   			stereoDownmix: jsii.String("stereoDownmix"),
//   			surroundExMode: jsii.String("surroundExMode"),
//   			surroundMode: jsii.String("surroundMode"),
//   		},
//   		mp2Settings: &mp2SettingsProperty{
//   			bitrate: jsii.Number(123),
//   			codingMode: jsii.String("codingMode"),
//   			sampleRate: jsii.Number(123),
//   		},
//   		passThroughSettings: &passThroughSettingsProperty{
//   		},
//   		wavSettings: &wavSettingsProperty{
//   			bitDepth: jsii.Number(123),
//   			codingMode: jsii.String("codingMode"),
//   			sampleRate: jsii.Number(123),
//   		},
//   	},
//   	languageCode: jsii.String("languageCode"),
//   	languageCodeControl: jsii.String("languageCodeControl"),
//   	name: jsii.String("name"),
//   	remixSettings: &remixSettingsProperty{
//   		channelMappings: []interface{}{
//   			&audioChannelMappingProperty{
//   				inputChannelLevels: []interface{}{
//   					&inputChannelLevelProperty{
//   						gain: jsii.Number(123),
//   						inputChannel: jsii.Number(123),
//   					},
//   				},
//   				outputChannel: jsii.Number(123),
//   			},
//   		},
//   		channelsIn: jsii.Number(123),
//   		channelsOut: jsii.Number(123),
//   	},
//   	streamName: jsii.String("streamName"),
//   }
//
type CfnChannel_AudioDescriptionProperty struct {
	// The advanced audio normalization settings.
	AudioNormalizationSettings interface{} `field:"optional" json:"audioNormalizationSettings" yaml:"audioNormalizationSettings"`
	// The name of the AudioSelector that is used as the source for this AudioDescription.
	AudioSelectorName *string `field:"optional" json:"audioSelectorName" yaml:"audioSelectorName"`
	// Applies only if audioTypeControl is useConfigured.
	//
	// The values for audioType are defined in ISO-IEC 13818-1.
	AudioType *string `field:"optional" json:"audioType" yaml:"audioType"`
	// Determines how audio type is determined.
	//
	// followInput: If the input contains an ISO 639 audioType, then that value is passed through to the output. If the input contains no ISO 639 audioType, the value in Audio Type is included in the output. useConfigured: The value in Audio Type is included in the output. Note that this field and audioType are both ignored if inputType is broadcasterMixedAd.
	AudioTypeControl *string `field:"optional" json:"audioTypeControl" yaml:"audioTypeControl"`
	// Settings to configure one or more solutions that insert audio watermarks in the audio encode.
	AudioWatermarkingSettings interface{} `field:"optional" json:"audioWatermarkingSettings" yaml:"audioWatermarkingSettings"`
	// The audio codec settings.
	CodecSettings interface{} `field:"optional" json:"codecSettings" yaml:"codecSettings"`
	// Indicates the language of the audio output track.
	//
	// Used only if languageControlMode is useConfigured, or there is no ISO 639 language code specified in the input.
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// Choosing followInput causes the ISO 639 language code of the output to follow the ISO 639 language code of the input.
	//
	// The languageCode setting is used when useConfigured is set, or when followInput is selected but there is no ISO 639 language code specified by the input.
	LanguageCodeControl *string `field:"optional" json:"languageCodeControl" yaml:"languageCodeControl"`
	// The name of this AudioDescription.
	//
	// Outputs use this name to uniquely identify this AudioDescription. Description names should be unique within this channel.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The settings that control how input audio channels are remixed into the output audio channels.
	RemixSettings interface{} `field:"optional" json:"remixSettings" yaml:"remixSettings"`
	// Used for Microsoft Smooth and Apple HLS outputs.
	//
	// Indicates the name displayed by the player (for example, English or Director Commentary).
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
}

// Selector for HLS audio rendition.
//
// The parent of this entity is AudioSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioHlsRenditionSelectionProperty := &audioHlsRenditionSelectionProperty{
//   	groupId: jsii.String("groupId"),
//   	name: jsii.String("name"),
//   }
//
type CfnChannel_AudioHlsRenditionSelectionProperty struct {
	// Specifies the GROUP-ID in the #EXT-X-MEDIA tag of the target HLS audio rendition.
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// Specifies the NAME in the #EXT-X-MEDIA tag of the target HLS audio rendition.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

// Information about the audio language to extract.
//
// The parent of this entity is AudioSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioLanguageSelectionProperty := &audioLanguageSelectionProperty{
//   	languageCode: jsii.String("languageCode"),
//   	languageSelectionPolicy: jsii.String("languageSelectionPolicy"),
//   }
//
type CfnChannel_AudioLanguageSelectionProperty struct {
	// Selects a specific three-letter language code from within an audio source.
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// When set to "strict," the transport stream demux strictly identifies audio streams by their language descriptor.
	//
	// If a PMT update occurs such that an audio stream matching the initially selected language is no longer present, then mute is encoded until the language returns. If set to "loose," then on a PMT update the demux chooses another audio stream in the program with the same stream type if it can't find one with the same language.
	LanguageSelectionPolicy *string `field:"optional" json:"languageSelectionPolicy" yaml:"languageSelectionPolicy"`
}

// The settings for normalizing video.
//
// The parent of this entity is AudioDescription.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioNormalizationSettingsProperty := &audioNormalizationSettingsProperty{
//   	algorithm: jsii.String("algorithm"),
//   	algorithmControl: jsii.String("algorithmControl"),
//   	targetLkfs: jsii.Number(123),
//   }
//
type CfnChannel_AudioNormalizationSettingsProperty struct {
	// The audio normalization algorithm to use.
	//
	// itu17701 conforms to the CALM Act specification. itu17702 conforms to the EBU R-128 specification.
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// When set to correctAudio, the output audio is corrected using the chosen algorithm.
	//
	// If set to measureOnly, the audio is measured but not adjusted.
	AlgorithmControl *string `field:"optional" json:"algorithmControl" yaml:"algorithmControl"`
	// The Target LKFS(loudness) to adjust volume to.
	//
	// If no value is entered, a default value is used according to the chosen algorithm. The CALM Act (1770-1) recommends a target of -24 LKFS. The EBU R-128 specification (1770-2) recommends a target of -23 LKFS.
	TargetLkfs *float64 `field:"optional" json:"targetLkfs" yaml:"targetLkfs"`
}

// The configuration of an audio-only HLS output.
//
// The parent of this entity is HlsSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioOnlyHlsSettingsProperty := &audioOnlyHlsSettingsProperty{
//   	audioGroupId: jsii.String("audioGroupId"),
//   	audioOnlyImage: &inputLocationProperty{
//   		passwordParam: jsii.String("passwordParam"),
//   		uri: jsii.String("uri"),
//   		username: jsii.String("username"),
//   	},
//   	audioTrackType: jsii.String("audioTrackType"),
//   	segmentType: jsii.String("segmentType"),
//   }
//
type CfnChannel_AudioOnlyHlsSettingsProperty struct {
	// Specifies the group that the audio rendition belongs to.
	AudioGroupId *string `field:"optional" json:"audioGroupId" yaml:"audioGroupId"`
	// Used with an audio-only stream.
	//
	// It must be a .jpg or .png file. If given, this image is used as the cover art for the audio-only output. Ideally, it should be formatted for an iPhone screen for two reasons. The iPhone does not resize the image; instead, it crops a centered image on the top/bottom and left/right. Additionally, this image file gets saved bit-for-bit into every 10-second segment file, so it increases bandwidth by {image file size} * {segment count} * {user count.}.
	AudioOnlyImage interface{} `field:"optional" json:"audioOnlyImage" yaml:"audioOnlyImage"`
	// Four types of audio-only tracks are supported: Audio-Only Variant Stream The client can play back this audio-only stream instead of video in low-bandwidth scenarios.
	//
	// Represented as an EXT-X-STREAM-INF in the HLS manifest. Alternate Audio, Auto Select, Default Alternate rendition that the client should try to play back by default. Represented as an EXT-X-MEDIA in the HLS manifest with DEFAULT=YES, AUTOSELECT=YES Alternate Audio, Auto Select, Not Default Alternate rendition that the client might try to play back by default. Represented as an EXT-X-MEDIA in the HLS manifest with DEFAULT=NO, AUTOSELECT=YES Alternate Audio, not Auto Select Alternate rendition that the client will not try to play back by default. Represented as an EXT-X-MEDIA in the HLS manifest with DEFAULT=NO, AUTOSELECT=NO.
	AudioTrackType *string `field:"optional" json:"audioTrackType" yaml:"audioTrackType"`
	// Specifies the segment type.
	SegmentType *string `field:"optional" json:"segmentType" yaml:"segmentType"`
}

// Used to extract audio by The PID.
//
// The parent of this entity is AudioSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioPidSelectionProperty := &audioPidSelectionProperty{
//   	pid: jsii.Number(123),
//   }
//
type CfnChannel_AudioPidSelectionProperty struct {
	// Select the audio by this PID.
	Pid *float64 `field:"optional" json:"pid" yaml:"pid"`
}

// Information about one audio to extract from the input.
//
// The parent of this entity is InputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioSelectorProperty := &audioSelectorProperty{
//   	name: jsii.String("name"),
//   	selectorSettings: &audioSelectorSettingsProperty{
//   		audioHlsRenditionSelection: &audioHlsRenditionSelectionProperty{
//   			groupId: jsii.String("groupId"),
//   			name: jsii.String("name"),
//   		},
//   		audioLanguageSelection: &audioLanguageSelectionProperty{
//   			languageCode: jsii.String("languageCode"),
//   			languageSelectionPolicy: jsii.String("languageSelectionPolicy"),
//   		},
//   		audioPidSelection: &audioPidSelectionProperty{
//   			pid: jsii.Number(123),
//   		},
//   		audioTrackSelection: &audioTrackSelectionProperty{
//   			tracks: []interface{}{
//   				&audioTrackProperty{
//   					track: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnChannel_AudioSelectorProperty struct {
	// A name for this AudioSelector.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Information about the specific audio to extract from the input.
	SelectorSettings interface{} `field:"optional" json:"selectorSettings" yaml:"selectorSettings"`
}

// Information about the audio to extract from the input.
//
// The parent of this entity is AudioSelector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioSelectorSettingsProperty := &audioSelectorSettingsProperty{
//   	audioHlsRenditionSelection: &audioHlsRenditionSelectionProperty{
//   		groupId: jsii.String("groupId"),
//   		name: jsii.String("name"),
//   	},
//   	audioLanguageSelection: &audioLanguageSelectionProperty{
//   		languageCode: jsii.String("languageCode"),
//   		languageSelectionPolicy: jsii.String("languageSelectionPolicy"),
//   	},
//   	audioPidSelection: &audioPidSelectionProperty{
//   		pid: jsii.Number(123),
//   	},
//   	audioTrackSelection: &audioTrackSelectionProperty{
//   		tracks: []interface{}{
//   			&audioTrackProperty{
//   				track: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
type CfnChannel_AudioSelectorSettingsProperty struct {
	// Selector for HLS audio rendition.
	AudioHlsRenditionSelection interface{} `field:"optional" json:"audioHlsRenditionSelection" yaml:"audioHlsRenditionSelection"`
	// The language code of the audio to select.
	AudioLanguageSelection interface{} `field:"optional" json:"audioLanguageSelection" yaml:"audioLanguageSelection"`
	// The PID of the audio to select.
	AudioPidSelection interface{} `field:"optional" json:"audioPidSelection" yaml:"audioPidSelection"`
	// Information about the audio track to extract.
	AudioTrackSelection interface{} `field:"optional" json:"audioTrackSelection" yaml:"audioTrackSelection"`
}

// MediaLive will perform a failover if audio is not detected in this input for the specified period.
//
// The parent of this entity is FailoverConditionSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioSilenceFailoverSettingsProperty := &audioSilenceFailoverSettingsProperty{
//   	audioSelectorName: jsii.String("audioSelectorName"),
//   	audioSilenceThresholdMsec: jsii.Number(123),
//   }
//
type CfnChannel_AudioSilenceFailoverSettingsProperty struct {
	// The name of the audio selector in the input that MediaLive should monitor to detect silence.
	//
	// Select your most important rendition. If you didn't create an audio selector in this input, leave blank.
	AudioSelectorName *string `field:"optional" json:"audioSelectorName" yaml:"audioSelectorName"`
	// The amount of time (in milliseconds) that the active input must be silent before automatic input failover occurs.
	//
	// Silence is defined as audio loss or audio quieter than -50 dBFS.
	AudioSilenceThresholdMsec *float64 `field:"optional" json:"audioSilenceThresholdMsec" yaml:"audioSilenceThresholdMsec"`
}

// Information about one audio track to extract. You can select multiple tracks.
//
// The parent of this entity is AudioTrackSelection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioTrackProperty := &audioTrackProperty{
//   	track: jsii.Number(123),
//   }
//
type CfnChannel_AudioTrackProperty struct {
	// 1-based integer value that maps to a specific audio track.
	Track *float64 `field:"optional" json:"track" yaml:"track"`
}

// Information about the audio track to extract.
//
// The parent of this entity is AudioSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioTrackSelectionProperty := &audioTrackSelectionProperty{
//   	tracks: []interface{}{
//   		&audioTrackProperty{
//   			track: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnChannel_AudioTrackSelectionProperty struct {
	// Selects one or more unique audio tracks from within a source.
	Tracks interface{} `field:"optional" json:"tracks" yaml:"tracks"`
}

// Audio Watermark Settings.
//
// The parent of this entity is AudioDescription.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioWatermarkSettingsProperty := &audioWatermarkSettingsProperty{
//   	nielsenWatermarksSettings: &nielsenWatermarksSettingsProperty{
//   		nielsenCbetSettings: &nielsenCBETProperty{
//   			cbetCheckDigitString: jsii.String("cbetCheckDigitString"),
//   			cbetStepaside: jsii.String("cbetStepaside"),
//   			csid: jsii.String("csid"),
//   		},
//   		nielsenDistributionType: jsii.String("nielsenDistributionType"),
//   		nielsenNaesIiNwSettings: &nielsenNaesIiNwProperty{
//   			checkDigitString: jsii.String("checkDigitString"),
//   			sid: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnChannel_AudioWatermarkSettingsProperty struct {
	// Settings to configure Nielsen Watermarks in the audio encode.
	NielsenWatermarksSettings interface{} `field:"optional" json:"nielsenWatermarksSettings" yaml:"nielsenWatermarksSettings"`
}

// Settings to configure the conditions that will define the input as unhealthy and that will make MediaLive fail over to the other input in the input failover pair.
//
// The parent of this entity is InputAttachment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   automaticInputFailoverSettingsProperty := &automaticInputFailoverSettingsProperty{
//   	errorClearTimeMsec: jsii.Number(123),
//   	failoverConditions: []interface{}{
//   		&failoverConditionProperty{
//   			failoverConditionSettings: &failoverConditionSettingsProperty{
//   				audioSilenceSettings: &audioSilenceFailoverSettingsProperty{
//   					audioSelectorName: jsii.String("audioSelectorName"),
//   					audioSilenceThresholdMsec: jsii.Number(123),
//   				},
//   				inputLossSettings: &inputLossFailoverSettingsProperty{
//   					inputLossThresholdMsec: jsii.Number(123),
//   				},
//   				videoBlackSettings: &videoBlackFailoverSettingsProperty{
//   					blackDetectThreshold: jsii.Number(123),
//   					videoBlackThresholdMsec: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	inputPreference: jsii.String("inputPreference"),
//   	secondaryInputId: jsii.String("secondaryInputId"),
//   }
//
type CfnChannel_AutomaticInputFailoverSettingsProperty struct {
	// This clear time defines the requirement a recovered input must meet to be considered healthy.
	//
	// The input must have no failover conditions for this length of time. Enter a time in milliseconds. This value is particularly important if the input_preference for the failover pair is set to PRIMARY_INPUT_PREFERRED, because after this time, MediaLive will switch back to the primary input.
	ErrorClearTimeMsec *float64 `field:"optional" json:"errorClearTimeMsec" yaml:"errorClearTimeMsec"`
	// A list of failover conditions.
	//
	// If any of these conditions occur, MediaLive will perform a failover to the other input.
	FailoverConditions interface{} `field:"optional" json:"failoverConditions" yaml:"failoverConditions"`
	// Input preference when deciding which input to make active when a previously failed input has recovered.
	InputPreference *string `field:"optional" json:"inputPreference" yaml:"inputPreference"`
	// The input ID of the secondary input in the automatic input failover pair.
	SecondaryInputId *string `field:"optional" json:"secondaryInputId" yaml:"secondaryInputId"`
}

// The configuration of ad avail blanking in the output.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   availBlankingProperty := &availBlankingProperty{
//   	availBlankingImage: &inputLocationProperty{
//   		passwordParam: jsii.String("passwordParam"),
//   		uri: jsii.String("uri"),
//   		username: jsii.String("username"),
//   	},
//   	state: jsii.String("state"),
//   }
//
type CfnChannel_AvailBlankingProperty struct {
	// The blanking image to be used.
	//
	// Keep empty for solid black. Only .bmp and .png images are supported.
	AvailBlankingImage interface{} `field:"optional" json:"availBlankingImage" yaml:"availBlankingImage"`
	// When set to enabled, the video, audio, and captions are blanked when insertion metadata is added.
	State *string `field:"optional" json:"state" yaml:"state"`
}

// The setup of ad avail handling in the output.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   availConfigurationProperty := &availConfigurationProperty{
//   	availSettings: &availSettingsProperty{
//   		scte35SpliceInsert: &scte35SpliceInsertProperty{
//   			adAvailOffset: jsii.Number(123),
//   			noRegionalBlackoutFlag: jsii.String("noRegionalBlackoutFlag"),
//   			webDeliveryAllowedFlag: jsii.String("webDeliveryAllowedFlag"),
//   		},
//   		scte35TimeSignalApos: &scte35TimeSignalAposProperty{
//   			adAvailOffset: jsii.Number(123),
//   			noRegionalBlackoutFlag: jsii.String("noRegionalBlackoutFlag"),
//   			webDeliveryAllowedFlag: jsii.String("webDeliveryAllowedFlag"),
//   		},
//   	},
//   }
//
type CfnChannel_AvailConfigurationProperty struct {
	// The setup of ad avail handling in the output.
	AvailSettings interface{} `field:"optional" json:"availSettings" yaml:"availSettings"`
}

// The settings for the ad avail setup in the output.
//
// The parent of this entity is AvailConfiguration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   availSettingsProperty := &availSettingsProperty{
//   	scte35SpliceInsert: &scte35SpliceInsertProperty{
//   		adAvailOffset: jsii.Number(123),
//   		noRegionalBlackoutFlag: jsii.String("noRegionalBlackoutFlag"),
//   		webDeliveryAllowedFlag: jsii.String("webDeliveryAllowedFlag"),
//   	},
//   	scte35TimeSignalApos: &scte35TimeSignalAposProperty{
//   		adAvailOffset: jsii.Number(123),
//   		noRegionalBlackoutFlag: jsii.String("noRegionalBlackoutFlag"),
//   		webDeliveryAllowedFlag: jsii.String("webDeliveryAllowedFlag"),
//   	},
//   }
//
type CfnChannel_AvailSettingsProperty struct {
	// The setup for SCTE-35 splice insert handling.
	Scte35SpliceInsert interface{} `field:"optional" json:"scte35SpliceInsert" yaml:"scte35SpliceInsert"`
	// The setup for SCTE-35 time signal APOS handling.
	Scte35TimeSignalApos interface{} `field:"optional" json:"scte35TimeSignalApos" yaml:"scte35TimeSignalApos"`
}

// The settings for a blackout slate.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   blackoutSlateProperty := &blackoutSlateProperty{
//   	blackoutSlateImage: &inputLocationProperty{
//   		passwordParam: jsii.String("passwordParam"),
//   		uri: jsii.String("uri"),
//   		username: jsii.String("username"),
//   	},
//   	networkEndBlackout: jsii.String("networkEndBlackout"),
//   	networkEndBlackoutImage: &inputLocationProperty{
//   		passwordParam: jsii.String("passwordParam"),
//   		uri: jsii.String("uri"),
//   		username: jsii.String("username"),
//   	},
//   	networkId: jsii.String("networkId"),
//   	state: jsii.String("state"),
//   }
//
type CfnChannel_BlackoutSlateProperty struct {
	// The blackout slate image to be used.
	//
	// Keep empty for solid black. Only .bmp and .png images are supported.
	BlackoutSlateImage interface{} `field:"optional" json:"blackoutSlateImage" yaml:"blackoutSlateImage"`
	// Setting to enabled causes MediaLive to blackout the video, audio, and captions, and raise the "Network Blackout Image" slate when an SCTE104/35 Network End Segmentation Descriptor is encountered.
	//
	// The blackout is lifted when the Network Start Segmentation Descriptor is encountered. The Network End and Network Start descriptors must contain a network ID that matches the value entered in Network ID.
	NetworkEndBlackout *string `field:"optional" json:"networkEndBlackout" yaml:"networkEndBlackout"`
	// The path to the local file to use as the Network End Blackout image.
	//
	// The image is scaled to fill the entire output raster.
	NetworkEndBlackoutImage interface{} `field:"optional" json:"networkEndBlackoutImage" yaml:"networkEndBlackoutImage"`
	// Provides a Network ID that matches EIDR ID format (for example, "10.XXXX/XXXX-XXXX-XXXX-XXXX-XXXX-C").
	NetworkId *string `field:"optional" json:"networkId" yaml:"networkId"`
	// When set to enabled, this causes video, audio, and captions to be blanked when indicated by program metadata.
	State *string `field:"optional" json:"state" yaml:"state"`
}

// The settings for burn-in captions in the output.
//
// The parent of this entity is CaptionDestinationSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   burnInDestinationSettingsProperty := &burnInDestinationSettingsProperty{
//   	alignment: jsii.String("alignment"),
//   	backgroundColor: jsii.String("backgroundColor"),
//   	backgroundOpacity: jsii.Number(123),
//   	font: &inputLocationProperty{
//   		passwordParam: jsii.String("passwordParam"),
//   		uri: jsii.String("uri"),
//   		username: jsii.String("username"),
//   	},
//   	fontColor: jsii.String("fontColor"),
//   	fontOpacity: jsii.Number(123),
//   	fontResolution: jsii.Number(123),
//   	fontSize: jsii.String("fontSize"),
//   	outlineColor: jsii.String("outlineColor"),
//   	outlineSize: jsii.Number(123),
//   	shadowColor: jsii.String("shadowColor"),
//   	shadowOpacity: jsii.Number(123),
//   	shadowXOffset: jsii.Number(123),
//   	shadowYOffset: jsii.Number(123),
//   	teletextGridControl: jsii.String("teletextGridControl"),
//   	xPosition: jsii.Number(123),
//   	yPosition: jsii.Number(123),
//   }
//
type CfnChannel_BurnInDestinationSettingsProperty struct {
	// If no explicit xPosition or yPosition is provided, setting alignment to centered places the captions at the bottom center of the output.
	//
	// Similarly, setting a left alignment aligns captions to the bottom left of the output. If x and y positions are specified in conjunction with the alignment parameter, the font is justified (either left or centered) relative to those coordinates. Selecting "smart" justification left-justifies live subtitles and center-justifies pre-recorded subtitles. All burn-in and DVB-Sub font settings must match.
	Alignment *string `field:"optional" json:"alignment" yaml:"alignment"`
	// Specifies the color of the rectangle behind the captions.
	//
	// All burn-in and DVB-Sub font settings must match.
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// Specifies the opacity of the background rectangle.
	//
	// 255 is opaque; 0 is transparent. Keeping this parameter blank is equivalent to setting it to 0 (transparent). All burn-in and DVB-Sub font settings must match.
	BackgroundOpacity *float64 `field:"optional" json:"backgroundOpacity" yaml:"backgroundOpacity"`
	// The external font file that is used for captions burn-in.
	//
	// The file extension must be .ttf or .tte. Although you can select output fonts for many different types of input captions, embedded, STL, and Teletext sources use a strict grid system. Using external fonts with these captions sources could cause an unexpected display of proportional fonts. All burn-in and DVB-Sub font settings must match.
	Font interface{} `field:"optional" json:"font" yaml:"font"`
	// Specifies the color of the burned-in captions.
	//
	// This option is not valid for source captions that are STL, 608/embedded, or Teletext. These source settings are already pre-defined by the captions stream. All burn-in and DVB-Sub font settings must match.
	FontColor *string `field:"optional" json:"fontColor" yaml:"fontColor"`
	// Specifies the opacity of the burned-in captions.
	//
	// 255 is opaque; 0 is transparent. All burn-in and DVB-Sub font settings must match.
	FontOpacity *float64 `field:"optional" json:"fontOpacity" yaml:"fontOpacity"`
	// The font resolution in DPI (dots per inch).
	//
	// The default is 96 dpi. All burn-in and DVB-Sub font settings must match.
	FontResolution *float64 `field:"optional" json:"fontResolution" yaml:"fontResolution"`
	// When set to auto, fontSize scales depending on the size of the output.
	//
	// Providing a positive integer specifies the exact font size in points. All burn-in and DVB-Sub font settings must match.
	FontSize *string `field:"optional" json:"fontSize" yaml:"fontSize"`
	// Specifies the font outline color.
	//
	// This option is not valid for source captions that are either 608/embedded or Teletext. These source settings are already pre-defined by the captions stream. All burn-in and DVB-Sub font settings must match.
	OutlineColor *string `field:"optional" json:"outlineColor" yaml:"outlineColor"`
	// Specifies font outline size in pixels.
	//
	// This option is not valid for source captions that are either 608/embedded or Teletext. These source settings are already pre-defined by the captions stream. All burn-in and DVB-Sub font settings must match.
	OutlineSize *float64 `field:"optional" json:"outlineSize" yaml:"outlineSize"`
	// Specifies the color of the shadow cast by the captions.
	//
	// All burn-in and DVB-Sub font settings must match.
	ShadowColor *string `field:"optional" json:"shadowColor" yaml:"shadowColor"`
	// Specifies the opacity of the shadow.
	//
	// 255 is opaque; 0 is transparent. Keeping this parameter blank is equivalent to setting it to 0 (transparent). All burn-in and DVB-Sub font settings must match.
	ShadowOpacity *float64 `field:"optional" json:"shadowOpacity" yaml:"shadowOpacity"`
	// Specifies the horizontal offset of the shadow that is relative to the captions in pixels.
	//
	// A value of -2 would result in a shadow offset 2 pixels to the left. All burn-in and DVB-Sub font settings must match.
	ShadowXOffset *float64 `field:"optional" json:"shadowXOffset" yaml:"shadowXOffset"`
	// Specifies the vertical offset of the shadow that is relative to the captions in pixels.
	//
	// A value of -2 would result in a shadow offset 2 pixels above the text. All burn-in and DVB-Sub font settings must match.
	ShadowYOffset *float64 `field:"optional" json:"shadowYOffset" yaml:"shadowYOffset"`
	// Controls whether a fixed grid size is used to generate the output subtitles bitmap.
	//
	// This applies only to Teletext inputs and DVB-Sub/Burn-in outputs.
	TeletextGridControl *string `field:"optional" json:"teletextGridControl" yaml:"teletextGridControl"`
	// Specifies the horizontal position of the captions relative to the left side of the output in pixels.
	//
	// A value of 10 would result in the captions starting 10 pixels from the left of the output. If no explicit xPosition is provided, the horizontal captions position is determined by the alignment parameter. All burn-in and DVB-Sub font settings must match.
	XPosition *float64 `field:"optional" json:"xPosition" yaml:"xPosition"`
	// Specifies the vertical position of the captions relative to the top of the output in pixels.
	//
	// A value of 10 would result in the captions starting 10 pixels from the top of the output. If no explicit yPosition is provided, the captions are positioned towards the bottom of the output. All burn-in and DVB-Sub font settings must match.
	YPosition *float64 `field:"optional" json:"yPosition" yaml:"yPosition"`
}

// The encoding information for output captions.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   captionDescriptionProperty := &captionDescriptionProperty{
//   	captionSelectorName: jsii.String("captionSelectorName"),
//   	destinationSettings: &captionDestinationSettingsProperty{
//   		aribDestinationSettings: &aribDestinationSettingsProperty{
//   		},
//   		burnInDestinationSettings: &burnInDestinationSettingsProperty{
//   			alignment: jsii.String("alignment"),
//   			backgroundColor: jsii.String("backgroundColor"),
//   			backgroundOpacity: jsii.Number(123),
//   			font: &inputLocationProperty{
//   				passwordParam: jsii.String("passwordParam"),
//   				uri: jsii.String("uri"),
//   				username: jsii.String("username"),
//   			},
//   			fontColor: jsii.String("fontColor"),
//   			fontOpacity: jsii.Number(123),
//   			fontResolution: jsii.Number(123),
//   			fontSize: jsii.String("fontSize"),
//   			outlineColor: jsii.String("outlineColor"),
//   			outlineSize: jsii.Number(123),
//   			shadowColor: jsii.String("shadowColor"),
//   			shadowOpacity: jsii.Number(123),
//   			shadowXOffset: jsii.Number(123),
//   			shadowYOffset: jsii.Number(123),
//   			teletextGridControl: jsii.String("teletextGridControl"),
//   			xPosition: jsii.Number(123),
//   			yPosition: jsii.Number(123),
//   		},
//   		dvbSubDestinationSettings: &dvbSubDestinationSettingsProperty{
//   			alignment: jsii.String("alignment"),
//   			backgroundColor: jsii.String("backgroundColor"),
//   			backgroundOpacity: jsii.Number(123),
//   			font: &inputLocationProperty{
//   				passwordParam: jsii.String("passwordParam"),
//   				uri: jsii.String("uri"),
//   				username: jsii.String("username"),
//   			},
//   			fontColor: jsii.String("fontColor"),
//   			fontOpacity: jsii.Number(123),
//   			fontResolution: jsii.Number(123),
//   			fontSize: jsii.String("fontSize"),
//   			outlineColor: jsii.String("outlineColor"),
//   			outlineSize: jsii.Number(123),
//   			shadowColor: jsii.String("shadowColor"),
//   			shadowOpacity: jsii.Number(123),
//   			shadowXOffset: jsii.Number(123),
//   			shadowYOffset: jsii.Number(123),
//   			teletextGridControl: jsii.String("teletextGridControl"),
//   			xPosition: jsii.Number(123),
//   			yPosition: jsii.Number(123),
//   		},
//   		ebuTtDDestinationSettings: &ebuTtDDestinationSettingsProperty{
//   			copyrightHolder: jsii.String("copyrightHolder"),
//   			fillLineGap: jsii.String("fillLineGap"),
//   			fontFamily: jsii.String("fontFamily"),
//   			styleControl: jsii.String("styleControl"),
//   		},
//   		embeddedDestinationSettings: &embeddedDestinationSettingsProperty{
//   		},
//   		embeddedPlusScte20DestinationSettings: &embeddedPlusScte20DestinationSettingsProperty{
//   		},
//   		rtmpCaptionInfoDestinationSettings: &rtmpCaptionInfoDestinationSettingsProperty{
//   		},
//   		scte20PlusEmbeddedDestinationSettings: &scte20PlusEmbeddedDestinationSettingsProperty{
//   		},
//   		scte27DestinationSettings: &scte27DestinationSettingsProperty{
//   		},
//   		smpteTtDestinationSettings: &smpteTtDestinationSettingsProperty{
//   		},
//   		teletextDestinationSettings: &teletextDestinationSettingsProperty{
//   		},
//   		ttmlDestinationSettings: &ttmlDestinationSettingsProperty{
//   			styleControl: jsii.String("styleControl"),
//   		},
//   		webvttDestinationSettings: &webvttDestinationSettingsProperty{
//   			styleControl: jsii.String("styleControl"),
//   		},
//   	},
//   	languageCode: jsii.String("languageCode"),
//   	languageDescription: jsii.String("languageDescription"),
//   	name: jsii.String("name"),
//   }
//
type CfnChannel_CaptionDescriptionProperty struct {
	// Specifies which input captions selector to use as a captions source when generating output captions.
	//
	// This field should match a captionSelector name.
	CaptionSelectorName *string `field:"optional" json:"captionSelectorName" yaml:"captionSelectorName"`
	// Additional settings for a captions destination that depend on the destination type.
	DestinationSettings interface{} `field:"optional" json:"destinationSettings" yaml:"destinationSettings"`
	// An ISO 639-2 three-digit code.
	//
	// For more information, see http://www.loc.gov/standards/iso639-2/.
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// Human-readable information to indicate the captions that are available for players (for example, English or Spanish).
	LanguageDescription *string `field:"optional" json:"languageDescription" yaml:"languageDescription"`
	// The name of the captions description.
	//
	// The name is used to associate a captions description with an output. Names must be unique within a channel.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

// The configuration of one captions encode in the output.
//
// The parent of this entity is CaptionDescription.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   captionDestinationSettingsProperty := &captionDestinationSettingsProperty{
//   	aribDestinationSettings: &aribDestinationSettingsProperty{
//   	},
//   	burnInDestinationSettings: &burnInDestinationSettingsProperty{
//   		alignment: jsii.String("alignment"),
//   		backgroundColor: jsii.String("backgroundColor"),
//   		backgroundOpacity: jsii.Number(123),
//   		font: &inputLocationProperty{
//   			passwordParam: jsii.String("passwordParam"),
//   			uri: jsii.String("uri"),
//   			username: jsii.String("username"),
//   		},
//   		fontColor: jsii.String("fontColor"),
//   		fontOpacity: jsii.Number(123),
//   		fontResolution: jsii.Number(123),
//   		fontSize: jsii.String("fontSize"),
//   		outlineColor: jsii.String("outlineColor"),
//   		outlineSize: jsii.Number(123),
//   		shadowColor: jsii.String("shadowColor"),
//   		shadowOpacity: jsii.Number(123),
//   		shadowXOffset: jsii.Number(123),
//   		shadowYOffset: jsii.Number(123),
//   		teletextGridControl: jsii.String("teletextGridControl"),
//   		xPosition: jsii.Number(123),
//   		yPosition: jsii.Number(123),
//   	},
//   	dvbSubDestinationSettings: &dvbSubDestinationSettingsProperty{
//   		alignment: jsii.String("alignment"),
//   		backgroundColor: jsii.String("backgroundColor"),
//   		backgroundOpacity: jsii.Number(123),
//   		font: &inputLocationProperty{
//   			passwordParam: jsii.String("passwordParam"),
//   			uri: jsii.String("uri"),
//   			username: jsii.String("username"),
//   		},
//   		fontColor: jsii.String("fontColor"),
//   		fontOpacity: jsii.Number(123),
//   		fontResolution: jsii.Number(123),
//   		fontSize: jsii.String("fontSize"),
//   		outlineColor: jsii.String("outlineColor"),
//   		outlineSize: jsii.Number(123),
//   		shadowColor: jsii.String("shadowColor"),
//   		shadowOpacity: jsii.Number(123),
//   		shadowXOffset: jsii.Number(123),
//   		shadowYOffset: jsii.Number(123),
//   		teletextGridControl: jsii.String("teletextGridControl"),
//   		xPosition: jsii.Number(123),
//   		yPosition: jsii.Number(123),
//   	},
//   	ebuTtDDestinationSettings: &ebuTtDDestinationSettingsProperty{
//   		copyrightHolder: jsii.String("copyrightHolder"),
//   		fillLineGap: jsii.String("fillLineGap"),
//   		fontFamily: jsii.String("fontFamily"),
//   		styleControl: jsii.String("styleControl"),
//   	},
//   	embeddedDestinationSettings: &embeddedDestinationSettingsProperty{
//   	},
//   	embeddedPlusScte20DestinationSettings: &embeddedPlusScte20DestinationSettingsProperty{
//   	},
//   	rtmpCaptionInfoDestinationSettings: &rtmpCaptionInfoDestinationSettingsProperty{
//   	},
//   	scte20PlusEmbeddedDestinationSettings: &scte20PlusEmbeddedDestinationSettingsProperty{
//   	},
//   	scte27DestinationSettings: &scte27DestinationSettingsProperty{
//   	},
//   	smpteTtDestinationSettings: &smpteTtDestinationSettingsProperty{
//   	},
//   	teletextDestinationSettings: &teletextDestinationSettingsProperty{
//   	},
//   	ttmlDestinationSettings: &ttmlDestinationSettingsProperty{
//   		styleControl: jsii.String("styleControl"),
//   	},
//   	webvttDestinationSettings: &webvttDestinationSettingsProperty{
//   		styleControl: jsii.String("styleControl"),
//   	},
//   }
//
type CfnChannel_CaptionDestinationSettingsProperty struct {
	// The configuration of one ARIB captions encode in the output.
	AribDestinationSettings interface{} `field:"optional" json:"aribDestinationSettings" yaml:"aribDestinationSettings"`
	// The configuration of one burn-in captions encode in the output.
	BurnInDestinationSettings interface{} `field:"optional" json:"burnInDestinationSettings" yaml:"burnInDestinationSettings"`
	// The configuration of one DVB Sub captions encode in the output.
	DvbSubDestinationSettings interface{} `field:"optional" json:"dvbSubDestinationSettings" yaml:"dvbSubDestinationSettings"`
	// Settings for EBU-TT captions in the output.
	EbuTtDDestinationSettings interface{} `field:"optional" json:"ebuTtDDestinationSettings" yaml:"ebuTtDDestinationSettings"`
	// The configuration of one embedded captions encode in the output.
	EmbeddedDestinationSettings interface{} `field:"optional" json:"embeddedDestinationSettings" yaml:"embeddedDestinationSettings"`
	// The configuration of one embedded plus SCTE-20 captions encode in the output.
	EmbeddedPlusScte20DestinationSettings interface{} `field:"optional" json:"embeddedPlusScte20DestinationSettings" yaml:"embeddedPlusScte20DestinationSettings"`
	// The configuration of one RTMPCaptionInfo captions encode in the output.
	RtmpCaptionInfoDestinationSettings interface{} `field:"optional" json:"rtmpCaptionInfoDestinationSettings" yaml:"rtmpCaptionInfoDestinationSettings"`
	// The configuration of one SCTE20 plus embedded captions encode in the output.
	Scte20PlusEmbeddedDestinationSettings interface{} `field:"optional" json:"scte20PlusEmbeddedDestinationSettings" yaml:"scte20PlusEmbeddedDestinationSettings"`
	// The configuration of one SCTE-27 captions encode in the output.
	Scte27DestinationSettings interface{} `field:"optional" json:"scte27DestinationSettings" yaml:"scte27DestinationSettings"`
	// The configuration of one SMPTE-TT captions encode in the output.
	SmpteTtDestinationSettings interface{} `field:"optional" json:"smpteTtDestinationSettings" yaml:"smpteTtDestinationSettings"`
	// The configuration of one Teletext captions encode in the output.
	TeletextDestinationSettings interface{} `field:"optional" json:"teletextDestinationSettings" yaml:"teletextDestinationSettings"`
	// The configuration of one TTML captions encode in the output.
	TtmlDestinationSettings interface{} `field:"optional" json:"ttmlDestinationSettings" yaml:"ttmlDestinationSettings"`
	// The configuration of one WebVTT captions encode in the output.
	WebvttDestinationSettings interface{} `field:"optional" json:"webvttDestinationSettings" yaml:"webvttDestinationSettings"`
}

// Maps a captions channel to an ISO 693-2 language code (http://www.loc.gov/standards/iso639-2), with an optional description.
//
// The parent of this entity is HlsGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   captionLanguageMappingProperty := &captionLanguageMappingProperty{
//   	captionChannel: jsii.Number(123),
//   	languageCode: jsii.String("languageCode"),
//   	languageDescription: jsii.String("languageDescription"),
//   }
//
type CfnChannel_CaptionLanguageMappingProperty struct {
	// The closed caption channel being described by this CaptionLanguageMapping.
	//
	// Each channel mapping must have a unique channel number (maximum of 4).
	CaptionChannel *float64 `field:"optional" json:"captionChannel" yaml:"captionChannel"`
	// A three-character ISO 639-2 language code (see http://www.loc.gov/standards/iso639-2).
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// The textual description of language.
	LanguageDescription *string `field:"optional" json:"languageDescription" yaml:"languageDescription"`
}

// Settings to configure the caption rectangle for an output captions that will be created using this Teletext source captions.
//
// The parent of this entity is TeletextSourceSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   captionRectangleProperty := &captionRectangleProperty{
//   	height: jsii.Number(123),
//   	leftOffset: jsii.Number(123),
//   	topOffset: jsii.Number(123),
//   	width: jsii.Number(123),
//   }
//
type CfnChannel_CaptionRectangleProperty struct {
	// See the description in leftOffset.
	//
	// For height, specify the entire height of the rectangle as a percentage of the underlying frame height. For example, \"80\" means the rectangle height is 80% of the underlying frame height. The topOffset and rectangleHeight must add up to 100% or less. This field corresponds to tts:extent - Y in the TTML standard.
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// Applies only if you plan to convert these source captions to EBU-TT-D or TTML in an output.
	//
	// (Make sure to leave the default if you don't have either of these formats in the output.) You can define a display rectangle for the captions that is smaller than the underlying video frame. You define the rectangle by specifying the position of the left edge, top edge, bottom edge, and right edge of the rectangle, all within the underlying video frame. The units for the measurements are percentages. If you specify a value for one of these fields, you must specify a value for all of them.
	//
	// For leftOffset, specify the position of the left edge of the rectangle, as a percentage of the underlying frame width, and relative to the left edge of the frame. For example, \"10\" means the measurement is 10% of the underlying frame width. The rectangle left edge starts at that position from the left edge of the frame. This field corresponds to tts:origin - X in the TTML standard.
	LeftOffset *float64 `field:"optional" json:"leftOffset" yaml:"leftOffset"`
	// See the description in leftOffset.
	//
	// For topOffset, specify the position of the top edge of the rectangle, as a percentage of the underlying frame height, and relative to the top edge of the frame. For example, \"10\" means the measurement is 10% of the underlying frame height. The rectangle top edge starts at that position from the top edge of the frame. This field corresponds to tts:origin - Y in the TTML standard.
	TopOffset *float64 `field:"optional" json:"topOffset" yaml:"topOffset"`
	// See the description in leftOffset.
	//
	// For width, specify the entire width of the rectangle as a percentage of the underlying frame width. For example, \"80\" means the rectangle width is 80% of the underlying frame width. The leftOffset and rectangleWidth must add up to 100% or less. This field corresponds to tts:extent - X in the TTML standard.
	Width *float64 `field:"optional" json:"width" yaml:"width"`
}

// Information about one caption to extract from the input.
//
// The parent of this entity is InputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   captionSelectorProperty := &captionSelectorProperty{
//   	languageCode: jsii.String("languageCode"),
//   	name: jsii.String("name"),
//   	selectorSettings: &captionSelectorSettingsProperty{
//   		ancillarySourceSettings: &ancillarySourceSettingsProperty{
//   			sourceAncillaryChannelNumber: jsii.Number(123),
//   		},
//   		aribSourceSettings: &aribSourceSettingsProperty{
//   		},
//   		dvbSubSourceSettings: &dvbSubSourceSettingsProperty{
//   			ocrLanguage: jsii.String("ocrLanguage"),
//   			pid: jsii.Number(123),
//   		},
//   		embeddedSourceSettings: &embeddedSourceSettingsProperty{
//   			convert608To708: jsii.String("convert608To708"),
//   			scte20Detection: jsii.String("scte20Detection"),
//   			source608ChannelNumber: jsii.Number(123),
//   			source608TrackNumber: jsii.Number(123),
//   		},
//   		scte20SourceSettings: &scte20SourceSettingsProperty{
//   			convert608To708: jsii.String("convert608To708"),
//   			source608ChannelNumber: jsii.Number(123),
//   		},
//   		scte27SourceSettings: &scte27SourceSettingsProperty{
//   			ocrLanguage: jsii.String("ocrLanguage"),
//   			pid: jsii.Number(123),
//   		},
//   		teletextSourceSettings: &teletextSourceSettingsProperty{
//   			outputRectangle: &captionRectangleProperty{
//   				height: jsii.Number(123),
//   				leftOffset: jsii.Number(123),
//   				topOffset: jsii.Number(123),
//   				width: jsii.Number(123),
//   			},
//   			pageNumber: jsii.String("pageNumber"),
//   		},
//   	},
//   }
//
type CfnChannel_CaptionSelectorProperty struct {
	// When specified, this field indicates the three-letter language code of the captions track to extract from the source.
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// The name identifier for a captions selector.
	//
	// This name is used to associate this captions selector with one or more captions descriptions. Names must be unique within a channel.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Information about the specific audio to extract from the input.
	SelectorSettings interface{} `field:"optional" json:"selectorSettings" yaml:"selectorSettings"`
}

// Captions Selector Settings.
//
// The parent of this entity is CaptionSelector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   captionSelectorSettingsProperty := &captionSelectorSettingsProperty{
//   	ancillarySourceSettings: &ancillarySourceSettingsProperty{
//   		sourceAncillaryChannelNumber: jsii.Number(123),
//   	},
//   	aribSourceSettings: &aribSourceSettingsProperty{
//   	},
//   	dvbSubSourceSettings: &dvbSubSourceSettingsProperty{
//   		ocrLanguage: jsii.String("ocrLanguage"),
//   		pid: jsii.Number(123),
//   	},
//   	embeddedSourceSettings: &embeddedSourceSettingsProperty{
//   		convert608To708: jsii.String("convert608To708"),
//   		scte20Detection: jsii.String("scte20Detection"),
//   		source608ChannelNumber: jsii.Number(123),
//   		source608TrackNumber: jsii.Number(123),
//   	},
//   	scte20SourceSettings: &scte20SourceSettingsProperty{
//   		convert608To708: jsii.String("convert608To708"),
//   		source608ChannelNumber: jsii.Number(123),
//   	},
//   	scte27SourceSettings: &scte27SourceSettingsProperty{
//   		ocrLanguage: jsii.String("ocrLanguage"),
//   		pid: jsii.Number(123),
//   	},
//   	teletextSourceSettings: &teletextSourceSettingsProperty{
//   		outputRectangle: &captionRectangleProperty{
//   			height: jsii.Number(123),
//   			leftOffset: jsii.Number(123),
//   			topOffset: jsii.Number(123),
//   			width: jsii.Number(123),
//   		},
//   		pageNumber: jsii.String("pageNumber"),
//   	},
//   }
//
type CfnChannel_CaptionSelectorSettingsProperty struct {
	// Information about the ancillary captions to extract from the input.
	AncillarySourceSettings interface{} `field:"optional" json:"ancillarySourceSettings" yaml:"ancillarySourceSettings"`
	// Information about the ARIB captions to extract from the input.
	AribSourceSettings interface{} `field:"optional" json:"aribSourceSettings" yaml:"aribSourceSettings"`
	// Information about the DVB Sub captions to extract from the input.
	DvbSubSourceSettings interface{} `field:"optional" json:"dvbSubSourceSettings" yaml:"dvbSubSourceSettings"`
	// Information about the embedded captions to extract from the input.
	EmbeddedSourceSettings interface{} `field:"optional" json:"embeddedSourceSettings" yaml:"embeddedSourceSettings"`
	// Information about the SCTE-20 captions to extract from the input.
	Scte20SourceSettings interface{} `field:"optional" json:"scte20SourceSettings" yaml:"scte20SourceSettings"`
	// Information about the SCTE-27 captions to extract from the input.
	Scte27SourceSettings interface{} `field:"optional" json:"scte27SourceSettings" yaml:"scte27SourceSettings"`
	// Information about the Teletext captions to extract from the input.
	TeletextSourceSettings interface{} `field:"optional" json:"teletextSourceSettings" yaml:"teletextSourceSettings"`
}

// The input specification for this channel.
//
// It specifies the key characteristics of CDI inputs for this channel, when those characteristics are different from other inputs.
//
// This entity is at the top level in the channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cdiInputSpecificationProperty := &cdiInputSpecificationProperty{
//   	resolution: jsii.String("resolution"),
//   }
//
type CfnChannel_CdiInputSpecificationProperty struct {
	// Maximum CDI input resolution.
	Resolution *string `field:"optional" json:"resolution" yaml:"resolution"`
}

// Passthrough applies no color space conversion to the output.
//
// The parents of this entity are H264ColorSpaceSettings and H265ColorSpaceSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   colorSpacePassthroughSettingsProperty := &colorSpacePassthroughSettingsProperty{
//   }
//
type CfnChannel_ColorSpacePassthroughSettingsProperty struct {
}

// The configuration of DVB NIT.
//
// The parent of this entity is M2tsSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dvbNitSettingsProperty := &dvbNitSettingsProperty{
//   	networkId: jsii.Number(123),
//   	networkName: jsii.String("networkName"),
//   	repInterval: jsii.Number(123),
//   }
//
type CfnChannel_DvbNitSettingsProperty struct {
	// The numeric value placed in the Network Information Table (NIT).
	NetworkId *float64 `field:"optional" json:"networkId" yaml:"networkId"`
	// The network name text placed in the networkNameDescriptor inside the Network Information Table (NIT).
	//
	// The maximum length is 256 characters.
	NetworkName *string `field:"optional" json:"networkName" yaml:"networkName"`
	// The number of milliseconds between instances of this table in the output transport stream.
	RepInterval *float64 `field:"optional" json:"repInterval" yaml:"repInterval"`
}

// A DVB Service Description Table (SDT).
//
// The parent of this entity is M2tsSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dvbSdtSettingsProperty := &dvbSdtSettingsProperty{
//   	outputSdt: jsii.String("outputSdt"),
//   	repInterval: jsii.Number(123),
//   	serviceName: jsii.String("serviceName"),
//   	serviceProviderName: jsii.String("serviceProviderName"),
//   }
//
type CfnChannel_DvbSdtSettingsProperty struct {
	// Selects a method of inserting SDT information into an output stream.
	//
	// The sdtFollow setting copies SDT information from input stream to output stream. The sdtFollowIfPresent setting copies SDT information from input stream to output stream if SDT information is present in the input. Otherwise, it falls back on the user-defined values. The sdtManual setting means that the user will enter the SDT information. The sdtNone setting means that the output stream will not contain SDT information.
	OutputSdt *string `field:"optional" json:"outputSdt" yaml:"outputSdt"`
	// The number of milliseconds between instances of this table in the output transport stream.
	RepInterval *float64 `field:"optional" json:"repInterval" yaml:"repInterval"`
	// The service name placed in the serviceDescriptor in the Service Description Table (SDT).
	//
	// The maximum length is 256 characters.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// The service provider name placed in the serviceDescriptor in the Service Description Table (SDT).
	//
	// The maximum length is 256 characters.
	ServiceProviderName *string `field:"optional" json:"serviceProviderName" yaml:"serviceProviderName"`
}

// The settings for DVB Sub captions in the output.
//
// The parent of this entity is CaptionDestinationSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dvbSubDestinationSettingsProperty := &dvbSubDestinationSettingsProperty{
//   	alignment: jsii.String("alignment"),
//   	backgroundColor: jsii.String("backgroundColor"),
//   	backgroundOpacity: jsii.Number(123),
//   	font: &inputLocationProperty{
//   		passwordParam: jsii.String("passwordParam"),
//   		uri: jsii.String("uri"),
//   		username: jsii.String("username"),
//   	},
//   	fontColor: jsii.String("fontColor"),
//   	fontOpacity: jsii.Number(123),
//   	fontResolution: jsii.Number(123),
//   	fontSize: jsii.String("fontSize"),
//   	outlineColor: jsii.String("outlineColor"),
//   	outlineSize: jsii.Number(123),
//   	shadowColor: jsii.String("shadowColor"),
//   	shadowOpacity: jsii.Number(123),
//   	shadowXOffset: jsii.Number(123),
//   	shadowYOffset: jsii.Number(123),
//   	teletextGridControl: jsii.String("teletextGridControl"),
//   	xPosition: jsii.Number(123),
//   	yPosition: jsii.Number(123),
//   }
//
type CfnChannel_DvbSubDestinationSettingsProperty struct {
	// If no explicit xPosition or yPosition is provided, setting the alignment to centered places the captions at the bottom center of the output.
	//
	// Similarly, setting a left alignment aligns captions to the bottom left of the output. If x and y positions are specified in conjunction with the alignment parameter, the font is justified (either left or centered) relative to those coordinates. Selecting "smart" justification left-justifies live subtitles and center-justifies pre-recorded subtitles. This option is not valid for source captions that are STL or 608/embedded. These source settings are already pre-defined by the captions stream. All burn-in and DVB-Sub font settings must match.
	Alignment *string `field:"optional" json:"alignment" yaml:"alignment"`
	// Specifies the color of the rectangle behind the captions.
	//
	// All burn-in and DVB-Sub font settings must match.
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// Specifies the opacity of the background rectangle.
	//
	// 255 is opaque; 0 is transparent. Keeping this parameter blank is equivalent to setting it to 0 (transparent). All burn-in and DVB-Sub font settings must match.
	BackgroundOpacity *float64 `field:"optional" json:"backgroundOpacity" yaml:"backgroundOpacity"`
	// The external font file that is used for captions burn-in.
	//
	// The file extension must be .ttf or .tte. Although you can select output fonts for many different types of input captions, embedded, STL, and Teletext sources use a strict grid system. Using external fonts with these captions sources could cause an unexpected display of proportional fonts. All burn-in and DVB-Sub font settings must match.
	Font interface{} `field:"optional" json:"font" yaml:"font"`
	// Specifies the color of the burned-in captions.
	//
	// This option is not valid for source captions that are STL, 608/embedded, or Teletext. These source settings are already pre-defined by the captions stream. All burn-in and DVB-Sub font settings must match.
	FontColor *string `field:"optional" json:"fontColor" yaml:"fontColor"`
	// Specifies the opacity of the burned-in captions.
	//
	// 255 is opaque; 0 is transparent. All burn-in and DVB-Sub font settings must match.
	FontOpacity *float64 `field:"optional" json:"fontOpacity" yaml:"fontOpacity"`
	// The font resolution in DPI (dots per inch).
	//
	// The default is 96 dpi. All burn-in and DVB-Sub font settings must match.
	FontResolution *float64 `field:"optional" json:"fontResolution" yaml:"fontResolution"`
	// When set to auto, fontSize scales depending on the size of the output.
	//
	// Providing a positive integer specifies the exact font size in points. All burn-in and DVB-Sub font settings must match.
	FontSize *string `field:"optional" json:"fontSize" yaml:"fontSize"`
	// Specifies the font outline color.
	//
	// This option is not valid for source captions that are either 608/embedded or Teletext. These source settings are already pre-defined by the captions stream. All burn-in and DVB-Sub font settings must match.
	OutlineColor *string `field:"optional" json:"outlineColor" yaml:"outlineColor"`
	// Specifies the font outline size in pixels.
	//
	// This option is not valid for source captions that are either 608/embedded or Teletext. These source settings are already pre-defined by the captions stream. All burn-in and DVB-Sub font settings must match.
	OutlineSize *float64 `field:"optional" json:"outlineSize" yaml:"outlineSize"`
	// Specifies the color of the shadow that is cast by the captions.
	//
	// All burn-in and DVB-Sub font settings must match.
	ShadowColor *string `field:"optional" json:"shadowColor" yaml:"shadowColor"`
	// Specifies the opacity of the shadow.
	//
	// 255 is opaque; 0 is transparent. Keeping this parameter blank is equivalent to setting it to 0 (transparent). All burn-in and DVB-Sub font settings must match.
	ShadowOpacity *float64 `field:"optional" json:"shadowOpacity" yaml:"shadowOpacity"`
	// Specifies the horizontal offset of the shadow relative to the captions in pixels.
	//
	// A value of -2 would result in a shadow offset 2 pixels to the left. All burn-in and DVB-Sub font settings must match.
	ShadowXOffset *float64 `field:"optional" json:"shadowXOffset" yaml:"shadowXOffset"`
	// Specifies the vertical offset of the shadow relative to the captions in pixels.
	//
	// A value of -2 would result in a shadow offset 2 pixels above the text. All burn-in and DVB-Sub font settings must match.
	ShadowYOffset *float64 `field:"optional" json:"shadowYOffset" yaml:"shadowYOffset"`
	// Controls whether a fixed grid size is used to generate the output subtitles bitmap.
	//
	// This applies to only Teletext inputs and DVB-Sub/Burn-in outputs.
	TeletextGridControl *string `field:"optional" json:"teletextGridControl" yaml:"teletextGridControl"`
	// Specifies the horizontal position of the captions relative to the left side of the output in pixels.
	//
	// A value of 10 would result in the captions starting 10 pixels from the left of the output. If no explicit xPosition is provided, the horizontal captions position is determined by the alignment parameter. This option is not valid for source captions that are STL, 608/embedded, or Teletext. These source settings are already pre-defined by the captions stream. All burn-in and DVB-Sub font settings must match.
	XPosition *float64 `field:"optional" json:"xPosition" yaml:"xPosition"`
	// Specifies the vertical position of the captions relative to the top of the output in pixels.
	//
	// A value of 10 would result in the captions starting 10 pixels from the top of the output. If no explicit yPosition is provided, the captions are positioned towards the bottom of the output. This option is not valid for source captions that are STL, 608/embedded, or Teletext. These source settings are already pre-defined by the captions stream. All burn-in and DVB-Sub font settings must match.
	YPosition *float64 `field:"optional" json:"yPosition" yaml:"yPosition"`
}

// Information about the DVB Sub captions to extract from the input.
//
// The parent of this entity is CaptionSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dvbSubSourceSettingsProperty := &dvbSubSourceSettingsProperty{
//   	ocrLanguage: jsii.String("ocrLanguage"),
//   	pid: jsii.Number(123),
//   }
//
type CfnChannel_DvbSubSourceSettingsProperty struct {
	// If you will configure a WebVTT caption description that references this caption selector, use this field to provide the language to consider when translating the image-based source to text.
	OcrLanguage *string `field:"optional" json:"ocrLanguage" yaml:"ocrLanguage"`
	// When using DVB-Sub with burn-in or SMPTE-TT, use this PID for the source content.
	//
	// It is unused for DVB-Sub passthrough. All DVB-Sub content is passed through, regardless of selectors.
	Pid *float64 `field:"optional" json:"pid" yaml:"pid"`
}

// The DVB Time and Date Table (TDT).
//
// The parent of this entity is M2tsSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dvbTdtSettingsProperty := &dvbTdtSettingsProperty{
//   	repInterval: jsii.Number(123),
//   }
//
type CfnChannel_DvbTdtSettingsProperty struct {
	// The number of milliseconds between instances of this table in the output transport stream.
	RepInterval *float64 `field:"optional" json:"repInterval" yaml:"repInterval"`
}

// The settings for an EAC3 audio encode in the output.
//
// The parent of this entity is AudioCodecSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eac3SettingsProperty := &eac3SettingsProperty{
//   	attenuationControl: jsii.String("attenuationControl"),
//   	bitrate: jsii.Number(123),
//   	bitstreamMode: jsii.String("bitstreamMode"),
//   	codingMode: jsii.String("codingMode"),
//   	dcFilter: jsii.String("dcFilter"),
//   	dialnorm: jsii.Number(123),
//   	drcLine: jsii.String("drcLine"),
//   	drcRf: jsii.String("drcRf"),
//   	lfeControl: jsii.String("lfeControl"),
//   	lfeFilter: jsii.String("lfeFilter"),
//   	loRoCenterMixLevel: jsii.Number(123),
//   	loRoSurroundMixLevel: jsii.Number(123),
//   	ltRtCenterMixLevel: jsii.Number(123),
//   	ltRtSurroundMixLevel: jsii.Number(123),
//   	metadataControl: jsii.String("metadataControl"),
//   	passthroughControl: jsii.String("passthroughControl"),
//   	phaseControl: jsii.String("phaseControl"),
//   	stereoDownmix: jsii.String("stereoDownmix"),
//   	surroundExMode: jsii.String("surroundExMode"),
//   	surroundMode: jsii.String("surroundMode"),
//   }
//
type CfnChannel_Eac3SettingsProperty struct {
	// When set to attenuate3Db, applies a 3 dB attenuation to the surround channels.
	//
	// Used only for the 3/2 coding mode.
	AttenuationControl *string `field:"optional" json:"attenuationControl" yaml:"attenuationControl"`
	// The average bitrate in bits/second.
	//
	// Valid bitrates depend on the coding mode.
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// Specifies the bitstream mode (bsmod) for the emitted E-AC-3 stream.
	//
	// For more information, see ATSC A/52-2012 (Annex E).
	BitstreamMode *string `field:"optional" json:"bitstreamMode" yaml:"bitstreamMode"`
	// The Dolby Digital Plus coding mode.
	//
	// This mode determines the number of channels.
	CodingMode *string `field:"optional" json:"codingMode" yaml:"codingMode"`
	// When set to enabled, activates a DC highpass filter for all input channels.
	DcFilter *string `field:"optional" json:"dcFilter" yaml:"dcFilter"`
	// Sets the dialnorm for the output.
	//
	// If blank and the input audio is Dolby Digital Plus, dialnorm will be passed through.
	Dialnorm *float64 `field:"optional" json:"dialnorm" yaml:"dialnorm"`
	// Sets the Dolby dynamic range compression profile.
	DrcLine *string `field:"optional" json:"drcLine" yaml:"drcLine"`
	// Sets the profile for heavy Dolby dynamic range compression, ensuring that the instantaneous signal peaks do not exceed specified levels.
	DrcRf *string `field:"optional" json:"drcRf" yaml:"drcRf"`
	// When encoding 3/2 audio, setting to lfe enables the LFE channel.
	LfeControl *string `field:"optional" json:"lfeControl" yaml:"lfeControl"`
	// When set to enabled, applies a 120Hz lowpass filter to the LFE channel prior to encoding.
	//
	// Valid only with a codingMode32 coding mode.
	LfeFilter *string `field:"optional" json:"lfeFilter" yaml:"lfeFilter"`
	// The Left only/Right only center mix level.
	//
	// Used only for the 3/2 coding mode.
	LoRoCenterMixLevel *float64 `field:"optional" json:"loRoCenterMixLevel" yaml:"loRoCenterMixLevel"`
	// The Left only/Right only surround mix level.
	//
	// Used only for a 3/2 coding mode.
	LoRoSurroundMixLevel *float64 `field:"optional" json:"loRoSurroundMixLevel" yaml:"loRoSurroundMixLevel"`
	// The Left total/Right total center mix level.
	//
	// Used only for a 3/2 coding mode.
	LtRtCenterMixLevel *float64 `field:"optional" json:"ltRtCenterMixLevel" yaml:"ltRtCenterMixLevel"`
	// The Left total/Right total surround mix level.
	//
	// Used only for the 3/2 coding mode.
	LtRtSurroundMixLevel *float64 `field:"optional" json:"ltRtSurroundMixLevel" yaml:"ltRtSurroundMixLevel"`
	// When set to followInput, encoder metadata is sourced from the DD, DD+, or DolbyE decoder that supplies this audio data.
	//
	// If the audio is not supplied from one of these streams, then the static metadata settings are used.
	MetadataControl *string `field:"optional" json:"metadataControl" yaml:"metadataControl"`
	// When set to whenPossible, input DD+ audio will be passed through if it is present on the input.
	//
	// This detection is dynamic over the life of the transcode. Inputs that alternate between DD+ and non-DD+ content will have a consistent DD+ output as the system alternates between passthrough and encoding.
	PassthroughControl *string `field:"optional" json:"passthroughControl" yaml:"passthroughControl"`
	// When set to shift90Degrees, applies a 90-degree phase shift to the surround channels.
	//
	// Used only for a 3/2 coding mode.
	PhaseControl *string `field:"optional" json:"phaseControl" yaml:"phaseControl"`
	// A stereo downmix preference.
	//
	// Used only for the 3/2 coding mode.
	StereoDownmix *string `field:"optional" json:"stereoDownmix" yaml:"stereoDownmix"`
	// When encoding 3/2 audio, sets whether an extra center back surround channel is matrix encoded into the left and right surround channels.
	SurroundExMode *string `field:"optional" json:"surroundExMode" yaml:"surroundExMode"`
	// When encoding 2/0 audio, sets whether Dolby Surround is matrix-encoded into the two channels.
	SurroundMode *string `field:"optional" json:"surroundMode" yaml:"surroundMode"`
}

// Settings for EBU-TT captions in the output.
//
// The parent of this entity is CaptionDestinationSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ebuTtDDestinationSettingsProperty := &ebuTtDDestinationSettingsProperty{
//   	copyrightHolder: jsii.String("copyrightHolder"),
//   	fillLineGap: jsii.String("fillLineGap"),
//   	fontFamily: jsii.String("fontFamily"),
//   	styleControl: jsii.String("styleControl"),
//   }
//
type CfnChannel_EbuTtDDestinationSettingsProperty struct {
	// Applies only if you plan to convert these source captions to EBU-TT-D or TTML in an output.
	//
	// Complete this field if you want to include the name of the copyright holder in the copyright metadata tag in the TTML.
	CopyrightHolder *string `field:"optional" json:"copyrightHolder" yaml:"copyrightHolder"`
	// Specifies how to handle the gap between the lines (in multi-line captions).
	//
	// - enabled: Fill with the captions background color (as specified in the input captions).
	// - disabled: Leave the gap unfilled.
	FillLineGap *string `field:"optional" json:"fillLineGap" yaml:"fillLineGap"`
	// Specifies the font family to include in the font data attached to the EBU-TT captions.
	//
	// Valid only if styleControl is set to include. If you leave this field empty, the font family is set to "monospaced". (If styleControl is set to exclude, the font family is always set to "monospaced".) You specify only the font family. All other style information (color, bold, position and so on) is copied from the input captions. The size is always set to 100% to allow the downstream player to choose the size. - Enter a list of font families, as a comma-separated list of font names, in order of preference. The name can be a font family (such as “Arial”), or a generic font family (such as “serif”), or “default” (to let the downstream player choose the font).
	// - Leave blank to set the family to “monospace”.
	FontFamily *string `field:"optional" json:"fontFamily" yaml:"fontFamily"`
	// Specifies the style information (font color, font position, and so on) to include in the font data that is attached to the EBU-TT captions.
	//
	// - include: Take the style information (font color, font position, and so on) from the source captions and include that information in the font data attached to the EBU-TT captions. This option is valid only if the source captions are Embedded or Teletext.
	// - exclude: In the font data attached to the EBU-TT captions, set the font family to "monospaced". Do not include any other style information.
	StyleControl *string `field:"optional" json:"styleControl" yaml:"styleControl"`
}

// The configuration of embedded captions in the output.
//
// The parent of this entity is CaptionDestinationSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   embeddedDestinationSettingsProperty := &embeddedDestinationSettingsProperty{
//   }
//
type CfnChannel_EmbeddedDestinationSettingsProperty struct {
}

// The settings for embedded plus SCTE-20 captions in the output.
//
// The parent of this entity is CaptionDestinationSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   embeddedPlusScte20DestinationSettingsProperty := &embeddedPlusScte20DestinationSettingsProperty{
//   }
//
type CfnChannel_EmbeddedPlusScte20DestinationSettingsProperty struct {
}

// Information about the embedded captions to extract from the input.
//
// The parent of this entity is CaptionSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   embeddedSourceSettingsProperty := &embeddedSourceSettingsProperty{
//   	convert608To708: jsii.String("convert608To708"),
//   	scte20Detection: jsii.String("scte20Detection"),
//   	source608ChannelNumber: jsii.Number(123),
//   	source608TrackNumber: jsii.Number(123),
//   }
//
type CfnChannel_EmbeddedSourceSettingsProperty struct {
	// If this is upconvert, 608 data is both passed through the "608 compatibility bytes" fields of the 708 wrapper as well as translated into 708.
	//
	// If 708 data is present in the source content, it is discarded.
	Convert608To708 *string `field:"optional" json:"convert608To708" yaml:"convert608To708"`
	// Set to "auto" to handle streams with intermittent or non-aligned SCTE-20 and embedded captions.
	Scte20Detection *string `field:"optional" json:"scte20Detection" yaml:"scte20Detection"`
	// Specifies the 608/708 channel number within the video track from which to extract captions.
	//
	// This is unused for passthrough.
	Source608ChannelNumber *float64 `field:"optional" json:"source608ChannelNumber" yaml:"source608ChannelNumber"`
	// This field is unused and deprecated.
	Source608TrackNumber *float64 `field:"optional" json:"source608TrackNumber" yaml:"source608TrackNumber"`
}

// The settings for the encoding of outputs.
//
// This entity is at the top level in the channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encoderSettingsProperty := &encoderSettingsProperty{
//   	audioDescriptions: []interface{}{
//   		&audioDescriptionProperty{
//   			audioNormalizationSettings: &audioNormalizationSettingsProperty{
//   				algorithm: jsii.String("algorithm"),
//   				algorithmControl: jsii.String("algorithmControl"),
//   				targetLkfs: jsii.Number(123),
//   			},
//   			audioSelectorName: jsii.String("audioSelectorName"),
//   			audioType: jsii.String("audioType"),
//   			audioTypeControl: jsii.String("audioTypeControl"),
//   			audioWatermarkingSettings: &audioWatermarkSettingsProperty{
//   				nielsenWatermarksSettings: &nielsenWatermarksSettingsProperty{
//   					nielsenCbetSettings: &nielsenCBETProperty{
//   						cbetCheckDigitString: jsii.String("cbetCheckDigitString"),
//   						cbetStepaside: jsii.String("cbetStepaside"),
//   						csid: jsii.String("csid"),
//   					},
//   					nielsenDistributionType: jsii.String("nielsenDistributionType"),
//   					nielsenNaesIiNwSettings: &nielsenNaesIiNwProperty{
//   						checkDigitString: jsii.String("checkDigitString"),
//   						sid: jsii.Number(123),
//   					},
//   				},
//   			},
//   			codecSettings: &audioCodecSettingsProperty{
//   				aacSettings: &aacSettingsProperty{
//   					bitrate: jsii.Number(123),
//   					codingMode: jsii.String("codingMode"),
//   					inputType: jsii.String("inputType"),
//   					profile: jsii.String("profile"),
//   					rateControlMode: jsii.String("rateControlMode"),
//   					rawFormat: jsii.String("rawFormat"),
//   					sampleRate: jsii.Number(123),
//   					spec: jsii.String("spec"),
//   					vbrQuality: jsii.String("vbrQuality"),
//   				},
//   				ac3Settings: &ac3SettingsProperty{
//   					bitrate: jsii.Number(123),
//   					bitstreamMode: jsii.String("bitstreamMode"),
//   					codingMode: jsii.String("codingMode"),
//   					dialnorm: jsii.Number(123),
//   					drcProfile: jsii.String("drcProfile"),
//   					lfeFilter: jsii.String("lfeFilter"),
//   					metadataControl: jsii.String("metadataControl"),
//   				},
//   				eac3Settings: &eac3SettingsProperty{
//   					attenuationControl: jsii.String("attenuationControl"),
//   					bitrate: jsii.Number(123),
//   					bitstreamMode: jsii.String("bitstreamMode"),
//   					codingMode: jsii.String("codingMode"),
//   					dcFilter: jsii.String("dcFilter"),
//   					dialnorm: jsii.Number(123),
//   					drcLine: jsii.String("drcLine"),
//   					drcRf: jsii.String("drcRf"),
//   					lfeControl: jsii.String("lfeControl"),
//   					lfeFilter: jsii.String("lfeFilter"),
//   					loRoCenterMixLevel: jsii.Number(123),
//   					loRoSurroundMixLevel: jsii.Number(123),
//   					ltRtCenterMixLevel: jsii.Number(123),
//   					ltRtSurroundMixLevel: jsii.Number(123),
//   					metadataControl: jsii.String("metadataControl"),
//   					passthroughControl: jsii.String("passthroughControl"),
//   					phaseControl: jsii.String("phaseControl"),
//   					stereoDownmix: jsii.String("stereoDownmix"),
//   					surroundExMode: jsii.String("surroundExMode"),
//   					surroundMode: jsii.String("surroundMode"),
//   				},
//   				mp2Settings: &mp2SettingsProperty{
//   					bitrate: jsii.Number(123),
//   					codingMode: jsii.String("codingMode"),
//   					sampleRate: jsii.Number(123),
//   				},
//   				passThroughSettings: &passThroughSettingsProperty{
//   				},
//   				wavSettings: &wavSettingsProperty{
//   					bitDepth: jsii.Number(123),
//   					codingMode: jsii.String("codingMode"),
//   					sampleRate: jsii.Number(123),
//   				},
//   			},
//   			languageCode: jsii.String("languageCode"),
//   			languageCodeControl: jsii.String("languageCodeControl"),
//   			name: jsii.String("name"),
//   			remixSettings: &remixSettingsProperty{
//   				channelMappings: []interface{}{
//   					&audioChannelMappingProperty{
//   						inputChannelLevels: []interface{}{
//   							&inputChannelLevelProperty{
//   								gain: jsii.Number(123),
//   								inputChannel: jsii.Number(123),
//   							},
//   						},
//   						outputChannel: jsii.Number(123),
//   					},
//   				},
//   				channelsIn: jsii.Number(123),
//   				channelsOut: jsii.Number(123),
//   			},
//   			streamName: jsii.String("streamName"),
//   		},
//   	},
//   	availBlanking: &availBlankingProperty{
//   		availBlankingImage: &inputLocationProperty{
//   			passwordParam: jsii.String("passwordParam"),
//   			uri: jsii.String("uri"),
//   			username: jsii.String("username"),
//   		},
//   		state: jsii.String("state"),
//   	},
//   	availConfiguration: &availConfigurationProperty{
//   		availSettings: &availSettingsProperty{
//   			scte35SpliceInsert: &scte35SpliceInsertProperty{
//   				adAvailOffset: jsii.Number(123),
//   				noRegionalBlackoutFlag: jsii.String("noRegionalBlackoutFlag"),
//   				webDeliveryAllowedFlag: jsii.String("webDeliveryAllowedFlag"),
//   			},
//   			scte35TimeSignalApos: &scte35TimeSignalAposProperty{
//   				adAvailOffset: jsii.Number(123),
//   				noRegionalBlackoutFlag: jsii.String("noRegionalBlackoutFlag"),
//   				webDeliveryAllowedFlag: jsii.String("webDeliveryAllowedFlag"),
//   			},
//   		},
//   	},
//   	blackoutSlate: &blackoutSlateProperty{
//   		blackoutSlateImage: &inputLocationProperty{
//   			passwordParam: jsii.String("passwordParam"),
//   			uri: jsii.String("uri"),
//   			username: jsii.String("username"),
//   		},
//   		networkEndBlackout: jsii.String("networkEndBlackout"),
//   		networkEndBlackoutImage: &inputLocationProperty{
//   			passwordParam: jsii.String("passwordParam"),
//   			uri: jsii.String("uri"),
//   			username: jsii.String("username"),
//   		},
//   		networkId: jsii.String("networkId"),
//   		state: jsii.String("state"),
//   	},
//   	captionDescriptions: []interface{}{
//   		&captionDescriptionProperty{
//   			captionSelectorName: jsii.String("captionSelectorName"),
//   			destinationSettings: &captionDestinationSettingsProperty{
//   				aribDestinationSettings: &aribDestinationSettingsProperty{
//   				},
//   				burnInDestinationSettings: &burnInDestinationSettingsProperty{
//   					alignment: jsii.String("alignment"),
//   					backgroundColor: jsii.String("backgroundColor"),
//   					backgroundOpacity: jsii.Number(123),
//   					font: &inputLocationProperty{
//   						passwordParam: jsii.String("passwordParam"),
//   						uri: jsii.String("uri"),
//   						username: jsii.String("username"),
//   					},
//   					fontColor: jsii.String("fontColor"),
//   					fontOpacity: jsii.Number(123),
//   					fontResolution: jsii.Number(123),
//   					fontSize: jsii.String("fontSize"),
//   					outlineColor: jsii.String("outlineColor"),
//   					outlineSize: jsii.Number(123),
//   					shadowColor: jsii.String("shadowColor"),
//   					shadowOpacity: jsii.Number(123),
//   					shadowXOffset: jsii.Number(123),
//   					shadowYOffset: jsii.Number(123),
//   					teletextGridControl: jsii.String("teletextGridControl"),
//   					xPosition: jsii.Number(123),
//   					yPosition: jsii.Number(123),
//   				},
//   				dvbSubDestinationSettings: &dvbSubDestinationSettingsProperty{
//   					alignment: jsii.String("alignment"),
//   					backgroundColor: jsii.String("backgroundColor"),
//   					backgroundOpacity: jsii.Number(123),
//   					font: &inputLocationProperty{
//   						passwordParam: jsii.String("passwordParam"),
//   						uri: jsii.String("uri"),
//   						username: jsii.String("username"),
//   					},
//   					fontColor: jsii.String("fontColor"),
//   					fontOpacity: jsii.Number(123),
//   					fontResolution: jsii.Number(123),
//   					fontSize: jsii.String("fontSize"),
//   					outlineColor: jsii.String("outlineColor"),
//   					outlineSize: jsii.Number(123),
//   					shadowColor: jsii.String("shadowColor"),
//   					shadowOpacity: jsii.Number(123),
//   					shadowXOffset: jsii.Number(123),
//   					shadowYOffset: jsii.Number(123),
//   					teletextGridControl: jsii.String("teletextGridControl"),
//   					xPosition: jsii.Number(123),
//   					yPosition: jsii.Number(123),
//   				},
//   				ebuTtDDestinationSettings: &ebuTtDDestinationSettingsProperty{
//   					copyrightHolder: jsii.String("copyrightHolder"),
//   					fillLineGap: jsii.String("fillLineGap"),
//   					fontFamily: jsii.String("fontFamily"),
//   					styleControl: jsii.String("styleControl"),
//   				},
//   				embeddedDestinationSettings: &embeddedDestinationSettingsProperty{
//   				},
//   				embeddedPlusScte20DestinationSettings: &embeddedPlusScte20DestinationSettingsProperty{
//   				},
//   				rtmpCaptionInfoDestinationSettings: &rtmpCaptionInfoDestinationSettingsProperty{
//   				},
//   				scte20PlusEmbeddedDestinationSettings: &scte20PlusEmbeddedDestinationSettingsProperty{
//   				},
//   				scte27DestinationSettings: &scte27DestinationSettingsProperty{
//   				},
//   				smpteTtDestinationSettings: &smpteTtDestinationSettingsProperty{
//   				},
//   				teletextDestinationSettings: &teletextDestinationSettingsProperty{
//   				},
//   				ttmlDestinationSettings: &ttmlDestinationSettingsProperty{
//   					styleControl: jsii.String("styleControl"),
//   				},
//   				webvttDestinationSettings: &webvttDestinationSettingsProperty{
//   					styleControl: jsii.String("styleControl"),
//   				},
//   			},
//   			languageCode: jsii.String("languageCode"),
//   			languageDescription: jsii.String("languageDescription"),
//   			name: jsii.String("name"),
//   		},
//   	},
//   	featureActivations: &featureActivationsProperty{
//   		inputPrepareScheduleActions: jsii.String("inputPrepareScheduleActions"),
//   	},
//   	globalConfiguration: &globalConfigurationProperty{
//   		initialAudioGain: jsii.Number(123),
//   		inputEndAction: jsii.String("inputEndAction"),
//   		inputLossBehavior: &inputLossBehaviorProperty{
//   			blackFrameMsec: jsii.Number(123),
//   			inputLossImageColor: jsii.String("inputLossImageColor"),
//   			inputLossImageSlate: &inputLocationProperty{
//   				passwordParam: jsii.String("passwordParam"),
//   				uri: jsii.String("uri"),
//   				username: jsii.String("username"),
//   			},
//   			inputLossImageType: jsii.String("inputLossImageType"),
//   			repeatFrameMsec: jsii.Number(123),
//   		},
//   		outputLockingMode: jsii.String("outputLockingMode"),
//   		outputTimingSource: jsii.String("outputTimingSource"),
//   		supportLowFramerateInputs: jsii.String("supportLowFramerateInputs"),
//   	},
//   	motionGraphicsConfiguration: &motionGraphicsConfigurationProperty{
//   		motionGraphicsInsertion: jsii.String("motionGraphicsInsertion"),
//   		motionGraphicsSettings: &motionGraphicsSettingsProperty{
//   			htmlMotionGraphicsSettings: &htmlMotionGraphicsSettingsProperty{
//   			},
//   		},
//   	},
//   	nielsenConfiguration: &nielsenConfigurationProperty{
//   		distributorId: jsii.String("distributorId"),
//   		nielsenPcmToId3Tagging: jsii.String("nielsenPcmToId3Tagging"),
//   	},
//   	outputGroups: []interface{}{
//   		&outputGroupProperty{
//   			name: jsii.String("name"),
//   			outputGroupSettings: &outputGroupSettingsProperty{
//   				archiveGroupSettings: &archiveGroupSettingsProperty{
//   					archiveCdnSettings: &archiveCdnSettingsProperty{
//   						archiveS3Settings: &archiveS3SettingsProperty{
//   							cannedAcl: jsii.String("cannedAcl"),
//   						},
//   					},
//   					destination: &outputLocationRefProperty{
//   						destinationRefId: jsii.String("destinationRefId"),
//   					},
//   					rolloverInterval: jsii.Number(123),
//   				},
//   				frameCaptureGroupSettings: &frameCaptureGroupSettingsProperty{
//   					destination: &outputLocationRefProperty{
//   						destinationRefId: jsii.String("destinationRefId"),
//   					},
//   					frameCaptureCdnSettings: &frameCaptureCdnSettingsProperty{
//   						frameCaptureS3Settings: &frameCaptureS3SettingsProperty{
//   							cannedAcl: jsii.String("cannedAcl"),
//   						},
//   					},
//   				},
//   				hlsGroupSettings: &hlsGroupSettingsProperty{
//   					adMarkers: []*string{
//   						jsii.String("adMarkers"),
//   					},
//   					baseUrlContent: jsii.String("baseUrlContent"),
//   					baseUrlContent1: jsii.String("baseUrlContent1"),
//   					baseUrlManifest: jsii.String("baseUrlManifest"),
//   					baseUrlManifest1: jsii.String("baseUrlManifest1"),
//   					captionLanguageMappings: []interface{}{
//   						&captionLanguageMappingProperty{
//   							captionChannel: jsii.Number(123),
//   							languageCode: jsii.String("languageCode"),
//   							languageDescription: jsii.String("languageDescription"),
//   						},
//   					},
//   					captionLanguageSetting: jsii.String("captionLanguageSetting"),
//   					clientCache: jsii.String("clientCache"),
//   					codecSpecification: jsii.String("codecSpecification"),
//   					constantIv: jsii.String("constantIv"),
//   					destination: &outputLocationRefProperty{
//   						destinationRefId: jsii.String("destinationRefId"),
//   					},
//   					directoryStructure: jsii.String("directoryStructure"),
//   					discontinuityTags: jsii.String("discontinuityTags"),
//   					encryptionType: jsii.String("encryptionType"),
//   					hlsCdnSettings: &hlsCdnSettingsProperty{
//   						hlsAkamaiSettings: &hlsAkamaiSettingsProperty{
//   							connectionRetryInterval: jsii.Number(123),
//   							filecacheDuration: jsii.Number(123),
//   							httpTransferMode: jsii.String("httpTransferMode"),
//   							numRetries: jsii.Number(123),
//   							restartDelay: jsii.Number(123),
//   							salt: jsii.String("salt"),
//   							token: jsii.String("token"),
//   						},
//   						hlsBasicPutSettings: &hlsBasicPutSettingsProperty{
//   							connectionRetryInterval: jsii.Number(123),
//   							filecacheDuration: jsii.Number(123),
//   							numRetries: jsii.Number(123),
//   							restartDelay: jsii.Number(123),
//   						},
//   						hlsMediaStoreSettings: &hlsMediaStoreSettingsProperty{
//   							connectionRetryInterval: jsii.Number(123),
//   							filecacheDuration: jsii.Number(123),
//   							mediaStoreStorageClass: jsii.String("mediaStoreStorageClass"),
//   							numRetries: jsii.Number(123),
//   							restartDelay: jsii.Number(123),
//   						},
//   						hlsS3Settings: &hlsS3SettingsProperty{
//   							cannedAcl: jsii.String("cannedAcl"),
//   						},
//   						hlsWebdavSettings: &hlsWebdavSettingsProperty{
//   							connectionRetryInterval: jsii.Number(123),
//   							filecacheDuration: jsii.Number(123),
//   							httpTransferMode: jsii.String("httpTransferMode"),
//   							numRetries: jsii.Number(123),
//   							restartDelay: jsii.Number(123),
//   						},
//   					},
//   					hlsId3SegmentTagging: jsii.String("hlsId3SegmentTagging"),
//   					iFrameOnlyPlaylists: jsii.String("iFrameOnlyPlaylists"),
//   					incompleteSegmentBehavior: jsii.String("incompleteSegmentBehavior"),
//   					indexNSegments: jsii.Number(123),
//   					inputLossAction: jsii.String("inputLossAction"),
//   					ivInManifest: jsii.String("ivInManifest"),
//   					ivSource: jsii.String("ivSource"),
//   					keepSegments: jsii.Number(123),
//   					keyFormat: jsii.String("keyFormat"),
//   					keyFormatVersions: jsii.String("keyFormatVersions"),
//   					keyProviderSettings: &keyProviderSettingsProperty{
//   						staticKeySettings: &staticKeySettingsProperty{
//   							keyProviderServer: &inputLocationProperty{
//   								passwordParam: jsii.String("passwordParam"),
//   								uri: jsii.String("uri"),
//   								username: jsii.String("username"),
//   							},
//   							staticKeyValue: jsii.String("staticKeyValue"),
//   						},
//   					},
//   					manifestCompression: jsii.String("manifestCompression"),
//   					manifestDurationFormat: jsii.String("manifestDurationFormat"),
//   					minSegmentLength: jsii.Number(123),
//   					mode: jsii.String("mode"),
//   					outputSelection: jsii.String("outputSelection"),
//   					programDateTime: jsii.String("programDateTime"),
//   					programDateTimeClock: jsii.String("programDateTimeClock"),
//   					programDateTimePeriod: jsii.Number(123),
//   					redundantManifest: jsii.String("redundantManifest"),
//   					segmentationMode: jsii.String("segmentationMode"),
//   					segmentLength: jsii.Number(123),
//   					segmentsPerSubdirectory: jsii.Number(123),
//   					streamInfResolution: jsii.String("streamInfResolution"),
//   					timedMetadataId3Frame: jsii.String("timedMetadataId3Frame"),
//   					timedMetadataId3Period: jsii.Number(123),
//   					timestampDeltaMilliseconds: jsii.Number(123),
//   					tsFileMode: jsii.String("tsFileMode"),
//   				},
//   				mediaPackageGroupSettings: &mediaPackageGroupSettingsProperty{
//   					destination: &outputLocationRefProperty{
//   						destinationRefId: jsii.String("destinationRefId"),
//   					},
//   				},
//   				msSmoothGroupSettings: &msSmoothGroupSettingsProperty{
//   					acquisitionPointId: jsii.String("acquisitionPointId"),
//   					audioOnlyTimecodeControl: jsii.String("audioOnlyTimecodeControl"),
//   					certificateMode: jsii.String("certificateMode"),
//   					connectionRetryInterval: jsii.Number(123),
//   					destination: &outputLocationRefProperty{
//   						destinationRefId: jsii.String("destinationRefId"),
//   					},
//   					eventId: jsii.String("eventId"),
//   					eventIdMode: jsii.String("eventIdMode"),
//   					eventStopBehavior: jsii.String("eventStopBehavior"),
//   					filecacheDuration: jsii.Number(123),
//   					fragmentLength: jsii.Number(123),
//   					inputLossAction: jsii.String("inputLossAction"),
//   					numRetries: jsii.Number(123),
//   					restartDelay: jsii.Number(123),
//   					segmentationMode: jsii.String("segmentationMode"),
//   					sendDelayMs: jsii.Number(123),
//   					sparseTrackType: jsii.String("sparseTrackType"),
//   					streamManifestBehavior: jsii.String("streamManifestBehavior"),
//   					timestampOffset: jsii.String("timestampOffset"),
//   					timestampOffsetMode: jsii.String("timestampOffsetMode"),
//   				},
//   				multiplexGroupSettings: &multiplexGroupSettingsProperty{
//   				},
//   				rtmpGroupSettings: &rtmpGroupSettingsProperty{
//   					adMarkers: []*string{
//   						jsii.String("adMarkers"),
//   					},
//   					authenticationScheme: jsii.String("authenticationScheme"),
//   					cacheFullBehavior: jsii.String("cacheFullBehavior"),
//   					cacheLength: jsii.Number(123),
//   					captionData: jsii.String("captionData"),
//   					inputLossAction: jsii.String("inputLossAction"),
//   					restartDelay: jsii.Number(123),
//   				},
//   				udpGroupSettings: &udpGroupSettingsProperty{
//   					inputLossAction: jsii.String("inputLossAction"),
//   					timedMetadataId3Frame: jsii.String("timedMetadataId3Frame"),
//   					timedMetadataId3Period: jsii.Number(123),
//   				},
//   			},
//   			outputs: []interface{}{
//   				&outputProperty{
//   					audioDescriptionNames: []*string{
//   						jsii.String("audioDescriptionNames"),
//   					},
//   					captionDescriptionNames: []*string{
//   						jsii.String("captionDescriptionNames"),
//   					},
//   					outputName: jsii.String("outputName"),
//   					outputSettings: &outputSettingsProperty{
//   						archiveOutputSettings: &archiveOutputSettingsProperty{
//   							containerSettings: &archiveContainerSettingsProperty{
//   								m2TsSettings: &m2tsSettingsProperty{
//   									absentInputAudioBehavior: jsii.String("absentInputAudioBehavior"),
//   									arib: jsii.String("arib"),
//   									aribCaptionsPid: jsii.String("aribCaptionsPid"),
//   									aribCaptionsPidControl: jsii.String("aribCaptionsPidControl"),
//   									audioBufferModel: jsii.String("audioBufferModel"),
//   									audioFramesPerPes: jsii.Number(123),
//   									audioPids: jsii.String("audioPids"),
//   									audioStreamType: jsii.String("audioStreamType"),
//   									bitrate: jsii.Number(123),
//   									bufferModel: jsii.String("bufferModel"),
//   									ccDescriptor: jsii.String("ccDescriptor"),
//   									dvbNitSettings: &dvbNitSettingsProperty{
//   										networkId: jsii.Number(123),
//   										networkName: jsii.String("networkName"),
//   										repInterval: jsii.Number(123),
//   									},
//   									dvbSdtSettings: &dvbSdtSettingsProperty{
//   										outputSdt: jsii.String("outputSdt"),
//   										repInterval: jsii.Number(123),
//   										serviceName: jsii.String("serviceName"),
//   										serviceProviderName: jsii.String("serviceProviderName"),
//   									},
//   									dvbSubPids: jsii.String("dvbSubPids"),
//   									dvbTdtSettings: &dvbTdtSettingsProperty{
//   										repInterval: jsii.Number(123),
//   									},
//   									dvbTeletextPid: jsii.String("dvbTeletextPid"),
//   									ebif: jsii.String("ebif"),
//   									ebpAudioInterval: jsii.String("ebpAudioInterval"),
//   									ebpLookaheadMs: jsii.Number(123),
//   									ebpPlacement: jsii.String("ebpPlacement"),
//   									ecmPid: jsii.String("ecmPid"),
//   									esRateInPes: jsii.String("esRateInPes"),
//   									etvPlatformPid: jsii.String("etvPlatformPid"),
//   									etvSignalPid: jsii.String("etvSignalPid"),
//   									fragmentTime: jsii.Number(123),
//   									klv: jsii.String("klv"),
//   									klvDataPids: jsii.String("klvDataPids"),
//   									nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   									nullPacketBitrate: jsii.Number(123),
//   									patInterval: jsii.Number(123),
//   									pcrControl: jsii.String("pcrControl"),
//   									pcrPeriod: jsii.Number(123),
//   									pcrPid: jsii.String("pcrPid"),
//   									pmtInterval: jsii.Number(123),
//   									pmtPid: jsii.String("pmtPid"),
//   									programNum: jsii.Number(123),
//   									rateMode: jsii.String("rateMode"),
//   									scte27Pids: jsii.String("scte27Pids"),
//   									scte35Control: jsii.String("scte35Control"),
//   									scte35Pid: jsii.String("scte35Pid"),
//   									segmentationMarkers: jsii.String("segmentationMarkers"),
//   									segmentationStyle: jsii.String("segmentationStyle"),
//   									segmentationTime: jsii.Number(123),
//   									timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   									timedMetadataPid: jsii.String("timedMetadataPid"),
//   									transportStreamId: jsii.Number(123),
//   									videoPid: jsii.String("videoPid"),
//   								},
//   								rawSettings: &rawSettingsProperty{
//   								},
//   							},
//   							extension: jsii.String("extension"),
//   							nameModifier: jsii.String("nameModifier"),
//   						},
//   						frameCaptureOutputSettings: &frameCaptureOutputSettingsProperty{
//   							nameModifier: jsii.String("nameModifier"),
//   						},
//   						hlsOutputSettings: &hlsOutputSettingsProperty{
//   							h265PackagingType: jsii.String("h265PackagingType"),
//   							hlsSettings: &hlsSettingsProperty{
//   								audioOnlyHlsSettings: &audioOnlyHlsSettingsProperty{
//   									audioGroupId: jsii.String("audioGroupId"),
//   									audioOnlyImage: &inputLocationProperty{
//   										passwordParam: jsii.String("passwordParam"),
//   										uri: jsii.String("uri"),
//   										username: jsii.String("username"),
//   									},
//   									audioTrackType: jsii.String("audioTrackType"),
//   									segmentType: jsii.String("segmentType"),
//   								},
//   								fmp4HlsSettings: &fmp4HlsSettingsProperty{
//   									audioRenditionSets: jsii.String("audioRenditionSets"),
//   									nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   									timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   								},
//   								frameCaptureHlsSettings: &frameCaptureHlsSettingsProperty{
//   								},
//   								standardHlsSettings: &standardHlsSettingsProperty{
//   									audioRenditionSets: jsii.String("audioRenditionSets"),
//   									m3U8Settings: &m3u8SettingsProperty{
//   										audioFramesPerPes: jsii.Number(123),
//   										audioPids: jsii.String("audioPids"),
//   										ecmPid: jsii.String("ecmPid"),
//   										nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   										patInterval: jsii.Number(123),
//   										pcrControl: jsii.String("pcrControl"),
//   										pcrPeriod: jsii.Number(123),
//   										pcrPid: jsii.String("pcrPid"),
//   										pmtInterval: jsii.Number(123),
//   										pmtPid: jsii.String("pmtPid"),
//   										programNum: jsii.Number(123),
//   										scte35Behavior: jsii.String("scte35Behavior"),
//   										scte35Pid: jsii.String("scte35Pid"),
//   										timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   										timedMetadataPid: jsii.String("timedMetadataPid"),
//   										transportStreamId: jsii.Number(123),
//   										videoPid: jsii.String("videoPid"),
//   									},
//   								},
//   							},
//   							nameModifier: jsii.String("nameModifier"),
//   							segmentModifier: jsii.String("segmentModifier"),
//   						},
//   						mediaPackageOutputSettings: &mediaPackageOutputSettingsProperty{
//   						},
//   						msSmoothOutputSettings: &msSmoothOutputSettingsProperty{
//   							h265PackagingType: jsii.String("h265PackagingType"),
//   							nameModifier: jsii.String("nameModifier"),
//   						},
//   						multiplexOutputSettings: &multiplexOutputSettingsProperty{
//   							destination: &outputLocationRefProperty{
//   								destinationRefId: jsii.String("destinationRefId"),
//   							},
//   						},
//   						rtmpOutputSettings: &rtmpOutputSettingsProperty{
//   							certificateMode: jsii.String("certificateMode"),
//   							connectionRetryInterval: jsii.Number(123),
//   							destination: &outputLocationRefProperty{
//   								destinationRefId: jsii.String("destinationRefId"),
//   							},
//   							numRetries: jsii.Number(123),
//   						},
//   						udpOutputSettings: &udpOutputSettingsProperty{
//   							bufferMsec: jsii.Number(123),
//   							containerSettings: &udpContainerSettingsProperty{
//   								m2TsSettings: &m2tsSettingsProperty{
//   									absentInputAudioBehavior: jsii.String("absentInputAudioBehavior"),
//   									arib: jsii.String("arib"),
//   									aribCaptionsPid: jsii.String("aribCaptionsPid"),
//   									aribCaptionsPidControl: jsii.String("aribCaptionsPidControl"),
//   									audioBufferModel: jsii.String("audioBufferModel"),
//   									audioFramesPerPes: jsii.Number(123),
//   									audioPids: jsii.String("audioPids"),
//   									audioStreamType: jsii.String("audioStreamType"),
//   									bitrate: jsii.Number(123),
//   									bufferModel: jsii.String("bufferModel"),
//   									ccDescriptor: jsii.String("ccDescriptor"),
//   									dvbNitSettings: &dvbNitSettingsProperty{
//   										networkId: jsii.Number(123),
//   										networkName: jsii.String("networkName"),
//   										repInterval: jsii.Number(123),
//   									},
//   									dvbSdtSettings: &dvbSdtSettingsProperty{
//   										outputSdt: jsii.String("outputSdt"),
//   										repInterval: jsii.Number(123),
//   										serviceName: jsii.String("serviceName"),
//   										serviceProviderName: jsii.String("serviceProviderName"),
//   									},
//   									dvbSubPids: jsii.String("dvbSubPids"),
//   									dvbTdtSettings: &dvbTdtSettingsProperty{
//   										repInterval: jsii.Number(123),
//   									},
//   									dvbTeletextPid: jsii.String("dvbTeletextPid"),
//   									ebif: jsii.String("ebif"),
//   									ebpAudioInterval: jsii.String("ebpAudioInterval"),
//   									ebpLookaheadMs: jsii.Number(123),
//   									ebpPlacement: jsii.String("ebpPlacement"),
//   									ecmPid: jsii.String("ecmPid"),
//   									esRateInPes: jsii.String("esRateInPes"),
//   									etvPlatformPid: jsii.String("etvPlatformPid"),
//   									etvSignalPid: jsii.String("etvSignalPid"),
//   									fragmentTime: jsii.Number(123),
//   									klv: jsii.String("klv"),
//   									klvDataPids: jsii.String("klvDataPids"),
//   									nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   									nullPacketBitrate: jsii.Number(123),
//   									patInterval: jsii.Number(123),
//   									pcrControl: jsii.String("pcrControl"),
//   									pcrPeriod: jsii.Number(123),
//   									pcrPid: jsii.String("pcrPid"),
//   									pmtInterval: jsii.Number(123),
//   									pmtPid: jsii.String("pmtPid"),
//   									programNum: jsii.Number(123),
//   									rateMode: jsii.String("rateMode"),
//   									scte27Pids: jsii.String("scte27Pids"),
//   									scte35Control: jsii.String("scte35Control"),
//   									scte35Pid: jsii.String("scte35Pid"),
//   									segmentationMarkers: jsii.String("segmentationMarkers"),
//   									segmentationStyle: jsii.String("segmentationStyle"),
//   									segmentationTime: jsii.Number(123),
//   									timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   									timedMetadataPid: jsii.String("timedMetadataPid"),
//   									transportStreamId: jsii.Number(123),
//   									videoPid: jsii.String("videoPid"),
//   								},
//   							},
//   							destination: &outputLocationRefProperty{
//   								destinationRefId: jsii.String("destinationRefId"),
//   							},
//   							fecOutputSettings: &fecOutputSettingsProperty{
//   								columnDepth: jsii.Number(123),
//   								includeFec: jsii.String("includeFec"),
//   								rowLength: jsii.Number(123),
//   							},
//   						},
//   					},
//   					videoDescriptionName: jsii.String("videoDescriptionName"),
//   				},
//   			},
//   		},
//   	},
//   	timecodeConfig: &timecodeConfigProperty{
//   		source: jsii.String("source"),
//   		syncThreshold: jsii.Number(123),
//   	},
//   	videoDescriptions: []interface{}{
//   		&videoDescriptionProperty{
//   			codecSettings: &videoCodecSettingsProperty{
//   				frameCaptureSettings: &frameCaptureSettingsProperty{
//   					captureInterval: jsii.Number(123),
//   					captureIntervalUnits: jsii.String("captureIntervalUnits"),
//   				},
//   				h264Settings: &h264SettingsProperty{
//   					adaptiveQuantization: jsii.String("adaptiveQuantization"),
//   					afdSignaling: jsii.String("afdSignaling"),
//   					bitrate: jsii.Number(123),
//   					bufFillPct: jsii.Number(123),
//   					bufSize: jsii.Number(123),
//   					colorMetadata: jsii.String("colorMetadata"),
//   					colorSpaceSettings: &h264ColorSpaceSettingsProperty{
//   						colorSpacePassthroughSettings: &colorSpacePassthroughSettingsProperty{
//   						},
//   						rec601Settings: &rec601SettingsProperty{
//   						},
//   						rec709Settings: &rec709SettingsProperty{
//   						},
//   					},
//   					entropyEncoding: jsii.String("entropyEncoding"),
//   					filterSettings: &h264FilterSettingsProperty{
//   						temporalFilterSettings: &temporalFilterSettingsProperty{
//   							postFilterSharpening: jsii.String("postFilterSharpening"),
//   							strength: jsii.String("strength"),
//   						},
//   					},
//   					fixedAfd: jsii.String("fixedAfd"),
//   					flickerAq: jsii.String("flickerAq"),
//   					forceFieldPictures: jsii.String("forceFieldPictures"),
//   					framerateControl: jsii.String("framerateControl"),
//   					framerateDenominator: jsii.Number(123),
//   					framerateNumerator: jsii.Number(123),
//   					gopBReference: jsii.String("gopBReference"),
//   					gopClosedCadence: jsii.Number(123),
//   					gopNumBFrames: jsii.Number(123),
//   					gopSize: jsii.Number(123),
//   					gopSizeUnits: jsii.String("gopSizeUnits"),
//   					level: jsii.String("level"),
//   					lookAheadRateControl: jsii.String("lookAheadRateControl"),
//   					maxBitrate: jsii.Number(123),
//   					minIInterval: jsii.Number(123),
//   					numRefFrames: jsii.Number(123),
//   					parControl: jsii.String("parControl"),
//   					parDenominator: jsii.Number(123),
//   					parNumerator: jsii.Number(123),
//   					profile: jsii.String("profile"),
//   					qualityLevel: jsii.String("qualityLevel"),
//   					qvbrQualityLevel: jsii.Number(123),
//   					rateControlMode: jsii.String("rateControlMode"),
//   					scanType: jsii.String("scanType"),
//   					sceneChangeDetect: jsii.String("sceneChangeDetect"),
//   					slices: jsii.Number(123),
//   					softness: jsii.Number(123),
//   					spatialAq: jsii.String("spatialAq"),
//   					subgopLength: jsii.String("subgopLength"),
//   					syntax: jsii.String("syntax"),
//   					temporalAq: jsii.String("temporalAq"),
//   					timecodeInsertion: jsii.String("timecodeInsertion"),
//   				},
//   				h265Settings: &h265SettingsProperty{
//   					adaptiveQuantization: jsii.String("adaptiveQuantization"),
//   					afdSignaling: jsii.String("afdSignaling"),
//   					alternativeTransferFunction: jsii.String("alternativeTransferFunction"),
//   					bitrate: jsii.Number(123),
//   					bufSize: jsii.Number(123),
//   					colorMetadata: jsii.String("colorMetadata"),
//   					colorSpaceSettings: &h265ColorSpaceSettingsProperty{
//   						colorSpacePassthroughSettings: &colorSpacePassthroughSettingsProperty{
//   						},
//   						hdr10Settings: &hdr10SettingsProperty{
//   							maxCll: jsii.Number(123),
//   							maxFall: jsii.Number(123),
//   						},
//   						rec601Settings: &rec601SettingsProperty{
//   						},
//   						rec709Settings: &rec709SettingsProperty{
//   						},
//   					},
//   					filterSettings: &h265FilterSettingsProperty{
//   						temporalFilterSettings: &temporalFilterSettingsProperty{
//   							postFilterSharpening: jsii.String("postFilterSharpening"),
//   							strength: jsii.String("strength"),
//   						},
//   					},
//   					fixedAfd: jsii.String("fixedAfd"),
//   					flickerAq: jsii.String("flickerAq"),
//   					framerateDenominator: jsii.Number(123),
//   					framerateNumerator: jsii.Number(123),
//   					gopClosedCadence: jsii.Number(123),
//   					gopSize: jsii.Number(123),
//   					gopSizeUnits: jsii.String("gopSizeUnits"),
//   					level: jsii.String("level"),
//   					lookAheadRateControl: jsii.String("lookAheadRateControl"),
//   					maxBitrate: jsii.Number(123),
//   					minIInterval: jsii.Number(123),
//   					parDenominator: jsii.Number(123),
//   					parNumerator: jsii.Number(123),
//   					profile: jsii.String("profile"),
//   					qvbrQualityLevel: jsii.Number(123),
//   					rateControlMode: jsii.String("rateControlMode"),
//   					scanType: jsii.String("scanType"),
//   					sceneChangeDetect: jsii.String("sceneChangeDetect"),
//   					slices: jsii.Number(123),
//   					tier: jsii.String("tier"),
//   					timecodeInsertion: jsii.String("timecodeInsertion"),
//   				},
//   				mpeg2Settings: &mpeg2SettingsProperty{
//   					adaptiveQuantization: jsii.String("adaptiveQuantization"),
//   					afdSignaling: jsii.String("afdSignaling"),
//   					colorMetadata: jsii.String("colorMetadata"),
//   					colorSpace: jsii.String("colorSpace"),
//   					displayAspectRatio: jsii.String("displayAspectRatio"),
//   					filterSettings: &mpeg2FilterSettingsProperty{
//   						temporalFilterSettings: &temporalFilterSettingsProperty{
//   							postFilterSharpening: jsii.String("postFilterSharpening"),
//   							strength: jsii.String("strength"),
//   						},
//   					},
//   					fixedAfd: jsii.String("fixedAfd"),
//   					framerateDenominator: jsii.Number(123),
//   					framerateNumerator: jsii.Number(123),
//   					gopClosedCadence: jsii.Number(123),
//   					gopNumBFrames: jsii.Number(123),
//   					gopSize: jsii.Number(123),
//   					gopSizeUnits: jsii.String("gopSizeUnits"),
//   					scanType: jsii.String("scanType"),
//   					subgopLength: jsii.String("subgopLength"),
//   					timecodeInsertion: jsii.String("timecodeInsertion"),
//   				},
//   			},
//   			height: jsii.Number(123),
//   			name: jsii.String("name"),
//   			respondToAfd: jsii.String("respondToAfd"),
//   			scalingBehavior: jsii.String("scalingBehavior"),
//   			sharpness: jsii.Number(123),
//   			width: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnChannel_EncoderSettingsProperty struct {
	// The encoding information for output audio.
	AudioDescriptions interface{} `field:"optional" json:"audioDescriptions" yaml:"audioDescriptions"`
	// The settings for ad avail blanking.
	AvailBlanking interface{} `field:"optional" json:"availBlanking" yaml:"availBlanking"`
	// The configuration settings for the ad avail handling.
	AvailConfiguration interface{} `field:"optional" json:"availConfiguration" yaml:"availConfiguration"`
	// The settings for the blackout slate.
	BlackoutSlate interface{} `field:"optional" json:"blackoutSlate" yaml:"blackoutSlate"`
	// The encoding information for output captions.
	CaptionDescriptions interface{} `field:"optional" json:"captionDescriptions" yaml:"captionDescriptions"`
	// Settings to enable specific features.
	FeatureActivations interface{} `field:"optional" json:"featureActivations" yaml:"featureActivations"`
	// The configuration settings that apply to the entire channel.
	GlobalConfiguration interface{} `field:"optional" json:"globalConfiguration" yaml:"globalConfiguration"`
	// Settings to enable and configure the motion graphics overlay feature in the channel.
	MotionGraphicsConfiguration interface{} `field:"optional" json:"motionGraphicsConfiguration" yaml:"motionGraphicsConfiguration"`
	// The settings to configure Nielsen watermarks.
	NielsenConfiguration interface{} `field:"optional" json:"nielsenConfiguration" yaml:"nielsenConfiguration"`
	// The settings for the output groups in the channel.
	OutputGroups interface{} `field:"optional" json:"outputGroups" yaml:"outputGroups"`
	// Contains settings used to acquire and adjust timecode information from the inputs.
	TimecodeConfig interface{} `field:"optional" json:"timecodeConfig" yaml:"timecodeConfig"`
	// The encoding information for output videos.
	VideoDescriptions interface{} `field:"optional" json:"videoDescriptions" yaml:"videoDescriptions"`
}

// Failover Condition settings. There can be multiple failover conditions inside AutomaticInputFailoverSettings.
//
// The parent of this entity is AutomaticInputFailoverSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   failoverConditionProperty := &failoverConditionProperty{
//   	failoverConditionSettings: &failoverConditionSettingsProperty{
//   		audioSilenceSettings: &audioSilenceFailoverSettingsProperty{
//   			audioSelectorName: jsii.String("audioSelectorName"),
//   			audioSilenceThresholdMsec: jsii.Number(123),
//   		},
//   		inputLossSettings: &inputLossFailoverSettingsProperty{
//   			inputLossThresholdMsec: jsii.Number(123),
//   		},
//   		videoBlackSettings: &videoBlackFailoverSettingsProperty{
//   			blackDetectThreshold: jsii.Number(123),
//   			videoBlackThresholdMsec: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnChannel_FailoverConditionProperty struct {
	// Settings for a specific failover condition.
	FailoverConditionSettings interface{} `field:"optional" json:"failoverConditionSettings" yaml:"failoverConditionSettings"`
}

// Settings for one failover condition.
//
// The parent of this entity is FailoverCondition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   failoverConditionSettingsProperty := &failoverConditionSettingsProperty{
//   	audioSilenceSettings: &audioSilenceFailoverSettingsProperty{
//   		audioSelectorName: jsii.String("audioSelectorName"),
//   		audioSilenceThresholdMsec: jsii.Number(123),
//   	},
//   	inputLossSettings: &inputLossFailoverSettingsProperty{
//   		inputLossThresholdMsec: jsii.Number(123),
//   	},
//   	videoBlackSettings: &videoBlackFailoverSettingsProperty{
//   		blackDetectThreshold: jsii.Number(123),
//   		videoBlackThresholdMsec: jsii.Number(123),
//   	},
//   }
//
type CfnChannel_FailoverConditionSettingsProperty struct {
	// MediaLive will perform a failover if the specified audio selector is silent for the specified period.
	AudioSilenceSettings interface{} `field:"optional" json:"audioSilenceSettings" yaml:"audioSilenceSettings"`
	// MediaLive will perform a failover if content is not detected in this input for the specified period.
	InputLossSettings interface{} `field:"optional" json:"inputLossSettings" yaml:"inputLossSettings"`
	// MediaLive will perform a failover if content is considered black for the specified period.
	VideoBlackSettings interface{} `field:"optional" json:"videoBlackSettings" yaml:"videoBlackSettings"`
}

// Settings to enable specific features. You can't configure these features until you have enabled them in the channel.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   featureActivationsProperty := &featureActivationsProperty{
//   	inputPrepareScheduleActions: jsii.String("inputPrepareScheduleActions"),
//   }
//
type CfnChannel_FeatureActivationsProperty struct {
	// Enables the Input Prepare feature.
	//
	// You can create Input Prepare actions in the schedule only if this feature is enabled.
	// If you disable the feature on an existing schedule, make sure that you first delete all input prepare actions from the schedule.
	InputPrepareScheduleActions *string `field:"optional" json:"inputPrepareScheduleActions" yaml:"inputPrepareScheduleActions"`
}

// The settings for FEC.
//
// The parent of this entity is UdpOutputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fecOutputSettingsProperty := &fecOutputSettingsProperty{
//   	columnDepth: jsii.Number(123),
//   	includeFec: jsii.String("includeFec"),
//   	rowLength: jsii.Number(123),
//   }
//
type CfnChannel_FecOutputSettingsProperty struct {
	// The parameter D from SMPTE 2022-1.
	//
	// The height of the FEC protection matrix. The number of transport stream packets per column error correction packet. The number must be between 4 and 20, inclusive.
	ColumnDepth *float64 `field:"optional" json:"columnDepth" yaml:"columnDepth"`
	// Enables column only or column and row-based FEC.
	IncludeFec *string `field:"optional" json:"includeFec" yaml:"includeFec"`
	// The parameter L from SMPTE 2022-1.
	//
	// The width of the FEC protection matrix. Must be between 1 and 20, inclusive. If only Column FEC is used, then larger values increase robustness. If Row FEC is used, then this is the number of transport stream packets per row error correction packet, and the value must be between 4 and 20, inclusive, if includeFec is columnAndRow. If includeFec is column, this value must be 1 to 20, inclusive.
	RowLength *float64 `field:"optional" json:"rowLength" yaml:"rowLength"`
}

// Settings for the fMP4 containers.
//
// The parent of this entity is HlsSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fmp4HlsSettingsProperty := &fmp4HlsSettingsProperty{
//   	audioRenditionSets: jsii.String("audioRenditionSets"),
//   	nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   	timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   }
//
type CfnChannel_Fmp4HlsSettingsProperty struct {
	// List all the audio groups that are used with the video output stream.
	//
	// Input all the audio GROUP-IDs that are associated to the video, separate by ','.
	AudioRenditionSets *string `field:"optional" json:"audioRenditionSets" yaml:"audioRenditionSets"`
	// If set to passthrough, Nielsen inaudible tones for media tracking will be detected in the input audio and an equivalent ID3 tag will be inserted in the output.
	NielsenId3Behavior *string `field:"optional" json:"nielsenId3Behavior" yaml:"nielsenId3Behavior"`
	// When set to passthrough, timed metadata is passed through from input to output.
	TimedMetadataBehavior *string `field:"optional" json:"timedMetadataBehavior" yaml:"timedMetadataBehavior"`
}

// Settings to configure the destination of a Frame Capture output.
//
// The parent of this entity is FrameCaptureGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   frameCaptureCdnSettingsProperty := &frameCaptureCdnSettingsProperty{
//   	frameCaptureS3Settings: &frameCaptureS3SettingsProperty{
//   		cannedAcl: jsii.String("cannedAcl"),
//   	},
//   }
//
type CfnChannel_FrameCaptureCdnSettingsProperty struct {
	// Sets up Amazon S3 as the destination for this Frame Capture output.
	FrameCaptureS3Settings interface{} `field:"optional" json:"frameCaptureS3Settings" yaml:"frameCaptureS3Settings"`
}

// The settings for a frame capture output group.
//
// The parent of this entity is OutputGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   frameCaptureGroupSettingsProperty := &frameCaptureGroupSettingsProperty{
//   	destination: &outputLocationRefProperty{
//   		destinationRefId: jsii.String("destinationRefId"),
//   	},
//   	frameCaptureCdnSettings: &frameCaptureCdnSettingsProperty{
//   		frameCaptureS3Settings: &frameCaptureS3SettingsProperty{
//   			cannedAcl: jsii.String("cannedAcl"),
//   		},
//   	},
//   }
//
type CfnChannel_FrameCaptureGroupSettingsProperty struct {
	// The destination for the frame capture files.
	//
	// The destination is either the URI for an Amazon S3 bucket and object, plus a file name prefix (for example, s3ssl://sportsDelivery/highlights/20180820/curling_) or the URI for a MediaStore container, plus a file name prefix (for example, mediastoressl://sportsDelivery/20180820/curling_). The final file names consist of the prefix from the destination field (for example, "curling_") + name modifier + the counter (5 digits, starting from 00001) + extension (which is always .jpg). For example, curlingLow.00001.jpg.
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// Settings to configure the destination of a Frame Capture output.
	FrameCaptureCdnSettings interface{} `field:"optional" json:"frameCaptureCdnSettings" yaml:"frameCaptureCdnSettings"`
}

// Settings for a frame capture output in an HLS output group.
//
// The parent of this entity is HlsSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   frameCaptureHlsSettingsProperty := &frameCaptureHlsSettingsProperty{
//   }
//
type CfnChannel_FrameCaptureHlsSettingsProperty struct {
}

// The frame capture output settings.
//
// The parent of this entity is OutputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   frameCaptureOutputSettingsProperty := &frameCaptureOutputSettingsProperty{
//   	nameModifier: jsii.String("nameModifier"),
//   }
//
type CfnChannel_FrameCaptureOutputSettingsProperty struct {
	// Required if the output group contains more than one output.
	//
	// This modifier forms part of the output file name.
	NameModifier *string `field:"optional" json:"nameModifier" yaml:"nameModifier"`
}

// Sets up Amazon S3 as the destination for this Frame Capture output.
//
// The parent of this entity is FrameCaptureCdnSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   frameCaptureS3SettingsProperty := &frameCaptureS3SettingsProperty{
//   	cannedAcl: jsii.String("cannedAcl"),
//   }
//
type CfnChannel_FrameCaptureS3SettingsProperty struct {
	// Specify the canned ACL to apply to each S3 request.
	//
	// Defaults to none.
	CannedAcl *string `field:"optional" json:"cannedAcl" yaml:"cannedAcl"`
}

// The frame capture settings.
//
// The parent of this entity is VideoCodecSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   frameCaptureSettingsProperty := &frameCaptureSettingsProperty{
//   	captureInterval: jsii.Number(123),
//   	captureIntervalUnits: jsii.String("captureIntervalUnits"),
//   }
//
type CfnChannel_FrameCaptureSettingsProperty struct {
	// The frequency, in seconds, for capturing frames for inclusion in the output.
	//
	// For example, "10" means capture a frame every 10 seconds.
	CaptureInterval *float64 `field:"optional" json:"captureInterval" yaml:"captureInterval"`
	// Unit for the frame capture interval.
	CaptureIntervalUnits *string `field:"optional" json:"captureIntervalUnits" yaml:"captureIntervalUnits"`
}

// The configuration settings that apply to the entire channel.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   globalConfigurationProperty := &globalConfigurationProperty{
//   	initialAudioGain: jsii.Number(123),
//   	inputEndAction: jsii.String("inputEndAction"),
//   	inputLossBehavior: &inputLossBehaviorProperty{
//   		blackFrameMsec: jsii.Number(123),
//   		inputLossImageColor: jsii.String("inputLossImageColor"),
//   		inputLossImageSlate: &inputLocationProperty{
//   			passwordParam: jsii.String("passwordParam"),
//   			uri: jsii.String("uri"),
//   			username: jsii.String("username"),
//   		},
//   		inputLossImageType: jsii.String("inputLossImageType"),
//   		repeatFrameMsec: jsii.Number(123),
//   	},
//   	outputLockingMode: jsii.String("outputLockingMode"),
//   	outputTimingSource: jsii.String("outputTimingSource"),
//   	supportLowFramerateInputs: jsii.String("supportLowFramerateInputs"),
//   }
//
type CfnChannel_GlobalConfigurationProperty struct {
	// The value to set the initial audio gain for the channel.
	InitialAudioGain *float64 `field:"optional" json:"initialAudioGain" yaml:"initialAudioGain"`
	// Indicates the action to take when the current input completes (for example, end-of-file).
	//
	// When switchAndLoopInputs is configured, MediaLive restarts at the beginning of the first input. When "none" is configured, MediaLive transcodes either black, a solid color, or a user-specified slate images per the "Input Loss Behavior" configuration until the next input switch occurs (which is controlled through the Channel Schedule API).
	InputEndAction *string `field:"optional" json:"inputEndAction" yaml:"inputEndAction"`
	// The settings for system actions when the input is lost.
	InputLossBehavior interface{} `field:"optional" json:"inputLossBehavior" yaml:"inputLossBehavior"`
	// Indicates how MediaLive pipelines are synchronized.
	//
	// PIPELINELOCKING - MediaLive attempts to synchronize the output of each pipeline to the other. EPOCHLOCKING - MediaLive attempts to synchronize the output of each pipeline to the Unix epoch.
	OutputLockingMode *string `field:"optional" json:"outputLockingMode" yaml:"outputLockingMode"`
	// Indicates whether the rate of frames emitted by the Live encoder should be paced by its system clock (which optionally might be locked to another source through NTP) or should be locked to the clock of the source that is providing the input stream.
	OutputTimingSource *string `field:"optional" json:"outputTimingSource" yaml:"outputTimingSource"`
	// Adjusts the video input buffer for streams with very low video frame rates.
	//
	// This is commonly set to enabled for music channels with less than one video frame per second.
	SupportLowFramerateInputs *string `field:"optional" json:"supportLowFramerateInputs" yaml:"supportLowFramerateInputs"`
}

// Settings for configuring color space in an H264 video encode.
//
// The parent of this entity is H264Settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   h264ColorSpaceSettingsProperty := &h264ColorSpaceSettingsProperty{
//   	colorSpacePassthroughSettings: &colorSpacePassthroughSettingsProperty{
//   	},
//   	rec601Settings: &rec601SettingsProperty{
//   	},
//   	rec709Settings: &rec709SettingsProperty{
//   	},
//   }
//
type CfnChannel_H264ColorSpaceSettingsProperty struct {
	// Passthrough applies no color space conversion to the output.
	ColorSpacePassthroughSettings interface{} `field:"optional" json:"colorSpacePassthroughSettings" yaml:"colorSpacePassthroughSettings"`
	// Settings to configure the handling of Rec601 color space.
	Rec601Settings interface{} `field:"optional" json:"rec601Settings" yaml:"rec601Settings"`
	// Settings to configure the handling of Rec709 color space.
	Rec709Settings interface{} `field:"optional" json:"rec709Settings" yaml:"rec709Settings"`
}

// Settings to configure video filters that apply to the H264 codec.
//
// The parent of this entity is H264Settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   h264FilterSettingsProperty := &h264FilterSettingsProperty{
//   	temporalFilterSettings: &temporalFilterSettingsProperty{
//   		postFilterSharpening: jsii.String("postFilterSharpening"),
//   		strength: jsii.String("strength"),
//   	},
//   }
//
type CfnChannel_H264FilterSettingsProperty struct {
	// Settings for applying the temporal filter to the video.
	TemporalFilterSettings interface{} `field:"optional" json:"temporalFilterSettings" yaml:"temporalFilterSettings"`
}

// The settings for the H.264 codec in the output.
//
// The parent of this entity is VideoCodecSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   h264SettingsProperty := &h264SettingsProperty{
//   	adaptiveQuantization: jsii.String("adaptiveQuantization"),
//   	afdSignaling: jsii.String("afdSignaling"),
//   	bitrate: jsii.Number(123),
//   	bufFillPct: jsii.Number(123),
//   	bufSize: jsii.Number(123),
//   	colorMetadata: jsii.String("colorMetadata"),
//   	colorSpaceSettings: &h264ColorSpaceSettingsProperty{
//   		colorSpacePassthroughSettings: &colorSpacePassthroughSettingsProperty{
//   		},
//   		rec601Settings: &rec601SettingsProperty{
//   		},
//   		rec709Settings: &rec709SettingsProperty{
//   		},
//   	},
//   	entropyEncoding: jsii.String("entropyEncoding"),
//   	filterSettings: &h264FilterSettingsProperty{
//   		temporalFilterSettings: &temporalFilterSettingsProperty{
//   			postFilterSharpening: jsii.String("postFilterSharpening"),
//   			strength: jsii.String("strength"),
//   		},
//   	},
//   	fixedAfd: jsii.String("fixedAfd"),
//   	flickerAq: jsii.String("flickerAq"),
//   	forceFieldPictures: jsii.String("forceFieldPictures"),
//   	framerateControl: jsii.String("framerateControl"),
//   	framerateDenominator: jsii.Number(123),
//   	framerateNumerator: jsii.Number(123),
//   	gopBReference: jsii.String("gopBReference"),
//   	gopClosedCadence: jsii.Number(123),
//   	gopNumBFrames: jsii.Number(123),
//   	gopSize: jsii.Number(123),
//   	gopSizeUnits: jsii.String("gopSizeUnits"),
//   	level: jsii.String("level"),
//   	lookAheadRateControl: jsii.String("lookAheadRateControl"),
//   	maxBitrate: jsii.Number(123),
//   	minIInterval: jsii.Number(123),
//   	numRefFrames: jsii.Number(123),
//   	parControl: jsii.String("parControl"),
//   	parDenominator: jsii.Number(123),
//   	parNumerator: jsii.Number(123),
//   	profile: jsii.String("profile"),
//   	qualityLevel: jsii.String("qualityLevel"),
//   	qvbrQualityLevel: jsii.Number(123),
//   	rateControlMode: jsii.String("rateControlMode"),
//   	scanType: jsii.String("scanType"),
//   	sceneChangeDetect: jsii.String("sceneChangeDetect"),
//   	slices: jsii.Number(123),
//   	softness: jsii.Number(123),
//   	spatialAq: jsii.String("spatialAq"),
//   	subgopLength: jsii.String("subgopLength"),
//   	syntax: jsii.String("syntax"),
//   	temporalAq: jsii.String("temporalAq"),
//   	timecodeInsertion: jsii.String("timecodeInsertion"),
//   }
//
type CfnChannel_H264SettingsProperty struct {
	// The adaptive quantization.
	//
	// This allows intra-frame quantizers to vary to improve visual quality.
	AdaptiveQuantization *string `field:"optional" json:"adaptiveQuantization" yaml:"adaptiveQuantization"`
	// Indicates that AFD values will be written into the output stream.
	//
	// If afdSignaling is auto, the system tries to preserve the input AFD value (in cases where multiple AFD values are valid). If set to fixed, the AFD value is the value configured in the fixedAfd parameter.
	AfdSignaling *string `field:"optional" json:"afdSignaling" yaml:"afdSignaling"`
	// The average bitrate in bits/second.
	//
	// This is required when the rate control mode is VBR or CBR. It isn't used for QVBR. In a Microsoft Smooth output group, each output must have a unique value when its bitrate is rounded down to the nearest multiple of 1000.
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// The percentage of the buffer that should initially be filled (HRD buffer model).
	BufFillPct *float64 `field:"optional" json:"bufFillPct" yaml:"bufFillPct"`
	// The size of the buffer (HRD buffer model) in bits/second.
	BufSize *float64 `field:"optional" json:"bufSize" yaml:"bufSize"`
	// Includes color space metadata in the output.
	ColorMetadata *string `field:"optional" json:"colorMetadata" yaml:"colorMetadata"`
	// Settings to configure the color space handling for the video.
	ColorSpaceSettings interface{} `field:"optional" json:"colorSpaceSettings" yaml:"colorSpaceSettings"`
	// The entropy encoding mode.
	//
	// Use cabac (must be in Main or High profile) or cavlc.
	EntropyEncoding *string `field:"optional" json:"entropyEncoding" yaml:"entropyEncoding"`
	// Optional filters that you can apply to an encode.
	FilterSettings interface{} `field:"optional" json:"filterSettings" yaml:"filterSettings"`
	// A four-bit AFD value to write on all frames of video in the output stream.
	//
	// Valid only when afdSignaling is set to Fixed.
	FixedAfd *string `field:"optional" json:"fixedAfd" yaml:"fixedAfd"`
	// If set to enabled, adjusts the quantization within each frame to reduce flicker or pop on I-frames.
	FlickerAq *string `field:"optional" json:"flickerAq" yaml:"flickerAq"`
	// This setting applies only when scan type is "interlaced." It controls whether coding is performed on a field basis or on a frame basis. (When the video is progressive, the coding is always performed on a frame basis.) enabled: Force MediaLive to code on a field basis, so that odd and even sets of fields are coded separately. disabled: Code the two sets of fields separately (on a field basis) or together (on a frame basis using PAFF), depending on what is most appropriate for the content.
	ForceFieldPictures *string `field:"optional" json:"forceFieldPictures" yaml:"forceFieldPictures"`
	// Indicates how the output video frame rate is specified.
	//
	// If you select "specified," the output video frame rate is determined by framerateNumerator and framerateDenominator. If you select "initializeFromSource," the output video frame rate is set equal to the input video frame rate of the first input.
	FramerateControl *string `field:"optional" json:"framerateControl" yaml:"framerateControl"`
	// The frame rate denominator.
	FramerateDenominator *float64 `field:"optional" json:"framerateDenominator" yaml:"framerateDenominator"`
	// The frame rate numerator.
	//
	// The frame rate is a fraction, for example, 24000/1001 = 23.976 fps.
	FramerateNumerator *float64 `field:"optional" json:"framerateNumerator" yaml:"framerateNumerator"`
	// If enabled, uses reference B frames for GOP structures that have B frames > 1.
	GopBReference *string `field:"optional" json:"gopBReference" yaml:"gopBReference"`
	// The frequency of closed GOPs.
	//
	// In streaming applications, we recommend that you set this to 1 so that a decoder joining mid-stream will receive an IDR frame as quickly as possible. Setting this value to 0 will break output segmenting.
	GopClosedCadence *float64 `field:"optional" json:"gopClosedCadence" yaml:"gopClosedCadence"`
	// The number of B-frames between reference frames.
	GopNumBFrames *float64 `field:"optional" json:"gopNumBFrames" yaml:"gopNumBFrames"`
	// The GOP size (keyframe interval) in units of either frames or seconds per gopSizeUnits.
	//
	// The value must be greater than zero.
	GopSize *float64 `field:"optional" json:"gopSize" yaml:"gopSize"`
	// Indicates if the gopSize is specified in frames or seconds.
	//
	// If seconds, the system converts the gopSize into a frame count at runtime.
	GopSizeUnits *string `field:"optional" json:"gopSizeUnits" yaml:"gopSizeUnits"`
	// The H.264 level.
	Level *string `field:"optional" json:"level" yaml:"level"`
	// The amount of lookahead.
	//
	// A value of low can decrease latency and memory usage, while high can produce better quality for certain content.
	LookAheadRateControl *string `field:"optional" json:"lookAheadRateControl" yaml:"lookAheadRateControl"`
	// For QVBR: See the tooltip for Quality level.
	//
	// For VBR: Set the maximum bitrate in order to accommodate expected spikes in the complexity of the video.
	MaxBitrate *float64 `field:"optional" json:"maxBitrate" yaml:"maxBitrate"`
	// Meaningful only if sceneChangeDetect is set to enabled.
	//
	// This setting enforces separation between repeated (cadence) I-frames and I-frames inserted by Scene Change Detection. If a scene change I-frame is within I-interval frames of a cadence I-frame, the GOP is shrunk or stretched to the scene change I-frame. GOP stretch requires enabling lookahead as well as setting the I-interval. The normal cadence resumes for the next GOP. Note that the maximum GOP stretch = GOP size + Min-I-interval - 1.
	MinIInterval *float64 `field:"optional" json:"minIInterval" yaml:"minIInterval"`
	// The number of reference frames to use.
	//
	// The encoder might use more than requested if you use B-frames or interlaced encoding.
	NumRefFrames *float64 `field:"optional" json:"numRefFrames" yaml:"numRefFrames"`
	// Indicates how the output pixel aspect ratio is specified.
	//
	// If "specified" is selected, the output video pixel aspect ratio is determined by parNumerator and parDenominator. If "initializeFromSource" is selected, the output pixels aspect ratio will be set equal to the input video pixel aspect ratio of the first input.
	ParControl *string `field:"optional" json:"parControl" yaml:"parControl"`
	// The Pixel Aspect Ratio denominator.
	ParDenominator *float64 `field:"optional" json:"parDenominator" yaml:"parDenominator"`
	// The Pixel Aspect Ratio numerator.
	ParNumerator *float64 `field:"optional" json:"parNumerator" yaml:"parNumerator"`
	// An H.264 profile.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Leave as STANDARD_QUALITY or choose a different value (which might result in additional costs to run the channel).
	//
	// - ENHANCED_QUALITY: Produces a slightly better video quality without an increase in the bitrate. Has an effect only when the Rate control mode is QVBR or CBR. If this channel is in a MediaLive multiplex, the value must be ENHANCED_QUALITY.
	// - STANDARD_QUALITY: Valid for any Rate control mode.
	QualityLevel *string `field:"optional" json:"qualityLevel" yaml:"qualityLevel"`
	// Controls the target quality for the video encode.
	//
	// This applies only when the rate control mode is QVBR. Set values for the QVBR quality level field and Max bitrate field that suit your most important viewing devices. Recommended values are: - Primary screen: Quality level: 8 to 10. Max bitrate: 4M - PC or tablet: Quality level: 7. Max bitrate: 1.5M to 3M - Smartphone: Quality level: 6. Max bitrate: 1M to 1.5M.
	QvbrQualityLevel *float64 `field:"optional" json:"qvbrQualityLevel" yaml:"qvbrQualityLevel"`
	// The rate control mode.
	//
	// QVBR: The quality will match the specified quality level except when it is constrained by the maximum bitrate. We recommend this if you or your viewers pay for bandwidth. VBR: The quality and bitrate vary, depending on the video complexity. We recommend this instead of QVBR if you want to maintain a specific average bitrate over the duration of the channel. CBR: The quality varies, depending on the video complexity. We recommend this only if you distribute your assets to devices that can't handle variable bitrates.
	RateControlMode *string `field:"optional" json:"rateControlMode" yaml:"rateControlMode"`
	// Sets the scan type of the output to progressive or top-field-first interlaced.
	ScanType *string `field:"optional" json:"scanType" yaml:"scanType"`
	// The scene change detection.
	//
	// On: inserts I-frames when the scene change is detected. Off: does not force an I-frame when the scene change is detected.
	SceneChangeDetect *string `field:"optional" json:"sceneChangeDetect" yaml:"sceneChangeDetect"`
	// The number of slices per picture.
	//
	// The number must be less than or equal to the number of macroblock rows for progressive pictures, and less than or equal to half the number of macroblock rows for interlaced pictures. This field is optional. If you don't specify a value, MediaLive chooses the number of slices based on the encode resolution.
	Slices *float64 `field:"optional" json:"slices" yaml:"slices"`
	// Softness.
	//
	// Selects a quantizer matrix. Larger values reduce high-frequency content in the encoded image.
	Softness *float64 `field:"optional" json:"softness" yaml:"softness"`
	// If set to enabled, adjusts quantization within each frame based on the spatial variation of content complexity.
	SpatialAq *string `field:"optional" json:"spatialAq" yaml:"spatialAq"`
	// If set to fixed, uses gopNumBFrames B-frames per sub-GOP.
	//
	// If set to dynamic, optimizes the number of B-frames used for each sub-GOP to improve visual quality.
	SubgopLength *string `field:"optional" json:"subgopLength" yaml:"subgopLength"`
	// Produces a bitstream that is compliant with SMPTE RP-2027.
	Syntax *string `field:"optional" json:"syntax" yaml:"syntax"`
	// If set to enabled, adjusts quantization within each frame based on the temporal variation of content complexity.
	TemporalAq *string `field:"optional" json:"temporalAq" yaml:"temporalAq"`
	// Determines how timecodes should be inserted into the video elementary stream.
	//
	// disabled: don't include timecodes. picTimingSei: pass through picture timing SEI messages from the source specified in Timecode Config.
	TimecodeInsertion *string `field:"optional" json:"timecodeInsertion" yaml:"timecodeInsertion"`
}

// H265 Color Space Settings.
//
// The parent of this entity is H265Settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   h265ColorSpaceSettingsProperty := &h265ColorSpaceSettingsProperty{
//   	colorSpacePassthroughSettings: &colorSpacePassthroughSettingsProperty{
//   	},
//   	hdr10Settings: &hdr10SettingsProperty{
//   		maxCll: jsii.Number(123),
//   		maxFall: jsii.Number(123),
//   	},
//   	rec601Settings: &rec601SettingsProperty{
//   	},
//   	rec709Settings: &rec709SettingsProperty{
//   	},
//   }
//
type CfnChannel_H265ColorSpaceSettingsProperty struct {
	// Passthrough applies no color space conversion to the output.
	ColorSpacePassthroughSettings interface{} `field:"optional" json:"colorSpacePassthroughSettings" yaml:"colorSpacePassthroughSettings"`
	// Settings to configure the handling of HDR10 color space.
	Hdr10Settings interface{} `field:"optional" json:"hdr10Settings" yaml:"hdr10Settings"`
	// Settings to configure the handling of Rec601 color space.
	Rec601Settings interface{} `field:"optional" json:"rec601Settings" yaml:"rec601Settings"`
	// Settings to configure the handling of Rec709 color space.
	Rec709Settings interface{} `field:"optional" json:"rec709Settings" yaml:"rec709Settings"`
}

// Settings to configure video filters that apply to the H265 codec.
//
// The parent of this entity is H265Settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   h265FilterSettingsProperty := &h265FilterSettingsProperty{
//   	temporalFilterSettings: &temporalFilterSettingsProperty{
//   		postFilterSharpening: jsii.String("postFilterSharpening"),
//   		strength: jsii.String("strength"),
//   	},
//   }
//
type CfnChannel_H265FilterSettingsProperty struct {
	// Settings for applying the temporal filter to the video.
	TemporalFilterSettings interface{} `field:"optional" json:"temporalFilterSettings" yaml:"temporalFilterSettings"`
}

// H265 Settings.
//
// The parent of this entity is VideoCodecSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   h265SettingsProperty := &h265SettingsProperty{
//   	adaptiveQuantization: jsii.String("adaptiveQuantization"),
//   	afdSignaling: jsii.String("afdSignaling"),
//   	alternativeTransferFunction: jsii.String("alternativeTransferFunction"),
//   	bitrate: jsii.Number(123),
//   	bufSize: jsii.Number(123),
//   	colorMetadata: jsii.String("colorMetadata"),
//   	colorSpaceSettings: &h265ColorSpaceSettingsProperty{
//   		colorSpacePassthroughSettings: &colorSpacePassthroughSettingsProperty{
//   		},
//   		hdr10Settings: &hdr10SettingsProperty{
//   			maxCll: jsii.Number(123),
//   			maxFall: jsii.Number(123),
//   		},
//   		rec601Settings: &rec601SettingsProperty{
//   		},
//   		rec709Settings: &rec709SettingsProperty{
//   		},
//   	},
//   	filterSettings: &h265FilterSettingsProperty{
//   		temporalFilterSettings: &temporalFilterSettingsProperty{
//   			postFilterSharpening: jsii.String("postFilterSharpening"),
//   			strength: jsii.String("strength"),
//   		},
//   	},
//   	fixedAfd: jsii.String("fixedAfd"),
//   	flickerAq: jsii.String("flickerAq"),
//   	framerateDenominator: jsii.Number(123),
//   	framerateNumerator: jsii.Number(123),
//   	gopClosedCadence: jsii.Number(123),
//   	gopSize: jsii.Number(123),
//   	gopSizeUnits: jsii.String("gopSizeUnits"),
//   	level: jsii.String("level"),
//   	lookAheadRateControl: jsii.String("lookAheadRateControl"),
//   	maxBitrate: jsii.Number(123),
//   	minIInterval: jsii.Number(123),
//   	parDenominator: jsii.Number(123),
//   	parNumerator: jsii.Number(123),
//   	profile: jsii.String("profile"),
//   	qvbrQualityLevel: jsii.Number(123),
//   	rateControlMode: jsii.String("rateControlMode"),
//   	scanType: jsii.String("scanType"),
//   	sceneChangeDetect: jsii.String("sceneChangeDetect"),
//   	slices: jsii.Number(123),
//   	tier: jsii.String("tier"),
//   	timecodeInsertion: jsii.String("timecodeInsertion"),
//   }
//
type CfnChannel_H265SettingsProperty struct {
	// Adaptive quantization.
	//
	// Allows intra-frame quantizers to vary to improve visual quality.
	AdaptiveQuantization *string `field:"optional" json:"adaptiveQuantization" yaml:"adaptiveQuantization"`
	// Indicates that AFD values will be written into the output stream.
	//
	// If afdSignaling is "auto", the system will try to preserve the input AFD value (in cases where multiple AFD values are valid). If set to "fixed", the AFD value will be the value configured in the fixedAfd parameter.
	AfdSignaling *string `field:"optional" json:"afdSignaling" yaml:"afdSignaling"`
	// Whether or not EML should insert an Alternative Transfer Function SEI message to support backwards compatibility with non-HDR decoders and displays.
	AlternativeTransferFunction *string `field:"optional" json:"alternativeTransferFunction" yaml:"alternativeTransferFunction"`
	// Average bitrate in bits/second.
	//
	// Required when the rate control mode is VBR or CBR. Not used for QVBR. In an MS Smooth output group, each output must have a unique value when its bitrate is rounded down to the nearest multiple of 1000.
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// Size of buffer (HRD buffer model) in bits.
	BufSize *float64 `field:"optional" json:"bufSize" yaml:"bufSize"`
	// Includes colorspace metadata in the output.
	ColorMetadata *string `field:"optional" json:"colorMetadata" yaml:"colorMetadata"`
	// Color Space settings.
	ColorSpaceSettings interface{} `field:"optional" json:"colorSpaceSettings" yaml:"colorSpaceSettings"`
	// Optional filters that you can apply to an encode.
	FilterSettings interface{} `field:"optional" json:"filterSettings" yaml:"filterSettings"`
	// Four bit AFD value to write on all frames of video in the output stream.
	//
	// Only valid when afdSignaling is set to 'Fixed'.
	FixedAfd *string `field:"optional" json:"fixedAfd" yaml:"fixedAfd"`
	// If set to enabled, adjust quantization within each frame to reduce flicker or 'pop' on I-frames.
	FlickerAq *string `field:"optional" json:"flickerAq" yaml:"flickerAq"`
	// Framerate denominator.
	FramerateDenominator *float64 `field:"optional" json:"framerateDenominator" yaml:"framerateDenominator"`
	// Framerate numerator - framerate is a fraction, e.g. 24000 / 1001 = 23.976 fps.
	FramerateNumerator *float64 `field:"optional" json:"framerateNumerator" yaml:"framerateNumerator"`
	// Frequency of closed GOPs.
	//
	// In streaming applications, it is recommended that this be set to 1 so a decoder joining mid-stream will receive an IDR frame as quickly as possible. Setting this value to 0 will break output segmenting.
	GopClosedCadence *float64 `field:"optional" json:"gopClosedCadence" yaml:"gopClosedCadence"`
	// GOP size (keyframe interval) in units of either frames or seconds per gopSizeUnits.
	//
	// If gopSizeUnits is frames, gopSize must be an integer and must be greater than or equal to 1.
	// If gopSizeUnits is seconds, gopSize must be greater than 0, but need not be an integer.
	GopSize *float64 `field:"optional" json:"gopSize" yaml:"gopSize"`
	// Indicates if the gopSize is specified in frames or seconds.
	//
	// If seconds the system will convert the gopSize into a frame count at run time.
	GopSizeUnits *string `field:"optional" json:"gopSizeUnits" yaml:"gopSizeUnits"`
	// H.265 Level.
	Level *string `field:"optional" json:"level" yaml:"level"`
	// Amount of lookahead.
	//
	// A value of low can decrease latency and memory usage, while high can produce better quality for certain content.
	LookAheadRateControl *string `field:"optional" json:"lookAheadRateControl" yaml:"lookAheadRateControl"`
	// For QVBR: See the tooltip for Quality level.
	MaxBitrate *float64 `field:"optional" json:"maxBitrate" yaml:"maxBitrate"`
	// Only meaningful if sceneChangeDetect is set to enabled.
	//
	// Defaults to 5 if multiplex rate control is used. Enforces separation between repeated (cadence) I-frames and I-frames inserted by Scene Change Detection. If a scene change I-frame is within I-interval frames of a cadence I-frame, the GOP is shrunk and/or stretched to the scene change I-frame. GOP stretch requires enabling lookahead as well as setting I-interval. The normal cadence resumes for the next GOP. Note: Maximum GOP stretch = GOP size + Min-I-interval - 1
	MinIInterval *float64 `field:"optional" json:"minIInterval" yaml:"minIInterval"`
	// Pixel Aspect Ratio denominator.
	ParDenominator *float64 `field:"optional" json:"parDenominator" yaml:"parDenominator"`
	// Pixel Aspect Ratio numerator.
	ParNumerator *float64 `field:"optional" json:"parNumerator" yaml:"parNumerator"`
	// H.265 Profile.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Controls the target quality for the video encode.
	//
	// Applies only when the rate control mode is QVBR. Set values for the QVBR quality level field and Max bitrate field that suit your most important viewing devices. Recommended values are:
	// - Primary screen: Quality level: 8 to 10. Max bitrate: 4M
	// - PC or tablet: Quality level: 7. Max bitrate: 1.5M to 3M
	// - Smartphone: Quality level: 6. Max bitrate: 1M to 1.5M
	QvbrQualityLevel *float64 `field:"optional" json:"qvbrQualityLevel" yaml:"qvbrQualityLevel"`
	// Rate control mode.
	//
	// QVBR: Quality will match the specified quality level except when it is constrained by the
	// maximum bitrate. Recommended if you or your viewers pay for bandwidth. CBR: Quality varies, depending on the video complexity. Recommended only if you distribute
	// your assets to devices that cannot handle variable bitrates. Multiplex: This rate control mode is only supported (and is required) when the video is being
	// delivered to a MediaLive Multiplex in which case the rate control configuration is controlled
	// by the properties within the Multiplex Program.
	RateControlMode *string `field:"optional" json:"rateControlMode" yaml:"rateControlMode"`
	// Sets the scan type of the output to progressive or top-field-first interlaced.
	ScanType *string `field:"optional" json:"scanType" yaml:"scanType"`
	// Scene change detection.
	SceneChangeDetect *string `field:"optional" json:"sceneChangeDetect" yaml:"sceneChangeDetect"`
	// Number of slices per picture.
	//
	// Must be less than or equal to the number of macroblock rows for progressive pictures, and less than or equal to half the number of macroblock rows for interlaced pictures.
	// This field is optional; when no value is specified the encoder will choose the number of slices based on encode resolution.
	Slices *float64 `field:"optional" json:"slices" yaml:"slices"`
	// H.265 Tier.
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
	// Determines how timecodes should be inserted into the video elementary stream.
	//
	// - 'disabled': Do not include timecodes
	// - 'picTimingSei': Pass through picture timing SEI messages from the source specified in Timecode Config.
	TimecodeInsertion *string `field:"optional" json:"timecodeInsertion" yaml:"timecodeInsertion"`
}

// Hdr10 Settings.
//
// The parents of this entity are H265ColorSpaceSettings (for color space settings in the output) and VideoSelectorColorSpaceSettings (for color space settings in the input).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hdr10SettingsProperty := &hdr10SettingsProperty{
//   	maxCll: jsii.Number(123),
//   	maxFall: jsii.Number(123),
//   }
//
type CfnChannel_Hdr10SettingsProperty struct {
	// Maximum Content Light Level An integer metadata value defining the maximum light level, in nits, of any single pixel within an encoded HDR video stream or file.
	MaxCll *float64 `field:"optional" json:"maxCll" yaml:"maxCll"`
	// Maximum Frame Average Light Level An integer metadata value defining the maximum average light level, in nits, for any single frame within an encoded HDR video stream or file.
	MaxFall *float64 `field:"optional" json:"maxFall" yaml:"maxFall"`
}

// The Akamai settings in an HLS output.
//
// The parent of this entity is HlsCdnSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsAkamaiSettingsProperty := &hlsAkamaiSettingsProperty{
//   	connectionRetryInterval: jsii.Number(123),
//   	filecacheDuration: jsii.Number(123),
//   	httpTransferMode: jsii.String("httpTransferMode"),
//   	numRetries: jsii.Number(123),
//   	restartDelay: jsii.Number(123),
//   	salt: jsii.String("salt"),
//   	token: jsii.String("token"),
//   }
//
type CfnChannel_HlsAkamaiSettingsProperty struct {
	// The number of seconds to wait before retrying a connection to the CDN if the connection is lost.
	ConnectionRetryInterval *float64 `field:"optional" json:"connectionRetryInterval" yaml:"connectionRetryInterval"`
	// The size, in seconds, of the file cache for streaming outputs.
	FilecacheDuration *float64 `field:"optional" json:"filecacheDuration" yaml:"filecacheDuration"`
	// Specifies whether to use chunked transfer encoding to Akamai.
	//
	// To enable this feature, contact Akamai.
	HttpTransferMode *string `field:"optional" json:"httpTransferMode" yaml:"httpTransferMode"`
	// The number of retry attempts that will be made before the channel is put into an error state.
	NumRetries *float64 `field:"optional" json:"numRetries" yaml:"numRetries"`
	// If a streaming output fails, the number of seconds to wait until a restart is initiated.
	//
	// A value of 0 means never restart.
	RestartDelay *float64 `field:"optional" json:"restartDelay" yaml:"restartDelay"`
	// The salt for authenticated Akamai.
	Salt *string `field:"optional" json:"salt" yaml:"salt"`
	// The token parameter for authenticated Akamai.
	//
	// If this is not specified, _gda_ is used.
	Token *string `field:"optional" json:"token" yaml:"token"`
}

// The configuration of HLS Basic Put Settings.
//
// The parent of this entity is HlsCdnSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsBasicPutSettingsProperty := &hlsBasicPutSettingsProperty{
//   	connectionRetryInterval: jsii.Number(123),
//   	filecacheDuration: jsii.Number(123),
//   	numRetries: jsii.Number(123),
//   	restartDelay: jsii.Number(123),
//   }
//
type CfnChannel_HlsBasicPutSettingsProperty struct {
	// The number of seconds to wait before retrying a connection to the CDN if the connection is lost.
	ConnectionRetryInterval *float64 `field:"optional" json:"connectionRetryInterval" yaml:"connectionRetryInterval"`
	// The size, in seconds, of the file cache for streaming outputs.
	FilecacheDuration *float64 `field:"optional" json:"filecacheDuration" yaml:"filecacheDuration"`
	// The number of retry attempts that MediaLive makes before the channel is put into an error state.
	NumRetries *float64 `field:"optional" json:"numRetries" yaml:"numRetries"`
	// If a streaming output fails, the number of seconds to wait until a restart is initiated.
	//
	// A value of 0 means never restart.
	RestartDelay *float64 `field:"optional" json:"restartDelay" yaml:"restartDelay"`
}

// The settings for the CDN of an HLS output.
//
// The parent of this entity is HlsGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsCdnSettingsProperty := &hlsCdnSettingsProperty{
//   	hlsAkamaiSettings: &hlsAkamaiSettingsProperty{
//   		connectionRetryInterval: jsii.Number(123),
//   		filecacheDuration: jsii.Number(123),
//   		httpTransferMode: jsii.String("httpTransferMode"),
//   		numRetries: jsii.Number(123),
//   		restartDelay: jsii.Number(123),
//   		salt: jsii.String("salt"),
//   		token: jsii.String("token"),
//   	},
//   	hlsBasicPutSettings: &hlsBasicPutSettingsProperty{
//   		connectionRetryInterval: jsii.Number(123),
//   		filecacheDuration: jsii.Number(123),
//   		numRetries: jsii.Number(123),
//   		restartDelay: jsii.Number(123),
//   	},
//   	hlsMediaStoreSettings: &hlsMediaStoreSettingsProperty{
//   		connectionRetryInterval: jsii.Number(123),
//   		filecacheDuration: jsii.Number(123),
//   		mediaStoreStorageClass: jsii.String("mediaStoreStorageClass"),
//   		numRetries: jsii.Number(123),
//   		restartDelay: jsii.Number(123),
//   	},
//   	hlsS3Settings: &hlsS3SettingsProperty{
//   		cannedAcl: jsii.String("cannedAcl"),
//   	},
//   	hlsWebdavSettings: &hlsWebdavSettingsProperty{
//   		connectionRetryInterval: jsii.Number(123),
//   		filecacheDuration: jsii.Number(123),
//   		httpTransferMode: jsii.String("httpTransferMode"),
//   		numRetries: jsii.Number(123),
//   		restartDelay: jsii.Number(123),
//   	},
//   }
//
type CfnChannel_HlsCdnSettingsProperty struct {
	// Sets up Akamai as the downstream system for the HLS output group.
	HlsAkamaiSettings interface{} `field:"optional" json:"hlsAkamaiSettings" yaml:"hlsAkamaiSettings"`
	// The settings for Basic Put for the HLS output.
	HlsBasicPutSettings interface{} `field:"optional" json:"hlsBasicPutSettings" yaml:"hlsBasicPutSettings"`
	// Sets up MediaStore as the destination for the HLS output.
	HlsMediaStoreSettings interface{} `field:"optional" json:"hlsMediaStoreSettings" yaml:"hlsMediaStoreSettings"`
	// Sets up Amazon S3 as the destination for this HLS output.
	HlsS3Settings interface{} `field:"optional" json:"hlsS3Settings" yaml:"hlsS3Settings"`
	// The settings for Web VTT captions in the HLS output group.
	//
	// The parent of this entity is HlsGroupSettings.
	HlsWebdavSettings interface{} `field:"optional" json:"hlsWebdavSettings" yaml:"hlsWebdavSettings"`
}

// The settings for an HLS output group.
//
// The parent of this entity is OutputGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsGroupSettingsProperty := &hlsGroupSettingsProperty{
//   	adMarkers: []*string{
//   		jsii.String("adMarkers"),
//   	},
//   	baseUrlContent: jsii.String("baseUrlContent"),
//   	baseUrlContent1: jsii.String("baseUrlContent1"),
//   	baseUrlManifest: jsii.String("baseUrlManifest"),
//   	baseUrlManifest1: jsii.String("baseUrlManifest1"),
//   	captionLanguageMappings: []interface{}{
//   		&captionLanguageMappingProperty{
//   			captionChannel: jsii.Number(123),
//   			languageCode: jsii.String("languageCode"),
//   			languageDescription: jsii.String("languageDescription"),
//   		},
//   	},
//   	captionLanguageSetting: jsii.String("captionLanguageSetting"),
//   	clientCache: jsii.String("clientCache"),
//   	codecSpecification: jsii.String("codecSpecification"),
//   	constantIv: jsii.String("constantIv"),
//   	destination: &outputLocationRefProperty{
//   		destinationRefId: jsii.String("destinationRefId"),
//   	},
//   	directoryStructure: jsii.String("directoryStructure"),
//   	discontinuityTags: jsii.String("discontinuityTags"),
//   	encryptionType: jsii.String("encryptionType"),
//   	hlsCdnSettings: &hlsCdnSettingsProperty{
//   		hlsAkamaiSettings: &hlsAkamaiSettingsProperty{
//   			connectionRetryInterval: jsii.Number(123),
//   			filecacheDuration: jsii.Number(123),
//   			httpTransferMode: jsii.String("httpTransferMode"),
//   			numRetries: jsii.Number(123),
//   			restartDelay: jsii.Number(123),
//   			salt: jsii.String("salt"),
//   			token: jsii.String("token"),
//   		},
//   		hlsBasicPutSettings: &hlsBasicPutSettingsProperty{
//   			connectionRetryInterval: jsii.Number(123),
//   			filecacheDuration: jsii.Number(123),
//   			numRetries: jsii.Number(123),
//   			restartDelay: jsii.Number(123),
//   		},
//   		hlsMediaStoreSettings: &hlsMediaStoreSettingsProperty{
//   			connectionRetryInterval: jsii.Number(123),
//   			filecacheDuration: jsii.Number(123),
//   			mediaStoreStorageClass: jsii.String("mediaStoreStorageClass"),
//   			numRetries: jsii.Number(123),
//   			restartDelay: jsii.Number(123),
//   		},
//   		hlsS3Settings: &hlsS3SettingsProperty{
//   			cannedAcl: jsii.String("cannedAcl"),
//   		},
//   		hlsWebdavSettings: &hlsWebdavSettingsProperty{
//   			connectionRetryInterval: jsii.Number(123),
//   			filecacheDuration: jsii.Number(123),
//   			httpTransferMode: jsii.String("httpTransferMode"),
//   			numRetries: jsii.Number(123),
//   			restartDelay: jsii.Number(123),
//   		},
//   	},
//   	hlsId3SegmentTagging: jsii.String("hlsId3SegmentTagging"),
//   	iFrameOnlyPlaylists: jsii.String("iFrameOnlyPlaylists"),
//   	incompleteSegmentBehavior: jsii.String("incompleteSegmentBehavior"),
//   	indexNSegments: jsii.Number(123),
//   	inputLossAction: jsii.String("inputLossAction"),
//   	ivInManifest: jsii.String("ivInManifest"),
//   	ivSource: jsii.String("ivSource"),
//   	keepSegments: jsii.Number(123),
//   	keyFormat: jsii.String("keyFormat"),
//   	keyFormatVersions: jsii.String("keyFormatVersions"),
//   	keyProviderSettings: &keyProviderSettingsProperty{
//   		staticKeySettings: &staticKeySettingsProperty{
//   			keyProviderServer: &inputLocationProperty{
//   				passwordParam: jsii.String("passwordParam"),
//   				uri: jsii.String("uri"),
//   				username: jsii.String("username"),
//   			},
//   			staticKeyValue: jsii.String("staticKeyValue"),
//   		},
//   	},
//   	manifestCompression: jsii.String("manifestCompression"),
//   	manifestDurationFormat: jsii.String("manifestDurationFormat"),
//   	minSegmentLength: jsii.Number(123),
//   	mode: jsii.String("mode"),
//   	outputSelection: jsii.String("outputSelection"),
//   	programDateTime: jsii.String("programDateTime"),
//   	programDateTimeClock: jsii.String("programDateTimeClock"),
//   	programDateTimePeriod: jsii.Number(123),
//   	redundantManifest: jsii.String("redundantManifest"),
//   	segmentationMode: jsii.String("segmentationMode"),
//   	segmentLength: jsii.Number(123),
//   	segmentsPerSubdirectory: jsii.Number(123),
//   	streamInfResolution: jsii.String("streamInfResolution"),
//   	timedMetadataId3Frame: jsii.String("timedMetadataId3Frame"),
//   	timedMetadataId3Period: jsii.Number(123),
//   	timestampDeltaMilliseconds: jsii.Number(123),
//   	tsFileMode: jsii.String("tsFileMode"),
//   }
//
type CfnChannel_HlsGroupSettingsProperty struct {
	// Chooses one or more ad marker types to pass SCTE35 signals through to this group of Apple HLS outputs.
	AdMarkers *[]*string `field:"optional" json:"adMarkers" yaml:"adMarkers"`
	// A partial URI prefix that will be prepended to each output in the media .m3u8 file. The partial URI prefix can be used if the base manifest is delivered from a different URL than the main .m3u8 file.
	BaseUrlContent *string `field:"optional" json:"baseUrlContent" yaml:"baseUrlContent"`
	// Optional.
	//
	// One value per output group. This field is required only if you are completing Base URL content A, and the downstream system has notified you that the media files for pipeline 1 of all outputs are in a location different from the media files for pipeline 0.
	BaseUrlContent1 *string `field:"optional" json:"baseUrlContent1" yaml:"baseUrlContent1"`
	// A partial URI prefix that will be prepended to each output in the media .m3u8 file. The partial URI prefix can be used if the base manifest is delivered from a different URL than the main .m3u8 file.
	BaseUrlManifest *string `field:"optional" json:"baseUrlManifest" yaml:"baseUrlManifest"`
	// Optional.
	//
	// One value per output group. Complete this field only if you are completing Base URL manifest A, and the downstream system has notified you that the child manifest files for pipeline 1 of all outputs are in a location different from the child manifest files for pipeline 0.
	BaseUrlManifest1 *string `field:"optional" json:"baseUrlManifest1" yaml:"baseUrlManifest1"`
	// A mapping of up to 4 captions channels to captions languages.
	//
	// This is meaningful only if captionLanguageSetting is set to "insert."
	CaptionLanguageMappings interface{} `field:"optional" json:"captionLanguageMappings" yaml:"captionLanguageMappings"`
	// Applies only to 608 embedded output captions.
	//
	// Insert: Include CLOSED-CAPTIONS lines in the manifest. Specify at least one language in the CC1 Language Code field. One CLOSED-CAPTION line is added for each Language Code that you specify. Make sure to specify the languages in the order in which they appear in the original source (if the source is embedded format) or the order of the captions selectors (if the source is other than embedded). Otherwise, languages in the manifest will not match properly with the output captions. None: Include the CLOSED-CAPTIONS=NONE line in the manifest. Omit: Omit any CLOSED-CAPTIONS line from the manifest.
	CaptionLanguageSetting *string `field:"optional" json:"captionLanguageSetting" yaml:"captionLanguageSetting"`
	// When set to "disabled," sets the #EXT-X-ALLOW-CACHE:no tag in the manifest, which prevents clients from saving media segments for later replay.
	ClientCache *string `field:"optional" json:"clientCache" yaml:"clientCache"`
	// The specification to use (RFC-6381 or the default RFC-4281) during m3u8 playlist generation.
	CodecSpecification *string `field:"optional" json:"codecSpecification" yaml:"codecSpecification"`
	// Used with encryptionType.
	//
	// This is a 128-bit, 16-byte hex value that is represented by a 32-character text string. If ivSource is set to "explicit," this parameter is required and is used as the IV for encryption.
	ConstantIv *string `field:"optional" json:"constantIv" yaml:"constantIv"`
	// A directory or HTTP destination for the HLS segments, manifest files, and encryption keys (if enabled).
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// Places segments in subdirectories.
	DirectoryStructure *string `field:"optional" json:"directoryStructure" yaml:"directoryStructure"`
	// Specifies whether to insert EXT-X-DISCONTINUITY tags in the HLS child manifests for this output group.
	//
	// Typically, choose Insert because these tags are required in the manifest (according to the HLS specification) and serve an important purpose.
	// Choose Never Insert only if the downstream system is doing real-time failover (without using the MediaLive automatic failover feature) and only if that downstream system has advised you to exclude the tags.
	DiscontinuityTags *string `field:"optional" json:"discontinuityTags" yaml:"discontinuityTags"`
	// Encrypts the segments with the specified encryption scheme.
	//
	// Exclude this parameter if you don't want encryption.
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// The parameters that control interactions with the CDN.
	HlsCdnSettings interface{} `field:"optional" json:"hlsCdnSettings" yaml:"hlsCdnSettings"`
	// State of HLS ID3 Segment Tagging.
	HlsId3SegmentTagging *string `field:"optional" json:"hlsId3SegmentTagging" yaml:"hlsId3SegmentTagging"`
	// DISABLED: Don't create an I-frame-only manifest, but do create the master and media manifests (according to the Output Selection field).
	//
	// STANDARD: Create an I-frame-only manifest for each output that contains video, as well as the other manifests (according to the Output Selection field). The I-frame manifest contains a #EXT-X-I-FRAMES-ONLY tag to indicate it is I-frame only, and one or more #EXT-X-BYTERANGE entries identifying the I-frame position. For example, #EXT-X-BYTERANGE:160364@1461888".
	IFrameOnlyPlaylists *string `field:"optional" json:"iFrameOnlyPlaylists" yaml:"iFrameOnlyPlaylists"`
	// Specifies whether to include the final (incomplete) segment in the media output when the pipeline stops producing output because of a channel stop, a channel pause or a loss of input to the pipeline.
	//
	// Auto means that MediaLive decides whether to include the final segment, depending on the channel class and the types of output groups.
	// Suppress means to never include the incomplete segment. We recommend you choose Auto and let MediaLive control the behavior.
	IncompleteSegmentBehavior *string `field:"optional" json:"incompleteSegmentBehavior" yaml:"incompleteSegmentBehavior"`
	// Applies only if the Mode field is LIVE.
	//
	// Specifies the maximum number of segments in the media manifest file. After this maximum, older segments are removed from the media manifest. This number must be less than or equal to the Keep Segments field.
	IndexNSegments *float64 `field:"optional" json:"indexNSegments" yaml:"indexNSegments"`
	// A parameter that controls output group behavior on an input loss.
	InputLossAction *string `field:"optional" json:"inputLossAction" yaml:"inputLossAction"`
	// Used with encryptionType.
	//
	// The IV (initialization vector) is a 128-bit number used in conjunction with the key for encrypting blocks. If set to "include," the IV is listed in the manifest. Otherwise, the IV is not in the manifest.
	IvInManifest *string `field:"optional" json:"ivInManifest" yaml:"ivInManifest"`
	// Used with encryptionType.
	//
	// The IV (initialization vector) is a 128-bit number used in conjunction with the key for encrypting blocks. If this setting is "followsSegmentNumber," it causes the IV to change every segment (to match the segment number). If this is set to "explicit," you must enter a constantIv value.
	IvSource *string `field:"optional" json:"ivSource" yaml:"ivSource"`
	// Applies only if the Mode field is LIVE.
	//
	// Specifies the number of media segments (.ts files) to retain in the destination directory.
	KeepSegments *float64 `field:"optional" json:"keepSegments" yaml:"keepSegments"`
	// Specifies how the key is represented in the resource identified by the URI.
	//
	// If the parameter is absent, an implicit value of "identity" is used. A reverse DNS string can also be specified.
	KeyFormat *string `field:"optional" json:"keyFormat" yaml:"keyFormat"`
	// Either a single positive integer version value or a slash-delimited list of version values (1/2/3).
	KeyFormatVersions *string `field:"optional" json:"keyFormatVersions" yaml:"keyFormatVersions"`
	// The key provider settings.
	KeyProviderSettings interface{} `field:"optional" json:"keyProviderSettings" yaml:"keyProviderSettings"`
	// When set to gzip, compresses HLS playlist.
	ManifestCompression *string `field:"optional" json:"manifestCompression" yaml:"manifestCompression"`
	// Indicates whether the output manifest should use a floating point or integer values for segment duration.
	ManifestDurationFormat *string `field:"optional" json:"manifestDurationFormat" yaml:"manifestDurationFormat"`
	// When set, minimumSegmentLength is enforced by looking ahead and back within the specified range for a nearby avail and extending the segment size if needed.
	MinSegmentLength *float64 `field:"optional" json:"minSegmentLength" yaml:"minSegmentLength"`
	// If "vod," all segments are indexed and kept permanently in the destination and manifest.
	//
	// If "live," only the number segments specified in keepSegments and indexNSegments are kept. Newer segments replace older segments, which might prevent players from rewinding all the way to the beginning of the channel. VOD mode uses HLS EXT-X-PLAYLIST-TYPE of EVENT while the channel is running, converting it to a "VOD" type manifest on completion of the stream.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// MANIFESTSANDSEGMENTS: Generates manifests (the master manifest, if applicable, and media manifests) for this output group.
	//
	// SEGMENTSONLY: Doesn't generate any manifests for this output group.
	OutputSelection *string `field:"optional" json:"outputSelection" yaml:"outputSelection"`
	// Includes or excludes the EXT-X-PROGRAM-DATE-TIME tag in .m3u8 manifest files. The value is calculated as follows: Either the program date and time are initialized using the input timecode source, or the time is initialized using the input timecode source and the date is initialized using the timestampOffset.
	ProgramDateTime *string `field:"optional" json:"programDateTime" yaml:"programDateTime"`
	// Specifies the algorithm used to drive the HLS EXT-X-PROGRAM-DATE-TIME clock.
	//
	// Options include: INITIALIZE_FROM_OUTPUT_TIMECODE: The PDT clock is initialized as a function of the first output timecode, then incremented by the EXTINF duration of each encoded segment. SYSTEM_CLOCK: The PDT clock is initialized as a function of the UTC wall clock, then incremented by the EXTINF duration of each encoded segment. If the PDT clock diverges from the wall clock by more than 500ms, it is resynchronized to the wall clock.
	ProgramDateTimeClock *string `field:"optional" json:"programDateTimeClock" yaml:"programDateTimeClock"`
	// The period of insertion of the EXT-X-PROGRAM-DATE-TIME entry, in seconds.
	ProgramDateTimePeriod *float64 `field:"optional" json:"programDateTimePeriod" yaml:"programDateTimePeriod"`
	// ENABLED: The master manifest (.m3u8 file) for each pipeline includes information about both pipelines: first its own media files, then the media files of the other pipeline. This feature allows a playout device that supports stale manifest detection to switch from one manifest to the other, when the current manifest seems to be stale. There are still two destinations and two master manifests, but both master manifests reference the media files from both pipelines. DISABLED: The master manifest (.m3u8 file) for each pipeline includes information about its own pipeline only. For an HLS output group with MediaPackage as the destination, the DISABLED behavior is always followed. MediaPackage regenerates the manifests it serves to players, so a redundant manifest from MediaLive is irrelevant.
	RedundantManifest *string `field:"optional" json:"redundantManifest" yaml:"redundantManifest"`
	// useInputSegmentation has been deprecated.
	//
	// The configured segment size is always used.
	SegmentationMode *string `field:"optional" json:"segmentationMode" yaml:"segmentationMode"`
	// The length of the MPEG-2 Transport Stream segments to create, in seconds.
	//
	// Note that segments will end on the next keyframe after this number of seconds, so the actual segment length might be longer.
	SegmentLength *float64 `field:"optional" json:"segmentLength" yaml:"segmentLength"`
	// The number of segments to write to a subdirectory before starting a new one.
	//
	// For this setting to have an effect, directoryStructure must be subdirectoryPerStream.
	SegmentsPerSubdirectory *float64 `field:"optional" json:"segmentsPerSubdirectory" yaml:"segmentsPerSubdirectory"`
	// The include or exclude RESOLUTION attribute for a video in the EXT-X-STREAM-INF tag of a variant manifest.
	StreamInfResolution *string `field:"optional" json:"streamInfResolution" yaml:"streamInfResolution"`
	// Indicates the ID3 frame that has the timecode.
	TimedMetadataId3Frame *string `field:"optional" json:"timedMetadataId3Frame" yaml:"timedMetadataId3Frame"`
	// The timed metadata interval, in seconds.
	TimedMetadataId3Period *float64 `field:"optional" json:"timedMetadataId3Period" yaml:"timedMetadataId3Period"`
	// Provides an extra millisecond delta offset to fine tune the timestamps.
	TimestampDeltaMilliseconds *float64 `field:"optional" json:"timestampDeltaMilliseconds" yaml:"timestampDeltaMilliseconds"`
	// SEGMENTEDFILES: Emits the program as segments -multiple .ts media files. SINGLEFILE: Applies only if the Mode field is VOD. Emits the program as a single .ts media file. The media manifest includes #EXT-X-BYTERANGE tags to index segments for playback. A typical use for this value is when sending the output to AWS Elemental MediaConvert, which can accept only a single media file. Playback while the channel is running is not guaranteed due to HTTP server caching.
	TsFileMode *string `field:"optional" json:"tsFileMode" yaml:"tsFileMode"`
}

// Information about how to connect to the upstream system.
//
// The parent of this entity is NetworkInputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsInputSettingsProperty := &hlsInputSettingsProperty{
//   	bandwidth: jsii.Number(123),
//   	bufferSegments: jsii.Number(123),
//   	retries: jsii.Number(123),
//   	retryInterval: jsii.Number(123),
//   	scte35Source: jsii.String("scte35Source"),
//   }
//
type CfnChannel_HlsInputSettingsProperty struct {
	// When specified, the HLS stream with the m3u8 bandwidth that most closely matches this value is chosen.
	//
	// Otherwise, the highest bandwidth stream in the m3u8 is chosen. The bitrate is specified in bits per second, as in an HLS manifest.
	Bandwidth *float64 `field:"optional" json:"bandwidth" yaml:"bandwidth"`
	// When specified, reading of the HLS input begins this many buffer segments from the end (most recently written segment).
	//
	// When not specified, the HLS input begins with the first segment specified in the m3u8.
	BufferSegments *float64 `field:"optional" json:"bufferSegments" yaml:"bufferSegments"`
	// The number of consecutive times that attempts to read a manifest or segment must fail before the input is considered unavailable.
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// The number of seconds between retries when an attempt to read a manifest or segment fails.
	RetryInterval *float64 `field:"optional" json:"retryInterval" yaml:"retryInterval"`
	// Identifies the source for the SCTE-35 messages that MediaLive will ingest.
	//
	// Messages can be ingested from the content segments (in the stream) or from tags in the playlist (the HLS manifest). MediaLive ignores SCTE-35 information in the source that is not selected.
	Scte35Source *string `field:"optional" json:"scte35Source" yaml:"scte35Source"`
}

// The configuration of a MediaStore container as the destination for an HLS output.
//
// The parent of this entity is HlsCdnSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsMediaStoreSettingsProperty := &hlsMediaStoreSettingsProperty{
//   	connectionRetryInterval: jsii.Number(123),
//   	filecacheDuration: jsii.Number(123),
//   	mediaStoreStorageClass: jsii.String("mediaStoreStorageClass"),
//   	numRetries: jsii.Number(123),
//   	restartDelay: jsii.Number(123),
//   }
//
type CfnChannel_HlsMediaStoreSettingsProperty struct {
	// The number of seconds to wait before retrying a connection to the CDN if the connection is lost.
	ConnectionRetryInterval *float64 `field:"optional" json:"connectionRetryInterval" yaml:"connectionRetryInterval"`
	// The size, in seconds, of the file cache for streaming outputs.
	FilecacheDuration *float64 `field:"optional" json:"filecacheDuration" yaml:"filecacheDuration"`
	// When set to temporal, output files are stored in non-persistent memory for faster reading and writing.
	MediaStoreStorageClass *string `field:"optional" json:"mediaStoreStorageClass" yaml:"mediaStoreStorageClass"`
	// The number of retry attempts that are made before the channel is put into an error state.
	NumRetries *float64 `field:"optional" json:"numRetries" yaml:"numRetries"`
	// If a streaming output fails, the number of seconds to wait until a restart is initiated.
	//
	// A value of 0 means never restart.
	RestartDelay *float64 `field:"optional" json:"restartDelay" yaml:"restartDelay"`
}

// The settings for an HLS output.
//
// The parent of this entity is OutputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsOutputSettingsProperty := &hlsOutputSettingsProperty{
//   	h265PackagingType: jsii.String("h265PackagingType"),
//   	hlsSettings: &hlsSettingsProperty{
//   		audioOnlyHlsSettings: &audioOnlyHlsSettingsProperty{
//   			audioGroupId: jsii.String("audioGroupId"),
//   			audioOnlyImage: &inputLocationProperty{
//   				passwordParam: jsii.String("passwordParam"),
//   				uri: jsii.String("uri"),
//   				username: jsii.String("username"),
//   			},
//   			audioTrackType: jsii.String("audioTrackType"),
//   			segmentType: jsii.String("segmentType"),
//   		},
//   		fmp4HlsSettings: &fmp4HlsSettingsProperty{
//   			audioRenditionSets: jsii.String("audioRenditionSets"),
//   			nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   			timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   		},
//   		frameCaptureHlsSettings: &frameCaptureHlsSettingsProperty{
//   		},
//   		standardHlsSettings: &standardHlsSettingsProperty{
//   			audioRenditionSets: jsii.String("audioRenditionSets"),
//   			m3U8Settings: &m3u8SettingsProperty{
//   				audioFramesPerPes: jsii.Number(123),
//   				audioPids: jsii.String("audioPids"),
//   				ecmPid: jsii.String("ecmPid"),
//   				nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   				patInterval: jsii.Number(123),
//   				pcrControl: jsii.String("pcrControl"),
//   				pcrPeriod: jsii.Number(123),
//   				pcrPid: jsii.String("pcrPid"),
//   				pmtInterval: jsii.Number(123),
//   				pmtPid: jsii.String("pmtPid"),
//   				programNum: jsii.Number(123),
//   				scte35Behavior: jsii.String("scte35Behavior"),
//   				scte35Pid: jsii.String("scte35Pid"),
//   				timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   				timedMetadataPid: jsii.String("timedMetadataPid"),
//   				transportStreamId: jsii.Number(123),
//   				videoPid: jsii.String("videoPid"),
//   			},
//   		},
//   	},
//   	nameModifier: jsii.String("nameModifier"),
//   	segmentModifier: jsii.String("segmentModifier"),
//   }
//
type CfnChannel_HlsOutputSettingsProperty struct {
	// Only applicable when this output is referencing an H.265 video description. Specifies whether MP4 segments should be packaged as HEV1 or HVC1.
	H265PackagingType *string `field:"optional" json:"h265PackagingType" yaml:"h265PackagingType"`
	// The settings regarding the underlying stream.
	//
	// These settings are different for audio-only outputs.
	HlsSettings interface{} `field:"optional" json:"hlsSettings" yaml:"hlsSettings"`
	// A string that is concatenated to the end of the destination file name.
	//
	// Accepts \"Format Identifiers\":#formatIdentifierParameters.
	NameModifier *string `field:"optional" json:"nameModifier" yaml:"nameModifier"`
	// A string that is concatenated to the end of segment file names.
	SegmentModifier *string `field:"optional" json:"segmentModifier" yaml:"segmentModifier"`
}

// Sets up Amazon S3 as the destination for this HLS output.
//
// The parent of this entity is HlsCdnSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsS3SettingsProperty := &hlsS3SettingsProperty{
//   	cannedAcl: jsii.String("cannedAcl"),
//   }
//
type CfnChannel_HlsS3SettingsProperty struct {
	// Specify the canned ACL to apply to each S3 request.
	//
	// Defaults to none.
	CannedAcl *string `field:"optional" json:"cannedAcl" yaml:"cannedAcl"`
}

// The settings for an HLS output.
//
// The parent of this entity is HlsOutputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsSettingsProperty := &hlsSettingsProperty{
//   	audioOnlyHlsSettings: &audioOnlyHlsSettingsProperty{
//   		audioGroupId: jsii.String("audioGroupId"),
//   		audioOnlyImage: &inputLocationProperty{
//   			passwordParam: jsii.String("passwordParam"),
//   			uri: jsii.String("uri"),
//   			username: jsii.String("username"),
//   		},
//   		audioTrackType: jsii.String("audioTrackType"),
//   		segmentType: jsii.String("segmentType"),
//   	},
//   	fmp4HlsSettings: &fmp4HlsSettingsProperty{
//   		audioRenditionSets: jsii.String("audioRenditionSets"),
//   		nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   		timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   	},
//   	frameCaptureHlsSettings: &frameCaptureHlsSettingsProperty{
//   	},
//   	standardHlsSettings: &standardHlsSettingsProperty{
//   		audioRenditionSets: jsii.String("audioRenditionSets"),
//   		m3U8Settings: &m3u8SettingsProperty{
//   			audioFramesPerPes: jsii.Number(123),
//   			audioPids: jsii.String("audioPids"),
//   			ecmPid: jsii.String("ecmPid"),
//   			nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   			patInterval: jsii.Number(123),
//   			pcrControl: jsii.String("pcrControl"),
//   			pcrPeriod: jsii.Number(123),
//   			pcrPid: jsii.String("pcrPid"),
//   			pmtInterval: jsii.Number(123),
//   			pmtPid: jsii.String("pmtPid"),
//   			programNum: jsii.Number(123),
//   			scte35Behavior: jsii.String("scte35Behavior"),
//   			scte35Pid: jsii.String("scte35Pid"),
//   			timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   			timedMetadataPid: jsii.String("timedMetadataPid"),
//   			transportStreamId: jsii.Number(123),
//   			videoPid: jsii.String("videoPid"),
//   		},
//   	},
//   }
//
type CfnChannel_HlsSettingsProperty struct {
	// The settings for an audio-only output.
	AudioOnlyHlsSettings interface{} `field:"optional" json:"audioOnlyHlsSettings" yaml:"audioOnlyHlsSettings"`
	// The settings for an fMP4 container.
	Fmp4HlsSettings interface{} `field:"optional" json:"fmp4HlsSettings" yaml:"fmp4HlsSettings"`
	// Settings for a frame capture output in an HLS output group.
	FrameCaptureHlsSettings interface{} `field:"optional" json:"frameCaptureHlsSettings" yaml:"frameCaptureHlsSettings"`
	// The settings for a standard output (an output that is not audio-only).
	StandardHlsSettings interface{} `field:"optional" json:"standardHlsSettings" yaml:"standardHlsSettings"`
}

// The configuration of a WebDav server as the downstream system for an HLS output.
//
// The parent of this entity is HlsCdnSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsWebdavSettingsProperty := &hlsWebdavSettingsProperty{
//   	connectionRetryInterval: jsii.Number(123),
//   	filecacheDuration: jsii.Number(123),
//   	httpTransferMode: jsii.String("httpTransferMode"),
//   	numRetries: jsii.Number(123),
//   	restartDelay: jsii.Number(123),
//   }
//
type CfnChannel_HlsWebdavSettingsProperty struct {
	// The number of seconds to wait before retrying a connection to the CDN if the connection is lost.
	ConnectionRetryInterval *float64 `field:"optional" json:"connectionRetryInterval" yaml:"connectionRetryInterval"`
	// The size, in seconds, of the file cache for streaming outputs.
	FilecacheDuration *float64 `field:"optional" json:"filecacheDuration" yaml:"filecacheDuration"`
	// Specifies whether to use chunked transfer encoding to WebDAV.
	HttpTransferMode *string `field:"optional" json:"httpTransferMode" yaml:"httpTransferMode"`
	// The number of retry attempts that are made before the channel is put into an error state.
	NumRetries *float64 `field:"optional" json:"numRetries" yaml:"numRetries"`
	// If a streaming output fails, the number of seconds to wait until a restart is initiated.
	//
	// A value of 0 means never restart.
	RestartDelay *float64 `field:"optional" json:"restartDelay" yaml:"restartDelay"`
}

// Settings to configure the motion graphics overlay to use an HTML asset.
//
// The parent of this entity is MotionGraphicsSetting.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   htmlMotionGraphicsSettingsProperty := &htmlMotionGraphicsSettingsProperty{
//   }
//
type CfnChannel_HtmlMotionGraphicsSettingsProperty struct {
}

// An input to attach to this channel.
//
// This entity is at the top level in the channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputAttachmentProperty := &inputAttachmentProperty{
//   	automaticInputFailoverSettings: &automaticInputFailoverSettingsProperty{
//   		errorClearTimeMsec: jsii.Number(123),
//   		failoverConditions: []interface{}{
//   			&failoverConditionProperty{
//   				failoverConditionSettings: &failoverConditionSettingsProperty{
//   					audioSilenceSettings: &audioSilenceFailoverSettingsProperty{
//   						audioSelectorName: jsii.String("audioSelectorName"),
//   						audioSilenceThresholdMsec: jsii.Number(123),
//   					},
//   					inputLossSettings: &inputLossFailoverSettingsProperty{
//   						inputLossThresholdMsec: jsii.Number(123),
//   					},
//   					videoBlackSettings: &videoBlackFailoverSettingsProperty{
//   						blackDetectThreshold: jsii.Number(123),
//   						videoBlackThresholdMsec: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   		inputPreference: jsii.String("inputPreference"),
//   		secondaryInputId: jsii.String("secondaryInputId"),
//   	},
//   	inputAttachmentName: jsii.String("inputAttachmentName"),
//   	inputId: jsii.String("inputId"),
//   	inputSettings: &inputSettingsProperty{
//   		audioSelectors: []interface{}{
//   			&audioSelectorProperty{
//   				name: jsii.String("name"),
//   				selectorSettings: &audioSelectorSettingsProperty{
//   					audioHlsRenditionSelection: &audioHlsRenditionSelectionProperty{
//   						groupId: jsii.String("groupId"),
//   						name: jsii.String("name"),
//   					},
//   					audioLanguageSelection: &audioLanguageSelectionProperty{
//   						languageCode: jsii.String("languageCode"),
//   						languageSelectionPolicy: jsii.String("languageSelectionPolicy"),
//   					},
//   					audioPidSelection: &audioPidSelectionProperty{
//   						pid: jsii.Number(123),
//   					},
//   					audioTrackSelection: &audioTrackSelectionProperty{
//   						tracks: []interface{}{
//   							&audioTrackProperty{
//   								track: jsii.Number(123),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		captionSelectors: []interface{}{
//   			&captionSelectorProperty{
//   				languageCode: jsii.String("languageCode"),
//   				name: jsii.String("name"),
//   				selectorSettings: &captionSelectorSettingsProperty{
//   					ancillarySourceSettings: &ancillarySourceSettingsProperty{
//   						sourceAncillaryChannelNumber: jsii.Number(123),
//   					},
//   					aribSourceSettings: &aribSourceSettingsProperty{
//   					},
//   					dvbSubSourceSettings: &dvbSubSourceSettingsProperty{
//   						ocrLanguage: jsii.String("ocrLanguage"),
//   						pid: jsii.Number(123),
//   					},
//   					embeddedSourceSettings: &embeddedSourceSettingsProperty{
//   						convert608To708: jsii.String("convert608To708"),
//   						scte20Detection: jsii.String("scte20Detection"),
//   						source608ChannelNumber: jsii.Number(123),
//   						source608TrackNumber: jsii.Number(123),
//   					},
//   					scte20SourceSettings: &scte20SourceSettingsProperty{
//   						convert608To708: jsii.String("convert608To708"),
//   						source608ChannelNumber: jsii.Number(123),
//   					},
//   					scte27SourceSettings: &scte27SourceSettingsProperty{
//   						ocrLanguage: jsii.String("ocrLanguage"),
//   						pid: jsii.Number(123),
//   					},
//   					teletextSourceSettings: &teletextSourceSettingsProperty{
//   						outputRectangle: &captionRectangleProperty{
//   							height: jsii.Number(123),
//   							leftOffset: jsii.Number(123),
//   							topOffset: jsii.Number(123),
//   							width: jsii.Number(123),
//   						},
//   						pageNumber: jsii.String("pageNumber"),
//   					},
//   				},
//   			},
//   		},
//   		deblockFilter: jsii.String("deblockFilter"),
//   		denoiseFilter: jsii.String("denoiseFilter"),
//   		filterStrength: jsii.Number(123),
//   		inputFilter: jsii.String("inputFilter"),
//   		networkInputSettings: &networkInputSettingsProperty{
//   			hlsInputSettings: &hlsInputSettingsProperty{
//   				bandwidth: jsii.Number(123),
//   				bufferSegments: jsii.Number(123),
//   				retries: jsii.Number(123),
//   				retryInterval: jsii.Number(123),
//   				scte35Source: jsii.String("scte35Source"),
//   			},
//   			serverValidation: jsii.String("serverValidation"),
//   		},
//   		scte35Pid: jsii.Number(123),
//   		smpte2038DataPreference: jsii.String("smpte2038DataPreference"),
//   		sourceEndBehavior: jsii.String("sourceEndBehavior"),
//   		videoSelector: &videoSelectorProperty{
//   			colorSpace: jsii.String("colorSpace"),
//   			colorSpaceSettings: &videoSelectorColorSpaceSettingsProperty{
//   				hdr10Settings: &hdr10SettingsProperty{
//   					maxCll: jsii.Number(123),
//   					maxFall: jsii.Number(123),
//   				},
//   			},
//   			colorSpaceUsage: jsii.String("colorSpaceUsage"),
//   			selectorSettings: &videoSelectorSettingsProperty{
//   				videoSelectorPid: &videoSelectorPidProperty{
//   					pid: jsii.Number(123),
//   				},
//   				videoSelectorProgramId: &videoSelectorProgramIdProperty{
//   					programId: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnChannel_InputAttachmentProperty struct {
	// Settings to implement automatic input failover in this input.
	AutomaticInputFailoverSettings interface{} `field:"optional" json:"automaticInputFailoverSettings" yaml:"automaticInputFailoverSettings"`
	// A name for the attachment.
	//
	// This is required if you want to use this input in an input switch action.
	InputAttachmentName *string `field:"optional" json:"inputAttachmentName" yaml:"inputAttachmentName"`
	// The ID of the input to attach.
	InputId *string `field:"optional" json:"inputId" yaml:"inputId"`
	// Information about the content to extract from the input and about the general handling of the content.
	InputSettings interface{} `field:"optional" json:"inputSettings" yaml:"inputSettings"`
}

// The setting to remix the audio.
//
// The parent of this entity is AudioChannelMappings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputChannelLevelProperty := &inputChannelLevelProperty{
//   	gain: jsii.Number(123),
//   	inputChannel: jsii.Number(123),
//   }
//
type CfnChannel_InputChannelLevelProperty struct {
	// The remixing value.
	//
	// Units are in dB, and acceptable values are within the range from -60 (mute) to 6 dB.
	Gain *float64 `field:"optional" json:"gain" yaml:"gain"`
	// The index of the input channel that is used as a source.
	InputChannel *float64 `field:"optional" json:"inputChannel" yaml:"inputChannel"`
}

// The input location.
//
// The parent of this entity is InputLossBehavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputLocationProperty := &inputLocationProperty{
//   	passwordParam: jsii.String("passwordParam"),
//   	uri: jsii.String("uri"),
//   	username: jsii.String("username"),
//   }
//
type CfnChannel_InputLocationProperty struct {
	// The password parameter that holds the password for accessing the downstream system.
	//
	// This applies only if the downstream system requires credentials.
	PasswordParam *string `field:"optional" json:"passwordParam" yaml:"passwordParam"`
	// The URI should be a path to a file that is accessible to the Live system (for example, an http:// URI) depending on the output type.
	//
	// For example, an RTMP destination should have a URI similar to rtmp://fmsserver/live.
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
	// The user name to connect to the downstream system.
	//
	// This applies only if the downstream system requires credentials.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

// The configuration of channel behavior when the input is lost.
//
// The parent of this entity is GlobalConfiguration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputLossBehaviorProperty := &inputLossBehaviorProperty{
//   	blackFrameMsec: jsii.Number(123),
//   	inputLossImageColor: jsii.String("inputLossImageColor"),
//   	inputLossImageSlate: &inputLocationProperty{
//   		passwordParam: jsii.String("passwordParam"),
//   		uri: jsii.String("uri"),
//   		username: jsii.String("username"),
//   	},
//   	inputLossImageType: jsii.String("inputLossImageType"),
//   	repeatFrameMsec: jsii.Number(123),
//   }
//
type CfnChannel_InputLossBehaviorProperty struct {
	// On input loss, the number of milliseconds to substitute black into the output before switching to the frame specified by inputLossImageType.
	//
	// A value x, where 0 <= x <= 1,000,000 and a value of 1,000,000, is interpreted as infinite.
	BlackFrameMsec *float64 `field:"optional" json:"blackFrameMsec" yaml:"blackFrameMsec"`
	// When the input loss image type is "color," this field specifies the color to use.
	//
	// Value: 6 hex characters that represent the values of RGB.
	InputLossImageColor *string `field:"optional" json:"inputLossImageColor" yaml:"inputLossImageColor"`
	// When the input loss image type is "slate," these fields specify the parameters for accessing the slate.
	InputLossImageSlate interface{} `field:"optional" json:"inputLossImageSlate" yaml:"inputLossImageSlate"`
	// Indicates whether to substitute a solid color or a slate into the output after the input loss exceeds blackFrameMsec.
	InputLossImageType *string `field:"optional" json:"inputLossImageType" yaml:"inputLossImageType"`
	// On input loss, the number of milliseconds to repeat the previous picture before substituting black into the output.
	//
	// A value x, where 0 <= x <= 1,000,000 and a value of 1,000,000, is interpreted as infinite.
	RepeatFrameMsec *float64 `field:"optional" json:"repeatFrameMsec" yaml:"repeatFrameMsec"`
}

// MediaLive will perform a failover if content is not detected in this input for the specified period.
//
// The parent of this entity is FailoverConditionSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputLossFailoverSettingsProperty := &inputLossFailoverSettingsProperty{
//   	inputLossThresholdMsec: jsii.Number(123),
//   }
//
type CfnChannel_InputLossFailoverSettingsProperty struct {
	// The amount of time (in milliseconds) that no input is detected.
	//
	// After that time, an input failover will occur.
	InputLossThresholdMsec *float64 `field:"optional" json:"inputLossThresholdMsec" yaml:"inputLossThresholdMsec"`
}

// Information about extracting content from the input and about handling the content.
//
// The parent of this entity is InputAttachment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputSettingsProperty := &inputSettingsProperty{
//   	audioSelectors: []interface{}{
//   		&audioSelectorProperty{
//   			name: jsii.String("name"),
//   			selectorSettings: &audioSelectorSettingsProperty{
//   				audioHlsRenditionSelection: &audioHlsRenditionSelectionProperty{
//   					groupId: jsii.String("groupId"),
//   					name: jsii.String("name"),
//   				},
//   				audioLanguageSelection: &audioLanguageSelectionProperty{
//   					languageCode: jsii.String("languageCode"),
//   					languageSelectionPolicy: jsii.String("languageSelectionPolicy"),
//   				},
//   				audioPidSelection: &audioPidSelectionProperty{
//   					pid: jsii.Number(123),
//   				},
//   				audioTrackSelection: &audioTrackSelectionProperty{
//   					tracks: []interface{}{
//   						&audioTrackProperty{
//   							track: jsii.Number(123),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	captionSelectors: []interface{}{
//   		&captionSelectorProperty{
//   			languageCode: jsii.String("languageCode"),
//   			name: jsii.String("name"),
//   			selectorSettings: &captionSelectorSettingsProperty{
//   				ancillarySourceSettings: &ancillarySourceSettingsProperty{
//   					sourceAncillaryChannelNumber: jsii.Number(123),
//   				},
//   				aribSourceSettings: &aribSourceSettingsProperty{
//   				},
//   				dvbSubSourceSettings: &dvbSubSourceSettingsProperty{
//   					ocrLanguage: jsii.String("ocrLanguage"),
//   					pid: jsii.Number(123),
//   				},
//   				embeddedSourceSettings: &embeddedSourceSettingsProperty{
//   					convert608To708: jsii.String("convert608To708"),
//   					scte20Detection: jsii.String("scte20Detection"),
//   					source608ChannelNumber: jsii.Number(123),
//   					source608TrackNumber: jsii.Number(123),
//   				},
//   				scte20SourceSettings: &scte20SourceSettingsProperty{
//   					convert608To708: jsii.String("convert608To708"),
//   					source608ChannelNumber: jsii.Number(123),
//   				},
//   				scte27SourceSettings: &scte27SourceSettingsProperty{
//   					ocrLanguage: jsii.String("ocrLanguage"),
//   					pid: jsii.Number(123),
//   				},
//   				teletextSourceSettings: &teletextSourceSettingsProperty{
//   					outputRectangle: &captionRectangleProperty{
//   						height: jsii.Number(123),
//   						leftOffset: jsii.Number(123),
//   						topOffset: jsii.Number(123),
//   						width: jsii.Number(123),
//   					},
//   					pageNumber: jsii.String("pageNumber"),
//   				},
//   			},
//   		},
//   	},
//   	deblockFilter: jsii.String("deblockFilter"),
//   	denoiseFilter: jsii.String("denoiseFilter"),
//   	filterStrength: jsii.Number(123),
//   	inputFilter: jsii.String("inputFilter"),
//   	networkInputSettings: &networkInputSettingsProperty{
//   		hlsInputSettings: &hlsInputSettingsProperty{
//   			bandwidth: jsii.Number(123),
//   			bufferSegments: jsii.Number(123),
//   			retries: jsii.Number(123),
//   			retryInterval: jsii.Number(123),
//   			scte35Source: jsii.String("scte35Source"),
//   		},
//   		serverValidation: jsii.String("serverValidation"),
//   	},
//   	scte35Pid: jsii.Number(123),
//   	smpte2038DataPreference: jsii.String("smpte2038DataPreference"),
//   	sourceEndBehavior: jsii.String("sourceEndBehavior"),
//   	videoSelector: &videoSelectorProperty{
//   		colorSpace: jsii.String("colorSpace"),
//   		colorSpaceSettings: &videoSelectorColorSpaceSettingsProperty{
//   			hdr10Settings: &hdr10SettingsProperty{
//   				maxCll: jsii.Number(123),
//   				maxFall: jsii.Number(123),
//   			},
//   		},
//   		colorSpaceUsage: jsii.String("colorSpaceUsage"),
//   		selectorSettings: &videoSelectorSettingsProperty{
//   			videoSelectorPid: &videoSelectorPidProperty{
//   				pid: jsii.Number(123),
//   			},
//   			videoSelectorProgramId: &videoSelectorProgramIdProperty{
//   				programId: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
type CfnChannel_InputSettingsProperty struct {
	// Information about the specific audio to extract from the input.
	//
	// The parent of this entity is InputSettings.
	AudioSelectors interface{} `field:"optional" json:"audioSelectors" yaml:"audioSelectors"`
	// Information about the specific captions to extract from the input.
	CaptionSelectors interface{} `field:"optional" json:"captionSelectors" yaml:"captionSelectors"`
	// Enables or disables the deblock filter when filtering.
	DeblockFilter *string `field:"optional" json:"deblockFilter" yaml:"deblockFilter"`
	// Enables or disables the denoise filter when filtering.
	DenoiseFilter *string `field:"optional" json:"denoiseFilter" yaml:"denoiseFilter"`
	// Adjusts the magnitude of filtering from 1 (minimal) to 5 (strongest).
	FilterStrength *float64 `field:"optional" json:"filterStrength" yaml:"filterStrength"`
	// Turns on the filter for this input.
	//
	// MPEG-2 inputs have the deblocking filter enabled by default. 1) auto - filtering is applied depending on input type/quality 2) disabled - no filtering is applied to the input 3) forced - filtering is applied regardless of the input type.
	InputFilter *string `field:"optional" json:"inputFilter" yaml:"inputFilter"`
	// Information about how to connect to the upstream system.
	NetworkInputSettings interface{} `field:"optional" json:"networkInputSettings" yaml:"networkInputSettings"`
	// `CfnChannel.InputSettingsProperty.Scte35Pid`.
	Scte35Pid *float64 `field:"optional" json:"scte35Pid" yaml:"scte35Pid"`
	// Specifies whether to extract applicable ancillary data from a SMPTE-2038 source in this input.
	//
	// Applicable data types are captions, timecode, AFD, and SCTE-104 messages.
	// - PREFER: Extract from SMPTE-2038 if present in this input, otherwise extract from another source (if any).
	// - IGNORE: Never extract any ancillary data from SMPTE-2038.
	Smpte2038DataPreference *string `field:"optional" json:"smpte2038DataPreference" yaml:"smpte2038DataPreference"`
	// The loop input if it is a file.
	SourceEndBehavior *string `field:"optional" json:"sourceEndBehavior" yaml:"sourceEndBehavior"`
	// Information about one video to extract from the input.
	VideoSelector interface{} `field:"optional" json:"videoSelector" yaml:"videoSelector"`
}

// The input specification for this channel.
//
// It specifies the key characteristics of the inputs for this channel: the maximum bitrate, the resolution, and the codec.
//
// This entity is at the top level in the channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputSpecificationProperty := &inputSpecificationProperty{
//   	codec: jsii.String("codec"),
//   	maximumBitrate: jsii.String("maximumBitrate"),
//   	resolution: jsii.String("resolution"),
//   }
//
type CfnChannel_InputSpecificationProperty struct {
	// The codec to include in the input specification for this channel.
	Codec *string `field:"optional" json:"codec" yaml:"codec"`
	// The maximum input bitrate for any input attached to this channel.
	MaximumBitrate *string `field:"optional" json:"maximumBitrate" yaml:"maximumBitrate"`
	// The resolution for any input attached to this channel.
	Resolution *string `field:"optional" json:"resolution" yaml:"resolution"`
}

// The configuration of key provider settings.
//
// The parent of this entity is HlsGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyProviderSettingsProperty := &keyProviderSettingsProperty{
//   	staticKeySettings: &staticKeySettingsProperty{
//   		keyProviderServer: &inputLocationProperty{
//   			passwordParam: jsii.String("passwordParam"),
//   			uri: jsii.String("uri"),
//   			username: jsii.String("username"),
//   		},
//   		staticKeyValue: jsii.String("staticKeyValue"),
//   	},
//   }
//
type CfnChannel_KeyProviderSettingsProperty struct {
	// The configuration of static key settings.
	StaticKeySettings interface{} `field:"optional" json:"staticKeySettings" yaml:"staticKeySettings"`
}

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

// Settings for the M3U8 container.
//
// The parent of this entity is StandardHlsSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   m3u8SettingsProperty := &m3u8SettingsProperty{
//   	audioFramesPerPes: jsii.Number(123),
//   	audioPids: jsii.String("audioPids"),
//   	ecmPid: jsii.String("ecmPid"),
//   	nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   	patInterval: jsii.Number(123),
//   	pcrControl: jsii.String("pcrControl"),
//   	pcrPeriod: jsii.Number(123),
//   	pcrPid: jsii.String("pcrPid"),
//   	pmtInterval: jsii.Number(123),
//   	pmtPid: jsii.String("pmtPid"),
//   	programNum: jsii.Number(123),
//   	scte35Behavior: jsii.String("scte35Behavior"),
//   	scte35Pid: jsii.String("scte35Pid"),
//   	timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   	timedMetadataPid: jsii.String("timedMetadataPid"),
//   	transportStreamId: jsii.Number(123),
//   	videoPid: jsii.String("videoPid"),
//   }
//
type CfnChannel_M3u8SettingsProperty struct {
	// The number of audio frames to insert for each PES packet.
	AudioFramesPerPes *float64 `field:"optional" json:"audioFramesPerPes" yaml:"audioFramesPerPes"`
	// The PID of the elementary audio streams in the transport stream.
	//
	// Multiple values are accepted, and can be entered in ranges or by comma separation. You can enter the value as a decimal or hexadecimal value.
	AudioPids *string `field:"optional" json:"audioPids" yaml:"audioPids"`
	// This parameter is unused and deprecated.
	EcmPid *string `field:"optional" json:"ecmPid" yaml:"ecmPid"`
	// If set to passthrough, Nielsen inaudible tones for media tracking will be detected in the input audio and an equivalent ID3 tag will be inserted in the output.
	NielsenId3Behavior *string `field:"optional" json:"nielsenId3Behavior" yaml:"nielsenId3Behavior"`
	// The number of milliseconds between instances of this table in the output transport stream.
	//
	// A value of \"0\" writes out the PMT once per segment file.
	PatInterval *float64 `field:"optional" json:"patInterval" yaml:"patInterval"`
	// When set to pcrEveryPesPacket, a Program Clock Reference value is inserted for every Packetized Elementary Stream (PES) header.
	//
	// This parameter is effective only when the PCR PID is the same as the video or audio elementary stream.
	PcrControl *string `field:"optional" json:"pcrControl" yaml:"pcrControl"`
	// The maximum time, in milliseconds, between Program Clock References (PCRs) inserted into the transport stream.
	PcrPeriod *float64 `field:"optional" json:"pcrPeriod" yaml:"pcrPeriod"`
	// The PID of the Program Clock Reference (PCR) in the transport stream.
	//
	// When no value is given, MediaLive assigns the same value as the video PID. You can enter the value as a decimal or hexadecimal value.
	PcrPid *string `field:"optional" json:"pcrPid" yaml:"pcrPid"`
	// The number of milliseconds between instances of this table in the output transport stream.
	//
	// A value of \"0\" writes out the PMT once per segment file.
	PmtInterval *float64 `field:"optional" json:"pmtInterval" yaml:"pmtInterval"`
	// The PID for the Program Map Table (PMT) in the transport stream.
	//
	// You can enter the value as a decimal or hexadecimal value.
	PmtPid *string `field:"optional" json:"pmtPid" yaml:"pmtPid"`
	// The value of the program number field in the Program Map Table (PMT).
	ProgramNum *float64 `field:"optional" json:"programNum" yaml:"programNum"`
	// If set to passthrough, passes any SCTE-35 signals from the input source to this output.
	Scte35Behavior *string `field:"optional" json:"scte35Behavior" yaml:"scte35Behavior"`
	// The PID of the SCTE-35 stream in the transport stream.
	//
	// You can enter the value as a decimal or hexadecimal value.
	Scte35Pid *string `field:"optional" json:"scte35Pid" yaml:"scte35Pid"`
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
	// You can enter the value as a decimal or hexadecimal value.
	VideoPid *string `field:"optional" json:"videoPid" yaml:"videoPid"`
}

// The settings for the MediaPackage group.
//
// The parent of this entity is OutputGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mediaPackageGroupSettingsProperty := &mediaPackageGroupSettingsProperty{
//   	destination: &outputLocationRefProperty{
//   		destinationRefId: jsii.String("destinationRefId"),
//   	},
//   }
//
type CfnChannel_MediaPackageGroupSettingsProperty struct {
	// The MediaPackage channel destination.
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
}

// Destination settings for a MediaPackage output.
//
// The parent of this entity is OutputDestination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mediaPackageOutputDestinationSettingsProperty := &mediaPackageOutputDestinationSettingsProperty{
//   	channelId: jsii.String("channelId"),
//   }
//
type CfnChannel_MediaPackageOutputDestinationSettingsProperty struct {
	// The ID of the channel in MediaPackage that is the destination for this output group.
	//
	// You don't need to specify the individual inputs in MediaPackage; MediaLive handles the connection of the two MediaLive pipelines to the two MediaPackage inputs. The MediaPackage channel and MediaLive channel must be in the same Region.
	ChannelId *string `field:"optional" json:"channelId" yaml:"channelId"`
}

// The settings for a MediaPackage output.
//
// The parent of this entity is OutputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mediaPackageOutputSettingsProperty := &mediaPackageOutputSettingsProperty{
//   }
//
type CfnChannel_MediaPackageOutputSettingsProperty struct {
}

// Settings to enable and configure the motion graphics overlay feature in the channel.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   motionGraphicsConfigurationProperty := &motionGraphicsConfigurationProperty{
//   	motionGraphicsInsertion: jsii.String("motionGraphicsInsertion"),
//   	motionGraphicsSettings: &motionGraphicsSettingsProperty{
//   		htmlMotionGraphicsSettings: &htmlMotionGraphicsSettingsProperty{
//   		},
//   	},
//   }
//
type CfnChannel_MotionGraphicsConfigurationProperty struct {
	// Enables or disables the motion graphics overlay feature in the channel.
	MotionGraphicsInsertion *string `field:"optional" json:"motionGraphicsInsertion" yaml:"motionGraphicsInsertion"`
	// Settings to enable and configure the motion graphics overlay feature in the channel.
	MotionGraphicsSettings interface{} `field:"optional" json:"motionGraphicsSettings" yaml:"motionGraphicsSettings"`
}

// Settings to enable and configure the motion graphics overlay feature in the channel.
//
// The parent of this entity is MotionGraphicsConfiguration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   motionGraphicsSettingsProperty := &motionGraphicsSettingsProperty{
//   	htmlMotionGraphicsSettings: &htmlMotionGraphicsSettingsProperty{
//   	},
//   }
//
type CfnChannel_MotionGraphicsSettingsProperty struct {
	// Settings to configure the motion graphics overlay to use an HTML asset.
	HtmlMotionGraphicsSettings interface{} `field:"optional" json:"htmlMotionGraphicsSettings" yaml:"htmlMotionGraphicsSettings"`
}

// The configuration for this MP2 audio.
//
// The parent of this entity is AudioCodecSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mp2SettingsProperty := &mp2SettingsProperty{
//   	bitrate: jsii.Number(123),
//   	codingMode: jsii.String("codingMode"),
//   	sampleRate: jsii.Number(123),
//   }
//
type CfnChannel_Mp2SettingsProperty struct {
	// The average bitrate in bits/second.
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// The MPEG2 Audio coding mode.
	//
	// Valid values are codingMode10 (for mono) or codingMode20 (for stereo).
	CodingMode *string `field:"optional" json:"codingMode" yaml:"codingMode"`
	// The sample rate in Hz.
	SampleRate *float64 `field:"optional" json:"sampleRate" yaml:"sampleRate"`
}

// Settings to configure video filters that apply to the MPEG-2 codec.
//
// The parent of this entity is Mpeg2FilterSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mpeg2FilterSettingsProperty := &mpeg2FilterSettingsProperty{
//   	temporalFilterSettings: &temporalFilterSettingsProperty{
//   		postFilterSharpening: jsii.String("postFilterSharpening"),
//   		strength: jsii.String("strength"),
//   	},
//   }
//
type CfnChannel_Mpeg2FilterSettingsProperty struct {
	// Settings for applying the temporal filter to the video.
	TemporalFilterSettings interface{} `field:"optional" json:"temporalFilterSettings" yaml:"temporalFilterSettings"`
}

// The settings for the MPEG-2 codec in the output.
//
// The parent of this entity is VideoCodecSetting.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mpeg2SettingsProperty := &mpeg2SettingsProperty{
//   	adaptiveQuantization: jsii.String("adaptiveQuantization"),
//   	afdSignaling: jsii.String("afdSignaling"),
//   	colorMetadata: jsii.String("colorMetadata"),
//   	colorSpace: jsii.String("colorSpace"),
//   	displayAspectRatio: jsii.String("displayAspectRatio"),
//   	filterSettings: &mpeg2FilterSettingsProperty{
//   		temporalFilterSettings: &temporalFilterSettingsProperty{
//   			postFilterSharpening: jsii.String("postFilterSharpening"),
//   			strength: jsii.String("strength"),
//   		},
//   	},
//   	fixedAfd: jsii.String("fixedAfd"),
//   	framerateDenominator: jsii.Number(123),
//   	framerateNumerator: jsii.Number(123),
//   	gopClosedCadence: jsii.Number(123),
//   	gopNumBFrames: jsii.Number(123),
//   	gopSize: jsii.Number(123),
//   	gopSizeUnits: jsii.String("gopSizeUnits"),
//   	scanType: jsii.String("scanType"),
//   	subgopLength: jsii.String("subgopLength"),
//   	timecodeInsertion: jsii.String("timecodeInsertion"),
//   }
//
type CfnChannel_Mpeg2SettingsProperty struct {
	// Choose Off to disable adaptive quantization.
	//
	// Or choose another value to enable the quantizer and set its strength. The strengths are: Auto, Off, Low, Medium, High. When you enable this field, MediaLive allows intra-frame quantizers to vary, which might improve visual quality.
	AdaptiveQuantization *string `field:"optional" json:"adaptiveQuantization" yaml:"adaptiveQuantization"`
	// Indicates the AFD values that MediaLive will write into the video encode.
	//
	// If you do not know what AFD signaling is, or if your downstream system has not given you guidance, choose AUTO.
	// AUTO: MediaLive will try to preserve the input AFD value (in cases where multiple AFD values are valid).
	// FIXED: MediaLive will use the value you specify in fixedAFD.
	AfdSignaling *string `field:"optional" json:"afdSignaling" yaml:"afdSignaling"`
	// Specifies whether to include the color space metadata.
	//
	// The metadata describes the color space that applies to the video (the colorSpace field). We recommend that you insert the metadata.
	ColorMetadata *string `field:"optional" json:"colorMetadata" yaml:"colorMetadata"`
	// Choose the type of color space conversion to apply to the output.
	//
	// For detailed information on setting up both the input and the output to obtain the desired color space in the output, see the section on \"MediaLive Features - Video - color space\" in the MediaLive User Guide.
	// PASSTHROUGH: Keep the color space of the input content - do not convert it.
	// AUTO:Convert all content that is SD to rec 601, and convert all content that is HD to rec 709.
	ColorSpace *string `field:"optional" json:"colorSpace" yaml:"colorSpace"`
	// Sets the pixel aspect ratio for the encode.
	DisplayAspectRatio *string `field:"optional" json:"displayAspectRatio" yaml:"displayAspectRatio"`
	// Optionally specify a noise reduction filter, which can improve quality of compressed content.
	//
	// If you do not choose a filter, no filter will be applied.
	// TEMPORAL: This filter is useful for both source content that is noisy (when it has excessive digital artifacts) and source content that is clean.
	// When the content is noisy, the filter cleans up the source content before the encoding phase, with these two effects: First, it improves the output video quality because the content has been cleaned up. Secondly, it decreases the bandwidth because MediaLive does not waste bits on encoding noise.
	// When the content is reasonably clean, the filter tends to decrease the bitrate.
	FilterSettings interface{} `field:"optional" json:"filterSettings" yaml:"filterSettings"`
	// Complete this field only when afdSignaling is set to FIXED.
	//
	// Enter the AFD value (4 bits) to write on all frames of the video encode.
	FixedAfd *string `field:"optional" json:"fixedAfd" yaml:"fixedAfd"`
	// description": "The framerate denominator.
	//
	// For example, 1001. The framerate is the numerator divided by the denominator. For example, 24000 / 1001 = 23.976 FPS.
	FramerateDenominator *float64 `field:"optional" json:"framerateDenominator" yaml:"framerateDenominator"`
	// The framerate numerator.
	//
	// For example, 24000. The framerate is the numerator divided by the denominator. For example, 24000 / 1001 = 23.976 FPS.
	FramerateNumerator *float64 `field:"optional" json:"framerateNumerator" yaml:"framerateNumerator"`
	// MPEG2: default is open GOP.
	GopClosedCadence *float64 `field:"optional" json:"gopClosedCadence" yaml:"gopClosedCadence"`
	// Relates to the GOP structure.
	//
	// The number of B-frames between reference frames. If you do not know what a B-frame is, use the default.
	GopNumBFrames *float64 `field:"optional" json:"gopNumBFrames" yaml:"gopNumBFrames"`
	// Relates to the GOP structure.
	//
	// The GOP size (keyframe interval) in the units specified in gopSizeUnits. If you do not know what GOP is, use the default.
	// If gopSizeUnits is frames, then the gopSize must be an integer and must be greater than or equal to 1.
	// If gopSizeUnits is seconds, the gopSize must be greater than 0, but does not need to be an integer.
	GopSize *float64 `field:"optional" json:"gopSize" yaml:"gopSize"`
	// Relates to the GOP structure.
	//
	// Specifies whether the gopSize is specified in frames or seconds. If you do not plan to change the default gopSize, leave the default. If you specify SECONDS, MediaLive will internally convert the gop size to a frame count.
	GopSizeUnits *string `field:"optional" json:"gopSizeUnits" yaml:"gopSizeUnits"`
	// Set the scan type of the output to PROGRESSIVE or INTERLACED (top field first).
	ScanType *string `field:"optional" json:"scanType" yaml:"scanType"`
	// Relates to the GOP structure.
	//
	// If you do not know what GOP is, use the default.
	// FIXED: Set the number of B-frames in each sub-GOP to the value in gopNumBFrames.
	// DYNAMIC: Let MediaLive optimize the number of B-frames in each sub-GOP, to improve visual quality.
	SubgopLength *string `field:"optional" json:"subgopLength" yaml:"subgopLength"`
	// Determines how MediaLive inserts timecodes in the output video.
	//
	// For detailed information about setting up the input and the output for a timecode, see the section on \"MediaLive Features - Timecode configuration\" in the MediaLive User Guide.
	// DISABLED: do not include timecodes.
	// GOP_TIMECODE: Include timecode metadata in the GOP header.
	TimecodeInsertion *string `field:"optional" json:"timecodeInsertion" yaml:"timecodeInsertion"`
}

// The settings for a Microsoft Smooth output group.
//
// The parent of this entity is OutputGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   msSmoothGroupSettingsProperty := &msSmoothGroupSettingsProperty{
//   	acquisitionPointId: jsii.String("acquisitionPointId"),
//   	audioOnlyTimecodeControl: jsii.String("audioOnlyTimecodeControl"),
//   	certificateMode: jsii.String("certificateMode"),
//   	connectionRetryInterval: jsii.Number(123),
//   	destination: &outputLocationRefProperty{
//   		destinationRefId: jsii.String("destinationRefId"),
//   	},
//   	eventId: jsii.String("eventId"),
//   	eventIdMode: jsii.String("eventIdMode"),
//   	eventStopBehavior: jsii.String("eventStopBehavior"),
//   	filecacheDuration: jsii.Number(123),
//   	fragmentLength: jsii.Number(123),
//   	inputLossAction: jsii.String("inputLossAction"),
//   	numRetries: jsii.Number(123),
//   	restartDelay: jsii.Number(123),
//   	segmentationMode: jsii.String("segmentationMode"),
//   	sendDelayMs: jsii.Number(123),
//   	sparseTrackType: jsii.String("sparseTrackType"),
//   	streamManifestBehavior: jsii.String("streamManifestBehavior"),
//   	timestampOffset: jsii.String("timestampOffset"),
//   	timestampOffsetMode: jsii.String("timestampOffsetMode"),
//   }
//
type CfnChannel_MsSmoothGroupSettingsProperty struct {
	// The value of the Acquisition Point Identity element that is used in each message placed in the sparse track.
	//
	// Enabled only if sparseTrackType is not "none."
	AcquisitionPointId *string `field:"optional" json:"acquisitionPointId" yaml:"acquisitionPointId"`
	// If set to passthrough for an audio-only Microsoft Smooth output, the fragment absolute time is set to the current timecode.
	//
	// This option does not write timecodes to the audio elementary stream.
	AudioOnlyTimecodeControl *string `field:"optional" json:"audioOnlyTimecodeControl" yaml:"audioOnlyTimecodeControl"`
	// If set to verifyAuthenticity, verifies the HTTPS certificate chain to a trusted certificate authority (CA).
	//
	// This causes HTTPS outputs to self-signed certificates to fail.
	CertificateMode *string `field:"optional" json:"certificateMode" yaml:"certificateMode"`
	// The number of seconds to wait before retrying the connection to the IIS server if the connection is lost.
	//
	// Content is cached during this time, and the cache is delivered to the IIS server after the connection is re-established.
	ConnectionRetryInterval *float64 `field:"optional" json:"connectionRetryInterval" yaml:"connectionRetryInterval"`
	// The Smooth Streaming publish point on an IIS server.
	//
	// MediaLive acts as a "Push" encoder to IIS.
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// The Microsoft Smooth channel ID that is sent to the IIS server.
	//
	// Specify the ID only if eventIdMode is set to useConfigured.
	EventId *string `field:"optional" json:"eventId" yaml:"eventId"`
	// Specifies whether to send a channel ID to the IIS server.
	//
	// If no channel ID is sent and the same channel is used without changing the publishing point, clients might see cached video from the previous run. Options: - "useConfigured" - use the value provided in eventId - "useTimestamp" - generate and send a channel ID based on the current timestamp - "noEventId" - do not send a channel ID to the IIS server.
	EventIdMode *string `field:"optional" json:"eventIdMode" yaml:"eventIdMode"`
	// When set to sendEos, sends an EOS signal to an IIS server when stopping the channel.
	EventStopBehavior *string `field:"optional" json:"eventStopBehavior" yaml:"eventStopBehavior"`
	// The size, in seconds, of the file cache for streaming outputs.
	FilecacheDuration *float64 `field:"optional" json:"filecacheDuration" yaml:"filecacheDuration"`
	// The length, in seconds, of mp4 fragments to generate.
	//
	// The fragment length must be compatible with GOP size and frame rate.
	FragmentLength *float64 `field:"optional" json:"fragmentLength" yaml:"fragmentLength"`
	// A parameter that controls output group behavior on an input loss.
	InputLossAction *string `field:"optional" json:"inputLossAction" yaml:"inputLossAction"`
	// The number of retry attempts.
	NumRetries *float64 `field:"optional" json:"numRetries" yaml:"numRetries"`
	// The number of seconds before initiating a restart due to output failure, due to exhausting the numRetries on one segment, or exceeding filecacheDuration.
	RestartDelay *float64 `field:"optional" json:"restartDelay" yaml:"restartDelay"`
	// useInputSegmentation has been deprecated.
	//
	// The configured segment size is always used.
	SegmentationMode *string `field:"optional" json:"segmentationMode" yaml:"segmentationMode"`
	// The number of milliseconds to delay the output from the second pipeline.
	SendDelayMs *float64 `field:"optional" json:"sendDelayMs" yaml:"sendDelayMs"`
	// If set to scte35, uses incoming SCTE-35 messages to generate a sparse track in this group of Microsoft Smooth outputs.
	SparseTrackType *string `field:"optional" json:"sparseTrackType" yaml:"sparseTrackType"`
	// When set to send, sends a stream manifest so that the publishing point doesn't start until all streams start.
	StreamManifestBehavior *string `field:"optional" json:"streamManifestBehavior" yaml:"streamManifestBehavior"`
	// The timestamp offset for the channel.
	//
	// Used only if timestampOffsetMode is set to useConfiguredOffset.
	TimestampOffset *string `field:"optional" json:"timestampOffset" yaml:"timestampOffset"`
	// The type of timestamp date offset to use.
	//
	// - useEventStartDate: Use the date the channel was started as the offset - useConfiguredOffset: Use an explicitly configured date as the offset.
	TimestampOffsetMode *string `field:"optional" json:"timestampOffsetMode" yaml:"timestampOffsetMode"`
}

// Configuration of a Microsoft Smooth output.
//
// The parent of this entity is OutputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   msSmoothOutputSettingsProperty := &msSmoothOutputSettingsProperty{
//   	h265PackagingType: jsii.String("h265PackagingType"),
//   	nameModifier: jsii.String("nameModifier"),
//   }
//
type CfnChannel_MsSmoothOutputSettingsProperty struct {
	// Only applicable when this output is referencing an H.265 video description. Specifies whether MP4 segments should be packaged as HEV1 or HVC1.
	H265PackagingType *string `field:"optional" json:"h265PackagingType" yaml:"h265PackagingType"`
	// A string that is concatenated to the end of the destination file name.
	//
	// This is required for multiple outputs of the same type.
	NameModifier *string `field:"optional" json:"nameModifier" yaml:"nameModifier"`
}

// The settings for a Multiplex output group.
//
// The parent of this entity is OutputGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiplexGroupSettingsProperty := &multiplexGroupSettingsProperty{
//   }
//
type CfnChannel_MultiplexGroupSettingsProperty struct {
}

// Configuration of a Multiplex output.
//
// The parent of this entity is OutputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiplexOutputSettingsProperty := &multiplexOutputSettingsProperty{
//   	destination: &outputLocationRefProperty{
//   		destinationRefId: jsii.String("destinationRefId"),
//   	},
//   }
//
type CfnChannel_MultiplexOutputSettingsProperty struct {
	// Destination is a Multiplex.
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
}

// Destination settings for a Multiplex output.
//
// The parent of this entity is OutputDestination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiplexProgramChannelDestinationSettingsProperty := &multiplexProgramChannelDestinationSettingsProperty{
//   	multiplexId: jsii.String("multiplexId"),
//   	programName: jsii.String("programName"),
//   }
//
type CfnChannel_MultiplexProgramChannelDestinationSettingsProperty struct {
	// The ID of the Multiplex that the encoder is providing output to.
	//
	// You do not need to specify the individual inputs to the Multiplex; MediaLive will handle the connection of the two MediaLive pipelines to the two Multiplex instances.
	// The Multiplex must be in the same region as the Channel.
	MultiplexId *string `field:"optional" json:"multiplexId" yaml:"multiplexId"`
	// The program name of the Multiplex program that the encoder is providing output to.
	ProgramName *string `field:"optional" json:"programName" yaml:"programName"`
}

// Information about how to connect to the upstream system.
//
// The parent of this entity is InputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkInputSettingsProperty := &networkInputSettingsProperty{
//   	hlsInputSettings: &hlsInputSettingsProperty{
//   		bandwidth: jsii.Number(123),
//   		bufferSegments: jsii.Number(123),
//   		retries: jsii.Number(123),
//   		retryInterval: jsii.Number(123),
//   		scte35Source: jsii.String("scte35Source"),
//   	},
//   	serverValidation: jsii.String("serverValidation"),
//   }
//
type CfnChannel_NetworkInputSettingsProperty struct {
	// Information about how to connect to the upstream system.
	HlsInputSettings interface{} `field:"optional" json:"hlsInputSettings" yaml:"hlsInputSettings"`
	// Checks HTTPS server certificates.
	//
	// When set to checkCryptographyOnly, cryptography in the certificate is checked, but not the server's name. Certain subdomains (notably S3 buckets that use dots in the bucket name) don't strictly match the corresponding certificate's wildcard pattern and would otherwise cause the channel to error. This setting is ignored for protocols that do not use HTTPS.
	ServerValidation *string `field:"optional" json:"serverValidation" yaml:"serverValidation"`
}

// Complete these fields only if you want to insert watermarks of type Nielsen CBET.
//
// The parent of this entity is NielsenWatermarksSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nielsenCBETProperty := &nielsenCBETProperty{
//   	cbetCheckDigitString: jsii.String("cbetCheckDigitString"),
//   	cbetStepaside: jsii.String("cbetStepaside"),
//   	csid: jsii.String("csid"),
//   }
//
type CfnChannel_NielsenCBETProperty struct {
	// Enter the CBET check digits to use in the watermark.
	CbetCheckDigitString *string `field:"optional" json:"cbetCheckDigitString" yaml:"cbetCheckDigitString"`
	// Determines the method of CBET insertion mode when prior encoding is detected on the same layer.
	CbetStepaside *string `field:"optional" json:"cbetStepaside" yaml:"cbetStepaside"`
	// Enter the CBET Source ID (CSID) to use in the watermark.
	Csid *string `field:"optional" json:"csid" yaml:"csid"`
}

// The settings to configure Nielsen watermarks.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nielsenConfigurationProperty := &nielsenConfigurationProperty{
//   	distributorId: jsii.String("distributorId"),
//   	nielsenPcmToId3Tagging: jsii.String("nielsenPcmToId3Tagging"),
//   }
//
type CfnChannel_NielsenConfigurationProperty struct {
	// Enter the Distributor ID assigned to your organization by Nielsen.
	DistributorId *string `field:"optional" json:"distributorId" yaml:"distributorId"`
	// Enables Nielsen PCM to ID3 tagging.
	NielsenPcmToId3Tagging *string `field:"optional" json:"nielsenPcmToId3Tagging" yaml:"nielsenPcmToId3Tagging"`
}

// Complete these fields only if you want to insert watermarks of type Nielsen NAES II (N2) and Nielsen NAES VI (NW).
//
// The parent of this entity is NielsenWatermarksSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nielsenNaesIiNwProperty := &nielsenNaesIiNwProperty{
//   	checkDigitString: jsii.String("checkDigitString"),
//   	sid: jsii.Number(123),
//   }
//
type CfnChannel_NielsenNaesIiNwProperty struct {
	// Enter the check digit string for the watermark.
	CheckDigitString *string `field:"optional" json:"checkDigitString" yaml:"checkDigitString"`
	// Enter the Nielsen Source ID (SID) to include in the watermark.
	Sid *float64 `field:"optional" json:"sid" yaml:"sid"`
}

// Settings to configure Nielsen Watermarks in the audio encode.
//
// The parent of this entity is AudioWatermarkSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nielsenWatermarksSettingsProperty := &nielsenWatermarksSettingsProperty{
//   	nielsenCbetSettings: &nielsenCBETProperty{
//   		cbetCheckDigitString: jsii.String("cbetCheckDigitString"),
//   		cbetStepaside: jsii.String("cbetStepaside"),
//   		csid: jsii.String("csid"),
//   	},
//   	nielsenDistributionType: jsii.String("nielsenDistributionType"),
//   	nielsenNaesIiNwSettings: &nielsenNaesIiNwProperty{
//   		checkDigitString: jsii.String("checkDigitString"),
//   		sid: jsii.Number(123),
//   	},
//   }
//
type CfnChannel_NielsenWatermarksSettingsProperty struct {
	// Complete these fields only if you want to insert watermarks of type Nielsen CBET.
	NielsenCbetSettings interface{} `field:"optional" json:"nielsenCbetSettings" yaml:"nielsenCbetSettings"`
	// Choose the distribution types that you want to assign to the watermarks: - PROGRAM_CONTENT - FINAL_DISTRIBUTOR.
	NielsenDistributionType *string `field:"optional" json:"nielsenDistributionType" yaml:"nielsenDistributionType"`
	// Complete these fields only if you want to insert watermarks of type Nielsen NAES II (N2) and Nielsen NAES VI (NW).
	NielsenNaesIiNwSettings interface{} `field:"optional" json:"nielsenNaesIiNwSettings" yaml:"nielsenNaesIiNwSettings"`
}

// Configuration information for an output.
//
// This entity is at the top level in the channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputDestinationProperty := &outputDestinationProperty{
//   	id: jsii.String("id"),
//   	mediaPackageSettings: []interface{}{
//   		&mediaPackageOutputDestinationSettingsProperty{
//   			channelId: jsii.String("channelId"),
//   		},
//   	},
//   	multiplexSettings: &multiplexProgramChannelDestinationSettingsProperty{
//   		multiplexId: jsii.String("multiplexId"),
//   		programName: jsii.String("programName"),
//   	},
//   	settings: []interface{}{
//   		&outputDestinationSettingsProperty{
//   			passwordParam: jsii.String("passwordParam"),
//   			streamName: jsii.String("streamName"),
//   			url: jsii.String("url"),
//   			username: jsii.String("username"),
//   		},
//   	},
//   }
//
type CfnChannel_OutputDestinationProperty struct {
	// The ID for this destination.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The destination settings for a MediaPackage output.
	MediaPackageSettings interface{} `field:"optional" json:"mediaPackageSettings" yaml:"mediaPackageSettings"`
	// Destination settings for a Multiplex output;
	//
	// one destination for both encoders.
	MultiplexSettings interface{} `field:"optional" json:"multiplexSettings" yaml:"multiplexSettings"`
	// The destination settings for an output.
	Settings interface{} `field:"optional" json:"settings" yaml:"settings"`
}

// The configuration information for this output.
//
// The parent of this entity is OutputDestination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputDestinationSettingsProperty := &outputDestinationSettingsProperty{
//   	passwordParam: jsii.String("passwordParam"),
//   	streamName: jsii.String("streamName"),
//   	url: jsii.String("url"),
//   	username: jsii.String("username"),
//   }
//
type CfnChannel_OutputDestinationSettingsProperty struct {
	// The password parameter that holds the password for accessing the downstream system.
	//
	// This password parameter applies only if the downstream system requires credentials.
	PasswordParam *string `field:"optional" json:"passwordParam" yaml:"passwordParam"`
	// The stream name for the content.
	//
	// This applies only to RTMP outputs.
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
	// The URL for the destination.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// The user name to connect to the downstream system.
	//
	// This applies only if the downstream system requires credentials.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

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

// The configuration of the output group.
//
// The parent of this entity is OutputGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputGroupSettingsProperty := &outputGroupSettingsProperty{
//   	archiveGroupSettings: &archiveGroupSettingsProperty{
//   		archiveCdnSettings: &archiveCdnSettingsProperty{
//   			archiveS3Settings: &archiveS3SettingsProperty{
//   				cannedAcl: jsii.String("cannedAcl"),
//   			},
//   		},
//   		destination: &outputLocationRefProperty{
//   			destinationRefId: jsii.String("destinationRefId"),
//   		},
//   		rolloverInterval: jsii.Number(123),
//   	},
//   	frameCaptureGroupSettings: &frameCaptureGroupSettingsProperty{
//   		destination: &outputLocationRefProperty{
//   			destinationRefId: jsii.String("destinationRefId"),
//   		},
//   		frameCaptureCdnSettings: &frameCaptureCdnSettingsProperty{
//   			frameCaptureS3Settings: &frameCaptureS3SettingsProperty{
//   				cannedAcl: jsii.String("cannedAcl"),
//   			},
//   		},
//   	},
//   	hlsGroupSettings: &hlsGroupSettingsProperty{
//   		adMarkers: []*string{
//   			jsii.String("adMarkers"),
//   		},
//   		baseUrlContent: jsii.String("baseUrlContent"),
//   		baseUrlContent1: jsii.String("baseUrlContent1"),
//   		baseUrlManifest: jsii.String("baseUrlManifest"),
//   		baseUrlManifest1: jsii.String("baseUrlManifest1"),
//   		captionLanguageMappings: []interface{}{
//   			&captionLanguageMappingProperty{
//   				captionChannel: jsii.Number(123),
//   				languageCode: jsii.String("languageCode"),
//   				languageDescription: jsii.String("languageDescription"),
//   			},
//   		},
//   		captionLanguageSetting: jsii.String("captionLanguageSetting"),
//   		clientCache: jsii.String("clientCache"),
//   		codecSpecification: jsii.String("codecSpecification"),
//   		constantIv: jsii.String("constantIv"),
//   		destination: &outputLocationRefProperty{
//   			destinationRefId: jsii.String("destinationRefId"),
//   		},
//   		directoryStructure: jsii.String("directoryStructure"),
//   		discontinuityTags: jsii.String("discontinuityTags"),
//   		encryptionType: jsii.String("encryptionType"),
//   		hlsCdnSettings: &hlsCdnSettingsProperty{
//   			hlsAkamaiSettings: &hlsAkamaiSettingsProperty{
//   				connectionRetryInterval: jsii.Number(123),
//   				filecacheDuration: jsii.Number(123),
//   				httpTransferMode: jsii.String("httpTransferMode"),
//   				numRetries: jsii.Number(123),
//   				restartDelay: jsii.Number(123),
//   				salt: jsii.String("salt"),
//   				token: jsii.String("token"),
//   			},
//   			hlsBasicPutSettings: &hlsBasicPutSettingsProperty{
//   				connectionRetryInterval: jsii.Number(123),
//   				filecacheDuration: jsii.Number(123),
//   				numRetries: jsii.Number(123),
//   				restartDelay: jsii.Number(123),
//   			},
//   			hlsMediaStoreSettings: &hlsMediaStoreSettingsProperty{
//   				connectionRetryInterval: jsii.Number(123),
//   				filecacheDuration: jsii.Number(123),
//   				mediaStoreStorageClass: jsii.String("mediaStoreStorageClass"),
//   				numRetries: jsii.Number(123),
//   				restartDelay: jsii.Number(123),
//   			},
//   			hlsS3Settings: &hlsS3SettingsProperty{
//   				cannedAcl: jsii.String("cannedAcl"),
//   			},
//   			hlsWebdavSettings: &hlsWebdavSettingsProperty{
//   				connectionRetryInterval: jsii.Number(123),
//   				filecacheDuration: jsii.Number(123),
//   				httpTransferMode: jsii.String("httpTransferMode"),
//   				numRetries: jsii.Number(123),
//   				restartDelay: jsii.Number(123),
//   			},
//   		},
//   		hlsId3SegmentTagging: jsii.String("hlsId3SegmentTagging"),
//   		iFrameOnlyPlaylists: jsii.String("iFrameOnlyPlaylists"),
//   		incompleteSegmentBehavior: jsii.String("incompleteSegmentBehavior"),
//   		indexNSegments: jsii.Number(123),
//   		inputLossAction: jsii.String("inputLossAction"),
//   		ivInManifest: jsii.String("ivInManifest"),
//   		ivSource: jsii.String("ivSource"),
//   		keepSegments: jsii.Number(123),
//   		keyFormat: jsii.String("keyFormat"),
//   		keyFormatVersions: jsii.String("keyFormatVersions"),
//   		keyProviderSettings: &keyProviderSettingsProperty{
//   			staticKeySettings: &staticKeySettingsProperty{
//   				keyProviderServer: &inputLocationProperty{
//   					passwordParam: jsii.String("passwordParam"),
//   					uri: jsii.String("uri"),
//   					username: jsii.String("username"),
//   				},
//   				staticKeyValue: jsii.String("staticKeyValue"),
//   			},
//   		},
//   		manifestCompression: jsii.String("manifestCompression"),
//   		manifestDurationFormat: jsii.String("manifestDurationFormat"),
//   		minSegmentLength: jsii.Number(123),
//   		mode: jsii.String("mode"),
//   		outputSelection: jsii.String("outputSelection"),
//   		programDateTime: jsii.String("programDateTime"),
//   		programDateTimeClock: jsii.String("programDateTimeClock"),
//   		programDateTimePeriod: jsii.Number(123),
//   		redundantManifest: jsii.String("redundantManifest"),
//   		segmentationMode: jsii.String("segmentationMode"),
//   		segmentLength: jsii.Number(123),
//   		segmentsPerSubdirectory: jsii.Number(123),
//   		streamInfResolution: jsii.String("streamInfResolution"),
//   		timedMetadataId3Frame: jsii.String("timedMetadataId3Frame"),
//   		timedMetadataId3Period: jsii.Number(123),
//   		timestampDeltaMilliseconds: jsii.Number(123),
//   		tsFileMode: jsii.String("tsFileMode"),
//   	},
//   	mediaPackageGroupSettings: &mediaPackageGroupSettingsProperty{
//   		destination: &outputLocationRefProperty{
//   			destinationRefId: jsii.String("destinationRefId"),
//   		},
//   	},
//   	msSmoothGroupSettings: &msSmoothGroupSettingsProperty{
//   		acquisitionPointId: jsii.String("acquisitionPointId"),
//   		audioOnlyTimecodeControl: jsii.String("audioOnlyTimecodeControl"),
//   		certificateMode: jsii.String("certificateMode"),
//   		connectionRetryInterval: jsii.Number(123),
//   		destination: &outputLocationRefProperty{
//   			destinationRefId: jsii.String("destinationRefId"),
//   		},
//   		eventId: jsii.String("eventId"),
//   		eventIdMode: jsii.String("eventIdMode"),
//   		eventStopBehavior: jsii.String("eventStopBehavior"),
//   		filecacheDuration: jsii.Number(123),
//   		fragmentLength: jsii.Number(123),
//   		inputLossAction: jsii.String("inputLossAction"),
//   		numRetries: jsii.Number(123),
//   		restartDelay: jsii.Number(123),
//   		segmentationMode: jsii.String("segmentationMode"),
//   		sendDelayMs: jsii.Number(123),
//   		sparseTrackType: jsii.String("sparseTrackType"),
//   		streamManifestBehavior: jsii.String("streamManifestBehavior"),
//   		timestampOffset: jsii.String("timestampOffset"),
//   		timestampOffsetMode: jsii.String("timestampOffsetMode"),
//   	},
//   	multiplexGroupSettings: &multiplexGroupSettingsProperty{
//   	},
//   	rtmpGroupSettings: &rtmpGroupSettingsProperty{
//   		adMarkers: []*string{
//   			jsii.String("adMarkers"),
//   		},
//   		authenticationScheme: jsii.String("authenticationScheme"),
//   		cacheFullBehavior: jsii.String("cacheFullBehavior"),
//   		cacheLength: jsii.Number(123),
//   		captionData: jsii.String("captionData"),
//   		inputLossAction: jsii.String("inputLossAction"),
//   		restartDelay: jsii.Number(123),
//   	},
//   	udpGroupSettings: &udpGroupSettingsProperty{
//   		inputLossAction: jsii.String("inputLossAction"),
//   		timedMetadataId3Frame: jsii.String("timedMetadataId3Frame"),
//   		timedMetadataId3Period: jsii.Number(123),
//   	},
//   }
//
type CfnChannel_OutputGroupSettingsProperty struct {
	// The configuration of an archive output group.
	//
	// The parent of this entity is OutputGroupSettings.
	ArchiveGroupSettings interface{} `field:"optional" json:"archiveGroupSettings" yaml:"archiveGroupSettings"`
	// The configuration of a frame capture output group.
	FrameCaptureGroupSettings interface{} `field:"optional" json:"frameCaptureGroupSettings" yaml:"frameCaptureGroupSettings"`
	// The configuration of an HLS output group.
	HlsGroupSettings interface{} `field:"optional" json:"hlsGroupSettings" yaml:"hlsGroupSettings"`
	// The configuration of a MediaPackage output group.
	MediaPackageGroupSettings interface{} `field:"optional" json:"mediaPackageGroupSettings" yaml:"mediaPackageGroupSettings"`
	// The configuration of a Microsoft Smooth output group.
	MsSmoothGroupSettings interface{} `field:"optional" json:"msSmoothGroupSettings" yaml:"msSmoothGroupSettings"`
	// The settings for a Multiplex output group.
	MultiplexGroupSettings interface{} `field:"optional" json:"multiplexGroupSettings" yaml:"multiplexGroupSettings"`
	// The configuration of an RTMP output group.
	RtmpGroupSettings interface{} `field:"optional" json:"rtmpGroupSettings" yaml:"rtmpGroupSettings"`
	// The configuration of a UDP output group.
	UdpGroupSettings interface{} `field:"optional" json:"udpGroupSettings" yaml:"udpGroupSettings"`
}

// A reference to an OutputDestination ID that is defined in the channel.
//
// This entity is used by ArchiveGroupSettings, FrameCaptureGroupSettings, HlsGroupSettings, MediaPackageGroupSettings, MSSmoothGroupSettings, RtmpOutputSettings, and UdpOutputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputLocationRefProperty := &outputLocationRefProperty{
//   	destinationRefId: jsii.String("destinationRefId"),
//   }
//
type CfnChannel_OutputLocationRefProperty struct {
	// A reference ID for this destination.
	DestinationRefId *string `field:"optional" json:"destinationRefId" yaml:"destinationRefId"`
}

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

// The output settings.
//
// The parent of this entity is Output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputSettingsProperty := &outputSettingsProperty{
//   	archiveOutputSettings: &archiveOutputSettingsProperty{
//   		containerSettings: &archiveContainerSettingsProperty{
//   			m2TsSettings: &m2tsSettingsProperty{
//   				absentInputAudioBehavior: jsii.String("absentInputAudioBehavior"),
//   				arib: jsii.String("arib"),
//   				aribCaptionsPid: jsii.String("aribCaptionsPid"),
//   				aribCaptionsPidControl: jsii.String("aribCaptionsPidControl"),
//   				audioBufferModel: jsii.String("audioBufferModel"),
//   				audioFramesPerPes: jsii.Number(123),
//   				audioPids: jsii.String("audioPids"),
//   				audioStreamType: jsii.String("audioStreamType"),
//   				bitrate: jsii.Number(123),
//   				bufferModel: jsii.String("bufferModel"),
//   				ccDescriptor: jsii.String("ccDescriptor"),
//   				dvbNitSettings: &dvbNitSettingsProperty{
//   					networkId: jsii.Number(123),
//   					networkName: jsii.String("networkName"),
//   					repInterval: jsii.Number(123),
//   				},
//   				dvbSdtSettings: &dvbSdtSettingsProperty{
//   					outputSdt: jsii.String("outputSdt"),
//   					repInterval: jsii.Number(123),
//   					serviceName: jsii.String("serviceName"),
//   					serviceProviderName: jsii.String("serviceProviderName"),
//   				},
//   				dvbSubPids: jsii.String("dvbSubPids"),
//   				dvbTdtSettings: &dvbTdtSettingsProperty{
//   					repInterval: jsii.Number(123),
//   				},
//   				dvbTeletextPid: jsii.String("dvbTeletextPid"),
//   				ebif: jsii.String("ebif"),
//   				ebpAudioInterval: jsii.String("ebpAudioInterval"),
//   				ebpLookaheadMs: jsii.Number(123),
//   				ebpPlacement: jsii.String("ebpPlacement"),
//   				ecmPid: jsii.String("ecmPid"),
//   				esRateInPes: jsii.String("esRateInPes"),
//   				etvPlatformPid: jsii.String("etvPlatformPid"),
//   				etvSignalPid: jsii.String("etvSignalPid"),
//   				fragmentTime: jsii.Number(123),
//   				klv: jsii.String("klv"),
//   				klvDataPids: jsii.String("klvDataPids"),
//   				nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   				nullPacketBitrate: jsii.Number(123),
//   				patInterval: jsii.Number(123),
//   				pcrControl: jsii.String("pcrControl"),
//   				pcrPeriod: jsii.Number(123),
//   				pcrPid: jsii.String("pcrPid"),
//   				pmtInterval: jsii.Number(123),
//   				pmtPid: jsii.String("pmtPid"),
//   				programNum: jsii.Number(123),
//   				rateMode: jsii.String("rateMode"),
//   				scte27Pids: jsii.String("scte27Pids"),
//   				scte35Control: jsii.String("scte35Control"),
//   				scte35Pid: jsii.String("scte35Pid"),
//   				segmentationMarkers: jsii.String("segmentationMarkers"),
//   				segmentationStyle: jsii.String("segmentationStyle"),
//   				segmentationTime: jsii.Number(123),
//   				timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   				timedMetadataPid: jsii.String("timedMetadataPid"),
//   				transportStreamId: jsii.Number(123),
//   				videoPid: jsii.String("videoPid"),
//   			},
//   			rawSettings: &rawSettingsProperty{
//   			},
//   		},
//   		extension: jsii.String("extension"),
//   		nameModifier: jsii.String("nameModifier"),
//   	},
//   	frameCaptureOutputSettings: &frameCaptureOutputSettingsProperty{
//   		nameModifier: jsii.String("nameModifier"),
//   	},
//   	hlsOutputSettings: &hlsOutputSettingsProperty{
//   		h265PackagingType: jsii.String("h265PackagingType"),
//   		hlsSettings: &hlsSettingsProperty{
//   			audioOnlyHlsSettings: &audioOnlyHlsSettingsProperty{
//   				audioGroupId: jsii.String("audioGroupId"),
//   				audioOnlyImage: &inputLocationProperty{
//   					passwordParam: jsii.String("passwordParam"),
//   					uri: jsii.String("uri"),
//   					username: jsii.String("username"),
//   				},
//   				audioTrackType: jsii.String("audioTrackType"),
//   				segmentType: jsii.String("segmentType"),
//   			},
//   			fmp4HlsSettings: &fmp4HlsSettingsProperty{
//   				audioRenditionSets: jsii.String("audioRenditionSets"),
//   				nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   				timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   			},
//   			frameCaptureHlsSettings: &frameCaptureHlsSettingsProperty{
//   			},
//   			standardHlsSettings: &standardHlsSettingsProperty{
//   				audioRenditionSets: jsii.String("audioRenditionSets"),
//   				m3U8Settings: &m3u8SettingsProperty{
//   					audioFramesPerPes: jsii.Number(123),
//   					audioPids: jsii.String("audioPids"),
//   					ecmPid: jsii.String("ecmPid"),
//   					nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   					patInterval: jsii.Number(123),
//   					pcrControl: jsii.String("pcrControl"),
//   					pcrPeriod: jsii.Number(123),
//   					pcrPid: jsii.String("pcrPid"),
//   					pmtInterval: jsii.Number(123),
//   					pmtPid: jsii.String("pmtPid"),
//   					programNum: jsii.Number(123),
//   					scte35Behavior: jsii.String("scte35Behavior"),
//   					scte35Pid: jsii.String("scte35Pid"),
//   					timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   					timedMetadataPid: jsii.String("timedMetadataPid"),
//   					transportStreamId: jsii.Number(123),
//   					videoPid: jsii.String("videoPid"),
//   				},
//   			},
//   		},
//   		nameModifier: jsii.String("nameModifier"),
//   		segmentModifier: jsii.String("segmentModifier"),
//   	},
//   	mediaPackageOutputSettings: &mediaPackageOutputSettingsProperty{
//   	},
//   	msSmoothOutputSettings: &msSmoothOutputSettingsProperty{
//   		h265PackagingType: jsii.String("h265PackagingType"),
//   		nameModifier: jsii.String("nameModifier"),
//   	},
//   	multiplexOutputSettings: &multiplexOutputSettingsProperty{
//   		destination: &outputLocationRefProperty{
//   			destinationRefId: jsii.String("destinationRefId"),
//   		},
//   	},
//   	rtmpOutputSettings: &rtmpOutputSettingsProperty{
//   		certificateMode: jsii.String("certificateMode"),
//   		connectionRetryInterval: jsii.Number(123),
//   		destination: &outputLocationRefProperty{
//   			destinationRefId: jsii.String("destinationRefId"),
//   		},
//   		numRetries: jsii.Number(123),
//   	},
//   	udpOutputSettings: &udpOutputSettingsProperty{
//   		bufferMsec: jsii.Number(123),
//   		containerSettings: &udpContainerSettingsProperty{
//   			m2TsSettings: &m2tsSettingsProperty{
//   				absentInputAudioBehavior: jsii.String("absentInputAudioBehavior"),
//   				arib: jsii.String("arib"),
//   				aribCaptionsPid: jsii.String("aribCaptionsPid"),
//   				aribCaptionsPidControl: jsii.String("aribCaptionsPidControl"),
//   				audioBufferModel: jsii.String("audioBufferModel"),
//   				audioFramesPerPes: jsii.Number(123),
//   				audioPids: jsii.String("audioPids"),
//   				audioStreamType: jsii.String("audioStreamType"),
//   				bitrate: jsii.Number(123),
//   				bufferModel: jsii.String("bufferModel"),
//   				ccDescriptor: jsii.String("ccDescriptor"),
//   				dvbNitSettings: &dvbNitSettingsProperty{
//   					networkId: jsii.Number(123),
//   					networkName: jsii.String("networkName"),
//   					repInterval: jsii.Number(123),
//   				},
//   				dvbSdtSettings: &dvbSdtSettingsProperty{
//   					outputSdt: jsii.String("outputSdt"),
//   					repInterval: jsii.Number(123),
//   					serviceName: jsii.String("serviceName"),
//   					serviceProviderName: jsii.String("serviceProviderName"),
//   				},
//   				dvbSubPids: jsii.String("dvbSubPids"),
//   				dvbTdtSettings: &dvbTdtSettingsProperty{
//   					repInterval: jsii.Number(123),
//   				},
//   				dvbTeletextPid: jsii.String("dvbTeletextPid"),
//   				ebif: jsii.String("ebif"),
//   				ebpAudioInterval: jsii.String("ebpAudioInterval"),
//   				ebpLookaheadMs: jsii.Number(123),
//   				ebpPlacement: jsii.String("ebpPlacement"),
//   				ecmPid: jsii.String("ecmPid"),
//   				esRateInPes: jsii.String("esRateInPes"),
//   				etvPlatformPid: jsii.String("etvPlatformPid"),
//   				etvSignalPid: jsii.String("etvSignalPid"),
//   				fragmentTime: jsii.Number(123),
//   				klv: jsii.String("klv"),
//   				klvDataPids: jsii.String("klvDataPids"),
//   				nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   				nullPacketBitrate: jsii.Number(123),
//   				patInterval: jsii.Number(123),
//   				pcrControl: jsii.String("pcrControl"),
//   				pcrPeriod: jsii.Number(123),
//   				pcrPid: jsii.String("pcrPid"),
//   				pmtInterval: jsii.Number(123),
//   				pmtPid: jsii.String("pmtPid"),
//   				programNum: jsii.Number(123),
//   				rateMode: jsii.String("rateMode"),
//   				scte27Pids: jsii.String("scte27Pids"),
//   				scte35Control: jsii.String("scte35Control"),
//   				scte35Pid: jsii.String("scte35Pid"),
//   				segmentationMarkers: jsii.String("segmentationMarkers"),
//   				segmentationStyle: jsii.String("segmentationStyle"),
//   				segmentationTime: jsii.Number(123),
//   				timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   				timedMetadataPid: jsii.String("timedMetadataPid"),
//   				transportStreamId: jsii.Number(123),
//   				videoPid: jsii.String("videoPid"),
//   			},
//   		},
//   		destination: &outputLocationRefProperty{
//   			destinationRefId: jsii.String("destinationRefId"),
//   		},
//   		fecOutputSettings: &fecOutputSettingsProperty{
//   			columnDepth: jsii.Number(123),
//   			includeFec: jsii.String("includeFec"),
//   			rowLength: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnChannel_OutputSettingsProperty struct {
	// The settings for an archive output.
	ArchiveOutputSettings interface{} `field:"optional" json:"archiveOutputSettings" yaml:"archiveOutputSettings"`
	// The settings for a frame capture output.
	//
	// The parent of this entity is OutputGroupSettings.
	FrameCaptureOutputSettings interface{} `field:"optional" json:"frameCaptureOutputSettings" yaml:"frameCaptureOutputSettings"`
	// The settings for an HLS output.
	//
	// The parent of this entity is OutputGroupSettings.
	HlsOutputSettings interface{} `field:"optional" json:"hlsOutputSettings" yaml:"hlsOutputSettings"`
	// The settings for a MediaPackage output.
	//
	// The parent of this entity is OutputGroupSettings.
	MediaPackageOutputSettings interface{} `field:"optional" json:"mediaPackageOutputSettings" yaml:"mediaPackageOutputSettings"`
	// The settings for a Microsoft Smooth output.
	MsSmoothOutputSettings interface{} `field:"optional" json:"msSmoothOutputSettings" yaml:"msSmoothOutputSettings"`
	// Configuration of a Multiplex output.
	MultiplexOutputSettings interface{} `field:"optional" json:"multiplexOutputSettings" yaml:"multiplexOutputSettings"`
	// The settings for an RTMP output.
	//
	// The parent of this entity is OutputGroupSettings.
	RtmpOutputSettings interface{} `field:"optional" json:"rtmpOutputSettings" yaml:"rtmpOutputSettings"`
	// The settings for a UDP output.
	//
	// The parent of this entity is OutputGroupSettings.
	UdpOutputSettings interface{} `field:"optional" json:"udpOutputSettings" yaml:"udpOutputSettings"`
}

// The settings for passing through audio to the output.
//
// The parent of this entity is AudioCodecSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   passThroughSettingsProperty := &passThroughSettingsProperty{
//   }
//
type CfnChannel_PassThroughSettingsProperty struct {
}

// The container for WAV audio in the output group.
//
// The parent of this entity is ArchiveContainerSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rawSettingsProperty := &rawSettingsProperty{
//   }
//
type CfnChannel_RawSettingsProperty struct {
}

// Rec601 Settings.
//
// The parents of this entity are H264ColorSpaceSettings and H265ColorSpaceSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rec601SettingsProperty := &rec601SettingsProperty{
//   }
//
type CfnChannel_Rec601SettingsProperty struct {
}

// Rec709 Settings.
//
// The parents of this entity are H264ColorSpaceSettings and H265ColorSpaceSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rec709SettingsProperty := &rec709SettingsProperty{
//   }
//
type CfnChannel_Rec709SettingsProperty struct {
}

// The settings for remixing audio in the output.
//
// The parent of this entity is AudioDescription.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   remixSettingsProperty := &remixSettingsProperty{
//   	channelMappings: []interface{}{
//   		&audioChannelMappingProperty{
//   			inputChannelLevels: []interface{}{
//   				&inputChannelLevelProperty{
//   					gain: jsii.Number(123),
//   					inputChannel: jsii.Number(123),
//   				},
//   			},
//   			outputChannel: jsii.Number(123),
//   		},
//   	},
//   	channelsIn: jsii.Number(123),
//   	channelsOut: jsii.Number(123),
//   }
//
type CfnChannel_RemixSettingsProperty struct {
	// A mapping of input channels to output channels, with appropriate gain adjustments.
	ChannelMappings interface{} `field:"optional" json:"channelMappings" yaml:"channelMappings"`
	// The number of input channels to be used.
	ChannelsIn *float64 `field:"optional" json:"channelsIn" yaml:"channelsIn"`
	// The number of output channels to be produced.
	//
	// Valid values: 1, 2, 4, 6, 8.
	ChannelsOut *float64 `field:"optional" json:"channelsOut" yaml:"channelsOut"`
}

// The settings for RTMPCaptionInfo captions encode in the output.
//
// The parent of this entity is CaptionDestinationSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rtmpCaptionInfoDestinationSettingsProperty := &rtmpCaptionInfoDestinationSettingsProperty{
//   }
//
type CfnChannel_RtmpCaptionInfoDestinationSettingsProperty struct {
}

// The configuration of an RTMP output group.
//
// The parent of this entity is OutputGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rtmpGroupSettingsProperty := &rtmpGroupSettingsProperty{
//   	adMarkers: []*string{
//   		jsii.String("adMarkers"),
//   	},
//   	authenticationScheme: jsii.String("authenticationScheme"),
//   	cacheFullBehavior: jsii.String("cacheFullBehavior"),
//   	cacheLength: jsii.Number(123),
//   	captionData: jsii.String("captionData"),
//   	inputLossAction: jsii.String("inputLossAction"),
//   	restartDelay: jsii.Number(123),
//   }
//
type CfnChannel_RtmpGroupSettingsProperty struct {
	// Choose the ad marker type for this output group.
	//
	// MediaLive will create a message based on the content of each SCTE-35 message, format it for that marker type, and insert it in the datastream.
	AdMarkers *[]*string `field:"optional" json:"adMarkers" yaml:"adMarkers"`
	// An authentication scheme to use when connecting with a CDN.
	AuthenticationScheme *string `field:"optional" json:"authenticationScheme" yaml:"authenticationScheme"`
	// Controls behavior when the content cache fills up.
	//
	// If a remote origin server stalls the RTMP connection and doesn't accept content fast enough, the media cache fills up. When the cache reaches the duration specified by cacheLength, the cache stops accepting new content. If set to disconnectImmediately, the RTMP output forces a disconnect. Clear the media cache, and reconnect after restartDelay seconds. If set to waitForServer, the RTMP output waits up to 5 minutes to allow the origin server to begin accepting data again.
	CacheFullBehavior *string `field:"optional" json:"cacheFullBehavior" yaml:"cacheFullBehavior"`
	// The cache length, in seconds, that is used to calculate buffer size.
	CacheLength *float64 `field:"optional" json:"cacheLength" yaml:"cacheLength"`
	// Controls the types of data that pass to onCaptionInfo outputs.
	//
	// If set to all, 608 and 708 carried DTVCC data is passed. If set to field1AndField2608, DTVCC data is stripped out, but 608 data from both fields is passed. If set to field1608, only the data carried in 608 from field 1 video is passed.
	CaptionData *string `field:"optional" json:"captionData" yaml:"captionData"`
	// Controls the behavior of this RTMP group if the input becomes unavailable.
	//
	// emitOutput: Emit a slate until the input returns. pauseOutput: Stop transmitting data until the input returns. This does not close the underlying RTMP connection.
	InputLossAction *string `field:"optional" json:"inputLossAction" yaml:"inputLossAction"`
	// If a streaming output fails, the number of seconds to wait until a restart is initiated.
	//
	// A value of 0 means never restart.
	RestartDelay *float64 `field:"optional" json:"restartDelay" yaml:"restartDelay"`
}

// The settings for one RTMP output.
//
// The parent of this entity is OutputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rtmpOutputSettingsProperty := &rtmpOutputSettingsProperty{
//   	certificateMode: jsii.String("certificateMode"),
//   	connectionRetryInterval: jsii.Number(123),
//   	destination: &outputLocationRefProperty{
//   		destinationRefId: jsii.String("destinationRefId"),
//   	},
//   	numRetries: jsii.Number(123),
//   }
//
type CfnChannel_RtmpOutputSettingsProperty struct {
	// If set to verifyAuthenticity, verifies the TLS certificate chain to a trusted certificate authority (CA).
	//
	// This causes RTMPS outputs with self-signed certificates to fail.
	CertificateMode *string `field:"optional" json:"certificateMode" yaml:"certificateMode"`
	// The number of seconds to wait before retrying a connection to the Flash Media server if the connection is lost.
	ConnectionRetryInterval *float64 `field:"optional" json:"connectionRetryInterval" yaml:"connectionRetryInterval"`
	// The RTMP endpoint excluding the stream name (for example, rtmp://host/appname).
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// The number of retry attempts.
	NumRetries *float64 `field:"optional" json:"numRetries" yaml:"numRetries"`
}

// The configuration of SCTE-20 plus embedded captions in the output.
//
// The parent of this entity is CaptionDestinationSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scte20PlusEmbeddedDestinationSettingsProperty := &scte20PlusEmbeddedDestinationSettingsProperty{
//   }
//
type CfnChannel_Scte20PlusEmbeddedDestinationSettingsProperty struct {
}

// Information about the SCTE-20 captions to extract from the input.
//
// The parent of this entity is CaptionSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scte20SourceSettingsProperty := &scte20SourceSettingsProperty{
//   	convert608To708: jsii.String("convert608To708"),
//   	source608ChannelNumber: jsii.Number(123),
//   }
//
type CfnChannel_Scte20SourceSettingsProperty struct {
	// If upconvert, 608 data is both passed through the "608 compatibility bytes" fields of the 708 wrapper as well as translated into 708.
	//
	// Any 708 data present in the source content is discarded.
	Convert608To708 *string `field:"optional" json:"convert608To708" yaml:"convert608To708"`
	// Specifies the 608/708 channel number within the video track from which to extract captions.
	Source608ChannelNumber *float64 `field:"optional" json:"source608ChannelNumber" yaml:"source608ChannelNumber"`
}

// The configuration of SCTE-27 captions in the output.
//
// The parent of this entity is CaptionDestinationSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scte27DestinationSettingsProperty := &scte27DestinationSettingsProperty{
//   }
//
type CfnChannel_Scte27DestinationSettingsProperty struct {
}

// Information about the SCTE-27 captions to extract from the input.
//
// The parent of this entity is CaptionSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scte27SourceSettingsProperty := &scte27SourceSettingsProperty{
//   	ocrLanguage: jsii.String("ocrLanguage"),
//   	pid: jsii.Number(123),
//   }
//
type CfnChannel_Scte27SourceSettingsProperty struct {
	// If you will configure a WebVTT caption description that references this caption selector, use this field to provide the language to consider when translating the image-based source to text.
	OcrLanguage *string `field:"optional" json:"ocrLanguage" yaml:"ocrLanguage"`
	// The PID field is used in conjunction with the captions selector languageCode field as follows: Specify PID and Language: Extracts captions from that PID;
	//
	// the language is "informational." Specify PID and omit Language: Extracts the specified PID. Omit PID and specify Language: Extracts the specified language, whichever PID that happens to be. Omit PID and omit Language: Valid only if source is DVB-Sub that is being passed through; all languages are passed through.
	Pid *float64 `field:"optional" json:"pid" yaml:"pid"`
}

// The setup of SCTE-35 splice insert handling.
//
// The parent of this entity is AvailSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scte35SpliceInsertProperty := &scte35SpliceInsertProperty{
//   	adAvailOffset: jsii.Number(123),
//   	noRegionalBlackoutFlag: jsii.String("noRegionalBlackoutFlag"),
//   	webDeliveryAllowedFlag: jsii.String("webDeliveryAllowedFlag"),
//   }
//
type CfnChannel_Scte35SpliceInsertProperty struct {
	// When specified, this offset (in milliseconds) is added to the input ad avail PTS time.
	//
	// This applies only to embedded SCTE 104/35 messages. It doesn't apply to OOB messages.
	AdAvailOffset *float64 `field:"optional" json:"adAvailOffset" yaml:"adAvailOffset"`
	// When set to ignore, segment descriptors with noRegionalBlackoutFlag set to 0 no longer trigger blackouts or ad avail slates.
	NoRegionalBlackoutFlag *string `field:"optional" json:"noRegionalBlackoutFlag" yaml:"noRegionalBlackoutFlag"`
	// When set to ignore, segment descriptors with webDeliveryAllowedFlag set to 0 no longer trigger blackouts or ad avail slates.
	WebDeliveryAllowedFlag *string `field:"optional" json:"webDeliveryAllowedFlag" yaml:"webDeliveryAllowedFlag"`
}

// The settings for the SCTE-35 time signal APOS mode.
//
// The parent of this entity is AvailSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scte35TimeSignalAposProperty := &scte35TimeSignalAposProperty{
//   	adAvailOffset: jsii.Number(123),
//   	noRegionalBlackoutFlag: jsii.String("noRegionalBlackoutFlag"),
//   	webDeliveryAllowedFlag: jsii.String("webDeliveryAllowedFlag"),
//   }
//
type CfnChannel_Scte35TimeSignalAposProperty struct {
	// When specified, this offset (in milliseconds) is added to the input ad avail PTS time.
	//
	// This applies only to embedded SCTE 104/35 messages. It doesn't apply to OOB messages.
	AdAvailOffset *float64 `field:"optional" json:"adAvailOffset" yaml:"adAvailOffset"`
	// When set to ignore, segment descriptors with noRegionalBlackoutFlag set to 0 no longer trigger blackouts or ad avail slates.
	NoRegionalBlackoutFlag *string `field:"optional" json:"noRegionalBlackoutFlag" yaml:"noRegionalBlackoutFlag"`
	// When set to ignore, segment descriptors with webDeliveryAllowedFlag set to 0 no longer trigger blackouts or ad avail slates.
	WebDeliveryAllowedFlag *string `field:"optional" json:"webDeliveryAllowedFlag" yaml:"webDeliveryAllowedFlag"`
}

// The setup of SMPTE-TT captions in the output.
//
// The parent of this entity is CaptionDestinationSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   smpteTtDestinationSettingsProperty := &smpteTtDestinationSettingsProperty{
//   }
//
type CfnChannel_SmpteTtDestinationSettingsProperty struct {
}

// The configuration of an HLS output that is a standard output (not an audio-only output).
//
// The parent of this entity is HlsSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   standardHlsSettingsProperty := &standardHlsSettingsProperty{
//   	audioRenditionSets: jsii.String("audioRenditionSets"),
//   	m3U8Settings: &m3u8SettingsProperty{
//   		audioFramesPerPes: jsii.Number(123),
//   		audioPids: jsii.String("audioPids"),
//   		ecmPid: jsii.String("ecmPid"),
//   		nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   		patInterval: jsii.Number(123),
//   		pcrControl: jsii.String("pcrControl"),
//   		pcrPeriod: jsii.Number(123),
//   		pcrPid: jsii.String("pcrPid"),
//   		pmtInterval: jsii.Number(123),
//   		pmtPid: jsii.String("pmtPid"),
//   		programNum: jsii.Number(123),
//   		scte35Behavior: jsii.String("scte35Behavior"),
//   		scte35Pid: jsii.String("scte35Pid"),
//   		timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   		timedMetadataPid: jsii.String("timedMetadataPid"),
//   		transportStreamId: jsii.Number(123),
//   		videoPid: jsii.String("videoPid"),
//   	},
//   }
//
type CfnChannel_StandardHlsSettingsProperty struct {
	// Lists all the audio groups that are used with the video output stream.
	//
	// This inputs all the audio GROUP-IDs that are associated with the video, separated by a comma (,).
	AudioRenditionSets *string `field:"optional" json:"audioRenditionSets" yaml:"audioRenditionSets"`
	// Settings for the M3U8 container.
	M3U8Settings interface{} `field:"optional" json:"m3U8Settings" yaml:"m3U8Settings"`
}

// The static key settings.
//
// The parent of this entity is KeyProviderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   staticKeySettingsProperty := &staticKeySettingsProperty{
//   	keyProviderServer: &inputLocationProperty{
//   		passwordParam: jsii.String("passwordParam"),
//   		uri: jsii.String("uri"),
//   		username: jsii.String("username"),
//   	},
//   	staticKeyValue: jsii.String("staticKeyValue"),
//   }
//
type CfnChannel_StaticKeySettingsProperty struct {
	// The URL of the license server that is used for protecting content.
	KeyProviderServer interface{} `field:"optional" json:"keyProviderServer" yaml:"keyProviderServer"`
	// The static key value as a 32 character hexadecimal string.
	StaticKeyValue *string `field:"optional" json:"staticKeyValue" yaml:"staticKeyValue"`
}

// The settings for a Teletext captions output encode.
//
// The parent of this entity is CaptionDestinationSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   teletextDestinationSettingsProperty := &teletextDestinationSettingsProperty{
//   }
//
type CfnChannel_TeletextDestinationSettingsProperty struct {
}

// Information about the Teletext captions to extract from the input.
//
// The parent of this entity is CaptionSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   teletextSourceSettingsProperty := &teletextSourceSettingsProperty{
//   	outputRectangle: &captionRectangleProperty{
//   		height: jsii.Number(123),
//   		leftOffset: jsii.Number(123),
//   		topOffset: jsii.Number(123),
//   		width: jsii.Number(123),
//   	},
//   	pageNumber: jsii.String("pageNumber"),
//   }
//
type CfnChannel_TeletextSourceSettingsProperty struct {
	// Settings to configure the caption rectangle for an output captions that will be created using this Teletext source captions.
	OutputRectangle interface{} `field:"optional" json:"outputRectangle" yaml:"outputRectangle"`
	// Specifies the Teletext page number within the data stream from which to extract captions.
	//
	// The range is 0x100 (256) to 0x8FF (2303). This is unused for passthrough. It should be specified as a hexadecimal string with no "0x" prefix.
	PageNumber *string `field:"optional" json:"pageNumber" yaml:"pageNumber"`
}

// Settings for the temporal filter to apply to the video.
//
// The parents of this entity are H264FilterSettings, H265FilterSettings, and Mpeg2FilterSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   temporalFilterSettingsProperty := &temporalFilterSettingsProperty{
//   	postFilterSharpening: jsii.String("postFilterSharpening"),
//   	strength: jsii.String("strength"),
//   }
//
type CfnChannel_TemporalFilterSettingsProperty struct {
	// If you enable this filter, the results are the following: - If the source content is noisy (it contains excessive digital artifacts), the filter cleans up the source.
	//
	// - If the source content is already clean, the filter tends to decrease the bitrate, especially when the rate control mode is QVBR.
	PostFilterSharpening *string `field:"optional" json:"postFilterSharpening" yaml:"postFilterSharpening"`
	// Choose a filter strength.
	//
	// We recommend a strength of 1 or 2. A higher strength might take out good information, resulting in an image that is overly soft.
	Strength *string `field:"optional" json:"strength" yaml:"strength"`
}

// The configuration of the timecode in the output.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timecodeConfigProperty := &timecodeConfigProperty{
//   	source: jsii.String("source"),
//   	syncThreshold: jsii.Number(123),
//   }
//
type CfnChannel_TimecodeConfigProperty struct {
	// Identifies the source for the timecode that will be associated with the channel outputs.
	//
	// Embedded (embedded): Initialize the output timecode with timecode from the source. If no embedded timecode is detected in the source, the system falls back to using "Start at 0" (zerobased). System Clock (systemclock): Use the UTC time. Start at 0 (zerobased): The time of the first frame of the channel will be 00:00:00:00.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// The threshold in frames beyond which output timecode is resynchronized to the input timecode.
	//
	// Discrepancies below this threshold are permitted to avoid unnecessary discontinuities in the output timecode. There is no timecode sync when this is not specified.
	SyncThreshold *float64 `field:"optional" json:"syncThreshold" yaml:"syncThreshold"`
}

// The setup of TTML captions in the output.
//
// The parent of this entity is CaptionDestinationSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ttmlDestinationSettingsProperty := &ttmlDestinationSettingsProperty{
//   	styleControl: jsii.String("styleControl"),
//   }
//
type CfnChannel_TtmlDestinationSettingsProperty struct {
	// When set to passthrough, passes through style and position information from a TTML-like input source (TTML, SMPTE-TT, CFF-TT) to the CFF-TT output or TTML output.
	StyleControl *string `field:"optional" json:"styleControl" yaml:"styleControl"`
}

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

// The configuration of a UDP output group.
//
// The parent of this entity is OutputGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   udpGroupSettingsProperty := &udpGroupSettingsProperty{
//   	inputLossAction: jsii.String("inputLossAction"),
//   	timedMetadataId3Frame: jsii.String("timedMetadataId3Frame"),
//   	timedMetadataId3Period: jsii.Number(123),
//   }
//
type CfnChannel_UdpGroupSettingsProperty struct {
	// Specifies the behavior of the last resort when the input video is lost, and no more backup inputs are available.
	//
	// When dropTs is selected, the entire transport stream stops emitting. When dropProgram is selected, the program can be dropped from the transport stream (and replaced with null packets to meet the TS bitrate requirement). Or when emitProgram is selected, the transport stream continues to be produced normally with repeat frames, black frames, or slate frames substituted for the absent input video.
	InputLossAction *string `field:"optional" json:"inputLossAction" yaml:"inputLossAction"`
	// Indicates the ID3 frame that has the timecode.
	TimedMetadataId3Frame *string `field:"optional" json:"timedMetadataId3Frame" yaml:"timedMetadataId3Frame"`
	// The timed metadata interval in seconds.
	TimedMetadataId3Period *float64 `field:"optional" json:"timedMetadataId3Period" yaml:"timedMetadataId3Period"`
}

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

// MediaLive will perform a failover if content is considered black for the specified period.
//
// The parent of this entity is FailoverConditionSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   videoBlackFailoverSettingsProperty := &videoBlackFailoverSettingsProperty{
//   	blackDetectThreshold: jsii.Number(123),
//   	videoBlackThresholdMsec: jsii.Number(123),
//   }
//
type CfnChannel_VideoBlackFailoverSettingsProperty struct {
	// A value used in calculating the threshold below which MediaLive considers a pixel to be 'black'.
	//
	// For the input to be considered black, every pixel in a frame must be below this threshold. The threshold is calculated as a percentage (expressed as a decimal) of white. Therefore .1 means 10% white (or 90% black). Note how the formula works for any color depth. For example, if you set this field to 0.1 in 10-bit color depth: (1023*0.1=102.3), which means a pixel value of 102 or less is 'black'. If you set this field to .1 in an 8-bit color depth: (255*0.1=25.5), which means a pixel value of 25 or less is 'black'. The range is 0.0 to 1.0, with any number of decimal places.
	BlackDetectThreshold *float64 `field:"optional" json:"blackDetectThreshold" yaml:"blackDetectThreshold"`
	// The amount of time (in milliseconds) that the active input must be black before automatic input failover occurs.
	VideoBlackThresholdMsec *float64 `field:"optional" json:"videoBlackThresholdMsec" yaml:"videoBlackThresholdMsec"`
}

// The settings for the video codec in the output.
//
// The parent of this entity is VideoDescription.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   videoCodecSettingsProperty := &videoCodecSettingsProperty{
//   	frameCaptureSettings: &frameCaptureSettingsProperty{
//   		captureInterval: jsii.Number(123),
//   		captureIntervalUnits: jsii.String("captureIntervalUnits"),
//   	},
//   	h264Settings: &h264SettingsProperty{
//   		adaptiveQuantization: jsii.String("adaptiveQuantization"),
//   		afdSignaling: jsii.String("afdSignaling"),
//   		bitrate: jsii.Number(123),
//   		bufFillPct: jsii.Number(123),
//   		bufSize: jsii.Number(123),
//   		colorMetadata: jsii.String("colorMetadata"),
//   		colorSpaceSettings: &h264ColorSpaceSettingsProperty{
//   			colorSpacePassthroughSettings: &colorSpacePassthroughSettingsProperty{
//   			},
//   			rec601Settings: &rec601SettingsProperty{
//   			},
//   			rec709Settings: &rec709SettingsProperty{
//   			},
//   		},
//   		entropyEncoding: jsii.String("entropyEncoding"),
//   		filterSettings: &h264FilterSettingsProperty{
//   			temporalFilterSettings: &temporalFilterSettingsProperty{
//   				postFilterSharpening: jsii.String("postFilterSharpening"),
//   				strength: jsii.String("strength"),
//   			},
//   		},
//   		fixedAfd: jsii.String("fixedAfd"),
//   		flickerAq: jsii.String("flickerAq"),
//   		forceFieldPictures: jsii.String("forceFieldPictures"),
//   		framerateControl: jsii.String("framerateControl"),
//   		framerateDenominator: jsii.Number(123),
//   		framerateNumerator: jsii.Number(123),
//   		gopBReference: jsii.String("gopBReference"),
//   		gopClosedCadence: jsii.Number(123),
//   		gopNumBFrames: jsii.Number(123),
//   		gopSize: jsii.Number(123),
//   		gopSizeUnits: jsii.String("gopSizeUnits"),
//   		level: jsii.String("level"),
//   		lookAheadRateControl: jsii.String("lookAheadRateControl"),
//   		maxBitrate: jsii.Number(123),
//   		minIInterval: jsii.Number(123),
//   		numRefFrames: jsii.Number(123),
//   		parControl: jsii.String("parControl"),
//   		parDenominator: jsii.Number(123),
//   		parNumerator: jsii.Number(123),
//   		profile: jsii.String("profile"),
//   		qualityLevel: jsii.String("qualityLevel"),
//   		qvbrQualityLevel: jsii.Number(123),
//   		rateControlMode: jsii.String("rateControlMode"),
//   		scanType: jsii.String("scanType"),
//   		sceneChangeDetect: jsii.String("sceneChangeDetect"),
//   		slices: jsii.Number(123),
//   		softness: jsii.Number(123),
//   		spatialAq: jsii.String("spatialAq"),
//   		subgopLength: jsii.String("subgopLength"),
//   		syntax: jsii.String("syntax"),
//   		temporalAq: jsii.String("temporalAq"),
//   		timecodeInsertion: jsii.String("timecodeInsertion"),
//   	},
//   	h265Settings: &h265SettingsProperty{
//   		adaptiveQuantization: jsii.String("adaptiveQuantization"),
//   		afdSignaling: jsii.String("afdSignaling"),
//   		alternativeTransferFunction: jsii.String("alternativeTransferFunction"),
//   		bitrate: jsii.Number(123),
//   		bufSize: jsii.Number(123),
//   		colorMetadata: jsii.String("colorMetadata"),
//   		colorSpaceSettings: &h265ColorSpaceSettingsProperty{
//   			colorSpacePassthroughSettings: &colorSpacePassthroughSettingsProperty{
//   			},
//   			hdr10Settings: &hdr10SettingsProperty{
//   				maxCll: jsii.Number(123),
//   				maxFall: jsii.Number(123),
//   			},
//   			rec601Settings: &rec601SettingsProperty{
//   			},
//   			rec709Settings: &rec709SettingsProperty{
//   			},
//   		},
//   		filterSettings: &h265FilterSettingsProperty{
//   			temporalFilterSettings: &temporalFilterSettingsProperty{
//   				postFilterSharpening: jsii.String("postFilterSharpening"),
//   				strength: jsii.String("strength"),
//   			},
//   		},
//   		fixedAfd: jsii.String("fixedAfd"),
//   		flickerAq: jsii.String("flickerAq"),
//   		framerateDenominator: jsii.Number(123),
//   		framerateNumerator: jsii.Number(123),
//   		gopClosedCadence: jsii.Number(123),
//   		gopSize: jsii.Number(123),
//   		gopSizeUnits: jsii.String("gopSizeUnits"),
//   		level: jsii.String("level"),
//   		lookAheadRateControl: jsii.String("lookAheadRateControl"),
//   		maxBitrate: jsii.Number(123),
//   		minIInterval: jsii.Number(123),
//   		parDenominator: jsii.Number(123),
//   		parNumerator: jsii.Number(123),
//   		profile: jsii.String("profile"),
//   		qvbrQualityLevel: jsii.Number(123),
//   		rateControlMode: jsii.String("rateControlMode"),
//   		scanType: jsii.String("scanType"),
//   		sceneChangeDetect: jsii.String("sceneChangeDetect"),
//   		slices: jsii.Number(123),
//   		tier: jsii.String("tier"),
//   		timecodeInsertion: jsii.String("timecodeInsertion"),
//   	},
//   	mpeg2Settings: &mpeg2SettingsProperty{
//   		adaptiveQuantization: jsii.String("adaptiveQuantization"),
//   		afdSignaling: jsii.String("afdSignaling"),
//   		colorMetadata: jsii.String("colorMetadata"),
//   		colorSpace: jsii.String("colorSpace"),
//   		displayAspectRatio: jsii.String("displayAspectRatio"),
//   		filterSettings: &mpeg2FilterSettingsProperty{
//   			temporalFilterSettings: &temporalFilterSettingsProperty{
//   				postFilterSharpening: jsii.String("postFilterSharpening"),
//   				strength: jsii.String("strength"),
//   			},
//   		},
//   		fixedAfd: jsii.String("fixedAfd"),
//   		framerateDenominator: jsii.Number(123),
//   		framerateNumerator: jsii.Number(123),
//   		gopClosedCadence: jsii.Number(123),
//   		gopNumBFrames: jsii.Number(123),
//   		gopSize: jsii.Number(123),
//   		gopSizeUnits: jsii.String("gopSizeUnits"),
//   		scanType: jsii.String("scanType"),
//   		subgopLength: jsii.String("subgopLength"),
//   		timecodeInsertion: jsii.String("timecodeInsertion"),
//   	},
//   }
//
type CfnChannel_VideoCodecSettingsProperty struct {
	// The settings for the video codec in a frame capture output.
	FrameCaptureSettings interface{} `field:"optional" json:"frameCaptureSettings" yaml:"frameCaptureSettings"`
	// The settings for the H.264 codec in the output.
	H264Settings interface{} `field:"optional" json:"h264Settings" yaml:"h264Settings"`
	// Settings for video encoded with the H265 codec.
	H265Settings interface{} `field:"optional" json:"h265Settings" yaml:"h265Settings"`
	// Settings for video encoded with the MPEG-2 codec.
	Mpeg2Settings interface{} `field:"optional" json:"mpeg2Settings" yaml:"mpeg2Settings"`
}

// Encoding information for one output video.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   videoDescriptionProperty := &videoDescriptionProperty{
//   	codecSettings: &videoCodecSettingsProperty{
//   		frameCaptureSettings: &frameCaptureSettingsProperty{
//   			captureInterval: jsii.Number(123),
//   			captureIntervalUnits: jsii.String("captureIntervalUnits"),
//   		},
//   		h264Settings: &h264SettingsProperty{
//   			adaptiveQuantization: jsii.String("adaptiveQuantization"),
//   			afdSignaling: jsii.String("afdSignaling"),
//   			bitrate: jsii.Number(123),
//   			bufFillPct: jsii.Number(123),
//   			bufSize: jsii.Number(123),
//   			colorMetadata: jsii.String("colorMetadata"),
//   			colorSpaceSettings: &h264ColorSpaceSettingsProperty{
//   				colorSpacePassthroughSettings: &colorSpacePassthroughSettingsProperty{
//   				},
//   				rec601Settings: &rec601SettingsProperty{
//   				},
//   				rec709Settings: &rec709SettingsProperty{
//   				},
//   			},
//   			entropyEncoding: jsii.String("entropyEncoding"),
//   			filterSettings: &h264FilterSettingsProperty{
//   				temporalFilterSettings: &temporalFilterSettingsProperty{
//   					postFilterSharpening: jsii.String("postFilterSharpening"),
//   					strength: jsii.String("strength"),
//   				},
//   			},
//   			fixedAfd: jsii.String("fixedAfd"),
//   			flickerAq: jsii.String("flickerAq"),
//   			forceFieldPictures: jsii.String("forceFieldPictures"),
//   			framerateControl: jsii.String("framerateControl"),
//   			framerateDenominator: jsii.Number(123),
//   			framerateNumerator: jsii.Number(123),
//   			gopBReference: jsii.String("gopBReference"),
//   			gopClosedCadence: jsii.Number(123),
//   			gopNumBFrames: jsii.Number(123),
//   			gopSize: jsii.Number(123),
//   			gopSizeUnits: jsii.String("gopSizeUnits"),
//   			level: jsii.String("level"),
//   			lookAheadRateControl: jsii.String("lookAheadRateControl"),
//   			maxBitrate: jsii.Number(123),
//   			minIInterval: jsii.Number(123),
//   			numRefFrames: jsii.Number(123),
//   			parControl: jsii.String("parControl"),
//   			parDenominator: jsii.Number(123),
//   			parNumerator: jsii.Number(123),
//   			profile: jsii.String("profile"),
//   			qualityLevel: jsii.String("qualityLevel"),
//   			qvbrQualityLevel: jsii.Number(123),
//   			rateControlMode: jsii.String("rateControlMode"),
//   			scanType: jsii.String("scanType"),
//   			sceneChangeDetect: jsii.String("sceneChangeDetect"),
//   			slices: jsii.Number(123),
//   			softness: jsii.Number(123),
//   			spatialAq: jsii.String("spatialAq"),
//   			subgopLength: jsii.String("subgopLength"),
//   			syntax: jsii.String("syntax"),
//   			temporalAq: jsii.String("temporalAq"),
//   			timecodeInsertion: jsii.String("timecodeInsertion"),
//   		},
//   		h265Settings: &h265SettingsProperty{
//   			adaptiveQuantization: jsii.String("adaptiveQuantization"),
//   			afdSignaling: jsii.String("afdSignaling"),
//   			alternativeTransferFunction: jsii.String("alternativeTransferFunction"),
//   			bitrate: jsii.Number(123),
//   			bufSize: jsii.Number(123),
//   			colorMetadata: jsii.String("colorMetadata"),
//   			colorSpaceSettings: &h265ColorSpaceSettingsProperty{
//   				colorSpacePassthroughSettings: &colorSpacePassthroughSettingsProperty{
//   				},
//   				hdr10Settings: &hdr10SettingsProperty{
//   					maxCll: jsii.Number(123),
//   					maxFall: jsii.Number(123),
//   				},
//   				rec601Settings: &rec601SettingsProperty{
//   				},
//   				rec709Settings: &rec709SettingsProperty{
//   				},
//   			},
//   			filterSettings: &h265FilterSettingsProperty{
//   				temporalFilterSettings: &temporalFilterSettingsProperty{
//   					postFilterSharpening: jsii.String("postFilterSharpening"),
//   					strength: jsii.String("strength"),
//   				},
//   			},
//   			fixedAfd: jsii.String("fixedAfd"),
//   			flickerAq: jsii.String("flickerAq"),
//   			framerateDenominator: jsii.Number(123),
//   			framerateNumerator: jsii.Number(123),
//   			gopClosedCadence: jsii.Number(123),
//   			gopSize: jsii.Number(123),
//   			gopSizeUnits: jsii.String("gopSizeUnits"),
//   			level: jsii.String("level"),
//   			lookAheadRateControl: jsii.String("lookAheadRateControl"),
//   			maxBitrate: jsii.Number(123),
//   			minIInterval: jsii.Number(123),
//   			parDenominator: jsii.Number(123),
//   			parNumerator: jsii.Number(123),
//   			profile: jsii.String("profile"),
//   			qvbrQualityLevel: jsii.Number(123),
//   			rateControlMode: jsii.String("rateControlMode"),
//   			scanType: jsii.String("scanType"),
//   			sceneChangeDetect: jsii.String("sceneChangeDetect"),
//   			slices: jsii.Number(123),
//   			tier: jsii.String("tier"),
//   			timecodeInsertion: jsii.String("timecodeInsertion"),
//   		},
//   		mpeg2Settings: &mpeg2SettingsProperty{
//   			adaptiveQuantization: jsii.String("adaptiveQuantization"),
//   			afdSignaling: jsii.String("afdSignaling"),
//   			colorMetadata: jsii.String("colorMetadata"),
//   			colorSpace: jsii.String("colorSpace"),
//   			displayAspectRatio: jsii.String("displayAspectRatio"),
//   			filterSettings: &mpeg2FilterSettingsProperty{
//   				temporalFilterSettings: &temporalFilterSettingsProperty{
//   					postFilterSharpening: jsii.String("postFilterSharpening"),
//   					strength: jsii.String("strength"),
//   				},
//   			},
//   			fixedAfd: jsii.String("fixedAfd"),
//   			framerateDenominator: jsii.Number(123),
//   			framerateNumerator: jsii.Number(123),
//   			gopClosedCadence: jsii.Number(123),
//   			gopNumBFrames: jsii.Number(123),
//   			gopSize: jsii.Number(123),
//   			gopSizeUnits: jsii.String("gopSizeUnits"),
//   			scanType: jsii.String("scanType"),
//   			subgopLength: jsii.String("subgopLength"),
//   			timecodeInsertion: jsii.String("timecodeInsertion"),
//   		},
//   	},
//   	height: jsii.Number(123),
//   	name: jsii.String("name"),
//   	respondToAfd: jsii.String("respondToAfd"),
//   	scalingBehavior: jsii.String("scalingBehavior"),
//   	sharpness: jsii.Number(123),
//   	width: jsii.Number(123),
//   }
//
type CfnChannel_VideoDescriptionProperty struct {
	// The video codec settings.
	CodecSettings interface{} `field:"optional" json:"codecSettings" yaml:"codecSettings"`
	// The output video height, in pixels.
	//
	// This must be an even number. For most codecs, you can keep this field and width blank in order to use the height and width (resolution) from the source. Note that we don't recommend keeping the field blank. For the Frame Capture codec, height and width are required.
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// The name of this VideoDescription.
	//
	// Outputs use this name to uniquely identify this description. Description names should be unique within this channel.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Indicates how to respond to the AFD values in the input stream.
	//
	// RESPOND causes input video to be clipped, depending on the AFD value, input display aspect ratio, and output display aspect ratio, and (except for the FRAMECAPTURE codec) includes the values in the output. PASSTHROUGH (does not apply to FRAMECAPTURE codec) ignores the AFD values and includes the values in the output, so input video is not clipped. NONE ignores the AFD values and does not include the values through to the output, so input video is not clipped.
	RespondToAfd *string `field:"optional" json:"respondToAfd" yaml:"respondToAfd"`
	// STRETCHTOOUTPUT configures the output position to stretch the video to the specified output resolution (height and width).
	//
	// This option overrides any position value. DEFAULT might insert black boxes (pillar boxes or letter boxes) around the video to provide the specified output resolution.
	ScalingBehavior *string `field:"optional" json:"scalingBehavior" yaml:"scalingBehavior"`
	// Changes the strength of the anti-alias filter used for scaling.
	//
	// 0 is the softest setting, and 100 is the sharpest. We recommend a setting of 50 for most content.
	Sharpness *float64 `field:"optional" json:"sharpness" yaml:"sharpness"`
	// The output video width, in pixels.
	//
	// It must be an even number. For most codecs, you can keep this field and height blank in order to use the height and width (resolution) from the source. Note that we don't recommend keeping the field blank. For the Frame Capture codec, height and width are required.
	Width *float64 `field:"optional" json:"width" yaml:"width"`
}

// Settings to configure color space settings in the incoming video.
//
// The parent of this entity is VideoSelector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   videoSelectorColorSpaceSettingsProperty := &videoSelectorColorSpaceSettingsProperty{
//   	hdr10Settings: &hdr10SettingsProperty{
//   		maxCll: jsii.Number(123),
//   		maxFall: jsii.Number(123),
//   	},
//   }
//
type CfnChannel_VideoSelectorColorSpaceSettingsProperty struct {
	// Settings to configure color space settings in the incoming video.
	Hdr10Settings interface{} `field:"optional" json:"hdr10Settings" yaml:"hdr10Settings"`
}

// Selects a specific PID from within a video source.
//
// The parent of this entity is VideoSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   videoSelectorPidProperty := &videoSelectorPidProperty{
//   	pid: jsii.Number(123),
//   }
//
type CfnChannel_VideoSelectorPidProperty struct {
	// Selects a specific PID from within a video source.
	Pid *float64 `field:"optional" json:"pid" yaml:"pid"`
}

// Used to extract video by the program ID.
//
// The parent of this entity is VideoSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   videoSelectorProgramIdProperty := &videoSelectorProgramIdProperty{
//   	programId: jsii.Number(123),
//   }
//
type CfnChannel_VideoSelectorProgramIdProperty struct {
	// Selects a specific program from within a multi-program transport stream.
	//
	// If the program doesn't exist, MediaLive selects the first program within the transport stream by default.
	ProgramId *float64 `field:"optional" json:"programId" yaml:"programId"`
}

// Information about the video to extract from the input. An input can contain only one video selector.
//
// The parent of this entity is InputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   videoSelectorProperty := &videoSelectorProperty{
//   	colorSpace: jsii.String("colorSpace"),
//   	colorSpaceSettings: &videoSelectorColorSpaceSettingsProperty{
//   		hdr10Settings: &hdr10SettingsProperty{
//   			maxCll: jsii.Number(123),
//   			maxFall: jsii.Number(123),
//   		},
//   	},
//   	colorSpaceUsage: jsii.String("colorSpaceUsage"),
//   	selectorSettings: &videoSelectorSettingsProperty{
//   		videoSelectorPid: &videoSelectorPidProperty{
//   			pid: jsii.Number(123),
//   		},
//   		videoSelectorProgramId: &videoSelectorProgramIdProperty{
//   			programId: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnChannel_VideoSelectorProperty struct {
	// Specifies the color space of an input.
	//
	// This setting works in tandem with colorSpaceConversion to determine if MediaLive will perform any conversion.
	ColorSpace *string `field:"optional" json:"colorSpace" yaml:"colorSpace"`
	// Settings to configure color space settings in the incoming video.
	ColorSpaceSettings interface{} `field:"optional" json:"colorSpaceSettings" yaml:"colorSpaceSettings"`
	// Applies only if colorSpace is a value other than Follow.
	//
	// This field controls how the value in the colorSpace field is used. Fallback means that when the input does include color space data, that data is used, but when the input has no color space data, the value in colorSpace is used. Choose fallback if your input is sometimes missing color space data, but when it does have color space data, that data is correct. Force means to always use the value in colorSpace. Choose force if your input usually has no color space data or might have unreliable color space data.
	ColorSpaceUsage *string `field:"optional" json:"colorSpaceUsage" yaml:"colorSpaceUsage"`
	// Information about the video to select from the content.
	SelectorSettings interface{} `field:"optional" json:"selectorSettings" yaml:"selectorSettings"`
}

// Information about the video to extract from the input.
//
// The parent of this entity is VideoSelector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   videoSelectorSettingsProperty := &videoSelectorSettingsProperty{
//   	videoSelectorPid: &videoSelectorPidProperty{
//   		pid: jsii.Number(123),
//   	},
//   	videoSelectorProgramId: &videoSelectorProgramIdProperty{
//   		programId: jsii.Number(123),
//   	},
//   }
//
type CfnChannel_VideoSelectorSettingsProperty struct {
	// Used to extract video by PID.
	VideoSelectorPid interface{} `field:"optional" json:"videoSelectorPid" yaml:"videoSelectorPid"`
	// Used to extract video by program ID.
	VideoSelectorProgramId interface{} `field:"optional" json:"videoSelectorProgramId" yaml:"videoSelectorProgramId"`
}

// Settings to enable VPC mode in the channel, so that the endpoints for all outputs are in your VPC.
//
// This entity is at the top level in the channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcOutputSettingsProperty := &vpcOutputSettingsProperty{
//   	publicAddressAllocationIds: []*string{
//   		jsii.String("publicAddressAllocationIds"),
//   	},
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   }
//
type CfnChannel_VpcOutputSettingsProperty struct {
	// List of public address allocation IDs to associate with ENIs that will be created in Output VPC.
	//
	// Must specify one for SINGLE_PIPELINE, two for STANDARD channels.
	PublicAddressAllocationIds *[]*string `field:"optional" json:"publicAddressAllocationIds" yaml:"publicAddressAllocationIds"`
	// A list of up to 5 EC2 VPC security group IDs to attach to the Output VPC network interfaces.
	//
	// If none are specified then the VPC default security group will be used.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// A list of VPC subnet IDs from the same VPC.
	//
	// If STANDARD channel, subnet IDs must be mapped to two unique availability zones (AZ).
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

// The setup of WAV audio in the output.
//
// The parent of this entity is AudioCodecSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   wavSettingsProperty := &wavSettingsProperty{
//   	bitDepth: jsii.Number(123),
//   	codingMode: jsii.String("codingMode"),
//   	sampleRate: jsii.Number(123),
//   }
//
type CfnChannel_WavSettingsProperty struct {
	// Bits per sample.
	BitDepth *float64 `field:"optional" json:"bitDepth" yaml:"bitDepth"`
	// The audio coding mode for the WAV audio.
	//
	// The mode determines the number of channels in the audio.
	CodingMode *string `field:"optional" json:"codingMode" yaml:"codingMode"`
	// Sample rate in Hz.
	SampleRate *float64 `field:"optional" json:"sampleRate" yaml:"sampleRate"`
}

// The configuration of Web VTT captions in the output.
//
// The parent of this entity is CaptionDestinationSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webvttDestinationSettingsProperty := &webvttDestinationSettingsProperty{
//   	styleControl: jsii.String("styleControl"),
//   }
//
type CfnChannel_WebvttDestinationSettingsProperty struct {
	// Controls whether the color and position of the source captions is passed through to the WebVTT output captions.
	//
	// PASSTHROUGH - Valid only if the source captions are EMBEDDED or TELETEXT. NO_STYLE_DATA - Don't pass through the style. The output captions will not contain any font styling information.
	StyleControl *string `field:"optional" json:"styleControl" yaml:"styleControl"`
}

// Properties for defining a `CfnChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnChannelProps := &cfnChannelProps{
//   	cdiInputSpecification: &cdiInputSpecificationProperty{
//   		resolution: jsii.String("resolution"),
//   	},
//   	channelClass: jsii.String("channelClass"),
//   	destinations: []interface{}{
//   		&outputDestinationProperty{
//   			id: jsii.String("id"),
//   			mediaPackageSettings: []interface{}{
//   				&mediaPackageOutputDestinationSettingsProperty{
//   					channelId: jsii.String("channelId"),
//   				},
//   			},
//   			multiplexSettings: &multiplexProgramChannelDestinationSettingsProperty{
//   				multiplexId: jsii.String("multiplexId"),
//   				programName: jsii.String("programName"),
//   			},
//   			settings: []interface{}{
//   				&outputDestinationSettingsProperty{
//   					passwordParam: jsii.String("passwordParam"),
//   					streamName: jsii.String("streamName"),
//   					url: jsii.String("url"),
//   					username: jsii.String("username"),
//   				},
//   			},
//   		},
//   	},
//   	encoderSettings: &encoderSettingsProperty{
//   		audioDescriptions: []interface{}{
//   			&audioDescriptionProperty{
//   				audioNormalizationSettings: &audioNormalizationSettingsProperty{
//   					algorithm: jsii.String("algorithm"),
//   					algorithmControl: jsii.String("algorithmControl"),
//   					targetLkfs: jsii.Number(123),
//   				},
//   				audioSelectorName: jsii.String("audioSelectorName"),
//   				audioType: jsii.String("audioType"),
//   				audioTypeControl: jsii.String("audioTypeControl"),
//   				audioWatermarkingSettings: &audioWatermarkSettingsProperty{
//   					nielsenWatermarksSettings: &nielsenWatermarksSettingsProperty{
//   						nielsenCbetSettings: &nielsenCBETProperty{
//   							cbetCheckDigitString: jsii.String("cbetCheckDigitString"),
//   							cbetStepaside: jsii.String("cbetStepaside"),
//   							csid: jsii.String("csid"),
//   						},
//   						nielsenDistributionType: jsii.String("nielsenDistributionType"),
//   						nielsenNaesIiNwSettings: &nielsenNaesIiNwProperty{
//   							checkDigitString: jsii.String("checkDigitString"),
//   							sid: jsii.Number(123),
//   						},
//   					},
//   				},
//   				codecSettings: &audioCodecSettingsProperty{
//   					aacSettings: &aacSettingsProperty{
//   						bitrate: jsii.Number(123),
//   						codingMode: jsii.String("codingMode"),
//   						inputType: jsii.String("inputType"),
//   						profile: jsii.String("profile"),
//   						rateControlMode: jsii.String("rateControlMode"),
//   						rawFormat: jsii.String("rawFormat"),
//   						sampleRate: jsii.Number(123),
//   						spec: jsii.String("spec"),
//   						vbrQuality: jsii.String("vbrQuality"),
//   					},
//   					ac3Settings: &ac3SettingsProperty{
//   						bitrate: jsii.Number(123),
//   						bitstreamMode: jsii.String("bitstreamMode"),
//   						codingMode: jsii.String("codingMode"),
//   						dialnorm: jsii.Number(123),
//   						drcProfile: jsii.String("drcProfile"),
//   						lfeFilter: jsii.String("lfeFilter"),
//   						metadataControl: jsii.String("metadataControl"),
//   					},
//   					eac3Settings: &eac3SettingsProperty{
//   						attenuationControl: jsii.String("attenuationControl"),
//   						bitrate: jsii.Number(123),
//   						bitstreamMode: jsii.String("bitstreamMode"),
//   						codingMode: jsii.String("codingMode"),
//   						dcFilter: jsii.String("dcFilter"),
//   						dialnorm: jsii.Number(123),
//   						drcLine: jsii.String("drcLine"),
//   						drcRf: jsii.String("drcRf"),
//   						lfeControl: jsii.String("lfeControl"),
//   						lfeFilter: jsii.String("lfeFilter"),
//   						loRoCenterMixLevel: jsii.Number(123),
//   						loRoSurroundMixLevel: jsii.Number(123),
//   						ltRtCenterMixLevel: jsii.Number(123),
//   						ltRtSurroundMixLevel: jsii.Number(123),
//   						metadataControl: jsii.String("metadataControl"),
//   						passthroughControl: jsii.String("passthroughControl"),
//   						phaseControl: jsii.String("phaseControl"),
//   						stereoDownmix: jsii.String("stereoDownmix"),
//   						surroundExMode: jsii.String("surroundExMode"),
//   						surroundMode: jsii.String("surroundMode"),
//   					},
//   					mp2Settings: &mp2SettingsProperty{
//   						bitrate: jsii.Number(123),
//   						codingMode: jsii.String("codingMode"),
//   						sampleRate: jsii.Number(123),
//   					},
//   					passThroughSettings: &passThroughSettingsProperty{
//   					},
//   					wavSettings: &wavSettingsProperty{
//   						bitDepth: jsii.Number(123),
//   						codingMode: jsii.String("codingMode"),
//   						sampleRate: jsii.Number(123),
//   					},
//   				},
//   				languageCode: jsii.String("languageCode"),
//   				languageCodeControl: jsii.String("languageCodeControl"),
//   				name: jsii.String("name"),
//   				remixSettings: &remixSettingsProperty{
//   					channelMappings: []interface{}{
//   						&audioChannelMappingProperty{
//   							inputChannelLevels: []interface{}{
//   								&inputChannelLevelProperty{
//   									gain: jsii.Number(123),
//   									inputChannel: jsii.Number(123),
//   								},
//   							},
//   							outputChannel: jsii.Number(123),
//   						},
//   					},
//   					channelsIn: jsii.Number(123),
//   					channelsOut: jsii.Number(123),
//   				},
//   				streamName: jsii.String("streamName"),
//   			},
//   		},
//   		availBlanking: &availBlankingProperty{
//   			availBlankingImage: &inputLocationProperty{
//   				passwordParam: jsii.String("passwordParam"),
//   				uri: jsii.String("uri"),
//   				username: jsii.String("username"),
//   			},
//   			state: jsii.String("state"),
//   		},
//   		availConfiguration: &availConfigurationProperty{
//   			availSettings: &availSettingsProperty{
//   				scte35SpliceInsert: &scte35SpliceInsertProperty{
//   					adAvailOffset: jsii.Number(123),
//   					noRegionalBlackoutFlag: jsii.String("noRegionalBlackoutFlag"),
//   					webDeliveryAllowedFlag: jsii.String("webDeliveryAllowedFlag"),
//   				},
//   				scte35TimeSignalApos: &scte35TimeSignalAposProperty{
//   					adAvailOffset: jsii.Number(123),
//   					noRegionalBlackoutFlag: jsii.String("noRegionalBlackoutFlag"),
//   					webDeliveryAllowedFlag: jsii.String("webDeliveryAllowedFlag"),
//   				},
//   			},
//   		},
//   		blackoutSlate: &blackoutSlateProperty{
//   			blackoutSlateImage: &inputLocationProperty{
//   				passwordParam: jsii.String("passwordParam"),
//   				uri: jsii.String("uri"),
//   				username: jsii.String("username"),
//   			},
//   			networkEndBlackout: jsii.String("networkEndBlackout"),
//   			networkEndBlackoutImage: &inputLocationProperty{
//   				passwordParam: jsii.String("passwordParam"),
//   				uri: jsii.String("uri"),
//   				username: jsii.String("username"),
//   			},
//   			networkId: jsii.String("networkId"),
//   			state: jsii.String("state"),
//   		},
//   		captionDescriptions: []interface{}{
//   			&captionDescriptionProperty{
//   				captionSelectorName: jsii.String("captionSelectorName"),
//   				destinationSettings: &captionDestinationSettingsProperty{
//   					aribDestinationSettings: &aribDestinationSettingsProperty{
//   					},
//   					burnInDestinationSettings: &burnInDestinationSettingsProperty{
//   						alignment: jsii.String("alignment"),
//   						backgroundColor: jsii.String("backgroundColor"),
//   						backgroundOpacity: jsii.Number(123),
//   						font: &inputLocationProperty{
//   							passwordParam: jsii.String("passwordParam"),
//   							uri: jsii.String("uri"),
//   							username: jsii.String("username"),
//   						},
//   						fontColor: jsii.String("fontColor"),
//   						fontOpacity: jsii.Number(123),
//   						fontResolution: jsii.Number(123),
//   						fontSize: jsii.String("fontSize"),
//   						outlineColor: jsii.String("outlineColor"),
//   						outlineSize: jsii.Number(123),
//   						shadowColor: jsii.String("shadowColor"),
//   						shadowOpacity: jsii.Number(123),
//   						shadowXOffset: jsii.Number(123),
//   						shadowYOffset: jsii.Number(123),
//   						teletextGridControl: jsii.String("teletextGridControl"),
//   						xPosition: jsii.Number(123),
//   						yPosition: jsii.Number(123),
//   					},
//   					dvbSubDestinationSettings: &dvbSubDestinationSettingsProperty{
//   						alignment: jsii.String("alignment"),
//   						backgroundColor: jsii.String("backgroundColor"),
//   						backgroundOpacity: jsii.Number(123),
//   						font: &inputLocationProperty{
//   							passwordParam: jsii.String("passwordParam"),
//   							uri: jsii.String("uri"),
//   							username: jsii.String("username"),
//   						},
//   						fontColor: jsii.String("fontColor"),
//   						fontOpacity: jsii.Number(123),
//   						fontResolution: jsii.Number(123),
//   						fontSize: jsii.String("fontSize"),
//   						outlineColor: jsii.String("outlineColor"),
//   						outlineSize: jsii.Number(123),
//   						shadowColor: jsii.String("shadowColor"),
//   						shadowOpacity: jsii.Number(123),
//   						shadowXOffset: jsii.Number(123),
//   						shadowYOffset: jsii.Number(123),
//   						teletextGridControl: jsii.String("teletextGridControl"),
//   						xPosition: jsii.Number(123),
//   						yPosition: jsii.Number(123),
//   					},
//   					ebuTtDDestinationSettings: &ebuTtDDestinationSettingsProperty{
//   						copyrightHolder: jsii.String("copyrightHolder"),
//   						fillLineGap: jsii.String("fillLineGap"),
//   						fontFamily: jsii.String("fontFamily"),
//   						styleControl: jsii.String("styleControl"),
//   					},
//   					embeddedDestinationSettings: &embeddedDestinationSettingsProperty{
//   					},
//   					embeddedPlusScte20DestinationSettings: &embeddedPlusScte20DestinationSettingsProperty{
//   					},
//   					rtmpCaptionInfoDestinationSettings: &rtmpCaptionInfoDestinationSettingsProperty{
//   					},
//   					scte20PlusEmbeddedDestinationSettings: &scte20PlusEmbeddedDestinationSettingsProperty{
//   					},
//   					scte27DestinationSettings: &scte27DestinationSettingsProperty{
//   					},
//   					smpteTtDestinationSettings: &smpteTtDestinationSettingsProperty{
//   					},
//   					teletextDestinationSettings: &teletextDestinationSettingsProperty{
//   					},
//   					ttmlDestinationSettings: &ttmlDestinationSettingsProperty{
//   						styleControl: jsii.String("styleControl"),
//   					},
//   					webvttDestinationSettings: &webvttDestinationSettingsProperty{
//   						styleControl: jsii.String("styleControl"),
//   					},
//   				},
//   				languageCode: jsii.String("languageCode"),
//   				languageDescription: jsii.String("languageDescription"),
//   				name: jsii.String("name"),
//   			},
//   		},
//   		featureActivations: &featureActivationsProperty{
//   			inputPrepareScheduleActions: jsii.String("inputPrepareScheduleActions"),
//   		},
//   		globalConfiguration: &globalConfigurationProperty{
//   			initialAudioGain: jsii.Number(123),
//   			inputEndAction: jsii.String("inputEndAction"),
//   			inputLossBehavior: &inputLossBehaviorProperty{
//   				blackFrameMsec: jsii.Number(123),
//   				inputLossImageColor: jsii.String("inputLossImageColor"),
//   				inputLossImageSlate: &inputLocationProperty{
//   					passwordParam: jsii.String("passwordParam"),
//   					uri: jsii.String("uri"),
//   					username: jsii.String("username"),
//   				},
//   				inputLossImageType: jsii.String("inputLossImageType"),
//   				repeatFrameMsec: jsii.Number(123),
//   			},
//   			outputLockingMode: jsii.String("outputLockingMode"),
//   			outputTimingSource: jsii.String("outputTimingSource"),
//   			supportLowFramerateInputs: jsii.String("supportLowFramerateInputs"),
//   		},
//   		motionGraphicsConfiguration: &motionGraphicsConfigurationProperty{
//   			motionGraphicsInsertion: jsii.String("motionGraphicsInsertion"),
//   			motionGraphicsSettings: &motionGraphicsSettingsProperty{
//   				htmlMotionGraphicsSettings: &htmlMotionGraphicsSettingsProperty{
//   				},
//   			},
//   		},
//   		nielsenConfiguration: &nielsenConfigurationProperty{
//   			distributorId: jsii.String("distributorId"),
//   			nielsenPcmToId3Tagging: jsii.String("nielsenPcmToId3Tagging"),
//   		},
//   		outputGroups: []interface{}{
//   			&outputGroupProperty{
//   				name: jsii.String("name"),
//   				outputGroupSettings: &outputGroupSettingsProperty{
//   					archiveGroupSettings: &archiveGroupSettingsProperty{
//   						archiveCdnSettings: &archiveCdnSettingsProperty{
//   							archiveS3Settings: &archiveS3SettingsProperty{
//   								cannedAcl: jsii.String("cannedAcl"),
//   							},
//   						},
//   						destination: &outputLocationRefProperty{
//   							destinationRefId: jsii.String("destinationRefId"),
//   						},
//   						rolloverInterval: jsii.Number(123),
//   					},
//   					frameCaptureGroupSettings: &frameCaptureGroupSettingsProperty{
//   						destination: &outputLocationRefProperty{
//   							destinationRefId: jsii.String("destinationRefId"),
//   						},
//   						frameCaptureCdnSettings: &frameCaptureCdnSettingsProperty{
//   							frameCaptureS3Settings: &frameCaptureS3SettingsProperty{
//   								cannedAcl: jsii.String("cannedAcl"),
//   							},
//   						},
//   					},
//   					hlsGroupSettings: &hlsGroupSettingsProperty{
//   						adMarkers: []*string{
//   							jsii.String("adMarkers"),
//   						},
//   						baseUrlContent: jsii.String("baseUrlContent"),
//   						baseUrlContent1: jsii.String("baseUrlContent1"),
//   						baseUrlManifest: jsii.String("baseUrlManifest"),
//   						baseUrlManifest1: jsii.String("baseUrlManifest1"),
//   						captionLanguageMappings: []interface{}{
//   							&captionLanguageMappingProperty{
//   								captionChannel: jsii.Number(123),
//   								languageCode: jsii.String("languageCode"),
//   								languageDescription: jsii.String("languageDescription"),
//   							},
//   						},
//   						captionLanguageSetting: jsii.String("captionLanguageSetting"),
//   						clientCache: jsii.String("clientCache"),
//   						codecSpecification: jsii.String("codecSpecification"),
//   						constantIv: jsii.String("constantIv"),
//   						destination: &outputLocationRefProperty{
//   							destinationRefId: jsii.String("destinationRefId"),
//   						},
//   						directoryStructure: jsii.String("directoryStructure"),
//   						discontinuityTags: jsii.String("discontinuityTags"),
//   						encryptionType: jsii.String("encryptionType"),
//   						hlsCdnSettings: &hlsCdnSettingsProperty{
//   							hlsAkamaiSettings: &hlsAkamaiSettingsProperty{
//   								connectionRetryInterval: jsii.Number(123),
//   								filecacheDuration: jsii.Number(123),
//   								httpTransferMode: jsii.String("httpTransferMode"),
//   								numRetries: jsii.Number(123),
//   								restartDelay: jsii.Number(123),
//   								salt: jsii.String("salt"),
//   								token: jsii.String("token"),
//   							},
//   							hlsBasicPutSettings: &hlsBasicPutSettingsProperty{
//   								connectionRetryInterval: jsii.Number(123),
//   								filecacheDuration: jsii.Number(123),
//   								numRetries: jsii.Number(123),
//   								restartDelay: jsii.Number(123),
//   							},
//   							hlsMediaStoreSettings: &hlsMediaStoreSettingsProperty{
//   								connectionRetryInterval: jsii.Number(123),
//   								filecacheDuration: jsii.Number(123),
//   								mediaStoreStorageClass: jsii.String("mediaStoreStorageClass"),
//   								numRetries: jsii.Number(123),
//   								restartDelay: jsii.Number(123),
//   							},
//   							hlsS3Settings: &hlsS3SettingsProperty{
//   								cannedAcl: jsii.String("cannedAcl"),
//   							},
//   							hlsWebdavSettings: &hlsWebdavSettingsProperty{
//   								connectionRetryInterval: jsii.Number(123),
//   								filecacheDuration: jsii.Number(123),
//   								httpTransferMode: jsii.String("httpTransferMode"),
//   								numRetries: jsii.Number(123),
//   								restartDelay: jsii.Number(123),
//   							},
//   						},
//   						hlsId3SegmentTagging: jsii.String("hlsId3SegmentTagging"),
//   						iFrameOnlyPlaylists: jsii.String("iFrameOnlyPlaylists"),
//   						incompleteSegmentBehavior: jsii.String("incompleteSegmentBehavior"),
//   						indexNSegments: jsii.Number(123),
//   						inputLossAction: jsii.String("inputLossAction"),
//   						ivInManifest: jsii.String("ivInManifest"),
//   						ivSource: jsii.String("ivSource"),
//   						keepSegments: jsii.Number(123),
//   						keyFormat: jsii.String("keyFormat"),
//   						keyFormatVersions: jsii.String("keyFormatVersions"),
//   						keyProviderSettings: &keyProviderSettingsProperty{
//   							staticKeySettings: &staticKeySettingsProperty{
//   								keyProviderServer: &inputLocationProperty{
//   									passwordParam: jsii.String("passwordParam"),
//   									uri: jsii.String("uri"),
//   									username: jsii.String("username"),
//   								},
//   								staticKeyValue: jsii.String("staticKeyValue"),
//   							},
//   						},
//   						manifestCompression: jsii.String("manifestCompression"),
//   						manifestDurationFormat: jsii.String("manifestDurationFormat"),
//   						minSegmentLength: jsii.Number(123),
//   						mode: jsii.String("mode"),
//   						outputSelection: jsii.String("outputSelection"),
//   						programDateTime: jsii.String("programDateTime"),
//   						programDateTimeClock: jsii.String("programDateTimeClock"),
//   						programDateTimePeriod: jsii.Number(123),
//   						redundantManifest: jsii.String("redundantManifest"),
//   						segmentationMode: jsii.String("segmentationMode"),
//   						segmentLength: jsii.Number(123),
//   						segmentsPerSubdirectory: jsii.Number(123),
//   						streamInfResolution: jsii.String("streamInfResolution"),
//   						timedMetadataId3Frame: jsii.String("timedMetadataId3Frame"),
//   						timedMetadataId3Period: jsii.Number(123),
//   						timestampDeltaMilliseconds: jsii.Number(123),
//   						tsFileMode: jsii.String("tsFileMode"),
//   					},
//   					mediaPackageGroupSettings: &mediaPackageGroupSettingsProperty{
//   						destination: &outputLocationRefProperty{
//   							destinationRefId: jsii.String("destinationRefId"),
//   						},
//   					},
//   					msSmoothGroupSettings: &msSmoothGroupSettingsProperty{
//   						acquisitionPointId: jsii.String("acquisitionPointId"),
//   						audioOnlyTimecodeControl: jsii.String("audioOnlyTimecodeControl"),
//   						certificateMode: jsii.String("certificateMode"),
//   						connectionRetryInterval: jsii.Number(123),
//   						destination: &outputLocationRefProperty{
//   							destinationRefId: jsii.String("destinationRefId"),
//   						},
//   						eventId: jsii.String("eventId"),
//   						eventIdMode: jsii.String("eventIdMode"),
//   						eventStopBehavior: jsii.String("eventStopBehavior"),
//   						filecacheDuration: jsii.Number(123),
//   						fragmentLength: jsii.Number(123),
//   						inputLossAction: jsii.String("inputLossAction"),
//   						numRetries: jsii.Number(123),
//   						restartDelay: jsii.Number(123),
//   						segmentationMode: jsii.String("segmentationMode"),
//   						sendDelayMs: jsii.Number(123),
//   						sparseTrackType: jsii.String("sparseTrackType"),
//   						streamManifestBehavior: jsii.String("streamManifestBehavior"),
//   						timestampOffset: jsii.String("timestampOffset"),
//   						timestampOffsetMode: jsii.String("timestampOffsetMode"),
//   					},
//   					multiplexGroupSettings: &multiplexGroupSettingsProperty{
//   					},
//   					rtmpGroupSettings: &rtmpGroupSettingsProperty{
//   						adMarkers: []*string{
//   							jsii.String("adMarkers"),
//   						},
//   						authenticationScheme: jsii.String("authenticationScheme"),
//   						cacheFullBehavior: jsii.String("cacheFullBehavior"),
//   						cacheLength: jsii.Number(123),
//   						captionData: jsii.String("captionData"),
//   						inputLossAction: jsii.String("inputLossAction"),
//   						restartDelay: jsii.Number(123),
//   					},
//   					udpGroupSettings: &udpGroupSettingsProperty{
//   						inputLossAction: jsii.String("inputLossAction"),
//   						timedMetadataId3Frame: jsii.String("timedMetadataId3Frame"),
//   						timedMetadataId3Period: jsii.Number(123),
//   					},
//   				},
//   				outputs: []interface{}{
//   					&outputProperty{
//   						audioDescriptionNames: []*string{
//   							jsii.String("audioDescriptionNames"),
//   						},
//   						captionDescriptionNames: []*string{
//   							jsii.String("captionDescriptionNames"),
//   						},
//   						outputName: jsii.String("outputName"),
//   						outputSettings: &outputSettingsProperty{
//   							archiveOutputSettings: &archiveOutputSettingsProperty{
//   								containerSettings: &archiveContainerSettingsProperty{
//   									m2TsSettings: &m2tsSettingsProperty{
//   										absentInputAudioBehavior: jsii.String("absentInputAudioBehavior"),
//   										arib: jsii.String("arib"),
//   										aribCaptionsPid: jsii.String("aribCaptionsPid"),
//   										aribCaptionsPidControl: jsii.String("aribCaptionsPidControl"),
//   										audioBufferModel: jsii.String("audioBufferModel"),
//   										audioFramesPerPes: jsii.Number(123),
//   										audioPids: jsii.String("audioPids"),
//   										audioStreamType: jsii.String("audioStreamType"),
//   										bitrate: jsii.Number(123),
//   										bufferModel: jsii.String("bufferModel"),
//   										ccDescriptor: jsii.String("ccDescriptor"),
//   										dvbNitSettings: &dvbNitSettingsProperty{
//   											networkId: jsii.Number(123),
//   											networkName: jsii.String("networkName"),
//   											repInterval: jsii.Number(123),
//   										},
//   										dvbSdtSettings: &dvbSdtSettingsProperty{
//   											outputSdt: jsii.String("outputSdt"),
//   											repInterval: jsii.Number(123),
//   											serviceName: jsii.String("serviceName"),
//   											serviceProviderName: jsii.String("serviceProviderName"),
//   										},
//   										dvbSubPids: jsii.String("dvbSubPids"),
//   										dvbTdtSettings: &dvbTdtSettingsProperty{
//   											repInterval: jsii.Number(123),
//   										},
//   										dvbTeletextPid: jsii.String("dvbTeletextPid"),
//   										ebif: jsii.String("ebif"),
//   										ebpAudioInterval: jsii.String("ebpAudioInterval"),
//   										ebpLookaheadMs: jsii.Number(123),
//   										ebpPlacement: jsii.String("ebpPlacement"),
//   										ecmPid: jsii.String("ecmPid"),
//   										esRateInPes: jsii.String("esRateInPes"),
//   										etvPlatformPid: jsii.String("etvPlatformPid"),
//   										etvSignalPid: jsii.String("etvSignalPid"),
//   										fragmentTime: jsii.Number(123),
//   										klv: jsii.String("klv"),
//   										klvDataPids: jsii.String("klvDataPids"),
//   										nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   										nullPacketBitrate: jsii.Number(123),
//   										patInterval: jsii.Number(123),
//   										pcrControl: jsii.String("pcrControl"),
//   										pcrPeriod: jsii.Number(123),
//   										pcrPid: jsii.String("pcrPid"),
//   										pmtInterval: jsii.Number(123),
//   										pmtPid: jsii.String("pmtPid"),
//   										programNum: jsii.Number(123),
//   										rateMode: jsii.String("rateMode"),
//   										scte27Pids: jsii.String("scte27Pids"),
//   										scte35Control: jsii.String("scte35Control"),
//   										scte35Pid: jsii.String("scte35Pid"),
//   										segmentationMarkers: jsii.String("segmentationMarkers"),
//   										segmentationStyle: jsii.String("segmentationStyle"),
//   										segmentationTime: jsii.Number(123),
//   										timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   										timedMetadataPid: jsii.String("timedMetadataPid"),
//   										transportStreamId: jsii.Number(123),
//   										videoPid: jsii.String("videoPid"),
//   									},
//   									rawSettings: &rawSettingsProperty{
//   									},
//   								},
//   								extension: jsii.String("extension"),
//   								nameModifier: jsii.String("nameModifier"),
//   							},
//   							frameCaptureOutputSettings: &frameCaptureOutputSettingsProperty{
//   								nameModifier: jsii.String("nameModifier"),
//   							},
//   							hlsOutputSettings: &hlsOutputSettingsProperty{
//   								h265PackagingType: jsii.String("h265PackagingType"),
//   								hlsSettings: &hlsSettingsProperty{
//   									audioOnlyHlsSettings: &audioOnlyHlsSettingsProperty{
//   										audioGroupId: jsii.String("audioGroupId"),
//   										audioOnlyImage: &inputLocationProperty{
//   											passwordParam: jsii.String("passwordParam"),
//   											uri: jsii.String("uri"),
//   											username: jsii.String("username"),
//   										},
//   										audioTrackType: jsii.String("audioTrackType"),
//   										segmentType: jsii.String("segmentType"),
//   									},
//   									fmp4HlsSettings: &fmp4HlsSettingsProperty{
//   										audioRenditionSets: jsii.String("audioRenditionSets"),
//   										nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   										timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   									},
//   									frameCaptureHlsSettings: &frameCaptureHlsSettingsProperty{
//   									},
//   									standardHlsSettings: &standardHlsSettingsProperty{
//   										audioRenditionSets: jsii.String("audioRenditionSets"),
//   										m3U8Settings: &m3u8SettingsProperty{
//   											audioFramesPerPes: jsii.Number(123),
//   											audioPids: jsii.String("audioPids"),
//   											ecmPid: jsii.String("ecmPid"),
//   											nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   											patInterval: jsii.Number(123),
//   											pcrControl: jsii.String("pcrControl"),
//   											pcrPeriod: jsii.Number(123),
//   											pcrPid: jsii.String("pcrPid"),
//   											pmtInterval: jsii.Number(123),
//   											pmtPid: jsii.String("pmtPid"),
//   											programNum: jsii.Number(123),
//   											scte35Behavior: jsii.String("scte35Behavior"),
//   											scte35Pid: jsii.String("scte35Pid"),
//   											timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   											timedMetadataPid: jsii.String("timedMetadataPid"),
//   											transportStreamId: jsii.Number(123),
//   											videoPid: jsii.String("videoPid"),
//   										},
//   									},
//   								},
//   								nameModifier: jsii.String("nameModifier"),
//   								segmentModifier: jsii.String("segmentModifier"),
//   							},
//   							mediaPackageOutputSettings: &mediaPackageOutputSettingsProperty{
//   							},
//   							msSmoothOutputSettings: &msSmoothOutputSettingsProperty{
//   								h265PackagingType: jsii.String("h265PackagingType"),
//   								nameModifier: jsii.String("nameModifier"),
//   							},
//   							multiplexOutputSettings: &multiplexOutputSettingsProperty{
//   								destination: &outputLocationRefProperty{
//   									destinationRefId: jsii.String("destinationRefId"),
//   								},
//   							},
//   							rtmpOutputSettings: &rtmpOutputSettingsProperty{
//   								certificateMode: jsii.String("certificateMode"),
//   								connectionRetryInterval: jsii.Number(123),
//   								destination: &outputLocationRefProperty{
//   									destinationRefId: jsii.String("destinationRefId"),
//   								},
//   								numRetries: jsii.Number(123),
//   							},
//   							udpOutputSettings: &udpOutputSettingsProperty{
//   								bufferMsec: jsii.Number(123),
//   								containerSettings: &udpContainerSettingsProperty{
//   									m2TsSettings: &m2tsSettingsProperty{
//   										absentInputAudioBehavior: jsii.String("absentInputAudioBehavior"),
//   										arib: jsii.String("arib"),
//   										aribCaptionsPid: jsii.String("aribCaptionsPid"),
//   										aribCaptionsPidControl: jsii.String("aribCaptionsPidControl"),
//   										audioBufferModel: jsii.String("audioBufferModel"),
//   										audioFramesPerPes: jsii.Number(123),
//   										audioPids: jsii.String("audioPids"),
//   										audioStreamType: jsii.String("audioStreamType"),
//   										bitrate: jsii.Number(123),
//   										bufferModel: jsii.String("bufferModel"),
//   										ccDescriptor: jsii.String("ccDescriptor"),
//   										dvbNitSettings: &dvbNitSettingsProperty{
//   											networkId: jsii.Number(123),
//   											networkName: jsii.String("networkName"),
//   											repInterval: jsii.Number(123),
//   										},
//   										dvbSdtSettings: &dvbSdtSettingsProperty{
//   											outputSdt: jsii.String("outputSdt"),
//   											repInterval: jsii.Number(123),
//   											serviceName: jsii.String("serviceName"),
//   											serviceProviderName: jsii.String("serviceProviderName"),
//   										},
//   										dvbSubPids: jsii.String("dvbSubPids"),
//   										dvbTdtSettings: &dvbTdtSettingsProperty{
//   											repInterval: jsii.Number(123),
//   										},
//   										dvbTeletextPid: jsii.String("dvbTeletextPid"),
//   										ebif: jsii.String("ebif"),
//   										ebpAudioInterval: jsii.String("ebpAudioInterval"),
//   										ebpLookaheadMs: jsii.Number(123),
//   										ebpPlacement: jsii.String("ebpPlacement"),
//   										ecmPid: jsii.String("ecmPid"),
//   										esRateInPes: jsii.String("esRateInPes"),
//   										etvPlatformPid: jsii.String("etvPlatformPid"),
//   										etvSignalPid: jsii.String("etvSignalPid"),
//   										fragmentTime: jsii.Number(123),
//   										klv: jsii.String("klv"),
//   										klvDataPids: jsii.String("klvDataPids"),
//   										nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   										nullPacketBitrate: jsii.Number(123),
//   										patInterval: jsii.Number(123),
//   										pcrControl: jsii.String("pcrControl"),
//   										pcrPeriod: jsii.Number(123),
//   										pcrPid: jsii.String("pcrPid"),
//   										pmtInterval: jsii.Number(123),
//   										pmtPid: jsii.String("pmtPid"),
//   										programNum: jsii.Number(123),
//   										rateMode: jsii.String("rateMode"),
//   										scte27Pids: jsii.String("scte27Pids"),
//   										scte35Control: jsii.String("scte35Control"),
//   										scte35Pid: jsii.String("scte35Pid"),
//   										segmentationMarkers: jsii.String("segmentationMarkers"),
//   										segmentationStyle: jsii.String("segmentationStyle"),
//   										segmentationTime: jsii.Number(123),
//   										timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   										timedMetadataPid: jsii.String("timedMetadataPid"),
//   										transportStreamId: jsii.Number(123),
//   										videoPid: jsii.String("videoPid"),
//   									},
//   								},
//   								destination: &outputLocationRefProperty{
//   									destinationRefId: jsii.String("destinationRefId"),
//   								},
//   								fecOutputSettings: &fecOutputSettingsProperty{
//   									columnDepth: jsii.Number(123),
//   									includeFec: jsii.String("includeFec"),
//   									rowLength: jsii.Number(123),
//   								},
//   							},
//   						},
//   						videoDescriptionName: jsii.String("videoDescriptionName"),
//   					},
//   				},
//   			},
//   		},
//   		timecodeConfig: &timecodeConfigProperty{
//   			source: jsii.String("source"),
//   			syncThreshold: jsii.Number(123),
//   		},
//   		videoDescriptions: []interface{}{
//   			&videoDescriptionProperty{
//   				codecSettings: &videoCodecSettingsProperty{
//   					frameCaptureSettings: &frameCaptureSettingsProperty{
//   						captureInterval: jsii.Number(123),
//   						captureIntervalUnits: jsii.String("captureIntervalUnits"),
//   					},
//   					h264Settings: &h264SettingsProperty{
//   						adaptiveQuantization: jsii.String("adaptiveQuantization"),
//   						afdSignaling: jsii.String("afdSignaling"),
//   						bitrate: jsii.Number(123),
//   						bufFillPct: jsii.Number(123),
//   						bufSize: jsii.Number(123),
//   						colorMetadata: jsii.String("colorMetadata"),
//   						colorSpaceSettings: &h264ColorSpaceSettingsProperty{
//   							colorSpacePassthroughSettings: &colorSpacePassthroughSettingsProperty{
//   							},
//   							rec601Settings: &rec601SettingsProperty{
//   							},
//   							rec709Settings: &rec709SettingsProperty{
//   							},
//   						},
//   						entropyEncoding: jsii.String("entropyEncoding"),
//   						filterSettings: &h264FilterSettingsProperty{
//   							temporalFilterSettings: &temporalFilterSettingsProperty{
//   								postFilterSharpening: jsii.String("postFilterSharpening"),
//   								strength: jsii.String("strength"),
//   							},
//   						},
//   						fixedAfd: jsii.String("fixedAfd"),
//   						flickerAq: jsii.String("flickerAq"),
//   						forceFieldPictures: jsii.String("forceFieldPictures"),
//   						framerateControl: jsii.String("framerateControl"),
//   						framerateDenominator: jsii.Number(123),
//   						framerateNumerator: jsii.Number(123),
//   						gopBReference: jsii.String("gopBReference"),
//   						gopClosedCadence: jsii.Number(123),
//   						gopNumBFrames: jsii.Number(123),
//   						gopSize: jsii.Number(123),
//   						gopSizeUnits: jsii.String("gopSizeUnits"),
//   						level: jsii.String("level"),
//   						lookAheadRateControl: jsii.String("lookAheadRateControl"),
//   						maxBitrate: jsii.Number(123),
//   						minIInterval: jsii.Number(123),
//   						numRefFrames: jsii.Number(123),
//   						parControl: jsii.String("parControl"),
//   						parDenominator: jsii.Number(123),
//   						parNumerator: jsii.Number(123),
//   						profile: jsii.String("profile"),
//   						qualityLevel: jsii.String("qualityLevel"),
//   						qvbrQualityLevel: jsii.Number(123),
//   						rateControlMode: jsii.String("rateControlMode"),
//   						scanType: jsii.String("scanType"),
//   						sceneChangeDetect: jsii.String("sceneChangeDetect"),
//   						slices: jsii.Number(123),
//   						softness: jsii.Number(123),
//   						spatialAq: jsii.String("spatialAq"),
//   						subgopLength: jsii.String("subgopLength"),
//   						syntax: jsii.String("syntax"),
//   						temporalAq: jsii.String("temporalAq"),
//   						timecodeInsertion: jsii.String("timecodeInsertion"),
//   					},
//   					h265Settings: &h265SettingsProperty{
//   						adaptiveQuantization: jsii.String("adaptiveQuantization"),
//   						afdSignaling: jsii.String("afdSignaling"),
//   						alternativeTransferFunction: jsii.String("alternativeTransferFunction"),
//   						bitrate: jsii.Number(123),
//   						bufSize: jsii.Number(123),
//   						colorMetadata: jsii.String("colorMetadata"),
//   						colorSpaceSettings: &h265ColorSpaceSettingsProperty{
//   							colorSpacePassthroughSettings: &colorSpacePassthroughSettingsProperty{
//   							},
//   							hdr10Settings: &hdr10SettingsProperty{
//   								maxCll: jsii.Number(123),
//   								maxFall: jsii.Number(123),
//   							},
//   							rec601Settings: &rec601SettingsProperty{
//   							},
//   							rec709Settings: &rec709SettingsProperty{
//   							},
//   						},
//   						filterSettings: &h265FilterSettingsProperty{
//   							temporalFilterSettings: &temporalFilterSettingsProperty{
//   								postFilterSharpening: jsii.String("postFilterSharpening"),
//   								strength: jsii.String("strength"),
//   							},
//   						},
//   						fixedAfd: jsii.String("fixedAfd"),
//   						flickerAq: jsii.String("flickerAq"),
//   						framerateDenominator: jsii.Number(123),
//   						framerateNumerator: jsii.Number(123),
//   						gopClosedCadence: jsii.Number(123),
//   						gopSize: jsii.Number(123),
//   						gopSizeUnits: jsii.String("gopSizeUnits"),
//   						level: jsii.String("level"),
//   						lookAheadRateControl: jsii.String("lookAheadRateControl"),
//   						maxBitrate: jsii.Number(123),
//   						minIInterval: jsii.Number(123),
//   						parDenominator: jsii.Number(123),
//   						parNumerator: jsii.Number(123),
//   						profile: jsii.String("profile"),
//   						qvbrQualityLevel: jsii.Number(123),
//   						rateControlMode: jsii.String("rateControlMode"),
//   						scanType: jsii.String("scanType"),
//   						sceneChangeDetect: jsii.String("sceneChangeDetect"),
//   						slices: jsii.Number(123),
//   						tier: jsii.String("tier"),
//   						timecodeInsertion: jsii.String("timecodeInsertion"),
//   					},
//   					mpeg2Settings: &mpeg2SettingsProperty{
//   						adaptiveQuantization: jsii.String("adaptiveQuantization"),
//   						afdSignaling: jsii.String("afdSignaling"),
//   						colorMetadata: jsii.String("colorMetadata"),
//   						colorSpace: jsii.String("colorSpace"),
//   						displayAspectRatio: jsii.String("displayAspectRatio"),
//   						filterSettings: &mpeg2FilterSettingsProperty{
//   							temporalFilterSettings: &temporalFilterSettingsProperty{
//   								postFilterSharpening: jsii.String("postFilterSharpening"),
//   								strength: jsii.String("strength"),
//   							},
//   						},
//   						fixedAfd: jsii.String("fixedAfd"),
//   						framerateDenominator: jsii.Number(123),
//   						framerateNumerator: jsii.Number(123),
//   						gopClosedCadence: jsii.Number(123),
//   						gopNumBFrames: jsii.Number(123),
//   						gopSize: jsii.Number(123),
//   						gopSizeUnits: jsii.String("gopSizeUnits"),
//   						scanType: jsii.String("scanType"),
//   						subgopLength: jsii.String("subgopLength"),
//   						timecodeInsertion: jsii.String("timecodeInsertion"),
//   					},
//   				},
//   				height: jsii.Number(123),
//   				name: jsii.String("name"),
//   				respondToAfd: jsii.String("respondToAfd"),
//   				scalingBehavior: jsii.String("scalingBehavior"),
//   				sharpness: jsii.Number(123),
//   				width: jsii.Number(123),
//   			},
//   		},
//   	},
//   	inputAttachments: []interface{}{
//   		&inputAttachmentProperty{
//   			automaticInputFailoverSettings: &automaticInputFailoverSettingsProperty{
//   				errorClearTimeMsec: jsii.Number(123),
//   				failoverConditions: []interface{}{
//   					&failoverConditionProperty{
//   						failoverConditionSettings: &failoverConditionSettingsProperty{
//   							audioSilenceSettings: &audioSilenceFailoverSettingsProperty{
//   								audioSelectorName: jsii.String("audioSelectorName"),
//   								audioSilenceThresholdMsec: jsii.Number(123),
//   							},
//   							inputLossSettings: &inputLossFailoverSettingsProperty{
//   								inputLossThresholdMsec: jsii.Number(123),
//   							},
//   							videoBlackSettings: &videoBlackFailoverSettingsProperty{
//   								blackDetectThreshold: jsii.Number(123),
//   								videoBlackThresholdMsec: jsii.Number(123),
//   							},
//   						},
//   					},
//   				},
//   				inputPreference: jsii.String("inputPreference"),
//   				secondaryInputId: jsii.String("secondaryInputId"),
//   			},
//   			inputAttachmentName: jsii.String("inputAttachmentName"),
//   			inputId: jsii.String("inputId"),
//   			inputSettings: &inputSettingsProperty{
//   				audioSelectors: []interface{}{
//   					&audioSelectorProperty{
//   						name: jsii.String("name"),
//   						selectorSettings: &audioSelectorSettingsProperty{
//   							audioHlsRenditionSelection: &audioHlsRenditionSelectionProperty{
//   								groupId: jsii.String("groupId"),
//   								name: jsii.String("name"),
//   							},
//   							audioLanguageSelection: &audioLanguageSelectionProperty{
//   								languageCode: jsii.String("languageCode"),
//   								languageSelectionPolicy: jsii.String("languageSelectionPolicy"),
//   							},
//   							audioPidSelection: &audioPidSelectionProperty{
//   								pid: jsii.Number(123),
//   							},
//   							audioTrackSelection: &audioTrackSelectionProperty{
//   								tracks: []interface{}{
//   									&audioTrackProperty{
//   										track: jsii.Number(123),
//   									},
//   								},
//   							},
//   						},
//   					},
//   				},
//   				captionSelectors: []interface{}{
//   					&captionSelectorProperty{
//   						languageCode: jsii.String("languageCode"),
//   						name: jsii.String("name"),
//   						selectorSettings: &captionSelectorSettingsProperty{
//   							ancillarySourceSettings: &ancillarySourceSettingsProperty{
//   								sourceAncillaryChannelNumber: jsii.Number(123),
//   							},
//   							aribSourceSettings: &aribSourceSettingsProperty{
//   							},
//   							dvbSubSourceSettings: &dvbSubSourceSettingsProperty{
//   								ocrLanguage: jsii.String("ocrLanguage"),
//   								pid: jsii.Number(123),
//   							},
//   							embeddedSourceSettings: &embeddedSourceSettingsProperty{
//   								convert608To708: jsii.String("convert608To708"),
//   								scte20Detection: jsii.String("scte20Detection"),
//   								source608ChannelNumber: jsii.Number(123),
//   								source608TrackNumber: jsii.Number(123),
//   							},
//   							scte20SourceSettings: &scte20SourceSettingsProperty{
//   								convert608To708: jsii.String("convert608To708"),
//   								source608ChannelNumber: jsii.Number(123),
//   							},
//   							scte27SourceSettings: &scte27SourceSettingsProperty{
//   								ocrLanguage: jsii.String("ocrLanguage"),
//   								pid: jsii.Number(123),
//   							},
//   							teletextSourceSettings: &teletextSourceSettingsProperty{
//   								outputRectangle: &captionRectangleProperty{
//   									height: jsii.Number(123),
//   									leftOffset: jsii.Number(123),
//   									topOffset: jsii.Number(123),
//   									width: jsii.Number(123),
//   								},
//   								pageNumber: jsii.String("pageNumber"),
//   							},
//   						},
//   					},
//   				},
//   				deblockFilter: jsii.String("deblockFilter"),
//   				denoiseFilter: jsii.String("denoiseFilter"),
//   				filterStrength: jsii.Number(123),
//   				inputFilter: jsii.String("inputFilter"),
//   				networkInputSettings: &networkInputSettingsProperty{
//   					hlsInputSettings: &hlsInputSettingsProperty{
//   						bandwidth: jsii.Number(123),
//   						bufferSegments: jsii.Number(123),
//   						retries: jsii.Number(123),
//   						retryInterval: jsii.Number(123),
//   						scte35Source: jsii.String("scte35Source"),
//   					},
//   					serverValidation: jsii.String("serverValidation"),
//   				},
//   				scte35Pid: jsii.Number(123),
//   				smpte2038DataPreference: jsii.String("smpte2038DataPreference"),
//   				sourceEndBehavior: jsii.String("sourceEndBehavior"),
//   				videoSelector: &videoSelectorProperty{
//   					colorSpace: jsii.String("colorSpace"),
//   					colorSpaceSettings: &videoSelectorColorSpaceSettingsProperty{
//   						hdr10Settings: &hdr10SettingsProperty{
//   							maxCll: jsii.Number(123),
//   							maxFall: jsii.Number(123),
//   						},
//   					},
//   					colorSpaceUsage: jsii.String("colorSpaceUsage"),
//   					selectorSettings: &videoSelectorSettingsProperty{
//   						videoSelectorPid: &videoSelectorPidProperty{
//   							pid: jsii.Number(123),
//   						},
//   						videoSelectorProgramId: &videoSelectorProgramIdProperty{
//   							programId: jsii.Number(123),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	inputSpecification: &inputSpecificationProperty{
//   		codec: jsii.String("codec"),
//   		maximumBitrate: jsii.String("maximumBitrate"),
//   		resolution: jsii.String("resolution"),
//   	},
//   	logLevel: jsii.String("logLevel"),
//   	name: jsii.String("name"),
//   	roleArn: jsii.String("roleArn"),
//   	tags: tags,
//   	vpc: &vpcOutputSettingsProperty{
//   		publicAddressAllocationIds: []*string{
//   			jsii.String("publicAddressAllocationIds"),
//   		},
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
type CfnChannelProps struct {
	// Specification of CDI inputs for this channel.
	CdiInputSpecification interface{} `field:"optional" json:"cdiInputSpecification" yaml:"cdiInputSpecification"`
	// The class for this channel.
	//
	// For a channel with two pipelines, the class is STANDARD. For a channel with one pipeline, the class is SINGLE_PIPELINE.
	ChannelClass *string `field:"optional" json:"channelClass" yaml:"channelClass"`
	// The settings that identify the destination for the outputs in this MediaLive output package.
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// The encoding configuration for the output content.
	EncoderSettings interface{} `field:"optional" json:"encoderSettings" yaml:"encoderSettings"`
	// The list of input attachments for the channel.
	InputAttachments interface{} `field:"optional" json:"inputAttachments" yaml:"inputAttachments"`
	// The input specification for this channel.
	//
	// It specifies the key characteristics of the inputs for this channel: the maximum bitrate, the resolution, and the codec.
	InputSpecification interface{} `field:"optional" json:"inputSpecification" yaml:"inputSpecification"`
	// The verbosity for logging activity for this channel.
	//
	// Charges for logging (which are generated through Amazon CloudWatch Logging) are higher for higher verbosities.
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// A name for this audio selector.
	//
	// The AudioDescription (in an output) references this name in order to identify a specific input audio to include in that output.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The IAM role for MediaLive to assume when running this channel.
	//
	// The role is identified by its ARN.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// A collection of tags for this channel.
	//
	// Each tag is a key-value pair.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Settings to enable VPC mode in the channel, so that the endpoints for all outputs are in your VPC.
	Vpc interface{} `field:"optional" json:"vpc" yaml:"vpc"`
}

// A CloudFormation `AWS::MediaLive::Input`.
//
// The AWS::MediaLive::Input resource is a MediaLive resource type that creates an input.
//
// A MediaLive input holds information that describes how the MediaLive channel is connected to the upstream system that is providing the source content that is to be transcoded.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnInput := awscdk.Aws_medialive.NewCfnInput(this, jsii.String("MyCfnInput"), &cfnInputProps{
//   	destinations: []interface{}{
//   		&inputDestinationRequestProperty{
//   			streamName: jsii.String("streamName"),
//   		},
//   	},
//   	inputDevices: []interface{}{
//   		&inputDeviceSettingsProperty{
//   			id: jsii.String("id"),
//   		},
//   	},
//   	inputSecurityGroups: []*string{
//   		jsii.String("inputSecurityGroups"),
//   	},
//   	mediaConnectFlows: []interface{}{
//   		&mediaConnectFlowRequestProperty{
//   			flowArn: jsii.String("flowArn"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	roleArn: jsii.String("roleArn"),
//   	sources: []interface{}{
//   		&inputSourceRequestProperty{
//   			passwordParam: jsii.String("passwordParam"),
//   			url: jsii.String("url"),
//   			username: jsii.String("username"),
//   		},
//   	},
//   	tags: tags,
//   	type: jsii.String("type"),
//   	vpc: &inputVpcRequestProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   })
//
type CfnInput interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ARN of the MediaLive input.
	//
	// For example: arn:aws:medialive:us-west-1:111122223333:medialive:input:1234567. MediaLive creates this ARN when it creates the input.
	AttrArn() *string
	// For a push input, the the destination or destinations for the input.
	//
	// The destinations are the URLs of locations on MediaLive where the upstream system pushes the content to, for this input. MediaLive creates these addresses when it creates the input.
	AttrDestinations() *[]*string
	// For a pull input, the source or sources for the input.
	//
	// The sources are the URLs of locations on the upstream system where MediaLive pulls the content from, for this input. You included these URLs in the create request.
	AttrSources() *[]*string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Settings that apply only if the input is a push type of input.
	Destinations() interface{}
	SetDestinations(val interface{})
	// Settings that apply only if the input is an Elemental Link input.
	InputDevices() interface{}
	SetInputDevices(val interface{})
	// The list of input security groups (referenced by IDs) to attach to the input if the input is a push type.
	InputSecurityGroups() *[]*string
	SetInputSecurityGroups(val *[]*string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// Settings that apply only if the input is a MediaConnect input.
	MediaConnectFlows() interface{}
	SetMediaConnectFlows(val interface{})
	// A name for the input.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The IAM role for MediaLive to assume when creating a MediaConnect input or Amazon VPC input.
	//
	// This doesn't apply to other types of inputs. The role is identified by its ARN.
	RoleArn() *string
	SetRoleArn(val *string)
	// Settings that apply only if the input is a pull type of input.
	Sources() interface{}
	SetSources(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// A collection of tags for this input.
	//
	// Each tag is a key-value pair.
	Tags() awscdk.TagManager
	// The type for this input.
	Type() *string
	SetType(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Settings that apply only if the input is an push input where the source is on Amazon VPC.
	Vpc() interface{}
	SetVpc(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnInput
type jsiiProxy_CfnInput struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnInput) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) AttrDestinations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrDestinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) AttrSources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) Destinations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) InputDevices() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputDevices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) InputSecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) MediaConnectFlows() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mediaConnectFlows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) Sources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) Vpc() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


// Create a new `AWS::MediaLive::Input`.
func NewCfnInput(scope constructs.Construct, id *string, props *CfnInputProps) CfnInput {
	_init_.Initialize()

	j := jsiiProxy_CfnInput{}

	_jsii_.Create(
		"aws-cdk-lib.aws_medialive.CfnInput",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MediaLive::Input`.
func NewCfnInput_Override(c CfnInput, scope constructs.Construct, id *string, props *CfnInputProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_medialive.CfnInput",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnInput) SetDestinations(val interface{}) {
	_jsii_.Set(
		j,
		"destinations",
		val,
	)
}

func (j *jsiiProxy_CfnInput) SetInputDevices(val interface{}) {
	_jsii_.Set(
		j,
		"inputDevices",
		val,
	)
}

func (j *jsiiProxy_CfnInput) SetInputSecurityGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"inputSecurityGroups",
		val,
	)
}

func (j *jsiiProxy_CfnInput) SetMediaConnectFlows(val interface{}) {
	_jsii_.Set(
		j,
		"mediaConnectFlows",
		val,
	)
}

func (j *jsiiProxy_CfnInput) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnInput) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnInput) SetSources(val interface{}) {
	_jsii_.Set(
		j,
		"sources",
		val,
	)
}

func (j *jsiiProxy_CfnInput) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_CfnInput) SetVpc(val interface{}) {
	_jsii_.Set(
		j,
		"vpc",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnInput_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_medialive.CfnInput",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnInput_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_medialive.CfnInput",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnInput_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_medialive.CfnInput",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInput_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_medialive.CfnInput",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInput) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnInput) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnInput) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnInput) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnInput) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnInput) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnInput) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnInput) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInput) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInput) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnInput) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnInput) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInput) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInput) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInput) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Settings that apply only if the input is a push type of input.
//
// The parent of this entity is Input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputDestinationRequestProperty := &inputDestinationRequestProperty{
//   	streamName: jsii.String("streamName"),
//   }
//
type CfnInput_InputDestinationRequestProperty struct {
	// The stream name (application name/application instance) for the location the RTMP source content will be pushed to in MediaLive.
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
}

// This entity is not used.
//
// Ignore it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputDeviceRequestProperty := &inputDeviceRequestProperty{
//   	id: jsii.String("id"),
//   }
//
type CfnInput_InputDeviceRequestProperty struct {
	// This property is not used.
	//
	// Ignore it.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

// Settings that apply only if the input is an Elemental Link input.
//
// The parent of this entity is Input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputDeviceSettingsProperty := &inputDeviceSettingsProperty{
//   	id: jsii.String("id"),
//   }
//
type CfnInput_InputDeviceSettingsProperty struct {
	// The unique ID for the device.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

// Settings that apply only if the input is a pull type of input.
//
// The parent of this entity is Input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputSourceRequestProperty := &inputSourceRequestProperty{
//   	passwordParam: jsii.String("passwordParam"),
//   	url: jsii.String("url"),
//   	username: jsii.String("username"),
//   }
//
type CfnInput_InputSourceRequestProperty struct {
	// The password parameter that holds the password for accessing the upstream system.
	//
	// The password parameter applies only if the upstream system requires credentials.
	PasswordParam *string `field:"optional" json:"passwordParam" yaml:"passwordParam"`
	// For a pull input, the URL where MediaLive pulls the source content from.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// The user name to connect to the upstream system.
	//
	// The user name applies only if the upstream system requires credentials.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

// Settings that apply only if the input is an push input where the source is on Amazon VPC.
//
// The parent of this entity is Input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputVpcRequestProperty := &inputVpcRequestProperty{
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   }
//
type CfnInput_InputVpcRequestProperty struct {
	// The list of up to five VPC security group IDs to attach to the input VPC network interfaces.
	//
	// The security groups require subnet IDs. If none are specified, MediaLive uses the VPC default security group.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The list of two VPC subnet IDs from the same VPC.
	//
	// You must associate subnet IDs to two unique Availability Zones.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

// Settings that apply only if the input is a MediaConnect input.
//
// The parent of this entity is Input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mediaConnectFlowRequestProperty := &mediaConnectFlowRequestProperty{
//   	flowArn: jsii.String("flowArn"),
//   }
//
type CfnInput_MediaConnectFlowRequestProperty struct {
	// The ARN of one or two MediaConnect flows that are the sources for this MediaConnect input.
	FlowArn *string `field:"optional" json:"flowArn" yaml:"flowArn"`
}

// Properties for defining a `CfnInput`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnInputProps := &cfnInputProps{
//   	destinations: []interface{}{
//   		&inputDestinationRequestProperty{
//   			streamName: jsii.String("streamName"),
//   		},
//   	},
//   	inputDevices: []interface{}{
//   		&inputDeviceSettingsProperty{
//   			id: jsii.String("id"),
//   		},
//   	},
//   	inputSecurityGroups: []*string{
//   		jsii.String("inputSecurityGroups"),
//   	},
//   	mediaConnectFlows: []interface{}{
//   		&mediaConnectFlowRequestProperty{
//   			flowArn: jsii.String("flowArn"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	roleArn: jsii.String("roleArn"),
//   	sources: []interface{}{
//   		&inputSourceRequestProperty{
//   			passwordParam: jsii.String("passwordParam"),
//   			url: jsii.String("url"),
//   			username: jsii.String("username"),
//   		},
//   	},
//   	tags: tags,
//   	type: jsii.String("type"),
//   	vpc: &inputVpcRequestProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
type CfnInputProps struct {
	// Settings that apply only if the input is a push type of input.
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// Settings that apply only if the input is an Elemental Link input.
	InputDevices interface{} `field:"optional" json:"inputDevices" yaml:"inputDevices"`
	// The list of input security groups (referenced by IDs) to attach to the input if the input is a push type.
	InputSecurityGroups *[]*string `field:"optional" json:"inputSecurityGroups" yaml:"inputSecurityGroups"`
	// Settings that apply only if the input is a MediaConnect input.
	MediaConnectFlows interface{} `field:"optional" json:"mediaConnectFlows" yaml:"mediaConnectFlows"`
	// A name for the input.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The IAM role for MediaLive to assume when creating a MediaConnect input or Amazon VPC input.
	//
	// This doesn't apply to other types of inputs. The role is identified by its ARN.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Settings that apply only if the input is a pull type of input.
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
	// A collection of tags for this input.
	//
	// Each tag is a key-value pair.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The type for this input.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Settings that apply only if the input is an push input where the source is on Amazon VPC.
	Vpc interface{} `field:"optional" json:"vpc" yaml:"vpc"`
}

// A CloudFormation `AWS::MediaLive::InputSecurityGroup`.
//
// The AWS::MediaLive::InputSecurityGroup is a MediaLive resource type that creates an input security group.
//
// A MediaLive input security group is associated with a MediaLive input. The input security group is an "allow list" of IP addresses that controls whether an external IP address can push content to the associated MediaLive input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnInputSecurityGroup := awscdk.Aws_medialive.NewCfnInputSecurityGroup(this, jsii.String("MyCfnInputSecurityGroup"), &cfnInputSecurityGroupProps{
//   	tags: tags,
//   	whitelistRules: []interface{}{
//   		&inputWhitelistRuleCidrProperty{
//   			cidr: jsii.String("cidr"),
//   		},
//   	},
//   })
//
type CfnInputSecurityGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ARN of the MediaLive input security group.
	//
	// For example: arn:aws:medialive:us-west-1:111122223333:medialive:inputSecurityGroup:1234567.
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// A collection of tags for this input security group.
	//
	// Each tag is a key-value pair.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// The list of IPv4 CIDR addresses to include in the input security group as "allowed" addresses.
	WhitelistRules() interface{}
	SetWhitelistRules(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnInputSecurityGroup
type jsiiProxy_CfnInputSecurityGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnInputSecurityGroup) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInputSecurityGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInputSecurityGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInputSecurityGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInputSecurityGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInputSecurityGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInputSecurityGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInputSecurityGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInputSecurityGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInputSecurityGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInputSecurityGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInputSecurityGroup) WhitelistRules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"whitelistRules",
		&returns,
	)
	return returns
}


// Create a new `AWS::MediaLive::InputSecurityGroup`.
func NewCfnInputSecurityGroup(scope constructs.Construct, id *string, props *CfnInputSecurityGroupProps) CfnInputSecurityGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnInputSecurityGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_medialive.CfnInputSecurityGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MediaLive::InputSecurityGroup`.
func NewCfnInputSecurityGroup_Override(c CfnInputSecurityGroup, scope constructs.Construct, id *string, props *CfnInputSecurityGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_medialive.CfnInputSecurityGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnInputSecurityGroup) SetWhitelistRules(val interface{}) {
	_jsii_.Set(
		j,
		"whitelistRules",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnInputSecurityGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_medialive.CfnInputSecurityGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnInputSecurityGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_medialive.CfnInputSecurityGroup",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnInputSecurityGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_medialive.CfnInputSecurityGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInputSecurityGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_medialive.CfnInputSecurityGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInputSecurityGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnInputSecurityGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnInputSecurityGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnInputSecurityGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnInputSecurityGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnInputSecurityGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnInputSecurityGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnInputSecurityGroup) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInputSecurityGroup) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInputSecurityGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnInputSecurityGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnInputSecurityGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInputSecurityGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInputSecurityGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInputSecurityGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// An IPv4 CIDR range to include in this input security group.
//
// The parent of this entity is InputSecurityGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputWhitelistRuleCidrProperty := &inputWhitelistRuleCidrProperty{
//   	cidr: jsii.String("cidr"),
//   }
//
type CfnInputSecurityGroup_InputWhitelistRuleCidrProperty struct {
	// An IPv4 CIDR range to include in this input security group.
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
}

// Properties for defining a `CfnInputSecurityGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnInputSecurityGroupProps := &cfnInputSecurityGroupProps{
//   	tags: tags,
//   	whitelistRules: []interface{}{
//   		&inputWhitelistRuleCidrProperty{
//   			cidr: jsii.String("cidr"),
//   		},
//   	},
//   }
//
type CfnInputSecurityGroupProps struct {
	// A collection of tags for this input security group.
	//
	// Each tag is a key-value pair.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The list of IPv4 CIDR addresses to include in the input security group as "allowed" addresses.
	WhitelistRules interface{} `field:"optional" json:"whitelistRules" yaml:"whitelistRules"`
}

