package awsaps

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnAnomalyDetectorMixinProps",
		reflect.TypeOf((*CfnAnomalyDetectorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnAnomalyDetectorPropsMixin",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAnomalyDetectorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnAnomalyDetectorPropsMixin.AnomalyDetectorConfigurationProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_AnomalyDetectorConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnAnomalyDetectorPropsMixin.IgnoreNearExpectedProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_IgnoreNearExpectedProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnAnomalyDetectorPropsMixin.LabelProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_LabelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnAnomalyDetectorPropsMixin.MissingDataActionProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_MissingDataActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnAnomalyDetectorPropsMixin.RandomCutForestConfigurationProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_RandomCutForestConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnResourcePolicyMixinProps",
		reflect.TypeOf((*CfnResourcePolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnResourcePolicyPropsMixin",
		reflect.TypeOf((*CfnResourcePolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnResourcePolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnRuleGroupsNamespaceMixinProps",
		reflect.TypeOf((*CfnRuleGroupsNamespaceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnRuleGroupsNamespacePropsMixin",
		reflect.TypeOf((*CfnRuleGroupsNamespacePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRuleGroupsNamespacePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnScraperMixinProps",
		reflect.TypeOf((*CfnScraperMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnScraperPropsMixin",
		reflect.TypeOf((*CfnScraperPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnScraperPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnScraperPropsMixin.AmpConfigurationProperty",
		reflect.TypeOf((*CfnScraperPropsMixin_AmpConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnScraperPropsMixin.CloudWatchLogDestinationProperty",
		reflect.TypeOf((*CfnScraperPropsMixin_CloudWatchLogDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnScraperPropsMixin.ComponentConfigProperty",
		reflect.TypeOf((*CfnScraperPropsMixin_ComponentConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnScraperPropsMixin.DestinationProperty",
		reflect.TypeOf((*CfnScraperPropsMixin_DestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnScraperPropsMixin.EksConfigurationProperty",
		reflect.TypeOf((*CfnScraperPropsMixin_EksConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnScraperPropsMixin.RoleConfigurationProperty",
		reflect.TypeOf((*CfnScraperPropsMixin_RoleConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnScraperPropsMixin.ScrapeConfigurationProperty",
		reflect.TypeOf((*CfnScraperPropsMixin_ScrapeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnScraperPropsMixin.ScraperComponentProperty",
		reflect.TypeOf((*CfnScraperPropsMixin_ScraperComponentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnScraperPropsMixin.ScraperLoggingConfigurationProperty",
		reflect.TypeOf((*CfnScraperPropsMixin_ScraperLoggingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnScraperPropsMixin.ScraperLoggingDestinationProperty",
		reflect.TypeOf((*CfnScraperPropsMixin_ScraperLoggingDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnScraperPropsMixin.SourceProperty",
		reflect.TypeOf((*CfnScraperPropsMixin_SourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnScraperPropsMixin.VpcConfigurationProperty",
		reflect.TypeOf((*CfnScraperPropsMixin_VpcConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnWorkspaceMixinProps",
		reflect.TypeOf((*CfnWorkspaceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnWorkspacePropsMixin",
		reflect.TypeOf((*CfnWorkspacePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWorkspacePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnWorkspacePropsMixin.CloudWatchLogDestinationProperty",
		reflect.TypeOf((*CfnWorkspacePropsMixin_CloudWatchLogDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnWorkspacePropsMixin.LabelProperty",
		reflect.TypeOf((*CfnWorkspacePropsMixin_LabelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnWorkspacePropsMixin.LimitsPerLabelSetEntryProperty",
		reflect.TypeOf((*CfnWorkspacePropsMixin_LimitsPerLabelSetEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnWorkspacePropsMixin.LimitsPerLabelSetProperty",
		reflect.TypeOf((*CfnWorkspacePropsMixin_LimitsPerLabelSetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnWorkspacePropsMixin.LoggingConfigurationProperty",
		reflect.TypeOf((*CfnWorkspacePropsMixin_LoggingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnWorkspacePropsMixin.LoggingDestinationProperty",
		reflect.TypeOf((*CfnWorkspacePropsMixin_LoggingDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnWorkspacePropsMixin.LoggingFilterProperty",
		reflect.TypeOf((*CfnWorkspacePropsMixin_LoggingFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnWorkspacePropsMixin.QueryLoggingConfigurationProperty",
		reflect.TypeOf((*CfnWorkspacePropsMixin_QueryLoggingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_aps.CfnWorkspacePropsMixin.WorkspaceConfigurationProperty",
		reflect.TypeOf((*CfnWorkspacePropsMixin_WorkspaceConfigurationProperty)(nil)).Elem(),
	)
}
