package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Configurations for sources.
//
// Example:
//   var stack Stack
//   var securityGroup ISecurityGroup
//   var subnet ISubnet
//   var role IRole
//
//
//   vpcInterface := awsmediaconnectalpha.VpcInterface_Define(&VpcInterfaceDefineProps{
//   	VpcInterfaceName: jsii.String("my-vpc-interface"),
//   	Role: role,
//   	SecurityGroups: []ISecurityGroup{
//   		securityGroup,
//   	},
//   	Subnet: subnet,
//   })
//
//   flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyFlow"), &FlowProps{
//   	Source: awsmediaconnectalpha.SourceConfiguration_Rist(&SourceRist{
//   		FlowSourceName: jsii.String("vpc-source"),
//   		Description: jsii.String("VPC-based source"),
//   		Port: jsii.Number(5000),
//   		MaxLatency: awscdk.Duration_Millis(jsii.Number(2000)),
//   		Network: awsmediaconnectalpha.NetworkConfiguration_Vpc(vpcInterface),
//   	}),
//   	VpcInterfaces: []VpcInterfaceConfig{
//   		vpcInterface,
//   	},
//   })
//
// Experimental.
type SourceConfiguration interface {
	// The name of this source, if one was set on the configuration.
	// Experimental.
	FlowSourceName() *string
}

// The jsii proxy struct for SourceConfiguration
type jsiiProxy_SourceConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_SourceConfiguration) FlowSourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowSourceName",
		&returns,
	)
	return returns
}


// Source option for CDI input.
// Experimental.
func SourceConfiguration_Cdi(input *SourceCdi) SourceConfiguration {
	_init_.Initialize()

	if err := validateSourceConfiguration_CdiParameters(input); err != nil {
		panic(err)
	}
	var returns SourceConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.SourceConfiguration",
		"cdi",
		[]interface{}{input},
		&returns,
	)

	return returns
}

// The entitlement that allows you to subscribe to content that comes from another AWS account.
//
// The entitlement is set by the
// content originator and the ARN is generated as part of the originator's flow.
// Experimental.
func SourceConfiguration_Entitlement(input *EntitlementSource) SourceConfiguration {
	_init_.Initialize()

	if err := validateSourceConfiguration_EntitlementParameters(input); err != nil {
		panic(err)
	}
	var returns SourceConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.SourceConfiguration",
		"entitlement",
		[]interface{}{input},
		&returns,
	)

	return returns
}

// The source configuration for cloud flows receiving a stream from a bridge.
// Experimental.
func SourceConfiguration_GatewayBridge(input *GatewayBridgeSource) SourceConfiguration {
	_init_.Initialize()

	if err := validateSourceConfiguration_GatewayBridgeParameters(input); err != nil {
		panic(err)
	}
	var returns SourceConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.SourceConfiguration",
		"gatewayBridge",
		[]interface{}{input},
		&returns,
	)

	return returns
}

// Source option for Jpeg-XS input.
// Experimental.
func SourceConfiguration_JpegXs(input *SourceJpegXs) SourceConfiguration {
	_init_.Initialize()

	if err := validateSourceConfiguration_JpegXsParameters(input); err != nil {
		panic(err)
	}
	var returns SourceConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.SourceConfiguration",
		"jpegXs",
		[]interface{}{input},
		&returns,
	)

	return returns
}

// Source option for NDI (SpeedHQ) input.
//
// The flow must be configured with `flowSize: FlowSize.LARGE` and
// `ndiConfig.ndiState = State.ENABLED` with at least one NDI discovery server.
// Experimental.
func SourceConfiguration_Ndi(input *SourceNdi) SourceConfiguration {
	_init_.Initialize()

	if err := validateSourceConfiguration_NdiParameters(input); err != nil {
		panic(err)
	}
	var returns SourceConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.SourceConfiguration",
		"ndi",
		[]interface{}{input},
		&returns,
	)

	return returns
}

// Source option for RIST input.
// Experimental.
func SourceConfiguration_Rist(input *SourceRist) SourceConfiguration {
	_init_.Initialize()

	if err := validateSourceConfiguration_RistParameters(input); err != nil {
		panic(err)
	}
	var returns SourceConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.SourceConfiguration",
		"rist",
		[]interface{}{input},
		&returns,
	)

	return returns
}

// The source configuration for flows receiving a stream from router.
// Experimental.
func SourceConfiguration_Router(input *RouterSource) SourceConfiguration {
	_init_.Initialize()

	if err := validateSourceConfiguration_RouterParameters(input); err != nil {
		panic(err)
	}
	var returns SourceConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.SourceConfiguration",
		"router",
		[]interface{}{input},
		&returns,
	)

	return returns
}

// Source option for RTP input.
// Experimental.
func SourceConfiguration_Rtp(input *SourceRtp) SourceConfiguration {
	_init_.Initialize()

	if err := validateSourceConfiguration_RtpParameters(input); err != nil {
		panic(err)
	}
	var returns SourceConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.SourceConfiguration",
		"rtp",
		[]interface{}{input},
		&returns,
	)

	return returns
}

// Source option for RTP-FEC input.
// Experimental.
func SourceConfiguration_RtpFec(input *SourceRtp) SourceConfiguration {
	_init_.Initialize()

	if err := validateSourceConfiguration_RtpFecParameters(input); err != nil {
		panic(err)
	}
	var returns SourceConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.SourceConfiguration",
		"rtpFec",
		[]interface{}{input},
		&returns,
	)

	return returns
}

// Source option for SRT Caller input.
// Experimental.
func SourceConfiguration_SrtCaller(input *SourceSrtCaller) SourceConfiguration {
	_init_.Initialize()

	if err := validateSourceConfiguration_SrtCallerParameters(input); err != nil {
		panic(err)
	}
	var returns SourceConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.SourceConfiguration",
		"srtCaller",
		[]interface{}{input},
		&returns,
	)

	return returns
}

// Source option for SRT Listener input.
// Experimental.
func SourceConfiguration_SrtListener(input *SourceSrtListener) SourceConfiguration {
	_init_.Initialize()

	if err := validateSourceConfiguration_SrtListenerParameters(input); err != nil {
		panic(err)
	}
	var returns SourceConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.SourceConfiguration",
		"srtListener",
		[]interface{}{input},
		&returns,
	)

	return returns
}

// Source option for Zixi Push input.
// Experimental.
func SourceConfiguration_ZixiPush(input *SourceZixiPush) SourceConfiguration {
	_init_.Initialize()

	if err := validateSourceConfiguration_ZixiPushParameters(input); err != nil {
		panic(err)
	}
	var returns SourceConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.SourceConfiguration",
		"zixiPush",
		[]interface{}{input},
		&returns,
	)

	return returns
}

