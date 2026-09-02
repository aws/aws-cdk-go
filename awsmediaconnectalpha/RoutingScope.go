package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Routing scope for the Router Input.
//
// Example:
//   var stack Stack
//   var mediaLiveChannel IChannel
//
//
//   input := awsmediaconnectalpha.NewRouterInput(stack, jsii.String("ChannelInput"), &RouterInputProps{
//   	RouterInputName: jsii.String("channel-input"),
//   	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(20)),
//   	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
//   	Tier: awsmediaconnectalpha.RouterInputTier_INPUT_50(),
//   	Configuration: awsmediaconnectalpha.RouterInputConfiguration_MediaLiveChannel(&MediaLiveChannelConfigurationProps{
//   		Channel: mediaLiveChannel,
//   		OutputName: jsii.String("router-ts"),
//   		Pipeline: awsmediaconnectalpha.MediaLivePipeline_PIPELINE_0,
//   	}),
//   })
//
// Experimental.
type RoutingScope interface {
	// The routing scope string value.
	// Experimental.
	Value() *string
	// Returns the string value.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RoutingScope
type jsiiProxy_RoutingScope struct {
	_ byte // padding
}

func (j *jsiiProxy_RoutingScope) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Use a custom routing scope value.
// Experimental.
func RoutingScope_Of(value *string) RoutingScope {
	_init_.Initialize()

	if err := validateRoutingScope_OfParameters(value); err != nil {
		panic(err)
	}
	var returns RoutingScope

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.RoutingScope",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func RoutingScope_GLOBAL() RoutingScope {
	_init_.Initialize()
	var returns RoutingScope
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.RoutingScope",
		"GLOBAL",
		&returns,
	)
	return returns
}

func RoutingScope_REGIONAL() RoutingScope {
	_init_.Initialize()
	var returns RoutingScope
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.RoutingScope",
		"REGIONAL",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RoutingScope) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

