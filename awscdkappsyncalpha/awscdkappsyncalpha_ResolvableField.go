// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkappsyncalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Resolvable Fields build upon Graphql Types and provide fields that can resolve into operations on a data source.
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
type ResolvableField interface {
	Field
	IField
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

// The jsii proxy struct for ResolvableField
type jsiiProxy_ResolvableField struct {
	jsiiProxy_Field
	jsiiProxy_IField
}

func (j *jsiiProxy_ResolvableField) FieldOptions() *ResolvableFieldOptions {
	var returns *ResolvableFieldOptions
	_jsii_.Get(
		j,
		"fieldOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResolvableField) IntermediateType() IIntermediateType {
	var returns IIntermediateType
	_jsii_.Get(
		j,
		"intermediateType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResolvableField) IsList() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResolvableField) IsRequired() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResolvableField) IsRequiredList() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isRequiredList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResolvableField) Type() Type {
	var returns Type
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Experimental.
func NewResolvableField(options *ResolvableFieldOptions) ResolvableField {
	_init_.Initialize()

	if err := validateNewResolvableFieldParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_ResolvableField{}

	_jsii_.Create(
		"@aws-cdk/aws-appsync-alpha.ResolvableField",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewResolvableField_Override(r ResolvableField, options *ResolvableFieldOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-appsync-alpha.ResolvableField",
		[]interface{}{options},
		r,
	)
}

// `AWSDate` scalar type represents a valid extended `ISO 8601 Date` string.
//
// In other words, accepts date strings in the form of `YYYY-MM-DD`. It accepts time zone offsets.
// Experimental.
func ResolvableField_AwsDate(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	if err := validateResolvableField_AwsDateParameters(options); err != nil {
		panic(err)
	}
	var returns GraphqlType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.ResolvableField",
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
func ResolvableField_AwsDateTime(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	if err := validateResolvableField_AwsDateTimeParameters(options); err != nil {
		panic(err)
	}
	var returns GraphqlType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.ResolvableField",
		"awsDateTime",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSEmail` scalar type represents an email address string (i.e.`username@example.com`).
// Experimental.
func ResolvableField_AwsEmail(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	if err := validateResolvableField_AwsEmailParameters(options); err != nil {
		panic(err)
	}
	var returns GraphqlType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.ResolvableField",
		"awsEmail",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSIPAddress` scalar type respresents a valid `IPv4` of `IPv6` address string.
// Experimental.
func ResolvableField_AwsIpAddress(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	if err := validateResolvableField_AwsIpAddressParameters(options); err != nil {
		panic(err)
	}
	var returns GraphqlType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.ResolvableField",
		"awsIpAddress",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSJson` scalar type represents a JSON string.
// Experimental.
func ResolvableField_AwsJson(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	if err := validateResolvableField_AwsJsonParameters(options); err != nil {
		panic(err)
	}
	var returns GraphqlType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.ResolvableField",
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
func ResolvableField_AwsPhone(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	if err := validateResolvableField_AwsPhoneParameters(options); err != nil {
		panic(err)
	}
	var returns GraphqlType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.ResolvableField",
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
func ResolvableField_AwsTime(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	if err := validateResolvableField_AwsTimeParameters(options); err != nil {
		panic(err)
	}
	var returns GraphqlType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.ResolvableField",
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
func ResolvableField_AwsTimestamp(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	if err := validateResolvableField_AwsTimestampParameters(options); err != nil {
		panic(err)
	}
	var returns GraphqlType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.ResolvableField",
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
func ResolvableField_AwsUrl(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	if err := validateResolvableField_AwsUrlParameters(options); err != nil {
		panic(err)
	}
	var returns GraphqlType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.ResolvableField",
		"awsUrl",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `Boolean` scalar type is a boolean value: true or false.
// Experimental.
func ResolvableField_Boolean(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	if err := validateResolvableField_BooleanParameters(options); err != nil {
		panic(err)
	}
	var returns GraphqlType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.ResolvableField",
		"boolean",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `Float` scalar type is a signed double-precision fractional value.
// Experimental.
func ResolvableField_Float(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	if err := validateResolvableField_FloatParameters(options); err != nil {
		panic(err)
	}
	var returns GraphqlType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.ResolvableField",
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
func ResolvableField_Id(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	if err := validateResolvableField_IdParameters(options); err != nil {
		panic(err)
	}
	var returns GraphqlType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.ResolvableField",
		"id",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `Int` scalar type is a signed non-fractional numerical value.
// Experimental.
func ResolvableField_Int(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	if err := validateResolvableField_IntParameters(options); err != nil {
		panic(err)
	}
	var returns GraphqlType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.ResolvableField",
		"int",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// an intermediate type to be added as an attribute (i.e. an interface or an object type).
// Experimental.
func ResolvableField_Intermediate(options *GraphqlTypeOptions) GraphqlType {
	_init_.Initialize()

	if err := validateResolvableField_IntermediateParameters(options); err != nil {
		panic(err)
	}
	var returns GraphqlType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.ResolvableField",
		"intermediate",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `String` scalar type is a free-form human-readable text.
// Experimental.
func ResolvableField_String(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	if err := validateResolvableField_StringParameters(options); err != nil {
		panic(err)
	}
	var returns GraphqlType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.ResolvableField",
		"string",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResolvableField) ArgsToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"argsToString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResolvableField) DirectivesToString(modes *[]AuthorizationType) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"directivesToString",
		[]interface{}{modes},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResolvableField) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

