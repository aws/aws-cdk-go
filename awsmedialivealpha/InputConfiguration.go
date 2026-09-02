package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Defines the input configuration for a MediaLive Input.
//
// Use the static factory methods to create the appropriate configuration for your input type.
//
// Example:
//   var stack Stack
//   var passphrase ISecret
//
//
//   sg := medialive.NewInputSecurityGroup(stack, jsii.String("SrtSg"), &InputSecurityGroupProps{
//   	AllowlistRules: []*string{
//   		jsii.String("203.0.113.0/24"),
//   	},
//   })
//
//   medialive.NewInput(stack, jsii.String("SrtListenerInput"), &InputProps{
//   	InputName: jsii.String("srt-listener"),
//   	Input: medialive.InputConfiguration_SrtListener(&SrtListenerInputProps{
//   		InputSecurityGroups: []IInputSecurityGroupRef{
//   			sg,
//   		},
//   		MinimumLatency: awscdk.Duration_Millis(jsii.Number(500)),
//   		StreamId: jsii.String("my-stream-id"),
//   		Decryption: &SrtDecryptionProps{
//   			Algorithm: medialive.SrtDecryptionAlgorithm_AES256(),
//   			PassphraseSecret: passphrase,
//   		},
//   	}),
//   })
//
// Experimental.
type InputConfiguration interface {
}

// The jsii proxy struct for InputConfiguration
type jsiiProxy_InputConfiguration struct {
	_ byte // padding
}

