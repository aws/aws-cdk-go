package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents the data for an attribute.
//
// Each attribute value is described as a name-value pair.
// The name is the data type, and the value is the data itself.
//
// Example:
//   var myTable table
//
//   tasks.NewDynamoDeleteItem(this, jsii.String("DeleteItem"), &DynamoDeleteItemProps{
//   	Key: map[string]dynamoAttributeValue{
//   		"MessageId": tasks.*dynamoAttributeValue_fromString(jsii.String("message-007")),
//   	},
//   	Table: myTable,
//   	ResultPath: sfn.JsonPath_DISCARD(),
//   })
//
// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_AttributeValue.html
//
type DynamoAttributeValue interface {
	// Represents the data for the attribute.
	//
	// Data can be
	// i.e. "S": "Hello"
	AttributeValue() interface{}
	// Returns the DynamoDB attribute value.
	ToObject() interface{}
}

// The jsii proxy struct for DynamoAttributeValue
type jsiiProxy_DynamoAttributeValue struct {
	_ byte // padding
}

func (j *jsiiProxy_DynamoAttributeValue) AttributeValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributeValue",
		&returns,
	)
	return returns
}


// Sets an attribute of type Boolean from state input through JSONata expression.
//
// For example:  "BOOL": true.
func DynamoAttributeValue_BooleanFromJsonata(value *string) DynamoAttributeValue {
	_init_.Initialize()

	if err := validateDynamoAttributeValue_BooleanFromJsonataParameters(value); err != nil {
		panic(err)
	}
	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoAttributeValue",
		"booleanFromJsonata",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Sets an attribute of type Boolean from state input through Json path.
//
// For example:  "BOOL": true.
func DynamoAttributeValue_BooleanFromJsonPath(value *string) DynamoAttributeValue {
	_init_.Initialize()

	if err := validateDynamoAttributeValue_BooleanFromJsonPathParameters(value); err != nil {
		panic(err)
	}
	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoAttributeValue",
		"booleanFromJsonPath",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Sets an attribute of type Binary.
//
// For example:  "B": "dGhpcyB0ZXh0IGlzIGJhc2U2NC1lbmNvZGVk".
func DynamoAttributeValue_FromBinary(value *string) DynamoAttributeValue {
	_init_.Initialize()

	if err := validateDynamoAttributeValue_FromBinaryParameters(value); err != nil {
		panic(err)
	}
	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoAttributeValue",
		"fromBinary",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Sets an attribute of type Binary Set.
//
// For example:  "BS": ["U3Vubnk=", "UmFpbnk=", "U25vd3k="].
func DynamoAttributeValue_FromBinarySet(value *[]*string) DynamoAttributeValue {
	_init_.Initialize()

	if err := validateDynamoAttributeValue_FromBinarySetParameters(value); err != nil {
		panic(err)
	}
	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoAttributeValue",
		"fromBinarySet",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Sets an attribute of type Boolean.
//
// For example:  "BOOL": true.
func DynamoAttributeValue_FromBoolean(value *bool) DynamoAttributeValue {
	_init_.Initialize()

	if err := validateDynamoAttributeValue_FromBooleanParameters(value); err != nil {
		panic(err)
	}
	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoAttributeValue",
		"fromBoolean",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Sets an attribute of type List.
//
// For example:  "L": [ {"S": "Cookies"} , {"S": "Coffee"}, {"N", "3.14159"}]
func DynamoAttributeValue_FromList(value *[]DynamoAttributeValue) DynamoAttributeValue {
	_init_.Initialize()

	if err := validateDynamoAttributeValue_FromListParameters(value); err != nil {
		panic(err)
	}
	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoAttributeValue",
		"fromList",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Sets an attribute of type Map.
//
// For example:  "M": {"Name": {"S": "Joe"}, "Age": {"N": "35"}}.
func DynamoAttributeValue_FromMap(value *map[string]DynamoAttributeValue) DynamoAttributeValue {
	_init_.Initialize()

	if err := validateDynamoAttributeValue_FromMapParameters(value); err != nil {
		panic(err)
	}
	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoAttributeValue",
		"fromMap",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Sets an attribute of type Null.
//
// For example:  "NULL": true.
func DynamoAttributeValue_FromNull(value *bool) DynamoAttributeValue {
	_init_.Initialize()

	if err := validateDynamoAttributeValue_FromNullParameters(value); err != nil {
		panic(err)
	}
	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoAttributeValue",
		"fromNull",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Sets a literal number.
//
// For example: 1234
// Numbers are sent across the network to DynamoDB as strings,
// to maximize compatibility across languages and libraries.
// However, DynamoDB treats them as number type attributes for mathematical operations.
func DynamoAttributeValue_FromNumber(value *float64) DynamoAttributeValue {
	_init_.Initialize()

	if err := validateDynamoAttributeValue_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoAttributeValue",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Sets an attribute of type Number Set.
//
// For example:  "NS": ["42.2", "-19", "7.5", "3.14"]
// Numbers are sent across the network to DynamoDB as strings,
// to maximize compatibility across languages and libraries.
// However, DynamoDB treats them as number type attributes for mathematical operations.
func DynamoAttributeValue_FromNumberSet(value *[]*float64) DynamoAttributeValue {
	_init_.Initialize()

	if err := validateDynamoAttributeValue_FromNumberSetParameters(value); err != nil {
		panic(err)
	}
	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoAttributeValue",
		"fromNumberSet",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Sets an attribute of type String.
//
// For example:  "S": "Hello"
// Strings may be literal values or as JSONata expression or as JsonPath. Example values:
//
// - `DynamoAttributeValue.fromString('someValue')`
// - `DynamoAttributeValue.fromString('{% $bar %}')`
// - `DynamoAttributeValue.fromString(JsonPath.stringAt('$.bar'))`
func DynamoAttributeValue_FromString(value *string) DynamoAttributeValue {
	_init_.Initialize()

	if err := validateDynamoAttributeValue_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoAttributeValue",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Sets an attribute of type String Set.
//
// For example:  "SS": ["Giraffe", "Hippo" ,"Zebra"].
func DynamoAttributeValue_FromStringSet(value *[]*string) DynamoAttributeValue {
	_init_.Initialize()

	if err := validateDynamoAttributeValue_FromStringSetParameters(value); err != nil {
		panic(err)
	}
	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoAttributeValue",
		"fromStringSet",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Sets an attribute of type List.
//
// For example:  "L": [ {"S": "Cookies"} , {"S": "Coffee"}, {"S", "Veggies"}].
func DynamoAttributeValue_ListFromJsonata(value *string) DynamoAttributeValue {
	_init_.Initialize()

	if err := validateDynamoAttributeValue_ListFromJsonataParameters(value); err != nil {
		panic(err)
	}
	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoAttributeValue",
		"listFromJsonata",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Sets an attribute of type List.
//
// For example:  "L": [ {"S": "Cookies"} , {"S": "Coffee"}, {"S", "Veggies"}].
func DynamoAttributeValue_ListFromJsonPath(value *string) DynamoAttributeValue {
	_init_.Initialize()

	if err := validateDynamoAttributeValue_ListFromJsonPathParameters(value); err != nil {
		panic(err)
	}
	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoAttributeValue",
		"listFromJsonPath",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Sets an attribute of type Map.
//
// For example:  "M": {"Name": {"S": "Joe"}, "Age": {"N": "35"}}.
func DynamoAttributeValue_MapFromJsonata(value *string) DynamoAttributeValue {
	_init_.Initialize()

	if err := validateDynamoAttributeValue_MapFromJsonataParameters(value); err != nil {
		panic(err)
	}
	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoAttributeValue",
		"mapFromJsonata",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Sets an attribute of type Map.
//
// For example:  "M": {"Name": {"S": "Joe"}, "Age": {"N": "35"}}.
func DynamoAttributeValue_MapFromJsonPath(value *string) DynamoAttributeValue {
	_init_.Initialize()

	if err := validateDynamoAttributeValue_MapFromJsonPathParameters(value); err != nil {
		panic(err)
	}
	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoAttributeValue",
		"mapFromJsonPath",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Sets an attribute of type Number.
//
// For example:  "N": "123.45"
// Numbers are sent across the network to DynamoDB as strings,
// to maximize compatibility across languages and libraries.
// However, DynamoDB treats them as number type attributes for mathematical operations.
//
// Numbers may be expressed as literal strings or as JSONata expression or as JsonPath.
func DynamoAttributeValue_NumberFromString(value *string) DynamoAttributeValue {
	_init_.Initialize()

	if err := validateDynamoAttributeValue_NumberFromStringParameters(value); err != nil {
		panic(err)
	}
	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoAttributeValue",
		"numberFromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Sets an attribute of type Number Set.
//
// For example:  "NS": ["42.2", "-19", "7.5", "3.14"]
// Numbers are sent across the network to DynamoDB as strings,
// to maximize compatibility across languages and libraries.
// However, DynamoDB treats them as number type attributes for mathematical operations.
//
// Numbers may be expressed as literal strings or as JsonPath.
func DynamoAttributeValue_NumberSetFromStrings(value *[]*string) DynamoAttributeValue {
	_init_.Initialize()

	if err := validateDynamoAttributeValue_NumberSetFromStringsParameters(value); err != nil {
		panic(err)
	}
	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoAttributeValue",
		"numberSetFromStrings",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoAttributeValue) ToObject() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toObject",
		nil, // no parameters
		&returns,
	)

	return returns
}

