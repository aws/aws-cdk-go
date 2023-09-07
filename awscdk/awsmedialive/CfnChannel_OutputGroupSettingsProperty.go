package awsmedialive


// The configuration of the output group.
//
// The parent of this entity is OutputGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputGroupSettingsProperty := &OutputGroupSettingsProperty{
//   	ArchiveGroupSettings: &ArchiveGroupSettingsProperty{
//   		ArchiveCdnSettings: &ArchiveCdnSettingsProperty{
//   			ArchiveS3Settings: &ArchiveS3SettingsProperty{
//   				CannedAcl: jsii.String("cannedAcl"),
//   			},
//   		},
//   		Destination: &OutputLocationRefProperty{
//   			DestinationRefId: jsii.String("destinationRefId"),
//   		},
//   		RolloverInterval: jsii.Number(123),
//   	},
//   	FrameCaptureGroupSettings: &FrameCaptureGroupSettingsProperty{
//   		Destination: &OutputLocationRefProperty{
//   			DestinationRefId: jsii.String("destinationRefId"),
//   		},
//   		FrameCaptureCdnSettings: &FrameCaptureCdnSettingsProperty{
//   			FrameCaptureS3Settings: &FrameCaptureS3SettingsProperty{
//   				CannedAcl: jsii.String("cannedAcl"),
//   			},
//   		},
//   	},
//   	HlsGroupSettings: &HlsGroupSettingsProperty{
//   		AdMarkers: []*string{
//   			jsii.String("adMarkers"),
//   		},
//   		BaseUrlContent: jsii.String("baseUrlContent"),
//   		BaseUrlContent1: jsii.String("baseUrlContent1"),
//   		BaseUrlManifest: jsii.String("baseUrlManifest"),
//   		BaseUrlManifest1: jsii.String("baseUrlManifest1"),
//   		CaptionLanguageMappings: []interface{}{
//   			&CaptionLanguageMappingProperty{
//   				CaptionChannel: jsii.Number(123),
//   				LanguageCode: jsii.String("languageCode"),
//   				LanguageDescription: jsii.String("languageDescription"),
//   			},
//   		},
//   		CaptionLanguageSetting: jsii.String("captionLanguageSetting"),
//   		ClientCache: jsii.String("clientCache"),
//   		CodecSpecification: jsii.String("codecSpecification"),
//   		ConstantIv: jsii.String("constantIv"),
//   		Destination: &OutputLocationRefProperty{
//   			DestinationRefId: jsii.String("destinationRefId"),
//   		},
//   		DirectoryStructure: jsii.String("directoryStructure"),
//   		DiscontinuityTags: jsii.String("discontinuityTags"),
//   		EncryptionType: jsii.String("encryptionType"),
//   		HlsCdnSettings: &HlsCdnSettingsProperty{
//   			HlsAkamaiSettings: &HlsAkamaiSettingsProperty{
//   				ConnectionRetryInterval: jsii.Number(123),
//   				FilecacheDuration: jsii.Number(123),
//   				HttpTransferMode: jsii.String("httpTransferMode"),
//   				NumRetries: jsii.Number(123),
//   				RestartDelay: jsii.Number(123),
//   				Salt: jsii.String("salt"),
//   				Token: jsii.String("token"),
//   			},
//   			HlsBasicPutSettings: &HlsBasicPutSettingsProperty{
//   				ConnectionRetryInterval: jsii.Number(123),
//   				FilecacheDuration: jsii.Number(123),
//   				NumRetries: jsii.Number(123),
//   				RestartDelay: jsii.Number(123),
//   			},
//   			HlsMediaStoreSettings: &HlsMediaStoreSettingsProperty{
//   				ConnectionRetryInterval: jsii.Number(123),
//   				FilecacheDuration: jsii.Number(123),
//   				MediaStoreStorageClass: jsii.String("mediaStoreStorageClass"),
//   				NumRetries: jsii.Number(123),
//   				RestartDelay: jsii.Number(123),
//   			},
//   			HlsS3Settings: &HlsS3SettingsProperty{
//   				CannedAcl: jsii.String("cannedAcl"),
//   			},
//   			HlsWebdavSettings: &HlsWebdavSettingsProperty{
//   				ConnectionRetryInterval: jsii.Number(123),
//   				FilecacheDuration: jsii.Number(123),
//   				HttpTransferMode: jsii.String("httpTransferMode"),
//   				NumRetries: jsii.Number(123),
//   				RestartDelay: jsii.Number(123),
//   			},
//   		},
//   		HlsId3SegmentTagging: jsii.String("hlsId3SegmentTagging"),
//   		IFrameOnlyPlaylists: jsii.String("iFrameOnlyPlaylists"),
//   		IncompleteSegmentBehavior: jsii.String("incompleteSegmentBehavior"),
//   		IndexNSegments: jsii.Number(123),
//   		InputLossAction: jsii.String("inputLossAction"),
//   		IvInManifest: jsii.String("ivInManifest"),
//   		IvSource: jsii.String("ivSource"),
//   		KeepSegments: jsii.Number(123),
//   		KeyFormat: jsii.String("keyFormat"),
//   		KeyFormatVersions: jsii.String("keyFormatVersions"),
//   		KeyProviderSettings: &KeyProviderSettingsProperty{
//   			StaticKeySettings: &StaticKeySettingsProperty{
//   				KeyProviderServer: &InputLocationProperty{
//   					PasswordParam: jsii.String("passwordParam"),
//   					Uri: jsii.String("uri"),
//   					Username: jsii.String("username"),
//   				},
//   				StaticKeyValue: jsii.String("staticKeyValue"),
//   			},
//   		},
//   		ManifestCompression: jsii.String("manifestCompression"),
//   		ManifestDurationFormat: jsii.String("manifestDurationFormat"),
//   		MinSegmentLength: jsii.Number(123),
//   		Mode: jsii.String("mode"),
//   		OutputSelection: jsii.String("outputSelection"),
//   		ProgramDateTime: jsii.String("programDateTime"),
//   		ProgramDateTimeClock: jsii.String("programDateTimeClock"),
//   		ProgramDateTimePeriod: jsii.Number(123),
//   		RedundantManifest: jsii.String("redundantManifest"),
//   		SegmentationMode: jsii.String("segmentationMode"),
//   		SegmentLength: jsii.Number(123),
//   		SegmentsPerSubdirectory: jsii.Number(123),
//   		StreamInfResolution: jsii.String("streamInfResolution"),
//   		TimedMetadataId3Frame: jsii.String("timedMetadataId3Frame"),
//   		TimedMetadataId3Period: jsii.Number(123),
//   		TimestampDeltaMilliseconds: jsii.Number(123),
//   		TsFileMode: jsii.String("tsFileMode"),
//   	},
//   	MediaPackageGroupSettings: &MediaPackageGroupSettingsProperty{
//   		Destination: &OutputLocationRefProperty{
//   			DestinationRefId: jsii.String("destinationRefId"),
//   		},
//   	},
//   	MsSmoothGroupSettings: &MsSmoothGroupSettingsProperty{
//   		AcquisitionPointId: jsii.String("acquisitionPointId"),
//   		AudioOnlyTimecodeControl: jsii.String("audioOnlyTimecodeControl"),
//   		CertificateMode: jsii.String("certificateMode"),
//   		ConnectionRetryInterval: jsii.Number(123),
//   		Destination: &OutputLocationRefProperty{
//   			DestinationRefId: jsii.String("destinationRefId"),
//   		},
//   		EventId: jsii.String("eventId"),
//   		EventIdMode: jsii.String("eventIdMode"),
//   		EventStopBehavior: jsii.String("eventStopBehavior"),
//   		FilecacheDuration: jsii.Number(123),
//   		FragmentLength: jsii.Number(123),
//   		InputLossAction: jsii.String("inputLossAction"),
//   		NumRetries: jsii.Number(123),
//   		RestartDelay: jsii.Number(123),
//   		SegmentationMode: jsii.String("segmentationMode"),
//   		SendDelayMs: jsii.Number(123),
//   		SparseTrackType: jsii.String("sparseTrackType"),
//   		StreamManifestBehavior: jsii.String("streamManifestBehavior"),
//   		TimestampOffset: jsii.String("timestampOffset"),
//   		TimestampOffsetMode: jsii.String("timestampOffsetMode"),
//   	},
//   	MultiplexGroupSettings: &MultiplexGroupSettingsProperty{
//   	},
//   	RtmpGroupSettings: &RtmpGroupSettingsProperty{
//   		AdMarkers: []*string{
//   			jsii.String("adMarkers"),
//   		},
//   		AuthenticationScheme: jsii.String("authenticationScheme"),
//   		CacheFullBehavior: jsii.String("cacheFullBehavior"),
//   		CacheLength: jsii.Number(123),
//   		CaptionData: jsii.String("captionData"),
//   		IncludeFillerNalUnits: jsii.String("includeFillerNalUnits"),
//   		InputLossAction: jsii.String("inputLossAction"),
//   		RestartDelay: jsii.Number(123),
//   	},
//   	UdpGroupSettings: &UdpGroupSettingsProperty{
//   		InputLossAction: jsii.String("inputLossAction"),
//   		TimedMetadataId3Frame: jsii.String("timedMetadataId3Frame"),
//   		TimedMetadataId3Period: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputgroupsettings.html
//
type CfnChannel_OutputGroupSettingsProperty struct {
	// The configuration of an archive output group.
	//
	// The parent of this entity is OutputGroupSettings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputgroupsettings.html#cfn-medialive-channel-outputgroupsettings-archivegroupsettings
	//
	ArchiveGroupSettings interface{} `field:"optional" json:"archiveGroupSettings" yaml:"archiveGroupSettings"`
	// The configuration of a frame capture output group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputgroupsettings.html#cfn-medialive-channel-outputgroupsettings-framecapturegroupsettings
	//
	FrameCaptureGroupSettings interface{} `field:"optional" json:"frameCaptureGroupSettings" yaml:"frameCaptureGroupSettings"`
	// The configuration of an HLS output group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputgroupsettings.html#cfn-medialive-channel-outputgroupsettings-hlsgroupsettings
	//
	HlsGroupSettings interface{} `field:"optional" json:"hlsGroupSettings" yaml:"hlsGroupSettings"`
	// The configuration of a MediaPackage output group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputgroupsettings.html#cfn-medialive-channel-outputgroupsettings-mediapackagegroupsettings
	//
	MediaPackageGroupSettings interface{} `field:"optional" json:"mediaPackageGroupSettings" yaml:"mediaPackageGroupSettings"`
	// The configuration of a Microsoft Smooth output group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputgroupsettings.html#cfn-medialive-channel-outputgroupsettings-mssmoothgroupsettings
	//
	MsSmoothGroupSettings interface{} `field:"optional" json:"msSmoothGroupSettings" yaml:"msSmoothGroupSettings"`
	// The settings for a Multiplex output group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputgroupsettings.html#cfn-medialive-channel-outputgroupsettings-multiplexgroupsettings
	//
	MultiplexGroupSettings interface{} `field:"optional" json:"multiplexGroupSettings" yaml:"multiplexGroupSettings"`
	// The configuration of an RTMP output group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputgroupsettings.html#cfn-medialive-channel-outputgroupsettings-rtmpgroupsettings
	//
	RtmpGroupSettings interface{} `field:"optional" json:"rtmpGroupSettings" yaml:"rtmpGroupSettings"`
	// The configuration of a UDP output group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputgroupsettings.html#cfn-medialive-channel-outputgroupsettings-udpgroupsettings
	//
	UdpGroupSettings interface{} `field:"optional" json:"udpGroupSettings" yaml:"udpGroupSettings"`
}

