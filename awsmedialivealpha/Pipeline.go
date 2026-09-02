package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// MediaLive pipeline (channel pipeline 0 or 1).
//
// `STANDARD` channels run two redundant pipelines (`PIPELINE_0`, `PIPELINE_1`).
// `SINGLE_PIPELINE` channels run only `PIPELINE_0`.
//
// Example:
//   var channel Channel
//   var stack Stack
//
//
//   channel.metricDroppedFrames(medialive.Pipeline_PIPELINE_0()).CreateAlarm(stack, jsii.String("DroppedFrames"), &CreateAlarmOptions{
//   	Threshold: jsii.Number(1),
//   	EvaluationPeriods: jsii.Number(2),
//   })
//
//   channel.metricSvqTime(medialive.Pipeline_PIPELINE_0()).CreateAlarm(stack, jsii.String("SvqTime"), &CreateAlarmOptions{
//   	Threshold: jsii.Number(0),
//   	EvaluationPeriods: jsii.Number(1),
//   })
//
//   // Custom metric by name with sum statistic
//   channel.metric(jsii.String("Output4xxErrors"), medialive.Pipeline_PIPELINE_0(), &MetricOptions{
//   	Statistic: jsii.String("sum"),
//   })
//
// Experimental.
type Pipeline interface {
	// Returns the CloudWatch dimension value for this pipeline (`'0'` or `'1'`).
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Pipeline
type jsiiProxy_Pipeline struct {
	_ byte // padding
}

func Pipeline_PIPELINE_0() Pipeline {
	_init_.Initialize()
	var returns Pipeline
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Pipeline",
		"PIPELINE_0",
		&returns,
	)
	return returns
}

func Pipeline_PIPELINE_1() Pipeline {
	_init_.Initialize()
	var returns Pipeline
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Pipeline",
		"PIPELINE_1",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_Pipeline) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

