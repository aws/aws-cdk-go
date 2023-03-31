package awsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A Graphql Field.
// Experimental.
type IField interface {
	// Generate the arguments for this field.
	// Experimental.
	ArgsToString() *string
	// Generate the directives for this field.
	// Experimental.
	DirectivesToString(modes *[]AuthorizationType) *string
	// Generate the string for this attribute.
	// Experimental.
	ToString() *string
	// The options to make this field resolvable.
	// Experimental.
	FieldOptions() *ResolvableFieldOptions
	// the intermediate type linked to this attribute (i.e. an interface or an object).
	// Experimental.
	IntermediateType() IIntermediateType
	// property determining if this attribute is a list i.e. if true, attribute would be `[Type]`.
	// Experimental.
	IsList() *bool
	// property determining if this attribute is non-nullable i.e. if true, attribute would be `Type!` and this attribute must always have a value.
	// Experimental.
	IsRequired() *bool
	// property determining if this attribute is a non-nullable list i.e. if true, attribute would be `[ Type ]!` and this attribute's list must always have a value.
	// Experimental.
	IsRequiredList() *bool
	// the type of attribute.
	// Experimental.
	Type() Type
}

// The jsii proxy for IField
type jsiiProxy_IField struct {
	_ byte // padding
}

func (i *jsiiProxy_IField) ArgsToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"argsToString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IField) DirectivesToString(modes *[]AuthorizationType) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"directivesToString",
		[]interface{}{modes},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IField) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IField) FieldOptions() *ResolvableFieldOptions {
	var returns *ResolvableFieldOptions
	_jsii_.Get(
		j,
		"fieldOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IField) IntermediateType() IIntermediateType {
	var returns IIntermediateType
	_jsii_.Get(
		j,
		"intermediateType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IField) IsList() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IField) IsRequired() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IField) IsRequiredList() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isRequiredList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IField) Type() Type {
	var returns Type
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

