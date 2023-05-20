package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a CloudFormation condition, for resources which must be conditionally created and the determination must be made at deploy time.
//
// Example:
//   rawBucket := s3.NewCfnBucket(this, jsii.String("Bucket"), &CfnBucketProps{
//   })
//   // -or-
//   rawBucketAlt := myBucket.Node.defaultChild.(cfnBucket)
//
//   // then
//   rawBucket.CfnOptions.Condition = awscdk.NewCfnCondition(this, jsii.String("EnableBucket"), &CfnConditionProps{
//   })
//   rawBucket.CfnOptions.Metadata = map[string]interface{}{
//   	"metadataKey": jsii.String("MetadataValue"),
//   }
//
type CfnCondition interface {
	CfnElement
	ICfnConditionExpression
	IResolvable
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The condition statement.
	Expression() ICfnConditionExpression
	SetExpression(val ICfnConditionExpression)
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
	// The tree node.
	Node() constructs.Node
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Synthesizes the condition.
	Resolve(_context IResolveContext) interface{}
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CfnCondition
type jsiiProxy_CfnCondition struct {
	jsiiProxy_CfnElement
	jsiiProxy_ICfnConditionExpression
	jsiiProxy_IResolvable
}

func (j *jsiiProxy_CfnCondition) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCondition) Expression() ICfnConditionExpression {
	var returns ICfnConditionExpression
	_jsii_.Get(
		j,
		"expression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCondition) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCondition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCondition) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Build a new condition.
//
// The condition must be constructed with a condition token,
// that the condition is based on.
func NewCfnCondition(scope constructs.Construct, id *string, props *CfnConditionProps) CfnCondition {
	_init_.Initialize()

	if err := validateNewCfnConditionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCondition{}

	_jsii_.Create(
		"aws-cdk-lib.CfnCondition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Build a new condition.
//
// The condition must be constructed with a condition token,
// that the condition is based on.
func NewCfnCondition_Override(c CfnCondition, scope constructs.Construct, id *string, props *CfnConditionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnCondition",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCondition)SetExpression(val ICfnConditionExpression) {
	_jsii_.Set(
		j,
		"expression",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnCondition_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCondition_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnCondition",
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
func CfnCondition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCondition_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnCondition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCondition) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCondition) Resolve(_context IResolveContext) interface{} {
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

func (c *jsiiProxy_CfnCondition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

