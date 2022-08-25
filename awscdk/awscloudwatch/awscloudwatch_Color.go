package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A set of standard colours that can be used in annotations in a GraphWidget.
//
// Example:
//   var dashboard dashboard
//   var executionCountMetric metric
//   var errorCountMetric metric
//
//
//   dashboard.addWidgets(cloudwatch.NewGraphWidget(&graphWidgetProps{
//   	title: jsii.String("Executions vs error rate"),
//
//   	left: []iMetric{
//   		executionCountMetric,
//   	},
//
//   	right: []*iMetric{
//   		errorCountMetric.with(&metricOptions{
//   			statistic: jsii.String("average"),
//   			label: jsii.String("Error rate"),
//   			color: cloudwatch.color_GREEN(),
//   		}),
//   	},
//   }))
//
type Color interface {
}

// The jsii proxy struct for Color
type jsiiProxy_Color struct {
	_ byte // padding
}

func Color_BLUE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudwatch.Color",
		"BLUE",
		&returns,
	)
	return returns
}

func Color_BROWN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudwatch.Color",
		"BROWN",
		&returns,
	)
	return returns
}

func Color_GREEN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudwatch.Color",
		"GREEN",
		&returns,
	)
	return returns
}

func Color_GREY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudwatch.Color",
		"GREY",
		&returns,
	)
	return returns
}

func Color_ORANGE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudwatch.Color",
		"ORANGE",
		&returns,
	)
	return returns
}

func Color_PINK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudwatch.Color",
		"PINK",
		&returns,
	)
	return returns
}

func Color_PURPLE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudwatch.Color",
		"PURPLE",
		&returns,
	)
	return returns
}

func Color_RED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudwatch.Color",
		"RED",
		&returns,
	)
	return returns
}

