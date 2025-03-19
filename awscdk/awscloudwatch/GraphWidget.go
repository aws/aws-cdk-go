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
//   dashboard.AddWidgets(cloudwatch.NewGraphWidget(&GraphWidgetProps{
//   	// ...
//
//   	LegendPosition: cloudwatch.LegendPosition_RIGHT,
//   }))
//
type GraphWidget interface {
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
	// Add another metric to the left Y axis of the GraphWidget.
	AddLeftMetric(metric IMetric)
	// Add another metric to the right Y axis of the GraphWidget.
	AddRightMetric(metric IMetric)
	// Copy the warnings from the given metric.
	CopyMetricWarnings(ms ...IMetric)
	// Place the widget at a given position.
	Position(x *float64, y *float64)
	// Return the widget JSON for use in the dashboard.
	ToJson() *[]interface{}
}

// The jsii proxy struct for GraphWidget
type jsiiProxy_GraphWidget struct {
	jsiiProxy_ConcreteWidget
}

func (j *jsiiProxy_GraphWidget) Height() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphWidget) Warnings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"warnings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphWidget) WarningsV2() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"warningsV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphWidget) Width() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"width",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphWidget) X() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"x",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphWidget) Y() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"y",
		&returns,
	)
	return returns
}


func NewGraphWidget(props *GraphWidgetProps) GraphWidget {
	_init_.Initialize()

	if err := validateNewGraphWidgetParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_GraphWidget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.GraphWidget",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewGraphWidget_Override(g GraphWidget, props *GraphWidgetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.GraphWidget",
		[]interface{}{props},
		g,
	)
}

func (j *jsiiProxy_GraphWidget)SetX(val *float64) {
	_jsii_.Set(
		j,
		"x",
		val,
	)
}

func (j *jsiiProxy_GraphWidget)SetY(val *float64) {
	_jsii_.Set(
		j,
		"y",
		val,
	)
}

func (g *jsiiProxy_GraphWidget) AddLeftMetric(metric IMetric) {
	if err := g.validateAddLeftMetricParameters(metric); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addLeftMetric",
		[]interface{}{metric},
	)
}

func (g *jsiiProxy_GraphWidget) AddRightMetric(metric IMetric) {
	if err := g.validateAddRightMetricParameters(metric); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addRightMetric",
		[]interface{}{metric},
	)
}

func (g *jsiiProxy_GraphWidget) CopyMetricWarnings(ms ...IMetric) {
	args := []interface{}{}
	for _, a := range ms {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		g,
		"copyMetricWarnings",
		args,
	)
}

func (g *jsiiProxy_GraphWidget) Position(x *float64, y *float64) {
	if err := g.validatePositionParameters(x, y); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"position",
		[]interface{}{x, y},
	)
}

func (g *jsiiProxy_GraphWidget) ToJson() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		g,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

