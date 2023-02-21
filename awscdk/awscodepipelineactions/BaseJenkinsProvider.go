package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipelineactions/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

type BaseJenkinsProvider interface {
	constructs.Construct
	IJenkinsProvider
	// The tree node.
	Node() constructs.Node
	ProviderName() *string
	ServerUrl() *string
	Version() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for BaseJenkinsProvider
type jsiiProxy_BaseJenkinsProvider struct {
	internal.Type__constructsConstruct
	jsiiProxy_IJenkinsProvider
}

func (j *jsiiProxy_BaseJenkinsProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseJenkinsProvider) ProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseJenkinsProvider) ServerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseJenkinsProvider) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


func NewBaseJenkinsProvider_Override(b BaseJenkinsProvider, scope constructs.Construct, id *string, version *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.BaseJenkinsProvider",
		[]interface{}{scope, id, version},
		b,
	)
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
func BaseJenkinsProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBaseJenkinsProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.BaseJenkinsProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseJenkinsProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

