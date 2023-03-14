package assertions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Suite of assertions that can be run on a CDK stack.
//
// Typically used, as part of unit tests, to validate that the rendered
// CloudFormation template has expected resources and properties.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stack := awscdk.NewStack()
//   // ...
//   template := awscdk.Template_FromStack(stack)
//
type Template interface {
	// Assert that all resources of the given type contain the given definition in the CloudFormation template.
	//
	// By default, performs partial matching on the resource, via the `Match.objectLike()`.
	// To configure different behavior, use other matchers in the `Match` class.
	AllResources(type_ *string, props interface{})
	// Assert that all resources of the given type contain the given properties CloudFormation template.
	//
	// By default, performs partial matching on the `Properties` key of the resource, via the
	// `Match.objectLike()`. To configure different behavior, use other matchers in the `Match` class.
	AllResourcesProperties(type_ *string, props interface{})
	// Get the set of matching Conditions that match the given properties in the CloudFormation template.
	FindConditions(logicalId *string, props interface{}) *map[string]*map[string]interface{}
	// Get the set of matching Mappings that match the given properties in the CloudFormation template.
	FindMappings(logicalId *string, props interface{}) *map[string]*map[string]interface{}
	// Get the set of matching Outputs that match the given properties in the CloudFormation template.
	FindOutputs(logicalId *string, props interface{}) *map[string]*map[string]interface{}
	// Get the set of matching Parameters that match the given properties in the CloudFormation template.
	FindParameters(logicalId *string, props interface{}) *map[string]*map[string]interface{}
	// Get the set of matching resources of a given type and properties in the CloudFormation template.
	FindResources(type_ *string, props interface{}) *map[string]*map[string]interface{}
	// Assert that a Condition with the given properties exists in the CloudFormation template.
	//
	// By default, performs partial matching on the resource, via the `Match.objectLike()`.
	// To configure different behavior, use other matchers in the `Match` class.
	HasCondition(logicalId *string, props interface{})
	// Assert that a Mapping with the given properties exists in the CloudFormation template.
	//
	// By default, performs partial matching on the resource, via the `Match.objectLike()`.
	// To configure different behavior, use other matchers in the `Match` class.
	HasMapping(logicalId *string, props interface{})
	// Assert that an Output with the given properties exists in the CloudFormation template.
	//
	// By default, performs partial matching on the resource, via the `Match.objectLike()`.
	// To configure different behavior, use other matchers in the `Match` class.
	HasOutput(logicalId *string, props interface{})
	// Assert that a Parameter with the given properties exists in the CloudFormation template.
	//
	// By default, performs partial matching on the parameter, via the `Match.objectLike()`.
	// To configure different behavior, use other matchers in the `Match` class.
	HasParameter(logicalId *string, props interface{})
	// Assert that a resource of the given type and given definition exists in the CloudFormation template.
	//
	// By default, performs partial matching on the resource, via the `Match.objectLike()`.
	// To configure different behavior, use other matchers in the `Match` class.
	HasResource(type_ *string, props interface{})
	// Assert that a resource of the given type and properties exists in the CloudFormation template.
	//
	// By default, performs partial matching on the `Properties` key of the resource, via the
	// `Match.objectLike()`. To configure different behavior, use other matchers in the `Match` class.
	HasResourceProperties(type_ *string, props interface{})
	// Assert that the given number of resources of the given type exist in the template.
	ResourceCountIs(type_ *string, count *float64)
	// Assert that the given number of resources of the given type and properties exists in the CloudFormation template.
	ResourcePropertiesCountIs(type_ *string, props interface{}, count *float64)
	// Assert that the CloudFormation template matches the given value.
	TemplateMatches(expected interface{})
	// The CloudFormation template deserialized into an object.
	ToJSON() *map[string]interface{}
}

// The jsii proxy struct for Template
type jsiiProxy_Template struct {
	_ byte // padding
}

// Base your assertions from an existing CloudFormation template formatted as an in-memory JSON object.
func Template_FromJSON(template *map[string]interface{}, templateParsingOptions *TemplateParsingOptions) Template {
	_init_.Initialize()

	if err := validateTemplate_FromJSONParameters(template, templateParsingOptions); err != nil {
		panic(err)
	}
	var returns Template

	_jsii_.StaticInvoke(
		"aws-cdk-lib.assertions.Template",
		"fromJSON",
		[]interface{}{template, templateParsingOptions},
		&returns,
	)

	return returns
}

// Base your assertions on the CloudFormation template synthesized by a CDK `Stack`.
func Template_FromStack(stack awscdk.Stack, templateParsingOptions *TemplateParsingOptions) Template {
	_init_.Initialize()

	if err := validateTemplate_FromStackParameters(stack, templateParsingOptions); err != nil {
		panic(err)
	}
	var returns Template

	_jsii_.StaticInvoke(
		"aws-cdk-lib.assertions.Template",
		"fromStack",
		[]interface{}{stack, templateParsingOptions},
		&returns,
	)

	return returns
}

