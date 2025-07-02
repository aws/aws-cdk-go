package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Service connect app protocol.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appProtocol := awscdk.Aws_ecs.AppProtocol_Grpc()
//
type AppProtocol interface {
	// Custom value.
	Value() *string
}

// The jsii proxy struct for AppProtocol
type jsiiProxy_AppProtocol struct {
	_ byte // padding
}

func (j *jsiiProxy_AppProtocol) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NewAppProtocol(value *string) AppProtocol {
	_init_.Initialize()

	if err := validateNewAppProtocolParameters(value); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppProtocol{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.AppProtocol",
		[]interface{}{value},
		&j,
	)

	return &j
}

func NewAppProtocol_Override(a AppProtocol, value *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.AppProtocol",
		[]interface{}{value},
		a,
	)
}

func AppProtocol_Grpc() AppProtocol {
	_init_.Initialize()
	var returns AppProtocol
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.AppProtocol",
		"grpc",
		&returns,
	)
	return returns
}

func AppProtocol_SetGrpc(val AppProtocol) {
	_init_.Initialize()
	if err := validateAppProtocol_SetGrpcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.StaticSet(
		"aws-cdk-lib.aws_ecs.AppProtocol",
		"grpc",
		val,
	)
}

func AppProtocol_Http() AppProtocol {
	_init_.Initialize()
	var returns AppProtocol
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.AppProtocol",
		"http",
		&returns,
	)
	return returns
}

func AppProtocol_SetHttp(val AppProtocol) {
	_init_.Initialize()
	if err := validateAppProtocol_SetHttpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.StaticSet(
		"aws-cdk-lib.aws_ecs.AppProtocol",
		"http",
		val,
	)
}

func AppProtocol_Http2() AppProtocol {
	_init_.Initialize()
	var returns AppProtocol
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.AppProtocol",
		"http2",
		&returns,
	)
	return returns
}

func AppProtocol_SetHttp2(val AppProtocol) {
	_init_.Initialize()
	if err := validateAppProtocol_SetHttp2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.StaticSet(
		"aws-cdk-lib.aws_ecs.AppProtocol",
		"http2",
		val,
	)
}

