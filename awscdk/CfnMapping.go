package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a CloudFormation mapping.
//
// Example:
//   regionTable := awscdk.NewCfnMapping(this, jsii.String("RegionTable"), &CfnMappingProps{
//   	Mapping: map[string]map[string]interface{}{
//   		"us-east-1": map[string]interface{}{
//   			"regionName": jsii.String("US East (N. Virginia)"),
//   		},
//   		"us-east-2": map[string]interface{}{
//   			"regionName": jsii.String("US East (Ohio)"),
//   		},
//   	},
//   })
//
//   regionTable.FindInMap(awscdk.Aws_REGION(), jsii.String("regionName"))
//
type CfnMapping interface {
	CfnRefElement
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// Returns: A reference to a value in the map based on the two keys.
	// If mapping is lazy, the value from the map or default value is returned instead of the reference and the mapping is not rendered in the template.
	FindInMap(key1 *string, key2 *string, defaultValue *string) *string
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Sets a value in the map based on the two keys.
	SetValue(key1 *string, key2 *string, value interface{})
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CfnMapping
type jsiiProxy_CfnMapping struct {
	jsiiProxy_CfnRefElement
}

func (j *jsiiProxy_CfnMapping) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMapping) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMapping) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMapping) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMapping) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewCfnMapping(scope constructs.Construct, id *string, props *CfnMappingProps) CfnMapping {
	_init_.Initialize()

	if err := validateNewCfnMappingParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMapping{}

	_jsii_.Create(
		"aws-cdk-lib.CfnMapping",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnMapping_Override(c CfnMapping, scope constructs.Construct, id *string, props *CfnMappingProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnMapping",
		[]interface{}{scope, id, props},
		c,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnMapping_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMapping_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnMapping",
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
func CfnMapping_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMapping_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnMapping",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMapping) FindInMap(key1 *string, key2 *string, defaultValue *string) *string {
	if err := c.validateFindInMapParameters(key1, key2); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"findInMap",
		[]interface{}{key1, key2, defaultValue},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMapping) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnMapping) SetValue(key1 *string, key2 *string, value interface{}) {
	if err := c.validateSetValueParameters(key1, key2, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"setValue",
		[]interface{}{key1, key2, value},
	)
}

func (c *jsiiProxy_CfnMapping) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

