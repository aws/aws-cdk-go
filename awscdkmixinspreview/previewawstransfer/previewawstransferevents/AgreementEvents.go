package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawstransfer"
)

// EventBridge event patterns for Agreement.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var agreementRef IAgreementRef
//
//   agreementEvents := awscdkmixinspreview.Events.AgreementEvents_FromAgreement(agreementRef)
//
// Experimental.
type AgreementEvents interface {
	// EventBridge event pattern for Agreement AS2 MDN Send Completed.
	// Experimental.
	As2MDNSendCompletedPattern(options *AS2MDNSendCompleted_AS2MDNSendCompletedProps) *awsevents.EventPattern
	// EventBridge event pattern for Agreement AS2 MDN Send Failed.
	// Experimental.
	As2MDNSendFailedPattern(options *AS2MDNSendFailed_AS2MDNSendFailedProps) *awsevents.EventPattern
	// EventBridge event pattern for Agreement AS2 Payload Receive Completed.
	// Experimental.
	As2PayloadReceiveCompletedPattern(options *AS2PayloadReceiveCompleted_AS2PayloadReceiveCompletedProps) *awsevents.EventPattern
	// EventBridge event pattern for Agreement AS2 Payload Receive Failed.
	// Experimental.
	As2PayloadReceiveFailedPattern(options *AS2PayloadReceiveFailed_AS2PayloadReceiveFailedProps) *awsevents.EventPattern
}

// The jsii proxy struct for AgreementEvents
type jsiiProxy_AgreementEvents struct {
	_ byte // padding
}

// Create AgreementEvents from a Agreement reference.
// Experimental.
func AgreementEvents_FromAgreement(agreementRef interfacesawstransfer.IAgreementRef) AgreementEvents {
	_init_.Initialize()

	if err := validateAgreementEvents_FromAgreementParameters(agreementRef); err != nil {
		panic(err)
	}
	var returns AgreementEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.AgreementEvents",
		"fromAgreement",
		[]interface{}{agreementRef},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgreementEvents) As2MDNSendCompletedPattern(options *AS2MDNSendCompleted_AS2MDNSendCompletedProps) *awsevents.EventPattern {
	if err := a.validateAs2MDNSendCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		a,
		"as2MDNSendCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgreementEvents) As2MDNSendFailedPattern(options *AS2MDNSendFailed_AS2MDNSendFailedProps) *awsevents.EventPattern {
	if err := a.validateAs2MDNSendFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		a,
		"as2MDNSendFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgreementEvents) As2PayloadReceiveCompletedPattern(options *AS2PayloadReceiveCompleted_AS2PayloadReceiveCompletedProps) *awsevents.EventPattern {
	if err := a.validateAs2PayloadReceiveCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		a,
		"as2PayloadReceiveCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgreementEvents) As2PayloadReceiveFailedPattern(options *AS2PayloadReceiveFailed_AS2PayloadReceiveFailedProps) *awsevents.EventPattern {
	if err := a.validateAs2PayloadReceiveFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		a,
		"as2PayloadReceiveFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

