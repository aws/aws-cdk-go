package awscodebuild

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecrassets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CodeBuild image running ARM MacOS.
//
// This class has a bunch of public constants that represent the most popular images.
//
// You can also specify a custom image using one of the static methods:
//
// - MacBuildImage.fromDockerRegistry(image[, { secretsManagerCredentials }])
// - MacBuildImage.fromEcrRepository(repo[, tag])
// - MacBuildImage.fromAsset(parent, id, props)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var networkMode NetworkMode
//   var platform Platform
//
//   macBuildImage := awscdk.Aws_codebuild.MacBuildImage_FromAsset(this, jsii.String("MyMacBuildImage"), &DockerImageAssetProps{
//   	Directory: jsii.String("directory"),
//
//   	// the properties below are optional
//   	AssetName: jsii.String("assetName"),
//   	BuildArgs: map[string]*string{
//   		"buildArgsKey": jsii.String("buildArgs"),
//   	},
//   	BuildSecrets: map[string]*string{
//   		"buildSecretsKey": jsii.String("buildSecrets"),
//   	},
//   	BuildSsh: jsii.String("buildSsh"),
//   	CacheDisabled: jsii.Boolean(false),
//   	CacheFrom: []DockerCacheOption{
//   		&DockerCacheOption{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Params: map[string]*string{
//   				"paramsKey": jsii.String("params"),
//   			},
//   		},
//   	},
//   	CacheTo: &DockerCacheOption{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		Params: map[string]*string{
//   			"paramsKey": jsii.String("params"),
//   		},
//   	},
//   	DisplayName: jsii.String("displayName"),
//   	Exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	ExtraHash: jsii.String("extraHash"),
//   	File: jsii.String("file"),
//   	FollowSymlinks: cdk.SymlinkFollowMode_NEVER,
//   	IgnoreMode: cdk.IgnoreMode_GLOB,
//   	Invalidation: &DockerImageAssetInvalidationOptions{
//   		BuildArgs: jsii.Boolean(false),
//   		BuildSecrets: jsii.Boolean(false),
//   		BuildSsh: jsii.Boolean(false),
//   		ExtraHash: jsii.Boolean(false),
//   		File: jsii.Boolean(false),
//   		NetworkMode: jsii.Boolean(false),
//   		Outputs: jsii.Boolean(false),
//   		Platform: jsii.Boolean(false),
//   		RepositoryName: jsii.Boolean(false),
//   		Target: jsii.Boolean(false),
//   	},
//   	NetworkMode: networkMode,
//   	Outputs: []*string{
//   		jsii.String("outputs"),
//   	},
//   	Platform: platform,
//   	Target: jsii.String("target"),
//   })
//
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-available.html
//
type MacBuildImage interface {
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

// The jsii proxy struct for MacBuildImage
type jsiiProxy_MacBuildImage struct {
	jsiiProxy_IBuildImage
}

func (j *jsiiProxy_MacBuildImage) DefaultComputeType() ComputeType {
	var returns ComputeType
	_jsii_.Get(
		j,
		"defaultComputeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacBuildImage) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacBuildImage) ImagePullPrincipalType() ImagePullPrincipalType {
	var returns ImagePullPrincipalType
	_jsii_.Get(
		j,
		"imagePullPrincipalType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacBuildImage) Repository() awsecr.IRepository {
	var returns awsecr.IRepository
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacBuildImage) SecretsManagerCredentials() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"secretsManagerCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacBuildImage) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Uses an Docker image asset as a ARM MacOS build image.
func MacBuildImage_FromAsset(scope constructs.Construct, id *string, props *awsecrassets.DockerImageAssetProps) IBuildImage {
	_init_.Initialize()

	if err := validateMacBuildImage_FromAssetParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns IBuildImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.MacBuildImage",
		"fromAsset",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Makes an ARM MacOS build image from a Docker Hub image.
func MacBuildImage_FromDockerRegistry(name *string, options *DockerImageOptions) IBuildImage {
	_init_.Initialize()

	if err := validateMacBuildImage_FromDockerRegistryParameters(name, options); err != nil {
		panic(err)
	}
	var returns IBuildImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.MacBuildImage",
		"fromDockerRegistry",
		[]interface{}{name, options},
		&returns,
	)

	return returns
}

// Makes an ARM MacOS build image from an ECR repository.
func MacBuildImage_FromEcrRepository(repository awsecr.IRepository, tagOrDigest *string) IBuildImage {
	_init_.Initialize()

	if err := validateMacBuildImage_FromEcrRepositoryParameters(repository); err != nil {
		panic(err)
	}
	var returns IBuildImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.MacBuildImage",
		"fromEcrRepository",
		[]interface{}{repository, tagOrDigest},
		&returns,
	)

	return returns
}

func MacBuildImage_BASE_14() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.MacBuildImage",
		"BASE_14",
		&returns,
	)
	return returns
}

func MacBuildImage_BASE_15() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.MacBuildImage",
		"BASE_15",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MacBuildImage) RunScriptBuildspec(entrypoint *string) BuildSpec {
	if err := m.validateRunScriptBuildspecParameters(entrypoint); err != nil {
		panic(err)
	}
	var returns BuildSpec

	_jsii_.Invoke(
		m,
		"runScriptBuildspec",
		[]interface{}{entrypoint},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MacBuildImage) Validate(buildEnvironment *BuildEnvironment) *[]*string {
	if err := m.validateValidateParameters(buildEnvironment); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"validate",
		[]interface{}{buildEnvironment},
		&returns,
	)

	return returns
}

