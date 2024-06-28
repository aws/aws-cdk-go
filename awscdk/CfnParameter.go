package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation parameter.
//
// Use the optional Parameters section to customize your templates.
// Parameters enable you to input custom values to your template each time you create or
// update a stack.
//
// Example:
//   myTopic := sns.NewTopic(this, jsii.String("MyTopic"))
//   url := awscdk.NewCfnParameter(this, jsii.String("url-param"))
//
//   myTopic.AddSubscription(subscriptions.NewUrlSubscription(url.valueAsString))
//
type CfnParameter interface {
	CfnElement
	// A regular expression that represents the patterns to allow for String types.
	// Default: - No constraints on patterns allowed for parameter.
	//
	AllowedPattern() *string
	SetAllowedPattern(val *string)
	// An array containing the list of values allowed for the parameter.
	// Default: - No constraints on values allowed for parameter.
	//
	AllowedValues() *[]*string
	SetAllowedValues(val *[]*string)
	// A string that explains a constraint when the constraint is violated.
	//
	// For example, without a constraint description, a parameter that has an allowed
	// pattern of [A-Za-z0-9]+ displays the following error message when the user specifies
	// an invalid value:.
	// Default: - No description with customized error message when user specifies invalid values.
	//
	ConstraintDescription() *string
	SetConstraintDescription(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A value of the appropriate type for the template to use if no value is specified when a stack is created.
	//
	// If you define constraints for the parameter, you must specify
	// a value that adheres to those constraints.
	// Default: - No default value for parameter.
	//
	Default() interface{}
	SetDefault(val interface{})
	// A string of up to 4000 characters that describes the parameter.
	// Default: - No description for the parameter.
	//
	Description() *string
	SetDescription(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// An integer value that determines the largest number of characters you want to allow for String types.
	// Default: - None.
	//
	MaxLength() *float64
	SetMaxLength(val *float64)
	// A numeric value that determines the largest numeric value you want to allow for Number types.
	// Default: - None.
	//
	MaxValue() *float64
	SetMaxValue(val *float64)
	// An integer value that determines the smallest number of characters you want to allow for String types.
	// Default: - None.
	//
	MinLength() *float64
	SetMinLength(val *float64)
	// A numeric value that determines the smallest numeric value you want to allow for Number types.
	// Default: - None.
	//
	MinValue() *float64
	SetMinValue(val *float64)
	// The tree node.
	Node() constructs.Node
	// Indicates if this parameter is configured with "NoEcho" enabled.
	NoEcho() *bool
	SetNoEcho(val *bool)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// The data type for the parameter (DataType).
	// Default: String.
	//
	Type() *string
	SetType(val *string)
	// The parameter value as a Token.
	Value() IResolvable
	// The parameter value, if it represents a string list.
	ValueAsList() *[]*string
	// The parameter value, if it represents a number.
	ValueAsNumber() *float64
	// The parameter value, if it represents a string.
	ValueAsString() *string
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	Resolve(_context IResolveContext) interface{}
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CfnParameter
type jsiiProxy_CfnParameter struct {
	jsiiProxy_CfnElement
}

func (j *jsiiProxy_CfnParameter) AllowedPattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) AllowedValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) ConstraintDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constraintDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) Default() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"default",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) MaxLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) MaxValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) MinLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) MinValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) NoEcho() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"noEcho",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) Value() IResolvable {
	var returns IResolvable
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) ValueAsList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valueAsList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) ValueAsNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueAsNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) ValueAsString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueAsString",
		&returns,
	)
	return returns
}


// Creates a parameter construct.
//
// Note that the name (logical ID) of the parameter will derive from it's `coname` and location
// within the stack. Therefore, it is recommended that parameters are defined at the stack level.
func NewCfnParameter(scope constructs.Construct, id *string, props *CfnParameterProps) CfnParameter {
	_init_.Initialize()

	if err := validateNewCfnParameterParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnParameter{}

	_jsii_.Create(
		"aws-cdk-lib.CfnParameter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates a parameter construct.
//
// Note that the name (logical ID) of the parameter will derive from it's `coname` and location
// within the stack. Therefore, it is recommended that parameters are defined at the stack level.
func NewCfnParameter_Override(c CfnParameter, scope constructs.Construct, id *string, props *CfnParameterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnParameter",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnParameter)SetAllowedPattern(val *string) {
	_jsii_.Set(
		j,
		"allowedPattern",
		val,
	)
}

func (j *jsiiProxy_CfnParameter)SetAllowedValues(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedValues",
		val,
	)
}

func (j *jsiiProxy_CfnParameter)SetConstraintDescription(val *string) {
	_jsii_.Set(
		j,
		"constraintDescription",
		val,
	)
}

func (j *jsiiProxy_CfnParameter)SetDefault(val interface{}) {
	if err := j.validateSetDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"default",
		val,
	)
}

func (j *jsiiProxy_CfnParameter)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnParameter)SetMaxLength(val *float64) {
	_jsii_.Set(
		j,
		"maxLength",
		val,
	)
}

func (j *jsiiProxy_CfnParameter)SetMaxValue(val *float64) {
	_jsii_.Set(
		j,
		"maxValue",
		val,
	)
}

func (j *jsiiProxy_CfnParameter)SetMinLength(val *float64) {
	_jsii_.Set(
		j,
		"minLength",
		val,
	)
}

func (j *jsiiProxy_CfnParameter)SetMinValue(val *float64) {
	_jsii_.Set(
		j,
		"minValue",
		val,
	)
}

func (j *jsiiProxy_CfnParameter)SetNoEcho(val *bool) {
	if err := j.validateSetNoEchoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noEcho",
		val,
	)
}

func (j *jsiiProxy_CfnParameter)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnParameter_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnParameter_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnParameter",
		"isCfnElement",
		[]interface{}{x},
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
func CfnParameter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnParameter_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnParameter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnParameter) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnParameter) Resolve(_context IResolveContext) interface{} {
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

func (c *jsiiProxy_CfnParameter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

