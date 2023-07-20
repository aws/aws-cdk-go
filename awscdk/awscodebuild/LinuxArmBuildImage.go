package awscodebuild

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// A CodeBuild image running aarch64 Linux.
//
// This class has a bunch of public constants that represent the CodeBuild ARM images.
//
// You can also specify a custom image using the static method:
//
// - LinuxBuildImage.fromEcrRepository(repo[, tag])
// - LinuxBuildImage.fromDockerRegistry(image[, { secretsManagerCredentials }])
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   linuxArmBuildImage := awscdk.Aws_codebuild.LinuxArmBuildImage_FromCodeBuildImageId(jsii.String("id"))
//
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-available.html
//
type LinuxArmBuildImage interface {
	IBuildImage
	// The default `ComputeType` to use with this image, if one was not specified in `BuildEnvironment#computeType` explicitly.
	DefaultComputeType() ComputeType
	// The Docker image identifier that the build environment uses.
	ImageId() *string
	// The type of principal that CodeBuild will use to pull this build Docker image.
	ImagePullPrincipalType() ImagePullPrincipalType
	// An optional ECR repository that the image is hosted in.
	Repository() awsecr.IRepository
	// The secretsManagerCredentials for access to a private registry.
	SecretsManagerCredentials() awssecretsmanager.ISecret
	// The type of build environment.
	Type() *string
	// Make a buildspec to run the indicated script.
	RunScriptBuildspec(entrypoint *string) BuildSpec
	// Validates by checking the BuildEnvironment computeType as aarch64 images only support ComputeType.SMALL and ComputeType.LARGE.
	Validate(buildEnvironment *BuildEnvironment) *[]*string
}

// The jsii proxy struct for LinuxArmBuildImage
type jsiiProxy_LinuxArmBuildImage struct {
	jsiiProxy_IBuildImage
}

func (j *jsiiProxy_LinuxArmBuildImage) DefaultComputeType() ComputeType {
	var returns ComputeType
	_jsii_.Get(
		j,
		"defaultComputeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxArmBuildImage) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxArmBuildImage) ImagePullPrincipalType() ImagePullPrincipalType {
	var returns ImagePullPrincipalType
	_jsii_.Get(
		j,
		"imagePullPrincipalType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxArmBuildImage) Repository() awsecr.IRepository {
	var returns awsecr.IRepository
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxArmBuildImage) SecretsManagerCredentials() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"secretsManagerCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxArmBuildImage) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Uses a Docker image provided by CodeBuild.
//
// Returns: A Docker image provided by CodeBuild.
//
// Example:
//   "aws/codebuild/amazonlinux2-aarch64-standard:1.0"
//
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-available.html
//
func LinuxArmBuildImage_FromCodeBuildImageId(id *string) IBuildImage {
	_init_.Initialize()

	if err := validateLinuxArmBuildImage_FromCodeBuildImageIdParameters(id); err != nil {
		panic(err)
	}
	var returns IBuildImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.LinuxArmBuildImage",
		"fromCodeBuildImageId",
		[]interface{}{id},
		&returns,
	)

	return returns
}

// Returns: a x86-64 Linux build image from a Docker Hub image.
func LinuxArmBuildImage_FromDockerRegistry(name *string, options *DockerImageOptions) IBuildImage {
	_init_.Initialize()

	if err := validateLinuxArmBuildImage_FromDockerRegistryParameters(name, options); err != nil {
		panic(err)
	}
	var returns IBuildImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.LinuxArmBuildImage",
		"fromDockerRegistry",
		[]interface{}{name, options},
		&returns,
	)

	return returns
}

// Returns an ARM image running Linux from an ECR repository.
//
// NOTE: if the repository is external (i.e. imported), then we won't be able to add
// a resource policy statement for it so CodeBuild can pull the image.
//
// Returns: An aarch64 Linux build image from an ECR repository.
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-ecr.html
//
func LinuxArmBuildImage_FromEcrRepository(repository awsecr.IRepository, tagOrDigest *string) IBuildImage {
	_init_.Initialize()

	if err := validateLinuxArmBuildImage_FromEcrRepositoryParameters(repository); err != nil {
		panic(err)
	}
	var returns IBuildImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.LinuxArmBuildImage",
		"fromEcrRepository",
		[]interface{}{repository, tagOrDigest},
		&returns,
	)

	return returns
}

func LinuxArmBuildImage_AMAZON_LINUX_2_STANDARD_1_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxArmBuildImage",
		"AMAZON_LINUX_2_STANDARD_1_0",
		&returns,
	)
	return returns
}

func LinuxArmBuildImage_AMAZON_LINUX_2_STANDARD_2_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxArmBuildImage",
		"AMAZON_LINUX_2_STANDARD_2_0",
		&returns,
	)
	return returns
}

func LinuxArmBuildImage_AMAZON_LINUX_2_STANDARD_3_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxArmBuildImage",
		"AMAZON_LINUX_2_STANDARD_3_0",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LinuxArmBuildImage) RunScriptBuildspec(entrypoint *string) BuildSpec {
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

func (l *jsiiProxy_LinuxArmBuildImage) Validate(buildEnvironment *BuildEnvironment) *[]*string {
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

