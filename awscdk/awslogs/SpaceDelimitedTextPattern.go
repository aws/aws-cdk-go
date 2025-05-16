package awslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Space delimited text pattern.
//
// Example:
//   // Search for all events where the component is "HttpServer" and the
//   // result code is not equal to 200.
//   pattern := logs.FilterPattern_SpaceDelimited(jsii.String("time"), jsii.String("component"), jsii.String("..."), jsii.String("result_code"), jsii.String("latency")).WhereString(jsii.String("component"), jsii.String("="), jsii.String("HttpServer")).WhereNumber(jsii.String("result_code"), jsii.String("!="), jsii.Number(200))
//
type SpaceDelimitedTextPattern interface {
	IFilterPattern
	LogPatternString() *string
	// Restrict where the pattern applies.
	WhereNumber(columnName *string, comparison *string, value *float64) SpaceDelimitedTextPattern
	// Restrict where the pattern applies.
	WhereString(columnName *string, comparison *string, value *string) SpaceDelimitedTextPattern
}

// The jsii proxy struct for SpaceDelimitedTextPattern
type jsiiProxy_SpaceDelimitedTextPattern struct {
	jsiiProxy_IFilterPattern
}

func (j *jsiiProxy_SpaceDelimitedTextPattern) LogPatternString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logPatternString",
		&returns,
	)
	return returns
}


func NewSpaceDelimitedTextPattern(columns *[]*string, restrictions *map[string]*[]*ColumnRestriction) SpaceDelimitedTextPattern {
	_init_.Initialize()

	if err := validateNewSpaceDelimitedTextPatternParameters(columns, restrictions); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpaceDelimitedTextPattern{}

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.SpaceDelimitedTextPattern",
		[]interface{}{columns, restrictions},
		&j,
	)

	return &j
}

func NewSpaceDelimitedTextPattern_Override(s SpaceDelimitedTextPattern, columns *[]*string, restrictions *map[string]*[]*ColumnRestriction) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.SpaceDelimitedTextPattern",
		[]interface{}{columns, restrictions},
		s,
	)
}

// Construct a new instance of a space delimited text pattern.
//
// Since this class must be public, we can't rely on the user only creating it through
// the `LogPattern.spaceDelimited()` factory function. We must therefore validate the
// argument in the constructor. Since we're returning a copy on every mutation, and we
// don't want to re-validate the same things on every construction, we provide a limited
// set of mutator functions and only validate the new data every time.
func SpaceDelimitedTextPattern_Construct(columns *[]*string) SpaceDelimitedTextPattern {
	_init_.Initialize()

	if err := validateSpaceDelimitedTextPattern_ConstructParameters(columns); err != nil {
		panic(err)
	}
	var returns SpaceDelimitedTextPattern

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_logs.SpaceDelimitedTextPattern",
		"construct",
		[]interface{}{columns},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpaceDelimitedTextPattern) WhereNumber(columnName *string, comparison *string, value *float64) SpaceDelimitedTextPattern {
	if err := s.validateWhereNumberParameters(columnName, comparison, value); err != nil {
		panic(err)
	}
	var returns SpaceDelimitedTextPattern

	_jsii_.Invoke(
		s,
		"whereNumber",
		[]interface{}{columnName, comparison, value},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpaceDelimitedTextPattern) WhereString(columnName *string, comparison *string, value *string) SpaceDelimitedTextPattern {
	if err := s.validateWhereStringParameters(columnName, comparison, value); err != nil {
		panic(err)
	}
	var returns SpaceDelimitedTextPattern

	_jsii_.Invoke(
		s,
		"whereString",
		[]interface{}{columnName, comparison, value},
		&returns,
	)

	return returns
}

