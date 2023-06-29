package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/cloudassemblyschema"
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
// The name is poorly chosen -- it does still require bootstrapping, it just
// does not support assets.
//
// Used by the CodePipeline construct for the support stacks needed for
// cross-region replication S3 buckets. App builders do not need to use this
// synthesizer directly.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   bootstraplessSynthesizer := cdk.NewBootstraplessSynthesizer(&BootstraplessSynthesizerProps{
//   	CloudFormationExecutionRoleArn: jsii.String("cloudFormationExecutionRoleArn"),
//   	DeployRoleArn: jsii.String("deployRoleArn"),
//   })
//
type BootstraplessSynthesizer interface {
	DefaultStackSynthesizer
	// The qualifier used to bootstrap this stack.
	BootstrapQualifier() *string
	// Retrieve the bound stack.
	//
	// Fails if the stack hasn't been bound yet.
	BoundStack() Stack
	// Returns the ARN of the CFN execution Role.
	CloudFormationExecutionRoleArn() *string
	// Returns the ARN of the deploy Role.
	DeployRoleArn() *string
	// The role used to lookup for this stack.
	LookupRole() *string
	// Return the currently bound stack.
	// Deprecated: Use `boundStack` instead.
	Stack() Stack
	// Add a CfnRule to the bound stack that checks whether an SSM parameter exceeds a given version.
	//
	// This will modify the template, so must be called before the stack is synthesized.
	AddBootstrapVersionRule(requiredVersion *float64, bootstrapStackVersionSsmParameter *string)
	// Register a Docker Image Asset.
	//
	// Returns the parameters that can be used to refer to the asset inside the template.
	//
	// The synthesizer must rely on some out-of-band mechanism to make sure the given files
	// are actually placed in the returned location before the deployment happens. This can
	// be by writing the instructions to the asset manifest (for use by the `cdk-assets` tool),
	// by relying on the CLI to upload files (legacy behavior), or some other operator controlled
	// mechanism.
	AddDockerImageAsset(_asset *DockerImageAssetSource) *DockerImageAssetLocation
	// Register a File Asset.
	//
	// Returns the parameters that can be used to refer to the asset inside the template.
	//
	// The synthesizer must rely on some out-of-band mechanism to make sure the given files
	// are actually placed in the returned location before the deployment happens. This can
	// be by writing the instructions to the asset manifest (for use by the `cdk-assets` tool),
	// by relying on the CLI to upload files (legacy behavior), or some other operator controlled
	// mechanism.
	AddFileAsset(_asset *FileAssetSource) *FileAssetLocation
	// Bind to the stack this environment is going to be used on.
	//
	// Must be called before any of the other methods are called.
	Bind(stack Stack)
	// Turn a docker asset location into a CloudFormation representation of that location.
	//
	// If any of the fields contain placeholders, the result will be wrapped in a `Fn.sub`.
	CloudFormationLocationFromDockerImageAsset(dest *cloudassemblyschema.DockerImageDestination) *DockerImageAssetLocation
	// Turn a file asset location into a CloudFormation representation of that location.
	//
	// If any of the fields contain placeholders, the result will be wrapped in a `Fn.sub`.
	CloudFormationLocationFromFileAsset(location *cloudassemblyschema.FileDestination) *FileAssetLocation
	// Write the CloudFormation stack artifact to the session.
	//
	// Use default settings to add a CloudFormationStackArtifact artifact to
	// the given synthesis session. The Stack artifact will control the settings for the
	// CloudFormation deployment.
	EmitArtifact(session ISynthesisSession, options *SynthesizeStackArtifactOptions)
	// Write the stack artifact to the session.
	//
	// Use default settings to add a CloudFormationStackArtifact artifact to
	// the given synthesis session.
	// Deprecated: Use `emitArtifact` instead.
	EmitStackArtifact(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions)
	// Produce a bound Stack Synthesizer for the given stack.
	//
	// This method may be called more than once on the same object.
	ReusableBind(stack Stack) IBoundStackSynthesizer
	// Synthesize the associated stack to the session.
	Synthesize(session ISynthesisSession)
	// Synthesize the stack template to the given session, passing the configured lookup role ARN.
	SynthesizeStackTemplate(stack Stack, session ISynthesisSession)
	// Write the stack template to the given session.
	//
	// Return a descriptor that represents the stack template as a file asset
	// source, for adding to an asset manifest (if desired). This can be used to
	// have the asset manifest system (`cdk-assets`) upload the template to S3
	// using the appropriate role, so that afterwards only a CloudFormation
	// deployment is necessary.
	//
	// If the template is uploaded as an asset, the `stackTemplateAssetObjectUrl`
	// property should be set when calling `emitArtifact.`
	//
	// If the template is *NOT* uploaded as an asset first and the template turns
	// out to be >50KB, it will need to be uploaded to S3 anyway. At that point
	// the credentials will be the same identity that is doing the `UpdateStack`
	// call, which may not have the right permissions to write to S3.
	SynthesizeTemplate(session ISynthesisSession, lookupRoleArn *string) *FileAssetSource
}

