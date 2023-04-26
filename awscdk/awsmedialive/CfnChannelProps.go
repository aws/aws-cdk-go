package awsmedialive


// Properties for defining a `CfnChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnChannelProps := &CfnChannelProps{
//   	CdiInputSpecification: &CdiInputSpecificationProperty{
//   		Resolution: jsii.String("resolution"),
//   	},
//   	ChannelClass: jsii.String("channelClass"),
//   	Destinations: []interface{}{
//   		&OutputDestinationProperty{
//   			Id: jsii.String("id"),
//   			MediaPackageSettings: []interface{}{
//   				&MediaPackageOutputDestinationSettingsProperty{
//   					ChannelId: jsii.String("channelId"),
//   				},
//   			},
//   			MultiplexSettings: &MultiplexProgramChannelDestinationSettingsProperty{
//   				MultiplexId: jsii.String("multiplexId"),
//   				ProgramName: jsii.String("programName"),
//   			},
//   			Settings: []interface{}{
//   				&OutputDestinationSettingsProperty{
//   					PasswordParam: jsii.String("passwordParam"),
//   					StreamName: jsii.String("streamName"),
//   					Url: jsii.String("url"),
//   					Username: jsii.String("username"),
//   				},
//   			},
//   		},
//   	},
//   	EncoderSettings: &EncoderSettingsProperty{
//   		AudioDescriptions: []interface{}{
//   			&AudioDescriptionProperty{
//   				AudioNormalizationSettings: &AudioNormalizationSettingsProperty{
//   					Algorithm: jsii.String("algorithm"),
//   					AlgorithmControl: jsii.String("algorithmControl"),
//   					TargetLkfs: jsii.Number(123),
//   				},
//   				AudioSelectorName: jsii.String("audioSelectorName"),
//   				AudioType: jsii.String("audioType"),
//   				AudioTypeControl: jsii.String("audioTypeControl"),
//   				AudioWatermarkingSettings: &AudioWatermarkSettingsProperty{
//   					NielsenWatermarksSettings: &NielsenWatermarksSettingsProperty{
//   						NielsenCbetSettings: &NielsenCBETProperty{
//   							CbetCheckDigitString: jsii.String("cbetCheckDigitString"),
//   							CbetStepaside: jsii.String("cbetStepaside"),
//   							Csid: jsii.String("csid"),
//   						},
//   						NielsenDistributionType: jsii.String("nielsenDistributionType"),
//   						NielsenNaesIiNwSettings: &NielsenNaesIiNwProperty{
//   							CheckDigitString: jsii.String("checkDigitString"),
//   							Sid: jsii.Number(123),
//   						},
//   					},
//   				},
//   				CodecSettings: &AudioCodecSettingsProperty{
//   					AacSettings: &AacSettingsProperty{
//   						Bitrate: jsii.Number(123),
//   						CodingMode: jsii.String("codingMode"),
//   						InputType: jsii.String("inputType"),
//   						Profile: jsii.String("profile"),
//   						RateControlMode: jsii.String("rateControlMode"),
//   						RawFormat: jsii.String("rawFormat"),
//   						SampleRate: jsii.Number(123),
//   						Spec: jsii.String("spec"),
//   						VbrQuality: jsii.String("vbrQuality"),
//   					},
//   					Ac3Settings: &Ac3SettingsProperty{
//   						Bitrate: jsii.Number(123),
//   						BitstreamMode: jsii.String("bitstreamMode"),
//   						CodingMode: jsii.String("codingMode"),
//   						Dialnorm: jsii.Number(123),
//   						DrcProfile: jsii.String("drcProfile"),
//   						LfeFilter: jsii.String("lfeFilter"),
//   						MetadataControl: jsii.String("metadataControl"),
//   					},
//   					Eac3Settings: &Eac3SettingsProperty{
//   						AttenuationControl: jsii.String("attenuationControl"),
//   						Bitrate: jsii.Number(123),
//   						BitstreamMode: jsii.String("bitstreamMode"),
//   						CodingMode: jsii.String("codingMode"),
//   						DcFilter: jsii.String("dcFilter"),
//   						Dialnorm: jsii.Number(123),
//   						DrcLine: jsii.String("drcLine"),
//   						DrcRf: jsii.String("drcRf"),
//   						LfeControl: jsii.String("lfeControl"),
//   						LfeFilter: jsii.String("lfeFilter"),
//   						LoRoCenterMixLevel: jsii.Number(123),
//   						LoRoSurroundMixLevel: jsii.Number(123),
//   						LtRtCenterMixLevel: jsii.Number(123),
//   						LtRtSurroundMixLevel: jsii.Number(123),
//   						MetadataControl: jsii.String("metadataControl"),
//   						PassthroughControl: jsii.String("passthroughControl"),
//   						PhaseControl: jsii.String("phaseControl"),
//   						StereoDownmix: jsii.String("stereoDownmix"),
//   						SurroundExMode: jsii.String("surroundExMode"),
//   						SurroundMode: jsii.String("surroundMode"),
//   					},
//   					Mp2Settings: &Mp2SettingsProperty{
//   						Bitrate: jsii.Number(123),
//   						CodingMode: jsii.String("codingMode"),
//   						SampleRate: jsii.Number(123),
//   					},
//   					PassThroughSettings: &PassThroughSettingsProperty{
//   					},
//   					WavSettings: &WavSettingsProperty{
//   						BitDepth: jsii.Number(123),
//   						CodingMode: jsii.String("codingMode"),
//   						SampleRate: jsii.Number(123),
//   					},
//   				},
//   				LanguageCode: jsii.String("languageCode"),
//   				LanguageCodeControl: jsii.String("languageCodeControl"),
//   				Name: jsii.String("name"),
//   				RemixSettings: &RemixSettingsProperty{
//   					ChannelMappings: []interface{}{
//   						&AudioChannelMappingProperty{
//   							InputChannelLevels: []interface{}{
//   								&InputChannelLevelProperty{
//   									Gain: jsii.Number(123),
//   									InputChannel: jsii.Number(123),
//   								},
//   							},
//   							OutputChannel: jsii.Number(123),
//   						},
//   					},
//   					ChannelsIn: jsii.Number(123),
//   					ChannelsOut: jsii.Number(123),
//   				},
//   				StreamName: jsii.String("streamName"),
//   			},
//   		},
//   		AvailBlanking: &AvailBlankingProperty{
//   			AvailBlankingImage: &InputLocationProperty{
//   				PasswordParam: jsii.String("passwordParam"),
//   				Uri: jsii.String("uri"),
//   				Username: jsii.String("username"),
//   			},
//   			State: jsii.String("state"),
//   		},
//   		AvailConfiguration: &AvailConfigurationProperty{
//   			AvailSettings: &AvailSettingsProperty{
//   				Scte35SpliceInsert: &Scte35SpliceInsertProperty{
//   					AdAvailOffset: jsii.Number(123),
//   					NoRegionalBlackoutFlag: jsii.String("noRegionalBlackoutFlag"),
//   					WebDeliveryAllowedFlag: jsii.String("webDeliveryAllowedFlag"),
//   				},
//   				Scte35TimeSignalApos: &Scte35TimeSignalAposProperty{
//   					AdAvailOffset: jsii.Number(123),
//   					NoRegionalBlackoutFlag: jsii.String("noRegionalBlackoutFlag"),
//   					WebDeliveryAllowedFlag: jsii.String("webDeliveryAllowedFlag"),
//   				},
//   			},
//   		},
//   		BlackoutSlate: &BlackoutSlateProperty{
//   			BlackoutSlateImage: &InputLocationProperty{
//   				PasswordParam: jsii.String("passwordParam"),
//   				Uri: jsii.String("uri"),
//   				Username: jsii.String("username"),
//   			},
//   			NetworkEndBlackout: jsii.String("networkEndBlackout"),
//   			NetworkEndBlackoutImage: &InputLocationProperty{
//   				PasswordParam: jsii.String("passwordParam"),
//   				Uri: jsii.String("uri"),
//   				Username: jsii.String("username"),
//   			},
//   			NetworkId: jsii.String("networkId"),
//   			State: jsii.String("state"),
//   		},
//   		CaptionDescriptions: []interface{}{
//   			&CaptionDescriptionProperty{
//   				CaptionSelectorName: jsii.String("captionSelectorName"),
//   				DestinationSettings: &CaptionDestinationSettingsProperty{
//   					AribDestinationSettings: &AribDestinationSettingsProperty{
//   					},
//   					BurnInDestinationSettings: &BurnInDestinationSettingsProperty{
//   						Alignment: jsii.String("alignment"),
//   						BackgroundColor: jsii.String("backgroundColor"),
//   						BackgroundOpacity: jsii.Number(123),
//   						Font: &InputLocationProperty{
//   							PasswordParam: jsii.String("passwordParam"),
//   							Uri: jsii.String("uri"),
//   							Username: jsii.String("username"),
//   						},
//   						FontColor: jsii.String("fontColor"),
//   						FontOpacity: jsii.Number(123),
//   						FontResolution: jsii.Number(123),
//   						FontSize: jsii.String("fontSize"),
//   						OutlineColor: jsii.String("outlineColor"),
//   						OutlineSize: jsii.Number(123),
//   						ShadowColor: jsii.String("shadowColor"),
//   						ShadowOpacity: jsii.Number(123),
//   						ShadowXOffset: jsii.Number(123),
//   						ShadowYOffset: jsii.Number(123),
//   						TeletextGridControl: jsii.String("teletextGridControl"),
//   						XPosition: jsii.Number(123),
//   						YPosition: jsii.Number(123),
//   					},
//   					DvbSubDestinationSettings: &DvbSubDestinationSettingsProperty{
//   						Alignment: jsii.String("alignment"),
//   						BackgroundColor: jsii.String("backgroundColor"),
//   						BackgroundOpacity: jsii.Number(123),
//   						Font: &InputLocationProperty{
//   							PasswordParam: jsii.String("passwordParam"),
//   							Uri: jsii.String("uri"),
//   							Username: jsii.String("username"),
//   						},
//   						FontColor: jsii.String("fontColor"),
//   						FontOpacity: jsii.Number(123),
//   						FontResolution: jsii.Number(123),
//   						FontSize: jsii.String("fontSize"),
//   						OutlineColor: jsii.String("outlineColor"),
//   						OutlineSize: jsii.Number(123),
//   						ShadowColor: jsii.String("shadowColor"),
//   						ShadowOpacity: jsii.Number(123),
//   						ShadowXOffset: jsii.Number(123),
//   						ShadowYOffset: jsii.Number(123),
//   						TeletextGridControl: jsii.String("teletextGridControl"),
//   						XPosition: jsii.Number(123),
//   						YPosition: jsii.Number(123),
//   					},
//   					EbuTtDDestinationSettings: &EbuTtDDestinationSettingsProperty{
//   						CopyrightHolder: jsii.String("copyrightHolder"),
//   						FillLineGap: jsii.String("fillLineGap"),
//   						FontFamily: jsii.String("fontFamily"),
//   						StyleControl: jsii.String("styleControl"),
//   					},
//   					EmbeddedDestinationSettings: &EmbeddedDestinationSettingsProperty{
//   					},
//   					EmbeddedPlusScte20DestinationSettings: &EmbeddedPlusScte20DestinationSettingsProperty{
//   					},
//   					RtmpCaptionInfoDestinationSettings: &RtmpCaptionInfoDestinationSettingsProperty{
//   					},
//   					Scte20PlusEmbeddedDestinationSettings: &Scte20PlusEmbeddedDestinationSettingsProperty{
//   					},
//   					Scte27DestinationSettings: &Scte27DestinationSettingsProperty{
//   					},
//   					SmpteTtDestinationSettings: &SmpteTtDestinationSettingsProperty{
//   					},
//   					TeletextDestinationSettings: &TeletextDestinationSettingsProperty{
//   					},
//   					TtmlDestinationSettings: &TtmlDestinationSettingsProperty{
//   						StyleControl: jsii.String("styleControl"),
//   					},
//   					WebvttDestinationSettings: &WebvttDestinationSettingsProperty{
//   						StyleControl: jsii.String("styleControl"),
//   					},
//   				},
//   				LanguageCode: jsii.String("languageCode"),
//   				LanguageDescription: jsii.String("languageDescription"),
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		FeatureActivations: &FeatureActivationsProperty{
//   			InputPrepareScheduleActions: jsii.String("inputPrepareScheduleActions"),
//   		},
//   		GlobalConfiguration: &GlobalConfigurationProperty{
//   			InitialAudioGain: jsii.Number(123),
//   			InputEndAction: jsii.String("inputEndAction"),
//   			InputLossBehavior: &InputLossBehaviorProperty{
//   				BlackFrameMsec: jsii.Number(123),
//   				InputLossImageColor: jsii.String("inputLossImageColor"),
//   				InputLossImageSlate: &InputLocationProperty{
//   					PasswordParam: jsii.String("passwordParam"),
//   					Uri: jsii.String("uri"),
//   					Username: jsii.String("username"),
//   				},
//   				InputLossImageType: jsii.String("inputLossImageType"),
//   				RepeatFrameMsec: jsii.Number(123),
//   			},
//   			OutputLockingMode: jsii.String("outputLockingMode"),
//   			OutputTimingSource: jsii.String("outputTimingSource"),
//   			SupportLowFramerateInputs: jsii.String("supportLowFramerateInputs"),
//   		},
//   		MotionGraphicsConfiguration: &MotionGraphicsConfigurationProperty{
//   			MotionGraphicsInsertion: jsii.String("motionGraphicsInsertion"),
//   			MotionGraphicsSettings: &MotionGraphicsSettingsProperty{
//   				HtmlMotionGraphicsSettings: &HtmlMotionGraphicsSettingsProperty{
//   				},
//   			},
//   		},
//   		NielsenConfiguration: &NielsenConfigurationProperty{
//   			DistributorId: jsii.String("distributorId"),
//   			NielsenPcmToId3Tagging: jsii.String("nielsenPcmToId3Tagging"),
//   		},
//   		OutputGroups: []interface{}{
//   			&OutputGroupProperty{
//   				Name: jsii.String("name"),
//   				OutputGroupSettings: &OutputGroupSettingsProperty{
//   					ArchiveGroupSettings: &ArchiveGroupSettingsProperty{
//   						ArchiveCdnSettings: &ArchiveCdnSettingsProperty{
//   							ArchiveS3Settings: &ArchiveS3SettingsProperty{
//   								CannedAcl: jsii.String("cannedAcl"),
//   							},
//   						},
//   						Destination: &OutputLocationRefProperty{
//   							DestinationRefId: jsii.String("destinationRefId"),
//   						},
//   						RolloverInterval: jsii.Number(123),
//   					},
//   					FrameCaptureGroupSettings: &FrameCaptureGroupSettingsProperty{
//   						Destination: &OutputLocationRefProperty{
//   							DestinationRefId: jsii.String("destinationRefId"),
//   						},
//   						FrameCaptureCdnSettings: &FrameCaptureCdnSettingsProperty{
//   							FrameCaptureS3Settings: &FrameCaptureS3SettingsProperty{
//   								CannedAcl: jsii.String("cannedAcl"),
//   							},
//   						},
//   					},
//   					HlsGroupSettings: &HlsGroupSettingsProperty{
//   						AdMarkers: []*string{
//   							jsii.String("adMarkers"),
//   						},
//   						BaseUrlContent: jsii.String("baseUrlContent"),
//   						BaseUrlContent1: jsii.String("baseUrlContent1"),
//   						BaseUrlManifest: jsii.String("baseUrlManifest"),
//   						BaseUrlManifest1: jsii.String("baseUrlManifest1"),
//   						CaptionLanguageMappings: []interface{}{
//   							&CaptionLanguageMappingProperty{
//   								CaptionChannel: jsii.Number(123),
//   								LanguageCode: jsii.String("languageCode"),
//   								LanguageDescription: jsii.String("languageDescription"),
//   							},
//   						},
//   						CaptionLanguageSetting: jsii.String("captionLanguageSetting"),
//   						ClientCache: jsii.String("clientCache"),
//   						CodecSpecification: jsii.String("codecSpecification"),
//   						ConstantIv: jsii.String("constantIv"),
//   						Destination: &OutputLocationRefProperty{
//   							DestinationRefId: jsii.String("destinationRefId"),
//   						},
//   						DirectoryStructure: jsii.String("directoryStructure"),
//   						DiscontinuityTags: jsii.String("discontinuityTags"),
//   						EncryptionType: jsii.String("encryptionType"),
//   						HlsCdnSettings: &HlsCdnSettingsProperty{
//   							HlsAkamaiSettings: &HlsAkamaiSettingsProperty{
//   								ConnectionRetryInterval: jsii.Number(123),
//   								FilecacheDuration: jsii.Number(123),
//   								HttpTransferMode: jsii.String("httpTransferMode"),
//   								NumRetries: jsii.Number(123),
//   								RestartDelay: jsii.Number(123),
//   								Salt: jsii.String("salt"),
//   								Token: jsii.String("token"),
//   							},
//   							HlsBasicPutSettings: &HlsBasicPutSettingsProperty{
//   								ConnectionRetryInterval: jsii.Number(123),
//   								FilecacheDuration: jsii.Number(123),
//   								NumRetries: jsii.Number(123),
//   								RestartDelay: jsii.Number(123),
//   							},
//   							HlsMediaStoreSettings: &HlsMediaStoreSettingsProperty{
//   								ConnectionRetryInterval: jsii.Number(123),
//   								FilecacheDuration: jsii.Number(123),
//   								MediaStoreStorageClass: jsii.String("mediaStoreStorageClass"),
//   								NumRetries: jsii.Number(123),
//   								RestartDelay: jsii.Number(123),
//   							},
//   							HlsS3Settings: &HlsS3SettingsProperty{
//   								CannedAcl: jsii.String("cannedAcl"),
//   							},
//   							HlsWebdavSettings: &HlsWebdavSettingsProperty{
//   								ConnectionRetryInterval: jsii.Number(123),
//   								FilecacheDuration: jsii.Number(123),
//   								HttpTransferMode: jsii.String("httpTransferMode"),
//   								NumRetries: jsii.Number(123),
//   								RestartDelay: jsii.Number(123),
//   							},
//   						},
//   						HlsId3SegmentTagging: jsii.String("hlsId3SegmentTagging"),
//   						IFrameOnlyPlaylists: jsii.String("iFrameOnlyPlaylists"),
//   						IncompleteSegmentBehavior: jsii.String("incompleteSegmentBehavior"),
//   						IndexNSegments: jsii.Number(123),
//   						InputLossAction: jsii.String("inputLossAction"),
//   						IvInManifest: jsii.String("ivInManifest"),
//   						IvSource: jsii.String("ivSource"),
//   						KeepSegments: jsii.Number(123),
//   						KeyFormat: jsii.String("keyFormat"),
//   						KeyFormatVersions: jsii.String("keyFormatVersions"),
//   						KeyProviderSettings: &KeyProviderSettingsProperty{
//   							StaticKeySettings: &StaticKeySettingsProperty{
//   								KeyProviderServer: &InputLocationProperty{
//   									PasswordParam: jsii.String("passwordParam"),
//   									Uri: jsii.String("uri"),
//   									Username: jsii.String("username"),
//   								},
//   								StaticKeyValue: jsii.String("staticKeyValue"),
//   							},
//   						},
//   						ManifestCompression: jsii.String("manifestCompression"),
//   						ManifestDurationFormat: jsii.String("manifestDurationFormat"),
//   						MinSegmentLength: jsii.Number(123),
//   						Mode: jsii.String("mode"),
//   						OutputSelection: jsii.String("outputSelection"),
//   						ProgramDateTime: jsii.String("programDateTime"),
//   						ProgramDateTimeClock: jsii.String("programDateTimeClock"),
//   						ProgramDateTimePeriod: jsii.Number(123),
//   						RedundantManifest: jsii.String("redundantManifest"),
//   						SegmentationMode: jsii.String("segmentationMode"),
//   						SegmentLength: jsii.Number(123),
//   						SegmentsPerSubdirectory: jsii.Number(123),
//   						StreamInfResolution: jsii.String("streamInfResolution"),
//   						TimedMetadataId3Frame: jsii.String("timedMetadataId3Frame"),
//   						TimedMetadataId3Period: jsii.Number(123),
//   						TimestampDeltaMilliseconds: jsii.Number(123),
//   						TsFileMode: jsii.String("tsFileMode"),
//   					},
//   					MediaPackageGroupSettings: &MediaPackageGroupSettingsProperty{
//   						Destination: &OutputLocationRefProperty{
//   							DestinationRefId: jsii.String("destinationRefId"),
//   						},
//   					},
//   					MsSmoothGroupSettings: &MsSmoothGroupSettingsProperty{
//   						AcquisitionPointId: jsii.String("acquisitionPointId"),
//   						AudioOnlyTimecodeControl: jsii.String("audioOnlyTimecodeControl"),
//   						CertificateMode: jsii.String("certificateMode"),
//   						ConnectionRetryInterval: jsii.Number(123),
//   						Destination: &OutputLocationRefProperty{
//   							DestinationRefId: jsii.String("destinationRefId"),
//   						},
//   						EventId: jsii.String("eventId"),
//   						EventIdMode: jsii.String("eventIdMode"),
//   						EventStopBehavior: jsii.String("eventStopBehavior"),
//   						FilecacheDuration: jsii.Number(123),
//   						FragmentLength: jsii.Number(123),
//   						InputLossAction: jsii.String("inputLossAction"),
//   						NumRetries: jsii.Number(123),
//   						RestartDelay: jsii.Number(123),
//   						SegmentationMode: jsii.String("segmentationMode"),
//   						SendDelayMs: jsii.Number(123),
//   						SparseTrackType: jsii.String("sparseTrackType"),
//   						StreamManifestBehavior: jsii.String("streamManifestBehavior"),
//   						TimestampOffset: jsii.String("timestampOffset"),
//   						TimestampOffsetMode: jsii.String("timestampOffsetMode"),
//   					},
//   					MultiplexGroupSettings: &MultiplexGroupSettingsProperty{
//   					},
//   					RtmpGroupSettings: &RtmpGroupSettingsProperty{
//   						AdMarkers: []*string{
//   							jsii.String("adMarkers"),
//   						},
//   						AuthenticationScheme: jsii.String("authenticationScheme"),
//   						CacheFullBehavior: jsii.String("cacheFullBehavior"),
//   						CacheLength: jsii.Number(123),
//   						CaptionData: jsii.String("captionData"),
//   						InputLossAction: jsii.String("inputLossAction"),
//   						RestartDelay: jsii.Number(123),
//   					},
//   					UdpGroupSettings: &UdpGroupSettingsProperty{
//   						InputLossAction: jsii.String("inputLossAction"),
//   						TimedMetadataId3Frame: jsii.String("timedMetadataId3Frame"),
//   						TimedMetadataId3Period: jsii.Number(123),
//   					},
//   				},
//   				Outputs: []interface{}{
//   					&OutputProperty{
//   						AudioDescriptionNames: []*string{
//   							jsii.String("audioDescriptionNames"),
//   						},
//   						CaptionDescriptionNames: []*string{
//   							jsii.String("captionDescriptionNames"),
//   						},
//   						OutputName: jsii.String("outputName"),
//   						OutputSettings: &OutputSettingsProperty{
//   							ArchiveOutputSettings: &ArchiveOutputSettingsProperty{
//   								ContainerSettings: &ArchiveContainerSettingsProperty{
//   									M2TsSettings: &M2tsSettingsProperty{
//   										AbsentInputAudioBehavior: jsii.String("absentInputAudioBehavior"),
//   										Arib: jsii.String("arib"),
//   										AribCaptionsPid: jsii.String("aribCaptionsPid"),
//   										AribCaptionsPidControl: jsii.String("aribCaptionsPidControl"),
//   										AudioBufferModel: jsii.String("audioBufferModel"),
//   										AudioFramesPerPes: jsii.Number(123),
//   										AudioPids: jsii.String("audioPids"),
//   										AudioStreamType: jsii.String("audioStreamType"),
//   										Bitrate: jsii.Number(123),
//   										BufferModel: jsii.String("bufferModel"),
//   										CcDescriptor: jsii.String("ccDescriptor"),
//   										DvbNitSettings: &DvbNitSettingsProperty{
//   											NetworkId: jsii.Number(123),
//   											NetworkName: jsii.String("networkName"),
//   											RepInterval: jsii.Number(123),
//   										},
//   										DvbSdtSettings: &DvbSdtSettingsProperty{
//   											OutputSdt: jsii.String("outputSdt"),
//   											RepInterval: jsii.Number(123),
//   											ServiceName: jsii.String("serviceName"),
//   											ServiceProviderName: jsii.String("serviceProviderName"),
//   										},
//   										DvbSubPids: jsii.String("dvbSubPids"),
//   										DvbTdtSettings: &DvbTdtSettingsProperty{
//   											RepInterval: jsii.Number(123),
//   										},
//   										DvbTeletextPid: jsii.String("dvbTeletextPid"),
//   										Ebif: jsii.String("ebif"),
//   										EbpAudioInterval: jsii.String("ebpAudioInterval"),
//   										EbpLookaheadMs: jsii.Number(123),
//   										EbpPlacement: jsii.String("ebpPlacement"),
//   										EcmPid: jsii.String("ecmPid"),
//   										EsRateInPes: jsii.String("esRateInPes"),
//   										EtvPlatformPid: jsii.String("etvPlatformPid"),
//   										EtvSignalPid: jsii.String("etvSignalPid"),
//   										FragmentTime: jsii.Number(123),
//   										Klv: jsii.String("klv"),
//   										KlvDataPids: jsii.String("klvDataPids"),
//   										NielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   										NullPacketBitrate: jsii.Number(123),
//   										PatInterval: jsii.Number(123),
//   										PcrControl: jsii.String("pcrControl"),
//   										PcrPeriod: jsii.Number(123),
//   										PcrPid: jsii.String("pcrPid"),
//   										PmtInterval: jsii.Number(123),
//   										PmtPid: jsii.String("pmtPid"),
//   										ProgramNum: jsii.Number(123),
//   										RateMode: jsii.String("rateMode"),
//   										Scte27Pids: jsii.String("scte27Pids"),
//   										Scte35Control: jsii.String("scte35Control"),
//   										Scte35Pid: jsii.String("scte35Pid"),
//   										SegmentationMarkers: jsii.String("segmentationMarkers"),
//   										SegmentationStyle: jsii.String("segmentationStyle"),
//   										SegmentationTime: jsii.Number(123),
//   										TimedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   										TimedMetadataPid: jsii.String("timedMetadataPid"),
//   										TransportStreamId: jsii.Number(123),
//   										VideoPid: jsii.String("videoPid"),
//   									},
//   									RawSettings: &RawSettingsProperty{
//   									},
//   								},
//   								Extension: jsii.String("extension"),
//   								NameModifier: jsii.String("nameModifier"),
//   							},
//   							FrameCaptureOutputSettings: &FrameCaptureOutputSettingsProperty{
//   								NameModifier: jsii.String("nameModifier"),
//   							},
//   							HlsOutputSettings: &HlsOutputSettingsProperty{
//   								H265PackagingType: jsii.String("h265PackagingType"),
//   								HlsSettings: &HlsSettingsProperty{
//   									AudioOnlyHlsSettings: &AudioOnlyHlsSettingsProperty{
//   										AudioGroupId: jsii.String("audioGroupId"),
//   										AudioOnlyImage: &InputLocationProperty{
//   											PasswordParam: jsii.String("passwordParam"),
//   											Uri: jsii.String("uri"),
//   											Username: jsii.String("username"),
//   										},
//   										AudioTrackType: jsii.String("audioTrackType"),
//   										SegmentType: jsii.String("segmentType"),
//   									},
//   									Fmp4HlsSettings: &Fmp4HlsSettingsProperty{
//   										AudioRenditionSets: jsii.String("audioRenditionSets"),
//   										NielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   										TimedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   									},
//   									FrameCaptureHlsSettings: &FrameCaptureHlsSettingsProperty{
//   									},
//   									StandardHlsSettings: &StandardHlsSettingsProperty{
//   										AudioRenditionSets: jsii.String("audioRenditionSets"),
//   										M3U8Settings: &M3u8SettingsProperty{
//   											AudioFramesPerPes: jsii.Number(123),
//   											AudioPids: jsii.String("audioPids"),
//   											EcmPid: jsii.String("ecmPid"),
//   											NielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   											PatInterval: jsii.Number(123),
//   											PcrControl: jsii.String("pcrControl"),
//   											PcrPeriod: jsii.Number(123),
//   											PcrPid: jsii.String("pcrPid"),
//   											PmtInterval: jsii.Number(123),
//   											PmtPid: jsii.String("pmtPid"),
//   											ProgramNum: jsii.Number(123),
//   											Scte35Behavior: jsii.String("scte35Behavior"),
//   											Scte35Pid: jsii.String("scte35Pid"),
//   											TimedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   											TimedMetadataPid: jsii.String("timedMetadataPid"),
//   											TransportStreamId: jsii.Number(123),
//   											VideoPid: jsii.String("videoPid"),
//   										},
//   									},
//   								},
//   								NameModifier: jsii.String("nameModifier"),
//   								SegmentModifier: jsii.String("segmentModifier"),
//   							},
//   							MediaPackageOutputSettings: &MediaPackageOutputSettingsProperty{
//   							},
//   							MsSmoothOutputSettings: &MsSmoothOutputSettingsProperty{
//   								H265PackagingType: jsii.String("h265PackagingType"),
//   								NameModifier: jsii.String("nameModifier"),
//   							},
//   							MultiplexOutputSettings: &MultiplexOutputSettingsProperty{
//   								Destination: &OutputLocationRefProperty{
//   									DestinationRefId: jsii.String("destinationRefId"),
//   								},
//   							},
//   							RtmpOutputSettings: &RtmpOutputSettingsProperty{
//   								CertificateMode: jsii.String("certificateMode"),
//   								ConnectionRetryInterval: jsii.Number(123),
//   								Destination: &OutputLocationRefProperty{
//   									DestinationRefId: jsii.String("destinationRefId"),
//   								},
//   								NumRetries: jsii.Number(123),
//   							},
//   							UdpOutputSettings: &UdpOutputSettingsProperty{
//   								BufferMsec: jsii.Number(123),
//   								ContainerSettings: &UdpContainerSettingsProperty{
//   									M2TsSettings: &M2tsSettingsProperty{
//   										AbsentInputAudioBehavior: jsii.String("absentInputAudioBehavior"),
//   										Arib: jsii.String("arib"),
//   										AribCaptionsPid: jsii.String("aribCaptionsPid"),
//   										AribCaptionsPidControl: jsii.String("aribCaptionsPidControl"),
//   										AudioBufferModel: jsii.String("audioBufferModel"),
//   										AudioFramesPerPes: jsii.Number(123),
//   										AudioPids: jsii.String("audioPids"),
//   										AudioStreamType: jsii.String("audioStreamType"),
//   										Bitrate: jsii.Number(123),
//   										BufferModel: jsii.String("bufferModel"),
//   										CcDescriptor: jsii.String("ccDescriptor"),
//   										DvbNitSettings: &DvbNitSettingsProperty{
//   											NetworkId: jsii.Number(123),
//   											NetworkName: jsii.String("networkName"),
//   											RepInterval: jsii.Number(123),
//   										},
//   										DvbSdtSettings: &DvbSdtSettingsProperty{
//   											OutputSdt: jsii.String("outputSdt"),
//   											RepInterval: jsii.Number(123),
//   											ServiceName: jsii.String("serviceName"),
//   											ServiceProviderName: jsii.String("serviceProviderName"),
//   										},
//   										DvbSubPids: jsii.String("dvbSubPids"),
//   										DvbTdtSettings: &DvbTdtSettingsProperty{
//   											RepInterval: jsii.Number(123),
//   										},
//   										DvbTeletextPid: jsii.String("dvbTeletextPid"),
//   										Ebif: jsii.String("ebif"),
//   										EbpAudioInterval: jsii.String("ebpAudioInterval"),
//   										EbpLookaheadMs: jsii.Number(123),
//   										EbpPlacement: jsii.String("ebpPlacement"),
//   										EcmPid: jsii.String("ecmPid"),
//   										EsRateInPes: jsii.String("esRateInPes"),
//   										EtvPlatformPid: jsii.String("etvPlatformPid"),
//   										EtvSignalPid: jsii.String("etvSignalPid"),
//   										FragmentTime: jsii.Number(123),
//   										Klv: jsii.String("klv"),
//   										KlvDataPids: jsii.String("klvDataPids"),
//   										NielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   										NullPacketBitrate: jsii.Number(123),
//   										PatInterval: jsii.Number(123),
//   										PcrControl: jsii.String("pcrControl"),
//   										PcrPeriod: jsii.Number(123),
//   										PcrPid: jsii.String("pcrPid"),
//   										PmtInterval: jsii.Number(123),
//   										PmtPid: jsii.String("pmtPid"),
//   										ProgramNum: jsii.Number(123),
//   										RateMode: jsii.String("rateMode"),
//   										Scte27Pids: jsii.String("scte27Pids"),
//   										Scte35Control: jsii.String("scte35Control"),
//   										Scte35Pid: jsii.String("scte35Pid"),
//   										SegmentationMarkers: jsii.String("segmentationMarkers"),
//   										SegmentationStyle: jsii.String("segmentationStyle"),
//   										SegmentationTime: jsii.Number(123),
//   										TimedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   										TimedMetadataPid: jsii.String("timedMetadataPid"),
//   										TransportStreamId: jsii.Number(123),
//   										VideoPid: jsii.String("videoPid"),
//   									},
//   								},
//   								Destination: &OutputLocationRefProperty{
//   									DestinationRefId: jsii.String("destinationRefId"),
//   								},
//   								FecOutputSettings: &FecOutputSettingsProperty{
//   									ColumnDepth: jsii.Number(123),
//   									IncludeFec: jsii.String("includeFec"),
//   									RowLength: jsii.Number(123),
//   								},
//   							},
//   						},
//   						VideoDescriptionName: jsii.String("videoDescriptionName"),
//   					},
//   				},
//   			},
//   		},
//   		TimecodeConfig: &TimecodeConfigProperty{
//   			Source: jsii.String("source"),
//   			SyncThreshold: jsii.Number(123),
//   		},
//   		VideoDescriptions: []interface{}{
//   			&VideoDescriptionProperty{
//   				CodecSettings: &VideoCodecSettingsProperty{
//   					FrameCaptureSettings: &FrameCaptureSettingsProperty{
//   						CaptureInterval: jsii.Number(123),
//   						CaptureIntervalUnits: jsii.String("captureIntervalUnits"),
//   					},
//   					H264Settings: &H264SettingsProperty{
//   						AdaptiveQuantization: jsii.String("adaptiveQuantization"),
//   						AfdSignaling: jsii.String("afdSignaling"),
//   						Bitrate: jsii.Number(123),
//   						BufFillPct: jsii.Number(123),
//   						BufSize: jsii.Number(123),
//   						ColorMetadata: jsii.String("colorMetadata"),
//   						ColorSpaceSettings: &H264ColorSpaceSettingsProperty{
//   							ColorSpacePassthroughSettings: &ColorSpacePassthroughSettingsProperty{
//   							},
//   							Rec601Settings: &Rec601SettingsProperty{
//   							},
//   							Rec709Settings: &Rec709SettingsProperty{
//   							},
//   						},
//   						EntropyEncoding: jsii.String("entropyEncoding"),
//   						FilterSettings: &H264FilterSettingsProperty{
//   							TemporalFilterSettings: &TemporalFilterSettingsProperty{
//   								PostFilterSharpening: jsii.String("postFilterSharpening"),
//   								Strength: jsii.String("strength"),
//   							},
//   						},
//   						FixedAfd: jsii.String("fixedAfd"),
//   						FlickerAq: jsii.String("flickerAq"),
//   						ForceFieldPictures: jsii.String("forceFieldPictures"),
//   						FramerateControl: jsii.String("framerateControl"),
//   						FramerateDenominator: jsii.Number(123),
//   						FramerateNumerator: jsii.Number(123),
//   						GopBReference: jsii.String("gopBReference"),
//   						GopClosedCadence: jsii.Number(123),
//   						GopNumBFrames: jsii.Number(123),
//   						GopSize: jsii.Number(123),
//   						GopSizeUnits: jsii.String("gopSizeUnits"),
//   						Level: jsii.String("level"),
//   						LookAheadRateControl: jsii.String("lookAheadRateControl"),
//   						MaxBitrate: jsii.Number(123),
//   						MinIInterval: jsii.Number(123),
//   						NumRefFrames: jsii.Number(123),
//   						ParControl: jsii.String("parControl"),
//   						ParDenominator: jsii.Number(123),
//   						ParNumerator: jsii.Number(123),
//   						Profile: jsii.String("profile"),
//   						QualityLevel: jsii.String("qualityLevel"),
//   						QvbrQualityLevel: jsii.Number(123),
//   						RateControlMode: jsii.String("rateControlMode"),
//   						ScanType: jsii.String("scanType"),
//   						SceneChangeDetect: jsii.String("sceneChangeDetect"),
//   						Slices: jsii.Number(123),
//   						Softness: jsii.Number(123),
//   						SpatialAq: jsii.String("spatialAq"),
//   						SubgopLength: jsii.String("subgopLength"),
//   						Syntax: jsii.String("syntax"),
//   						TemporalAq: jsii.String("temporalAq"),
//   						TimecodeInsertion: jsii.String("timecodeInsertion"),
//   					},
//   					H265Settings: &H265SettingsProperty{
//   						AdaptiveQuantization: jsii.String("adaptiveQuantization"),
//   						AfdSignaling: jsii.String("afdSignaling"),
//   						AlternativeTransferFunction: jsii.String("alternativeTransferFunction"),
//   						Bitrate: jsii.Number(123),
//   						BufSize: jsii.Number(123),
//   						ColorMetadata: jsii.String("colorMetadata"),
//   						ColorSpaceSettings: &H265ColorSpaceSettingsProperty{
//   							ColorSpacePassthroughSettings: &ColorSpacePassthroughSettingsProperty{
//   							},
//   							Hdr10Settings: &Hdr10SettingsProperty{
//   								MaxCll: jsii.Number(123),
//   								MaxFall: jsii.Number(123),
//   							},
//   							Rec601Settings: &Rec601SettingsProperty{
//   							},
//   							Rec709Settings: &Rec709SettingsProperty{
//   							},
//   						},
//   						FilterSettings: &H265FilterSettingsProperty{
//   							TemporalFilterSettings: &TemporalFilterSettingsProperty{
//   								PostFilterSharpening: jsii.String("postFilterSharpening"),
//   								Strength: jsii.String("strength"),
//   							},
//   						},
//   						FixedAfd: jsii.String("fixedAfd"),
//   						FlickerAq: jsii.String("flickerAq"),
//   						FramerateDenominator: jsii.Number(123),
//   						FramerateNumerator: jsii.Number(123),
//   						GopClosedCadence: jsii.Number(123),
//   						GopSize: jsii.Number(123),
//   						GopSizeUnits: jsii.String("gopSizeUnits"),
//   						Level: jsii.String("level"),
//   						LookAheadRateControl: jsii.String("lookAheadRateControl"),
//   						MaxBitrate: jsii.Number(123),
//   						MinIInterval: jsii.Number(123),
//   						ParDenominator: jsii.Number(123),
//   						ParNumerator: jsii.Number(123),
//   						Profile: jsii.String("profile"),
//   						QvbrQualityLevel: jsii.Number(123),
//   						RateControlMode: jsii.String("rateControlMode"),
//   						ScanType: jsii.String("scanType"),
//   						SceneChangeDetect: jsii.String("sceneChangeDetect"),
//   						Slices: jsii.Number(123),
//   						Tier: jsii.String("tier"),
//   						TimecodeInsertion: jsii.String("timecodeInsertion"),
//   					},
//   					Mpeg2Settings: &Mpeg2SettingsProperty{
//   						AdaptiveQuantization: jsii.String("adaptiveQuantization"),
//   						AfdSignaling: jsii.String("afdSignaling"),
//   						ColorMetadata: jsii.String("colorMetadata"),
//   						ColorSpace: jsii.String("colorSpace"),
//   						DisplayAspectRatio: jsii.String("displayAspectRatio"),
//   						FilterSettings: &Mpeg2FilterSettingsProperty{
//   							TemporalFilterSettings: &TemporalFilterSettingsProperty{
//   								PostFilterSharpening: jsii.String("postFilterSharpening"),
//   								Strength: jsii.String("strength"),
//   							},
//   						},
//   						FixedAfd: jsii.String("fixedAfd"),
//   						FramerateDenominator: jsii.Number(123),
//   						FramerateNumerator: jsii.Number(123),
//   						GopClosedCadence: jsii.Number(123),
//   						GopNumBFrames: jsii.Number(123),
//   						GopSize: jsii.Number(123),
//   						GopSizeUnits: jsii.String("gopSizeUnits"),
//   						ScanType: jsii.String("scanType"),
//   						SubgopLength: jsii.String("subgopLength"),
//   						TimecodeInsertion: jsii.String("timecodeInsertion"),
//   					},
//   				},
//   				Height: jsii.Number(123),
//   				Name: jsii.String("name"),
//   				RespondToAfd: jsii.String("respondToAfd"),
//   				ScalingBehavior: jsii.String("scalingBehavior"),
//   				Sharpness: jsii.Number(123),
//   				Width: jsii.Number(123),
//   			},
//   		},
//   	},
//   	InputAttachments: []interface{}{
//   		&InputAttachmentProperty{
//   			AutomaticInputFailoverSettings: &AutomaticInputFailoverSettingsProperty{
//   				ErrorClearTimeMsec: jsii.Number(123),
//   				FailoverConditions: []interface{}{
//   					&FailoverConditionProperty{
//   						FailoverConditionSettings: &FailoverConditionSettingsProperty{
//   							AudioSilenceSettings: &AudioSilenceFailoverSettingsProperty{
//   								AudioSelectorName: jsii.String("audioSelectorName"),
//   								AudioSilenceThresholdMsec: jsii.Number(123),
//   							},
//   							InputLossSettings: &InputLossFailoverSettingsProperty{
//   								InputLossThresholdMsec: jsii.Number(123),
//   							},
//   							VideoBlackSettings: &VideoBlackFailoverSettingsProperty{
//   								BlackDetectThreshold: jsii.Number(123),
//   								VideoBlackThresholdMsec: jsii.Number(123),
//   							},
//   						},
//   					},
//   				},
//   				InputPreference: jsii.String("inputPreference"),
//   				SecondaryInputId: jsii.String("secondaryInputId"),
//   			},
//   			InputAttachmentName: jsii.String("inputAttachmentName"),
//   			InputId: jsii.String("inputId"),
//   			InputSettings: &InputSettingsProperty{
//   				AudioSelectors: []interface{}{
//   					&AudioSelectorProperty{
//   						Name: jsii.String("name"),
//   						SelectorSettings: &AudioSelectorSettingsProperty{
//   							AudioHlsRenditionSelection: &AudioHlsRenditionSelectionProperty{
//   								GroupId: jsii.String("groupId"),
//   								Name: jsii.String("name"),
//   							},
//   							AudioLanguageSelection: &AudioLanguageSelectionProperty{
//   								LanguageCode: jsii.String("languageCode"),
//   								LanguageSelectionPolicy: jsii.String("languageSelectionPolicy"),
//   							},
//   							AudioPidSelection: &AudioPidSelectionProperty{
//   								Pid: jsii.Number(123),
//   							},
//   							AudioTrackSelection: &AudioTrackSelectionProperty{
//   								Tracks: []interface{}{
//   									&AudioTrackProperty{
//   										Track: jsii.Number(123),
//   									},
//   								},
//   							},
//   						},
//   					},
//   				},
//   				CaptionSelectors: []interface{}{
//   					&CaptionSelectorProperty{
//   						LanguageCode: jsii.String("languageCode"),
//   						Name: jsii.String("name"),
//   						SelectorSettings: &CaptionSelectorSettingsProperty{
//   							AncillarySourceSettings: &AncillarySourceSettingsProperty{
//   								SourceAncillaryChannelNumber: jsii.Number(123),
//   							},
//   							AribSourceSettings: &AribSourceSettingsProperty{
//   							},
//   							DvbSubSourceSettings: &DvbSubSourceSettingsProperty{
//   								OcrLanguage: jsii.String("ocrLanguage"),
//   								Pid: jsii.Number(123),
//   							},
//   							EmbeddedSourceSettings: &EmbeddedSourceSettingsProperty{
//   								Convert608To708: jsii.String("convert608To708"),
//   								Scte20Detection: jsii.String("scte20Detection"),
//   								Source608ChannelNumber: jsii.Number(123),
//   								Source608TrackNumber: jsii.Number(123),
//   							},
//   							Scte20SourceSettings: &Scte20SourceSettingsProperty{
//   								Convert608To708: jsii.String("convert608To708"),
//   								Source608ChannelNumber: jsii.Number(123),
//   							},
//   							Scte27SourceSettings: &Scte27SourceSettingsProperty{
//   								OcrLanguage: jsii.String("ocrLanguage"),
//   								Pid: jsii.Number(123),
//   							},
//   							TeletextSourceSettings: &TeletextSourceSettingsProperty{
//   								OutputRectangle: &CaptionRectangleProperty{
//   									Height: jsii.Number(123),
//   									LeftOffset: jsii.Number(123),
//   									TopOffset: jsii.Number(123),
//   									Width: jsii.Number(123),
//   								},
//   								PageNumber: jsii.String("pageNumber"),
//   							},
//   						},
//   					},
//   				},
//   				DeblockFilter: jsii.String("deblockFilter"),
//   				DenoiseFilter: jsii.String("denoiseFilter"),
//   				FilterStrength: jsii.Number(123),
//   				InputFilter: jsii.String("inputFilter"),
//   				NetworkInputSettings: &NetworkInputSettingsProperty{
//   					HlsInputSettings: &HlsInputSettingsProperty{
//   						Bandwidth: jsii.Number(123),
//   						BufferSegments: jsii.Number(123),
//   						Retries: jsii.Number(123),
//   						RetryInterval: jsii.Number(123),
//   						Scte35Source: jsii.String("scte35Source"),
//   					},
//   					ServerValidation: jsii.String("serverValidation"),
//   				},
//   				Scte35Pid: jsii.Number(123),
//   				Smpte2038DataPreference: jsii.String("smpte2038DataPreference"),
//   				SourceEndBehavior: jsii.String("sourceEndBehavior"),
//   				VideoSelector: &VideoSelectorProperty{
//   					ColorSpace: jsii.String("colorSpace"),
//   					ColorSpaceSettings: &VideoSelectorColorSpaceSettingsProperty{
//   						Hdr10Settings: &Hdr10SettingsProperty{
//   							MaxCll: jsii.Number(123),
//   							MaxFall: jsii.Number(123),
//   						},
//   					},
//   					ColorSpaceUsage: jsii.String("colorSpaceUsage"),
//   					SelectorSettings: &VideoSelectorSettingsProperty{
//   						VideoSelectorPid: &VideoSelectorPidProperty{
//   							Pid: jsii.Number(123),
//   						},
//   						VideoSelectorProgramId: &VideoSelectorProgramIdProperty{
//   							ProgramId: jsii.Number(123),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	InputSpecification: &InputSpecificationProperty{
//   		Codec: jsii.String("codec"),
//   		MaximumBitrate: jsii.String("maximumBitrate"),
//   		Resolution: jsii.String("resolution"),
//   	},
//   	LogLevel: jsii.String("logLevel"),
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: tags,
//   	Vpc: &VpcOutputSettingsProperty{
//   		PublicAddressAllocationIds: []*string{
//   			jsii.String("publicAddressAllocationIds"),
//   		},
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
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

