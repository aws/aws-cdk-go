package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Display query results from Logs Insights.
//
// Example:
//   var dashboard dashboard
//
//
//   dashboard.AddWidgets(cloudwatch.NewLogQueryWidget(&LogQueryWidgetProps{
//   	LogGroupNames: []*string{
//   		jsii.String("my-log-group"),
//   	},
//   	View: cloudwatch.LogQueryVisualizationType_TABLE,
//   	// The lines will be automatically combined using '\n|'.
//   	QueryLines: []*string{
//   		jsii.String("fields @message"),
//   		jsii.String("filter @message like /Error/"),
//   	},
//   }))
//
type LogQueryWidget interface {
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

func (j *jsiiProxy_LogQueryWidget) WarningsV2() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"warningsV2",
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


func NewLogQueryWidget(props *LogQueryWidgetProps) LogQueryWidget {
	_init_.Initialize()

	if err := validateNewLogQueryWidgetParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogQueryWidget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.LogQueryWidget",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewLogQueryWidget_Override(l LogQueryWidget, props *LogQueryWidgetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.LogQueryWidget",
		[]interface{}{props},
		l,
	)
}

func (j *jsiiProxy_LogQueryWidget)SetX(val *float64) {
	_jsii_.Set(
		j,
		"x",
		val,
	)
}

func (j *jsiiProxy_LogQueryWidget)SetY(val *float64) {
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
	if err := l.validatePositionParameters(x, y); err != nil {
		panic(err)
	}
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

