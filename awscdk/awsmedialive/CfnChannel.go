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
//   cfnChannel := awscdk.Aws_medialive.NewCfnChannel(this, jsii.String("MyCfnChannel"), &CfnChannelProps{
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
//   							Timezone: jsii.String("timezone"),
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
//   					Eac3AtmosSettings: &Eac3AtmosSettingsProperty{
//   						Bitrate: jsii.Number(123),
//   						CodingMode: jsii.String("codingMode"),
//   						Dialnorm: jsii.Number(123),
//   						DrcLine: jsii.String("drcLine"),
//   						DrcRf: jsii.String("drcRf"),
//   						HeightTrim: jsii.Number(123),
//   						SurroundTrim: jsii.Number(123),
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
//   				Esam: &EsamProperty{
//   					AcquisitionPointId: jsii.String("acquisitionPointId"),
//   					AdAvailOffset: jsii.Number(123),
//   					PasswordParam: jsii.String("passwordParam"),
//   					PoisEndpoint: jsii.String("poisEndpoint"),
//   					Username: jsii.String("username"),
//   					ZoneIdentity: jsii.String("zoneIdentity"),
//   				},
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
//   				Accessibility: jsii.String("accessibility"),
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
//   										Scte35PrerollPullupMilliseconds: jsii.Number(123),
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
//   										Scte35PrerollPullupMilliseconds: jsii.Number(123),
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
//   						TimecodeBurninSettings: &TimecodeBurninSettingsProperty{
//   							FontSize: jsii.String("fontSize"),
//   							Position: jsii.String("position"),
//   							Prefix: jsii.String("prefix"),
//   						},
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
//   						TimecodeBurninSettings: &TimecodeBurninSettingsProperty{
//   							FontSize: jsii.String("fontSize"),
//   							Position: jsii.String("position"),
//   							Prefix: jsii.String("prefix"),
//   						},
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
//   							DolbyVision81Settings: &DolbyVision81SettingsProperty{
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
//   						TimecodeBurninSettings: &TimecodeBurninSettingsProperty{
//   							FontSize: jsii.String("fontSize"),
//   							Position: jsii.String("position"),
//   							Prefix: jsii.String("prefix"),
//   						},
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
//   						TimecodeBurninSettings: &TimecodeBurninSettingsProperty{
//   							FontSize: jsii.String("fontSize"),
//   							Position: jsii.String("position"),
//   							Prefix: jsii.String("prefix"),
//   						},
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
//   								DolbyEDecode: &AudioDolbyEDecodeProperty{
//   									ProgramSelection: jsii.String("programSelection"),
//   								},
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
//   	Maintenance: &MaintenanceCreateSettingsProperty{
//   		MaintenanceDay: jsii.String("maintenanceDay"),
//   		MaintenanceStartTime: jsii.String("maintenanceStartTime"),
//   	},
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
	// `AWS::MediaLive::Channel.Maintenance`.
	Maintenance() interface{}
	SetMaintenance(val interface{})
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
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
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
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
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
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
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
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
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

func (j *jsiiProxy_CfnChannel) Maintenance() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maintenance",
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

func (j *jsiiProxy_CfnChannel)SetMaintenance(val interface{}) {
	if err := j.validateSetMaintenanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenance",
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

func (c *jsiiProxy_CfnChannel) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
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

func (c *jsiiProxy_CfnChannel) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
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

func (c *jsiiProxy_CfnChannel) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnChannel) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
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

func (c *jsiiProxy_CfnChannel) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
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

func (c *jsiiProxy_CfnChannel) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
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

