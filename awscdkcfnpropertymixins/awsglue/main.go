package awsglue

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCatalogMixinProps",
		reflect.TypeOf((*CfnCatalogMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCatalogPropsMixin",
		reflect.TypeOf((*CfnCatalogPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCatalogPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCatalogPropsMixin.CatalogPropertiesProperty",
		reflect.TypeOf((*CfnCatalogPropsMixin_CatalogPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCatalogPropsMixin.DataLakeAccessPropertiesProperty",
		reflect.TypeOf((*CfnCatalogPropsMixin_DataLakeAccessPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCatalogPropsMixin.DataLakePrincipalProperty",
		reflect.TypeOf((*CfnCatalogPropsMixin_DataLakePrincipalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCatalogPropsMixin.FederatedCatalogProperty",
		reflect.TypeOf((*CfnCatalogPropsMixin_FederatedCatalogProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCatalogPropsMixin.PrincipalPermissionsProperty",
		reflect.TypeOf((*CfnCatalogPropsMixin_PrincipalPermissionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCatalogPropsMixin.TargetRedshiftCatalogProperty",
		reflect.TypeOf((*CfnCatalogPropsMixin_TargetRedshiftCatalogProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnClassifierMixinProps",
		reflect.TypeOf((*CfnClassifierMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnClassifierPropsMixin",
		reflect.TypeOf((*CfnClassifierPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnClassifierPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnClassifierPropsMixin.CsvClassifierProperty",
		reflect.TypeOf((*CfnClassifierPropsMixin_CsvClassifierProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnClassifierPropsMixin.GrokClassifierProperty",
		reflect.TypeOf((*CfnClassifierPropsMixin_GrokClassifierProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnClassifierPropsMixin.JsonClassifierProperty",
		reflect.TypeOf((*CfnClassifierPropsMixin_JsonClassifierProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnClassifierPropsMixin.XMLClassifierProperty",
		reflect.TypeOf((*CfnClassifierPropsMixin_XMLClassifierProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnConnectionMixinProps",
		reflect.TypeOf((*CfnConnectionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnConnectionPropsMixin",
		reflect.TypeOf((*CfnConnectionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConnectionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnConnectionPropsMixin.AuthenticationConfigurationInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_AuthenticationConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnConnectionPropsMixin.AuthorizationCodePropertiesProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_AuthorizationCodePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnConnectionPropsMixin.BasicAuthenticationCredentialsProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_BasicAuthenticationCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnConnectionPropsMixin.ConnectionInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_ConnectionInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnConnectionPropsMixin.OAuth2ClientApplicationProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_OAuth2ClientApplicationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnConnectionPropsMixin.OAuth2CredentialsProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_OAuth2CredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnConnectionPropsMixin.OAuth2PropertiesInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_OAuth2PropertiesInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnConnectionPropsMixin.PhysicalConnectionRequirementsProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_PhysicalConnectionRequirementsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCrawlerMixinProps",
		reflect.TypeOf((*CfnCrawlerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCrawlerPropsMixin",
		reflect.TypeOf((*CfnCrawlerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCrawlerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCrawlerPropsMixin.CatalogTargetProperty",
		reflect.TypeOf((*CfnCrawlerPropsMixin_CatalogTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCrawlerPropsMixin.DeltaTargetProperty",
		reflect.TypeOf((*CfnCrawlerPropsMixin_DeltaTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCrawlerPropsMixin.DynamoDBTargetProperty",
		reflect.TypeOf((*CfnCrawlerPropsMixin_DynamoDBTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCrawlerPropsMixin.HudiTargetProperty",
		reflect.TypeOf((*CfnCrawlerPropsMixin_HudiTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCrawlerPropsMixin.IcebergTargetProperty",
		reflect.TypeOf((*CfnCrawlerPropsMixin_IcebergTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCrawlerPropsMixin.JdbcTargetProperty",
		reflect.TypeOf((*CfnCrawlerPropsMixin_JdbcTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCrawlerPropsMixin.LakeFormationConfigurationProperty",
		reflect.TypeOf((*CfnCrawlerPropsMixin_LakeFormationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCrawlerPropsMixin.MongoDBTargetProperty",
		reflect.TypeOf((*CfnCrawlerPropsMixin_MongoDBTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCrawlerPropsMixin.RecrawlPolicyProperty",
		reflect.TypeOf((*CfnCrawlerPropsMixin_RecrawlPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCrawlerPropsMixin.S3TargetProperty",
		reflect.TypeOf((*CfnCrawlerPropsMixin_S3TargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCrawlerPropsMixin.ScheduleProperty",
		reflect.TypeOf((*CfnCrawlerPropsMixin_ScheduleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCrawlerPropsMixin.SchemaChangePolicyProperty",
		reflect.TypeOf((*CfnCrawlerPropsMixin_SchemaChangePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCrawlerPropsMixin.TargetsProperty",
		reflect.TypeOf((*CfnCrawlerPropsMixin_TargetsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCustomEntityTypeMixinProps",
		reflect.TypeOf((*CfnCustomEntityTypeMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCustomEntityTypePropsMixin",
		reflect.TypeOf((*CfnCustomEntityTypePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCustomEntityTypePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnDataCatalogEncryptionSettingsMixinProps",
		reflect.TypeOf((*CfnDataCatalogEncryptionSettingsMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnDataCatalogEncryptionSettingsPropsMixin",
		reflect.TypeOf((*CfnDataCatalogEncryptionSettingsPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataCatalogEncryptionSettingsPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnDataCatalogEncryptionSettingsPropsMixin.ConnectionPasswordEncryptionProperty",
		reflect.TypeOf((*CfnDataCatalogEncryptionSettingsPropsMixin_ConnectionPasswordEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnDataCatalogEncryptionSettingsPropsMixin.DataCatalogEncryptionSettingsProperty",
		reflect.TypeOf((*CfnDataCatalogEncryptionSettingsPropsMixin_DataCatalogEncryptionSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnDataCatalogEncryptionSettingsPropsMixin.EncryptionAtRestProperty",
		reflect.TypeOf((*CfnDataCatalogEncryptionSettingsPropsMixin_EncryptionAtRestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnDataQualityRulesetMixinProps",
		reflect.TypeOf((*CfnDataQualityRulesetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnDataQualityRulesetPropsMixin",
		reflect.TypeOf((*CfnDataQualityRulesetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataQualityRulesetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnDataQualityRulesetPropsMixin.DataQualityTargetTableProperty",
		reflect.TypeOf((*CfnDataQualityRulesetPropsMixin_DataQualityTargetTableProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnDatabaseMixinProps",
		reflect.TypeOf((*CfnDatabaseMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnDatabasePropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnDatabasePropsMixin.DataLakePrincipalProperty",
		reflect.TypeOf((*CfnDatabasePropsMixin_DataLakePrincipalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnDatabasePropsMixin.DatabaseIdentifierProperty",
		reflect.TypeOf((*CfnDatabasePropsMixin_DatabaseIdentifierProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnDatabasePropsMixin.DatabaseInputProperty",
		reflect.TypeOf((*CfnDatabasePropsMixin_DatabaseInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnDatabasePropsMixin.FederatedDatabaseProperty",
		reflect.TypeOf((*CfnDatabasePropsMixin_FederatedDatabaseProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnDatabasePropsMixin.PrincipalPrivilegesProperty",
		reflect.TypeOf((*CfnDatabasePropsMixin_PrincipalPrivilegesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnDevEndpointMixinProps",
		reflect.TypeOf((*CfnDevEndpointMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnDevEndpointPropsMixin",
		reflect.TypeOf((*CfnDevEndpointPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDevEndpointPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnIdentityCenterConfigurationMixinProps",
		reflect.TypeOf((*CfnIdentityCenterConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnIdentityCenterConfigurationPropsMixin",
		reflect.TypeOf((*CfnIdentityCenterConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIdentityCenterConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnIntegrationMixinProps",
		reflect.TypeOf((*CfnIntegrationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnIntegrationPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnIntegrationPropsMixin.IntegrationConfigProperty",
		reflect.TypeOf((*CfnIntegrationPropsMixin_IntegrationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnIntegrationResourcePropertyMixinProps",
		reflect.TypeOf((*CfnIntegrationResourcePropertyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnIntegrationResourcePropertyPropsMixin",
		reflect.TypeOf((*CfnIntegrationResourcePropertyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIntegrationResourcePropertyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnIntegrationResourcePropertyPropsMixin.SourceProcessingPropertiesProperty",
		reflect.TypeOf((*CfnIntegrationResourcePropertyPropsMixin_SourceProcessingPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnIntegrationResourcePropertyPropsMixin.TargetProcessingPropertiesProperty",
		reflect.TypeOf((*CfnIntegrationResourcePropertyPropsMixin_TargetProcessingPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnJobMixinProps",
		reflect.TypeOf((*CfnJobMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnJobPropsMixin",
		reflect.TypeOf((*CfnJobPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnJobPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnJobPropsMixin.ConnectionsListProperty",
		reflect.TypeOf((*CfnJobPropsMixin_ConnectionsListProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnJobPropsMixin.ExecutionPropertyProperty",
		reflect.TypeOf((*CfnJobPropsMixin_ExecutionPropertyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnJobPropsMixin.JobCommandProperty",
		reflect.TypeOf((*CfnJobPropsMixin_JobCommandProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnJobPropsMixin.NotificationPropertyProperty",
		reflect.TypeOf((*CfnJobPropsMixin_NotificationPropertyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnMLTransformMixinProps",
		reflect.TypeOf((*CfnMLTransformMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnMLTransformPropsMixin",
		reflect.TypeOf((*CfnMLTransformPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMLTransformPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnMLTransformPropsMixin.FindMatchesParametersProperty",
		reflect.TypeOf((*CfnMLTransformPropsMixin_FindMatchesParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnMLTransformPropsMixin.GlueTablesProperty",
		reflect.TypeOf((*CfnMLTransformPropsMixin_GlueTablesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnMLTransformPropsMixin.InputRecordTablesProperty",
		reflect.TypeOf((*CfnMLTransformPropsMixin_InputRecordTablesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnMLTransformPropsMixin.MLUserDataEncryptionProperty",
		reflect.TypeOf((*CfnMLTransformPropsMixin_MLUserDataEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnMLTransformPropsMixin.TransformEncryptionProperty",
		reflect.TypeOf((*CfnMLTransformPropsMixin_TransformEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnMLTransformPropsMixin.TransformParametersProperty",
		reflect.TypeOf((*CfnMLTransformPropsMixin_TransformParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnPartitionMixinProps",
		reflect.TypeOf((*CfnPartitionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnPartitionPropsMixin",
		reflect.TypeOf((*CfnPartitionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPartitionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnPartitionPropsMixin.ColumnProperty",
		reflect.TypeOf((*CfnPartitionPropsMixin_ColumnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnPartitionPropsMixin.OrderProperty",
		reflect.TypeOf((*CfnPartitionPropsMixin_OrderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnPartitionPropsMixin.PartitionInputProperty",
		reflect.TypeOf((*CfnPartitionPropsMixin_PartitionInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnPartitionPropsMixin.SchemaIdProperty",
		reflect.TypeOf((*CfnPartitionPropsMixin_SchemaIdProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnPartitionPropsMixin.SchemaReferenceProperty",
		reflect.TypeOf((*CfnPartitionPropsMixin_SchemaReferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnPartitionPropsMixin.SerdeInfoProperty",
		reflect.TypeOf((*CfnPartitionPropsMixin_SerdeInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnPartitionPropsMixin.SkewedInfoProperty",
		reflect.TypeOf((*CfnPartitionPropsMixin_SkewedInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnPartitionPropsMixin.StorageDescriptorProperty",
		reflect.TypeOf((*CfnPartitionPropsMixin_StorageDescriptorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnRegistryMixinProps",
		reflect.TypeOf((*CfnRegistryMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnRegistryPropsMixin",
		reflect.TypeOf((*CfnRegistryPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRegistryPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnSchemaMixinProps",
		reflect.TypeOf((*CfnSchemaMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnSchemaPropsMixin",
		reflect.TypeOf((*CfnSchemaPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSchemaPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnSchemaPropsMixin.RegistryProperty",
		reflect.TypeOf((*CfnSchemaPropsMixin_RegistryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnSchemaPropsMixin.SchemaVersionProperty",
		reflect.TypeOf((*CfnSchemaPropsMixin_SchemaVersionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnSchemaVersionMetadataMixinProps",
		reflect.TypeOf((*CfnSchemaVersionMetadataMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnSchemaVersionMetadataPropsMixin",
		reflect.TypeOf((*CfnSchemaVersionMetadataPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSchemaVersionMetadataPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnSchemaVersionMixinProps",
		reflect.TypeOf((*CfnSchemaVersionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnSchemaVersionPropsMixin",
		reflect.TypeOf((*CfnSchemaVersionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSchemaVersionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnSchemaVersionPropsMixin.SchemaProperty",
		reflect.TypeOf((*CfnSchemaVersionPropsMixin_SchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnSecurityConfigurationMixinProps",
		reflect.TypeOf((*CfnSecurityConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnSecurityConfigurationPropsMixin",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSecurityConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnSecurityConfigurationPropsMixin.CloudWatchEncryptionProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_CloudWatchEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnSecurityConfigurationPropsMixin.EncryptionConfigurationProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_EncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnSecurityConfigurationPropsMixin.JobBookmarksEncryptionProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_JobBookmarksEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnSecurityConfigurationPropsMixin.S3EncryptionProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_S3EncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTableMixinProps",
		reflect.TypeOf((*CfnTableMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTableOptimizerMixinProps",
		reflect.TypeOf((*CfnTableOptimizerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTableOptimizerPropsMixin",
		reflect.TypeOf((*CfnTableOptimizerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTableOptimizerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTableOptimizerPropsMixin.IcebergConfigurationProperty",
		reflect.TypeOf((*CfnTableOptimizerPropsMixin_IcebergConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTableOptimizerPropsMixin.IcebergRetentionConfigurationProperty",
		reflect.TypeOf((*CfnTableOptimizerPropsMixin_IcebergRetentionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTableOptimizerPropsMixin.OrphanFileDeletionConfigurationProperty",
		reflect.TypeOf((*CfnTableOptimizerPropsMixin_OrphanFileDeletionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTableOptimizerPropsMixin.RetentionConfigurationProperty",
		reflect.TypeOf((*CfnTableOptimizerPropsMixin_RetentionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTableOptimizerPropsMixin.TableOptimizerConfigurationProperty",
		reflect.TypeOf((*CfnTableOptimizerPropsMixin_TableOptimizerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTableOptimizerPropsMixin.VpcConfigurationProperty",
		reflect.TypeOf((*CfnTableOptimizerPropsMixin_VpcConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTablePropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTablePropsMixin.ColumnProperty",
		reflect.TypeOf((*CfnTablePropsMixin_ColumnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTablePropsMixin.IcebergInputProperty",
		reflect.TypeOf((*CfnTablePropsMixin_IcebergInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTablePropsMixin.IcebergPartitionFieldProperty",
		reflect.TypeOf((*CfnTablePropsMixin_IcebergPartitionFieldProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTablePropsMixin.IcebergPartitionSpecProperty",
		reflect.TypeOf((*CfnTablePropsMixin_IcebergPartitionSpecProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTablePropsMixin.IcebergSchemaProperty",
		reflect.TypeOf((*CfnTablePropsMixin_IcebergSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTablePropsMixin.IcebergSortFieldProperty",
		reflect.TypeOf((*CfnTablePropsMixin_IcebergSortFieldProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTablePropsMixin.IcebergSortOrderProperty",
		reflect.TypeOf((*CfnTablePropsMixin_IcebergSortOrderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTablePropsMixin.IcebergStructFieldProperty",
		reflect.TypeOf((*CfnTablePropsMixin_IcebergStructFieldProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTablePropsMixin.IcebergTableInputProperty",
		reflect.TypeOf((*CfnTablePropsMixin_IcebergTableInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTablePropsMixin.OpenTableFormatInputProperty",
		reflect.TypeOf((*CfnTablePropsMixin_OpenTableFormatInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTablePropsMixin.OrderProperty",
		reflect.TypeOf((*CfnTablePropsMixin_OrderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTablePropsMixin.SchemaIdProperty",
		reflect.TypeOf((*CfnTablePropsMixin_SchemaIdProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTablePropsMixin.SchemaReferenceProperty",
		reflect.TypeOf((*CfnTablePropsMixin_SchemaReferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTablePropsMixin.SerdeInfoProperty",
		reflect.TypeOf((*CfnTablePropsMixin_SerdeInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTablePropsMixin.SkewedInfoProperty",
		reflect.TypeOf((*CfnTablePropsMixin_SkewedInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTablePropsMixin.StorageDescriptorProperty",
		reflect.TypeOf((*CfnTablePropsMixin_StorageDescriptorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTablePropsMixin.TableIdentifierProperty",
		reflect.TypeOf((*CfnTablePropsMixin_TableIdentifierProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTablePropsMixin.TableInputProperty",
		reflect.TypeOf((*CfnTablePropsMixin_TableInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTablePropsMixin.ViewDefinitionProperty",
		reflect.TypeOf((*CfnTablePropsMixin_ViewDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTablePropsMixin.ViewRepresentationProperty",
		reflect.TypeOf((*CfnTablePropsMixin_ViewRepresentationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTriggerMixinProps",
		reflect.TypeOf((*CfnTriggerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTriggerPropsMixin",
		reflect.TypeOf((*CfnTriggerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTriggerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTriggerPropsMixin.ActionProperty",
		reflect.TypeOf((*CfnTriggerPropsMixin_ActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTriggerPropsMixin.ConditionProperty",
		reflect.TypeOf((*CfnTriggerPropsMixin_ConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTriggerPropsMixin.EventBatchingConditionProperty",
		reflect.TypeOf((*CfnTriggerPropsMixin_EventBatchingConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTriggerPropsMixin.NotificationPropertyProperty",
		reflect.TypeOf((*CfnTriggerPropsMixin_NotificationPropertyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnTriggerPropsMixin.PredicateProperty",
		reflect.TypeOf((*CfnTriggerPropsMixin_PredicateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnUsageProfileMixinProps",
		reflect.TypeOf((*CfnUsageProfileMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnUsageProfilePropsMixin",
		reflect.TypeOf((*CfnUsageProfilePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUsageProfilePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnUsageProfilePropsMixin.ConfigurationObjectProperty",
		reflect.TypeOf((*CfnUsageProfilePropsMixin_ConfigurationObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnUsageProfilePropsMixin.ProfileConfigurationProperty",
		reflect.TypeOf((*CfnUsageProfilePropsMixin_ProfileConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnWorkflowMixinProps",
		reflect.TypeOf((*CfnWorkflowMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnWorkflowPropsMixin",
		reflect.TypeOf((*CfnWorkflowPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWorkflowPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
