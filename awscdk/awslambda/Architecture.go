package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Architectures supported by AWS Lambda.
//
// Example:
//   lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_18_X(),
//   	Handler: jsii.String("index.handler"),
//   	Architecture: lambda.Architecture_ARM_64(),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   	InsightsVersion: lambda.LambdaInsightsVersion_VERSION_1_0_119_0(),
//   })
//
type Architecture interface {
	// The platform to use for this architecture when building with Docker.
	DockerPlatform() *string
	// The name of the architecture as recognized by the AWS Lambda service APIs.
	Name() *string
	// Returns a string representation of the architecture using the name.
	ToString() *string
}

// The jsii proxy struct for Architecture
type jsiiProxy_Architecture struct {
	_ byte // padding
}

func (j *jsiiProxy_Architecture) DockerPlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerPlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Architecture) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Used to specify a custom architecture name.
//
// Use this if the architecture name is not yet supported by the CDK.
func Architecture_Custom(name *string, dockerPlatform *string) Architecture {
	_init_.Initialize()

	if err := validateArchitecture_CustomParameters(name); err != nil {
		panic(err)
	}
	var returns Architecture

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Architecture",
		"custom",
		[]interface{}{name, dockerPlatform},
		&returns,
	)

	return returns
}

func Architecture_ARM_64() Architecture {
	_init_.Initialize()
	var returns Architecture
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Architecture",
		"ARM_64",
		&returns,
	)
	return returns
}

func Architecture_X86_64() Architecture {
	_init_.Initialize()
	var returns Architecture
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Architecture",
		"X86_64",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_Architecture) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

