package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A widget that doesn't display anything but takes up space.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spacer := awscdk.Aws_cloudwatch.NewSpacer(&SpacerProps{
//   	Height: jsii.Number(123),
//   	Width: jsii.Number(123),
//   })
//
type Spacer interface {
	IWidget
	// The amount of vertical grid units the widget will take up.
	Height() *float64
	// The amount of horizontal grid units the widget will take up.
	Width() *float64
	// Place the widget at a given position.
	Position(_x *float64, _y *float64)
	// Return the widget JSON for use in the dashboard.
	ToJson() *[]interface{}
}

// The jsii proxy struct for Spacer
type jsiiProxy_Spacer struct {
	jsiiProxy_IWidget
}

func (j *jsiiProxy_Spacer) Height() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Spacer) Width() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"width",
		&returns,
	)
	return returns
}


func NewSpacer(props *SpacerProps) Spacer {
	_init_.Initialize()

	if err := validateNewSpacerParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Spacer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.Spacer",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewSpacer_Override(s Spacer, props *SpacerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.Spacer",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_Spacer) Position(_x *float64, _y *float64) {
	if err := s.validatePositionParameters(_x, _y); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"position",
		[]interface{}{_x, _y},
	)
}

func (s *jsiiProxy_Spacer) ToJson() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		s,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

