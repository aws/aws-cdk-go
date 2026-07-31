package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Protocol options available for Router Output configurations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//
//   routerOutputProtocolOptions := mediaconnect_alpha.RouterOutputProtocolOptions_RIST()
//
// Experimental.
type RouterOutputProtocolOptions interface {
	// The protocol string value.
	// Experimental.
	Value() *string
	// Returns the string value.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RouterOutputProtocolOptions
type jsiiProxy_RouterOutputProtocolOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_RouterOutputProtocolOptions) Value() *string {
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
func RouterOutputProtocolOptions_Of(value *string) RouterOutputProtocolOptions {
	_init_.Initialize()

	if err := validateRouterOutputProtocolOptions_OfParameters(value); err != nil {
		panic(err)
	}
	var returns RouterOutputProtocolOptions

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.RouterOutputProtocolOptions",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func RouterOutputProtocolOptions_RIST() RouterOutputProtocolOptions {
	_init_.Initialize()
	var returns RouterOutputProtocolOptions
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.RouterOutputProtocolOptions",
		"RIST",
		&returns,
	)
	return returns
}

func RouterOutputProtocolOptions_RTP() RouterOutputProtocolOptions {
	_init_.Initialize()
	var returns RouterOutputProtocolOptions
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.RouterOutputProtocolOptions",
		"RTP",
		&returns,
	)
	return returns
}

func RouterOutputProtocolOptions_SRT_CALLER() RouterOutputProtocolOptions {
	_init_.Initialize()
	var returns RouterOutputProtocolOptions
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.RouterOutputProtocolOptions",
		"SRT_CALLER",
		&returns,
	)
	return returns
}

func RouterOutputProtocolOptions_SRT_LISTENER() RouterOutputProtocolOptions {
	_init_.Initialize()
	var returns RouterOutputProtocolOptions
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.RouterOutputProtocolOptions",
		"SRT_LISTENER",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RouterOutputProtocolOptions) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

