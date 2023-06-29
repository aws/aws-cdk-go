package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/cloudassemblyschema"
)

// Use the CDK classic way of referencing assets.
//
// This synthesizer will generate CloudFormation parameters for every referenced
// asset, and use the CLI's current credentials to deploy the stack.
//
// - It does not support cross-account deployment (the CLI must have credentials
//   to the account you are trying to deploy to).
// - It cannot be used with **CDK Pipelines**. To deploy using CDK Pipelines,
//   you must use the `DefaultStackSynthesizer`.
// - Each asset will take up a CloudFormation Parameter in your template. Keep in
//   mind that there is a maximum of 200 parameters in a CloudFormation template.
//   To use deterministic asset locations instead, use `CliCredentialsStackSynthesizer`.
//
// Be aware that your CLI credentials must be valid for the duration of the
// entire deployment. If you are using session credentials, make sure the
// session lifetime is long enough.
//
// This is the only StackSynthesizer that supports customizing asset behavior
// by overriding `Stack.addFileAsset()` and `Stack.addDockerImageAsset()`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   legacyStackSynthesizer := cdk.NewLegacyStackSynthesizer()
//
type LegacyStackSynthesizer interface {
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

// The jsii proxy struct for LegacyStackSynthesizer
type jsiiProxy_LegacyStackSynthesizer struct {
	jsiiProxy_StackSynthesizer
	jsiiProxy_IBoundStackSynthesizer
	jsiiProxy_IReusableStackSynthesizer
}

func (j *jsiiProxy_LegacyStackSynthesizer) BootstrapQualifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapQualifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyStackSynthesizer) BoundStack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"boundStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyStackSynthesizer) LookupRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookupRole",
		&returns,
	)
	return returns
}


func NewLegacyStackSynthesizer() LegacyStackSynthesizer {
	_init_.Initialize()

	j := jsiiProxy_LegacyStackSynthesizer{}

	_jsii_.Create(
		"aws-cdk-lib.LegacyStackSynthesizer",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewLegacyStackSynthesizer_Override(l LegacyStackSynthesizer) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.LegacyStackSynthesizer",
		nil, // no parameters
		l,
	)
}

func (l *jsiiProxy_LegacyStackSynthesizer) AddBootstrapVersionRule(requiredVersion *float64, bootstrapStackVersionSsmParameter *string) {
	if err := l.validateAddBootstrapVersionRuleParameters(requiredVersion, bootstrapStackVersionSsmParameter); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addBootstrapVersionRule",
		[]interface{}{requiredVersion, bootstrapStackVersionSsmParameter},
	)
}

func (l *jsiiProxy_LegacyStackSynthesizer) AddDockerImageAsset(asset *DockerImageAssetSource) *DockerImageAssetLocation {
	if err := l.validateAddDockerImageAssetParameters(asset); err != nil {
		panic(err)
	}
	var returns *DockerImageAssetLocation

	_jsii_.Invoke(
		l,
		"addDockerImageAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyStackSynthesizer) AddFileAsset(asset *FileAssetSource) *FileAssetLocation {
	if err := l.validateAddFileAssetParameters(asset); err != nil {
		panic(err)
	}
	var returns *FileAssetLocation

	_jsii_.Invoke(
		l,
		"addFileAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyStackSynthesizer) Bind(stack Stack) {
	if err := l.validateBindParameters(stack); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"bind",
		[]interface{}{stack},
	)
}

func (l *jsiiProxy_LegacyStackSynthesizer) CloudFormationLocationFromDockerImageAsset(dest *cloudassemblyschema.DockerImageDestination) *DockerImageAssetLocation {
	if err := l.validateCloudFormationLocationFromDockerImageAssetParameters(dest); err != nil {
		panic(err)
	}
	var returns *DockerImageAssetLocation

	_jsii_.Invoke(
		l,
		"cloudFormationLocationFromDockerImageAsset",
		[]interface{}{dest},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyStackSynthesizer) CloudFormationLocationFromFileAsset(location *cloudassemblyschema.FileDestination) *FileAssetLocation {
	if err := l.validateCloudFormationLocationFromFileAssetParameters(location); err != nil {
		panic(err)
	}
	var returns *FileAssetLocation

	_jsii_.Invoke(
		l,
		"cloudFormationLocationFromFileAsset",
		[]interface{}{location},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyStackSynthesizer) EmitArtifact(session ISynthesisSession, options *SynthesizeStackArtifactOptions) {
	if err := l.validateEmitArtifactParameters(session, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"emitArtifact",
		[]interface{}{session, options},
	)
}

func (l *jsiiProxy_LegacyStackSynthesizer) EmitStackArtifact(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions) {
	if err := l.validateEmitStackArtifactParameters(stack, session, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"emitStackArtifact",
		[]interface{}{stack, session, options},
	)
}

func (l *jsiiProxy_LegacyStackSynthesizer) ReusableBind(stack Stack) IBoundStackSynthesizer {
	if err := l.validateReusableBindParameters(stack); err != nil {
		panic(err)
	}
	var returns IBoundStackSynthesizer

	_jsii_.Invoke(
		l,
		"reusableBind",
		[]interface{}{stack},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyStackSynthesizer) Synthesize(session ISynthesisSession) {
	if err := l.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		[]interface{}{session},
	)
}

func (l *jsiiProxy_LegacyStackSynthesizer) SynthesizeStackTemplate(stack Stack, session ISynthesisSession) {
	if err := l.validateSynthesizeStackTemplateParameters(stack, session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"synthesizeStackTemplate",
		[]interface{}{stack, session},
	)
}

func (l *jsiiProxy_LegacyStackSynthesizer) SynthesizeTemplate(session ISynthesisSession, lookupRoleArn *string) *FileAssetSource {
	if err := l.validateSynthesizeTemplateParameters(session); err != nil {
		panic(err)
	}
	var returns *FileAssetSource

	_jsii_.Invoke(
		l,
		"synthesizeTemplate",
		[]interface{}{session, lookupRoleArn},
		&returns,
	)

	return returns
}

