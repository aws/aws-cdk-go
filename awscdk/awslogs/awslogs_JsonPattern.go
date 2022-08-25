package awslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Base class for patterns that only match JSON log events.
//
// Example:
//   // Search for all events where the component field is equal to
//   // "HttpServer" and either error is true or the latency is higher
//   // than 1000.
//   pattern := logs.filterPattern.all(logs.filterPattern.stringValue(jsii.String("$.component"), jsii.String("="), jsii.String("HttpServer")), logs.filterPattern.any(logs.filterPattern.booleanValue(jsii.String("$.error"), jsii.Boolean(true)), logs.filterPattern.numberValue(jsii.String("$.latency"), jsii.String(">"), jsii.Number(1000))))
//
type JsonPattern interface {
	IFilterPattern
	JsonPatternString() *string
	LogPatternString() *string
}

// The jsii proxy struct for JsonPattern
type jsiiProxy_JsonPattern struct {
	jsiiProxy_IFilterPattern
}

func (j *jsiiProxy_JsonPattern) JsonPatternString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsonPatternString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JsonPattern) LogPatternString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logPatternString",
		&returns,
	)
	return returns
}


func NewJsonPattern_Override(j JsonPattern, jsonPatternString *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.JsonPattern",
		[]interface{}{jsonPatternString},
		j,
	)
}

