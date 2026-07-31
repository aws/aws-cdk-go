package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Protocol options available for Router Input configurations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//
//   routerInputProtocolOptions := mediaconnect_alpha.RouterInputProtocolOptions_RIST()
//
// Experimental.
type RouterInputProtocolOptions interface {
	// The protocol string value.
	// Experimental.
	Value() *string
	// Returns the string value.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RouterInputProtocolOptions
type jsiiProxy_RouterInputProtocolOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_RouterInputProtocolOptions) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Use a custom protocol value.
// Experimental.
func RouterInputProtocolOptions_Of(value *string) RouterInputProtocolOptions {
	_init_.Initialize()

	if err := validateRouterInputProtocolOptions_OfParameters(value); err != nil {
		panic(err)
	}
	var returns RouterInputProtocolOptions

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInputProtocolOptions",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func RouterInputProtocolOptions_RIST() RouterInputProtocolOptions {
	_init_.Initialize()
	var returns RouterInputProtocolOptions
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInputProtocolOptions",
		"RIST",
		&returns,
	)
	return returns
}

func RouterInputProtocolOptions_RTP() RouterInputProtocolOptions {
	_init_.Initialize()
	var returns RouterInputProtocolOptions
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInputProtocolOptions",
		"RTP",
		&returns,
	)
	return returns
}

func RouterInputProtocolOptions_SRT_CALLER() RouterInputProtocolOptions {
	_init_.Initialize()
	var returns RouterInputProtocolOptions
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInputProtocolOptions",
		"SRT_CALLER",
		&returns,
	)
	return returns
}

func RouterInputProtocolOptions_SRT_LISTENER() RouterInputProtocolOptions {
	_init_.Initialize()
	var returns RouterInputProtocolOptions
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInputProtocolOptions",
		"SRT_LISTENER",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RouterInputProtocolOptions) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

