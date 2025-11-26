//go:build !no_runtime_type_checking

package previewawsvoiceidevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsvoiceid"
)

func (d *jsiiProxy_DomainEvents) validateVoiceIdBatchFraudsterRegistrationActionPatternParameters(options *DomainEvents_VoiceIdBatchFraudsterRegistrationAction_VoiceIdBatchFraudsterRegistrationActionProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DomainEvents) validateVoiceIdBatchSpeakerEnrollmentActionPatternParameters(options *DomainEvents_VoiceIdBatchSpeakerEnrollmentAction_VoiceIdBatchSpeakerEnrollmentActionProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DomainEvents) validateVoiceIdEvaluateSessionActionPatternParameters(options *DomainEvents_VoiceIdEvaluateSessionAction_VoiceIdEvaluateSessionActionProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DomainEvents) validateVoiceIdFraudsterActionPatternParameters(options *DomainEvents_VoiceIdFraudsterAction_VoiceIdFraudsterActionProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DomainEvents) validateVoiceIdSessionSpeakerEnrollmentActionPatternParameters(options *DomainEvents_VoiceIdSessionSpeakerEnrollmentAction_VoiceIdSessionSpeakerEnrollmentActionProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DomainEvents) validateVoiceIdSpeakerActionPatternParameters(options *DomainEvents_VoiceIdSpeakerAction_VoiceIdSpeakerActionProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DomainEvents) validateVoiceIdStartSessionActionPatternParameters(options *DomainEvents_VoiceIdStartSessionAction_VoiceIdStartSessionActionProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DomainEvents) validateVoiceIdUpdateSessionActionPatternParameters(options *DomainEvents_VoiceIdUpdateSessionAction_VoiceIdUpdateSessionActionProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateDomainEvents_FromDomainParameters(domainRef interfacesawsvoiceid.IDomainRef) error {
	if domainRef == nil {
		return fmt.Errorf("parameter domainRef is required, but nil was provided")
	}

	return nil
}

