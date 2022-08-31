package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The GraphQL Types in AppSync's GraphQL.
//
// GraphQL Types are the
// building blocks for object types, queries, mutations, etc. They are
// types like String, Int, Id or even Object Types you create.
//
// i.e. `String`, `String!`, `[String]`, `[String!]`, `[String]!`
//
// GraphQL Types are used to define the entirety of schema.
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
type GraphqlType interface {
	IField
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
	// Generate the arguments for this field.
	// Experimental.
	ArgsToString() *string
	// Generate the directives for this field.
	// Experimental.
	DirectivesToString(_modes *[]AuthorizationType) *string
	// Generate the string for this attribute.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GraphqlType
type jsiiProxy_GraphqlType struct {
	jsiiProxy_IField
}

func (j *jsiiProxy_GraphqlType) IntermediateType() IIntermediateType {
	var returns IIntermediateType
	_jsii_.Get(
		j,
		"intermediateType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlType) IsList() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlType) IsRequired() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlType) IsRequiredList() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isRequiredList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlType) Type() Type {
	var returns Type
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Experimental.
func NewGraphqlType(type_ Type, options *GraphqlTypeOptions) GraphqlType {
	_init_.Initialize()

	j := jsiiProxy_GraphqlType{}

	_jsii_.Create(
		"monocdk.aws_appsync.GraphqlType",
		[]interface{}{type_, options},
		&j,
	)

	return &j
}

// Experimental.
func NewGraphqlType_Override(g GraphqlType, type_ Type, options *GraphqlTypeOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.GraphqlType",
		[]interface{}{type_, options},
		g,
	)
}

// `AWSDate` scalar type represents a valid extended `ISO 8601 Date` string.
//
// In other words, accepts date strings in the form of `YYYY-MM-DD`. It accepts time zone offsets.
// Experimental.
func GraphqlType_AwsDate(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlType",
		"awsDate",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSDateTime` scalar type represents a valid extended `ISO 8601 DateTime` string.
//
// In other words, accepts date strings in the form of `YYYY-MM-DDThh:mm:ss.sssZ`. It accepts time zone offsets.
// Experimental.
func GraphqlType_AwsDateTime(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlType",
		"awsDateTime",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSEmail` scalar type represents an email address string (i.e.`username@example.com`).
// Experimental.
func GraphqlType_AwsEmail(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlType",
		"awsEmail",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSIPAddress` scalar type respresents a valid `IPv4` of `IPv6` address string.
// Experimental.
func GraphqlType_AwsIpAddress(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlType",
		"awsIpAddress",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSJson` scalar type represents a JSON string.
// Experimental.
func GraphqlType_AwsJson(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlType",
		"awsJson",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSPhone` scalar type represents a valid phone number. Phone numbers maybe be whitespace delimited or hyphenated.
//
// The number can specify a country code at the beginning, but is not required for US phone numbers.
// Experimental.
func GraphqlType_AwsPhone(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlType",
		"awsPhone",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSTime` scalar type represents a valid extended `ISO 8601 Time` string.
//
// In other words, accepts date strings in the form of `hh:mm:ss.sss`. It accepts time zone offsets.
// Experimental.
func GraphqlType_AwsTime(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlType",
		"awsTime",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSTimestamp` scalar type represents the number of seconds since `1970-01-01T00:00Z`.
//
// Timestamps are serialized and deserialized as numbers.
// Experimental.
func GraphqlType_AwsTimestamp(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlType",
		"awsTimestamp",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSURL` scalar type represetns a valid URL string.
//
// URLs wihtout schemes or contain double slashes are considered invalid.
// Experimental.
func GraphqlType_AwsUrl(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlType",
		"awsUrl",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `Boolean` scalar type is a boolean value: true or false.
// Experimental.
func GraphqlType_Boolean(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlType",
		"boolean",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `Float` scalar type is a signed double-precision fractional value.
// Experimental.
func GraphqlType_Float(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlType",
		"float",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `ID` scalar type is a unique identifier. `ID` type is serialized similar to `String`.
//
// Often used as a key for a cache and not intended to be human-readable.
// Experimental.
func GraphqlType_Id(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlType",
		"id",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `Int` scalar type is a signed non-fractional numerical value.
// Experimental.
func GraphqlType_Int(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlType",
		"int",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// an intermediate type to be added as an attribute (i.e. an interface or an object type).
// Experimental.
func GraphqlType_Intermediate(options *GraphqlTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlType",
		"intermediate",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `String` scalar type is a free-form human-readable text.
// Experimental.
func GraphqlType_String(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlType",
		"string",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GraphqlType) ArgsToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"argsToString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GraphqlType) DirectivesToString(_modes *[]AuthorizationType) *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"directivesToString",
		[]interface{}{_modes},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GraphqlType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

