// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkappsyncalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Input Types are abstract types that define complex objects.
//
// They are used in arguments to represent.
//
// Example:
//   var api graphqlApi
//
//   review := appsync.NewInputType(jsii.String("Review"), &intermediateTypeOptions{
//   	definition: map[string]iField{
//   		"stars": appsync.GraphqlType.int(&BaseTypeOptions{
//   			"isRequired": jsii.Boolean(true),
//   		}),
//   		"commentary": appsync.GraphqlType.string(),
//   	},
//   })
//   api.addType(review)
//
// Experimental.
type InputType interface {
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
	// Add a field to this Input Type.
	//
	// Input Types must have both fieldName and field options.
	// Experimental.
	AddField(options *AddFieldOptions)
	// Create a GraphQL Type representing this Input Type.
	// Experimental.
	Attribute(options *BaseTypeOptions) GraphqlType
	// Generate the string of this input type.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for InputType
type jsiiProxy_InputType struct {
	jsiiProxy_IIntermediateType
}

func (j *jsiiProxy_InputType) Definition() *map[string]IField {
	var returns *map[string]IField
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InputType) Modes() *[]AuthorizationType {
	var returns *[]AuthorizationType
	_jsii_.Get(
		j,
		"modes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InputType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Experimental.
func NewInputType(name *string, props *IntermediateTypeOptions) InputType {
	_init_.Initialize()

	if err := validateNewInputTypeParameters(name, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_InputType{}

	_jsii_.Create(
		"@aws-cdk/aws-appsync-alpha.InputType",
		[]interface{}{name, props},
		&j,
	)

	return &j
}

// Experimental.
func NewInputType_Override(i InputType, name *string, props *IntermediateTypeOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-appsync-alpha.InputType",
		[]interface{}{name, props},
		i,
	)
}

func (j *jsiiProxy_InputType)SetModes(val *[]AuthorizationType) {
	_jsii_.Set(
		j,
		"modes",
		val,
	)
}

func (i *jsiiProxy_InputType) AddField(options *AddFieldOptions) {
	if err := i.validateAddFieldParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addField",
		[]interface{}{options},
	)
}

func (i *jsiiProxy_InputType) Attribute(options *BaseTypeOptions) GraphqlType {
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

func (i *jsiiProxy_InputType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

