package awsmedialive


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

