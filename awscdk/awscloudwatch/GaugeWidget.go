package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A dashboard gauge widget that displays metrics.
//
// Example:
//   var dashboard dashboard
//   var errorAlarm alarm
//   var gaugeMetric metric
//
//
//   dashboard.AddWidgets(cloudwatch.NewGaugeWidget(&GaugeWidgetProps{
//   	Metrics: []iMetric{
//   		gaugeMetric,
//   	},
//   	LeftYAxis: &YAxisProps{
//   		Min: jsii.Number(0),
//   		Max: jsii.Number(1000),
//   	},
//   }))
//
type GaugeWidget interface {
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
	// Add another metric to the left Y axis of the GaugeWidget.
	AddMetric(metric IMetric)
	// Copy the warnings from the given metric.
	CopyMetricWarnings(ms ...IMetric)
	// Place the widget at a given position.
	Position(x *float64, y *float64)
	// Return the widget JSON for use in the dashboard.
	ToJson() *[]interface{}
}

// The jsii proxy struct for GaugeWidget
type jsiiProxy_GaugeWidget struct {
	jsiiProxy_ConcreteWidget
}

func (j *jsiiProxy_GaugeWidget) Height() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaugeWidget) Warnings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"warnings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaugeWidget) WarningsV2() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"warningsV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaugeWidget) Width() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"width",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaugeWidget) X() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"x",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaugeWidget) Y() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"y",
		&returns,
	)
	return returns
}


func NewGaugeWidget(props *GaugeWidgetProps) GaugeWidget {
	_init_.Initialize()

	if err := validateNewGaugeWidgetParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_GaugeWidget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.GaugeWidget",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewGaugeWidget_Override(g GaugeWidget, props *GaugeWidgetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.GaugeWidget",
		[]interface{}{props},
		g,
	)
}

func (j *jsiiProxy_GaugeWidget)SetX(val *float64) {
	_jsii_.Set(
		j,
		"x",
		val,
	)
}

func (j *jsiiProxy_GaugeWidget)SetY(val *float64) {
	_jsii_.Set(
		j,
		"y",
		val,
	)
}

func (g *jsiiProxy_GaugeWidget) AddMetric(metric IMetric) {
	if err := g.validateAddMetricParameters(metric); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMetric",
		[]interface{}{metric},
	)
}

func (g *jsiiProxy_GaugeWidget) CopyMetricWarnings(ms ...IMetric) {
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

func (g *jsiiProxy_GaugeWidget) Position(x *float64, y *float64) {
	if err := g.validatePositionParameters(x, y); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"position",
		[]interface{}{x, y},
	)
}

func (g *jsiiProxy_GaugeWidget) ToJson() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		g,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

