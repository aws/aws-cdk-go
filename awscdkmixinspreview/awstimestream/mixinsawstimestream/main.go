package mixinsawstimestream

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_timestream.mixins.CfnDatabaseMixinProps",
		reflect.TypeOf((*CfnDatabaseMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_timestream.mixins.CfnDatabasePropsMixin",
		reflect.TypeOf((*CfnDatabasePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDatabasePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_timestream.mixins.CfnInfluxDBInstanceMixinProps",
		reflect.TypeOf((*CfnInfluxDBInstanceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_timestream.mixins.CfnInfluxDBInstancePropsMixin",
		reflect.TypeOf((*CfnInfluxDBInstancePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInfluxDBInstancePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_timestream.mixins.CfnInfluxDBInstancePropsMixin.LogDeliveryConfigurationProperty",
		reflect.TypeOf((*CfnInfluxDBInstancePropsMixin_LogDeliveryConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_timestream.mixins.CfnInfluxDBInstancePropsMixin.S3ConfigurationProperty",
		reflect.TypeOf((*CfnInfluxDBInstancePropsMixin_S3ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_timestream.mixins.CfnScheduledQueryMixinProps",
		reflect.TypeOf((*CfnScheduledQueryMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_timestream.mixins.CfnScheduledQueryPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_timestream.mixins.CfnScheduledQueryPropsMixin.DimensionMappingProperty",
		reflect.TypeOf((*CfnScheduledQueryPropsMixin_DimensionMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_timestream.mixins.CfnScheduledQueryPropsMixin.ErrorReportConfigurationProperty",
		reflect.TypeOf((*CfnScheduledQueryPropsMixin_ErrorReportConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_timestream.mixins.CfnScheduledQueryPropsMixin.MixedMeasureMappingProperty",
		reflect.TypeOf((*CfnScheduledQueryPropsMixin_MixedMeasureMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_timestream.mixins.CfnScheduledQueryPropsMixin.MultiMeasureAttributeMappingProperty",
		reflect.TypeOf((*CfnScheduledQueryPropsMixin_MultiMeasureAttributeMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_timestream.mixins.CfnScheduledQueryPropsMixin.MultiMeasureMappingsProperty",
		reflect.TypeOf((*CfnScheduledQueryPropsMixin_MultiMeasureMappingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_timestream.mixins.CfnScheduledQueryPropsMixin.NotificationConfigurationProperty",
		reflect.TypeOf((*CfnScheduledQueryPropsMixin_NotificationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_timestream.mixins.CfnScheduledQueryPropsMixin.S3ConfigurationProperty",
		reflect.TypeOf((*CfnScheduledQueryPropsMixin_S3ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_timestream.mixins.CfnScheduledQueryPropsMixin.ScheduleConfigurationProperty",
		reflect.TypeOf((*CfnScheduledQueryPropsMixin_ScheduleConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_timestream.mixins.CfnScheduledQueryPropsMixin.SnsConfigurationProperty",
		reflect.TypeOf((*CfnScheduledQueryPropsMixin_SnsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_timestream.mixins.CfnScheduledQueryPropsMixin.TargetConfigurationProperty",
		reflect.TypeOf((*CfnScheduledQueryPropsMixin_TargetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_timestream.mixins.CfnScheduledQueryPropsMixin.TimestreamConfigurationProperty",
		reflect.TypeOf((*CfnScheduledQueryPropsMixin_TimestreamConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_timestream.mixins.CfnTableMixinProps",
		reflect.TypeOf((*CfnTableMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_timestream.mixins.CfnTablePropsMixin",
		reflect.TypeOf((*CfnTablePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTablePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_timestream.mixins.CfnTablePropsMixin.MagneticStoreRejectedDataLocationProperty",
		reflect.TypeOf((*CfnTablePropsMixin_MagneticStoreRejectedDataLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_timestream.mixins.CfnTablePropsMixin.MagneticStoreWritePropertiesProperty",
		reflect.TypeOf((*CfnTablePropsMixin_MagneticStoreWritePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_timestream.mixins.CfnTablePropsMixin.PartitionKeyProperty",
		reflect.TypeOf((*CfnTablePropsMixin_PartitionKeyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_timestream.mixins.CfnTablePropsMixin.RetentionPropertiesProperty",
		reflect.TypeOf((*CfnTablePropsMixin_RetentionPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_timestream.mixins.CfnTablePropsMixin.S3ConfigurationProperty",
		reflect.TypeOf((*CfnTablePropsMixin_S3ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_timestream.mixins.CfnTablePropsMixin.SchemaProperty",
		reflect.TypeOf((*CfnTablePropsMixin_SchemaProperty)(nil)).Elem(),
	)
}
