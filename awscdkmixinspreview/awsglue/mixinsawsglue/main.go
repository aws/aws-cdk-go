package mixinsawsglue

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnClassifierMixinProps",
		reflect.TypeOf((*CfnClassifierMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnClassifierPropsMixin",
		reflect.TypeOf((*CfnClassifierPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnClassifierPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnClassifierPropsMixin.CsvClassifierProperty",
		reflect.TypeOf((*CfnClassifierPropsMixin_CsvClassifierProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnClassifierPropsMixin.GrokClassifierProperty",
		reflect.TypeOf((*CfnClassifierPropsMixin_GrokClassifierProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnClassifierPropsMixin.JsonClassifierProperty",
		reflect.TypeOf((*CfnClassifierPropsMixin_JsonClassifierProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnClassifierPropsMixin.XMLClassifierProperty",
		reflect.TypeOf((*CfnClassifierPropsMixin_XMLClassifierProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnConnectionMixinProps",
		reflect.TypeOf((*CfnConnectionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnConnectionPropsMixin",
		reflect.TypeOf((*CfnConnectionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConnectionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnConnectionPropsMixin.AuthenticationConfigurationInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_AuthenticationConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnConnectionPropsMixin.AuthorizationCodePropertiesProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_AuthorizationCodePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnConnectionPropsMixin.BasicAuthenticationCredentialsProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_BasicAuthenticationCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnConnectionPropsMixin.ConnectionInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_ConnectionInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnConnectionPropsMixin.OAuth2ClientApplicationProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_OAuth2ClientApplicationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnConnectionPropsMixin.OAuth2CredentialsProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_OAuth2CredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnConnectionPropsMixin.OAuth2PropertiesInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_OAuth2PropertiesInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnConnectionPropsMixin.PhysicalConnectionRequirementsProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_PhysicalConnectionRequirementsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnCrawlerMixinProps",
		reflect.TypeOf((*CfnCrawlerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnCrawlerPropsMixin",
		reflect.TypeOf((*CfnCrawlerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCrawlerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnCrawlerPropsMixin.CatalogTargetProperty",
		reflect.TypeOf((*CfnCrawlerPropsMixin_CatalogTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnCrawlerPropsMixin.DeltaTargetProperty",
		reflect.TypeOf((*CfnCrawlerPropsMixin_DeltaTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnCrawlerPropsMixin.DynamoDBTargetProperty",
		reflect.TypeOf((*CfnCrawlerPropsMixin_DynamoDBTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnCrawlerPropsMixin.HudiTargetProperty",
		reflect.TypeOf((*CfnCrawlerPropsMixin_HudiTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnCrawlerPropsMixin.IcebergTargetProperty",
		reflect.TypeOf((*CfnCrawlerPropsMixin_IcebergTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnCrawlerPropsMixin.JdbcTargetProperty",
		reflect.TypeOf((*CfnCrawlerPropsMixin_JdbcTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnCrawlerPropsMixin.LakeFormationConfigurationProperty",
		reflect.TypeOf((*CfnCrawlerPropsMixin_LakeFormationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnCrawlerPropsMixin.MongoDBTargetProperty",
		reflect.TypeOf((*CfnCrawlerPropsMixin_MongoDBTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnCrawlerPropsMixin.RecrawlPolicyProperty",
		reflect.TypeOf((*CfnCrawlerPropsMixin_RecrawlPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnCrawlerPropsMixin.S3TargetProperty",
		reflect.TypeOf((*CfnCrawlerPropsMixin_S3TargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnCrawlerPropsMixin.ScheduleProperty",
		reflect.TypeOf((*CfnCrawlerPropsMixin_ScheduleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnCrawlerPropsMixin.SchemaChangePolicyProperty",
		reflect.TypeOf((*CfnCrawlerPropsMixin_SchemaChangePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnCrawlerPropsMixin.TargetsProperty",
		reflect.TypeOf((*CfnCrawlerPropsMixin_TargetsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnCustomEntityTypeMixinProps",
		reflect.TypeOf((*CfnCustomEntityTypeMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnCustomEntityTypePropsMixin",
		reflect.TypeOf((*CfnCustomEntityTypePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCustomEntityTypePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnDataCatalogEncryptionSettingsMixinProps",
		reflect.TypeOf((*CfnDataCatalogEncryptionSettingsMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnDataCatalogEncryptionSettingsPropsMixin",
		reflect.TypeOf((*CfnDataCatalogEncryptionSettingsPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataCatalogEncryptionSettingsPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnDataCatalogEncryptionSettingsPropsMixin.ConnectionPasswordEncryptionProperty",
		reflect.TypeOf((*CfnDataCatalogEncryptionSettingsPropsMixin_ConnectionPasswordEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnDataCatalogEncryptionSettingsPropsMixin.DataCatalogEncryptionSettingsProperty",
		reflect.TypeOf((*CfnDataCatalogEncryptionSettingsPropsMixin_DataCatalogEncryptionSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnDataCatalogEncryptionSettingsPropsMixin.EncryptionAtRestProperty",
		reflect.TypeOf((*CfnDataCatalogEncryptionSettingsPropsMixin_EncryptionAtRestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnDataQualityRulesetMixinProps",
		reflect.TypeOf((*CfnDataQualityRulesetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnDataQualityRulesetPropsMixin",
		reflect.TypeOf((*CfnDataQualityRulesetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataQualityRulesetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnDataQualityRulesetPropsMixin.DataQualityTargetTableProperty",
		reflect.TypeOf((*CfnDataQualityRulesetPropsMixin_DataQualityTargetTableProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnDatabaseMixinProps",
		reflect.TypeOf((*CfnDatabaseMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnDatabasePropsMixin",
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
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnDatabasePropsMixin.DataLakePrincipalProperty",
		reflect.TypeOf((*CfnDatabasePropsMixin_DataLakePrincipalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnDatabasePropsMixin.DatabaseIdentifierProperty",
		reflect.TypeOf((*CfnDatabasePropsMixin_DatabaseIdentifierProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnDatabasePropsMixin.DatabaseInputProperty",
		reflect.TypeOf((*CfnDatabasePropsMixin_DatabaseInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnDatabasePropsMixin.FederatedDatabaseProperty",
		reflect.TypeOf((*CfnDatabasePropsMixin_FederatedDatabaseProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnDatabasePropsMixin.PrincipalPrivilegesProperty",
		reflect.TypeOf((*CfnDatabasePropsMixin_PrincipalPrivilegesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnDevEndpointMixinProps",
		reflect.TypeOf((*CfnDevEndpointMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnDevEndpointPropsMixin",
		reflect.TypeOf((*CfnDevEndpointPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDevEndpointPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnIdentityCenterConfigurationMixinProps",
		reflect.TypeOf((*CfnIdentityCenterConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnIdentityCenterConfigurationPropsMixin",
		reflect.TypeOf((*CfnIdentityCenterConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIdentityCenterConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnIntegrationMixinProps",
		reflect.TypeOf((*CfnIntegrationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnIntegrationPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnIntegrationPropsMixin.IntegrationConfigProperty",
		reflect.TypeOf((*CfnIntegrationPropsMixin_IntegrationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnIntegrationResourcePropertyMixinProps",
		reflect.TypeOf((*CfnIntegrationResourcePropertyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnIntegrationResourcePropertyPropsMixin",
		reflect.TypeOf((*CfnIntegrationResourcePropertyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIntegrationResourcePropertyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnIntegrationResourcePropertyPropsMixin.SourceProcessingPropertiesProperty",
		reflect.TypeOf((*CfnIntegrationResourcePropertyPropsMixin_SourceProcessingPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnIntegrationResourcePropertyPropsMixin.TargetProcessingPropertiesProperty",
		reflect.TypeOf((*CfnIntegrationResourcePropertyPropsMixin_TargetProcessingPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnJobMixinProps",
		reflect.TypeOf((*CfnJobMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnJobPropsMixin",
		reflect.TypeOf((*CfnJobPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnJobPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnJobPropsMixin.ConnectionsListProperty",
		reflect.TypeOf((*CfnJobPropsMixin_ConnectionsListProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnJobPropsMixin.ExecutionPropertyProperty",
		reflect.TypeOf((*CfnJobPropsMixin_ExecutionPropertyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnJobPropsMixin.JobCommandProperty",
		reflect.TypeOf((*CfnJobPropsMixin_JobCommandProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnJobPropsMixin.NotificationPropertyProperty",
		reflect.TypeOf((*CfnJobPropsMixin_NotificationPropertyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnMLTransformMixinProps",
		reflect.TypeOf((*CfnMLTransformMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnMLTransformPropsMixin",
		reflect.TypeOf((*CfnMLTransformPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMLTransformPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnMLTransformPropsMixin.FindMatchesParametersProperty",
		reflect.TypeOf((*CfnMLTransformPropsMixin_FindMatchesParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnMLTransformPropsMixin.GlueTablesProperty",
		reflect.TypeOf((*CfnMLTransformPropsMixin_GlueTablesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnMLTransformPropsMixin.InputRecordTablesProperty",
		reflect.TypeOf((*CfnMLTransformPropsMixin_InputRecordTablesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnMLTransformPropsMixin.MLUserDataEncryptionProperty",
		reflect.TypeOf((*CfnMLTransformPropsMixin_MLUserDataEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnMLTransformPropsMixin.TransformEncryptionProperty",
		reflect.TypeOf((*CfnMLTransformPropsMixin_TransformEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnMLTransformPropsMixin.TransformParametersProperty",
		reflect.TypeOf((*CfnMLTransformPropsMixin_TransformParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnPartitionMixinProps",
		reflect.TypeOf((*CfnPartitionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnPartitionPropsMixin",
		reflect.TypeOf((*CfnPartitionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPartitionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnPartitionPropsMixin.ColumnProperty",
		reflect.TypeOf((*CfnPartitionPropsMixin_ColumnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnPartitionPropsMixin.OrderProperty",
		reflect.TypeOf((*CfnPartitionPropsMixin_OrderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnPartitionPropsMixin.PartitionInputProperty",
		reflect.TypeOf((*CfnPartitionPropsMixin_PartitionInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnPartitionPropsMixin.SchemaIdProperty",
		reflect.TypeOf((*CfnPartitionPropsMixin_SchemaIdProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnPartitionPropsMixin.SchemaReferenceProperty",
		reflect.TypeOf((*CfnPartitionPropsMixin_SchemaReferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnPartitionPropsMixin.SerdeInfoProperty",
		reflect.TypeOf((*CfnPartitionPropsMixin_SerdeInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnPartitionPropsMixin.SkewedInfoProperty",
		reflect.TypeOf((*CfnPartitionPropsMixin_SkewedInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnPartitionPropsMixin.StorageDescriptorProperty",
		reflect.TypeOf((*CfnPartitionPropsMixin_StorageDescriptorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnRegistryMixinProps",
		reflect.TypeOf((*CfnRegistryMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnRegistryPropsMixin",
		reflect.TypeOf((*CfnRegistryPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRegistryPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnSchemaMixinProps",
		reflect.TypeOf((*CfnSchemaMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnSchemaPropsMixin",
		reflect.TypeOf((*CfnSchemaPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSchemaPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnSchemaPropsMixin.RegistryProperty",
		reflect.TypeOf((*CfnSchemaPropsMixin_RegistryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnSchemaPropsMixin.SchemaVersionProperty",
		reflect.TypeOf((*CfnSchemaPropsMixin_SchemaVersionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnSchemaVersionMetadataMixinProps",
		reflect.TypeOf((*CfnSchemaVersionMetadataMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnSchemaVersionMetadataPropsMixin",
		reflect.TypeOf((*CfnSchemaVersionMetadataPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSchemaVersionMetadataPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnSchemaVersionMixinProps",
		reflect.TypeOf((*CfnSchemaVersionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnSchemaVersionPropsMixin",
		reflect.TypeOf((*CfnSchemaVersionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSchemaVersionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnSchemaVersionPropsMixin.SchemaProperty",
		reflect.TypeOf((*CfnSchemaVersionPropsMixin_SchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnSecurityConfigurationMixinProps",
		reflect.TypeOf((*CfnSecurityConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnSecurityConfigurationPropsMixin",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSecurityConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnSecurityConfigurationPropsMixin.CloudWatchEncryptionProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_CloudWatchEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnSecurityConfigurationPropsMixin.EncryptionConfigurationProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_EncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnSecurityConfigurationPropsMixin.JobBookmarksEncryptionProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_JobBookmarksEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnSecurityConfigurationPropsMixin.S3EncryptionProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_S3EncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTableMixinProps",
		reflect.TypeOf((*CfnTableMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTableOptimizerMixinProps",
		reflect.TypeOf((*CfnTableOptimizerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTableOptimizerPropsMixin",
		reflect.TypeOf((*CfnTableOptimizerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTableOptimizerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTableOptimizerPropsMixin.IcebergConfigurationProperty",
		reflect.TypeOf((*CfnTableOptimizerPropsMixin_IcebergConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTableOptimizerPropsMixin.IcebergRetentionConfigurationProperty",
		reflect.TypeOf((*CfnTableOptimizerPropsMixin_IcebergRetentionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTableOptimizerPropsMixin.OrphanFileDeletionConfigurationProperty",
		reflect.TypeOf((*CfnTableOptimizerPropsMixin_OrphanFileDeletionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTableOptimizerPropsMixin.RetentionConfigurationProperty",
		reflect.TypeOf((*CfnTableOptimizerPropsMixin_RetentionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTableOptimizerPropsMixin.TableOptimizerConfigurationProperty",
		reflect.TypeOf((*CfnTableOptimizerPropsMixin_TableOptimizerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTableOptimizerPropsMixin.VpcConfigurationProperty",
		reflect.TypeOf((*CfnTableOptimizerPropsMixin_VpcConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTablePropsMixin",
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
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTablePropsMixin.ColumnProperty",
		reflect.TypeOf((*CfnTablePropsMixin_ColumnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTablePropsMixin.IcebergInputProperty",
		reflect.TypeOf((*CfnTablePropsMixin_IcebergInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTablePropsMixin.OpenTableFormatInputProperty",
		reflect.TypeOf((*CfnTablePropsMixin_OpenTableFormatInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTablePropsMixin.OrderProperty",
		reflect.TypeOf((*CfnTablePropsMixin_OrderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTablePropsMixin.SchemaIdProperty",
		reflect.TypeOf((*CfnTablePropsMixin_SchemaIdProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTablePropsMixin.SchemaReferenceProperty",
		reflect.TypeOf((*CfnTablePropsMixin_SchemaReferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTablePropsMixin.SerdeInfoProperty",
		reflect.TypeOf((*CfnTablePropsMixin_SerdeInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTablePropsMixin.SkewedInfoProperty",
		reflect.TypeOf((*CfnTablePropsMixin_SkewedInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTablePropsMixin.StorageDescriptorProperty",
		reflect.TypeOf((*CfnTablePropsMixin_StorageDescriptorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTablePropsMixin.TableIdentifierProperty",
		reflect.TypeOf((*CfnTablePropsMixin_TableIdentifierProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTablePropsMixin.TableInputProperty",
		reflect.TypeOf((*CfnTablePropsMixin_TableInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTriggerMixinProps",
		reflect.TypeOf((*CfnTriggerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTriggerPropsMixin",
		reflect.TypeOf((*CfnTriggerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTriggerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTriggerPropsMixin.ActionProperty",
		reflect.TypeOf((*CfnTriggerPropsMixin_ActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTriggerPropsMixin.ConditionProperty",
		reflect.TypeOf((*CfnTriggerPropsMixin_ConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTriggerPropsMixin.EventBatchingConditionProperty",
		reflect.TypeOf((*CfnTriggerPropsMixin_EventBatchingConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTriggerPropsMixin.NotificationPropertyProperty",
		reflect.TypeOf((*CfnTriggerPropsMixin_NotificationPropertyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTriggerPropsMixin.PredicateProperty",
		reflect.TypeOf((*CfnTriggerPropsMixin_PredicateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnUsageProfileMixinProps",
		reflect.TypeOf((*CfnUsageProfileMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnUsageProfilePropsMixin",
		reflect.TypeOf((*CfnUsageProfilePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUsageProfilePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnUsageProfilePropsMixin.ConfigurationObjectProperty",
		reflect.TypeOf((*CfnUsageProfilePropsMixin_ConfigurationObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnUsageProfilePropsMixin.ProfileConfigurationProperty",
		reflect.TypeOf((*CfnUsageProfilePropsMixin_ProfileConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnWorkflowMixinProps",
		reflect.TypeOf((*CfnWorkflowMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnWorkflowPropsMixin",
		reflect.TypeOf((*CfnWorkflowPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWorkflowPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
}
