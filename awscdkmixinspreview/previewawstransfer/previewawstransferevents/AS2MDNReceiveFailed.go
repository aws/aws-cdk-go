package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@AS2MDNReceiveFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aS2MDNReceiveFailed := awscdkmixinspreview.Events.NewAS2MDNReceiveFailed()
//
// Experimental.
type AS2MDNReceiveFailed interface {
}

// The jsii proxy struct for AS2MDNReceiveFailed
type jsiiProxy_AS2MDNReceiveFailed struct {
	_ byte // padding
}

// Experimental.
func NewAS2MDNReceiveFailed() AS2MDNReceiveFailed {
	_init_.Initialize()

	j := jsiiProxy_AS2MDNReceiveFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.AS2MDNReceiveFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAS2MDNReceiveFailed_Override(a AS2MDNReceiveFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.AS2MDNReceiveFailed",
		nil, // no parameters
		a,
	)
}

// EventBridge event pattern for AS2 MDN Receive Failed.
// Experimental.
func AS2MDNReceiveFailed_As2MDNReceiveFailedPattern(options *AS2MDNReceiveFailed_AS2MDNReceiveFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateAS2MDNReceiveFailed_As2MDNReceiveFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.AS2MDNReceiveFailed",
		"as2MDNReceiveFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

