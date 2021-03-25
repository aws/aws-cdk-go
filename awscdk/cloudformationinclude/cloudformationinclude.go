package cloudformationinclude

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/cloudformationinclude/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Construct to import an existing CloudFormation template file into a CDK application.
//
// All resources defined in the template file can be retrieved by calling the {@link getResource} method.
// Any modifications made on the returned resource objects will be reflected in the resulting CDK template.
// Experimental.
type CfnInclude interface {
	awscdk.CfnElement
	CreationStack() *[]*string
	LogicalId() *string
	Node() awscdk.ConstructNode
	Stack() awscdk.Stack
	GetCondition(conditionName *string) awscdk.CfnCondition
	GetHook(hookLogicalId *string) awscdk.CfnHook
	GetMapping(mappingName *string) awscdk.CfnMapping
	GetNestedStack(logicalId *string) *IncludedNestedStack
	GetOutput(logicalId *string) awscdk.CfnOutput
	GetParameter(parameterName *string) awscdk.CfnParameter
	GetResource(logicalId *string) awscdk.CfnResource
	GetRule(ruleName *string) awscdk.CfnRule
	LoadNestedStack(logicalId *string, nestedStackProps *CfnIncludeProps) *IncludedNestedStack
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_CfnInclude) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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


// Experimental.
func NewCfnInclude(scope constructs.Construct, id *string, props *CfnIncludeProps) CfnInclude {
	_init_.Initialize()

	j := jsiiProxy_CfnInclude{}

	_jsii_.Create(
		"monocdk.cloudformation_include.CfnInclude",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCfnInclude_Override(c CfnInclude, scope constructs.Construct, id *string, props *CfnIncludeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.cloudformation_include.CfnInclude",
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
// Experimental.
func CfnInclude_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.cloudformation_include.CfnInclude",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnInclude_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.cloudformation_include.CfnInclude",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns the CfnCondition object from the 'Conditions' section of the CloudFormation template with the given name.
//
// Any modifications performed on that object will be reflected in the resulting CDK template.
//
// If a Condition with the given name is not present in the template,
// throws an exception.
// Experimental.
func (c *jsiiProxy_CfnInclude) GetCondition(conditionName *string) awscdk.CfnCondition {
	var returns awscdk.CfnCondition

	_jsii_.Invoke(
		c,
		"getCondition",
		[]interface{}{conditionName},
		&returns,
	)

	return returns
}

// Returns the CfnHook object from the 'Hooks' section of the included CloudFormation template with the given logical ID.
//
// Any modifications performed on the returned object will be reflected in the resulting CDK template.
//
// If a Hook with the given logical ID is not present in the template,
// an exception will be thrown.
// Experimental.
func (c *jsiiProxy_CfnInclude) GetHook(hookLogicalId *string) awscdk.CfnHook {
	var returns awscdk.CfnHook

	_jsii_.Invoke(
		c,
		"getHook",
		[]interface{}{hookLogicalId},
		&returns,
	)

	return returns
}

// Returns the CfnMapping object from the 'Mappings' section of the included template.
//
// Any modifications performed on that object will be reflected in the resulting CDK template.
//
// If a Mapping with the given name is not present in the template,
// an exception will be thrown.
// Experimental.
func (c *jsiiProxy_CfnInclude) GetMapping(mappingName *string) awscdk.CfnMapping {
	var returns awscdk.CfnMapping

	_jsii_.Invoke(
		c,
		"getMapping",
		[]interface{}{mappingName},
		&returns,
	)

	return returns
}

// Returns a loaded NestedStack with name logicalId.
//
// For a nested stack to be returned by this method,
// it must be specified either in the {@link CfnIncludeProps.loadNestedStacks} property,
// or through the {@link loadNestedStack} method.
// Experimental.
func (c *jsiiProxy_CfnInclude) GetNestedStack(logicalId *string) *IncludedNestedStack {
	var returns *IncludedNestedStack

	_jsii_.Invoke(
		c,
		"getNestedStack",
		[]interface{}{logicalId},
		&returns,
	)

	return returns
}

// Returns the CfnOutput object from the 'Outputs' section of the included template.
//
// Any modifications performed on that object will be reflected in the resulting CDK template.
//
// If an Output with the given name is not present in the template,
// throws an exception.
// Experimental.
func (c *jsiiProxy_CfnInclude) GetOutput(logicalId *string) awscdk.CfnOutput {
	var returns awscdk.CfnOutput

	_jsii_.Invoke(
		c,
		"getOutput",
		[]interface{}{logicalId},
		&returns,
	)

	return returns
}

// Returns the CfnParameter object from the 'Parameters' section of the included template.
//
// Any modifications performed on that object will be reflected in the resulting CDK template.
//
// If a Parameter with the given name is not present in the template,
// throws an exception.
// Experimental.
func (c *jsiiProxy_CfnInclude) GetParameter(parameterName *string) awscdk.CfnParameter {
	var returns awscdk.CfnParameter

	_jsii_.Invoke(
		c,
		"getParameter",
		[]interface{}{parameterName},
		&returns,
	)

	return returns
}

// Returns the low-level CfnResource from the template with the given logical ID.
//
// Any modifications performed on that resource will be reflected in the resulting CDK template.
//
// The returned object will be of the proper underlying class;
// you can always cast it to the correct type in your code:
//
//      // assume the template contains an AWS::S3::Bucket with logical ID 'Bucket'
//      const cfnBucket = cfnTemplate.getResource('Bucket') as s3.CfnBucket;
//      // cfnBucket is of type s3.CfnBucket
//
// If the template does not contain a resource with the given logical ID,
// an exception will be thrown.
// Experimental.
func (c *jsiiProxy_CfnInclude) GetResource(logicalId *string) awscdk.CfnResource {
	var returns awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"getResource",
		[]interface{}{logicalId},
		&returns,
	)

	return returns
}

// Returns the CfnRule object from the 'Rules' section of the CloudFormation template with the given name.
//
// Any modifications performed on that object will be reflected in the resulting CDK template.
//
// If a Rule with the given name is not present in the template,
// an exception will be thrown.
// Experimental.
func (c *jsiiProxy_CfnInclude) GetRule(ruleName *string) awscdk.CfnRule {
	var returns awscdk.CfnRule

	_jsii_.Invoke(
		c,
		"getRule",
		[]interface{}{ruleName},
		&returns,
	)

	return returns
}

// Includes a template for a child stack inside of this parent template.
//
// A child with this logical ID must exist in the template,
// and be of type AWS::CloudFormation::Stack.
// This is equivalent to specifying the value in the {@link CfnIncludeProps.loadNestedStacks}
// property on object construction.
//
// Returns: the same {@link IncludedNestedStack} object that {@link getNestedStack} returns for this logical ID
// Experimental.
func (c *jsiiProxy_CfnInclude) LoadNestedStack(logicalId *string, nestedStackProps *CfnIncludeProps) *IncludedNestedStack {
	var returns *IncludedNestedStack

	_jsii_.Invoke(
		c,
		"loadNestedStack",
		[]interface{}{logicalId, nestedStackProps},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnInclude) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnInclude) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnInclude) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnInclude) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnInclude) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnInclude) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnInclude) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties of {@link CfnInclude}.
