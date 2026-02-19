package appstagingsynthesizeralpha

import (
	_init_ "github.com/aws/aws-cdk-go/appstagingsynthesizeralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/appstagingsynthesizeralpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// This is a dummy construct meant to signify that a stack is utilizing the AppStagingSynthesizer.
//
// It does not do anything, and is not meant
// to be created on its own. This construct will be a part of the
// construct tree only and not the Cfn template. The construct tree is
// then encoded in the AWS::CDK::Metadata resource of the stack and
// injested in our metrics like every other construct.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import app_staging_synthesizer_alpha "github.com/aws/aws-cdk-go/appstagingsynthesizeralpha"
//
//   usingAppStagingSynthesizer := app_staging_synthesizer_alpha.NewUsingAppStagingSynthesizer(this, jsii.String("MyUsingAppStagingSynthesizer"))
//
// Experimental.
type UsingAppStagingSynthesizer interface {
	constructs.Construct
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	// Experimental.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for UsingAppStagingSynthesizer
type jsiiProxy_UsingAppStagingSynthesizer struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_UsingAppStagingSynthesizer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewUsingAppStagingSynthesizer(scope constructs.Construct, id *string) UsingAppStagingSynthesizer {
	_init_.Initialize()

	if err := validateNewUsingAppStagingSynthesizerParameters(scope, id); err != nil {
		panic(err)
	}
	j := jsiiProxy_UsingAppStagingSynthesizer{}

	_jsii_.Create(
		"@aws-cdk/app-staging-synthesizer-alpha.UsingAppStagingSynthesizer",
		[]interface{}{scope, id},
		&j,
	)

	return &j
}

// Experimental.
func NewUsingAppStagingSynthesizer_Override(u UsingAppStagingSynthesizer, scope constructs.Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/app-staging-synthesizer-alpha.UsingAppStagingSynthesizer",
		[]interface{}{scope, id},
		u,
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
// Experimental.
func UsingAppStagingSynthesizer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUsingAppStagingSynthesizer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/app-staging-synthesizer-alpha.UsingAppStagingSynthesizer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UsingAppStagingSynthesizer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UsingAppStagingSynthesizer) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		u,
		"with",
		args,
		&returns,
	)

	return returns
}

