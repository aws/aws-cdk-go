package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
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
//   tagParam := awscdk.NewCfnParameter(this, jsii.String("TagName"))
//
//   stringEquals := awscdk.NewCfnJson(this, jsii.String("ConditionJson"), &CfnJsonProps{
//   	Value: map[string]*bool{
//   		fmt.Sprintf("aws:PrincipalTag/%v", tagParam.valueAsString): jsii.Boolean(true),
//   	},
//   })
//
//   principal := iam.NewAccountRootPrincipal().WithConditions(map[string]interface{}{
//   	"StringEquals": stringEquals,
//   })
//
//   iam.NewRole(this, jsii.String("MyRole"), &RoleProps{
//   	AssumedBy: principal,
//   })
//
type CfnJson interface {
	constructs.Construct
	IResolvable
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// This may return an array with a single informational element indicating how
	// to get this property populated, if it was skipped for performance reasons.
	CreationStack() *[]*string
	// The tree node.
	Node() constructs.Node
	// An Fn::GetAtt to the JSON object passed through `value` and resolved during synthesis.
	//
	// Normally there is no need to use this property since `CfnJson` is an
	// IResolvable, so it can be simply used as a value.
	Value() Reference
	// Produce the Token's value at resolution time.
	Resolve(_context IResolveContext) interface{}
	// This is required in case someone JSON.stringifys an object which references this object. Otherwise, we'll get a cyclic JSON reference.
	ToJSON() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CfnJson
type jsiiProxy_CfnJson struct {
	internal.Type__constructsConstruct
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

func (j *jsiiProxy_CfnJson) Node() constructs.Node {
	var returns constructs.Node
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


func NewCfnJson(scope constructs.Construct, id *string, props *CfnJsonProps) CfnJson {
	_init_.Initialize()

	if err := validateNewCfnJsonParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnJson{}

	_jsii_.Create(
		"aws-cdk-lib.CfnJson",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnJson_Override(c CfnJson, scope constructs.Construct, id *string, props *CfnJsonProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnJson",
		[]interface{}{scope, id, props},
		c,
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
func CfnJson_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnJson_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnJson",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJson) Resolve(_context IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
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

