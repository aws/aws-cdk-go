package previewawslogsmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnAccountPolicyMixinProps",
		reflect.TypeOf((*CfnAccountPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnAccountPolicyPropsMixin",
		reflect.TypeOf((*CfnAccountPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAccountPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnDeliveryDestinationMixinProps",
		reflect.TypeOf((*CfnDeliveryDestinationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnDeliveryDestinationPropsMixin",
		reflect.TypeOf((*CfnDeliveryDestinationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDeliveryDestinationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnDeliveryDestinationPropsMixin.DestinationPolicyProperty",
		reflect.TypeOf((*CfnDeliveryDestinationPropsMixin_DestinationPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnDeliveryMixinProps",
		reflect.TypeOf((*CfnDeliveryMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnDeliveryPropsMixin",
		reflect.TypeOf((*CfnDeliveryPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDeliveryPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnDeliverySourceMixinProps",
		reflect.TypeOf((*CfnDeliverySourceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnDeliverySourcePropsMixin",
		reflect.TypeOf((*CfnDeliverySourcePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDeliverySourcePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnDestinationMixinProps",
		reflect.TypeOf((*CfnDestinationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnDestinationPropsMixin",
		reflect.TypeOf((*CfnDestinationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDestinationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnIntegrationMixinProps",
		reflect.TypeOf((*CfnIntegrationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnIntegrationPropsMixin",
		reflect.TypeOf((*CfnIntegrationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIntegrationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnIntegrationPropsMixin.OpenSearchResourceConfigProperty",
		reflect.TypeOf((*CfnIntegrationPropsMixin_OpenSearchResourceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnIntegrationPropsMixin.ResourceConfigProperty",
		reflect.TypeOf((*CfnIntegrationPropsMixin_ResourceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnLogAnomalyDetectorMixinProps",
		reflect.TypeOf((*CfnLogAnomalyDetectorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnLogAnomalyDetectorPropsMixin",
		reflect.TypeOf((*CfnLogAnomalyDetectorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLogAnomalyDetectorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnLogGroupMixinProps",
		reflect.TypeOf((*CfnLogGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnLogGroupPropsMixin",
		reflect.TypeOf((*CfnLogGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLogGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnLogStreamMixinProps",
		reflect.TypeOf((*CfnLogStreamMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnLogStreamPropsMixin",
		reflect.TypeOf((*CfnLogStreamPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLogStreamPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnMetricFilterMixinProps",
		reflect.TypeOf((*CfnMetricFilterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnMetricFilterPropsMixin",
		reflect.TypeOf((*CfnMetricFilterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMetricFilterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnMetricFilterPropsMixin.DimensionProperty",
		reflect.TypeOf((*CfnMetricFilterPropsMixin_DimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnMetricFilterPropsMixin.MetricTransformationProperty",
		reflect.TypeOf((*CfnMetricFilterPropsMixin_MetricTransformationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnQueryDefinitionMixinProps",
		reflect.TypeOf((*CfnQueryDefinitionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnQueryDefinitionPropsMixin",
		reflect.TypeOf((*CfnQueryDefinitionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnQueryDefinitionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnResourcePolicyMixinProps",
		reflect.TypeOf((*CfnResourcePolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnResourcePolicyPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnScheduledQueryMixinProps",
		reflect.TypeOf((*CfnScheduledQueryMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnScheduledQueryPropsMixin",
		reflect.TypeOf((*CfnScheduledQueryPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnScheduledQueryPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnScheduledQueryPropsMixin.DestinationConfigurationProperty",
		reflect.TypeOf((*CfnScheduledQueryPropsMixin_DestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnScheduledQueryPropsMixin.S3ConfigurationProperty",
		reflect.TypeOf((*CfnScheduledQueryPropsMixin_S3ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnScheduledQueryPropsMixin.TagsItemsProperty",
		reflect.TypeOf((*CfnScheduledQueryPropsMixin_TagsItemsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnSubscriptionFilterMixinProps",
		reflect.TypeOf((*CfnSubscriptionFilterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnSubscriptionFilterPropsMixin",
		reflect.TypeOf((*CfnSubscriptionFilterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSubscriptionFilterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerMixinProps",
		reflect.TypeOf((*CfnTransformerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.AddKeyEntryProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_AddKeyEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.AddKeysProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_AddKeysProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.CopyValueEntryProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_CopyValueEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.CopyValueProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_CopyValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.CsvProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_CsvProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.DateTimeConverterProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_DateTimeConverterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.DeleteKeysProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_DeleteKeysProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.GrokProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_GrokProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.ListToMapProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_ListToMapProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.LowerCaseStringProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_LowerCaseStringProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.MoveKeyEntryProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_MoveKeyEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.MoveKeysProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_MoveKeysProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.ParseCloudfrontProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_ParseCloudfrontProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.ParseJSONProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_ParseJSONProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.ParseKeyValueProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_ParseKeyValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.ParsePostgresProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_ParsePostgresProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.ParseRoute53Property",
		reflect.TypeOf((*CfnTransformerPropsMixin_ParseRoute53Property)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.ParseToOCSFProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_ParseToOCSFProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.ParseVPCProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_ParseVPCProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.ParseWAFProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_ParseWAFProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.ProcessorProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_ProcessorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.RenameKeyEntryProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_RenameKeyEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.RenameKeysProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_RenameKeysProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.SplitStringEntryProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_SplitStringEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.SplitStringProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_SplitStringProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.SubstituteStringEntryProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_SubstituteStringEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.SubstituteStringProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_SubstituteStringProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.TrimStringProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_TrimStringProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.TypeConverterEntryProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_TypeConverterEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.TypeConverterProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_TypeConverterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin.UpperCaseStringProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_UpperCaseStringProperty)(nil)).Elem(),
	)
}
