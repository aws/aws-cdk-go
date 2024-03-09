package awscodebuild

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecrassets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CodeBuild image running Windows.
//
// This class has a bunch of public constants that represent the most popular images.
//
// You can also specify a custom image using one of the static methods:
//
// - WindowsBuildImage.fromDockerRegistry(image[, { secretsManagerCredentials }, imageType])
// - WindowsBuildImage.fromEcrRepository(repo[, tag, imageType])
// - WindowsBuildImage.fromAsset(parent, id, props, [, imageType])
//
// Example:
//   var ecrRepository repository
//
//
//   codebuild.NewProject(this, jsii.String("Project"), &ProjectProps{
//   	Environment: &BuildEnvironment{
//   		BuildImage: codebuild.WindowsBuildImage_FromEcrRepository(ecrRepository, jsii.String("v1.0"), codebuild.WindowsImageType_SERVER_2019),
//   		// optional certificate to include in the build image
//   		Certificate: &BuildEnvironmentCertificate{
//   			Bucket: s3.Bucket_FromBucketName(this, jsii.String("Bucket"), jsii.String("my-bucket")),
//   			ObjectKey: jsii.String("path/to/cert.pem"),
//   		},
//   	},
//   })
//
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-available.html
//
type WindowsBuildImage interface {
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
	// Allows the image a chance to validate whether the passed configuration is correct.
	Validate(buildEnvironment *BuildEnvironment) *[]*string
}

// The jsii proxy struct for WindowsBuildImage
type jsiiProxy_WindowsBuildImage struct {
	jsiiProxy_IBuildImage
}

func (j *jsiiProxy_WindowsBuildImage) DefaultComputeType() ComputeType {
	var returns ComputeType
	_jsii_.Get(
		j,
		"defaultComputeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsBuildImage) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsBuildImage) ImagePullPrincipalType() ImagePullPrincipalType {
	var returns ImagePullPrincipalType
	_jsii_.Get(
		j,
		"imagePullPrincipalType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsBuildImage) Repository() awsecr.IRepository {
	var returns awsecr.IRepository
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsBuildImage) SecretsManagerCredentials() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"secretsManagerCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsBuildImage) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Uses an Docker image asset as a Windows build image.
func WindowsBuildImage_FromAsset(scope constructs.Construct, id *string, props *awsecrassets.DockerImageAssetProps, imageType WindowsImageType) IBuildImage {
	_init_.Initialize()

	if err := validateWindowsBuildImage_FromAssetParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns IBuildImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.WindowsBuildImage",
		"fromAsset",
		[]interface{}{scope, id, props, imageType},
		&returns,
	)

	return returns
}

// Returns: a Windows build image from a Docker Hub image.
func WindowsBuildImage_FromDockerRegistry(name *string, options *DockerImageOptions, imageType WindowsImageType) IBuildImage {
	_init_.Initialize()

	if err := validateWindowsBuildImage_FromDockerRegistryParameters(name, options); err != nil {
		panic(err)
	}
	var returns IBuildImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.WindowsBuildImage",
		"fromDockerRegistry",
		[]interface{}{name, options, imageType},
		&returns,
	)

	return returns
}

// Returns: A Windows build image from an ECR repository.
//
// NOTE: if the repository is external (i.e. imported), then we won't be able to add
// a resource policy statement for it so CodeBuild can pull the image.
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-ecr.html
//
func WindowsBuildImage_FromEcrRepository(repository awsecr.IRepository, tagOrDigest *string, imageType WindowsImageType) IBuildImage {
	_init_.Initialize()

	if err := validateWindowsBuildImage_FromEcrRepositoryParameters(repository); err != nil {
		panic(err)
	}
	var returns IBuildImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.WindowsBuildImage",
		"fromEcrRepository",
		[]interface{}{repository, tagOrDigest, imageType},
		&returns,
	)

	return returns
}

func WindowsBuildImage_WIN_SERVER_CORE_2019_BASE() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.WindowsBuildImage",
		"WIN_SERVER_CORE_2019_BASE",
		&returns,
	)
	return returns
}

func WindowsBuildImage_WIN_SERVER_CORE_2019_BASE_2_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.WindowsBuildImage",
		"WIN_SERVER_CORE_2019_BASE_2_0",
		&returns,
	)
	return returns
}

func WindowsBuildImage_WIN_SERVER_CORE_2019_BASE_3_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.WindowsBuildImage",
		"WIN_SERVER_CORE_2019_BASE_3_0",
		&returns,
	)
	return returns
}

func WindowsBuildImage_WINDOWS_BASE_2_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.WindowsBuildImage",
		"WINDOWS_BASE_2_0",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_WindowsBuildImage) RunScriptBuildspec(entrypoint *string) BuildSpec {
	if err := w.validateRunScriptBuildspecParameters(entrypoint); err != nil {
		panic(err)
	}
	var returns BuildSpec

	_jsii_.Invoke(
		w,
		"runScriptBuildspec",
		[]interface{}{entrypoint},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsBuildImage) Validate(buildEnvironment *BuildEnvironment) *[]*string {
	if err := w.validateValidateParameters(buildEnvironment); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"validate",
		[]interface{}{buildEnvironment},
		&returns,
	)

	return returns
}

