package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A dashboard widget that displays alarms in a grid view.
//
// Example:
//   var dashboard dashboard
//   var errorAlarm alarm
//
//
//   dashboard.AddWidgets(
//   cloudwatch.NewAlarmStatusWidget(&AlarmStatusWidgetProps{
//   	Alarms: []iAlarm{
//   		errorAlarm,
//   	},
//   }))
//
// Experimental.
type AlarmStatusWidget interface {
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

// The jsii proxy struct for AlarmStatusWidget
type jsiiProxy_AlarmStatusWidget struct {
	jsiiProxy_ConcreteWidget
}

func (j *jsiiProxy_AlarmStatusWidget) Height() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlarmStatusWidget) Warnings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"warnings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlarmStatusWidget) Width() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"width",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlarmStatusWidget) X() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"x",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlarmStatusWidget) Y() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"y",
		&returns,
	)
	return returns
}


// Experimental.
func NewAlarmStatusWidget(props *AlarmStatusWidgetProps) AlarmStatusWidget {
	_init_.Initialize()

	if err := validateNewAlarmStatusWidgetParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlarmStatusWidget{}

	_jsii_.Create(
		"monocdk.aws_cloudwatch.AlarmStatusWidget",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewAlarmStatusWidget_Override(a AlarmStatusWidget, props *AlarmStatusWidgetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cloudwatch.AlarmStatusWidget",
		[]interface{}{props},
		a,
	)
}

func (j *jsiiProxy_AlarmStatusWidget)SetX(val *float64) {
	_jsii_.Set(
		j,
		"x",
		val,
	)
}

func (j *jsiiProxy_AlarmStatusWidget)SetY(val *float64) {
	_jsii_.Set(
		j,
		"y",
		val,
	)
}

func (a *jsiiProxy_AlarmStatusWidget) CopyMetricWarnings(ms ...IMetric) {
	args := []interface{}{}
	for _, a := range ms {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"copyMetricWarnings",
		args,
	)
}

func (a *jsiiProxy_AlarmStatusWidget) Position(x *float64, y *float64) {
	if err := a.validatePositionParameters(x, y); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"position",
		[]interface{}{x, y},
	)
}

func (a *jsiiProxy_AlarmStatusWidget) ToJson() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		a,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

