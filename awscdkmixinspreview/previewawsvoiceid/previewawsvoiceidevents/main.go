package previewawsvoiceidevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents",
		reflect.TypeOf((*DomainEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "voiceIdBatchFraudsterRegistrationActionPattern", GoMethod: "VoiceIdBatchFraudsterRegistrationActionPattern"},
			_jsii_.MemberMethod{JsiiMethod: "voiceIdBatchSpeakerEnrollmentActionPattern", GoMethod: "VoiceIdBatchSpeakerEnrollmentActionPattern"},
			_jsii_.MemberMethod{JsiiMethod: "voiceIdEvaluateSessionActionPattern", GoMethod: "VoiceIdEvaluateSessionActionPattern"},
			_jsii_.MemberMethod{JsiiMethod: "voiceIdFraudsterActionPattern", GoMethod: "VoiceIdFraudsterActionPattern"},
			_jsii_.MemberMethod{JsiiMethod: "voiceIdSessionSpeakerEnrollmentActionPattern", GoMethod: "VoiceIdSessionSpeakerEnrollmentActionPattern"},
			_jsii_.MemberMethod{JsiiMethod: "voiceIdSpeakerActionPattern", GoMethod: "VoiceIdSpeakerActionPattern"},
			_jsii_.MemberMethod{JsiiMethod: "voiceIdStartSessionActionPattern", GoMethod: "VoiceIdStartSessionActionPattern"},
			_jsii_.MemberMethod{JsiiMethod: "voiceIdUpdateSessionActionPattern", GoMethod: "VoiceIdUpdateSessionActionPattern"},
		},
		func() interface{} {
			return &jsiiProxy_DomainEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdBatchFraudsterRegistrationAction",
		reflect.TypeOf((*DomainEvents_VoiceIdBatchFraudsterRegistrationAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DomainEvents_VoiceIdBatchFraudsterRegistrationAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdBatchFraudsterRegistrationAction.Data",
		reflect.TypeOf((*DomainEvents_VoiceIdBatchFraudsterRegistrationAction_Data)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdBatchFraudsterRegistrationAction.ErrorInfo",
		reflect.TypeOf((*DomainEvents_VoiceIdBatchFraudsterRegistrationAction_ErrorInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdBatchFraudsterRegistrationAction.InputDataConfig",
		reflect.TypeOf((*DomainEvents_VoiceIdBatchFraudsterRegistrationAction_InputDataConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdBatchFraudsterRegistrationAction.OutputDataConfig",
		reflect.TypeOf((*DomainEvents_VoiceIdBatchFraudsterRegistrationAction_OutputDataConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdBatchFraudsterRegistrationAction.RegistrationConfig",
		reflect.TypeOf((*DomainEvents_VoiceIdBatchFraudsterRegistrationAction_RegistrationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdBatchFraudsterRegistrationAction.VoiceIdBatchFraudsterRegistrationActionProps",
		reflect.TypeOf((*DomainEvents_VoiceIdBatchFraudsterRegistrationAction_VoiceIdBatchFraudsterRegistrationActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdBatchSpeakerEnrollmentAction",
		reflect.TypeOf((*DomainEvents_VoiceIdBatchSpeakerEnrollmentAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DomainEvents_VoiceIdBatchSpeakerEnrollmentAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdBatchSpeakerEnrollmentAction.Data",
		reflect.TypeOf((*DomainEvents_VoiceIdBatchSpeakerEnrollmentAction_Data)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdBatchSpeakerEnrollmentAction.EnrollmentConfig",
		reflect.TypeOf((*DomainEvents_VoiceIdBatchSpeakerEnrollmentAction_EnrollmentConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdBatchSpeakerEnrollmentAction.ErrorInfo",
		reflect.TypeOf((*DomainEvents_VoiceIdBatchSpeakerEnrollmentAction_ErrorInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdBatchSpeakerEnrollmentAction.FraudDetectionConfig",
		reflect.TypeOf((*DomainEvents_VoiceIdBatchSpeakerEnrollmentAction_FraudDetectionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdBatchSpeakerEnrollmentAction.InputDataConfig",
		reflect.TypeOf((*DomainEvents_VoiceIdBatchSpeakerEnrollmentAction_InputDataConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdBatchSpeakerEnrollmentAction.OutputDataConfig",
		reflect.TypeOf((*DomainEvents_VoiceIdBatchSpeakerEnrollmentAction_OutputDataConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdBatchSpeakerEnrollmentAction.VoiceIdBatchSpeakerEnrollmentActionProps",
		reflect.TypeOf((*DomainEvents_VoiceIdBatchSpeakerEnrollmentAction_VoiceIdBatchSpeakerEnrollmentActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdEvaluateSessionAction",
		reflect.TypeOf((*DomainEvents_VoiceIdEvaluateSessionAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DomainEvents_VoiceIdEvaluateSessionAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdEvaluateSessionAction.AuthenticationResult",
		reflect.TypeOf((*DomainEvents_VoiceIdEvaluateSessionAction_AuthenticationResult)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdEvaluateSessionAction.ConfigurationAuthentication",
		reflect.TypeOf((*DomainEvents_VoiceIdEvaluateSessionAction_ConfigurationAuthentication)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdEvaluateSessionAction.ConfigurationFraud",
		reflect.TypeOf((*DomainEvents_VoiceIdEvaluateSessionAction_ConfigurationFraud)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdEvaluateSessionAction.ErrorInfo",
		reflect.TypeOf((*DomainEvents_VoiceIdEvaluateSessionAction_ErrorInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdEvaluateSessionAction.FraudDetectionResult",
		reflect.TypeOf((*DomainEvents_VoiceIdEvaluateSessionAction_FraudDetectionResult)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdEvaluateSessionAction.KnownFraudsterRisk",
		reflect.TypeOf((*DomainEvents_VoiceIdEvaluateSessionAction_KnownFraudsterRisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdEvaluateSessionAction.RiskDetails",
		reflect.TypeOf((*DomainEvents_VoiceIdEvaluateSessionAction_RiskDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdEvaluateSessionAction.Session",
		reflect.TypeOf((*DomainEvents_VoiceIdEvaluateSessionAction_Session)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdEvaluateSessionAction.SystemAttributes",
		reflect.TypeOf((*DomainEvents_VoiceIdEvaluateSessionAction_SystemAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdEvaluateSessionAction.VoiceIdEvaluateSessionActionProps",
		reflect.TypeOf((*DomainEvents_VoiceIdEvaluateSessionAction_VoiceIdEvaluateSessionActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdEvaluateSessionAction.VoiceSpoofingRisk",
		reflect.TypeOf((*DomainEvents_VoiceIdEvaluateSessionAction_VoiceSpoofingRisk)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdFraudsterAction",
		reflect.TypeOf((*DomainEvents_VoiceIdFraudsterAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DomainEvents_VoiceIdFraudsterAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdFraudsterAction.Data",
		reflect.TypeOf((*DomainEvents_VoiceIdFraudsterAction_Data)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdFraudsterAction.ErrorInfo",
		reflect.TypeOf((*DomainEvents_VoiceIdFraudsterAction_ErrorInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdFraudsterAction.VoiceIdFraudsterActionProps",
		reflect.TypeOf((*DomainEvents_VoiceIdFraudsterAction_VoiceIdFraudsterActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdSessionSpeakerEnrollmentAction",
		reflect.TypeOf((*DomainEvents_VoiceIdSessionSpeakerEnrollmentAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DomainEvents_VoiceIdSessionSpeakerEnrollmentAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdSessionSpeakerEnrollmentAction.ErrorInfo",
		reflect.TypeOf((*DomainEvents_VoiceIdSessionSpeakerEnrollmentAction_ErrorInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdSessionSpeakerEnrollmentAction.SystemAttributes",
		reflect.TypeOf((*DomainEvents_VoiceIdSessionSpeakerEnrollmentAction_SystemAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdSessionSpeakerEnrollmentAction.VoiceIdSessionSpeakerEnrollmentActionProps",
		reflect.TypeOf((*DomainEvents_VoiceIdSessionSpeakerEnrollmentAction_VoiceIdSessionSpeakerEnrollmentActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdSpeakerAction",
		reflect.TypeOf((*DomainEvents_VoiceIdSpeakerAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DomainEvents_VoiceIdSpeakerAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdSpeakerAction.Data",
		reflect.TypeOf((*DomainEvents_VoiceIdSpeakerAction_Data)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdSpeakerAction.ErrorInfo",
		reflect.TypeOf((*DomainEvents_VoiceIdSpeakerAction_ErrorInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdSpeakerAction.SystemAttributes",
		reflect.TypeOf((*DomainEvents_VoiceIdSpeakerAction_SystemAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdSpeakerAction.VoiceIdSpeakerActionProps",
		reflect.TypeOf((*DomainEvents_VoiceIdSpeakerAction_VoiceIdSpeakerActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdStartSessionAction",
		reflect.TypeOf((*DomainEvents_VoiceIdStartSessionAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DomainEvents_VoiceIdStartSessionAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdStartSessionAction.AuthenticationAudioProgress",
		reflect.TypeOf((*DomainEvents_VoiceIdStartSessionAction_AuthenticationAudioProgress)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdStartSessionAction.AuthenticationConfiguration",
		reflect.TypeOf((*DomainEvents_VoiceIdStartSessionAction_AuthenticationConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdStartSessionAction.EnrollmentAudioProgress",
		reflect.TypeOf((*DomainEvents_VoiceIdStartSessionAction_EnrollmentAudioProgress)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdStartSessionAction.ErrorInfo",
		reflect.TypeOf((*DomainEvents_VoiceIdStartSessionAction_ErrorInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdStartSessionAction.FraudDetectionConfiguration",
		reflect.TypeOf((*DomainEvents_VoiceIdStartSessionAction_FraudDetectionConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdStartSessionAction.Session",
		reflect.TypeOf((*DomainEvents_VoiceIdStartSessionAction_Session)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdStartSessionAction.StreamingConfiguration",
		reflect.TypeOf((*DomainEvents_VoiceIdStartSessionAction_StreamingConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdStartSessionAction.SystemAttributes",
		reflect.TypeOf((*DomainEvents_VoiceIdStartSessionAction_SystemAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdStartSessionAction.VoiceIdStartSessionActionProps",
		reflect.TypeOf((*DomainEvents_VoiceIdStartSessionAction_VoiceIdStartSessionActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdUpdateSessionAction",
		reflect.TypeOf((*DomainEvents_VoiceIdUpdateSessionAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DomainEvents_VoiceIdUpdateSessionAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdUpdateSessionAction.AuthenticationConfiguration",
		reflect.TypeOf((*DomainEvents_VoiceIdUpdateSessionAction_AuthenticationConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdUpdateSessionAction.ErrorInfo",
		reflect.TypeOf((*DomainEvents_VoiceIdUpdateSessionAction_ErrorInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdUpdateSessionAction.FraudDetectionConfiguration",
		reflect.TypeOf((*DomainEvents_VoiceIdUpdateSessionAction_FraudDetectionConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdUpdateSessionAction.Session",
		reflect.TypeOf((*DomainEvents_VoiceIdUpdateSessionAction_Session)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdUpdateSessionAction.VoiceIdUpdateSessionActionProps",
		reflect.TypeOf((*DomainEvents_VoiceIdUpdateSessionAction_VoiceIdUpdateSessionActionProps)(nil)).Elem(),
	)
}
