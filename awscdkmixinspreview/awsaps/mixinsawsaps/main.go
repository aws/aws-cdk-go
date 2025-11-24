package mixinsawsaps

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnAnomalyDetectorMixinProps",
		reflect.TypeOf((*CfnAnomalyDetectorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnAnomalyDetectorPropsMixin",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAnomalyDetectorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnAnomalyDetectorPropsMixin.AnomalyDetectorConfigurationProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_AnomalyDetectorConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnAnomalyDetectorPropsMixin.IgnoreNearExpectedProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_IgnoreNearExpectedProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnAnomalyDetectorPropsMixin.LabelProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_LabelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnAnomalyDetectorPropsMixin.MissingDataActionProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_MissingDataActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnAnomalyDetectorPropsMixin.RandomCutForestConfigurationProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_RandomCutForestConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnResourcePolicyMixinProps",
		reflect.TypeOf((*CfnResourcePolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnResourcePolicyPropsMixin",
		reflect.TypeOf((*CfnResourcePolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnResourcePolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnRuleGroupsNamespaceMixinProps",
		reflect.TypeOf((*CfnRuleGroupsNamespaceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnRuleGroupsNamespacePropsMixin",
		reflect.TypeOf((*CfnRuleGroupsNamespacePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRuleGroupsNamespacePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperMixinProps",
		reflect.TypeOf((*CfnScraperMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperPropsMixin",
		reflect.TypeOf((*CfnScraperPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnScraperPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperPropsMixin.AmpConfigurationProperty",
		reflect.TypeOf((*CfnScraperPropsMixin_AmpConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperPropsMixin.CloudWatchLogDestinationProperty",
		reflect.TypeOf((*CfnScraperPropsMixin_CloudWatchLogDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperPropsMixin.ComponentConfigProperty",
		reflect.TypeOf((*CfnScraperPropsMixin_ComponentConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperPropsMixin.DestinationProperty",
		reflect.TypeOf((*CfnScraperPropsMixin_DestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperPropsMixin.EksConfigurationProperty",
		reflect.TypeOf((*CfnScraperPropsMixin_EksConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperPropsMixin.RoleConfigurationProperty",
		reflect.TypeOf((*CfnScraperPropsMixin_RoleConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperPropsMixin.ScrapeConfigurationProperty",
		reflect.TypeOf((*CfnScraperPropsMixin_ScrapeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperPropsMixin.ScraperComponentProperty",
		reflect.TypeOf((*CfnScraperPropsMixin_ScraperComponentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperPropsMixin.ScraperLoggingConfigurationProperty",
		reflect.TypeOf((*CfnScraperPropsMixin_ScraperLoggingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperPropsMixin.ScraperLoggingDestinationProperty",
		reflect.TypeOf((*CfnScraperPropsMixin_ScraperLoggingDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperPropsMixin.SourceProperty",
		reflect.TypeOf((*CfnScraperPropsMixin_SourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperPropsMixin.VpcConfigurationProperty",
		reflect.TypeOf((*CfnScraperPropsMixin_VpcConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnWorkspaceMixinProps",
		reflect.TypeOf((*CfnWorkspaceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnWorkspacePropsMixin",
		reflect.TypeOf((*CfnWorkspacePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWorkspacePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnWorkspacePropsMixin.CloudWatchLogDestinationProperty",
		reflect.TypeOf((*CfnWorkspacePropsMixin_CloudWatchLogDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnWorkspacePropsMixin.LabelProperty",
		reflect.TypeOf((*CfnWorkspacePropsMixin_LabelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnWorkspacePropsMixin.LimitsPerLabelSetEntryProperty",
		reflect.TypeOf((*CfnWorkspacePropsMixin_LimitsPerLabelSetEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnWorkspacePropsMixin.LimitsPerLabelSetProperty",
		reflect.TypeOf((*CfnWorkspacePropsMixin_LimitsPerLabelSetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnWorkspacePropsMixin.LoggingConfigurationProperty",
		reflect.TypeOf((*CfnWorkspacePropsMixin_LoggingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnWorkspacePropsMixin.LoggingDestinationProperty",
		reflect.TypeOf((*CfnWorkspacePropsMixin_LoggingDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnWorkspacePropsMixin.LoggingFilterProperty",
		reflect.TypeOf((*CfnWorkspacePropsMixin_LoggingFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnWorkspacePropsMixin.QueryLoggingConfigurationProperty",
		reflect.TypeOf((*CfnWorkspacePropsMixin_QueryLoggingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnWorkspacePropsMixin.WorkspaceConfigurationProperty",
		reflect.TypeOf((*CfnWorkspacePropsMixin_WorkspaceConfigurationProperty)(nil)).Elem(),
	)
}
