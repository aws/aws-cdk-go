package cloudformationinclude

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/cloudformationinclude/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Construct to import an existing CloudFormation template file into a CDK application.
//
// All resources defined in the template file can be retrieved by calling the `getResource` method.
// Any modifications made on the returned resource objects will be reflected in the resulting CDK template.
//
// Example:
//   cfnTemplate := cfn_inc.NewCfnInclude(this, jsii.String("Template"), &CfnIncludeProps{
//   	TemplateFile: jsii.String("my-template.json"),
//   })
//
type CfnInclude interface {
	awscdk.CfnElement
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
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Returns the CfnCondition object from the 'Conditions' section of the CloudFormation template with the given name.
	//
	// Any modifications performed on that object will be reflected in the resulting CDK template.
	//
	// If a Condition with the given name is not present in the template,
	// throws an exception.
	GetCondition(conditionName *string) awscdk.CfnCondition
	// Returns the CfnHook object from the 'Hooks' section of the included CloudFormation template with the given logical ID.
	//
	// Any modifications performed on the returned object will be reflected in the resulting CDK template.
	//
	// If a Hook with the given logical ID is not present in the template,
	// an exception will be thrown.
	GetHook(hookLogicalId *string) awscdk.CfnHook
	// Returns the CfnMapping object from the 'Mappings' section of the included template.
	//
	// Any modifications performed on that object will be reflected in the resulting CDK template.
	//
	// If a Mapping with the given name is not present in the template,
	// an exception will be thrown.
	GetMapping(mappingName *string) awscdk.CfnMapping
	// Returns a loaded NestedStack with name logicalId.
	//
	// For a nested stack to be returned by this method,
	// it must be specified either in the `CfnIncludeProps.loadNestedStacks` property,
	// or through the `loadNestedStack` method.
	GetNestedStack(logicalId *string) *IncludedNestedStack
	// Returns the CfnOutput object from the 'Outputs' section of the included template.
	//
	// Any modifications performed on that object will be reflected in the resulting CDK template.
	//
	// If an Output with the given name is not present in the template,
	// throws an exception.
	GetOutput(logicalId *string) awscdk.CfnOutput
	// Returns the CfnParameter object from the 'Parameters' section of the included template.
	//
	// Any modifications performed on that object will be reflected in the resulting CDK template.
	//
	// If a Parameter with the given name is not present in the template,
	// throws an exception.
	GetParameter(parameterName *string) awscdk.CfnParameter
	// Returns the low-level CfnResource from the template with the given logical ID.
	//
	// Any modifications performed on that resource will be reflected in the resulting CDK template.
	//
	// The returned object will be of the proper underlying class;
	// you can always cast it to the correct type in your code:
	//
	//     // assume the template contains an AWS::S3::Bucket with logical ID 'Bucket'
	//     const cfnBucket = cfnTemplate.getResource('Bucket') as s3.CfnBucket;
	//     // cfnBucket is of type s3.CfnBucket
	//
	// If the template does not contain a resource with the given logical ID,
	// an exception will be thrown.
	GetResource(logicalId *string) awscdk.CfnResource
	// Returns the CfnRule object from the 'Rules' section of the CloudFormation template with the given name.
	//
	// Any modifications performed on that object will be reflected in the resulting CDK template.
	//
	// If a Rule with the given name is not present in the template,
	// an exception will be thrown.
	GetRule(ruleName *string) awscdk.CfnRule
	// Includes a template for a child stack inside of this parent template.
	//
	// A child with this logical ID must exist in the template,
	// and be of type AWS::CloudFormation::Stack.
	// This is equivalent to specifying the value in the `CfnIncludeProps.loadNestedStacks`
	// property on object construction.
	//
	// Returns: the same `IncludedNestedStack` object that `getNestedStack` returns for this logical ID.
	LoadNestedStack(logicalId *string, nestedStackProps *CfnIncludeProps) *IncludedNestedStack
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CfnInclude
type jsiiProxy_CfnInclude struct {
	internal.Type__awscdkCfnElement
}

func (j *jsiiProxy_CfnInclude) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInclude) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInclude) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInclude) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewCfnInclude(scope constructs.Construct, id *string, props *CfnIncludeProps) CfnInclude {
	_init_.Initialize()

	if err := validateNewCfnIncludeParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnInclude{}

	_jsii_.Create(
		"aws-cdk-lib.cloudformation_include.CfnInclude",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnInclude_Override(c CfnInclude, scope constructs.Construct, id *string, props *CfnIncludeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.cloudformation_include.CfnInclude",
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
func CfnInclude_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInclude_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.cloudformation_include.CfnInclude",
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
func CfnInclude_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInclude_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.cloudformation_include.CfnInclude",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInclude) GetCondition(conditionName *string) awscdk.CfnCondition {
	if err := c.validateGetConditionParameters(conditionName); err != nil {
		panic(err)
	}
	var returns awscdk.CfnCondition

	_jsii_.Invoke(
		c,
		"getCondition",
		[]interface{}{conditionName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInclude) GetHook(hookLogicalId *string) awscdk.CfnHook {
	if err := c.validateGetHookParameters(hookLogicalId); err != nil {
		panic(err)
	}
	var returns awscdk.CfnHook

	_jsii_.Invoke(
		c,
		"getHook",
		[]interface{}{hookLogicalId},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInclude) GetMapping(mappingName *string) awscdk.CfnMapping {
	if err := c.validateGetMappingParameters(mappingName); err != nil {
		panic(err)
	}
	var returns awscdk.CfnMapping

	_jsii_.Invoke(
		c,
		"getMapping",
		[]interface{}{mappingName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInclude) GetNestedStack(logicalId *string) *IncludedNestedStack {
	if err := c.validateGetNestedStackParameters(logicalId); err != nil {
		panic(err)
	}
	var returns *IncludedNestedStack

	_jsii_.Invoke(
		c,
		"getNestedStack",
		[]interface{}{logicalId},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInclude) GetOutput(logicalId *string) awscdk.CfnOutput {
	if err := c.validateGetOutputParameters(logicalId); err != nil {
		panic(err)
	}
	var returns awscdk.CfnOutput

	_jsii_.Invoke(
		c,
		"getOutput",
		[]interface{}{logicalId},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInclude) GetParameter(parameterName *string) awscdk.CfnParameter {
	if err := c.validateGetParameterParameters(parameterName); err != nil {
		panic(err)
	}
	var returns awscdk.CfnParameter

	_jsii_.Invoke(
		c,
		"getParameter",
		[]interface{}{parameterName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInclude) GetResource(logicalId *string) awscdk.CfnResource {
	if err := c.validateGetResourceParameters(logicalId); err != nil {
		panic(err)
	}
	var returns awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"getResource",
		[]interface{}{logicalId},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInclude) GetRule(ruleName *string) awscdk.CfnRule {
	if err := c.validateGetRuleParameters(ruleName); err != nil {
		panic(err)
	}
	var returns awscdk.CfnRule

	_jsii_.Invoke(
		c,
		"getRule",
		[]interface{}{ruleName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInclude) LoadNestedStack(logicalId *string, nestedStackProps *CfnIncludeProps) *IncludedNestedStack {
	if err := c.validateLoadNestedStackParameters(logicalId, nestedStackProps); err != nil {
		panic(err)
	}
	var returns *IncludedNestedStack

	_jsii_.Invoke(
		c,
		"loadNestedStack",
		[]interface{}{logicalId, nestedStackProps},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInclude) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnInclude) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

