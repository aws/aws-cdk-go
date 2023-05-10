package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A dashboard widget that displays the most recent value for every metric.
//
// Example:
//   var dashboard dashboard
//   var visitorCount metric
//   var purchaseCount metric
//
//
//   dashboard.AddWidgets(cloudwatch.NewSingleValueWidget(&SingleValueWidgetProps{
//   	Metrics: []iMetric{
//   		visitorCount,
//   		purchaseCount,
//   	},
//   }))
//
// Experimental.
type SingleValueWidget interface {
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

// The jsii proxy struct for SingleValueWidget
type jsiiProxy_SingleValueWidget struct {
	jsiiProxy_ConcreteWidget
}

func (j *jsiiProxy_SingleValueWidget) Height() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingleValueWidget) Warnings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"warnings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingleValueWidget) Width() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"width",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingleValueWidget) X() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"x",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingleValueWidget) Y() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"y",
		&returns,
	)
	return returns
}


// Experimental.
func NewSingleValueWidget(props *SingleValueWidgetProps) SingleValueWidget {
	_init_.Initialize()

	if err := validateNewSingleValueWidgetParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SingleValueWidget{}

	_jsii_.Create(
		"monocdk.aws_cloudwatch.SingleValueWidget",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewSingleValueWidget_Override(s SingleValueWidget, props *SingleValueWidgetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cloudwatch.SingleValueWidget",
		[]interface{}{props},
		s,
	)
}

func (j *jsiiProxy_SingleValueWidget)SetX(val *float64) {
	_jsii_.Set(
		j,
		"x",
		val,
	)
}

func (j *jsiiProxy_SingleValueWidget)SetY(val *float64) {
	_jsii_.Set(
		j,
		"y",
		val,
	)
}

func (s *jsiiProxy_SingleValueWidget) CopyMetricWarnings(ms ...IMetric) {
	args := []interface{}{}
	for _, a := range ms {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		s,
		"copyMetricWarnings",
		args,
	)
}

func (s *jsiiProxy_SingleValueWidget) Position(x *float64, y *float64) {
	if err := s.validatePositionParameters(x, y); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"position",
		[]interface{}{x, y},
	)
}

func (s *jsiiProxy_SingleValueWidget) ToJson() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		s,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