// Experimental.
type CfnIncludeProps struct {
	// Path to the template file.
	//
	// Both JSON and YAML template formats are supported.
	// Experimental.
	TemplateFile *string `json:"templateFile"`
	// Specifies the template files that define nested stacks that should be included.
	//
	// If your template specifies a stack that isn't included here, it won't be created as a NestedStack
	// resource, and it won't be accessible from the {@link CfnInclude.getNestedStack} method
	// (but will still be accessible from the {@link CfnInclude.getResource} method).
	//
	// If you include a stack here with an ID that isn't in the template,
	// or is in the template but is not a nested stack,
	// template creation will fail and an error will be thrown.
	// Experimental.
	LoadNestedStacks *map[string]*CfnIncludeProps `json:"loadNestedStacks"`
	// Specifies parameters to be replaced by the values in this mapping.
	//
	// Any parameters in the template that aren't specified here will be left unmodified.
	// If you include a parameter here with an ID that isn't in the template,
	// template creation will fail and an error will be thrown.
	// Experimental.
	Parameters *map[string]interface{} `json:"parameters"`
	// Whether the resources should have the same logical IDs in the resulting CDK template as they did in the original CloudFormation template file.
	//
	// If you're vending a Construct using an existing CloudFormation template,
	// make sure to pass this as `false`.
	//
	// **Note**: regardless of whether this option is true or false,
	// the {@link CfnInclude.getResource} and related methods always uses the original logical ID of the resource/element,
	// as specified in the template file.
	// Experimental.
	PreserveLogicalIds *bool `json:"preserveLogicalIds"`
}

// The type returned from {@link CfnInclude.getNestedStack}. Contains both the NestedStack object and CfnInclude representations of the child stack.
// Experimental.
type IncludedNestedStack struct {
	// The CfnInclude that represents the template, which can be used to access Resources and other template elements.
	// Experimental.
	IncludedTemplate CfnInclude `json:"includedTemplate"`
	// The NestedStack object which represents the scope of the template.
	// Experimental.
	Stack awscdk.NestedStack `json:"stack"`
}

