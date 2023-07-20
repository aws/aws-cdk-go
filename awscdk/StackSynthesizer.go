package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/cloudassemblyschema"
)

// Base class for implementing an IStackSynthesizer.
//
// This class needs to exist to provide public surface area for external
// implementations of stack synthesizers. The protected methods give
// access to functions that are otherwise.
type StackSynthesizer interface {
	IStackSynthesizer
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

// The jsii proxy struct for StackSynthesizer
type jsiiProxy_StackSynthesizer struct {
	jsiiProxy_IStackSynthesizer
}

func (j *jsiiProxy_StackSynthesizer) BootstrapQualifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapQualifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackSynthesizer) BoundStack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"boundStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackSynthesizer) LookupRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookupRole",
		&returns,
	)
	return returns
}


func NewStackSynthesizer_Override(s StackSynthesizer) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.StackSynthesizer",
		nil, // no parameters
		s,
	)
}

func (s *jsiiProxy_StackSynthesizer) AddBootstrapVersionRule(requiredVersion *float64, bootstrapStackVersionSsmParameter *string) {
	if err := s.validateAddBootstrapVersionRuleParameters(requiredVersion, bootstrapStackVersionSsmParameter); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addBootstrapVersionRule",
		[]interface{}{requiredVersion, bootstrapStackVersionSsmParameter},
	)
}

func (s *jsiiProxy_StackSynthesizer) AddDockerImageAsset(asset *DockerImageAssetSource) *DockerImageAssetLocation {
	if err := s.validateAddDockerImageAssetParameters(asset); err != nil {
		panic(err)
	}
	var returns *DockerImageAssetLocation

	_jsii_.Invoke(
		s,
		"addDockerImageAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackSynthesizer) AddFileAsset(asset *FileAssetSource) *FileAssetLocation {
	if err := s.validateAddFileAssetParameters(asset); err != nil {
		panic(err)
	}
	var returns *FileAssetLocation

	_jsii_.Invoke(
		s,
		"addFileAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackSynthesizer) Bind(stack Stack) {
	if err := s.validateBindParameters(stack); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"bind",
		[]interface{}{stack},
	)
}

func (s *jsiiProxy_StackSynthesizer) CloudFormationLocationFromDockerImageAsset(dest *cloudassemblyschema.DockerImageDestination) *DockerImageAssetLocation {
	if err := s.validateCloudFormationLocationFromDockerImageAssetParameters(dest); err != nil {
		panic(err)
	}
	var returns *DockerImageAssetLocation

	_jsii_.Invoke(
		s,
		"cloudFormationLocationFromDockerImageAsset",
		[]interface{}{dest},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackSynthesizer) CloudFormationLocationFromFileAsset(location *cloudassemblyschema.FileDestination) *FileAssetLocation {
	if err := s.validateCloudFormationLocationFromFileAssetParameters(location); err != nil {
		panic(err)
	}
	var returns *FileAssetLocation

	_jsii_.Invoke(
		s,
		"cloudFormationLocationFromFileAsset",
		[]interface{}{location},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackSynthesizer) EmitArtifact(session ISynthesisSession, options *SynthesizeStackArtifactOptions) {
	if err := s.validateEmitArtifactParameters(session, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"emitArtifact",
		[]interface{}{session, options},
	)
}

func (s *jsiiProxy_StackSynthesizer) EmitStackArtifact(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions) {
	if err := s.validateEmitStackArtifactParameters(stack, session, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"emitStackArtifact",
		[]interface{}{stack, session, options},
	)
}

func (s *jsiiProxy_StackSynthesizer) Synthesize(session ISynthesisSession) {
	if err := s.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_StackSynthesizer) SynthesizeStackTemplate(stack Stack, session ISynthesisSession) {
	if err := s.validateSynthesizeStackTemplateParameters(stack, session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"synthesizeStackTemplate",
		[]interface{}{stack, session},
	)
}

func (s *jsiiProxy_StackSynthesizer) SynthesizeTemplate(session ISynthesisSession, lookupRoleArn *string) *FileAssetSource {
	if err := s.validateSynthesizeTemplateParameters(session); err != nil {
		panic(err)
	}
	var returns *FileAssetSource

	_jsii_.Invoke(
		s,
		"synthesizeTemplate",
		[]interface{}{session, lookupRoleArn},
		&returns,
	)

	return returns
}

