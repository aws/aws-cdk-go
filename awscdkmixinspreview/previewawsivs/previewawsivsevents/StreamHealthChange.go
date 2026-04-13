package previewawsivsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.ivs@StreamHealthChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   streamHealthChange := awscdkmixinspreview.Events.NewStreamHealthChange()
//
// Experimental.
type StreamHealthChange interface {
}

// The jsii proxy struct for StreamHealthChange
type jsiiProxy_StreamHealthChange struct {
	_ byte // padding
}

// Experimental.
func NewStreamHealthChange() StreamHealthChange {
	_init_.Initialize()

	j := jsiiProxy_StreamHealthChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ivs.events.StreamHealthChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewStreamHealthChange_Override(s StreamHealthChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ivs.events.StreamHealthChange",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for IVS Stream Health Change.
// Experimental.
func StreamHealthChange_EventPattern(options *StreamHealthChange_StreamHealthChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateStreamHealthChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ivs.events.StreamHealthChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

