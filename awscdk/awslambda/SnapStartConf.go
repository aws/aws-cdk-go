package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Example:
//   fn := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("handler.zip"))),
//   	Runtime: lambda.Runtime_JAVA_11(),
//   	Handler: jsii.String("example.Handler::handleRequest"),
//   	SnapStart: lambda.SnapStartConf_ON_PUBLISHED_VERSIONS(),
//   })
//
//   version := fn.currentVersion
//
type SnapStartConf interface {
}

// The jsii proxy struct for SnapStartConf
type jsiiProxy_SnapStartConf struct {
	_ byte // padding
}

func NewSnapStartConf_Override(s SnapStartConf) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.SnapStartConf",
		nil, // no parameters
		s,
	)
}

func SnapStartConf_ON_PUBLISHED_VERSIONS() SnapStartConf {
	_init_.Initialize()
	var returns SnapStartConf
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.SnapStartConf",
		"ON_PUBLISHED_VERSIONS",
		&returns,
	)
	return returns
}

