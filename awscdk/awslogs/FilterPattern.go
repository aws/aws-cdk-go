package awslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A collection of static methods to generate appropriate ILogPatterns.
//
// Example:
//   // Search for all events where the component field is equal to
//   // "HttpServer" and either error is true or the latency is higher
//   // than 1000.
//   pattern := logs.FilterPattern_All(logs.FilterPattern_StringValue(jsii.String("$.component"), jsii.String("="), jsii.String("HttpServer")), logs.FilterPattern_Any(logs.FilterPattern_BooleanValue(jsii.String("$.error"), jsii.Boolean(true)), logs.FilterPattern_NumberValue(jsii.String("$.latency"), jsii.String(">"), jsii.Number(1000))))
//
type FilterPattern interface {
}

// The jsii proxy struct for FilterPattern
type jsiiProxy_FilterPattern struct {
	_ byte // padding
}

func NewFilterPattern() FilterPattern {
	_init_.Initialize()

	j := jsiiProxy_FilterPattern{}

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.FilterPattern",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewFilterPattern_Override(f FilterPattern) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.FilterPattern",
		nil, // no parameters
		f,
	)
}

// A JSON log pattern that matches if all given JSON log patterns match.
func FilterPattern_All(patterns ...JsonPattern) JsonPattern {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range patterns {
		args = append(args, a)
	}

	var returns JsonPattern

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_logs.FilterPattern",
		"all",
		args,
		&returns,
	)

	return returns
}

// A log pattern that matches all events.
func FilterPattern_AllEvents() IFilterPattern {
	_init_.Initialize()

	var returns IFilterPattern

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_logs.FilterPattern",
		"allEvents",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A log pattern that matches if all the strings given appear in the event.
func FilterPattern_AllTerms(terms ...*string) IFilterPattern {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range terms {
		args = append(args, a)
	}

	var returns IFilterPattern

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_logs.FilterPattern",
		"allTerms",
		args,
		&returns,
	)

	return returns
}

// A JSON log pattern that matches if any of the given JSON log patterns match.
func FilterPattern_Any(patterns ...JsonPattern) JsonPattern {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range patterns {
		args = append(args, a)
	}

	var returns JsonPattern

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_logs.FilterPattern",
		"any",
		args,
		&returns,
	)

	return returns
}

// A log pattern that matches if any of the strings given appear in the event.
func FilterPattern_AnyTerm(terms ...*string) IFilterPattern {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range terms {
		args = append(args, a)
	}

	var returns IFilterPattern

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_logs.FilterPattern",
		"anyTerm",
		args,
		&returns,
	)

	return returns
}

// A log pattern that matches if any of the given term groups matches the event.
//
// A term group matches an event if all the terms in it appear in the event string.
func FilterPattern_AnyTermGroup(termGroups ...*[]*string) IFilterPattern {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range termGroups {
		args = append(args, a)
	}

	var returns IFilterPattern

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_logs.FilterPattern",
		"anyTermGroup",
		args,
		&returns,
	)

	return returns
}

// A JSON log pattern that matches if the field exists and equals the boolean value.
func FilterPattern_BooleanValue(jsonField *string, value *bool) JsonPattern {
	_init_.Initialize()

	if err := validateFilterPattern_BooleanValueParameters(jsonField, value); err != nil {
		panic(err)
	}
	var returns JsonPattern

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_logs.FilterPattern",
		"booleanValue",
		[]interface{}{jsonField, value},
		&returns,
	)

	return returns
}

// A JSON log patter that matches if the field exists.
//
// This is a readable convenience wrapper over 'field = *'.
func FilterPattern_Exists(jsonField *string) JsonPattern {
	_init_.Initialize()

	if err := validateFilterPattern_ExistsParameters(jsonField); err != nil {
		panic(err)
	}
	var returns JsonPattern

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_logs.FilterPattern",
		"exists",
		[]interface{}{jsonField},
		&returns,
	)

	return returns
}

// A JSON log pattern that matches if the field exists and has the special value 'null'.
func FilterPattern_IsNull(jsonField *string) JsonPattern {
	_init_.Initialize()

	if err := validateFilterPattern_IsNullParameters(jsonField); err != nil {
		panic(err)
	}
	var returns JsonPattern

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_logs.FilterPattern",
		"isNull",
		[]interface{}{jsonField},
		&returns,
	)

	return returns
}

// Use the given string as log pattern.
//
// See https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html
// for information on writing log patterns.
func FilterPattern_Literal(logPatternString *string) IFilterPattern {
	_init_.Initialize()

	if err := validateFilterPattern_LiteralParameters(logPatternString); err != nil {
		panic(err)
	}
	var returns IFilterPattern

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_logs.FilterPattern",
		"literal",
		[]interface{}{logPatternString},
		&returns,
	)

	return returns
}

// A JSON log pattern that matches if the field does not exist.
func FilterPattern_NotExists(jsonField *string) JsonPattern {
	_init_.Initialize()

	if err := validateFilterPattern_NotExistsParameters(jsonField); err != nil {
		panic(err)
	}
	var returns JsonPattern

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_logs.FilterPattern",
		"notExists",
		[]interface{}{jsonField},
		&returns,
	)

	return returns
}

// A JSON log pattern that compares numerical values.
//
// This pattern only matches if the event is a JSON event, and the indicated field inside
// compares with the value in the indicated way.
//
// Use '$' to indicate the root of the JSON structure. The comparison operator can only
// compare equality or inequality. The '*' wildcard may appear in the value may at the
// start or at the end.
//
// For more information, see:
//
// https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html
func FilterPattern_NumberValue(jsonField *string, comparison *string, value *float64) JsonPattern {
	_init_.Initialize()

	if err := validateFilterPattern_NumberValueParameters(jsonField, comparison, value); err != nil {
		panic(err)
	}
	var returns JsonPattern

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_logs.FilterPattern",
		"numberValue",
		[]interface{}{jsonField, comparison, value},
		&returns,
	)

	return returns
}

// A space delimited log pattern matcher.
//
// The log event is divided into space-delimited columns (optionally
// enclosed by "" or [] to capture spaces into column values), and names
// are given to each column.
//
// '...' may be specified once to match any number of columns.
//
// Afterwards, conditions may be added to individual columns.
func FilterPattern_SpaceDelimited(columns ...*string) SpaceDelimitedTextPattern {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range columns {
		args = append(args, a)
	}

	var returns SpaceDelimitedTextPattern

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_logs.FilterPattern",
		"spaceDelimited",
		args,
		&returns,
	)

	return returns
}

// A JSON log pattern that compares string values.
//
// This pattern only matches if the event is a JSON event, and the indicated field inside
// compares with the string value.
//
// Use '$' to indicate the root of the JSON structure. The comparison operator can only
// compare equality or inequality. The '*' wildcard may appear in the value may at the
// start or at the end.
//
// For more information, see:
//
// https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html
func FilterPattern_StringValue(jsonField *string, comparison *string, value *string) JsonPattern {
	_init_.Initialize()

	if err := validateFilterPattern_StringValueParameters(jsonField, comparison, value); err != nil {
		panic(err)
	}
	var returns JsonPattern

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_logs.FilterPattern",
		"stringValue",
		[]interface{}{jsonField, comparison, value},
		&returns,
	)

	return returns
}

