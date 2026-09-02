package awsmedialivealpha


// An input attachment definition for a channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var audioSelector AudioSelector
//   var bitrate Bitrate
//   var captionSelector CaptionSelector
//   var failoverCondition FailoverCondition
//   var hlsScte35Source HlsScte35Source
//   var input Input
//   var inputFilter InputFilter
//   var inputPreference InputPreference
//   var serverValidation ServerValidation
//   var smpte2038DataPreference Smpte2038DataPreference
//   var sourceEndBehavior SourceEndBehavior
//   var videoColorSpace VideoColorSpace
//   var videoColorSpaceUsage VideoColorSpaceUsage
//   var videoSelection VideoSelection
//
//   inputAttachment := &InputAttachment{
//   	Input: input,
//
//   	// the properties below are optional
//   	AudioSelectors: []AudioSelector{
//   		audioSelector,
//   	},
//   	AutomaticInputFailover: &AutomaticInputFailover{
//   		SecondaryInput: input,
//
//   		// the properties below are optional
//   		ErrorClearTime: cdk.Duration_Minutes(jsii.Number(30)),
//   		FailoverConditions: []FailoverCondition{
//   			failoverCondition,
//   		},
//   		InputPreference: inputPreference,
//   	},
//   	CaptionSelectors: []CaptionSelector{
//   		captionSelector,
//   	},
//   	DeblockFilter: jsii.Boolean(false),
//   	DenoiseFilter: jsii.Boolean(false),
//   	FilterStrength: jsii.Number(123),
//   	InputAttachmentName: jsii.String("inputAttachmentName"),
//   	InputFilter: inputFilter,
//   	LogicalInterfaceNames: []*string{
//   		jsii.String("logicalInterfaceNames"),
//   	},
//   	NetworkInputSettings: &NetworkInputSettings{
//   		HlsInputSettings: &HlsInputSettings{
//   			Bandwidth: bitrate,
//   			BufferSegments: jsii.Number(123),
//   			Retries: jsii.Number(123),
//   			RetryInterval: cdk.Duration_*Minutes(jsii.Number(30)),
//   			Scte35Source: hlsScte35Source,
//   		},
//   		MulticastSourceIp: jsii.String("multicastSourceIp"),
//   		ServerValidation: serverValidation,
//   	},
//   	Scte35Pid: jsii.Number(123),
//   	Smpte2038DataPreference: smpte2038DataPreference,
//   	SourceEndBehavior: sourceEndBehavior,
//   	VideoSelector: &VideoSelectorSettings{
//   		ColorSpace: videoColorSpace,
//   		ColorSpaceUsage: videoColorSpaceUsage,
//   		Hdr10: &Hdr10Settings{
//   			MaxContentLightLevel: jsii.Number(123),
//   			MaxFrameAverageLightLevel: jsii.Number(123),
//   		},
//   		SelectBy: videoSelection,
//   	},
//   }
//
// Experimental.
type InputAttachment struct {
	// The input to attach.
	//
	// [disable-awslint:prefer-ref-interface].
	// Experimental.
	Input IInput `field:"required" json:"input" yaml:"input"`
	// Audio selectors to extract specific audio tracks from the input.
	// Default: - no audio selectors (uses default audio).
	//
	// Experimental.
	AudioSelectors *[]AudioSelector `field:"optional" json:"audioSelectors" yaml:"audioSelectors"`
	// Automatic input failover to a secondary input.
	//
	// When the active input meets any of the
	// failover conditions, MediaLive switches to the secondary input without restarting the
	// channel. This is input-source redundancy, distinct from the pipeline redundancy of
	// `ChannelClass.STANDARD`.
	// Default: - no automatic input failover.
	//
	// Experimental.
	AutomaticInputFailover *AutomaticInputFailover `field:"optional" json:"automaticInputFailover" yaml:"automaticInputFailover"`
	// Caption selectors to extract specific caption tracks from the input.
	// Default: - no caption selectors.
	//
	// Experimental.
	CaptionSelectors *[]CaptionSelector `field:"optional" json:"captionSelectors" yaml:"captionSelectors"`
	// Whether to enable the deblock filter.
	// Default: false.
	//
	// Experimental.
	DeblockFilter *bool `field:"optional" json:"deblockFilter" yaml:"deblockFilter"`
	// Whether to enable the denoise filter.
	// Default: false.
	//
	// Experimental.
	DenoiseFilter *bool `field:"optional" json:"denoiseFilter" yaml:"denoiseFilter"`
	// The filter strength (1-5).
	//
	// 1 is minimal, 5 is strongest.
	// Default: 1.
	//
	// Experimental.
	FilterStrength *float64 `field:"optional" json:"filterStrength" yaml:"filterStrength"`
	// A name for this input attachment, used to reference it in schedule actions.
	// Default: - auto-generated.
	//
	// Experimental.
	InputAttachmentName *string `field:"optional" json:"inputAttachmentName" yaml:"inputAttachmentName"`
	// The input filter mode.
	// Default: InputFilter.AUTO
	//
	// Experimental.
	InputFilter InputFilter `field:"optional" json:"inputFilter" yaml:"inputFilter"`
	// The logical interface names (MediaLive Anywhere) this input is wired to.
	//
	// Each name maps the
	// input to a network interface on the channel's nodes.
	// Default: - no logical interface names.
	//
	// Experimental.
	LogicalInterfaceNames *[]*string `field:"optional" json:"logicalInterfaceNames" yaml:"logicalInterfaceNames"`
	// Network input settings (for URL pull inputs — HLS buffer, server validation).
	// Default: - no network input settings.
	//
	// Experimental.
	NetworkInputSettings *NetworkInputSettings `field:"optional" json:"networkInputSettings" yaml:"networkInputSettings"`
	// The SCTE-35 PID override for this input.
	// Default: - auto-detect.
	//
	// Experimental.
	Scte35Pid *float64 `field:"optional" json:"scte35Pid" yaml:"scte35Pid"`
	// SMPTE-2038 ancillary data preference.
	// Default: Smpte2038DataPreference.IGNORE
	//
	// Experimental.
	Smpte2038DataPreference Smpte2038DataPreference `field:"optional" json:"smpte2038DataPreference" yaml:"smpte2038DataPreference"`
	// The source end behavior for file-based inputs.
	// Default: SourceEndBehavior.LOOP for MP4 and TS file inputs, SourceEndBehavior.CONTINUE for all others
	//
	// Experimental.
	SourceEndBehavior SourceEndBehavior `field:"optional" json:"sourceEndBehavior" yaml:"sourceEndBehavior"`
	// Video selector settings for the input (color space, PID selection).
	// Default: - no video selector (use default video).
	//
	// Experimental.
	VideoSelector *VideoSelectorSettings `field:"optional" json:"videoSelector" yaml:"videoSelector"`
}

