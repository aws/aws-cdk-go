package previewawsdmsmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnCertificateMixinProps",
		reflect.TypeOf((*CfnCertificateMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnCertificatePropsMixin",
		reflect.TypeOf((*CfnCertificatePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCertificatePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnDataMigrationMixinProps",
		reflect.TypeOf((*CfnDataMigrationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnDataMigrationPropsMixin",
		reflect.TypeOf((*CfnDataMigrationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataMigrationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnDataMigrationPropsMixin.DataMigrationSettingsProperty",
		reflect.TypeOf((*CfnDataMigrationPropsMixin_DataMigrationSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnDataMigrationPropsMixin.SourceDataSettingsProperty",
		reflect.TypeOf((*CfnDataMigrationPropsMixin_SourceDataSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnDataProviderMixinProps",
		reflect.TypeOf((*CfnDataProviderMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnDataProviderPropsMixin",
		reflect.TypeOf((*CfnDataProviderPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataProviderPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnDataProviderPropsMixin.DocDbSettingsProperty",
		reflect.TypeOf((*CfnDataProviderPropsMixin_DocDbSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnDataProviderPropsMixin.IbmDb2LuwSettingsProperty",
		reflect.TypeOf((*CfnDataProviderPropsMixin_IbmDb2LuwSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnDataProviderPropsMixin.IbmDb2zOsSettingsProperty",
		reflect.TypeOf((*CfnDataProviderPropsMixin_IbmDb2zOsSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnDataProviderPropsMixin.MariaDbSettingsProperty",
		reflect.TypeOf((*CfnDataProviderPropsMixin_MariaDbSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnDataProviderPropsMixin.MicrosoftSqlServerSettingsProperty",
		reflect.TypeOf((*CfnDataProviderPropsMixin_MicrosoftSqlServerSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnDataProviderPropsMixin.MongoDbSettingsProperty",
		reflect.TypeOf((*CfnDataProviderPropsMixin_MongoDbSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnDataProviderPropsMixin.MySqlSettingsProperty",
		reflect.TypeOf((*CfnDataProviderPropsMixin_MySqlSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnDataProviderPropsMixin.OracleSettingsProperty",
		reflect.TypeOf((*CfnDataProviderPropsMixin_OracleSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnDataProviderPropsMixin.PostgreSqlSettingsProperty",
		reflect.TypeOf((*CfnDataProviderPropsMixin_PostgreSqlSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnDataProviderPropsMixin.RedshiftSettingsProperty",
		reflect.TypeOf((*CfnDataProviderPropsMixin_RedshiftSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnDataProviderPropsMixin.SettingsProperty",
		reflect.TypeOf((*CfnDataProviderPropsMixin_SettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnDataProviderPropsMixin.SybaseAseSettingsProperty",
		reflect.TypeOf((*CfnDataProviderPropsMixin_SybaseAseSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnEndpointMixinProps",
		reflect.TypeOf((*CfnEndpointMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnEndpointPropsMixin",
		reflect.TypeOf((*CfnEndpointPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEndpointPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnEndpointPropsMixin.DocDbSettingsProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_DocDbSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnEndpointPropsMixin.DynamoDbSettingsProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_DynamoDbSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnEndpointPropsMixin.ElasticsearchSettingsProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_ElasticsearchSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnEndpointPropsMixin.GcpMySQLSettingsProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_GcpMySQLSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnEndpointPropsMixin.IbmDb2SettingsProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_IbmDb2SettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnEndpointPropsMixin.KafkaSettingsProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_KafkaSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnEndpointPropsMixin.KinesisSettingsProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_KinesisSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnEndpointPropsMixin.MicrosoftSqlServerSettingsProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_MicrosoftSqlServerSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnEndpointPropsMixin.MongoDbSettingsProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_MongoDbSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnEndpointPropsMixin.MySqlSettingsProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_MySqlSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnEndpointPropsMixin.NeptuneSettingsProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_NeptuneSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnEndpointPropsMixin.OracleSettingsProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_OracleSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnEndpointPropsMixin.PostgreSqlSettingsProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_PostgreSqlSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnEndpointPropsMixin.RedisSettingsProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_RedisSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnEndpointPropsMixin.RedshiftSettingsProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_RedshiftSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnEndpointPropsMixin.S3SettingsProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_S3SettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnEndpointPropsMixin.SybaseSettingsProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_SybaseSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnEventSubscriptionMixinProps",
		reflect.TypeOf((*CfnEventSubscriptionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnEventSubscriptionPropsMixin",
		reflect.TypeOf((*CfnEventSubscriptionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEventSubscriptionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnInstanceProfileMixinProps",
		reflect.TypeOf((*CfnInstanceProfileMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnInstanceProfilePropsMixin",
		reflect.TypeOf((*CfnInstanceProfilePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInstanceProfilePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnMigrationProjectMixinProps",
		reflect.TypeOf((*CfnMigrationProjectMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnMigrationProjectPropsMixin",
		reflect.TypeOf((*CfnMigrationProjectPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMigrationProjectPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnMigrationProjectPropsMixin.DataProviderDescriptorProperty",
		reflect.TypeOf((*CfnMigrationProjectPropsMixin_DataProviderDescriptorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnMigrationProjectPropsMixin.SchemaConversionApplicationAttributesProperty",
		reflect.TypeOf((*CfnMigrationProjectPropsMixin_SchemaConversionApplicationAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnReplicationConfigMixinProps",
		reflect.TypeOf((*CfnReplicationConfigMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnReplicationConfigPropsMixin",
		reflect.TypeOf((*CfnReplicationConfigPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnReplicationConfigPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnReplicationConfigPropsMixin.ComputeConfigProperty",
		reflect.TypeOf((*CfnReplicationConfigPropsMixin_ComputeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnReplicationInstanceMixinProps",
		reflect.TypeOf((*CfnReplicationInstanceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnReplicationInstancePropsMixin",
		reflect.TypeOf((*CfnReplicationInstancePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnReplicationInstancePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnReplicationSubnetGroupMixinProps",
		reflect.TypeOf((*CfnReplicationSubnetGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnReplicationSubnetGroupPropsMixin",
		reflect.TypeOf((*CfnReplicationSubnetGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnReplicationSubnetGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnReplicationTaskMixinProps",
		reflect.TypeOf((*CfnReplicationTaskMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnReplicationTaskPropsMixin",
		reflect.TypeOf((*CfnReplicationTaskPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnReplicationTaskPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
}
