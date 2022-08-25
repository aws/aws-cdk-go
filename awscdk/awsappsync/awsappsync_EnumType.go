package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Enum Types are abstract types that includes a set of fields that represent the strings this type can create.
//
// Example:
//   var api graphqlApi
//
//   episode := appsync.NewEnumType(jsii.String("Episode"), &enumTypeOptions{
//   	definition: []*string{
//   		jsii.String("NEWHOPE"),
//   		jsii.String("EMPIRE"),
//   		jsii.String("JEDI"),
//   	},
//   })
//   api.addType(episode)
//
// Experimental.
type EnumType interface {
	IIntermediateType
	// the attributes of this type.
	// Experimental.
	Definition() *map[string]IField
	// the authorization modes for this intermediate type.
	// Experimental.
	Modes() *[]AuthorizationType
	// Experimental.
	SetModes(val *[]AuthorizationType)
	// the name of this type.
	// Experimental.
	Name() *string
	// Add a field to this Enum Type.
	//
	// To add a field to this Enum Type, you must only configure
	// addField with the fieldName options.
	// Experimental.
	AddField(options *AddFieldOptions)
	// Create an GraphQL Type representing this Enum Type.
	// Experimental.
	Attribute(options *BaseTypeOptions) GraphqlType
	// Generate the string of this enum type.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EnumType
type jsiiProxy_EnumType struct {
	jsiiProxy_IIntermediateType
}

func (j *jsiiProxy_EnumType) Definition() *map[string]IField {
	var returns *map[string]IField
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnumType) Modes() *[]AuthorizationType {
	var returns *[]AuthorizationType
	_jsii_.Get(
		j,
		"modes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnumType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Experimental.
func NewEnumType(name *string, options *EnumTypeOptions) EnumType {
	_init_.Initialize()

	j := jsiiProxy_EnumType{}

	_jsii_.Create(
		"monocdk.aws_appsync.EnumType",
		[]interface{}{name, options},
		&j,
	)

	return &j
}

// Experimental.
func NewEnumType_Override(e EnumType, name *string, options *EnumTypeOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.EnumType",
		[]interface{}{name, options},
		e,
	)
}

func (j *jsiiProxy_EnumType) SetModes(val *[]AuthorizationType) {
	_jsii_.Set(
		j,
		"modes",
		val,
	)
}

func (e *jsiiProxy_EnumType) AddField(options *AddFieldOptions) {
	_jsii_.InvokeVoid(
		e,
		"addField",
		[]interface{}{options},
	)
}

func (e *jsiiProxy_EnumType) Attribute(options *BaseTypeOptions) GraphqlType {
	var returns GraphqlType

	_jsii_.Invoke(
		e,
		"attribute",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnumType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

