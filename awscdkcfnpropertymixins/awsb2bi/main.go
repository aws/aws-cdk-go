package awsb2bi

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnCapabilityMixinProps",
		reflect.TypeOf((*CfnCapabilityMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnCapabilityPropsMixin",
		reflect.TypeOf((*CfnCapabilityPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCapabilityPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnCapabilityPropsMixin.CapabilityConfigurationProperty",
		reflect.TypeOf((*CfnCapabilityPropsMixin_CapabilityConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnCapabilityPropsMixin.EdiConfigurationProperty",
		reflect.TypeOf((*CfnCapabilityPropsMixin_EdiConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnCapabilityPropsMixin.EdiTypeProperty",
		reflect.TypeOf((*CfnCapabilityPropsMixin_EdiTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnCapabilityPropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnCapabilityPropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnCapabilityPropsMixin.X12DetailsProperty",
		reflect.TypeOf((*CfnCapabilityPropsMixin_X12DetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnPartnershipMixinProps",
		reflect.TypeOf((*CfnPartnershipMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnPartnershipPropsMixin",
		reflect.TypeOf((*CfnPartnershipPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPartnershipPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnPartnershipPropsMixin.CapabilityOptionsProperty",
		reflect.TypeOf((*CfnPartnershipPropsMixin_CapabilityOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnPartnershipPropsMixin.InboundEdiOptionsProperty",
		reflect.TypeOf((*CfnPartnershipPropsMixin_InboundEdiOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnPartnershipPropsMixin.OutboundEdiOptionsProperty",
		reflect.TypeOf((*CfnPartnershipPropsMixin_OutboundEdiOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnPartnershipPropsMixin.WrapOptionsProperty",
		reflect.TypeOf((*CfnPartnershipPropsMixin_WrapOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnPartnershipPropsMixin.X12AcknowledgmentOptionsProperty",
		reflect.TypeOf((*CfnPartnershipPropsMixin_X12AcknowledgmentOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnPartnershipPropsMixin.X12ControlNumbersProperty",
		reflect.TypeOf((*CfnPartnershipPropsMixin_X12ControlNumbersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnPartnershipPropsMixin.X12DelimitersProperty",
		reflect.TypeOf((*CfnPartnershipPropsMixin_X12DelimitersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnPartnershipPropsMixin.X12EnvelopeProperty",
		reflect.TypeOf((*CfnPartnershipPropsMixin_X12EnvelopeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnPartnershipPropsMixin.X12FunctionalGroupHeadersProperty",
		reflect.TypeOf((*CfnPartnershipPropsMixin_X12FunctionalGroupHeadersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnPartnershipPropsMixin.X12InboundEdiOptionsProperty",
		reflect.TypeOf((*CfnPartnershipPropsMixin_X12InboundEdiOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnPartnershipPropsMixin.X12InterchangeControlHeadersProperty",
		reflect.TypeOf((*CfnPartnershipPropsMixin_X12InterchangeControlHeadersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnPartnershipPropsMixin.X12OutboundEdiHeadersProperty",
		reflect.TypeOf((*CfnPartnershipPropsMixin_X12OutboundEdiHeadersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnProfileMixinProps",
		reflect.TypeOf((*CfnProfileMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnProfilePropsMixin",
		reflect.TypeOf((*CfnProfilePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProfilePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnTransformerMixinProps",
		reflect.TypeOf((*CfnTransformerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnTransformerPropsMixin",
		reflect.TypeOf((*CfnTransformerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTransformerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnTransformerPropsMixin.AdvancedOptionsProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_AdvancedOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnTransformerPropsMixin.EdiTypeProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_EdiTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnTransformerPropsMixin.FormatOptionsProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_FormatOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnTransformerPropsMixin.InputConversionProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_InputConversionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnTransformerPropsMixin.MappingProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_MappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnTransformerPropsMixin.OutputConversionProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_OutputConversionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnTransformerPropsMixin.SampleDocumentKeysProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_SampleDocumentKeysProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnTransformerPropsMixin.SampleDocumentsProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_SampleDocumentsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnTransformerPropsMixin.X12AdvancedOptionsProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_X12AdvancedOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnTransformerPropsMixin.X12CodeListValidationRuleProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_X12CodeListValidationRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnTransformerPropsMixin.X12DetailsProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_X12DetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnTransformerPropsMixin.X12ElementLengthValidationRuleProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_X12ElementLengthValidationRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnTransformerPropsMixin.X12ElementRequirementValidationRuleProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_X12ElementRequirementValidationRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnTransformerPropsMixin.X12SplitOptionsProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_X12SplitOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnTransformerPropsMixin.X12ValidationOptionsProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_X12ValidationOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_b2bi.CfnTransformerPropsMixin.X12ValidationRuleProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_X12ValidationRuleProperty)(nil)).Elem(),
	)
}
