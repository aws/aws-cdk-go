// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Synthesizer that reuses bootstrap roles from a different region.
//
// A special synthesizer that behaves similarly to `DefaultStackSynthesizer`,
// but doesn't require bootstrapping the environment it operates in. Instead,
// it will re-use the Roles that were created for a different region (which
// is possible because IAM is a global service).
//
// However, it will not assume asset buckets or repositories have been created,
// and therefore does not support assets.
//
// Used by the CodePipeline construct for the support stacks needed for
// cross-region replication S3 buckets. App builders do not need to use this
// synthesizer directly.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   bootstraplessSynthesizer := monocdk.NewBootstraplessSynthesizer(&bootstraplessSynthesizerProps{
//   	cloudFormationExecutionRoleArn: jsii.String("cloudFormationExecutionRoleArn"),
//   	deployRoleArn: jsii.String("deployRoleArn"),
//   })
//
// Experimental.
type BootstraplessSynthesizer interface {
	DefaultStackSynthesizer
	// Returns the ARN of the CFN execution Role.
	// Experimental.
	CloudFormationExecutionRoleArn() *string
	// Returns the ARN of the deploy Role.
	// Experimental.
	DeployRoleArn() *string
	// Experimental.
	Stack() Stack
	// Register a Docker Image Asset.
	//
	// Returns the parameters that can be used to refer to the asset inside the template.
	// Experimental.
	AddDockerImageAsset(_asset *DockerImageAssetSource) *DockerImageAssetLocation
	// Register a File Asset.
	//
	// Returns the parameters that can be used to refer to the asset inside the template.
	// Experimental.
	AddFileAsset(_asset *FileAssetSource) *FileAssetLocation
	// Bind to the stack this environment is going to be used on.
	//
	// Must be called before any of the other methods are called.
	// Experimental.
	Bind(stack Stack)
	// Write the stack artifact to the session.
	//
	// Use default settings to add a CloudFormationStackArtifact artifact to
	// the given synthesis session.
	// Experimental.
	EmitStackArtifact(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions)
	// Synthesize the associated stack to the session.
	// Experimental.
	Synthesize(session ISynthesisSession)
	// Have the stack write out its template.
	// Experimental.
	SynthesizeStackTemplate(stack Stack, session ISynthesisSession)
}

// The jsii proxy struct for BootstraplessSynthesizer
type jsiiProxy_BootstraplessSynthesizer struct {
	jsiiProxy_DefaultStackSynthesizer
}

func (j *jsiiProxy_BootstraplessSynthesizer) CloudFormationExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudFormationExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BootstraplessSynthesizer) DeployRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deployRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BootstraplessSynthesizer) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewBootstraplessSynthesizer(props *BootstraplessSynthesizerProps) BootstraplessSynthesizer {
	_init_.Initialize()

	if err := validateNewBootstraplessSynthesizerParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_BootstraplessSynthesizer{}

	_jsii_.Create(
		"monocdk.BootstraplessSynthesizer",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewBootstraplessSynthesizer_Override(b BootstraplessSynthesizer, props *BootstraplessSynthesizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.BootstraplessSynthesizer",
		[]interface{}{props},
		b,
	)
}

func BootstraplessSynthesizer_DEFAULT_BOOTSTRAP_STACK_VERSION_SSM_PARAMETER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.BootstraplessSynthesizer",
		"DEFAULT_BOOTSTRAP_STACK_VERSION_SSM_PARAMETER",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_CLOUDFORMATION_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.BootstraplessSynthesizer",
		"DEFAULT_CLOUDFORMATION_ROLE_ARN",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_DEPLOY_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.BootstraplessSynthesizer",
		"DEFAULT_DEPLOY_ROLE_ARN",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_DOCKER_ASSET_PREFIX() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.BootstraplessSynthesizer",
		"DEFAULT_DOCKER_ASSET_PREFIX",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_FILE_ASSET_KEY_ARN_EXPORT_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.BootstraplessSynthesizer",
		"DEFAULT_FILE_ASSET_KEY_ARN_EXPORT_NAME",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_FILE_ASSET_PREFIX() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.BootstraplessSynthesizer",
		"DEFAULT_FILE_ASSET_PREFIX",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_FILE_ASSET_PUBLISHING_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.BootstraplessSynthesizer",
		"DEFAULT_FILE_ASSET_PUBLISHING_ROLE_ARN",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_FILE_ASSETS_BUCKET_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.BootstraplessSynthesizer",
		"DEFAULT_FILE_ASSETS_BUCKET_NAME",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_IMAGE_ASSET_PUBLISHING_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.BootstraplessSynthesizer",
		"DEFAULT_IMAGE_ASSET_PUBLISHING_ROLE_ARN",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_IMAGE_ASSETS_REPOSITORY_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.BootstraplessSynthesizer",
		"DEFAULT_IMAGE_ASSETS_REPOSITORY_NAME",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_LOOKUP_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.BootstraplessSynthesizer",
		"DEFAULT_LOOKUP_ROLE_ARN",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_QUALIFIER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.BootstraplessSynthesizer",
		"DEFAULT_QUALIFIER",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BootstraplessSynthesizer) AddDockerImageAsset(_asset *DockerImageAssetSource) *DockerImageAssetLocation {
	if err := b.validateAddDockerImageAssetParameters(_asset); err != nil {
		panic(err)
	}
	var returns *DockerImageAssetLocation

	_jsii_.Invoke(
		b,
		"addDockerImageAsset",
		[]interface{}{_asset},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BootstraplessSynthesizer) AddFileAsset(_asset *FileAssetSource) *FileAssetLocation {
	if err := b.validateAddFileAssetParameters(_asset); err != nil {
		panic(err)
	}
	var returns *FileAssetLocation

	_jsii_.Invoke(
		b,
		"addFileAsset",
		[]interface{}{_asset},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BootstraplessSynthesizer) Bind(stack Stack) {
	if err := b.validateBindParameters(stack); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"bind",
		[]interface{}{stack},
	)
}

func (b *jsiiProxy_BootstraplessSynthesizer) EmitStackArtifact(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions) {
	if err := b.validateEmitStackArtifactParameters(stack, session, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"emitStackArtifact",
		[]interface{}{stack, session, options},
	)
}

func (b *jsiiProxy_BootstraplessSynthesizer) Synthesize(session ISynthesisSession) {
	if err := b.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"synthesize",
		[]interface{}{session},
	)
}

func (b *jsiiProxy_BootstraplessSynthesizer) SynthesizeStackTemplate(stack Stack, session ISynthesisSession) {
	if err := b.validateSynthesizeStackTemplateParameters(stack, session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"synthesizeStackTemplate",
		[]interface{}{stack, session},
	)
}

