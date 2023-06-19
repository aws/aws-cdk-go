package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Use the CDK classic way of referencing assets.
//
// This synthesizer will generate CloudFormation parameters for every referenced
// asset, and use the CLI's current credentials to deploy the stack.
//
// - It does not support cross-account deployment (the CLI must have credentials
//    to the account you are trying to deploy to).
// - It cannot be used with **CDK Pipelines**. To deploy using CDK Pipelines,
//    you must use the `DefaultStackSynthesizer`.
// - Each asset will take up a CloudFormation Parameter in your template. Keep in
//    mind that there is a maximum of 200 parameters in a CloudFormation template.
//    To use determinstic asset locations instead, use `CliCredentialsStackSynthesizer`.
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
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   legacyStackSynthesizer := monocdk.NewLegacyStackSynthesizer()
//
// Experimental.
type LegacyStackSynthesizer interface {
	StackSynthesizer
	// Register a Docker Image Asset.
	//
	// Returns the parameters that can be used to refer to the asset inside the template.
	// Experimental.
	AddDockerImageAsset(asset *DockerImageAssetSource) *DockerImageAssetLocation
	// Register a File Asset.
	//
	// Returns the parameters that can be used to refer to the asset inside the template.
	// Experimental.
	AddFileAsset(asset *FileAssetSource) *FileAssetLocation
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

// The jsii proxy struct for LegacyStackSynthesizer
type jsiiProxy_LegacyStackSynthesizer struct {
	jsiiProxy_StackSynthesizer
}

// Experimental.
func NewLegacyStackSynthesizer() LegacyStackSynthesizer {
	_init_.Initialize()

	j := jsiiProxy_LegacyStackSynthesizer{}

	_jsii_.Create(
		"monocdk.LegacyStackSynthesizer",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewLegacyStackSynthesizer_Override(l LegacyStackSynthesizer) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.LegacyStackSynthesizer",
		nil, // no parameters
		l,
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

