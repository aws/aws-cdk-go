package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Object Types are types declared by you.
//
// Example:
//   var api graphqlApi
//   var dummyRequest mappingTemplate
//   var dummyResponse mappingTemplate
//
//   info := appsync.NewObjectType(jsii.String("Info"), &objectTypeOptions{
//   	definition: map[string]iField{
//   		"node": appsync.NewResolvableField(&ResolvableFieldOptions{
//   			"returnType": appsync.GraphqlType.string(),
//   			"args": map[string]GraphqlType{
//   				"id": appsync.GraphqlType.string(),
//   			},
//   			"dataSource": api.addNoneDataSource(jsii.String("none")),
//   			"requestMappingTemplate": dummyRequest,
//   			"responseMappingTemplate": dummyResponse,
//   		}),
//   	},
//   })
//
// Experimental.
type ObjectType interface {
	InterfaceType
	IIntermediateType
	// the attributes of this type.
	// Experimental.
	Definition() *map[string]IField
	// the directives for this object type.
	// Experimental.
	Directives() *[]Directive
	// The Interface Types this Object Type implements.
	// Experimental.
	InterfaceTypes() *[]InterfaceType
	// the authorization modes for this intermediate type.
	// Experimental.
	Modes() *[]AuthorizationType
	// Experimental.
	SetModes(val *[]AuthorizationType)
	// the name of this type.
	// Experimental.
	Name() *string
	// The resolvers linked to this data source.
	// Experimental.
	Resolvers() *[]Resolver
	// Experimental.
	SetResolvers(val *[]Resolver)
	// Add a field to this Object Type.
	//
	// Object Types must have both fieldName and field options.
	// Experimental.
	AddField(options *AddFieldOptions)
	// Create a GraphQL Type representing this Intermediate Type.
	// Experimental.
	Attribute(options *BaseTypeOptions) GraphqlType
	// Generate the resolvers linked to this Object Type.
	// Experimental.
	GenerateResolver(api IGraphqlApi, fieldName *string, options *ResolvableFieldOptions) Resolver
	// Generate the string of this object type.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ObjectType
type jsiiProxy_ObjectType struct {
	jsiiProxy_InterfaceType
	jsiiProxy_IIntermediateType
}

func (j *jsiiProxy_ObjectType) Definition() *map[string]IField {
	var returns *map[string]IField
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObjectType) Directives() *[]Directive {
	var returns *[]Directive
	_jsii_.Get(
		j,
		"directives",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObjectType) InterfaceTypes() *[]InterfaceType {
	var returns *[]InterfaceType
	_jsii_.Get(
		j,
		"interfaceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObjectType) Modes() *[]AuthorizationType {
	var returns *[]AuthorizationType
	_jsii_.Get(
		j,
		"modes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObjectType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObjectType) Resolvers() *[]Resolver {
	var returns *[]Resolver
	_jsii_.Get(
		j,
		"resolvers",
		&returns,
	)
	return returns
}


// Experimental.
func NewObjectType(name *string, props *ObjectTypeOptions) ObjectType {
	_init_.Initialize()

	if err := validateNewObjectTypeParameters(name, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObjectType{}

	_jsii_.Create(
		"monocdk.aws_appsync.ObjectType",
		[]interface{}{name, props},
		&j,
	)

	return &j
}

// Experimental.
func NewObjectType_Override(o ObjectType, name *string, props *ObjectTypeOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.ObjectType",
		[]interface{}{name, props},
		o,
	)
}

func (j *jsiiProxy_ObjectType)SetModes(val *[]AuthorizationType) {
	_jsii_.Set(
		j,
		"modes",
		val,
	)
}

func (j *jsiiProxy_ObjectType)SetResolvers(val *[]Resolver) {
	_jsii_.Set(
		j,
		"resolvers",
		val,
	)
}

func (o *jsiiProxy_ObjectType) AddField(options *AddFieldOptions) {
	if err := o.validateAddFieldParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addField",
		[]interface{}{options},
	)
}

func (o *jsiiProxy_ObjectType) Attribute(options *BaseTypeOptions) GraphqlType {
	if err := o.validateAttributeParameters(options); err != nil {
		panic(err)
	}
	var returns GraphqlType

	_jsii_.Invoke(
		o,
		"attribute",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObjectType) GenerateResolver(api IGraphqlApi, fieldName *string, options *ResolvableFieldOptions) Resolver {
	if err := o.validateGenerateResolverParameters(api, fieldName, options); err != nil {
		panic(err)
	}
	var returns Resolver

	_jsii_.Invoke(
		o,
		"generateResolver",
		[]interface{}{api, fieldName, options},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObjectType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

