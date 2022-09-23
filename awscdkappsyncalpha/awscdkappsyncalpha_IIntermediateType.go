// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Intermediate Types are types that includes a certain set of fields that define the entirety of your schema.
// Experimental.
type IIntermediateType interface {
	// Add a field to this Intermediate Type.
	// Experimental.
	AddField(options *AddFieldOptions)
	// Create an GraphQL Type representing this Intermediate Type.
	// Experimental.
	Attribute(options *BaseTypeOptions) GraphqlType
	// Generate the string of this object type.
	// Experimental.
	ToString() *string
	// the attributes of this type.
	// Experimental.
	Definition() *map[string]IField
	// the directives for this object type.
	// Experimental.
	Directives() *[]Directive
	// The Interface Types this Intermediate Type implements.
	// Experimental.
	InterfaceTypes() *[]InterfaceType
	// the intermediate type linked to this attribute (i.e. an interface or an object).
	// Experimental.
	IntermediateType() IIntermediateType
	// the name of this type.
	// Experimental.
	Name() *string
	// The resolvers linked to this data source.
	// Experimental.
	Resolvers() *[]Resolver
	// Experimental.
	SetResolvers(r *[]Resolver)
}

// The jsii proxy for IIntermediateType
type jsiiProxy_IIntermediateType struct {
	_ byte // padding
}

func (i *jsiiProxy_IIntermediateType) AddField(options *AddFieldOptions) {
	if err := i.validateAddFieldParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addField",
		[]interface{}{options},
	)
}

func (i *jsiiProxy_IIntermediateType) Attribute(options *BaseTypeOptions) GraphqlType {
	if err := i.validateAttributeParameters(options); err != nil {
		panic(err)
	}
	var returns GraphqlType

	_jsii_.Invoke(
		i,
		"attribute",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IIntermediateType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IIntermediateType) Definition() *map[string]IField {
	var returns *map[string]IField
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIntermediateType) Directives() *[]Directive {
	var returns *[]Directive
	_jsii_.Get(
		j,
		"directives",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIntermediateType) InterfaceTypes() *[]InterfaceType {
	var returns *[]InterfaceType
	_jsii_.Get(
		j,
		"interfaceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIntermediateType) IntermediateType() IIntermediateType {
	var returns IIntermediateType
	_jsii_.Get(
		j,
		"intermediateType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIntermediateType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIntermediateType) Resolvers() *[]Resolver {
	var returns *[]Resolver
	_jsii_.Get(
		j,
		"resolvers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIntermediateType)SetResolvers(val *[]Resolver) {
	_jsii_.Set(
		j,
		"resolvers",
		val,
	)
}

