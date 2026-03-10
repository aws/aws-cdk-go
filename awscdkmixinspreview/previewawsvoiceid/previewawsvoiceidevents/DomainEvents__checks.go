//go:build !no_runtime_type_checking

package previewawsvoiceidevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsvoiceid"
)

func (d *jsiiProxy_DomainEvents) validateVoiceIdBatchFraudsterRegistrationActionPatternParameters(options *VoiceIdBatchFraudsterRegistrationAction_VoiceIdBatchFraudsterRegistrationActionProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DomainEvents) validateVoiceIdBatchSpeakerEnrollmentActionPatternParameters(options *VoiceIdBatchSpeakerEnrollmentAction_VoiceIdBatchSpeakerEnrollmentActionProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DomainEvents) validateVoiceIdEvaluateSessionActionPatternParameters(options *VoiceIdEvaluateSessionAction_VoiceIdEvaluateSessionActionProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DomainEvents) validateVoiceIdFraudsterActionPatternParameters(options *VoiceIdFraudsterAction_VoiceIdFraudsterActionProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DomainEvents) validateVoiceIdSessionSpeakerEnrollmentActionPatternParameters(options *VoiceIdSessionSpeakerEnrollmentAction_VoiceIdSessionSpeakerEnrollmentActionProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DomainEvents) validateVoiceIdSpeakerActionPatternParameters(options *VoiceIdSpeakerAction_VoiceIdSpeakerActionProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DomainEvents) validateVoiceIdStartSessionActionPatternParameters(options *VoiceIdStartSessionAction_VoiceIdStartSessionActionProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DomainEvents) validateVoiceIdUpdateSessionActionPatternParameters(options *VoiceIdUpdateSessionAction_VoiceIdUpdateSessionActionProps) error {
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

