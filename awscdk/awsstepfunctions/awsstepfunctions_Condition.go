package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A Condition for use in a Choice state branch.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//   var submitLambda function
//   var getStatusLambda function
//
//
//   submitJob := tasks.NewLambdaInvoke(this, jsii.String("Submit Job"), &lambdaInvokeProps{
//   	lambdaFunction: submitLambda,
//   	// Lambda's result is in the attribute `Payload`
//   	outputPath: jsii.String("$.Payload"),
//   })
//
//   waitX := sfn.NewWait(this, jsii.String("Wait X Seconds"), &waitProps{
//   	time: sfn.waitTime.secondsPath(jsii.String("$.waitSeconds")),
//   })
//
//   getStatus := tasks.NewLambdaInvoke(this, jsii.String("Get Job Status"), &lambdaInvokeProps{
//   	lambdaFunction: getStatusLambda,
//   	// Pass just the field named "guid" into the Lambda, put the
//   	// Lambda's result in a field called "status" in the response
//   	inputPath: jsii.String("$.guid"),
//   	outputPath: jsii.String("$.Payload"),
//   })
//
//   jobFailed := sfn.NewFail(this, jsii.String("Job Failed"), &failProps{
//   	cause: jsii.String("AWS Batch Job Failed"),
//   	error: jsii.String("DescribeJob returned FAILED"),
//   })
//
//   finalStatus := tasks.NewLambdaInvoke(this, jsii.String("Get Final Job Status"), &lambdaInvokeProps{
//   	lambdaFunction: getStatusLambda,
//   	// Use "guid" field as input
//   	inputPath: jsii.String("$.guid"),
//   	outputPath: jsii.String("$.Payload"),
//   })
//
//   definition := submitJob.next(waitX).next(getStatus).next(sfn.NewChoice(this, jsii.String("Job Complete?")).when(sfn.condition.stringEquals(jsii.String("$.status"), jsii.String("FAILED")), jobFailed).when(sfn.condition.stringEquals(jsii.String("$.status"), jsii.String("SUCCEEDED")), finalStatus).otherwise(waitX))
//
//   sfn.NewStateMachine(this, jsii.String("StateMachine"), &stateMachineProps{
//   	definition: definition,
//   	timeout: awscdk.Duration.minutes(jsii.Number(5)),
//   })
//
// Experimental.
type Condition interface {
	// Render Amazon States Language JSON for the condition.
	// Experimental.
	RenderCondition() interface{}
}

// The jsii proxy struct for Condition
type jsiiProxy_Condition struct {
	_ byte // padding
}

// Experimental.
func NewCondition_Override(c Condition) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions.Condition",
		nil, // no parameters
		c,
	)
}

// Combine two or more conditions with a logical AND.
// Experimental.
func Condition_And(conditions ...Condition) Condition {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range conditions {
		args = append(args, a)
	}

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"and",
		args,
		&returns,
	)

	return returns
}

// Matches if a boolean field has the given value.
// Experimental.
func Condition_BooleanEquals(variable *string, value *bool) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"booleanEquals",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a boolean field equals to a value at a given mapping path.
// Experimental.
func Condition_BooleanEqualsJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"booleanEqualsJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if variable is boolean.
// Experimental.
func Condition_IsBoolean(variable *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"isBoolean",
		[]interface{}{variable},
		&returns,
	)

	return returns
}

// Matches if variable is not boolean.
// Experimental.
func Condition_IsNotBoolean(variable *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"isNotBoolean",
		[]interface{}{variable},
		&returns,
	)

	return returns
}

// Matches if variable is not null.
// Experimental.
func Condition_IsNotNull(variable *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"isNotNull",
		[]interface{}{variable},
		&returns,
	)

	return returns
}

// Matches if variable is not numeric.
// Experimental.
func Condition_IsNotNumeric(variable *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"isNotNumeric",
		[]interface{}{variable},
		&returns,
	)

	return returns
}

// Matches if variable is not present.
// Experimental.
func Condition_IsNotPresent(variable *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"isNotPresent",
		[]interface{}{variable},
		&returns,
	)

	return returns
}

// Matches if variable is not a string.
// Experimental.
func Condition_IsNotString(variable *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"isNotString",
		[]interface{}{variable},
		&returns,
	)

	return returns
}

// Matches if variable is not a timestamp.
// Experimental.
func Condition_IsNotTimestamp(variable *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"isNotTimestamp",
		[]interface{}{variable},
		&returns,
	)

	return returns
}

// Matches if variable is Null.
// Experimental.
func Condition_IsNull(variable *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"isNull",
		[]interface{}{variable},
		&returns,
	)

	return returns
}

// Matches if variable is numeric.
// Experimental.
func Condition_IsNumeric(variable *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"isNumeric",
		[]interface{}{variable},
		&returns,
	)

	return returns
}

// Matches if variable is present.
// Experimental.
func Condition_IsPresent(variable *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"isPresent",
		[]interface{}{variable},
		&returns,
	)

	return returns
}

// Matches if variable is a string.
// Experimental.
func Condition_IsString(variable *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"isString",
		[]interface{}{variable},
		&returns,
	)

	return returns
}

// Matches if variable is a timestamp.
// Experimental.
func Condition_IsTimestamp(variable *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"isTimestamp",
		[]interface{}{variable},
		&returns,
	)

	return returns
}

