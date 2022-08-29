package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// A class representing Jenkins providers.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   jenkinsProvider := codepipeline_actions.NewJenkinsProvider(this, jsii.String("JenkinsProvider"), &jenkinsProviderProps{
//   	providerName: jsii.String("MyJenkinsProvider"),
//   	serverUrl: jsii.String("http://my-jenkins.com:8080"),
//   	version: jsii.String("2"),
//   })
//
// See: #import.
//
type JenkinsProvider interface {
	BaseJenkinsProvider
	// The tree node.
	Node() constructs.Node
	ProviderName() *string
	ServerUrl() *string
	Version() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for JenkinsProvider
type jsiiProxy_JenkinsProvider struct {
	jsiiProxy_BaseJenkinsProvider
}

func (j *jsiiProxy_JenkinsProvider) Node() constructs.Node {
	var returns constructs.Node
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


func NewJenkinsProvider(scope constructs.Construct, id *string, props *JenkinsProviderProps) JenkinsProvider {
	_init_.Initialize()

	j := jsiiProxy_JenkinsProvider{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.JenkinsProvider",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewJenkinsProvider_Override(j JenkinsProvider, scope constructs.Construct, id *string, props *JenkinsProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.JenkinsProvider",
		[]interface{}{scope, id, props},
		j,
	)
}

// Import a Jenkins provider registered either outside the CDK, or in a different CDK Stack.
//
// Returns: a new Construct representing a reference to an existing Jenkins provider.
func JenkinsProvider_FromJenkinsProviderAttributes(scope constructs.Construct, id *string, attrs *JenkinsProviderAttributes) IJenkinsProvider {
	_init_.Initialize()

	var returns IJenkinsProvider

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.JenkinsProvider",
		"fromJenkinsProviderAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func JenkinsProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.JenkinsProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
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

