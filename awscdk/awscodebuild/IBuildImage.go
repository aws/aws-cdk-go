package awscodebuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Represents a Docker image used for the CodeBuild Project builds.
//
// Use the concrete subclasses, either:
// `LinuxBuildImage` or `WindowsBuildImage`.
type IBuildImage interface {
	// Make a buildspec to run the indicated script.
	RunScriptBuildspec(entrypoint *string) BuildSpec
	// Allows the image a chance to validate whether the passed configuration is correct.
	Validate(buildEnvironment *BuildEnvironment) *[]*string
	// The default `ComputeType` to use with this image, if one was not specified in `BuildEnvironment#computeType` explicitly.
	DefaultComputeType() ComputeType
	// The Docker image identifier that the build environment uses.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-available.html
	//
	ImageId() *string
	// The type of principal that CodeBuild will use to pull this build Docker image.
	// Default: ImagePullPrincipalType.SERVICE_ROLE
	//
	ImagePullPrincipalType() ImagePullPrincipalType
	// An optional ECR repository that the image is hosted in.
	// Default: no repository.
	//
	Repository() awsecr.IRepository
	// The secretsManagerCredentials for access to a private registry.
	// Default: no credentials will be used.
	//
	SecretsManagerCredentials() awssecretsmanager.ISecret
	// The type of build environment.
	Type() *string
}

// The jsii proxy for IBuildImage
type jsiiProxy_IBuildImage struct {
	_ byte // padding
}

func (i *jsiiProxy_IBuildImage) RunScriptBuildspec(entrypoint *string) BuildSpec {
	if err := i.validateRunScriptBuildspecParameters(entrypoint); err != nil {
		panic(err)
	}
	var returns BuildSpec

	_jsii_.Invoke(
		i,
		"runScriptBuildspec",
		[]interface{}{entrypoint},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBuildImage) Validate(buildEnvironment *BuildEnvironment) *[]*string {
	if err := i.validateValidateParameters(buildEnvironment); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"validate",
		[]interface{}{buildEnvironment},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IBuildImage) DefaultComputeType() ComputeType {
	var returns ComputeType
	_jsii_.Get(
		j,
		"defaultComputeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBuildImage) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBuildImage) ImagePullPrincipalType() ImagePullPrincipalType {
	var returns ImagePullPrincipalType
	_jsii_.Get(
		j,
		"imagePullPrincipalType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBuildImage) Repository() awsecr.IRepository {
	var returns awsecr.IRepository
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBuildImage) SecretsManagerCredentials() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"secretsManagerCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBuildImage) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

