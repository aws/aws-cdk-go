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
//   outputGroupProperty := &OutputGroupProperty{
//   	Name: jsii.String("name"),
//   	OutputGroupSettings: &OutputGroupSettingsProperty{
//   		ArchiveGroupSettings: &ArchiveGroupSettingsProperty{
//   			ArchiveCdnSettings: &ArchiveCdnSettingsProperty{
//   				ArchiveS3Settings: &ArchiveS3SettingsProperty{
//   					CannedAcl: jsii.String("cannedAcl"),
//   				},
//   			},
//   			Destination: &OutputLocationRefProperty{
//   				DestinationRefId: jsii.String("destinationRefId"),
//   			},
//   			RolloverInterval: jsii.Number(123),
//   		},
//   		FrameCaptureGroupSettings: &FrameCaptureGroupSettingsProperty{
//   			Destination: &OutputLocationRefProperty{
//   				DestinationRefId: jsii.String("destinationRefId"),
//   			},
//   			FrameCaptureCdnSettings: &FrameCaptureCdnSettingsProperty{
//   				FrameCaptureS3Settings: &FrameCaptureS3SettingsProperty{
//   					CannedAcl: jsii.String("cannedAcl"),
//   				},
//   			},
//   		},
//   		HlsGroupSettings: &HlsGroupSettingsProperty{
//   			AdMarkers: []*string{
//   				jsii.String("adMarkers"),
//   			},
//   			BaseUrlContent: jsii.String("baseUrlContent"),
//   			BaseUrlContent1: jsii.String("baseUrlContent1"),
//   			BaseUrlManifest: jsii.String("baseUrlManifest"),
//   			BaseUrlManifest1: jsii.String("baseUrlManifest1"),
//   			CaptionLanguageMappings: []interface{}{
//   				&CaptionLanguageMappingProperty{
//   					CaptionChannel: jsii.Number(123),
//   					LanguageCode: jsii.String("languageCode"),
//   					LanguageDescription: jsii.String("languageDescription"),
//   				},
//   			},
//   			CaptionLanguageSetting: jsii.String("captionLanguageSetting"),
//   			ClientCache: jsii.String("clientCache"),
//   			CodecSpecification: jsii.String("codecSpecification"),
//   			ConstantIv: jsii.String("constantIv"),
//   			Destination: &OutputLocationRefProperty{
//   				DestinationRefId: jsii.String("destinationRefId"),
//   			},
//   			DirectoryStructure: jsii.String("directoryStructure"),
//   			DiscontinuityTags: jsii.String("discontinuityTags"),
//   			EncryptionType: jsii.String("encryptionType"),
//   			HlsCdnSettings: &HlsCdnSettingsProperty{
//   				HlsAkamaiSettings: &HlsAkamaiSettingsProperty{
//   					ConnectionRetryInterval: jsii.Number(123),
//   					FilecacheDuration: jsii.Number(123),
//   					HttpTransferMode: jsii.String("httpTransferMode"),
//   					NumRetries: jsii.Number(123),
//   					RestartDelay: jsii.Number(123),
//   					Salt: jsii.String("salt"),
//   					Token: jsii.String("token"),
//   				},
//   				HlsBasicPutSettings: &HlsBasicPutSettingsProperty{
//   					ConnectionRetryInterval: jsii.Number(123),
//   					FilecacheDuration: jsii.Number(123),
//   					NumRetries: jsii.Number(123),
//   					RestartDelay: jsii.Number(123),
//   				},
//   				HlsMediaStoreSettings: &HlsMediaStoreSettingsProperty{
//   					ConnectionRetryInterval: jsii.Number(123),
//   					FilecacheDuration: jsii.Number(123),
//   					MediaStoreStorageClass: jsii.String("mediaStoreStorageClass"),
//   					NumRetries: jsii.Number(123),
//   					RestartDelay: jsii.Number(123),
//   				},
//   				HlsS3Settings: &HlsS3SettingsProperty{
//   					CannedAcl: jsii.String("cannedAcl"),
//   				},
//   				HlsWebdavSettings: &HlsWebdavSettingsProperty{
//   					ConnectionRetryInterval: jsii.Number(123),
//   					FilecacheDuration: jsii.Number(123),
//   					HttpTransferMode: jsii.String("httpTransferMode"),
//   					NumRetries: jsii.Number(123),
//   					RestartDelay: jsii.Number(123),
//   				},
//   			},
//   			HlsId3SegmentTagging: jsii.String("hlsId3SegmentTagging"),
//   			IFrameOnlyPlaylists: jsii.String("iFrameOnlyPlaylists"),
//   			IncompleteSegmentBehavior: jsii.String("incompleteSegmentBehavior"),
//   			IndexNSegments: jsii.Number(123),
//   			InputLossAction: jsii.String("inputLossAction"),
//   			IvInManifest: jsii.String("ivInManifest"),
//   			IvSource: jsii.String("ivSource"),
//   			KeepSegments: jsii.Number(123),
//   			KeyFormat: jsii.String("keyFormat"),
//   			KeyFormatVersions: jsii.String("keyFormatVersions"),
//   			KeyProviderSettings: &KeyProviderSettingsProperty{
//   				StaticKeySettings: &StaticKeySettingsProperty{
//   					KeyProviderServer: &InputLocationProperty{
//   						PasswordParam: jsii.String("passwordParam"),
//   						Uri: jsii.String("uri"),
//   						Username: jsii.String("username"),
//   					},
//   					StaticKeyValue: jsii.String("staticKeyValue"),
//   				},
//   			},
//   			ManifestCompression: jsii.String("manifestCompression"),
//   			ManifestDurationFormat: jsii.String("manifestDurationFormat"),
//   			MinSegmentLength: jsii.Number(123),
//   			Mode: jsii.String("mode"),
//   			OutputSelection: jsii.String("outputSelection"),
//   			ProgramDateTime: jsii.String("programDateTime"),
//   			ProgramDateTimeClock: jsii.String("programDateTimeClock"),
//   			ProgramDateTimePeriod: jsii.Number(123),
//   			RedundantManifest: jsii.String("redundantManifest"),
//   			SegmentationMode: jsii.String("segmentationMode"),
//   			SegmentLength: jsii.Number(123),
//   			SegmentsPerSubdirectory: jsii.Number(123),
//   			StreamInfResolution: jsii.String("streamInfResolution"),
//   			TimedMetadataId3Frame: jsii.String("timedMetadataId3Frame"),
//   			TimedMetadataId3Period: jsii.Number(123),
//   			TimestampDeltaMilliseconds: jsii.Number(123),
//   			TsFileMode: jsii.String("tsFileMode"),
//   		},
//   		MediaPackageGroupSettings: &MediaPackageGroupSettingsProperty{
//   			Destination: &OutputLocationRefProperty{
//   				DestinationRefId: jsii.String("destinationRefId"),
//   			},
//   		},
//   		MsSmoothGroupSettings: &MsSmoothGroupSettingsProperty{
//   			AcquisitionPointId: jsii.String("acquisitionPointId"),
//   			AudioOnlyTimecodeControl: jsii.String("audioOnlyTimecodeControl"),
//   			CertificateMode: jsii.String("certificateMode"),
//   			ConnectionRetryInterval: jsii.Number(123),
//   			Destination: &OutputLocationRefProperty{
//   				DestinationRefId: jsii.String("destinationRefId"),
//   			},
//   			EventId: jsii.String("eventId"),
//   			EventIdMode: jsii.String("eventIdMode"),
//   			EventStopBehavior: jsii.String("eventStopBehavior"),
//   			FilecacheDuration: jsii.Number(123),
//   			FragmentLength: jsii.Number(123),
//   			InputLossAction: jsii.String("inputLossAction"),
//   			NumRetries: jsii.Number(123),
//   			RestartDelay: jsii.Number(123),
//   			SegmentationMode: jsii.String("segmentationMode"),
//   			SendDelayMs: jsii.Number(123),
//   			SparseTrackType: jsii.String("sparseTrackType"),
//   			StreamManifestBehavior: jsii.String("streamManifestBehavior"),
//   			TimestampOffset: jsii.String("timestampOffset"),
//   			TimestampOffsetMode: jsii.String("timestampOffsetMode"),
//   		},
//   		MultiplexGroupSettings: &MultiplexGroupSettingsProperty{
//   		},
//   		RtmpGroupSettings: &RtmpGroupSettingsProperty{
//   			AdMarkers: []*string{
//   				jsii.String("adMarkers"),
//   			},
//   			AuthenticationScheme: jsii.String("authenticationScheme"),
//   			CacheFullBehavior: jsii.String("cacheFullBehavior"),
//   			CacheLength: jsii.Number(123),
//   			CaptionData: jsii.String("captionData"),
//   			IncludeFillerNalUnits: jsii.String("includeFillerNalUnits"),
//   			InputLossAction: jsii.String("inputLossAction"),
//   			RestartDelay: jsii.Number(123),
//   		},
//   		UdpGroupSettings: &UdpGroupSettingsProperty{
//   			InputLossAction: jsii.String("inputLossAction"),
//   			TimedMetadataId3Frame: jsii.String("timedMetadataId3Frame"),
//   			TimedMetadataId3Period: jsii.Number(123),
//   		},
//   	},
//   	Outputs: []interface{}{
//   		&OutputProperty{
//   			AudioDescriptionNames: []*string{
//   				jsii.String("audioDescriptionNames"),
//   			},
//   			CaptionDescriptionNames: []*string{
//   				jsii.String("captionDescriptionNames"),
//   			},
//   			OutputName: jsii.String("outputName"),
//   			OutputSettings: &OutputSettingsProperty{
//   				ArchiveOutputSettings: &ArchiveOutputSettingsProperty{
//   					ContainerSettings: &ArchiveContainerSettingsProperty{
//   						M2TsSettings: &M2tsSettingsProperty{
//   							AbsentInputAudioBehavior: jsii.String("absentInputAudioBehavior"),
//   							Arib: jsii.String("arib"),
//   							AribCaptionsPid: jsii.String("aribCaptionsPid"),
//   							AribCaptionsPidControl: jsii.String("aribCaptionsPidControl"),
//   							AudioBufferModel: jsii.String("audioBufferModel"),
//   							AudioFramesPerPes: jsii.Number(123),
//   							AudioPids: jsii.String("audioPids"),
//   							AudioStreamType: jsii.String("audioStreamType"),
//   							Bitrate: jsii.Number(123),
//   							BufferModel: jsii.String("bufferModel"),
//   							CcDescriptor: jsii.String("ccDescriptor"),
//   							DvbNitSettings: &DvbNitSettingsProperty{
//   								NetworkId: jsii.Number(123),
//   								NetworkName: jsii.String("networkName"),
//   								RepInterval: jsii.Number(123),
//   							},
//   							DvbSdtSettings: &DvbSdtSettingsProperty{
//   								OutputSdt: jsii.String("outputSdt"),
//   								RepInterval: jsii.Number(123),
//   								ServiceName: jsii.String("serviceName"),
//   								ServiceProviderName: jsii.String("serviceProviderName"),
//   							},
//   							DvbSubPids: jsii.String("dvbSubPids"),
//   							DvbTdtSettings: &DvbTdtSettingsProperty{
//   								RepInterval: jsii.Number(123),
//   							},
//   							DvbTeletextPid: jsii.String("dvbTeletextPid"),
//   							Ebif: jsii.String("ebif"),
//   							EbpAudioInterval: jsii.String("ebpAudioInterval"),
//   							EbpLookaheadMs: jsii.Number(123),
//   							EbpPlacement: jsii.String("ebpPlacement"),
//   							EcmPid: jsii.String("ecmPid"),
//   							EsRateInPes: jsii.String("esRateInPes"),
//   							EtvPlatformPid: jsii.String("etvPlatformPid"),
//   							EtvSignalPid: jsii.String("etvSignalPid"),
//   							FragmentTime: jsii.Number(123),
//   							Klv: jsii.String("klv"),
//   							KlvDataPids: jsii.String("klvDataPids"),
//   							NielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   							NullPacketBitrate: jsii.Number(123),
//   							PatInterval: jsii.Number(123),
//   							PcrControl: jsii.String("pcrControl"),
//   							PcrPeriod: jsii.Number(123),
//   							PcrPid: jsii.String("pcrPid"),
//   							PmtInterval: jsii.Number(123),
//   							PmtPid: jsii.String("pmtPid"),
//   							ProgramNum: jsii.Number(123),
//   							RateMode: jsii.String("rateMode"),
//   							Scte27Pids: jsii.String("scte27Pids"),
//   							Scte35Control: jsii.String("scte35Control"),
//   							Scte35Pid: jsii.String("scte35Pid"),
//   							Scte35PrerollPullupMilliseconds: jsii.Number(123),
//   							SegmentationMarkers: jsii.String("segmentationMarkers"),
//   							SegmentationStyle: jsii.String("segmentationStyle"),
//   							SegmentationTime: jsii.Number(123),
//   							TimedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   							TimedMetadataPid: jsii.String("timedMetadataPid"),
//   							TransportStreamId: jsii.Number(123),
//   							VideoPid: jsii.String("videoPid"),
//   						},
//   						RawSettings: &RawSettingsProperty{
//   						},
//   					},
//   					Extension: jsii.String("extension"),
//   					NameModifier: jsii.String("nameModifier"),
//   				},
//   				FrameCaptureOutputSettings: &FrameCaptureOutputSettingsProperty{
//   					NameModifier: jsii.String("nameModifier"),
//   				},
//   				HlsOutputSettings: &HlsOutputSettingsProperty{
//   					H265PackagingType: jsii.String("h265PackagingType"),
//   					HlsSettings: &HlsSettingsProperty{
//   						AudioOnlyHlsSettings: &AudioOnlyHlsSettingsProperty{
//   							AudioGroupId: jsii.String("audioGroupId"),
//   							AudioOnlyImage: &InputLocationProperty{
//   								PasswordParam: jsii.String("passwordParam"),
//   								Uri: jsii.String("uri"),
//   								Username: jsii.String("username"),
//   							},
//   							AudioTrackType: jsii.String("audioTrackType"),
//   							SegmentType: jsii.String("segmentType"),
//   						},
//   						Fmp4HlsSettings: &Fmp4HlsSettingsProperty{
//   							AudioRenditionSets: jsii.String("audioRenditionSets"),
//   							NielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   							TimedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   						},
//   						FrameCaptureHlsSettings: &FrameCaptureHlsSettingsProperty{
//   						},
//   						StandardHlsSettings: &StandardHlsSettingsProperty{
//   							AudioRenditionSets: jsii.String("audioRenditionSets"),
//   							M3U8Settings: &M3u8SettingsProperty{
//   								AudioFramesPerPes: jsii.Number(123),
//   								AudioPids: jsii.String("audioPids"),
//   								EcmPid: jsii.String("ecmPid"),
//   								KlvBehavior: jsii.String("klvBehavior"),
//   								KlvDataPids: jsii.String("klvDataPids"),
//   								NielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   								PatInterval: jsii.Number(123),
//   								PcrControl: jsii.String("pcrControl"),
//   								PcrPeriod: jsii.Number(123),
//   								PcrPid: jsii.String("pcrPid"),
//   								PmtInterval: jsii.Number(123),
//   								PmtPid: jsii.String("pmtPid"),
//   								ProgramNum: jsii.Number(123),
//   								Scte35Behavior: jsii.String("scte35Behavior"),
//   								Scte35Pid: jsii.String("scte35Pid"),
//   								TimedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   								TimedMetadataPid: jsii.String("timedMetadataPid"),
//   								TransportStreamId: jsii.Number(123),
//   								VideoPid: jsii.String("videoPid"),
//   							},
//   						},
//   					},
//   					NameModifier: jsii.String("nameModifier"),
//   					SegmentModifier: jsii.String("segmentModifier"),
//   				},
//   				MediaPackageOutputSettings: &MediaPackageOutputSettingsProperty{
//   				},
//   				MsSmoothOutputSettings: &MsSmoothOutputSettingsProperty{
//   					H265PackagingType: jsii.String("h265PackagingType"),
//   					NameModifier: jsii.String("nameModifier"),
//   				},
//   				MultiplexOutputSettings: &MultiplexOutputSettingsProperty{
//   					Destination: &OutputLocationRefProperty{
//   						DestinationRefId: jsii.String("destinationRefId"),
//   					},
//   				},
//   				RtmpOutputSettings: &RtmpOutputSettingsProperty{
//   					CertificateMode: jsii.String("certificateMode"),
//   					ConnectionRetryInterval: jsii.Number(123),
//   					Destination: &OutputLocationRefProperty{
//   						DestinationRefId: jsii.String("destinationRefId"),
//   					},
//   					NumRetries: jsii.Number(123),
//   				},
//   				UdpOutputSettings: &UdpOutputSettingsProperty{
//   					BufferMsec: jsii.Number(123),
//   					ContainerSettings: &UdpContainerSettingsProperty{
//   						M2TsSettings: &M2tsSettingsProperty{
//   							AbsentInputAudioBehavior: jsii.String("absentInputAudioBehavior"),
//   							Arib: jsii.String("arib"),
//   							AribCaptionsPid: jsii.String("aribCaptionsPid"),
//   							AribCaptionsPidControl: jsii.String("aribCaptionsPidControl"),
//   							AudioBufferModel: jsii.String("audioBufferModel"),
//   							AudioFramesPerPes: jsii.Number(123),
//   							AudioPids: jsii.String("audioPids"),
//   							AudioStreamType: jsii.String("audioStreamType"),
//   							Bitrate: jsii.Number(123),
//   							BufferModel: jsii.String("bufferModel"),
//   							CcDescriptor: jsii.String("ccDescriptor"),
//   							DvbNitSettings: &DvbNitSettingsProperty{
//   								NetworkId: jsii.Number(123),
//   								NetworkName: jsii.String("networkName"),
//   								RepInterval: jsii.Number(123),
//   							},
//   							DvbSdtSettings: &DvbSdtSettingsProperty{
//   								OutputSdt: jsii.String("outputSdt"),
//   								RepInterval: jsii.Number(123),
//   								ServiceName: jsii.String("serviceName"),
//   								ServiceProviderName: jsii.String("serviceProviderName"),
//   							},
//   							DvbSubPids: jsii.String("dvbSubPids"),
//   							DvbTdtSettings: &DvbTdtSettingsProperty{
//   								RepInterval: jsii.Number(123),
//   							},
//   							DvbTeletextPid: jsii.String("dvbTeletextPid"),
//   							Ebif: jsii.String("ebif"),
//   							EbpAudioInterval: jsii.String("ebpAudioInterval"),
//   							EbpLookaheadMs: jsii.Number(123),
//   							EbpPlacement: jsii.String("ebpPlacement"),
//   							EcmPid: jsii.String("ecmPid"),
//   							EsRateInPes: jsii.String("esRateInPes"),
//   							EtvPlatformPid: jsii.String("etvPlatformPid"),
//   							EtvSignalPid: jsii.String("etvSignalPid"),
//   							FragmentTime: jsii.Number(123),
//   							Klv: jsii.String("klv"),
//   							KlvDataPids: jsii.String("klvDataPids"),
//   							NielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   							NullPacketBitrate: jsii.Number(123),
//   							PatInterval: jsii.Number(123),
//   							PcrControl: jsii.String("pcrControl"),
//   							PcrPeriod: jsii.Number(123),
//   							PcrPid: jsii.String("pcrPid"),
//   							PmtInterval: jsii.Number(123),
//   							PmtPid: jsii.String("pmtPid"),
//   							ProgramNum: jsii.Number(123),
//   							RateMode: jsii.String("rateMode"),
//   							Scte27Pids: jsii.String("scte27Pids"),
//   							Scte35Control: jsii.String("scte35Control"),
//   							Scte35Pid: jsii.String("scte35Pid"),
//   							Scte35PrerollPullupMilliseconds: jsii.Number(123),
//   							SegmentationMarkers: jsii.String("segmentationMarkers"),
//   							SegmentationStyle: jsii.String("segmentationStyle"),
//   							SegmentationTime: jsii.Number(123),
//   							TimedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   							TimedMetadataPid: jsii.String("timedMetadataPid"),
//   							TransportStreamId: jsii.Number(123),
//   							VideoPid: jsii.String("videoPid"),
//   						},
//   					},
//   					Destination: &OutputLocationRefProperty{
//   						DestinationRefId: jsii.String("destinationRefId"),
//   					},
//   					FecOutputSettings: &FecOutputSettingsProperty{
//   						ColumnDepth: jsii.Number(123),
//   						IncludeFec: jsii.String("includeFec"),
//   						RowLength: jsii.Number(123),
//   					},
//   				},
//   			},
//   			VideoDescriptionName: jsii.String("videoDescriptionName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputgroup.html
//
type CfnChannel_OutputGroupProperty struct {
	// A custom output group name that you can optionally define.
	//
	// Only letters, numbers, and the underscore character are allowed. The maximum length is 32 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputgroup.html#cfn-medialive-channel-outputgroup-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The settings associated with the output group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputgroup.html#cfn-medialive-channel-outputgroup-outputgroupsettings
	//
	OutputGroupSettings interface{} `field:"optional" json:"outputGroupSettings" yaml:"outputGroupSettings"`
	// The settings for the outputs in the output group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputgroup.html#cfn-medialive-channel-outputgroup-outputs
	//
	Outputs interface{} `field:"optional" json:"outputs" yaml:"outputs"`
}

