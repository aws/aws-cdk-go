package previewawsiotanalyticsmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnChannelMixinProps",
		reflect.TypeOf((*CfnChannelMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnChannelPropsMixin",
		reflect.TypeOf((*CfnChannelPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnChannelPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnChannelPropsMixin.ChannelStorageProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_ChannelStorageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnChannelPropsMixin.CustomerManagedS3Property",
		reflect.TypeOf((*CfnChannelPropsMixin_CustomerManagedS3Property)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnChannelPropsMixin.RetentionPeriodProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_RetentionPeriodProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatasetMixinProps",
		reflect.TypeOf((*CfnDatasetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatasetPropsMixin",
		reflect.TypeOf((*CfnDatasetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDatasetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatasetPropsMixin.ActionProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_ActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatasetPropsMixin.ContainerActionProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_ContainerActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatasetPropsMixin.DatasetContentDeliveryRuleDestinationProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_DatasetContentDeliveryRuleDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatasetPropsMixin.DatasetContentDeliveryRuleProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_DatasetContentDeliveryRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatasetPropsMixin.DatasetContentVersionValueProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_DatasetContentVersionValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatasetPropsMixin.DeltaTimeProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_DeltaTimeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatasetPropsMixin.DeltaTimeSessionWindowConfigurationProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_DeltaTimeSessionWindowConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatasetPropsMixin.FilterProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_FilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatasetPropsMixin.GlueConfigurationProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_GlueConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatasetPropsMixin.IotEventsDestinationConfigurationProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_IotEventsDestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatasetPropsMixin.LateDataRuleConfigurationProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_LateDataRuleConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatasetPropsMixin.LateDataRuleProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_LateDataRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatasetPropsMixin.OutputFileUriValueProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_OutputFileUriValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatasetPropsMixin.QueryActionProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_QueryActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatasetPropsMixin.ResourceConfigurationProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_ResourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatasetPropsMixin.RetentionPeriodProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_RetentionPeriodProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatasetPropsMixin.S3DestinationConfigurationProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_S3DestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatasetPropsMixin.ScheduleProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_ScheduleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatasetPropsMixin.TriggerProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_TriggerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatasetPropsMixin.TriggeringDatasetProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_TriggeringDatasetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatasetPropsMixin.VariableProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_VariableProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatasetPropsMixin.VersioningConfigurationProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_VersioningConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatastoreMixinProps",
		reflect.TypeOf((*CfnDatastoreMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatastorePropsMixin",
		reflect.TypeOf((*CfnDatastorePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDatastorePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatastorePropsMixin.ColumnProperty",
		reflect.TypeOf((*CfnDatastorePropsMixin_ColumnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatastorePropsMixin.CustomerManagedS3Property",
		reflect.TypeOf((*CfnDatastorePropsMixin_CustomerManagedS3Property)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatastorePropsMixin.CustomerManagedS3StorageProperty",
		reflect.TypeOf((*CfnDatastorePropsMixin_CustomerManagedS3StorageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatastorePropsMixin.DatastorePartitionProperty",
		reflect.TypeOf((*CfnDatastorePropsMixin_DatastorePartitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatastorePropsMixin.DatastorePartitionsProperty",
		reflect.TypeOf((*CfnDatastorePropsMixin_DatastorePartitionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatastorePropsMixin.DatastoreStorageProperty",
		reflect.TypeOf((*CfnDatastorePropsMixin_DatastoreStorageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatastorePropsMixin.FileFormatConfigurationProperty",
		reflect.TypeOf((*CfnDatastorePropsMixin_FileFormatConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatastorePropsMixin.IotSiteWiseMultiLayerStorageProperty",
		reflect.TypeOf((*CfnDatastorePropsMixin_IotSiteWiseMultiLayerStorageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatastorePropsMixin.ParquetConfigurationProperty",
		reflect.TypeOf((*CfnDatastorePropsMixin_ParquetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatastorePropsMixin.PartitionProperty",
		reflect.TypeOf((*CfnDatastorePropsMixin_PartitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatastorePropsMixin.RetentionPeriodProperty",
		reflect.TypeOf((*CfnDatastorePropsMixin_RetentionPeriodProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatastorePropsMixin.SchemaDefinitionProperty",
		reflect.TypeOf((*CfnDatastorePropsMixin_SchemaDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatastorePropsMixin.TimestampPartitionProperty",
		reflect.TypeOf((*CfnDatastorePropsMixin_TimestampPartitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnPipelineMixinProps",
		reflect.TypeOf((*CfnPipelineMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnPipelinePropsMixin",
		reflect.TypeOf((*CfnPipelinePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPipelinePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnPipelinePropsMixin.ActivityProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_ActivityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnPipelinePropsMixin.AddAttributesProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_AddAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnPipelinePropsMixin.ChannelProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_ChannelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnPipelinePropsMixin.DatastoreProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_DatastoreProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnPipelinePropsMixin.DeviceRegistryEnrichProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_DeviceRegistryEnrichProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnPipelinePropsMixin.DeviceShadowEnrichProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_DeviceShadowEnrichProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnPipelinePropsMixin.FilterProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_FilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnPipelinePropsMixin.LambdaProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_LambdaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnPipelinePropsMixin.MathProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_MathProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnPipelinePropsMixin.RemoveAttributesProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_RemoveAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnPipelinePropsMixin.SelectAttributesProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_SelectAttributesProperty)(nil)).Elem(),
	)
}
