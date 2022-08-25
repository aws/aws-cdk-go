package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A widget that contains other widgets in a vertical column.
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
//   column := awscdk.Aws_cloudwatch.NewColumn(widget)
//
// Experimental.
type Column interface {
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

// The jsii proxy struct for Column
type jsiiProxy_Column struct {
	jsiiProxy_IWidget
}

func (j *jsiiProxy_Column) Height() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Column) Widgets() *[]IWidget {
	var returns *[]IWidget
	_jsii_.Get(
		j,
		"widgets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Column) Width() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"width",
		&returns,
	)
	return returns
}


// Experimental.
func NewColumn(widgets ...IWidget) Column {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range widgets {
		args = append(args, a)
	}

	j := jsiiProxy_Column{}

	_jsii_.Create(
		"monocdk.aws_cloudwatch.Column",
		args,
		&j,
	)

	return &j
}

// Experimental.
func NewColumn_Override(c Column, widgets ...IWidget) {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range widgets {
		args = append(args, a)
	}

	_jsii_.Create(
		"monocdk.aws_cloudwatch.Column",
		args,
		c,
	)
}

func (c *jsiiProxy_Column) Position(x *float64, y *float64) {
	_jsii_.InvokeVoid(
		c,
		"position",
		[]interface{}{x, y},
	)
}

func (c *jsiiProxy_Column) ToJson() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

