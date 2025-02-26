package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A dashboard widget that displays metrics.
//
// Example:
//   var dashboard dashboard
//
//
//   dashboard.AddWidgets(cloudwatch.NewTableWidget(&TableWidgetProps{
//   	// ...
//
//   	Layout: cloudwatch.TableLayout_VERTICAL,
//   }))
//
type TableWidget interface {
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
	// Add another metric.
	AddMetric(metric IMetric)
	// Copy the warnings from the given metric.
	CopyMetricWarnings(ms ...IMetric)
	// Place the widget at a given position.
	Position(x *float64, y *float64)
	// Return the widget JSON for use in the dashboard.
	ToJson() *[]interface{}
}

// The jsii proxy struct for TableWidget
type jsiiProxy_TableWidget struct {
	jsiiProxy_ConcreteWidget
}

func (j *jsiiProxy_TableWidget) Height() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TableWidget) Warnings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"warnings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TableWidget) WarningsV2() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"warningsV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TableWidget) Width() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"width",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TableWidget) X() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"x",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TableWidget) Y() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"y",
		&returns,
	)
	return returns
}


func NewTableWidget(props *TableWidgetProps) TableWidget {
	_init_.Initialize()

	if err := validateNewTableWidgetParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_TableWidget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.TableWidget",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewTableWidget_Override(t TableWidget, props *TableWidgetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.TableWidget",
		[]interface{}{props},
		t,
	)
}

func (j *jsiiProxy_TableWidget)SetX(val *float64) {
	_jsii_.Set(
		j,
		"x",
		val,
	)
}

func (j *jsiiProxy_TableWidget)SetY(val *float64) {
	_jsii_.Set(
		j,
		"y",
		val,
	)
}

func (t *jsiiProxy_TableWidget) AddMetric(metric IMetric) {
	if err := t.validateAddMetricParameters(metric); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMetric",
		[]interface{}{metric},
	)
}

func (t *jsiiProxy_TableWidget) CopyMetricWarnings(ms ...IMetric) {
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

func (t *jsiiProxy_TableWidget) Position(x *float64, y *float64) {
	if err := t.validatePositionParameters(x, y); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"position",
		[]interface{}{x, y},
	)
}

func (t *jsiiProxy_TableWidget) ToJson() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		t,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

