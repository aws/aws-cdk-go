package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Configuration for an output group.
//
// Use the static factory methods to create.
//
// Example:
//   var stack Stack
//   var input IInput
//   var bucket IBucket
//   var video EncodeConfiguration
//
//
//   medialive.NewChannel(stack, jsii.String("Channel"), &ChannelProps{
//   	Inputs: []InputAttachment{
//   		&InputAttachment{
//   			Input: *Input,
//   			NetworkInputSettings: &NetworkInputSettings{
//   				ServerValidation: medialive.ServerValidation_CHECK_CRYPTOGRAPHY_AND_VALIDATE_NAME(),
//   				HlsInputSettings: &HlsInputSettings{
//   					Bandwidth: awscdk.Bitrate_Mbps(jsii.Number(5)),
//   					Scte35Source: medialive.HlsScte35Source_MANIFEST(),
//   				},
//   			},
//   			LogicalInterfaceNames: []*string{
//   				jsii.String("eth0"),
//   				jsii.String("eth1"),
//   			},
//   		},
//   	},
//   	OutputGroups: []OutputGroupConfiguration{
//   		medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
//   			Name: jsii.String("hls"),
//   			Destinations: []OutputDestination{
//   				medialive.OutputDestination_ToBucket(bucket, jsii.String("live/stream")),
//   			},
//   			Outputs: []HlsOutputDefinition{
//   				&HlsOutputDefinition{
//   					Encodes: []EncodeConfiguration{
//   						video,
//   					},
//   					OutputName: jsii.String("hls_out"),
//   				},
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type OutputGroupConfiguration interface {
}

// The jsii proxy struct for OutputGroupConfiguration
type jsiiProxy_OutputGroupConfiguration struct {
	_ byte // padding
}

// Experimental.
func NewOutputGroupConfiguration_Override(o OutputGroupConfiguration) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-medialive-alpha.OutputGroupConfiguration",
		nil, // no parameters
		o,
	)
}

// Create an Archive (S3) output group configuration.
// Experimental.
func OutputGroupConfiguration_Archive(props *ArchiveOutputGroupProps) OutputGroupConfiguration {
	_init_.Initialize()

	if err := validateOutputGroupConfiguration_ArchiveParameters(props); err != nil {
		panic(err)
	}
	var returns OutputGroupConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.OutputGroupConfiguration",
		"archive",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create a CMAF Ingest output group configuration.
// Experimental.
func OutputGroupConfiguration_CmafIngest(props *CmafIngestOutputGroupProps) OutputGroupConfiguration {
	_init_.Initialize()

	if err := validateOutputGroupConfiguration_CmafIngestParameters(props); err != nil {
		panic(err)
	}
	var returns OutputGroupConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.OutputGroupConfiguration",
		"cmafIngest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create a Frame Capture output group configuration.
//
// A channel that includes a Frame Capture output group must also include
// a separate video output group (e.g. Archive, HLS, UDP). Frame Capture
// cannot be the channel's only output group.
// Experimental.
func OutputGroupConfiguration_FrameCapture(props *FrameCaptureOutputGroupProps) OutputGroupConfiguration {
	_init_.Initialize()

	if err := validateOutputGroupConfiguration_FrameCaptureParameters(props); err != nil {
		panic(err)
	}
	var returns OutputGroupConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.OutputGroupConfiguration",
		"frameCapture",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create an HLS output group configuration.
// Experimental.
func OutputGroupConfiguration_Hls(props *HlsOutputGroupProps) OutputGroupConfiguration {
	_init_.Initialize()

	if err := validateOutputGroupConfiguration_HlsParameters(props); err != nil {
		panic(err)
	}
	var returns OutputGroupConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.OutputGroupConfiguration",
		"hls",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create a MediaConnect Router output group, delivering each channel pipeline to a MediaConnect Router.
//
// Transit encryption defaults to AUTOMATIC; override per pipeline via `routerSettings`.
//
// The downstream wiring (which router input each pipeline feeds) is configured on the
// MediaConnect side, referencing this output group by `name` and the pipeline id.
// Experimental.
func OutputGroupConfiguration_MediaConnectRouter(props *MediaConnectRouterOutputGroupProps) OutputGroupConfiguration {
	_init_.Initialize()

	if err := validateOutputGroupConfiguration_MediaConnectRouterParameters(props); err != nil {
		panic(err)
	}
	var returns OutputGroupConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.OutputGroupConfiguration",
		"mediaConnectRouter",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create a MediaPackage V2 output group that delivers to a single channel, auto-mapping each pipeline to a MediaPackage ingest endpoint based on the channel class.
// Experimental.
func OutputGroupConfiguration_MediaPackageV2(props *MediaPackageV2OutputGroupProps) OutputGroupConfiguration {
	_init_.Initialize()

	if err := validateOutputGroupConfiguration_MediaPackageV2Parameters(props); err != nil {
		panic(err)
	}
	var returns OutputGroupConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.OutputGroupConfiguration",
		"mediaPackageV2",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create a MediaPackage V2 output group with explicit per-pipeline destinations (channel + endpoint per pipeline).
//
// Use for cross-region delivery or pinning a pipeline to an endpoint.
// Experimental.
func OutputGroupConfiguration_MediaPackageV2PerPipeline(props *MediaPackageV2PerPipelineOutputGroupProps) OutputGroupConfiguration {
	_init_.Initialize()

	if err := validateOutputGroupConfiguration_MediaPackageV2PerPipelineParameters(props); err != nil {
		panic(err)
	}
	var returns OutputGroupConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.OutputGroupConfiguration",
		"mediaPackageV2PerPipeline",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create an MS Smooth output group configuration.
// Experimental.
func OutputGroupConfiguration_MsSmooth(props *MsSmoothOutputGroupProps) OutputGroupConfiguration {
	_init_.Initialize()

	if err := validateOutputGroupConfiguration_MsSmoothParameters(props); err != nil {
		panic(err)
	}
	var returns OutputGroupConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.OutputGroupConfiguration",
		"msSmooth",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create an RTMP output group configuration.
// Experimental.
func OutputGroupConfiguration_Rtmp(props *RtmpOutputGroupProps) OutputGroupConfiguration {
	_init_.Initialize()

	if err := validateOutputGroupConfiguration_RtmpParameters(props); err != nil {
		panic(err)
	}
	var returns OutputGroupConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.OutputGroupConfiguration",
		"rtmp",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create an SRT output group configuration.
// Experimental.
func OutputGroupConfiguration_Srt(props *SrtOutputGroupProps) OutputGroupConfiguration {
	_init_.Initialize()

	if err := validateOutputGroupConfiguration_SrtParameters(props); err != nil {
		panic(err)
	}
	var returns OutputGroupConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.OutputGroupConfiguration",
		"srt",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create a UDP output group configuration.
// Experimental.
func OutputGroupConfiguration_Udp(props *UdpOutputGroupProps) OutputGroupConfiguration {
	_init_.Initialize()

	if err := validateOutputGroupConfiguration_UdpParameters(props); err != nil {
		panic(err)
	}
	var returns OutputGroupConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.OutputGroupConfiguration",
		"udp",
		[]interface{}{props},
		&returns,
	)

	return returns
}

