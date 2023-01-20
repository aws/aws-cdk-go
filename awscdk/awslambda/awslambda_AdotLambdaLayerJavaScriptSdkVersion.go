package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// The collection of versions of the ADOT Lambda Layer for JavaScript SDK.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   fn := lambda.NewFunction(this, jsii.String("MyFunction"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_18_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromInline(jsii.String("exports.handler = function(event, ctx, cb) { return cb(null, \"hi\"); }")),
//   	adotInstrumentation: &adotInstrumentationConfig{
//   		layerVersion: awscdk.AdotLayerVersion.fromJavaScriptSdkLayerVersion(awscdk.AdotLambdaLayerJavaScriptSdkVersion_LATEST()),
//   		execWrapper: awscdk.AdotLambdaExecWrapper_REGULAR_HANDLER,
//   	},
//   })
//
type AdotLambdaLayerJavaScriptSdkVersion interface {
	LayerVersion() *string
	// The ARN of the Lambda layer.
	LayerArn(scope constructs.IConstruct, architecture Architecture) *string
}

// The jsii proxy struct for AdotLambdaLayerJavaScriptSdkVersion
type jsiiProxy_AdotLambdaLayerJavaScriptSdkVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_AdotLambdaLayerJavaScriptSdkVersion) LayerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layerVersion",
		&returns,
	)
	return returns
}


func AdotLambdaLayerJavaScriptSdkVersion_LATEST() AdotLambdaLayerJavaScriptSdkVersion {
	_init_.Initialize()
	var returns AdotLambdaLayerJavaScriptSdkVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.AdotLambdaLayerJavaScriptSdkVersion",
		"LATEST",
		&returns,
	)
	return returns
}

func AdotLambdaLayerJavaScriptSdkVersion_V1_7_0() AdotLambdaLayerJavaScriptSdkVersion {
	_init_.Initialize()
	var returns AdotLambdaLayerJavaScriptSdkVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.AdotLambdaLayerJavaScriptSdkVersion",
		"V1_7_0",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AdotLambdaLayerJavaScriptSdkVersion) LayerArn(scope constructs.IConstruct, architecture Architecture) *string {
	if err := a.validateLayerArnParameters(scope, architecture); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"layerArn",
		[]interface{}{scope, architecture},
		&returns,
	)

	return returns
}

