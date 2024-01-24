package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A Condition for use in a Choice state branch.
//
// Example:
//   map := sfn.NewMap(this, jsii.String("Map State"), &MapProps{
//   	MaxConcurrency: jsii.Number(1),
//   	ItemsPath: sfn.JsonPath_StringAt(jsii.String("$.inputForMap")),
//   	ItemSelector: map[string]interface{}{
//   		"item": sfn.JsonPath_*StringAt(jsii.String("$.Map.Item.Value")),
//   	},
//   	ResultPath: jsii.String("$.mapOutput"),
//   })
//
//   // The Map iterator can contain a IChainable, which can be an individual or multiple steps chained together.
//   // Below example is with a Choice and Pass step
//   choice := sfn.NewChoice(this, jsii.String("Choice"))
//   condition1 := sfn.Condition_StringEquals(jsii.String("$.item.status"), jsii.String("SUCCESS"))
//   step1 := sfn.NewPass(this, jsii.String("Step1"))
//   step2 := sfn.NewPass(this, jsii.String("Step2"))
//   finish := sfn.NewPass(this, jsii.String("Finish"))
//
//   definition := choice.When(condition1, step1).Otherwise(step2).Afterwards().Next(finish)
//
//   map.ItemProcessor(definition)
//
type Condition interface {
	// Render Amazon States Language JSON for the condition.
	RenderCondition() interface{}
}

// The jsii proxy struct for Condition
type jsiiProxy_Condition struct {
	_ byte // padding
}

func NewCondition_Override(c Condition) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		nil, // no parameters
		c,
	)
}

// Combine two or more conditions with a logical AND.
func Condition_And(conditions ...Condition) Condition {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range conditions {
		args = append(args, a)
	}

	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"and",
		args,
		&returns,
	)

	return returns
}

// Matches if a boolean field has the given value.
func Condition_BooleanEquals(variable *string, value *bool) Condition {
	_init_.Initialize()

	if err := validateCondition_BooleanEqualsParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"booleanEquals",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a boolean field equals to a value at a given mapping path.
func Condition_BooleanEqualsJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	if err := validateCondition_BooleanEqualsJsonPathParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"booleanEqualsJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if variable is boolean.
func Condition_IsBoolean(variable *string) Condition {
	_init_.Initialize()

	if err := validateCondition_IsBooleanParameters(variable); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"isBoolean",
		[]interface{}{variable},
		&returns,
	)

	return returns
}

// Matches if variable is not boolean.
func Condition_IsNotBoolean(variable *string) Condition {
	_init_.Initialize()

	if err := validateCondition_IsNotBooleanParameters(variable); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"isNotBoolean",
		[]interface{}{variable},
		&returns,
	)

	return returns
}

// Matches if variable is not null.
func Condition_IsNotNull(variable *string) Condition {
	_init_.Initialize()

	if err := validateCondition_IsNotNullParameters(variable); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"isNotNull",
		[]interface{}{variable},
		&returns,
	)

	return returns
}

// Matches if variable is not numeric.
func Condition_IsNotNumeric(variable *string) Condition {
	_init_.Initialize()

	if err := validateCondition_IsNotNumericParameters(variable); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"isNotNumeric",
		[]interface{}{variable},
		&returns,
	)

	return returns
}

// Matches if variable is not present.
func Condition_IsNotPresent(variable *string) Condition {
	_init_.Initialize()

	if err := validateCondition_IsNotPresentParameters(variable); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"isNotPresent",
		[]interface{}{variable},
		&returns,
	)

	return returns
}

// Matches if variable is not a string.
func Condition_IsNotString(variable *string) Condition {
	_init_.Initialize()

	if err := validateCondition_IsNotStringParameters(variable); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"isNotString",
		[]interface{}{variable},
		&returns,
	)

	return returns
}

// Matches if variable is not a timestamp.
func Condition_IsNotTimestamp(variable *string) Condition {
	_init_.Initialize()

	if err := validateCondition_IsNotTimestampParameters(variable); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"isNotTimestamp",
		[]interface{}{variable},
		&returns,
	)

	return returns
}

