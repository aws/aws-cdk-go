package mixinsawsguardduty

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnDetectorMixinProps",
		reflect.TypeOf((*CfnDetectorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnDetectorPropsMixin",
		reflect.TypeOf((*CfnDetectorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDetectorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnDetectorPropsMixin.CFNDataSourceConfigurationsProperty",
		reflect.TypeOf((*CfnDetectorPropsMixin_CFNDataSourceConfigurationsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnDetectorPropsMixin.CFNFeatureAdditionalConfigurationProperty",
		reflect.TypeOf((*CfnDetectorPropsMixin_CFNFeatureAdditionalConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnDetectorPropsMixin.CFNFeatureConfigurationProperty",
		reflect.TypeOf((*CfnDetectorPropsMixin_CFNFeatureConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnDetectorPropsMixin.CFNKubernetesAuditLogsConfigurationProperty",
		reflect.TypeOf((*CfnDetectorPropsMixin_CFNKubernetesAuditLogsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnDetectorPropsMixin.CFNKubernetesConfigurationProperty",
		reflect.TypeOf((*CfnDetectorPropsMixin_CFNKubernetesConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnDetectorPropsMixin.CFNMalwareProtectionConfigurationProperty",
		reflect.TypeOf((*CfnDetectorPropsMixin_CFNMalwareProtectionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnDetectorPropsMixin.CFNS3LogsConfigurationProperty",
		reflect.TypeOf((*CfnDetectorPropsMixin_CFNS3LogsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnDetectorPropsMixin.CFNScanEc2InstanceWithFindingsConfigurationProperty",
		reflect.TypeOf((*CfnDetectorPropsMixin_CFNScanEc2InstanceWithFindingsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnDetectorPropsMixin.TagItemProperty",
		reflect.TypeOf((*CfnDetectorPropsMixin_TagItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnFilterMixinProps",
		reflect.TypeOf((*CfnFilterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnFilterPropsMixin",
		reflect.TypeOf((*CfnFilterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFilterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnFilterPropsMixin.ConditionProperty",
		reflect.TypeOf((*CfnFilterPropsMixin_ConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnFilterPropsMixin.FindingCriteriaProperty",
		reflect.TypeOf((*CfnFilterPropsMixin_FindingCriteriaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnIPSetMixinProps",
		reflect.TypeOf((*CfnIPSetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnIPSetPropsMixin",
		reflect.TypeOf((*CfnIPSetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIPSetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnMalwareProtectionPlanMixinProps",
		reflect.TypeOf((*CfnMalwareProtectionPlanMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnMalwareProtectionPlanPropsMixin",
		reflect.TypeOf((*CfnMalwareProtectionPlanPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMalwareProtectionPlanPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnMalwareProtectionPlanPropsMixin.CFNActionsProperty",
		reflect.TypeOf((*CfnMalwareProtectionPlanPropsMixin_CFNActionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnMalwareProtectionPlanPropsMixin.CFNProtectedResourceProperty",
		reflect.TypeOf((*CfnMalwareProtectionPlanPropsMixin_CFNProtectedResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnMalwareProtectionPlanPropsMixin.CFNStatusReasonsProperty",
		reflect.TypeOf((*CfnMalwareProtectionPlanPropsMixin_CFNStatusReasonsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnMalwareProtectionPlanPropsMixin.CFNTaggingProperty",
		reflect.TypeOf((*CfnMalwareProtectionPlanPropsMixin_CFNTaggingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnMalwareProtectionPlanPropsMixin.S3BucketProperty",
		reflect.TypeOf((*CfnMalwareProtectionPlanPropsMixin_S3BucketProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnMalwareProtectionPlanPropsMixin.TagItemProperty",
		reflect.TypeOf((*CfnMalwareProtectionPlanPropsMixin_TagItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnMasterMixinProps",
		reflect.TypeOf((*CfnMasterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnMasterPropsMixin",
		reflect.TypeOf((*CfnMasterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMasterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnMemberMixinProps",
		reflect.TypeOf((*CfnMemberMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnMemberPropsMixin",
		reflect.TypeOf((*CfnMemberPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMemberPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnPublishingDestinationMixinProps",
		reflect.TypeOf((*CfnPublishingDestinationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnPublishingDestinationPropsMixin",
		reflect.TypeOf((*CfnPublishingDestinationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPublishingDestinationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnPublishingDestinationPropsMixin.CFNDestinationPropertiesProperty",
		reflect.TypeOf((*CfnPublishingDestinationPropsMixin_CFNDestinationPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnPublishingDestinationPropsMixin.TagItemProperty",
		reflect.TypeOf((*CfnPublishingDestinationPropsMixin_TagItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnThreatEntitySetMixinProps",
		reflect.TypeOf((*CfnThreatEntitySetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnThreatEntitySetPropsMixin",
		reflect.TypeOf((*CfnThreatEntitySetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnThreatEntitySetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnThreatEntitySetPropsMixin.TagItemProperty",
		reflect.TypeOf((*CfnThreatEntitySetPropsMixin_TagItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnThreatIntelSetMixinProps",
		reflect.TypeOf((*CfnThreatIntelSetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnThreatIntelSetPropsMixin",
		reflect.TypeOf((*CfnThreatIntelSetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnThreatIntelSetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnTrustedEntitySetMixinProps",
		reflect.TypeOf((*CfnTrustedEntitySetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnTrustedEntitySetPropsMixin",
		reflect.TypeOf((*CfnTrustedEntitySetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTrustedEntitySetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnTrustedEntitySetPropsMixin.TagItemProperty",
		reflect.TypeOf((*CfnTrustedEntitySetPropsMixin_TagItemProperty)(nil)).Elem(),
	)
}
