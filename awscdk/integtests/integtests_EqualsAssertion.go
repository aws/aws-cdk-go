package integtests

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/integtests/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Construct that creates a CustomResource to assert that two values are equal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var actualResult actualResult
//   var expectedResult expectedResult
//
//   equalsAssertion := awscdk.Integ_tests.NewEqualsAssertion(this, jsii.String("MyEqualsAssertion"), &equalsAssertionProps{
//   	actual: actualResult,
//   	expected: expectedResult,
//
//   	// the properties below are optional
//   	failDeployment: jsii.Boolean(false),
//   })
//
// Experimental.
type EqualsAssertion interface {
	awscdk.Construct
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The result of the assertion.
	// Experimental.
	Result() *string
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

// The jsii proxy struct for EqualsAssertion
type jsiiProxy_EqualsAssertion struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_EqualsAssertion) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EqualsAssertion) Result() *string {
	var returns *string
	_jsii_.Get(
		j,
		"result",
		&returns,
	)
	return returns
}


// Experimental.
func NewEqualsAssertion(scope constructs.Construct, id *string, props *EqualsAssertionProps) EqualsAssertion {
	_init_.Initialize()

	if err := validateNewEqualsAssertionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EqualsAssertion{}

	_jsii_.Create(
		"monocdk.integ_tests.EqualsAssertion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEqualsAssertion_Override(e EqualsAssertion, scope constructs.Construct, id *string, props *EqualsAssertionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.integ_tests.EqualsAssertion",
		[]interface{}{scope, id, props},
		e,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func EqualsAssertion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEqualsAssertion_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.integ_tests.EqualsAssertion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EqualsAssertion) OnPrepare() {
	_jsii_.InvokeVoid(
		e,
		"onPrepare",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EqualsAssertion) OnSynthesize(session constructs.ISynthesisSession) {
	if err := e.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (e *jsiiProxy_EqualsAssertion) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EqualsAssertion) Prepare() {
	_jsii_.InvokeVoid(
		e,
		"prepare",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EqualsAssertion) Synthesize(session awscdk.ISynthesisSession) {
	if err := e.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"synthesize",
		[]interface{}{session},
	)
}

func (e *jsiiProxy_EqualsAssertion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EqualsAssertion) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

