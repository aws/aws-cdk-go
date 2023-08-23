package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A CustomWidget shows the result of a AWS lambda function.
//
// Example:
//   var dashboard dashboard
//
//
//   // Import or create a lambda function
//   fn := lambda.Function_FromFunctionArn(dashboard, jsii.String("Function"), jsii.String("arn:aws:lambda:us-east-1:123456789012:function:MyFn"))
//
//   dashboard.AddWidgets(cloudwatch.NewCustomWidget(&CustomWidgetProps{
//   	FunctionArn: fn.FunctionArn,
//   	Title: jsii.String("My lambda baked widget"),
//   }))
//
type CustomWidget interface {
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

// The jsii proxy struct for CustomWidget
type jsiiProxy_CustomWidget struct {
	jsiiProxy_ConcreteWidget
}

func (j *jsiiProxy_CustomWidget) Height() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomWidget) Warnings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"warnings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomWidget) WarningsV2() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"warningsV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomWidget) Width() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"width",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomWidget) X() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"x",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomWidget) Y() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"y",
		&returns,
	)
	return returns
}


func NewCustomWidget(props *CustomWidgetProps) CustomWidget {
	_init_.Initialize()

	if err := validateNewCustomWidgetParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomWidget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.CustomWidget",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewCustomWidget_Override(c CustomWidget, props *CustomWidgetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.CustomWidget",
		[]interface{}{props},
		c,
	)
}

func (j *jsiiProxy_CustomWidget)SetX(val *float64) {
	_jsii_.Set(
		j,
		"x",
		val,
	)
}

func (j *jsiiProxy_CustomWidget)SetY(val *float64) {
	_jsii_.Set(
		j,
		"y",
		val,
	)
}

func (c *jsiiProxy_CustomWidget) CopyMetricWarnings(ms ...IMetric) {
	args := []interface{}{}
	for _, a := range ms {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"copyMetricWarnings",
		args,
	)
}

func (c *jsiiProxy_CustomWidget) Position(x *float64, y *float64) {
	if err := c.validatePositionParameters(x, y); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"position",
		[]interface{}{x, y},
	)
}

func (c *jsiiProxy_CustomWidget) ToJson() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

