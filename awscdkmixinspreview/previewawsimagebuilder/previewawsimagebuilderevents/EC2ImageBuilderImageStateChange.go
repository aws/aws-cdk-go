package previewawsimagebuilderevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.imagebuilder@EC2ImageBuilderImageStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eC2ImageBuilderImageStateChange := awscdkmixinspreview.Events.NewEC2ImageBuilderImageStateChange()
//
// Experimental.
type EC2ImageBuilderImageStateChange interface {
}

// The jsii proxy struct for EC2ImageBuilderImageStateChange
type jsiiProxy_EC2ImageBuilderImageStateChange struct {
	_ byte // padding
}

// Experimental.
func NewEC2ImageBuilderImageStateChange() EC2ImageBuilderImageStateChange {
	_init_.Initialize()

	j := jsiiProxy_EC2ImageBuilderImageStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_imagebuilder.events.EC2ImageBuilderImageStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewEC2ImageBuilderImageStateChange_Override(e EC2ImageBuilderImageStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_imagebuilder.events.EC2ImageBuilderImageStateChange",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for EC2 Image Builder Image State Change.
// Experimental.
func EC2ImageBuilderImageStateChange_EventPattern(options *EC2ImageBuilderImageStateChange_EC2ImageBuilderImageStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateEC2ImageBuilderImageStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_imagebuilder.events.EC2ImageBuilderImageStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

