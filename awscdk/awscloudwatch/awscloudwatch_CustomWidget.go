package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A CustomWidget shows the result of a AWS lambda function.
//
// Example:
//   var dashboard dashboard
//
//
//   // Import or create a lambda function
//   fn := lambda.function.fromFunctionArn(dashboard, jsii.String("Function"), jsii.String("arn:aws:lambda:us-east-1:123456789012:function:MyFn"))
//
//   dashboard.addWidgets(cloudwatch.NewCustomWidget(&customWidgetProps{
//   	functionArn: fn.functionArn,
//   	title: jsii.String("My lambda baked widget"),
//   }))
//
// Experimental.
type CustomWidget interface {
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


// Experimental.
func NewCustomWidget(props *CustomWidgetProps) CustomWidget {
	_init_.Initialize()

	if err := validateNewCustomWidgetParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomWidget{}

	_jsii_.Create(
		"monocdk.aws_cloudwatch.CustomWidget",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewCustomWidget_Override(c CustomWidget, props *CustomWidgetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cloudwatch.CustomWidget",
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