// Negate a condition.
// Experimental.
func Condition_Not(condition Condition) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"not",
		[]interface{}{condition},
		&returns,
	)

	return returns
}

// Matches if a numeric field has the given value.
// Experimental.
func Condition_NumberEquals(variable *string, value *float64) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"numberEquals",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a numeric field has the value in a given mapping path.
// Experimental.
func Condition_NumberEqualsJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"numberEqualsJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a numeric field is greater than the given value.
// Experimental.
func Condition_NumberGreaterThan(variable *string, value *float64) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"numberGreaterThan",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a numeric field is greater than or equal to the given value.
// Experimental.
func Condition_NumberGreaterThanEquals(variable *string, value *float64) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"numberGreaterThanEquals",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a numeric field is greater than or equal to the value at a given mapping path.
// Experimental.
func Condition_NumberGreaterThanEqualsJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"numberGreaterThanEqualsJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a numeric field is greater than the value at a given mapping path.
// Experimental.
func Condition_NumberGreaterThanJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"numberGreaterThanJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a numeric field is less than the given value.
// Experimental.
func Condition_NumberLessThan(variable *string, value *float64) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"numberLessThan",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a numeric field is less than or equal to the given value.
// Experimental.
func Condition_NumberLessThanEquals(variable *string, value *float64) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"numberLessThanEquals",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a numeric field is less than or equal to the numeric value at given mapping path.
// Experimental.
func Condition_NumberLessThanEqualsJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"numberLessThanEqualsJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a numeric field is less than the value at the given mapping path.
// Experimental.
func Condition_NumberLessThanJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"numberLessThanJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Combine two or more conditions with a logical OR.
// Experimental.
func Condition_Or(conditions ...Condition) Condition {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range conditions {
		args = append(args, a)
	}

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"or",
		args,
		&returns,
	)

	return returns
}

// Matches if a string field has the given value.
// Experimental.
func Condition_StringEquals(variable *string, value *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"stringEquals",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a string field equals to a value at a given mapping path.
// Experimental.
func Condition_StringEqualsJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"stringEqualsJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a string field sorts after a given value.
// Experimental.
func Condition_StringGreaterThan(variable *string, value *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"stringGreaterThan",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a string field sorts after or equal to a given value.
// Experimental.
func Condition_StringGreaterThanEquals(variable *string, value *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"stringGreaterThanEquals",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a string field sorts after or equal to value at a given mapping path.
// Experimental.
func Condition_StringGreaterThanEqualsJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"stringGreaterThanEqualsJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a string field sorts after a value at a given mapping path.
// Experimental.
func Condition_StringGreaterThanJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"stringGreaterThanJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a string field sorts before a given value.
// Experimental.
func Condition_StringLessThan(variable *string, value *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"stringLessThan",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a string field sorts equal to or before a given value.
// Experimental.
func Condition_StringLessThanEquals(variable *string, value *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"stringLessThanEquals",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a string field sorts equal to or before a given mapping.
// Experimental.
func Condition_StringLessThanEqualsJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"stringLessThanEqualsJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a string field sorts before a given value at a particular mapping.
// Experimental.
func Condition_StringLessThanJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"stringLessThanJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a field matches a string pattern that can contain a wild card (*) e.g: log-*.txt or *LATEST*. No other characters other than "*" have any special meaning - * can be escaped: \\*.
// Experimental.
func Condition_StringMatches(variable *string, value *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"stringMatches",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a timestamp field is the same time as the given timestamp.
// Experimental.
func Condition_TimestampEquals(variable *string, value *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"timestampEquals",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a timestamp field is the same time as the timestamp at a given mapping path.
// Experimental.
func Condition_TimestampEqualsJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"timestampEqualsJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a timestamp field is after the given timestamp.
// Experimental.
func Condition_TimestampGreaterThan(variable *string, value *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"timestampGreaterThan",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a timestamp field is after or equal to the given timestamp.
// Experimental.
func Condition_TimestampGreaterThanEquals(variable *string, value *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"timestampGreaterThanEquals",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a timestamp field is after or equal to the timestamp at a given mapping path.
// Experimental.
func Condition_TimestampGreaterThanEqualsJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"timestampGreaterThanEqualsJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a timestamp field is after the timestamp at a given mapping path.
// Experimental.
func Condition_TimestampGreaterThanJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"timestampGreaterThanJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a timestamp field is before the given timestamp.
// Experimental.
func Condition_TimestampLessThan(variable *string, value *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"timestampLessThan",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a timestamp field is before or equal to the given timestamp.
// Experimental.
func Condition_TimestampLessThanEquals(variable *string, value *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"timestampLessThanEquals",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a timestamp field is before or equal to the timestamp at a given mapping path.
// Experimental.
func Condition_TimestampLessThanEqualsJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
		"timestampLessThanEqualsJsonPath",
		[]interface{}{variable, value},
		&returns,
	)

	return returns
}

// Matches if a timestamp field is before the timestamp at a given mapping path.
// Experimental.
func Condition_TimestampLessThanJsonPath(variable *string, value *string) Condition {
	_init_.Initialize()

	var returns Condition

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Condition",
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

