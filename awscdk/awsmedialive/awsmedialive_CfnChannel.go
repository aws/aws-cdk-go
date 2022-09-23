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
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
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

func (j *jsiiProxy_CfnChannel) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
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

	if err := validateNewCfnChannelParameters(scope, id, props); err != nil {
		panic(err)
	}
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

func (j *jsiiProxy_CfnChannel)SetCdiInputSpecification(val interface{}) {
	if err := j.validateSetCdiInputSpecificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdiInputSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnChannel)SetChannelClass(val *string) {
	_jsii_.Set(
		j,
		"channelClass",
		val,
	)
}

func (j *jsiiProxy_CfnChannel)SetDestinations(val interface{}) {
	if err := j.validateSetDestinationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinations",
		val,
	)
}

func (j *jsiiProxy_CfnChannel)SetEncoderSettings(val interface{}) {
	if err := j.validateSetEncoderSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encoderSettings",
		val,
	)
}

func (j *jsiiProxy_CfnChannel)SetInputAttachments(val interface{}) {
	if err := j.validateSetInputAttachmentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputAttachments",
		val,
	)
}

func (j *jsiiProxy_CfnChannel)SetInputSpecification(val interface{}) {
	if err := j.validateSetInputSpecificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnChannel)SetLogLevel(val *string) {
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_CfnChannel)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnChannel)SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnChannel)SetVpc(val interface{}) {
	if err := j.validateSetVpcParameters(val); err != nil {
		panic(err)
	}
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

	if err := validateCfnChannel_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
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

	if err := validateCfnChannel_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
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
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnChannel_IsConstructParameters(x); err != nil {
		panic(err)
	}
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
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnChannel) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnChannel) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnChannel) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnChannel) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnChannel) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnChannel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnChannel) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
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
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
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
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnChannel) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnChannel) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
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
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

