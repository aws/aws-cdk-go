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

