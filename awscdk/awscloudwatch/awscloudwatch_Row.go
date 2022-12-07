package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A widget that contains other widgets in a horizontal row.
//
// Widgets will be laid out next to each other.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var widget iWidget
//
//   row := awscdk.Aws_cloudwatch.NewRow(widget)
//
// Experimental.
type Row interface {
	IWidget
	// The amount of vertical grid units the widget will take up.
	// Experimental.
	Height() *float64
	// List of contained widgets.
	// Experimental.
	Widgets() *[]IWidget
	// The amount of horizontal grid units the widget will take up.
	// Experimental.
	Width() *float64
	// Place the widget at a given position.
	// Experimental.
	Position(x *float64, y *float64)
	// Return the widget JSON for use in the dashboard.
	// Experimental.
	ToJson() *[]interface{}
}

// The jsii proxy struct for Row
type jsiiProxy_Row struct {
	jsiiProxy_IWidget
}

func (j *jsiiProxy_Row) Height() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Row) Widgets() *[]IWidget {
	var returns *[]IWidget
	_jsii_.Get(
		j,
		"widgets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Row) Width() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"width",
		&returns,
	)
	return returns
}


// Experimental.
func NewRow(widgets ...IWidget) Row {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range widgets {
		args = append(args, a)
	}

	j := jsiiProxy_Row{}

	_jsii_.Create(
		"monocdk.aws_cloudwatch.Row",
		args,
		&j,
	)

	return &j
}

// Experimental.
func NewRow_Override(r Row, widgets ...IWidget) {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range widgets {
		args = append(args, a)
	}

	_jsii_.Create(
		"monocdk.aws_cloudwatch.Row",
		args,
		r,
	)
}

func (r *jsiiProxy_Row) Position(x *float64, y *float64) {
	if err := r.validatePositionParameters(x, y); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"position",
		[]interface{}{x, y},
	)
}

func (r *jsiiProxy_Row) ToJson() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		r,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

