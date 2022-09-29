// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A synthesizer that uses conventional asset locations, but not conventional deployment roles.
//
// Instead of assuming the bootstrapped deployment roles, all stack operations will be performed
// using the CLI's current credentials.
//
// - This synthesizer does not support deploying to accounts to which the CLI does not have
//    credentials. It also does not support deploying using **CDK Pipelines**. For either of those
//    features, use `DefaultStackSynthesizer`.
// - This synthesizer requires an S3 bucket and ECR repository with well-known names. To
//    not depend on those, use `LegacyStackSynthesizer`.
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
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cliCredentialsStackSynthesizer := monocdk.NewCliCredentialsStackSynthesizer(&cliCredentialsStackSynthesizerProps{
//   	bucketPrefix: jsii.String("bucketPrefix"),
//   	dockerTagPrefix: jsii.String("dockerTagPrefix"),
//   	fileAssetsBucketName: jsii.String("fileAssetsBucketName"),
//   	imageAssetsRepositoryName: jsii.String("imageAssetsRepositoryName"),
//   	qualifier: jsii.String("qualifier"),
//   })
//
// Experimental.
type CliCredentialsStackSynthesizer interface {
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

// The jsii proxy struct for CliCredentialsStackSynthesizer
type jsiiProxy_CliCredentialsStackSynthesizer struct {
	jsiiProxy_StackSynthesizer
}

// Experimental.
func NewCliCredentialsStackSynthesizer(props *CliCredentialsStackSynthesizerProps) CliCredentialsStackSynthesizer {
	_init_.Initialize()

	if err := validateNewCliCredentialsStackSynthesizerParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CliCredentialsStackSynthesizer{}

	_jsii_.Create(
		"monocdk.CliCredentialsStackSynthesizer",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewCliCredentialsStackSynthesizer_Override(c CliCredentialsStackSynthesizer, props *CliCredentialsStackSynthesizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.CliCredentialsStackSynthesizer",
		[]interface{}{props},
		c,
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

