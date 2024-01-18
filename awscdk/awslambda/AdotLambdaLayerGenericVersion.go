package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// The collection of versions of the ADOT Lambda Layer for generic purpose.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   adotLambdaLayerGenericVersion := awscdk.Aws_lambda.AdotLambdaLayerGenericVersion_LATEST()
//
type AdotLambdaLayerGenericVersion interface {
	LayerVersion() *string
	// The ARN of the Lambda layer.
	LayerArn(scope constructs.IConstruct, architecture Architecture) *string
}

// The jsii proxy struct for AdotLambdaLayerGenericVersion
type jsiiProxy_AdotLambdaLayerGenericVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_AdotLambdaLayerGenericVersion) LayerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layerVersion",
		&returns,
	)
	return returns
}


func AdotLambdaLayerGenericVersion_LATEST() AdotLambdaLayerGenericVersion {
	_init_.Initialize()
	var returns AdotLambdaLayerGenericVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.AdotLambdaLayerGenericVersion",
		"LATEST",
		&returns,
	)
	return returns
}

func AdotLambdaLayerGenericVersion_V0_62_1() AdotLambdaLayerGenericVersion {
	_init_.Initialize()
	var returns AdotLambdaLayerGenericVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.AdotLambdaLayerGenericVersion",
		"V0_62_1",
		&returns,
	)
	return returns
}

func AdotLambdaLayerGenericVersion_V0_82_0() AdotLambdaLayerGenericVersion {
	_init_.Initialize()
	var returns AdotLambdaLayerGenericVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.AdotLambdaLayerGenericVersion",
		"V0_82_0",
		&returns,
	)
	return returns
}

func AdotLambdaLayerGenericVersion_V0_84_0() AdotLambdaLayerGenericVersion {
	_init_.Initialize()
	var returns AdotLambdaLayerGenericVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.AdotLambdaLayerGenericVersion",
		"V0_84_0",
		&returns,
	)
	return returns
}

func AdotLambdaLayerGenericVersion_V0_88_0() AdotLambdaLayerGenericVersion {
	_init_.Initialize()
	var returns AdotLambdaLayerGenericVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.AdotLambdaLayerGenericVersion",
		"V0_88_0",
		&returns,
	)
	return returns
}

func AdotLambdaLayerGenericVersion_V0_90_1() AdotLambdaLayerGenericVersion {
	_init_.Initialize()
	var returns AdotLambdaLayerGenericVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.AdotLambdaLayerGenericVersion",
		"V0_90_1",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AdotLambdaLayerGenericVersion) LayerArn(scope constructs.IConstruct, architecture Architecture) *string {
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

