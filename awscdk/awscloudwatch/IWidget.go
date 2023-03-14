package awscloudwatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A single dashboard widget.
type IWidget interface {
	// Place the widget at a given position.
	Position(x *float64, y *float64)
	// Return the widget JSON for use in the dashboard.
	ToJson() *[]interface{}
	// The amount of vertical grid units the widget will take up.
	Height() *float64
	// Any warnings that are produced as a result of putting together this widget.
	Warnings() *[]*string
	// The amount of horizontal grid units the widget will take up.
	Width() *float64
}

// The jsii proxy for IWidget
type jsiiProxy_IWidget struct {
	_ byte // padding
}

func (i *jsiiProxy_IWidget) Position(x *float64, y *float64) {
	if err := i.validatePositionParameters(x, y); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"position",
		[]interface{}{x, y},
	)
}

func (i *jsiiProxy_IWidget) ToJson() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		i,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IWidget) Height() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWidget) Warnings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"warnings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWidget) Width() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"width",
		&returns,
	)
	return returns
}

