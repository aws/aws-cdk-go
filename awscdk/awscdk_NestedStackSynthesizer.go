// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Synthesizer for a nested stack.
//
// Forwards all calls to the parent stack's synthesizer.
//
// This synthesizer is automatically used for `NestedStack` constructs.
// App builder do not need to use this class directly.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var stackSynthesizer stackSynthesizer
//
//   nestedStackSynthesizer := cdk.NewNestedStackSynthesizer(stackSynthesizer)
//
type NestedStackSynthesizer interface {
	StackSynthesizer
	// Register a Docker Image Asset.
	//
	// Returns the parameters that can be used to refer to the asset inside the template.
	AddDockerImageAsset(asset *DockerImageAssetSource) *DockerImageAssetLocation
	// Register a File Asset.
	//
	// Returns the parameters that can be used to refer to the asset inside the template.
	AddFileAsset(asset *FileAssetSource) *FileAssetLocation
	// Bind to the stack this environment is going to be used on.
	//
	// Must be called before any of the other methods are called.
	Bind(stack Stack)
	// Write the stack artifact to the session.
	//
	// Use default settings to add a CloudFormationStackArtifact artifact to
	// the given synthesis session.
	EmitStackArtifact(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions)
	// Synthesize the associated stack to the session.
	Synthesize(session ISynthesisSession)
	// Have the stack write out its template.
	SynthesizeStackTemplate(stack Stack, session ISynthesisSession)
}

// The jsii proxy struct for NestedStackSynthesizer
type jsiiProxy_NestedStackSynthesizer struct {
	jsiiProxy_StackSynthesizer
}

func NewNestedStackSynthesizer(parentDeployment IStackSynthesizer) NestedStackSynthesizer {
	_init_.Initialize()

	if err := validateNewNestedStackSynthesizerParameters(parentDeployment); err != nil {
		panic(err)
	}
	j := jsiiProxy_NestedStackSynthesizer{}

	_jsii_.Create(
		"aws-cdk-lib.NestedStackSynthesizer",
		[]interface{}{parentDeployment},
		&j,
	)

	return &j
}

func NewNestedStackSynthesizer_Override(n NestedStackSynthesizer, parentDeployment IStackSynthesizer) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.NestedStackSynthesizer",
		[]interface{}{parentDeployment},
		n,
	)
}

func (n *jsiiProxy_NestedStackSynthesizer) AddDockerImageAsset(asset *DockerImageAssetSource) *DockerImageAssetLocation {
	if err := n.validateAddDockerImageAssetParameters(asset); err != nil {
		panic(err)
	}
	var returns *DockerImageAssetLocation

	_jsii_.Invoke(
		n,
		"addDockerImageAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NestedStackSynthesizer) AddFileAsset(asset *FileAssetSource) *FileAssetLocation {
	if err := n.validateAddFileAssetParameters(asset); err != nil {
		panic(err)
	}
	var returns *FileAssetLocation

	_jsii_.Invoke(
		n,
		"addFileAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NestedStackSynthesizer) Bind(stack Stack) {
	if err := n.validateBindParameters(stack); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"bind",
		[]interface{}{stack},
	)
}

func (n *jsiiProxy_NestedStackSynthesizer) EmitStackArtifact(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions) {
	if err := n.validateEmitStackArtifactParameters(stack, session, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"emitStackArtifact",
		[]interface{}{stack, session, options},
	)
}

func (n *jsiiProxy_NestedStackSynthesizer) Synthesize(session ISynthesisSession) {
	if err := n.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"synthesize",
		[]interface{}{session},
	)
}

func (n *jsiiProxy_NestedStackSynthesizer) SynthesizeStackTemplate(stack Stack, session ISynthesisSession) {
	if err := n.validateSynthesizeStackTemplateParameters(stack, session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"synthesizeStackTemplate",
		[]interface{}{stack, session},
	)
}

