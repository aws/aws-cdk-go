package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@AS2MDNSendFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aS2MDNSendFailed := awscdkmixinspreview.Events.NewAS2MDNSendFailed()
//
// Experimental.
type AS2MDNSendFailed interface {
}

// The jsii proxy struct for AS2MDNSendFailed
type jsiiProxy_AS2MDNSendFailed struct {
	_ byte // padding
}

// Experimental.
func NewAS2MDNSendFailed() AS2MDNSendFailed {
	_init_.Initialize()

	j := jsiiProxy_AS2MDNSendFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.AS2MDNSendFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAS2MDNSendFailed_Override(a AS2MDNSendFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.AS2MDNSendFailed",
		nil, // no parameters
		a,
	)
}

// EventBridge event pattern for AS2 MDN Send Failed.
// Experimental.
func AS2MDNSendFailed_As2MDNSendFailedPattern(options *AS2MDNSendFailed_AS2MDNSendFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateAS2MDNSendFailed_As2MDNSendFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.AS2MDNSendFailed",
		"as2MDNSendFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

