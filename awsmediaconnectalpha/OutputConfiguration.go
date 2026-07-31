package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Configuration options to define a FlowOutput by protocol.
//
// Example:
//   var stack Stack
//   var flow Flow
//   var role IRole
//   var secret ISecret
//
//
//   // SRT Caller output with encryption
//   output := awsmediaconnectalpha.NewFlowOutput(stack, jsii.String("EncryptedOutput"), &FlowOutputProps{
//   	Flow: flow,
//   	Description: jsii.String("Encrypted SRT output"),
//   	Output: awsmediaconnectalpha.OutputConfiguration_SrtCaller(&SrtCallerOutputConfig{
//   		Destination: jsii.String("203.0.113.100"),
//   		Port: jsii.Number(7000),
//   		Encryption: &SrtPasswordEncryption{
//   			Role: *Role,
//   			Secret: *Secret,
//   		},
//   	}),
//   })
//
// Experimental.
type OutputConfiguration interface {
}

// The jsii proxy struct for OutputConfiguration
type jsiiProxy_OutputConfiguration struct {
	_ byte // padding
}

// Option for NDI configuration.
// Experimental.
func OutputConfiguration_Ndi(input *NdiOutputConfig) OutputConfiguration {
	_init_.Initialize()

	if err := validateOutputConfiguration_NdiParameters(input); err != nil {
		panic(err)
	}
	var returns OutputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.OutputConfiguration",
		"ndi",
		[]interface{}{input},
		&returns,
	)

	return returns
}

// Option for RIST configuration.
// Experimental.
func OutputConfiguration_Rist(input *RistOutputConfig) OutputConfiguration {
	_init_.Initialize()

	if err := validateOutputConfiguration_RistParameters(input); err != nil {
		panic(err)
	}
	var returns OutputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.OutputConfiguration",
		"rist",
		[]interface{}{input},
		&returns,
	)

	return returns
}

// The source configuration for flows receiving a stream from router.
// Experimental.
func OutputConfiguration_Router(input *RouterTransitConfig) OutputConfiguration {
	_init_.Initialize()

	if err := validateOutputConfiguration_RouterParameters(input); err != nil {
		panic(err)
	}
	var returns OutputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.OutputConfiguration",
		"router",
		[]interface{}{input},
		&returns,
	)

	return returns
}

// Option for RTP configuration.
// Experimental.
func OutputConfiguration_Rtp(input *RtpOutputConfig) OutputConfiguration {
	_init_.Initialize()

	if err := validateOutputConfiguration_RtpParameters(input); err != nil {
		panic(err)
	}
	var returns OutputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.OutputConfiguration",
		"rtp",
		[]interface{}{input},
		&returns,
	)

	return returns
}

// Option for RTP-FEC configuration.
// Experimental.
func OutputConfiguration_RtpFec(input *RtpFecOutputConfig) OutputConfiguration {
	_init_.Initialize()

	if err := validateOutputConfiguration_RtpFecParameters(input); err != nil {
		panic(err)
	}
	var returns OutputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.OutputConfiguration",
		"rtpFec",
		[]interface{}{input},
		&returns,
	)

	return returns
}

// Option for SRT Caller configuration.
// Experimental.
func OutputConfiguration_SrtCaller(input *SrtCallerOutputConfig) OutputConfiguration {
	_init_.Initialize()

	if err := validateOutputConfiguration_SrtCallerParameters(input); err != nil {
		panic(err)
	}
	var returns OutputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.OutputConfiguration",
		"srtCaller",
		[]interface{}{input},
		&returns,
	)

	return returns
}

// Option for SRT Listener configuration.
// Experimental.
func OutputConfiguration_SrtListener(input *SrtListenerOutputConfig) OutputConfiguration {
	_init_.Initialize()

	if err := validateOutputConfiguration_SrtListenerParameters(input); err != nil {
		panic(err)
	}
	var returns OutputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.OutputConfiguration",
		"srtListener",
		[]interface{}{input},
		&returns,
	)

	return returns
}

// Option for Zixi Pull configuration.
// Experimental.
func OutputConfiguration_ZixiPull(input *ZixiPullOutputConfig) OutputConfiguration {
	_init_.Initialize()

	if err := validateOutputConfiguration_ZixiPullParameters(input); err != nil {
		panic(err)
	}
	var returns OutputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.OutputConfiguration",
		"zixiPull",
		[]interface{}{input},
		&returns,
	)

	return returns
}

// Option for Zixi Push configuration.
// Experimental.
func OutputConfiguration_ZixiPush(input *ZixiPushOutputConfig) OutputConfiguration {
	_init_.Initialize()

	if err := validateOutputConfiguration_ZixiPushParameters(input); err != nil {
		panic(err)
	}
	var returns OutputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.OutputConfiguration",
		"zixiPush",
		[]interface{}{input},
		&returns,
	)

	return returns
}

