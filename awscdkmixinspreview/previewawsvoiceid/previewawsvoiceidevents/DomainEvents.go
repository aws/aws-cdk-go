package previewawsvoiceidevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsvoiceid"
)

// EventBridge event patterns for Domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var domainRef IDomainRef
//
//   domainEvents := awscdkmixinspreview.Events.DomainEvents_FromDomain(domainRef)
//
// Experimental.
type DomainEvents interface {
	// EventBridge event pattern for Domain VoiceId Batch Fraudster Registration Action.
	// Experimental.
	VoiceIdBatchFraudsterRegistrationActionPattern(options *DomainEvents_VoiceIdBatchFraudsterRegistrationAction_VoiceIdBatchFraudsterRegistrationActionProps) *awsevents.EventPattern
	// EventBridge event pattern for Domain VoiceId Batch Speaker Enrollment Action.
	// Experimental.
	VoiceIdBatchSpeakerEnrollmentActionPattern(options *DomainEvents_VoiceIdBatchSpeakerEnrollmentAction_VoiceIdBatchSpeakerEnrollmentActionProps) *awsevents.EventPattern
	// EventBridge event pattern for Domain VoiceId Evaluate Session Action.
	// Experimental.
	VoiceIdEvaluateSessionActionPattern(options *DomainEvents_VoiceIdEvaluateSessionAction_VoiceIdEvaluateSessionActionProps) *awsevents.EventPattern
	// EventBridge event pattern for Domain VoiceId Fraudster Action.
	// Experimental.
	VoiceIdFraudsterActionPattern(options *DomainEvents_VoiceIdFraudsterAction_VoiceIdFraudsterActionProps) *awsevents.EventPattern
	// EventBridge event pattern for Domain VoiceId Session Speaker Enrollment Action.
	// Experimental.
	VoiceIdSessionSpeakerEnrollmentActionPattern(options *DomainEvents_VoiceIdSessionSpeakerEnrollmentAction_VoiceIdSessionSpeakerEnrollmentActionProps) *awsevents.EventPattern
	// EventBridge event pattern for Domain VoiceId Speaker Action.
	// Experimental.
	VoiceIdSpeakerActionPattern(options *DomainEvents_VoiceIdSpeakerAction_VoiceIdSpeakerActionProps) *awsevents.EventPattern
	// EventBridge event pattern for Domain VoiceId Start Session Action.
	// Experimental.
	VoiceIdStartSessionActionPattern(options *DomainEvents_VoiceIdStartSessionAction_VoiceIdStartSessionActionProps) *awsevents.EventPattern
	// EventBridge event pattern for Domain VoiceId Update Session Action.
	// Experimental.
	VoiceIdUpdateSessionActionPattern(options *DomainEvents_VoiceIdUpdateSessionAction_VoiceIdUpdateSessionActionProps) *awsevents.EventPattern
}

// The jsii proxy struct for DomainEvents
type jsiiProxy_DomainEvents struct {
	_ byte // padding
}

// Create DomainEvents from a Domain reference.
// Experimental.
func DomainEvents_FromDomain(domainRef interfacesawsvoiceid.IDomainRef) DomainEvents {
	_init_.Initialize()

	if err := validateDomainEvents_FromDomainParameters(domainRef); err != nil {
		panic(err)
	}
	var returns DomainEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents",
		"fromDomain",
		[]interface{}{domainRef},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DomainEvents) VoiceIdBatchFraudsterRegistrationActionPattern(options *DomainEvents_VoiceIdBatchFraudsterRegistrationAction_VoiceIdBatchFraudsterRegistrationActionProps) *awsevents.EventPattern {
	if err := d.validateVoiceIdBatchFraudsterRegistrationActionPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"voiceIdBatchFraudsterRegistrationActionPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DomainEvents) VoiceIdBatchSpeakerEnrollmentActionPattern(options *DomainEvents_VoiceIdBatchSpeakerEnrollmentAction_VoiceIdBatchSpeakerEnrollmentActionProps) *awsevents.EventPattern {
	if err := d.validateVoiceIdBatchSpeakerEnrollmentActionPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"voiceIdBatchSpeakerEnrollmentActionPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DomainEvents) VoiceIdEvaluateSessionActionPattern(options *DomainEvents_VoiceIdEvaluateSessionAction_VoiceIdEvaluateSessionActionProps) *awsevents.EventPattern {
	if err := d.validateVoiceIdEvaluateSessionActionPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"voiceIdEvaluateSessionActionPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DomainEvents) VoiceIdFraudsterActionPattern(options *DomainEvents_VoiceIdFraudsterAction_VoiceIdFraudsterActionProps) *awsevents.EventPattern {
	if err := d.validateVoiceIdFraudsterActionPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"voiceIdFraudsterActionPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DomainEvents) VoiceIdSessionSpeakerEnrollmentActionPattern(options *DomainEvents_VoiceIdSessionSpeakerEnrollmentAction_VoiceIdSessionSpeakerEnrollmentActionProps) *awsevents.EventPattern {
	if err := d.validateVoiceIdSessionSpeakerEnrollmentActionPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"voiceIdSessionSpeakerEnrollmentActionPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DomainEvents) VoiceIdSpeakerActionPattern(options *DomainEvents_VoiceIdSpeakerAction_VoiceIdSpeakerActionProps) *awsevents.EventPattern {
	if err := d.validateVoiceIdSpeakerActionPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"voiceIdSpeakerActionPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DomainEvents) VoiceIdStartSessionActionPattern(options *DomainEvents_VoiceIdStartSessionAction_VoiceIdStartSessionActionProps) *awsevents.EventPattern {
	if err := d.validateVoiceIdStartSessionActionPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"voiceIdStartSessionActionPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DomainEvents) VoiceIdUpdateSessionActionPattern(options *DomainEvents_VoiceIdUpdateSessionAction_VoiceIdUpdateSessionActionProps) *awsevents.EventPattern {
	if err := d.validateVoiceIdUpdateSessionActionPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"voiceIdUpdateSessionActionPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

