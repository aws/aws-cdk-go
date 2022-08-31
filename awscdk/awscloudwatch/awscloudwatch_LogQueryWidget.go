package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Display query results from Logs Insights.
//
// Example:
//   var dashboard dashboard
//
//
//   dashboard.addWidgets(cloudwatch.NewLogQueryWidget(&logQueryWidgetProps{
//   	logGroupNames: []*string{
//   		jsii.String("my-log-group"),
//   	},
//   	view: cloudwatch.logQueryVisualizationType_TABLE,
//   	// The lines will be automatically combined using '\n|'.
//   	queryLines: []*string{
//   		jsii.String("fields @message"),
//   		jsii.String("filter @message like /Error/"),
//   	},
//   }))
//
// Experimental.
type LogQueryWidget interface {
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

// The jsii proxy struct for LogQueryWidget
type jsiiProxy_LogQueryWidget struct {
	jsiiProxy_ConcreteWidget
}

func (j *jsiiProxy_LogQueryWidget) Height() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogQueryWidget) Warnings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"warnings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogQueryWidget) Width() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"width",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogQueryWidget) X() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"x",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogQueryWidget) Y() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"y",
		&returns,
	)
	return returns
}


// Experimental.
func NewLogQueryWidget(props *LogQueryWidgetProps) LogQueryWidget {
	_init_.Initialize()

	j := jsiiProxy_LogQueryWidget{}

	_jsii_.Create(
		"monocdk.aws_cloudwatch.LogQueryWidget",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewLogQueryWidget_Override(l LogQueryWidget, props *LogQueryWidgetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cloudwatch.LogQueryWidget",
		[]interface{}{props},
		l,
	)
}

func (j *jsiiProxy_LogQueryWidget) SetX(val *float64) {
	_jsii_.Set(
		j,
		"x",
		val,
	)
}

func (j *jsiiProxy_LogQueryWidget) SetY(val *float64) {
	_jsii_.Set(
		j,
		"y",
		val,
	)
}

func (l *jsiiProxy_LogQueryWidget) CopyMetricWarnings(ms ...IMetric) {
	args := []interface{}{}
	for _, a := range ms {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		l,
		"copyMetricWarnings",
		args,
	)
}

func (l *jsiiProxy_LogQueryWidget) Position(x *float64, y *float64) {
	_jsii_.InvokeVoid(
		l,
		"position",
		[]interface{}{x, y},
	)
}

func (l *jsiiProxy_LogQueryWidget) ToJson() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		l,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

