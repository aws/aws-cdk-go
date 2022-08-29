// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkappsyncalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface Types are abstract types that includes a certain set of fields that other types must include if they implement the interface.
//
// Example:
//   node := appsync.NewInterfaceType(jsii.String("Node"), &intermediateTypeOptions{
//   	definition: map[string]iField{
//   		"id": appsync.GraphqlType.string(&BaseTypeOptions{
//   			"isRequired": jsii.Boolean(true),
//   		}),
//   	},
//   })
//   demo := appsync.NewObjectType(jsii.String("Demo"), &objectTypeOptions{
//   	interfaceTypes: []interfaceType{
//   		node,
//   	},
//   	definition: map[string]*iField{
//   		"version": appsync.GraphqlType.string(&BaseTypeOptions{
//   			"isRequired": jsii.Boolean(true),
//   		}),
//   	},
//   })
//
// Experimental.
type InterfaceType interface {
	IIntermediateType
	// the attributes of this type.
	// Experimental.
	Definition() *map[string]IField
	// the directives for this object type.
	// Experimental.
	Directives() *[]Directive
	// the authorization modes for this intermediate type.
	// Experimental.
	Modes() *[]AuthorizationType
	// Experimental.
	SetModes(val *[]AuthorizationType)
	// the name of this type.
	// Experimental.
	Name() *string
	// Add a field to this Interface Type.
	//
	// Interface Types must have both fieldName and field options.
	// Experimental.
	AddField(options *AddFieldOptions)
	// Create a GraphQL Type representing this Intermediate Type.
	// Experimental.
	Attribute(options *BaseTypeOptions) GraphqlType
	// Generate the string of this object type.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for InterfaceType
type jsiiProxy_InterfaceType struct {
	jsiiProxy_IIntermediateType
}

func (j *jsiiProxy_InterfaceType) Definition() *map[string]IField {
	var returns *map[string]IField
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InterfaceType) Directives() *[]Directive {
	var returns *[]Directive
	_jsii_.Get(
		j,
		"directives",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InterfaceType) Modes() *[]AuthorizationType {
	var returns *[]AuthorizationType
	_jsii_.Get(
		j,
		"modes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InterfaceType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Experimental.
func NewInterfaceType(name *string, props *IntermediateTypeOptions) InterfaceType {
	_init_.Initialize()

	j := jsiiProxy_InterfaceType{}

	_jsii_.Create(
		"@aws-cdk/aws-appsync-alpha.InterfaceType",
		[]interface{}{name, props},
		&j,
	)

	return &j
}

// Experimental.
func NewInterfaceType_Override(i InterfaceType, name *string, props *IntermediateTypeOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-appsync-alpha.InterfaceType",
		[]interface{}{name, props},
		i,
	)
}

func (j *jsiiProxy_InterfaceType) SetModes(val *[]AuthorizationType) {
	_jsii_.Set(
		j,
		"modes",
		val,
	)
}

func (i *jsiiProxy_InterfaceType) AddField(options *AddFieldOptions) {
	_jsii_.InvokeVoid(
		i,
		"addField",
		[]interface{}{options},
	)
}

func (i *jsiiProxy_InterfaceType) Attribute(options *BaseTypeOptions) GraphqlType {
	var returns GraphqlType

	_jsii_.Invoke(
		i,
		"attribute",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InterfaceType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

