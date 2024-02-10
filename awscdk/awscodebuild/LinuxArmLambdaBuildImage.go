package awscodebuild

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A CodeBuild image running aarch64 Lambda.
//
// This class has a bunch of public constants that represent the CodeBuild aarch64 Lambda images.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   linuxArmLambdaBuildImage := awscdk.Aws_codebuild.LinuxArmLambdaBuildImage_AMAZON_LINUX_2_CORRETTO_11()
//
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-available.html
//
type LinuxArmLambdaBuildImage interface {
	IBuildImage
	// The default `ComputeType` to use with this image, if one was not specified in `BuildEnvironment#computeType` explicitly.
	DefaultComputeType() ComputeType
	// The Docker image identifier that the build environment uses.
	ImageId() *string
	// The type of build environment.
	Type() *string
	// Make a buildspec to run the indicated script.
	RunScriptBuildspec(entrypoint *string) BuildSpec
	// Allows the image a chance to validate whether the passed configuration is correct.
	Validate(buildEnvironment *BuildEnvironment) *[]*string
}

// The jsii proxy struct for LinuxArmLambdaBuildImage
type jsiiProxy_LinuxArmLambdaBuildImage struct {
	jsiiProxy_IBuildImage
}

func (j *jsiiProxy_LinuxArmLambdaBuildImage) DefaultComputeType() ComputeType {
	var returns ComputeType
	_jsii_.Get(
		j,
		"defaultComputeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxArmLambdaBuildImage) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxArmLambdaBuildImage) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func LinuxArmLambdaBuildImage_AMAZON_LINUX_2_CORRETTO_11() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxArmLambdaBuildImage",
		"AMAZON_LINUX_2_CORRETTO_11",
		&returns,
	)
	return returns
}

func LinuxArmLambdaBuildImage_AMAZON_LINUX_2_CORRETTO_17() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxArmLambdaBuildImage",
		"AMAZON_LINUX_2_CORRETTO_17",
		&returns,
	)
	return returns
}

func LinuxArmLambdaBuildImage_AMAZON_LINUX_2_DOTNET_6() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxArmLambdaBuildImage",
		"AMAZON_LINUX_2_DOTNET_6",
		&returns,
	)
	return returns
}

func LinuxArmLambdaBuildImage_AMAZON_LINUX_2_GO_1_21() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxArmLambdaBuildImage",
		"AMAZON_LINUX_2_GO_1_21",
		&returns,
	)
	return returns
}

func LinuxArmLambdaBuildImage_AMAZON_LINUX_2_NODE_18() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxArmLambdaBuildImage",
		"AMAZON_LINUX_2_NODE_18",
		&returns,
	)
	return returns
}

func LinuxArmLambdaBuildImage_AMAZON_LINUX_2_PYTHON_3_11() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxArmLambdaBuildImage",
		"AMAZON_LINUX_2_PYTHON_3_11",
		&returns,
	)
	return returns
}

func LinuxArmLambdaBuildImage_AMAZON_LINUX_2_RUBY_3_2() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxArmLambdaBuildImage",
		"AMAZON_LINUX_2_RUBY_3_2",
		&returns,
	)
	return returns
}

func LinuxArmLambdaBuildImage_AMAZON_LINUX_2023_CORRETTO_21() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxArmLambdaBuildImage",
		"AMAZON_LINUX_2023_CORRETTO_21",
		&returns,
	)
	return returns
}

func LinuxArmLambdaBuildImage_AMAZON_LINUX_2023_NODE_20() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxArmLambdaBuildImage",
		"AMAZON_LINUX_2023_NODE_20",
		&returns,
	)
	return returns
}

func LinuxArmLambdaBuildImage_AMAZON_LINUX_2023_PYTHON_3_12() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxArmLambdaBuildImage",
		"AMAZON_LINUX_2023_PYTHON_3_12",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LinuxArmLambdaBuildImage) RunScriptBuildspec(entrypoint *string) BuildSpec {
	if err := l.validateRunScriptBuildspecParameters(entrypoint); err != nil {
		panic(err)
	}
	var returns BuildSpec

	_jsii_.Invoke(
		l,
		"runScriptBuildspec",
		[]interface{}{entrypoint},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxArmLambdaBuildImage) Validate(buildEnvironment *BuildEnvironment) *[]*string {
	if err := l.validateValidateParameters(buildEnvironment); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"validate",
		[]interface{}{buildEnvironment},
		&returns,
	)

	return returns
}

