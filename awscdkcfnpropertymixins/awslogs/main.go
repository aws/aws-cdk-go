package awslogs

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnAccountPolicyMixinProps",
		reflect.TypeOf((*CfnAccountPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnAccountPolicyPropsMixin",
		reflect.TypeOf((*CfnAccountPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAccountPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnDeliveryDestinationMixinProps",
		reflect.TypeOf((*CfnDeliveryDestinationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnDeliveryDestinationPropsMixin",
		reflect.TypeOf((*CfnDeliveryDestinationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDeliveryDestinationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnDeliveryDestinationPropsMixin.DestinationPolicyProperty",
		reflect.TypeOf((*CfnDeliveryDestinationPropsMixin_DestinationPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnDeliveryMixinProps",
		reflect.TypeOf((*CfnDeliveryMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnDeliveryPropsMixin",
		reflect.TypeOf((*CfnDeliveryPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDeliveryPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnDeliverySourceMixinProps",
		reflect.TypeOf((*CfnDeliverySourceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnDeliverySourcePropsMixin",
		reflect.TypeOf((*CfnDeliverySourcePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDeliverySourcePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnDestinationMixinProps",
		reflect.TypeOf((*CfnDestinationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnDestinationPropsMixin",
		reflect.TypeOf((*CfnDestinationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDestinationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnIntegrationMixinProps",
		reflect.TypeOf((*CfnIntegrationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnIntegrationPropsMixin",
		reflect.TypeOf((*CfnIntegrationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIntegrationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnIntegrationPropsMixin.OpenSearchResourceConfigProperty",
		reflect.TypeOf((*CfnIntegrationPropsMixin_OpenSearchResourceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnIntegrationPropsMixin.ResourceConfigProperty",
		reflect.TypeOf((*CfnIntegrationPropsMixin_ResourceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnLogAnomalyDetectorMixinProps",
		reflect.TypeOf((*CfnLogAnomalyDetectorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnLogAnomalyDetectorPropsMixin",
		reflect.TypeOf((*CfnLogAnomalyDetectorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLogAnomalyDetectorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnLogGroupMixinProps",
		reflect.TypeOf((*CfnLogGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnLogGroupPropsMixin",
		reflect.TypeOf((*CfnLogGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLogGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnLogStreamMixinProps",
		reflect.TypeOf((*CfnLogStreamMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnLogStreamPropsMixin",
		reflect.TypeOf((*CfnLogStreamPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLogStreamPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnMetricFilterMixinProps",
		reflect.TypeOf((*CfnMetricFilterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnMetricFilterPropsMixin",
		reflect.TypeOf((*CfnMetricFilterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMetricFilterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnMetricFilterPropsMixin.DimensionProperty",
		reflect.TypeOf((*CfnMetricFilterPropsMixin_DimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnMetricFilterPropsMixin.MetricTransformationProperty",
		reflect.TypeOf((*CfnMetricFilterPropsMixin_MetricTransformationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnQueryDefinitionMixinProps",
		reflect.TypeOf((*CfnQueryDefinitionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnQueryDefinitionPropsMixin",
		reflect.TypeOf((*CfnQueryDefinitionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnQueryDefinitionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnResourcePolicyMixinProps",
		reflect.TypeOf((*CfnResourcePolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnResourcePolicyPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnScheduledQueryMixinProps",
		reflect.TypeOf((*CfnScheduledQueryMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnScheduledQueryPropsMixin",
		reflect.TypeOf((*CfnScheduledQueryPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnScheduledQueryPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnScheduledQueryPropsMixin.DestinationConfigurationProperty",
		reflect.TypeOf((*CfnScheduledQueryPropsMixin_DestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnScheduledQueryPropsMixin.S3ConfigurationProperty",
		reflect.TypeOf((*CfnScheduledQueryPropsMixin_S3ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnScheduledQueryPropsMixin.TagsItemsProperty",
		reflect.TypeOf((*CfnScheduledQueryPropsMixin_TagsItemsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnSubscriptionFilterMixinProps",
		reflect.TypeOf((*CfnSubscriptionFilterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnSubscriptionFilterPropsMixin",
		reflect.TypeOf((*CfnSubscriptionFilterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSubscriptionFilterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerMixinProps",
		reflect.TypeOf((*CfnTransformerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.AddKeyEntryProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_AddKeyEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.AddKeysProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_AddKeysProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.CopyValueEntryProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_CopyValueEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.CopyValueProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_CopyValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.CsvProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_CsvProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.DateTimeConverterProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_DateTimeConverterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.DeleteKeysProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_DeleteKeysProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.GrokProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_GrokProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.ListToMapProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_ListToMapProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.LowerCaseStringProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_LowerCaseStringProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.MoveKeyEntryProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_MoveKeyEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.MoveKeysProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_MoveKeysProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.ParseCloudfrontProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_ParseCloudfrontProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.ParseJSONProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_ParseJSONProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.ParseKeyValueProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_ParseKeyValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.ParsePostgresProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_ParsePostgresProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.ParseRoute53Property",
		reflect.TypeOf((*CfnTransformerPropsMixin_ParseRoute53Property)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.ParseToOCSFProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_ParseToOCSFProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.ParseVPCProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_ParseVPCProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.ParseWAFProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_ParseWAFProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.ProcessorProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_ProcessorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.RenameKeyEntryProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_RenameKeyEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.RenameKeysProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_RenameKeysProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.SplitStringEntryProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_SplitStringEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.SplitStringProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_SplitStringProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.SubstituteStringEntryProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_SubstituteStringEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.SubstituteStringProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_SubstituteStringProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.TrimStringProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_TrimStringProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.TypeConverterEntryProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_TypeConverterEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.TypeConverterProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_TypeConverterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnTransformerPropsMixin.UpperCaseStringProperty",
		reflect.TypeOf((*CfnTransformerPropsMixin_UpperCaseStringProperty)(nil)).Elem(),
	)
}
