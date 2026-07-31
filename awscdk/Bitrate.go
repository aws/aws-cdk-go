package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a bitrate value.
//
// The amount can be specified either as a literal value (e.g: `10`) which
// cannot be negative, or as an unresolved number token.
//
// When the amount is passed as a token, unit conversion is not possible.
//
// Example:
//   var stack Stack
//   var networkInterface RouterNetworkInterface
//
//
//   output := awsmediaconnectalpha.NewRouterOutput(stack, jsii.String("SrtOutput"), &RouterOutputProps{
//   	RouterOutputName: jsii.String("srt-output"),
//   	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(10)),
//   	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
//   	// tier defaults to RouterOutputTier.OUTPUT_20 (lowest cost)
//   	Configuration: awsmediaconnectalpha.RouterOutputConfiguration_Standard(&StandardOutputConfigurationProps{
//   		Protocol: awsmediaconnectalpha.RouterOutputProtocol_SrtListener(&SrtListenerOutputProtocolProps{
//   			Port: jsii.Number(9001),
//   			MinimumLatency: awscdk.Duration_Millis(jsii.Number(200)),
//   		}),
//   		NetworkInterface: networkInterface,
//   	}),
//   })
//
type Bitrate interface {
	// Checks if bitrate is a token or a resolvable object.
	IsUnresolved() *bool
	// Return the total number of bits per second.
	ToBps() *float64
	// Return the total number of gigabits per second.
	ToGbps() *float64
	// Return the total number of kilobits per second.
	ToKbps() *float64
	// Return the total number of megabits per second.
	ToMbps() *float64
}

// The jsii proxy struct for Bitrate
type jsiiProxy_Bitrate struct {
	_ byte // padding
}

// Create a Bitrate representing an amount of bits per second.
func Bitrate_Bps(amount *float64) Bitrate {
	_init_.Initialize()

	if err := validateBitrate_BpsParameters(amount); err != nil {
		panic(err)
	}
	var returns Bitrate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Bitrate",
		"bps",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Create a Bitrate representing an amount of gigabits per second.
func Bitrate_Gbps(amount *float64) Bitrate {
	_init_.Initialize()

	if err := validateBitrate_GbpsParameters(amount); err != nil {
		panic(err)
	}
	var returns Bitrate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Bitrate",
		"gbps",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Create a Bitrate representing an amount of kilobits per second.
func Bitrate_Kbps(amount *float64) Bitrate {
	_init_.Initialize()

	if err := validateBitrate_KbpsParameters(amount); err != nil {
		panic(err)
	}
	var returns Bitrate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Bitrate",
		"kbps",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Create a Bitrate representing an amount of megabits per second.
func Bitrate_Mbps(amount *float64) Bitrate {
	_init_.Initialize()

	if err := validateBitrate_MbpsParameters(amount); err != nil {
		panic(err)
	}
	var returns Bitrate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Bitrate",
		"mbps",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bitrate) IsUnresolved() *bool {
	var returns *bool

	_jsii_.Invoke(
		b,
		"isUnresolved",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bitrate) ToBps() *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"toBps",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bitrate) ToGbps() *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"toGbps",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bitrate) ToKbps() *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"toKbps",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bitrate) ToMbps() *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"toMbps",
		nil, // no parameters
		&returns,
	)

	return returns
}

