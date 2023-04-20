// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Represents the building block of the construct graph.
//
// All constructs besides the root construct must be created within the scope of
// another construct.
//
// Example:
//   entry := "/path/to/function"
//   image := awscdk.DockerImage_FromBuild(entry)
//
//   lambda.NewPythonFunction(this, jsii.String("function"), &PythonFunctionProps{
//   	Entry: jsii.String(Entry),
//   	Runtime: awscdk.Runtime_PYTHON_3_8(),
//   	Bundling: &BundlingOptions{
//   		BuildArgs: map[string]*string{
//   			"PIP_INDEX_URL": jsii.String("https://your.index.url/simple/"),
//   			"PIP_EXTRA_INDEX_URL": jsii.String("https://your.extra-index.url/simple/"),
//   		},
//   	},
//   })
//
// Experimental.
type Construct interface {
	constructs.Construct
	IConstruct
	// The construct tree node associated with this construct.
	// Experimental.
	Node() ConstructNode
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
	Synthesize(session ISynthesisSession)
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

// The jsii proxy struct for Construct
type jsiiProxy_Construct struct {
	internal.Type__constructsConstruct
	jsiiProxy_IConstruct
}

func (j *jsiiProxy_Construct) Node() ConstructNode {
	var returns ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewConstruct(scope constructs.Construct, id *string) Construct {
	_init_.Initialize()

	if err := validateNewConstructParameters(scope, id); err != nil {
		panic(err)
	}
	j := jsiiProxy_Construct{}

	_jsii_.Create(
		"monocdk.Construct",
		[]interface{}{scope, id},
		&j,
	)

	return &j
}

// Experimental.
func NewConstruct_Override(c Construct, scope constructs.Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.Construct",
		[]interface{}{scope, id},
		c,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func Construct_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConstruct_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.Construct",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Construct) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Construct) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_Construct) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Construct) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Construct) Synthesize(session ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_Construct) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Construct) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

