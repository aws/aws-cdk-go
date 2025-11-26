//go:build !no_runtime_type_checking

package previewawstransferevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawstransfer"
)

func (a *jsiiProxy_AgreementEvents) validateAS2MDNSendCompletedPatternParameters(options *AgreementEvents_AS2MDNSendCompleted_AS2MDNSendCompletedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AgreementEvents) validateAS2MDNSendFailedPatternParameters(options *AgreementEvents_AS2MDNSendFailed_AS2MDNSendFailedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AgreementEvents) validateAS2PayloadReceiveCompletedPatternParameters(options *AgreementEvents_AS2PayloadReceiveCompleted_AS2PayloadReceiveCompletedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AgreementEvents) validateAS2PayloadReceiveFailedPatternParameters(options *AgreementEvents_AS2PayloadReceiveFailed_AS2PayloadReceiveFailedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateAgreementEvents_FromAgreementParameters(agreementRef interfacesawstransfer.IAgreementRef) error {
	if agreementRef == nil {
		return fmt.Errorf("parameter agreementRef is required, but nil was provided")
	}

	return nil
}

