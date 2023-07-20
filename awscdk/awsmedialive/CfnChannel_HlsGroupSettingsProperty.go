package awsmedialive


// The settings for an HLS output group.
//
// The parent of this entity is OutputGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsGroupSettingsProperty := &HlsGroupSettingsProperty{
//   	AdMarkers: []*string{
//   		jsii.String("adMarkers"),
//   	},
//   	BaseUrlContent: jsii.String("baseUrlContent"),
//   	BaseUrlContent1: jsii.String("baseUrlContent1"),
//   	BaseUrlManifest: jsii.String("baseUrlManifest"),
//   	BaseUrlManifest1: jsii.String("baseUrlManifest1"),
//   	CaptionLanguageMappings: []interface{}{
//   		&CaptionLanguageMappingProperty{
//   			CaptionChannel: jsii.Number(123),
//   			LanguageCode: jsii.String("languageCode"),
//   			LanguageDescription: jsii.String("languageDescription"),
//   		},
//   	},
//   	CaptionLanguageSetting: jsii.String("captionLanguageSetting"),
//   	ClientCache: jsii.String("clientCache"),
//   	CodecSpecification: jsii.String("codecSpecification"),
//   	ConstantIv: jsii.String("constantIv"),
//   	Destination: &OutputLocationRefProperty{
//   		DestinationRefId: jsii.String("destinationRefId"),
//   	},
//   	DirectoryStructure: jsii.String("directoryStructure"),
//   	DiscontinuityTags: jsii.String("discontinuityTags"),
//   	EncryptionType: jsii.String("encryptionType"),
//   	HlsCdnSettings: &HlsCdnSettingsProperty{
//   		HlsAkamaiSettings: &HlsAkamaiSettingsProperty{
//   			ConnectionRetryInterval: jsii.Number(123),
//   			FilecacheDuration: jsii.Number(123),
//   			HttpTransferMode: jsii.String("httpTransferMode"),
//   			NumRetries: jsii.Number(123),
//   			RestartDelay: jsii.Number(123),
//   			Salt: jsii.String("salt"),
//   			Token: jsii.String("token"),
//   		},
//   		HlsBasicPutSettings: &HlsBasicPutSettingsProperty{
//   			ConnectionRetryInterval: jsii.Number(123),
//   			FilecacheDuration: jsii.Number(123),
//   			NumRetries: jsii.Number(123),
//   			RestartDelay: jsii.Number(123),
//   		},
//   		HlsMediaStoreSettings: &HlsMediaStoreSettingsProperty{
//   			ConnectionRetryInterval: jsii.Number(123),
//   			FilecacheDuration: jsii.Number(123),
//   			MediaStoreStorageClass: jsii.String("mediaStoreStorageClass"),
//   			NumRetries: jsii.Number(123),
//   			RestartDelay: jsii.Number(123),
//   		},
//   		HlsS3Settings: &HlsS3SettingsProperty{
//   			CannedAcl: jsii.String("cannedAcl"),
//   		},
//   		HlsWebdavSettings: &HlsWebdavSettingsProperty{
//   			ConnectionRetryInterval: jsii.Number(123),
//   			FilecacheDuration: jsii.Number(123),
//   			HttpTransferMode: jsii.String("httpTransferMode"),
//   			NumRetries: jsii.Number(123),
//   			RestartDelay: jsii.Number(123),
//   		},
//   	},
//   	HlsId3SegmentTagging: jsii.String("hlsId3SegmentTagging"),
//   	IFrameOnlyPlaylists: jsii.String("iFrameOnlyPlaylists"),
//   	IncompleteSegmentBehavior: jsii.String("incompleteSegmentBehavior"),
//   	IndexNSegments: jsii.Number(123),
//   	InputLossAction: jsii.String("inputLossAction"),
//   	IvInManifest: jsii.String("ivInManifest"),
//   	IvSource: jsii.String("ivSource"),
//   	KeepSegments: jsii.Number(123),
//   	KeyFormat: jsii.String("keyFormat"),
//   	KeyFormatVersions: jsii.String("keyFormatVersions"),
//   	KeyProviderSettings: &KeyProviderSettingsProperty{
//   		StaticKeySettings: &StaticKeySettingsProperty{
//   			KeyProviderServer: &InputLocationProperty{
//   				PasswordParam: jsii.String("passwordParam"),
//   				Uri: jsii.String("uri"),
//   				Username: jsii.String("username"),
//   			},
//   			StaticKeyValue: jsii.String("staticKeyValue"),
//   		},
//   	},
//   	ManifestCompression: jsii.String("manifestCompression"),
//   	ManifestDurationFormat: jsii.String("manifestDurationFormat"),
//   	MinSegmentLength: jsii.Number(123),
//   	Mode: jsii.String("mode"),
//   	OutputSelection: jsii.String("outputSelection"),
//   	ProgramDateTime: jsii.String("programDateTime"),
//   	ProgramDateTimeClock: jsii.String("programDateTimeClock"),
//   	ProgramDateTimePeriod: jsii.Number(123),
//   	RedundantManifest: jsii.String("redundantManifest"),
//   	SegmentationMode: jsii.String("segmentationMode"),
//   	SegmentLength: jsii.Number(123),
//   	SegmentsPerSubdirectory: jsii.Number(123),
//   	StreamInfResolution: jsii.String("streamInfResolution"),
//   	TimedMetadataId3Frame: jsii.String("timedMetadataId3Frame"),
//   	TimedMetadataId3Period: jsii.Number(123),
//   	TimestampDeltaMilliseconds: jsii.Number(123),
//   	TsFileMode: jsii.String("tsFileMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html
//
type CfnChannel_HlsGroupSettingsProperty struct {
	// Chooses one or more ad marker types to pass SCTE35 signals through to this group of Apple HLS outputs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-admarkers
	//
	AdMarkers *[]*string `field:"optional" json:"adMarkers" yaml:"adMarkers"`
	// A partial URI prefix that will be prepended to each output in the media .m3u8 file. The partial URI prefix can be used if the base manifest is delivered from a different URL than the main .m3u8 file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-baseurlcontent
	//
	BaseUrlContent *string `field:"optional" json:"baseUrlContent" yaml:"baseUrlContent"`
	// Optional.
	//
	// One value per output group. This field is required only if you are completing Base URL content A, and the downstream system has notified you that the media files for pipeline 1 of all outputs are in a location different from the media files for pipeline 0.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-baseurlcontent1
	//
	BaseUrlContent1 *string `field:"optional" json:"baseUrlContent1" yaml:"baseUrlContent1"`
	// A partial URI prefix that will be prepended to each output in the media .m3u8 file. The partial URI prefix can be used if the base manifest is delivered from a different URL than the main .m3u8 file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-baseurlmanifest
	//
	BaseUrlManifest *string `field:"optional" json:"baseUrlManifest" yaml:"baseUrlManifest"`
	// Optional.
	//
	// One value per output group. Complete this field only if you are completing Base URL manifest A, and the downstream system has notified you that the child manifest files for pipeline 1 of all outputs are in a location different from the child manifest files for pipeline 0.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-baseurlmanifest1
	//
	BaseUrlManifest1 *string `field:"optional" json:"baseUrlManifest1" yaml:"baseUrlManifest1"`
	// A mapping of up to 4 captions channels to captions languages.
	//
	// This is meaningful only if captionLanguageSetting is set to "insert."
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-captionlanguagemappings
	//
	CaptionLanguageMappings interface{} `field:"optional" json:"captionLanguageMappings" yaml:"captionLanguageMappings"`
	// Applies only to 608 embedded output captions.
	//
	// Insert: Include CLOSED-CAPTIONS lines in the manifest. Specify at least one language in the CC1 Language Code field. One CLOSED-CAPTION line is added for each Language Code that you specify. Make sure to specify the languages in the order in which they appear in the original source (if the source is embedded format) or the order of the captions selectors (if the source is other than embedded). Otherwise, languages in the manifest will not match properly with the output captions. None: Include the CLOSED-CAPTIONS=NONE line in the manifest. Omit: Omit any CLOSED-CAPTIONS line from the manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-captionlanguagesetting
	//
	CaptionLanguageSetting *string `field:"optional" json:"captionLanguageSetting" yaml:"captionLanguageSetting"`
	// When set to "disabled," sets the #EXT-X-ALLOW-CACHE:no tag in the manifest, which prevents clients from saving media segments for later replay.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-clientcache
	//
	ClientCache *string `field:"optional" json:"clientCache" yaml:"clientCache"`
	// The specification to use (RFC-6381 or the default RFC-4281) during m3u8 playlist generation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-codecspecification
	//
	CodecSpecification *string `field:"optional" json:"codecSpecification" yaml:"codecSpecification"`
	// Used with encryptionType.
	//
	// This is a 128-bit, 16-byte hex value that is represented by a 32-character text string. If ivSource is set to "explicit," this parameter is required and is used as the IV for encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-constantiv
	//
	ConstantIv *string `field:"optional" json:"constantIv" yaml:"constantIv"`
	// A directory or HTTP destination for the HLS segments, manifest files, and encryption keys (if enabled).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-destination
	//
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// Places segments in subdirectories.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-directorystructure
	//
	DirectoryStructure *string `field:"optional" json:"directoryStructure" yaml:"directoryStructure"`
	// Specifies whether to insert EXT-X-DISCONTINUITY tags in the HLS child manifests for this output group.
	//
	// Typically, choose Insert because these tags are required in the manifest (according to the HLS specification) and serve an important purpose.
	// Choose Never Insert only if the downstream system is doing real-time failover (without using the MediaLive automatic failover feature) and only if that downstream system has advised you to exclude the tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-discontinuitytags
	//
	DiscontinuityTags *string `field:"optional" json:"discontinuityTags" yaml:"discontinuityTags"`
	// Encrypts the segments with the specified encryption scheme.
	//
	// Exclude this parameter if you don't want encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-encryptiontype
	//
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// The parameters that control interactions with the CDN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-hlscdnsettings
	//
	HlsCdnSettings interface{} `field:"optional" json:"hlsCdnSettings" yaml:"hlsCdnSettings"`
	// State of HLS ID3 Segment Tagging.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-hlsid3segmenttagging
	//
	HlsId3SegmentTagging *string `field:"optional" json:"hlsId3SegmentTagging" yaml:"hlsId3SegmentTagging"`
	// DISABLED: Don't create an I-frame-only manifest, but do create the master and media manifests (according to the Output Selection field).
	//
	// STANDARD: Create an I-frame-only manifest for each output that contains video, as well as the other manifests (according to the Output Selection field). The I-frame manifest contains a #EXT-X-I-FRAMES-ONLY tag to indicate it is I-frame only, and one or more #EXT-X-BYTERANGE entries identifying the I-frame position. For example, #EXT-X-BYTERANGE:160364@1461888".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-iframeonlyplaylists
	//
	IFrameOnlyPlaylists *string `field:"optional" json:"iFrameOnlyPlaylists" yaml:"iFrameOnlyPlaylists"`
	// Specifies whether to include the final (incomplete) segment in the media output when the pipeline stops producing output because of a channel stop, a channel pause or a loss of input to the pipeline.
	//
	// Auto means that MediaLive decides whether to include the final segment, depending on the channel class and the types of output groups.
	// Suppress means to never include the incomplete segment. We recommend you choose Auto and let MediaLive control the behavior.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-incompletesegmentbehavior
	//
	IncompleteSegmentBehavior *string `field:"optional" json:"incompleteSegmentBehavior" yaml:"incompleteSegmentBehavior"`
	// Applies only if the Mode field is LIVE.
	//
	// Specifies the maximum number of segments in the media manifest file. After this maximum, older segments are removed from the media manifest. This number must be less than or equal to the Keep Segments field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-indexnsegments
	//
	IndexNSegments *float64 `field:"optional" json:"indexNSegments" yaml:"indexNSegments"`
	// A parameter that controls output group behavior on an input loss.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-inputlossaction
	//
	InputLossAction *string `field:"optional" json:"inputLossAction" yaml:"inputLossAction"`
	// Used with encryptionType.
	//
	// The IV (initialization vector) is a 128-bit number used in conjunction with the key for encrypting blocks. If set to "include," the IV is listed in the manifest. Otherwise, the IV is not in the manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-ivinmanifest
	//
	IvInManifest *string `field:"optional" json:"ivInManifest" yaml:"ivInManifest"`
	// Used with encryptionType.
	//
	// The IV (initialization vector) is a 128-bit number used in conjunction with the key for encrypting blocks. If this setting is "followsSegmentNumber," it causes the IV to change every segment (to match the segment number). If this is set to "explicit," you must enter a constantIv value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-ivsource
	//
	IvSource *string `field:"optional" json:"ivSource" yaml:"ivSource"`
	// Applies only if the Mode field is LIVE.
	//
	// Specifies the number of media segments (.ts files) to retain in the destination directory.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-keepsegments
	//
	KeepSegments *float64 `field:"optional" json:"keepSegments" yaml:"keepSegments"`
	// Specifies how the key is represented in the resource identified by the URI.
	//
	// If the parameter is absent, an implicit value of "identity" is used. A reverse DNS string can also be specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-keyformat
	//
	KeyFormat *string `field:"optional" json:"keyFormat" yaml:"keyFormat"`
	// Either a single positive integer version value or a slash-delimited list of version values (1/2/3).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-keyformatversions
	//
	KeyFormatVersions *string `field:"optional" json:"keyFormatVersions" yaml:"keyFormatVersions"`
	// The key provider settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-keyprovidersettings
	//
	KeyProviderSettings interface{} `field:"optional" json:"keyProviderSettings" yaml:"keyProviderSettings"`
	// When set to gzip, compresses HLS playlist.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-manifestcompression
	//
	ManifestCompression *string `field:"optional" json:"manifestCompression" yaml:"manifestCompression"`
	// Indicates whether the output manifest should use a floating point or integer values for segment duration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-manifestdurationformat
	//
	ManifestDurationFormat *string `field:"optional" json:"manifestDurationFormat" yaml:"manifestDurationFormat"`
	// When set, minimumSegmentLength is enforced by looking ahead and back within the specified range for a nearby avail and extending the segment size if needed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-minsegmentlength
	//
	MinSegmentLength *float64 `field:"optional" json:"minSegmentLength" yaml:"minSegmentLength"`
	// If "vod," all segments are indexed and kept permanently in the destination and manifest.
	//
	// If "live," only the number segments specified in keepSegments and indexNSegments are kept. Newer segments replace older segments, which might prevent players from rewinding all the way to the beginning of the channel. VOD mode uses HLS EXT-X-PLAYLIST-TYPE of EVENT while the channel is running, converting it to a "VOD" type manifest on completion of the stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// MANIFESTSANDSEGMENTS: Generates manifests (the master manifest, if applicable, and media manifests) for this output group.
	//
	// SEGMENTSONLY: Doesn't generate any manifests for this output group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-outputselection
	//
	OutputSelection *string `field:"optional" json:"outputSelection" yaml:"outputSelection"`
	// Includes or excludes the EXT-X-PROGRAM-DATE-TIME tag in .m3u8 manifest files. The value is calculated as follows: Either the program date and time are initialized using the input timecode source, or the time is initialized using the input timecode source and the date is initialized using the timestampOffset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-programdatetime
	//
	ProgramDateTime *string `field:"optional" json:"programDateTime" yaml:"programDateTime"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-programdatetimeclock
	//
	ProgramDateTimeClock *string `field:"optional" json:"programDateTimeClock" yaml:"programDateTimeClock"`
	// The period of insertion of the EXT-X-PROGRAM-DATE-TIME entry, in seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-programdatetimeperiod
	//
	ProgramDateTimePeriod *float64 `field:"optional" json:"programDateTimePeriod" yaml:"programDateTimePeriod"`
	// ENABLED: The master manifest (.m3u8 file) for each pipeline includes information about both pipelines: first its own media files, then the media files of the other pipeline. This feature allows a playout device that supports stale manifest detection to switch from one manifest to the other, when the current manifest seems to be stale. There are still two destinations and two master manifests, but both master manifests reference the media files from both pipelines. DISABLED: The master manifest (.m3u8 file) for each pipeline includes information about its own pipeline only. For an HLS output group with MediaPackage as the destination, the DISABLED behavior is always followed. MediaPackage regenerates the manifests it serves to players, so a redundant manifest from MediaLive is irrelevant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-redundantmanifest
	//
	RedundantManifest *string `field:"optional" json:"redundantManifest" yaml:"redundantManifest"`
	// useInputSegmentation has been deprecated.
	//
	// The configured segment size is always used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-segmentationmode
	//
	SegmentationMode *string `field:"optional" json:"segmentationMode" yaml:"segmentationMode"`
	// The length of the MPEG-2 Transport Stream segments to create, in seconds.
	//
	// Note that segments will end on the next keyframe after this number of seconds, so the actual segment length might be longer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-segmentlength
	//
	SegmentLength *float64 `field:"optional" json:"segmentLength" yaml:"segmentLength"`
	// The number of segments to write to a subdirectory before starting a new one.
	//
	// For this setting to have an effect, directoryStructure must be subdirectoryPerStream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-segmentspersubdirectory
	//
	SegmentsPerSubdirectory *float64 `field:"optional" json:"segmentsPerSubdirectory" yaml:"segmentsPerSubdirectory"`
	// The include or exclude RESOLUTION attribute for a video in the EXT-X-STREAM-INF tag of a variant manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-streaminfresolution
	//
	StreamInfResolution *string `field:"optional" json:"streamInfResolution" yaml:"streamInfResolution"`
	// Indicates the ID3 frame that has the timecode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-timedmetadataid3frame
	//
	TimedMetadataId3Frame *string `field:"optional" json:"timedMetadataId3Frame" yaml:"timedMetadataId3Frame"`
	// The timed metadata interval, in seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-timedmetadataid3period
	//
	TimedMetadataId3Period *float64 `field:"optional" json:"timedMetadataId3Period" yaml:"timedMetadataId3Period"`
	// Provides an extra millisecond delta offset to fine tune the timestamps.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-timestampdeltamilliseconds
	//
	TimestampDeltaMilliseconds *float64 `field:"optional" json:"timestampDeltaMilliseconds" yaml:"timestampDeltaMilliseconds"`
	// SEGMENTEDFILES: Emits the program as segments -multiple .ts media files. SINGLEFILE: Applies only if the Mode field is VOD. Emits the program as a single .ts media file. The media manifest includes #EXT-X-BYTERANGE tags to index segments for playback. A typical use for this value is when sending the output to AWS Elemental MediaConvert, which can accept only a single media file. Playback while the channel is running is not guaranteed due to HTTP server caching.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsgroupsettings.html#cfn-medialive-channel-hlsgroupsettings-tsfilemode
	//
	TsFileMode *string `field:"optional" json:"tsFileMode" yaml:"tsFileMode"`
}