// Matches if variable is Null.
func Condition_IsNull(variable *string) Condition {
	_init_.Initialize()

	if err := validateCondition_IsNullParameters(variable); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"isNull",
		[]interface{}{variable},
		&returns,
	)

	return returns
}

// Matches if variable is numeric.
func Condition_IsNumeric(variable *string) Condition {
	_init_.Initialize()

	if err := validateCondition_IsNumericParameters(variable); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"isNumeric",
		[]interface{}{variable},
		&returns,
	)

	return returns
}

// Matches if variable is present.
func Condition_IsPresent(variable *string) Condition {
	_init_.Initialize()

	if err := validateCondition_IsPresentParameters(variable); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"isPresent",
		[]interface{}{variable},
		&returns,
	)

	return returns
}

// Matches if variable is a string.
func Condition_IsString(variable *string) Condition {
	_init_.Initialize()

	if err := validateCondition_IsStringParameters(variable); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"isString",
		[]interface{}{variable},
		&returns,
	)

	return returns
}

// Matches if variable is a timestamp.
func Condition_IsTimestamp(variable *string) Condition {
	_init_.Initialize()

	if err := validateCondition_IsTimestampParameters(variable); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"isTimestamp",
		[]interface{}{variable},
		&returns,
	)

	return returns
}

// Negate a condition.
func Condition_Not(condition Condition) Condition {
	_init_.Initialize()

	if err := validateCondition_NotParameters(condition); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"not",
		[]interface{}{condition},
		&returns,
	)

	return returns
}

