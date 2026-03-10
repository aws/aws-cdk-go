package previewawsivsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.ivs@StreamStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   streamStateChange := awscdkmixinspreview.Events.NewStreamStateChange()
//
// Experimental.
type StreamStateChange interface {
}

// The jsii proxy struct for StreamStateChange
type jsiiProxy_StreamStateChange struct {
	_ byte // padding
}

// Experimental.
func NewStreamStateChange() StreamStateChange {
	_init_.Initialize()

	j := jsiiProxy_StreamStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ivs.events.StreamStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewStreamStateChange_Override(s StreamStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ivs.events.StreamStateChange",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for IVS Stream State Change.
// Experimental.
func StreamStateChange_StreamStateChangePattern(options *StreamStateChange_StreamStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateStreamStateChange_StreamStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ivs.events.StreamStateChange",
		"streamStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

