// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkappsyncalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Union Types are abstract types that are similar to Interface Types, but they cannot to specify any common fields between types.
//
// Note that fields of a union type need to be object types. In other words,
// you can't create a union type out of interfaces, other unions, or inputs.
//
// Example:
//   var api graphqlApi
//
//   string := appsync.graphqlType.string()
//   human := appsync.NewObjectType(jsii.String("Human"), &objectTypeOptions{
//   	definition: map[string]iField{
//   		"name": string,
//   	},
//   })
//   droid := appsync.NewObjectType(jsii.String("Droid"), &objectTypeOptions{
//   	definition: map[string]*iField{
//   		"name": string,
//   	},
//   })
//   starship := appsync.NewObjectType(jsii.String("Starship"), &objectTypeOptions{
//   	definition: map[string]*iField{
//   		"name": string,
//   	},
//   })
//   search := appsync.NewUnionType(jsii.String("Search"), &unionTypeOptions{
//   	definition: []iIntermediateType{
//   		human,
//   		droid,
//   		starship,
//   	},
//   })
//   api.addType(search)
//
// Experimental.
type UnionType interface {
	IIntermediateType
	// the attributes of this type.
	// Experimental.
	Definition() *map[string]IField
	// the authorization modes supported by this intermediate type.
	// Experimental.
	Modes() *[]AuthorizationType
	// Experimental.
	SetModes(val *[]AuthorizationType)
	// the name of this type.
	// Experimental.
	Name() *string
	// Add a field to this Union Type.
	//
	// Input Types must have field options and the IField must be an Object Type.
	// Experimental.
	AddField(options *AddFieldOptions)
	// Create a GraphQL Type representing this Union Type.
	// Experimental.
	Attribute(options *BaseTypeOptions) GraphqlType
	// Generate the string of this Union type.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for UnionType
type jsiiProxy_UnionType struct {
	jsiiProxy_IIntermediateType
}

func (j *jsiiProxy_UnionType) Definition() *map[string]IField {
	var returns *map[string]IField
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UnionType) Modes() *[]AuthorizationType {
	var returns *[]AuthorizationType
	_jsii_.Get(
		j,
		"modes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UnionType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Experimental.
func NewUnionType(name *string, options *UnionTypeOptions) UnionType {
	_init_.Initialize()

	if err := validateNewUnionTypeParameters(name, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_UnionType{}

	_jsii_.Create(
		"@aws-cdk/aws-appsync-alpha.UnionType",
		[]interface{}{name, options},
		&j,
	)

	return &j
}

// Experimental.
func NewUnionType_Override(u UnionType, name *string, options *UnionTypeOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-appsync-alpha.UnionType",
		[]interface{}{name, options},
		u,
	)
}

func (j *jsiiProxy_UnionType)SetModes(val *[]AuthorizationType) {
	_jsii_.Set(
		j,
		"modes",
		val,
	)
}

func (u *jsiiProxy_UnionType) AddField(options *AddFieldOptions) {
	if err := u.validateAddFieldParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"addField",
		[]interface{}{options},
	)
}

func (u *jsiiProxy_UnionType) Attribute(options *BaseTypeOptions) GraphqlType {
	if err := u.validateAttributeParameters(options); err != nil {
		panic(err)
	}
	var returns GraphqlType

	_jsii_.Invoke(
		u,
		"attribute",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UnionType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

