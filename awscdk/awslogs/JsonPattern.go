package awslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Base class for patterns that only match JSON log events.
//
// Example:
//   awscdk.NewMetricFilter(this, jsii.String("MetricFilter"), &MetricFilterProps{
//   	LogGroup: LogGroup,
//   	MetricNamespace: jsii.String("MyApp"),
//   	MetricName: jsii.String("Latency"),
//   	FilterPattern: awscdk.FilterPattern_All(awscdk.FilterPattern_Exists(jsii.String("$.latency")), awscdk.FilterPattern_RegexValue(jsii.String("$.message"), jsii.String("="), jsii.String("bind: address already in use"))),
//   	MetricValue: jsii.String("$.latency"),
//   })
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

