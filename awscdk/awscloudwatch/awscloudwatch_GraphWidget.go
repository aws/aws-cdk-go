package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A dashboard widget that displays metrics.
//
// Example:
//   var dashboard dashboard
//
//
//   dashboard.addWidgets(cloudwatch.NewGraphWidget(&graphWidgetProps{
//   	// ...
//
//   	legendPosition: cloudwatch.legendPosition_RIGHT,
//   }))
//
// Experimental.
type GraphWidget interface {
	ConcreteWidget
	// The amount of vertical grid units the widget will take up.
	// Experimental.
	Height() *float64
	// Any warnings that are produced as a result of putting together this widget.
	// Experimental.
	Warnings() *[]*string
	// The amount of horizontal grid units the widget will take up.
	// Experimental.
	Width() *float64
	// Experimental.
	X() *float64
	// Experimental.
	SetX(val *float64)
	// Experimental.
	Y() *float64
	// Experimental.
	SetY(val *float64)
	// Add another metric to the left Y axis of the GraphWidget.
	// Experimental.
	AddLeftMetric(metric IMetric)
	// Add another metric to the right Y axis of the GraphWidget.
	// Experimental.
	AddRightMetric(metric IMetric)
	// Copy the warnings from the given metric.
	// Experimental.
	CopyMetricWarnings(ms ...IMetric)
	// Place the widget at a given position.
	// Experimental.
	Position(x *float64, y *float64)
	// Return the widget JSON for use in the dashboard.
	// Experimental.
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


// Experimental.
func NewGraphWidget(props *GraphWidgetProps) GraphWidget {
	_init_.Initialize()

	if err := validateNewGraphWidgetParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_GraphWidget{}

	_jsii_.Create(
		"monocdk.aws_cloudwatch.GraphWidget",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewGraphWidget_Override(g GraphWidget, props *GraphWidgetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cloudwatch.GraphWidget",
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

