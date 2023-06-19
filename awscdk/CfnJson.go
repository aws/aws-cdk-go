package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v3"
)

// Captures a synthesis-time JSON object a CloudFormation reference which resolves during deployment to the resolved values of the JSON object.
//
// The main use case for this is to overcome a limitation in CloudFormation that
// does not allow using intrinsic functions as dictionary keys (because
// dictionary keys in JSON must be strings). Specifically this is common in IAM
// conditions such as `StringEquals: { lhs: "rhs" }` where you want "lhs" to be
// a reference.
//
// This object is resolvable, so it can be used as a value.
//
// This construct is backed by a custom resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var value interface{}
//
//   cfnJson := monocdk.NewCfnJson(this, jsii.String("MyCfnJson"), &CfnJsonProps{
//   	Value: value,
//   })
//
// Experimental.
type CfnJson interface {
	Construct
	IResolvable
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// This may return an array with a single informational element indicating how
	// to get this property populated, if it was skipped for performance reasons.
	// Experimental.
	CreationStack() *[]*string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() ConstructNode
	// An Fn::GetAtt to the JSON object passed through `value` and resolved during synthesis.
	//
	// Normally there is no need to use this property since `CfnJson` is an
	// IResolvable, so it can be simply used as a value.
	// Experimental.
	Value() Reference
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_arg IResolveContext) interface{}
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session ISynthesisSession)
	// This is required in case someone JSON.stringifys an object which refrences this object. Otherwise, we'll get a cyclic JSON reference.
	// Experimental.
	ToJSON() *string
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

// The jsii proxy struct for CfnJson
type jsiiProxy_CfnJson struct {
	jsiiProxy_Construct
	jsiiProxy_IResolvable
}

func (j *jsiiProxy_CfnJson) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJson) Node() ConstructNode {
	var returns ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJson) Value() Reference {
	var returns Reference
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func NewCfnJson(scope constructs.Construct, id *string, props *CfnJsonProps) CfnJson {
	_init_.Initialize()

	if err := validateNewCfnJsonParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnJson{}

	_jsii_.Create(
		"monocdk.CfnJson",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCfnJson_Override(c CfnJson, scope constructs.Construct, id *string, props *CfnJsonProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.CfnJson",
		[]interface{}{scope, id, props},
		c,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func CfnJson_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnJson_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.CfnJson",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJson) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnJson) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnJson) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJson) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnJson) Resolve(_arg IResolveContext) interface{} {
	if err := c.validateResolveParameters(_arg); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_arg},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJson) Synthesize(session ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnJson) ToJSON() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJson) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJson) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

