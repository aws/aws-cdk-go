package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/cloudassemblyschema"
)

// A synthesizer that uses conventional asset locations, but not conventional deployment roles.
//
// Instead of assuming the bootstrapped deployment roles, all stack operations will be performed
// using the CLI's current credentials.
//
// - This synthesizer does not support deploying to accounts to which the CLI does not have
//   credentials. It also does not support deploying using **CDK Pipelines**. For either of those
//   features, use `DefaultStackSynthesizer`.
// - This synthesizer requires an S3 bucket and ECR repository with well-known names. To
//   not depend on those, use `LegacyStackSynthesizer`.
//
// Be aware that your CLI credentials must be valid for the duration of the
// entire deployment. If you are using session credentials, make sure the
// session lifetime is long enough.
//
// By default, expects the environment to have been bootstrapped with just the staging resources
// of the Bootstrap Stack V2 (also known as "modern bootstrap stack"). You can override
// the default names using the synthesizer's construction properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cliCredentialsStackSynthesizer := cdk.NewCliCredentialsStackSynthesizer(&CliCredentialsStackSynthesizerProps{
//   	BucketPrefix: jsii.String("bucketPrefix"),
//   	DockerTagPrefix: jsii.String("dockerTagPrefix"),
//   	FileAssetsBucketName: jsii.String("fileAssetsBucketName"),
//   	ImageAssetsRepositoryName: jsii.String("imageAssetsRepositoryName"),
//   	Qualifier: jsii.String("qualifier"),
//   })
//
type CliCredentialsStackSynthesizer interface {
	StackSynthesizer
	IBoundStackSynthesizer
	IReusableStackSynthesizer
	// The qualifier used to bootstrap this stack.
	BootstrapQualifier() *string
	// Retrieve the bound stack.
	//
	// Fails if the stack hasn't been bound yet.
	BoundStack() Stack
	// The role used to lookup for this stack.
	LookupRole() *string
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
	AddDockerImageAsset(asset *DockerImageAssetSource) *DockerImageAssetLocation
	// Register a File Asset.
	//
	// Returns the parameters that can be used to refer to the asset inside the template.
	//
	// The synthesizer must rely on some out-of-band mechanism to make sure the given files
	// are actually placed in the returned location before the deployment happens. This can
	// be by writing the instructions to the asset manifest (for use by the `cdk-assets` tool),
	// by relying on the CLI to upload files (legacy behavior), or some other operator controlled
	// mechanism.
	AddFileAsset(asset *FileAssetSource) *FileAssetLocation
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
	// Have the stack write out its template.
	// Deprecated: Use `synthesizeTemplate` instead.
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

// The jsii proxy struct for CliCredentialsStackSynthesizer
type jsiiProxy_CliCredentialsStackSynthesizer struct {
	jsiiProxy_StackSynthesizer
	jsiiProxy_IBoundStackSynthesizer
	jsiiProxy_IReusableStackSynthesizer
}

func (j *jsiiProxy_CliCredentialsStackSynthesizer) BootstrapQualifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapQualifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CliCredentialsStackSynthesizer) BoundStack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"boundStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CliCredentialsStackSynthesizer) LookupRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookupRole",
		&returns,
	)
	return returns
}


func NewCliCredentialsStackSynthesizer(props *CliCredentialsStackSynthesizerProps) CliCredentialsStackSynthesizer {
	_init_.Initialize()

	if err := validateNewCliCredentialsStackSynthesizerParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CliCredentialsStackSynthesizer{}

	_jsii_.Create(
		"aws-cdk-lib.CliCredentialsStackSynthesizer",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewCliCredentialsStackSynthesizer_Override(c CliCredentialsStackSynthesizer, props *CliCredentialsStackSynthesizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CliCredentialsStackSynthesizer",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) AddBootstrapVersionRule(requiredVersion *float64, bootstrapStackVersionSsmParameter *string) {
	if err := c.validateAddBootstrapVersionRuleParameters(requiredVersion, bootstrapStackVersionSsmParameter); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addBootstrapVersionRule",
		[]interface{}{requiredVersion, bootstrapStackVersionSsmParameter},
	)
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) AddDockerImageAsset(asset *DockerImageAssetSource) *DockerImageAssetLocation {
	if err := c.validateAddDockerImageAssetParameters(asset); err != nil {
		panic(err)
	}
	var returns *DockerImageAssetLocation

	_jsii_.Invoke(
		c,
		"addDockerImageAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) AddFileAsset(asset *FileAssetSource) *FileAssetLocation {
	if err := c.validateAddFileAssetParameters(asset); err != nil {
		panic(err)
	}
	var returns *FileAssetLocation

	_jsii_.Invoke(
		c,
		"addFileAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) Bind(stack Stack) {
	if err := c.validateBindParameters(stack); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"bind",
		[]interface{}{stack},
	)
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) CloudFormationLocationFromDockerImageAsset(dest *cloudassemblyschema.DockerImageDestination) *DockerImageAssetLocation {
	if err := c.validateCloudFormationLocationFromDockerImageAssetParameters(dest); err != nil {
		panic(err)
	}
	var returns *DockerImageAssetLocation

	_jsii_.Invoke(
		c,
		"cloudFormationLocationFromDockerImageAsset",
		[]interface{}{dest},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) CloudFormationLocationFromFileAsset(location *cloudassemblyschema.FileDestination) *FileAssetLocation {
	if err := c.validateCloudFormationLocationFromFileAssetParameters(location); err != nil {
		panic(err)
	}
	var returns *FileAssetLocation

	_jsii_.Invoke(
		c,
		"cloudFormationLocationFromFileAsset",
		[]interface{}{location},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) EmitArtifact(session ISynthesisSession, options *SynthesizeStackArtifactOptions) {
	if err := c.validateEmitArtifactParameters(session, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"emitArtifact",
		[]interface{}{session, options},
	)
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) EmitStackArtifact(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions) {
	if err := c.validateEmitStackArtifactParameters(stack, session, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"emitStackArtifact",
		[]interface{}{stack, session, options},
	)
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) ReusableBind(stack Stack) IBoundStackSynthesizer {
	if err := c.validateReusableBindParameters(stack); err != nil {
		panic(err)
	}
	var returns IBoundStackSynthesizer

	_jsii_.Invoke(
		c,
		"reusableBind",
		[]interface{}{stack},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) Synthesize(session ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) SynthesizeStackTemplate(stack Stack, session ISynthesisSession) {
	if err := c.validateSynthesizeStackTemplateParameters(stack, session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesizeStackTemplate",
		[]interface{}{stack, session},
	)
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) SynthesizeTemplate(session ISynthesisSession, lookupRoleArn *string) *FileAssetSource {
	if err := c.validateSynthesizeTemplateParameters(session); err != nil {
		panic(err)
	}
	var returns *FileAssetSource

	_jsii_.Invoke(
		c,
		"synthesizeTemplate",
		[]interface{}{session, lookupRoleArn},
		&returns,
	)

	return returns
}

