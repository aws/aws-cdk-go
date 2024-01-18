package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// The collection of versions of the ADOT Lambda Layer for Java auto-instrumentation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   adotLambdaLayerJavaAutoInstrumentationVersion := awscdk.Aws_lambda.AdotLambdaLayerJavaAutoInstrumentationVersion_LATEST()
//
type AdotLambdaLayerJavaAutoInstrumentationVersion interface {
	LayerVersion() *string
	// The ARN of the Lambda layer.
	LayerArn(scope constructs.IConstruct, architecture Architecture) *string
}

// The jsii proxy struct for AdotLambdaLayerJavaAutoInstrumentationVersion
type jsiiProxy_AdotLambdaLayerJavaAutoInstrumentationVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_AdotLambdaLayerJavaAutoInstrumentationVersion) LayerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layerVersion",
		&returns,
	)
	return returns
}


func AdotLambdaLayerJavaAutoInstrumentationVersion_LATEST() AdotLambdaLayerJavaAutoInstrumentationVersion {
	_init_.Initialize()
	var returns AdotLambdaLayerJavaAutoInstrumentationVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.AdotLambdaLayerJavaAutoInstrumentationVersion",
		"LATEST",
		&returns,
	)
	return returns
}

func AdotLambdaLayerJavaAutoInstrumentationVersion_V1_19_2() AdotLambdaLayerJavaAutoInstrumentationVersion {
	_init_.Initialize()
	var returns AdotLambdaLayerJavaAutoInstrumentationVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.AdotLambdaLayerJavaAutoInstrumentationVersion",
		"V1_19_2",
		&returns,
	)
	return returns
}

func AdotLambdaLayerJavaAutoInstrumentationVersion_V1_28_1() AdotLambdaLayerJavaAutoInstrumentationVersion {
	_init_.Initialize()
	var returns AdotLambdaLayerJavaAutoInstrumentationVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.AdotLambdaLayerJavaAutoInstrumentationVersion",
		"V1_28_1",
		&returns,
	)
	return returns
}

func AdotLambdaLayerJavaAutoInstrumentationVersion_V1_30_0() AdotLambdaLayerJavaAutoInstrumentationVersion {
	_init_.Initialize()
	var returns AdotLambdaLayerJavaAutoInstrumentationVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.AdotLambdaLayerJavaAutoInstrumentationVersion",
		"V1_30_0",
		&returns,
	)
	return returns
}

func AdotLambdaLayerJavaAutoInstrumentationVersion_V1_31_0() AdotLambdaLayerJavaAutoInstrumentationVersion {
	_init_.Initialize()
	var returns AdotLambdaLayerJavaAutoInstrumentationVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.AdotLambdaLayerJavaAutoInstrumentationVersion",
		"V1_31_0",
		&returns,
	)
	return returns
}

func AdotLambdaLayerJavaAutoInstrumentationVersion_V1_32_0() AdotLambdaLayerJavaAutoInstrumentationVersion {
	_init_.Initialize()
	var returns AdotLambdaLayerJavaAutoInstrumentationVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.AdotLambdaLayerJavaAutoInstrumentationVersion",
		"V1_32_0",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AdotLambdaLayerJavaAutoInstrumentationVersion) LayerArn(scope constructs.IConstruct, architecture Architecture) *string {
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

