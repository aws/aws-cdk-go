package previewawsb2bimixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnCapabilityMixinProps",
		reflect.TypeOf((*CfnCapabilityMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnCapabilityPropsMixin",
		reflect.TypeOf((*CfnCapabilityPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCapabilityPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnCapabilityPropsMixin.CapabilityConfigurationProperty",
		reflect.TypeOf((*CfnCapabilityPropsMixin_CapabilityConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnCapabilityPropsMixin.EdiConfigurationProperty",
		reflect.TypeOf((*CfnCapabilityPropsMixin_EdiConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnCapabilityPropsMixin.EdiTypeProperty",
		reflect.TypeOf((*CfnCapabilityPropsMixin_EdiTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnCapabilityPropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnCapabilityPropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnCapabilityPropsMixin.X12DetailsProperty",
		reflect.TypeOf((*CfnCapabilityPropsMixin_X12DetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnPartnershipMixinProps",
		reflect.TypeOf((*CfnPartnershipMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnPartnershipPropsMixin",
		reflect.TypeOf((*CfnPartnershipPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPartnershipPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnPartnershipPropsMixin.CapabilityOptionsProperty",
		reflect.TypeOf((*CfnPartnershipPropsMixin_CapabilityOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnPartnershipPropsMixin.InboundEdiOptionsProperty",
		reflect.TypeOf((*CfnPartnershipPropsMixin_InboundEdiOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnPartnershipPropsMixin.OutboundEdiOptionsProperty",
		reflect.TypeOf((*CfnPartnershipPropsMixin_OutboundEdiOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnPartnershipPropsMixin.WrapOptionsProperty",
		reflect.TypeOf((*CfnPartnershipPropsMixin_WrapOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnPartnershipPropsMixin.X12AcknowledgmentOptionsProperty",
		reflect.TypeOf((*CfnPartnershipPropsMixin_X12AcknowledgmentOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnPartnershipPropsMixin.X12ControlNumbersProperty",
		reflect.TypeOf((*CfnPartnershipPropsMixin_X12ControlNumbersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnPartnershipPropsMixin.X12DelimitersProperty",
		reflect.TypeOf((*CfnPartnershipPropsMixin_X12DelimitersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnPartnershipPropsMixin.X12EnvelopeProperty",
		reflect.TypeOf((*CfnPartnershipPropsMixin_X12EnvelopeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnPartnershipPropsMixin.X12FunctionalGroupHeadersProperty",
		reflect.TypeOf((*CfnPartnershipPropsMixin_X12FunctionalGroupHeadersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnPartnershipPropsMixin.X12InboundEdiOptionsProperty",
		reflect.TypeOf((*CfnPartnershipPropsMixin_X12InboundEdiOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnPartnershipPropsMixin.X12InterchangeControlHeadersProperty",
		reflect.TypeOf((*CfnPartnershipPropsMixin_X12InterchangeControlHeadersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnPartnershipPropsMixin.X12OutboundEdiHeadersProperty",
		reflect.TypeOf((*CfnPartnershipPropsMixin_X12OutboundEdiHeadersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnProfileMixinProps",
		reflect.TypeOf((*CfnProfileMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnProfilePropsMixin",
		reflect.TypeOf((*CfnProfilePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProfilePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerB2biExecutionLogs",
		reflect.TypeOf((*CfnTransformerB2biExecutionLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnTransformerB2biExecutionLogs{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerLogsMixin",
		reflect.TypeOf((*CfnTransformerLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTransformerLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerMixinProps",
		reflect.TypeOf((*CfnTransformerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerPropsMixin",
		reflect.TypeOf((*CfnTransformerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTransformerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerPropsMixin.AdvancedOptionsProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_AdvancedOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerPropsMixin.EdiTypeProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_EdiTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerPropsMixin.FormatOptionsProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_FormatOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerPropsMixin.InputConversionProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_InputConversionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerPropsMixin.MappingProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_MappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerPropsMixin.OutputConversionProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_OutputConversionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerPropsMixin.SampleDocumentKeysProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_SampleDocumentKeysProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerPropsMixin.SampleDocumentsProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_SampleDocumentsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerPropsMixin.X12AdvancedOptionsProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_X12AdvancedOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerPropsMixin.X12CodeListValidationRuleProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_X12CodeListValidationRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerPropsMixin.X12DetailsProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_X12DetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerPropsMixin.X12ElementLengthValidationRuleProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_X12ElementLengthValidationRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerPropsMixin.X12ElementRequirementValidationRuleProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_X12ElementRequirementValidationRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerPropsMixin.X12SplitOptionsProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_X12SplitOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerPropsMixin.X12ValidationOptionsProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_X12ValidationOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerPropsMixin.X12ValidationRuleProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_X12ValidationRuleProperty)(nil)).Elem(),
	)
}
