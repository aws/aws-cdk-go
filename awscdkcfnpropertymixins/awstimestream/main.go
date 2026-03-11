package awstimestream

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnDatabaseMixinProps",
		reflect.TypeOf((*CfnDatabaseMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnDatabasePropsMixin",
		reflect.TypeOf((*CfnDatabasePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDatabasePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnInfluxDBClusterMixinProps",
		reflect.TypeOf((*CfnInfluxDBClusterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnInfluxDBClusterPropsMixin",
		reflect.TypeOf((*CfnInfluxDBClusterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInfluxDBClusterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnInfluxDBClusterPropsMixin.LogDeliveryConfigurationProperty",
		reflect.TypeOf((*CfnInfluxDBClusterPropsMixin_LogDeliveryConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnInfluxDBClusterPropsMixin.S3ConfigurationProperty",
		reflect.TypeOf((*CfnInfluxDBClusterPropsMixin_S3ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnInfluxDBInstanceMixinProps",
		reflect.TypeOf((*CfnInfluxDBInstanceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnInfluxDBInstancePropsMixin",
		reflect.TypeOf((*CfnInfluxDBInstancePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInfluxDBInstancePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnInfluxDBInstancePropsMixin.LogDeliveryConfigurationProperty",
		reflect.TypeOf((*CfnInfluxDBInstancePropsMixin_LogDeliveryConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnInfluxDBInstancePropsMixin.S3ConfigurationProperty",
		reflect.TypeOf((*CfnInfluxDBInstancePropsMixin_S3ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnScheduledQueryMixinProps",
		reflect.TypeOf((*CfnScheduledQueryMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnScheduledQueryPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnScheduledQueryPropsMixin.DimensionMappingProperty",
		reflect.TypeOf((*CfnScheduledQueryPropsMixin_DimensionMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnScheduledQueryPropsMixin.ErrorReportConfigurationProperty",
		reflect.TypeOf((*CfnScheduledQueryPropsMixin_ErrorReportConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnScheduledQueryPropsMixin.MixedMeasureMappingProperty",
		reflect.TypeOf((*CfnScheduledQueryPropsMixin_MixedMeasureMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnScheduledQueryPropsMixin.MultiMeasureAttributeMappingProperty",
		reflect.TypeOf((*CfnScheduledQueryPropsMixin_MultiMeasureAttributeMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnScheduledQueryPropsMixin.MultiMeasureMappingsProperty",
		reflect.TypeOf((*CfnScheduledQueryPropsMixin_MultiMeasureMappingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnScheduledQueryPropsMixin.NotificationConfigurationProperty",
		reflect.TypeOf((*CfnScheduledQueryPropsMixin_NotificationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnScheduledQueryPropsMixin.S3ConfigurationProperty",
		reflect.TypeOf((*CfnScheduledQueryPropsMixin_S3ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnScheduledQueryPropsMixin.ScheduleConfigurationProperty",
		reflect.TypeOf((*CfnScheduledQueryPropsMixin_ScheduleConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnScheduledQueryPropsMixin.SnsConfigurationProperty",
		reflect.TypeOf((*CfnScheduledQueryPropsMixin_SnsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnScheduledQueryPropsMixin.TargetConfigurationProperty",
		reflect.TypeOf((*CfnScheduledQueryPropsMixin_TargetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnScheduledQueryPropsMixin.TimestreamConfigurationProperty",
		reflect.TypeOf((*CfnScheduledQueryPropsMixin_TimestreamConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnTableMixinProps",
		reflect.TypeOf((*CfnTableMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnTablePropsMixin",
		reflect.TypeOf((*CfnTablePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTablePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnTablePropsMixin.MagneticStoreRejectedDataLocationProperty",
		reflect.TypeOf((*CfnTablePropsMixin_MagneticStoreRejectedDataLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnTablePropsMixin.MagneticStoreWritePropertiesProperty",
		reflect.TypeOf((*CfnTablePropsMixin_MagneticStoreWritePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnTablePropsMixin.PartitionKeyProperty",
		reflect.TypeOf((*CfnTablePropsMixin_PartitionKeyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnTablePropsMixin.RetentionPropertiesProperty",
		reflect.TypeOf((*CfnTablePropsMixin_RetentionPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnTablePropsMixin.S3ConfigurationProperty",
		reflect.TypeOf((*CfnTablePropsMixin_S3ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnTablePropsMixin.SchemaProperty",
		reflect.TypeOf((*CfnTablePropsMixin_SchemaProperty)(nil)).Elem(),
	)
}
