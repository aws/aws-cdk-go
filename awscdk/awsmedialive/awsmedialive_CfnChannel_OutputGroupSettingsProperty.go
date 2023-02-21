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