// Matches if a numeric field has the given value.
func Condition_NumberEquals(variable *string, value *float64) Condition {
	_init_.Initialize()

	if err := validateCondition_NumberEqualsParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"numberEquals",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a numeric field has the value in a given mapping path.
func Condition_NumberEqualsJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	if err := validateCondition_NumberEqualsJsonPathParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"numberEqualsJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a numeric field is greater than the given value.
func Condition_NumberGreaterThan(variable *string, value *float64) Condition {
	_init_.Initialize()

	if err := validateCondition_NumberGreaterThanParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"numberGreaterThan",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a numeric field is greater than or equal to the given value.
func Condition_NumberGreaterThanEquals(variable *string, value *float64) Condition {
	_init_.Initialize()

	if err := validateCondition_NumberGreaterThanEqualsParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"numberGreaterThanEquals",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a numeric field is greater than or equal to the value at a given mapping path.
func Condition_NumberGreaterThanEqualsJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	if err := validateCondition_NumberGreaterThanEqualsJsonPathParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"numberGreaterThanEqualsJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a numeric field is greater than the value at a given mapping path.
func Condition_NumberGreaterThanJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	if err := validateCondition_NumberGreaterThanJsonPathParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"numberGreaterThanJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a numeric field is less than the given value.
func Condition_NumberLessThan(variable *string, value *float64) Condition {
	_init_.Initialize()

	if err := validateCondition_NumberLessThanParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"numberLessThan",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a numeric field is less than or equal to the given value.
func Condition_NumberLessThanEquals(variable *string, value *float64) Condition {
	_init_.Initialize()

	if err := validateCondition_NumberLessThanEqualsParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"numberLessThanEquals",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a numeric field is less than or equal to the numeric value at given mapping path.
func Condition_NumberLessThanEqualsJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	if err := validateCondition_NumberLessThanEqualsJsonPathParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"numberLessThanEqualsJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a numeric field is less than the value at the given mapping path.
func Condition_NumberLessThanJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	if err := validateCondition_NumberLessThanJsonPathParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"numberLessThanJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Combine two or more conditions with a logical OR.
func Condition_Or(conditions ...Condition) Condition {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range conditions {
		args = append(args, a)
	}

	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"or",
		args,
		&returns,
	)

	return returns
}

// Matches if a string field has the given value.
func Condition_StringEquals(variable *string, value *string) Condition {
	_init_.Initialize()

	if err := validateCondition_StringEqualsParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"stringEquals",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a string field equals to a value at a given mapping path.
func Condition_StringEqualsJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	if err := validateCondition_StringEqualsJsonPathParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"stringEqualsJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a string field sorts after a given value.
func Condition_StringGreaterThan(variable *string, value *string) Condition {
	_init_.Initialize()

	if err := validateCondition_StringGreaterThanParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"stringGreaterThan",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a string field sorts after or equal to a given value.
func Condition_StringGreaterThanEquals(variable *string, value *string) Condition {
	_init_.Initialize()

	if err := validateCondition_StringGreaterThanEqualsParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"stringGreaterThanEquals",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a string field sorts after or equal to value at a given mapping path.
func Condition_StringGreaterThanEqualsJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	if err := validateCondition_StringGreaterThanEqualsJsonPathParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"stringGreaterThanEqualsJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a string field sorts after a value at a given mapping path.
func Condition_StringGreaterThanJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	if err := validateCondition_StringGreaterThanJsonPathParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"stringGreaterThanJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a string field sorts before a given value.
func Condition_StringLessThan(variable *string, value *string) Condition {
	_init_.Initialize()

	if err := validateCondition_StringLessThanParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"stringLessThan",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a string field sorts equal to or before a given value.
func Condition_StringLessThanEquals(variable *string, value *string) Condition {
	_init_.Initialize()

	if err := validateCondition_StringLessThanEqualsParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"stringLessThanEquals",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a string field sorts equal to or before a given mapping.
func Condition_StringLessThanEqualsJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	if err := validateCondition_StringLessThanEqualsJsonPathParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"stringLessThanEqualsJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a string field sorts before a given value at a particular mapping.
func Condition_StringLessThanJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	if err := validateCondition_StringLessThanJsonPathParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"stringLessThanJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a field matches a string pattern that can contain a wild card (*) e.g: log-*.txt or *LATEST*. No other characters other than "*" have any special meaning - * can be escaped: \\*.
func Condition_StringMatches(variable *string, value *string) Condition {
	_init_.Initialize()

	if err := validateCondition_StringMatchesParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"stringMatches",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a timestamp field is the same time as the given timestamp.
func Condition_TimestampEquals(variable *string, value *string) Condition {
	_init_.Initialize()

	if err := validateCondition_TimestampEqualsParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"timestampEquals",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a timestamp field is the same time as the timestamp at a given mapping path.
func Condition_TimestampEqualsJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	if err := validateCondition_TimestampEqualsJsonPathParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"timestampEqualsJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a timestamp field is after the given timestamp.
func Condition_TimestampGreaterThan(variable *string, value *string) Condition {
	_init_.Initialize()

	if err := validateCondition_TimestampGreaterThanParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"timestampGreaterThan",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a timestamp field is after or equal to the given timestamp.
func Condition_TimestampGreaterThanEquals(variable *string, value *string) Condition {
	_init_.Initialize()

	if err := validateCondition_TimestampGreaterThanEqualsParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"timestampGreaterThanEquals",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a timestamp field is after or equal to the timestamp at a given mapping path.
func Condition_TimestampGreaterThanEqualsJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	if err := validateCondition_TimestampGreaterThanEqualsJsonPathParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"timestampGreaterThanEqualsJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a timestamp field is after the timestamp at a given mapping path.
func Condition_TimestampGreaterThanJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	if err := validateCondition_TimestampGreaterThanJsonPathParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"timestampGreaterThanJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a timestamp field is before the given timestamp.
func Condition_TimestampLessThan(variable *string, value *string) Condition {
	_init_.Initialize()

	if err := validateCondition_TimestampLessThanParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"timestampLessThan",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a timestamp field is before or equal to the given timestamp.
func Condition_TimestampLessThanEquals(variable *string, value *string) Condition {
	_init_.Initialize()

	if err := validateCondition_TimestampLessThanEqualsParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"timestampLessThanEquals",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a timestamp field is before or equal to the timestamp at a given mapping path.
func Condition_TimestampLessThanEqualsJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	if err := validateCondition_TimestampLessThanEqualsJsonPathParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"timestampLessThanEqualsJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a timestamp field is before the timestamp at a given mapping path.
func Condition_TimestampLessThanJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	if err := validateCondition_TimestampLessThanJsonPathParameters(variable, value); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Condition",
		"timestampLessThanJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Condition) RenderCondition() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderCondition",
		nil, // no parameters
		&returns,
	)

	return returns
}