// Create a CDI (uncompressed) input delivered into a VPC.
//
// MediaLive creates network interfaces in the supplied subnets and hands back the CDI push
// endpoints. The required EC2 permissions are granted to the role automatically.
// Experimental.
func InputConfiguration_Cdi(props *CdiInputProps) InputConfiguration {
	_init_.Initialize()

	if err := validateInputConfiguration_CdiParameters(props); err != nil {
		panic(err)
	}
	var returns InputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.InputConfiguration",
		"cdi",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create an Elemental Link input from one or two registered Link device IDs.
// Experimental.
func InputConfiguration_InputDevice(props *InputDeviceInputProps) InputConfiguration {
	_init_.Initialize()

	if err := validateInputConfiguration_InputDeviceParameters(props); err != nil {
		panic(err)
	}
	var returns InputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.InputConfiguration",
		"inputDevice",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create a MediaConnect flow input configuration.
//
// Example:
//   var role IRole
//   var flow IFlowRef
//
//
//   awsmedialivealpha.InputConfiguration_MediaConnect(&MediaConnectInputProps{
//   	Flows: []IFlowRef{
//   		flow,
//   	},
//   	Role: Role,
//   })
//
// Experimental.
func InputConfiguration_MediaConnect(props *MediaConnectInputProps) InputConfiguration {
	_init_.Initialize()

	if err := validateInputConfiguration_MediaConnectParameters(props); err != nil {
		panic(err)
	}
	var returns InputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.InputConfiguration",
		"mediaConnect",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create a MediaConnect router input configuration.
//
// Example:
//   // Redundant pipelines with explicit AZs
//   awsmedialivealpha.InputConfiguration_MediaConnectRouter(&MediaConnectRouterInputProps{
//   	AvailabilityZones: []*string{
//   		jsii.String("us-east-1a"),
//   		jsii.String("us-east-1b"),
//   	},
//   })
//
// Experimental.
func InputConfiguration_MediaConnectRouter(props *MediaConnectRouterInputProps) InputConfiguration {
	_init_.Initialize()

	if err := validateInputConfiguration_MediaConnectRouterParameters(props); err != nil {
		panic(err)
	}
	var returns InputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.InputConfiguration",
		"mediaConnectRouter",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create an MP4 file pull input.
// Experimental.
func InputConfiguration_Mp4File(sources *[]InputSource) InputConfiguration {
	_init_.Initialize()

	if err := validateInputConfiguration_Mp4FileParameters(sources); err != nil {
		panic(err)
	}
	var returns InputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.InputConfiguration",
		"mp4File",
		[]interface{}{sources},
		&returns,
	)

	return returns
}

// Create a multicast input.
//
// Requires `anywhereSettings` on the channel.
// Experimental.
func InputConfiguration_Multicast(props *MulticastInputProps) InputConfiguration {
	_init_.Initialize()

	if err := validateInputConfiguration_MulticastParameters(props); err != nil {
		panic(err)
	}
	var returns InputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.InputConfiguration",
		"multicast",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create an RTMP pull input.
// Experimental.
func InputConfiguration_RtmpPull(sources *[]InputSource) InputConfiguration {
	_init_.Initialize()

	if err := validateInputConfiguration_RtmpPullParameters(sources); err != nil {
		panic(err)
	}
	var returns InputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.InputConfiguration",
		"rtmpPull",
		[]interface{}{sources},
		&returns,
	)

	return returns
}

// Create an RTMP push input.
//
// Example:
//   var securityGroup InputSecurityGroup
//
//
//   awsmedialivealpha.InputConfiguration_RtmpPush(&PushInputProps{
//   	InputSecurityGroups: []IInputSecurityGroupRef{
//   		securityGroup,
//   	},
//   })
//
// Experimental.
func InputConfiguration_RtmpPush(props *PushInputProps) InputConfiguration {
	_init_.Initialize()

	if err := validateInputConfiguration_RtmpPushParameters(props); err != nil {
		panic(err)
	}
	var returns InputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.InputConfiguration",
		"rtmpPush",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create an RTP push input.
// Experimental.
func InputConfiguration_RtpPush(props *PushInputProps) InputConfiguration {
	_init_.Initialize()

	if err := validateInputConfiguration_RtpPushParameters(props); err != nil {
		panic(err)
	}
	var returns InputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.InputConfiguration",
		"rtpPush",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create an SDI input configuration.
// Experimental.
func InputConfiguration_Sdi(sources *[]ISdiSource) InputConfiguration {
	_init_.Initialize()

	if err := validateInputConfiguration_SdiParameters(sources); err != nil {
		panic(err)
	}
	var returns InputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.InputConfiguration",
		"sdi",
		[]interface{}{sources},
		&returns,
	)

	return returns
}

// Create a SMPTE 2110 receiver group input.
//
// Requires `anywhereSettings` on the channel.
// Experimental.
func InputConfiguration_Smpte2110ReceiverGroup(props *Smpte2110InputProps) InputConfiguration {
	_init_.Initialize()

	if err := validateInputConfiguration_Smpte2110ReceiverGroupParameters(props); err != nil {
		panic(err)
	}
	var returns InputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.InputConfiguration",
		"smpte2110ReceiverGroup",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create an SRT caller input configuration.
//
// Example:
//   awsmedialivealpha.InputConfiguration_SrtCaller([]SrtCallerSourceProps{
//   	&SrtCallerSourceProps{
//   		SrtListenerAddress: jsii.String("10.0.0.1"),
//   		SrtListenerPort: jsii.Number(5000),
//   	},
//   })
//
// Experimental.
func InputConfiguration_SrtCaller(sources *[]*SrtCallerSourceProps) InputConfiguration {
	_init_.Initialize()

	if err := validateInputConfiguration_SrtCallerParameters(sources); err != nil {
		panic(err)
	}
	var returns InputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.InputConfiguration",
		"srtCaller",
		[]interface{}{sources},
		&returns,
	)

	return returns
}

// Create an SRT listener input configuration.
// Experimental.
func InputConfiguration_SrtListener(props *SrtListenerInputProps) InputConfiguration {
	_init_.Initialize()

	if err := validateInputConfiguration_SrtListenerParameters(props); err != nil {
		panic(err)
	}
	var returns InputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.InputConfiguration",
		"srtListener",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create a TS file pull input.
// Experimental.
func InputConfiguration_TsFile(sources *[]InputSource) InputConfiguration {
	_init_.Initialize()

	if err := validateInputConfiguration_TsFileParameters(sources); err != nil {
		panic(err)
	}
	var returns InputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.InputConfiguration",
		"tsFile",
		[]interface{}{sources},
		&returns,
	)

	return returns
}

// Create a UDP push input.
// Experimental.
func InputConfiguration_UdpPush(props *PushInputProps) InputConfiguration {
	_init_.Initialize()

	if err := validateInputConfiguration_UdpPushParameters(props); err != nil {
		panic(err)
	}
	var returns InputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.InputConfiguration",
		"udpPush",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create a URL pull input for HLS or TS streams.
//
// Example:
//   awsmedialivealpha.InputConfiguration_UrlPull([]InputSource{
//   	awsmedialivealpha.InputSource_Url(jsii.String("https://example.com/stream.m3u8")),
//   })
//
// Experimental.
func InputConfiguration_UrlPull(sources *[]InputSource) InputConfiguration {
	_init_.Initialize()

	if err := validateInputConfiguration_UrlPullParameters(sources); err != nil {
		panic(err)
	}
	var returns InputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.InputConfiguration",
		"urlPull",
		[]interface{}{sources},
		&returns,
	)

	return returns
}

