package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A dashboard widget that displays MarkDown.
//
// Example:
//   var dashboard dashboard
//
//
//   dashboard.AddWidgets(cloudwatch.NewTextWidget(&TextWidgetProps{
//   	Markdown: jsii.String("# Key Performance Indicators"),
//   }))
//
type TextWidget interface {
	ConcreteWidget
	// The amount of vertical grid units the widget will take up.
	Height() *float64
	// Any warnings that are produced as a result of putting together this widget.
	Warnings() *[]*string
	// Any warnings that are produced as a result of putting together this widget.
	WarningsV2() *map[string]*string
	// The amount of horizontal grid units the widget will take up.
	Width() *float64
	X() *float64
	SetX(val *float64)
	Y() *float64
	SetY(val *float64)
	// Copy the warnings from the given metric.
	CopyMetricWarnings(ms ...IMetric)
	// Place the widget at a given position.
	Position(x *float64, y *float64)
	// Return the widget JSON for use in the dashboard.
	ToJson() *[]interface{}
}

// The jsii proxy struct for TextWidget
type jsiiProxy_TextWidget struct {
	jsiiProxy_ConcreteWidget
}

func (j *jsiiProxy_TextWidget) Height() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextWidget) Warnings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"warnings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextWidget) WarningsV2() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"warningsV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextWidget) Width() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"width",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextWidget) X() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"x",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextWidget) Y() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"y",
		&returns,
	)
	return returns
}


func NewTextWidget(props *TextWidgetProps) TextWidget {
	_init_.Initialize()

	if err := validateNewTextWidgetParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_TextWidget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.TextWidget",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewTextWidget_Override(t TextWidget, props *TextWidgetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.TextWidget",
		[]interface{}{props},
		t,
	)
}

func (j *jsiiProxy_TextWidget)SetX(val *float64) {
	_jsii_.Set(
		j,
		"x",
		val,
	)
}

func (j *jsiiProxy_TextWidget)SetY(val *float64) {
	_jsii_.Set(
		j,
		"y",
		val,
	)
}

func (t *jsiiProxy_TextWidget) CopyMetricWarnings(ms ...IMetric) {
	args := []interface{}{}
	for _, a := range ms {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		t,
		"copyMetricWarnings",
		args,
	)
}

func (t *jsiiProxy_TextWidget) Position(x *float64, y *float64) {
	if err := t.validatePositionParameters(x, y); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"position",
		[]interface{}{x, y},
	)
}

func (t *jsiiProxy_TextWidget) ToJson() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		t,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

