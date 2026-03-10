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
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdBatchFraudsterRegistrationAction",
		reflect.TypeOf((*VoiceIdBatchFraudsterRegistrationAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_VoiceIdBatchFraudsterRegistrationAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdBatchFraudsterRegistrationAction.Data",
		reflect.TypeOf((*VoiceIdBatchFraudsterRegistrationAction_Data)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdBatchFraudsterRegistrationAction.ErrorInfo",
		reflect.TypeOf((*VoiceIdBatchFraudsterRegistrationAction_ErrorInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdBatchFraudsterRegistrationAction.InputDataConfig",
		reflect.TypeOf((*VoiceIdBatchFraudsterRegistrationAction_InputDataConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdBatchFraudsterRegistrationAction.OutputDataConfig",
		reflect.TypeOf((*VoiceIdBatchFraudsterRegistrationAction_OutputDataConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdBatchFraudsterRegistrationAction.RegistrationConfig",
		reflect.TypeOf((*VoiceIdBatchFraudsterRegistrationAction_RegistrationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdBatchFraudsterRegistrationAction.VoiceIdBatchFraudsterRegistrationActionProps",
		reflect.TypeOf((*VoiceIdBatchFraudsterRegistrationAction_VoiceIdBatchFraudsterRegistrationActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdBatchSpeakerEnrollmentAction",
		reflect.TypeOf((*VoiceIdBatchSpeakerEnrollmentAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_VoiceIdBatchSpeakerEnrollmentAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdBatchSpeakerEnrollmentAction.Data",
		reflect.TypeOf((*VoiceIdBatchSpeakerEnrollmentAction_Data)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdBatchSpeakerEnrollmentAction.EnrollmentConfig",
		reflect.TypeOf((*VoiceIdBatchSpeakerEnrollmentAction_EnrollmentConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdBatchSpeakerEnrollmentAction.ErrorInfo",
		reflect.TypeOf((*VoiceIdBatchSpeakerEnrollmentAction_ErrorInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdBatchSpeakerEnrollmentAction.FraudDetectionConfig",
		reflect.TypeOf((*VoiceIdBatchSpeakerEnrollmentAction_FraudDetectionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdBatchSpeakerEnrollmentAction.InputDataConfig",
		reflect.TypeOf((*VoiceIdBatchSpeakerEnrollmentAction_InputDataConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdBatchSpeakerEnrollmentAction.OutputDataConfig",
		reflect.TypeOf((*VoiceIdBatchSpeakerEnrollmentAction_OutputDataConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdBatchSpeakerEnrollmentAction.VoiceIdBatchSpeakerEnrollmentActionProps",
		reflect.TypeOf((*VoiceIdBatchSpeakerEnrollmentAction_VoiceIdBatchSpeakerEnrollmentActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdEvaluateSessionAction",
		reflect.TypeOf((*VoiceIdEvaluateSessionAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_VoiceIdEvaluateSessionAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdEvaluateSessionAction.AuthenticationResult",
		reflect.TypeOf((*VoiceIdEvaluateSessionAction_AuthenticationResult)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdEvaluateSessionAction.ConfigurationAuthentication",
		reflect.TypeOf((*VoiceIdEvaluateSessionAction_ConfigurationAuthentication)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdEvaluateSessionAction.ConfigurationFraud",
		reflect.TypeOf((*VoiceIdEvaluateSessionAction_ConfigurationFraud)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdEvaluateSessionAction.ErrorInfo",
		reflect.TypeOf((*VoiceIdEvaluateSessionAction_ErrorInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdEvaluateSessionAction.FraudDetectionResult",
		reflect.TypeOf((*VoiceIdEvaluateSessionAction_FraudDetectionResult)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdEvaluateSessionAction.KnownFraudsterRisk",
		reflect.TypeOf((*VoiceIdEvaluateSessionAction_KnownFraudsterRisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdEvaluateSessionAction.RiskDetails",
		reflect.TypeOf((*VoiceIdEvaluateSessionAction_RiskDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdEvaluateSessionAction.Session",
		reflect.TypeOf((*VoiceIdEvaluateSessionAction_Session)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdEvaluateSessionAction.SystemAttributes",
		reflect.TypeOf((*VoiceIdEvaluateSessionAction_SystemAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdEvaluateSessionAction.VoiceIdEvaluateSessionActionProps",
		reflect.TypeOf((*VoiceIdEvaluateSessionAction_VoiceIdEvaluateSessionActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdEvaluateSessionAction.VoiceSpoofingRisk",
		reflect.TypeOf((*VoiceIdEvaluateSessionAction_VoiceSpoofingRisk)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdFraudsterAction",
		reflect.TypeOf((*VoiceIdFraudsterAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_VoiceIdFraudsterAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdFraudsterAction.Data",
		reflect.TypeOf((*VoiceIdFraudsterAction_Data)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdFraudsterAction.ErrorInfo",
		reflect.TypeOf((*VoiceIdFraudsterAction_ErrorInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdFraudsterAction.VoiceIdFraudsterActionProps",
		reflect.TypeOf((*VoiceIdFraudsterAction_VoiceIdFraudsterActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdSessionSpeakerEnrollmentAction",
		reflect.TypeOf((*VoiceIdSessionSpeakerEnrollmentAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_VoiceIdSessionSpeakerEnrollmentAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdSessionSpeakerEnrollmentAction.ErrorInfo",
		reflect.TypeOf((*VoiceIdSessionSpeakerEnrollmentAction_ErrorInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdSessionSpeakerEnrollmentAction.SystemAttributes",
		reflect.TypeOf((*VoiceIdSessionSpeakerEnrollmentAction_SystemAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdSessionSpeakerEnrollmentAction.VoiceIdSessionSpeakerEnrollmentActionProps",
		reflect.TypeOf((*VoiceIdSessionSpeakerEnrollmentAction_VoiceIdSessionSpeakerEnrollmentActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdSpeakerAction",
		reflect.TypeOf((*VoiceIdSpeakerAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_VoiceIdSpeakerAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdSpeakerAction.Data",
		reflect.TypeOf((*VoiceIdSpeakerAction_Data)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdSpeakerAction.ErrorInfo",
		reflect.TypeOf((*VoiceIdSpeakerAction_ErrorInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdSpeakerAction.SystemAttributes",
		reflect.TypeOf((*VoiceIdSpeakerAction_SystemAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdSpeakerAction.VoiceIdSpeakerActionProps",
		reflect.TypeOf((*VoiceIdSpeakerAction_VoiceIdSpeakerActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdStartSessionAction",
		reflect.TypeOf((*VoiceIdStartSessionAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_VoiceIdStartSessionAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdStartSessionAction.AuthenticationAudioProgress",
		reflect.TypeOf((*VoiceIdStartSessionAction_AuthenticationAudioProgress)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdStartSessionAction.AuthenticationConfiguration",
		reflect.TypeOf((*VoiceIdStartSessionAction_AuthenticationConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdStartSessionAction.EnrollmentAudioProgress",
		reflect.TypeOf((*VoiceIdStartSessionAction_EnrollmentAudioProgress)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdStartSessionAction.ErrorInfo",
		reflect.TypeOf((*VoiceIdStartSessionAction_ErrorInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdStartSessionAction.FraudDetectionConfiguration",
		reflect.TypeOf((*VoiceIdStartSessionAction_FraudDetectionConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdStartSessionAction.Session",
		reflect.TypeOf((*VoiceIdStartSessionAction_Session)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdStartSessionAction.StreamingConfiguration",
		reflect.TypeOf((*VoiceIdStartSessionAction_StreamingConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdStartSessionAction.SystemAttributes",
		reflect.TypeOf((*VoiceIdStartSessionAction_SystemAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdStartSessionAction.VoiceIdStartSessionActionProps",
		reflect.TypeOf((*VoiceIdStartSessionAction_VoiceIdStartSessionActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdUpdateSessionAction",
		reflect.TypeOf((*VoiceIdUpdateSessionAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_VoiceIdUpdateSessionAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdUpdateSessionAction.AuthenticationConfiguration",
		reflect.TypeOf((*VoiceIdUpdateSessionAction_AuthenticationConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdUpdateSessionAction.ErrorInfo",
		reflect.TypeOf((*VoiceIdUpdateSessionAction_ErrorInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdUpdateSessionAction.FraudDetectionConfiguration",
		reflect.TypeOf((*VoiceIdUpdateSessionAction_FraudDetectionConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdUpdateSessionAction.Session",
		reflect.TypeOf((*VoiceIdUpdateSessionAction_Session)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdUpdateSessionAction.VoiceIdUpdateSessionActionProps",
		reflect.TypeOf((*VoiceIdUpdateSessionAction_VoiceIdUpdateSessionActionProps)(nil)).Elem(),
	)
}
