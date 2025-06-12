package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An ADOT Lambda layer version that's specific to a lambda layer type and an architecture.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   fn := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_18_X(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromInline(jsii.String("exports.handler = function(event, ctx, cb) { return cb(null, \"hi\"); }")),
//   	AdotInstrumentation: &AdotInstrumentationConfig{
//   		LayerVersion: awscdk.AdotLayerVersion_FromJavaScriptSdkLayerVersion(awscdk.AdotLambdaLayerJavaScriptSdkVersion_LATEST()),
//   		ExecWrapper: awscdk.AdotLambdaExecWrapper_REGULAR_HANDLER,
//   	},
//   })
//
type AdotLayerVersion interface {
}

// The jsii proxy struct for AdotLayerVersion
type jsiiProxy_AdotLayerVersion struct {
	_ byte // padding
}

func NewAdotLayerVersion_Override(a AdotLayerVersion) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.AdotLayerVersion",
		nil, // no parameters
		a,
	)
}

// The ADOT Lambda layer for generic use cases.
func AdotLayerVersion_FromGenericLayerVersion(version AdotLambdaLayerGenericVersion) AdotLayerVersion {
	_init_.Initialize()

	if err := validateAdotLayerVersion_FromGenericLayerVersionParameters(version); err != nil {
		panic(err)
	}
	var returns AdotLayerVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.AdotLayerVersion",
		"fromGenericLayerVersion",
		[]interface{}{version},
		&returns,
	)

	return returns
}

// The ADOT Lambda layer for Java auto instrumentation.
func AdotLayerVersion_FromJavaAutoInstrumentationLayerVersion(version AdotLambdaLayerJavaAutoInstrumentationVersion) AdotLayerVersion {
	_init_.Initialize()

	if err := validateAdotLayerVersion_FromJavaAutoInstrumentationLayerVersionParameters(version); err != nil {
		panic(err)
	}
	var returns AdotLayerVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.AdotLayerVersion",
		"fromJavaAutoInstrumentationLayerVersion",
		[]interface{}{version},
		&returns,
	)

	return returns
}

// The ADOT Lambda layer for JavaScript SDK.
func AdotLayerVersion_FromJavaScriptSdkLayerVersion(version AdotLambdaLayerJavaScriptSdkVersion) AdotLayerVersion {
	_init_.Initialize()

	if err := validateAdotLayerVersion_FromJavaScriptSdkLayerVersionParameters(version); err != nil {
		panic(err)
	}
	var returns AdotLayerVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.AdotLayerVersion",
		"fromJavaScriptSdkLayerVersion",
		[]interface{}{version},
		&returns,
	)

	return returns
}

// The ADOT Lambda layer for Java SDK.
func AdotLayerVersion_FromJavaSdkLayerVersion(version AdotLambdaLayerJavaSdkVersion) AdotLayerVersion {
	_init_.Initialize()

	if err := validateAdotLayerVersion_FromJavaSdkLayerVersionParameters(version); err != nil {
		panic(err)
	}
	var returns AdotLayerVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.AdotLayerVersion",
		"fromJavaSdkLayerVersion",
		[]interface{}{version},
		&returns,
	)

	return returns
}

// The ADOT Lambda layer for Python SDK.
func AdotLayerVersion_FromPythonSdkLayerVersion(version AdotLambdaLayerPythonSdkVersion) AdotLayerVersion {
	_init_.Initialize()

	if err := validateAdotLayerVersion_FromPythonSdkLayerVersionParameters(version); err != nil {
		panic(err)
	}
	var returns AdotLayerVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.AdotLayerVersion",
		"fromPythonSdkLayerVersion",
		[]interface{}{version},
		&returns,
	)

	return returns
}