// The jsii proxy struct for BootstraplessSynthesizer
type jsiiProxy_BootstraplessSynthesizer struct {
	jsiiProxy_DefaultStackSynthesizer
}

func (j *jsiiProxy_BootstraplessSynthesizer) BootstrapQualifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapQualifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BootstraplessSynthesizer) BoundStack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"boundStack",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_BootstraplessSynthesizer) LookupRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookupRole",
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


func NewBootstraplessSynthesizer(props *BootstraplessSynthesizerProps) BootstraplessSynthesizer {
	_init_.Initialize()

	if err := validateNewBootstraplessSynthesizerParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_BootstraplessSynthesizer{}

	_jsii_.Create(
		"aws-cdk-lib.BootstraplessSynthesizer",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewBootstraplessSynthesizer_Override(b BootstraplessSynthesizer, props *BootstraplessSynthesizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.BootstraplessSynthesizer",
		[]interface{}{props},
		b,
	)
}

func BootstraplessSynthesizer_DEFAULT_BOOTSTRAP_STACK_VERSION_SSM_PARAMETER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_BOOTSTRAP_STACK_VERSION_SSM_PARAMETER",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_CLOUDFORMATION_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_CLOUDFORMATION_ROLE_ARN",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_DEPLOY_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_DEPLOY_ROLE_ARN",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_DOCKER_ASSET_PREFIX() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_DOCKER_ASSET_PREFIX",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_FILE_ASSET_KEY_ARN_EXPORT_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_FILE_ASSET_KEY_ARN_EXPORT_NAME",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_FILE_ASSET_PREFIX() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_FILE_ASSET_PREFIX",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_FILE_ASSET_PUBLISHING_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_FILE_ASSET_PUBLISHING_ROLE_ARN",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_FILE_ASSETS_BUCKET_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_FILE_ASSETS_BUCKET_NAME",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_IMAGE_ASSET_PUBLISHING_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_IMAGE_ASSET_PUBLISHING_ROLE_ARN",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_IMAGE_ASSETS_REPOSITORY_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_IMAGE_ASSETS_REPOSITORY_NAME",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_LOOKUP_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_LOOKUP_ROLE_ARN",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_QUALIFIER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_QUALIFIER",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BootstraplessSynthesizer) AddBootstrapVersionRule(requiredVersion *float64, bootstrapStackVersionSsmParameter *string) {
	if err := b.validateAddBootstrapVersionRuleParameters(requiredVersion, bootstrapStackVersionSsmParameter); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addBootstrapVersionRule",
		[]interface{}{requiredVersion, bootstrapStackVersionSsmParameter},
	)
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

func (b *jsiiProxy_BootstraplessSynthesizer) CloudFormationLocationFromDockerImageAsset(dest *cloudassemblyschema.DockerImageDestination) *DockerImageAssetLocation {
	if err := b.validateCloudFormationLocationFromDockerImageAssetParameters(dest); err != nil {
		panic(err)
	}
	var returns *DockerImageAssetLocation

	_jsii_.Invoke(
		b,
		"cloudFormationLocationFromDockerImageAsset",
		[]interface{}{dest},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BootstraplessSynthesizer) CloudFormationLocationFromFileAsset(location *cloudassemblyschema.FileDestination) *FileAssetLocation {
	if err := b.validateCloudFormationLocationFromFileAssetParameters(location); err != nil {
		panic(err)
	}
	var returns *FileAssetLocation

	_jsii_.Invoke(
		b,
		"cloudFormationLocationFromFileAsset",
		[]interface{}{location},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BootstraplessSynthesizer) EmitArtifact(session ISynthesisSession, options *SynthesizeStackArtifactOptions) {
	if err := b.validateEmitArtifactParameters(session, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"emitArtifact",
		[]interface{}{session, options},
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

func (b *jsiiProxy_BootstraplessSynthesizer) ReusableBind(stack Stack) IBoundStackSynthesizer {
	if err := b.validateReusableBindParameters(stack); err != nil {
		panic(err)
	}
	var returns IBoundStackSynthesizer

	_jsii_.Invoke(
		b,
		"reusableBind",
		[]interface{}{stack},
		&returns,
	)

	return returns
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

func (b *jsiiProxy_BootstraplessSynthesizer) SynthesizeTemplate(session ISynthesisSession, lookupRoleArn *string) *FileAssetSource {
	if err := b.validateSynthesizeTemplateParameters(session); err != nil {
		panic(err)
	}
	var returns *FileAssetSource

	_jsii_.Invoke(
		b,
		"synthesizeTemplate",
		[]interface{}{session, lookupRoleArn},
		&returns,
	)

	return returns
}

