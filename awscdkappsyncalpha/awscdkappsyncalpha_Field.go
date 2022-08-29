// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkappsyncalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Fields build upon Graphql Types and provide typing and arguments.
//
// Example:
//   field := appsync.NewField(&fieldOptions{
//   	returnType: appsync.graphqlType.string(),
//   	args: map[string]*graphqlType{
//   		"argument": appsync.*graphqlType.string(),
//   	},
//   })
//   type := appsync.NewInterfaceType(jsii.String("Node"), &intermediateTypeOptions{
//   	definition: map[string]iField{
//   		"test": field,
//   	},
//   })
//
// Experimental.
type Field interface {
	GraphqlType
	IField
	// The options for this field.
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
	// Generate the args string of this resolvable field.
	// Experimental.
	ArgsToString() *string
	// Generate the directives for this field.
	// Experimental.
	DirectivesToString(modes *[]AuthorizationType) *string
	// Generate the string for this attribute.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Field
type jsiiProxy_Field struct {
	jsiiProxy_GraphqlType
	jsiiProxy_IField
}

func (j *jsiiProxy_Field) FieldOptions() *ResolvableFieldOptions {
	var returns *ResolvableFieldOptions
	_jsii_.Get(
		j,
		"fieldOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Field) IntermediateType() IIntermediateType {
	var returns IIntermediateType
	_jsii_.Get(
		j,
		"intermediateType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Field) IsList() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Field) IsRequired() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Field) IsRequiredList() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isRequiredList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Field) Type() Type {
	var returns Type
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Experimental.
func NewField(options *FieldOptions) Field {
	_init_.Initialize()

	j := jsiiProxy_Field{}

	_jsii_.Create(
		"@aws-cdk/aws-appsync-alpha.Field",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewField_Override(f Field, options *FieldOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-appsync-alpha.Field",
		[]interface{}{options},
		f,
	)
}

// `AWSDate` scalar type represents a valid extended `ISO 8601 Date` string.
//
// In other words, accepts date strings in the form of `YYYY-MM-DD`. It accepts time zone offsets.
// Experimental.
func Field_AwsDate(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.Field",
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
func Field_AwsDateTime(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.Field",
		"awsDateTime",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSEmail` scalar type represents an email address string (i.e.`username@example.com`).
// Experimental.
func Field_AwsEmail(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.Field",
		"awsEmail",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSIPAddress` scalar type respresents a valid `IPv4` of `IPv6` address string.
// Experimental.
func Field_AwsIpAddress(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.Field",
		"awsIpAddress",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSJson` scalar type represents a JSON string.
// Experimental.
func Field_AwsJson(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.Field",
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
func Field_AwsPhone(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.Field",
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
func Field_AwsTime(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.Field",
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
func Field_AwsTimestamp(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.Field",
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
func Field_AwsUrl(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.Field",
		"awsUrl",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `Boolean` scalar type is a boolean value: true or false.
// Experimental.
func Field_Boolean(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.Field",
		"boolean",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `Float` scalar type is a signed double-precision fractional value.
// Experimental.
func Field_Float(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.Field",
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
func Field_Id(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.Field",
		"id",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `Int` scalar type is a signed non-fractional numerical value.
// Experimental.
func Field_Int(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.Field",
		"int",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// an intermediate type to be added as an attribute (i.e. an interface or an object type).
// Experimental.
func Field_Intermediate(options *GraphqlTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.Field",
		"intermediate",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `String` scalar type is a free-form human-readable text.
// Experimental.
func Field_String(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.Field",
		"string",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Field) ArgsToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"argsToString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Field) DirectivesToString(modes *[]AuthorizationType) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"directivesToString",
		[]interface{}{modes},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Field) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