// Base your assertions from an existing CloudFormation template formatted as a JSON string.
func Template_FromString(template *string, templateParsingOptions *TemplateParsingOptions) Template {
	_init_.Initialize()

	if err := validateTemplate_FromStringParameters(template, templateParsingOptions); err != nil {
		panic(err)
	}
	var returns Template

	_jsii_.StaticInvoke(
		"aws-cdk-lib.assertions.Template",
		"fromString",
		[]interface{}{template, templateParsingOptions},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Template) AllResources(type_ *string, props interface{}) {
	if err := t.validateAllResourcesParameters(type_, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"allResources",
		[]interface{}{type_, props},
	)
}

func (t *jsiiProxy_Template) AllResourcesProperties(type_ *string, props interface{}) {
	if err := t.validateAllResourcesPropertiesParameters(type_, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"allResourcesProperties",
		[]interface{}{type_, props},
	)
}

func (t *jsiiProxy_Template) FindConditions(logicalId *string, props interface{}) *map[string]*map[string]interface{} {
	if err := t.validateFindConditionsParameters(logicalId); err != nil {
		panic(err)
	}
	var returns *map[string]*map[string]interface{}

	_jsii_.Invoke(
		t,
		"findConditions",
		[]interface{}{logicalId, props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Template) FindMappings(logicalId *string, props interface{}) *map[string]*map[string]interface{} {
	if err := t.validateFindMappingsParameters(logicalId); err != nil {
		panic(err)
	}
	var returns *map[string]*map[string]interface{}

	_jsii_.Invoke(
		t,
		"findMappings",
		[]interface{}{logicalId, props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Template) FindOutputs(logicalId *string, props interface{}) *map[string]*map[string]interface{} {
	if err := t.validateFindOutputsParameters(logicalId); err != nil {
		panic(err)
	}
	var returns *map[string]*map[string]interface{}

	_jsii_.Invoke(
		t,
		"findOutputs",
		[]interface{}{logicalId, props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Template) FindParameters(logicalId *string, props interface{}) *map[string]*map[string]interface{} {
	if err := t.validateFindParametersParameters(logicalId); err != nil {
		panic(err)
	}
	var returns *map[string]*map[string]interface{}

	_jsii_.Invoke(
		t,
		"findParameters",
		[]interface{}{logicalId, props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Template) FindResources(type_ *string, props interface{}) *map[string]*map[string]interface{} {
	if err := t.validateFindResourcesParameters(type_); err != nil {
		panic(err)
	}
	var returns *map[string]*map[string]interface{}

	_jsii_.Invoke(
		t,
		"findResources",
		[]interface{}{type_, props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Template) HasCondition(logicalId *string, props interface{}) {
	if err := t.validateHasConditionParameters(logicalId, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"hasCondition",
		[]interface{}{logicalId, props},
	)
}

func (t *jsiiProxy_Template) HasMapping(logicalId *string, props interface{}) {
	if err := t.validateHasMappingParameters(logicalId, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"hasMapping",
		[]interface{}{logicalId, props},
	)
}

func (t *jsiiProxy_Template) HasOutput(logicalId *string, props interface{}) {
	if err := t.validateHasOutputParameters(logicalId, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"hasOutput",
		[]interface{}{logicalId, props},
	)
}

func (t *jsiiProxy_Template) HasParameter(logicalId *string, props interface{}) {
	if err := t.validateHasParameterParameters(logicalId, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"hasParameter",
		[]interface{}{logicalId, props},
	)
}

func (t *jsiiProxy_Template) HasResource(type_ *string, props interface{}) {
	if err := t.validateHasResourceParameters(type_, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"hasResource",
		[]interface{}{type_, props},
	)
}

func (t *jsiiProxy_Template) HasResourceProperties(type_ *string, props interface{}) {
	if err := t.validateHasResourcePropertiesParameters(type_, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"hasResourceProperties",
		[]interface{}{type_, props},
	)
}

func (t *jsiiProxy_Template) ResourceCountIs(type_ *string, count *float64) {
	if err := t.validateResourceCountIsParameters(type_, count); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"resourceCountIs",
		[]interface{}{type_, count},
	)
}

func (t *jsiiProxy_Template) ResourcePropertiesCountIs(type_ *string, props interface{}, count *float64) {
	if err := t.validateResourcePropertiesCountIsParameters(type_, props, count); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"resourcePropertiesCountIs",
		[]interface{}{type_, props, count},
	)
}

func (t *jsiiProxy_Template) TemplateMatches(expected interface{}) {
	if err := t.validateTemplateMatchesParameters(expected); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"templateMatches",
		[]interface{}{expected},
	)
}

func (t *jsiiProxy_Template) ToJSON() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

