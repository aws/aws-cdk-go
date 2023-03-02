package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/constructs-go/constructs/v3"
)

// A class representing Jenkins providers.
//
// Example:
//   jenkinsProvider := codepipeline_actions.NewJenkinsProvider(this, jsii.String("JenkinsProvider"), &jenkinsProviderProps{
//   	providerName: jsii.String("MyJenkinsProvider"),
//   	serverUrl: jsii.String("http://my-jenkins.com:8080"),
//   	version: jsii.String("2"),
//   })
//
// See: #import.
//
// Experimental.
type JenkinsProvider interface {
	BaseJenkinsProvider
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Experimental.
	ProviderName() *string
	// Experimental.
	ServerUrl() *string
	// Experimental.
	Version() *string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for JenkinsProvider
type jsiiProxy_JenkinsProvider struct {
	jsiiProxy_BaseJenkinsProvider
}

func (j *jsiiProxy_JenkinsProvider) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JenkinsProvider) ProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JenkinsProvider) ServerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JenkinsProvider) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Experimental.
func NewJenkinsProvider(scope constructs.Construct, id *string, props *JenkinsProviderProps) JenkinsProvider {
	_init_.Initialize()

	if err := validateNewJenkinsProviderParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_JenkinsProvider{}

	_jsii_.Create(
		"monocdk.aws_codepipeline_actions.JenkinsProvider",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewJenkinsProvider_Override(j JenkinsProvider, scope constructs.Construct, id *string, props *JenkinsProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codepipeline_actions.JenkinsProvider",
		[]interface{}{scope, id, props},
		j,
	)
}

// Import a Jenkins provider registered either outside the CDK, or in a different CDK Stack.
//
// Returns: a new Construct representing a reference to an existing Jenkins provider.
// Experimental.
func JenkinsProvider_FromJenkinsProviderAttributes(scope constructs.Construct, id *string, attrs *JenkinsProviderAttributes) IJenkinsProvider {
	_init_.Initialize()

	if err := validateJenkinsProvider_FromJenkinsProviderAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IJenkinsProvider

	_jsii_.StaticInvoke(
		"monocdk.aws_codepipeline_actions.JenkinsProvider",
		"fromJenkinsProviderAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func JenkinsProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJenkinsProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codepipeline_actions.JenkinsProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JenkinsProvider) OnPrepare() {
	_jsii_.InvokeVoid(
		j,
		"onPrepare",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JenkinsProvider) OnSynthesize(session constructs.ISynthesisSession) {
	if err := j.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (j *jsiiProxy_JenkinsProvider) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JenkinsProvider) Prepare() {
	_jsii_.InvokeVoid(
		j,
		"prepare",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JenkinsProvider) Synthesize(session awscdk.ISynthesisSession) {
	if err := j.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"synthesize",
		[]interface{}{session},
	)
}

func (j *jsiiProxy_JenkinsProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JenkinsProvider) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

