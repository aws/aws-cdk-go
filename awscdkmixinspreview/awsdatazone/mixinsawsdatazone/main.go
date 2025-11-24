package mixinsawsdatazone

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionMixinProps",
		reflect.TypeOf((*CfnConnectionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionPropsMixin.AmazonQPropertiesInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_AmazonQPropertiesInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionPropsMixin.AthenaPropertiesInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_AthenaPropertiesInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionPropsMixin.AuthenticationConfigurationInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_AuthenticationConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionPropsMixin.AuthorizationCodePropertiesProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_AuthorizationCodePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionPropsMixin.AwsLocationProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_AwsLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionPropsMixin.BasicAuthenticationCredentialsProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_BasicAuthenticationCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionPropsMixin.ConnectionPropertiesInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_ConnectionPropertiesInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionPropsMixin.GlueConnectionInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_GlueConnectionInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionPropsMixin.GlueOAuth2CredentialsProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_GlueOAuth2CredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionPropsMixin.GluePropertiesInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_GluePropertiesInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionPropsMixin.HyperPodPropertiesInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_HyperPodPropertiesInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionPropsMixin.IamPropertiesInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_IamPropertiesInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionPropsMixin.LineageSyncScheduleProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_LineageSyncScheduleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionPropsMixin.OAuth2ClientApplicationProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_OAuth2ClientApplicationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionPropsMixin.OAuth2PropertiesProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_OAuth2PropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionPropsMixin.PhysicalConnectionRequirementsProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_PhysicalConnectionRequirementsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionPropsMixin.RedshiftCredentialsProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_RedshiftCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionPropsMixin.RedshiftLineageSyncConfigurationInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_RedshiftLineageSyncConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionPropsMixin.RedshiftPropertiesInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_RedshiftPropertiesInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionPropsMixin.RedshiftStoragePropertiesProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_RedshiftStoragePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionPropsMixin.S3PropertiesInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_S3PropertiesInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionPropsMixin.SparkEmrPropertiesInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_SparkEmrPropertiesInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionPropsMixin.SparkGlueArgsProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_SparkGlueArgsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionPropsMixin.SparkGluePropertiesInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_SparkGluePropertiesInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionPropsMixin.UsernamePasswordProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_UsernamePasswordProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnDataSourceMixinProps",
		reflect.TypeOf((*CfnDataSourceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnDataSourcePropsMixin",
		reflect.TypeOf((*CfnDataSourcePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataSourcePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnDataSourcePropsMixin.DataSourceConfigurationInputProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_DataSourceConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnDataSourcePropsMixin.FilterExpressionProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_FilterExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnDataSourcePropsMixin.FormInputProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_FormInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnDataSourcePropsMixin.GlueRunConfigurationInputProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_GlueRunConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnDataSourcePropsMixin.RecommendationConfigurationProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_RecommendationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnDataSourcePropsMixin.RedshiftClusterStorageProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_RedshiftClusterStorageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnDataSourcePropsMixin.RedshiftCredentialConfigurationProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_RedshiftCredentialConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnDataSourcePropsMixin.RedshiftRunConfigurationInputProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_RedshiftRunConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnDataSourcePropsMixin.RedshiftServerlessStorageProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_RedshiftServerlessStorageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnDataSourcePropsMixin.RedshiftStorageProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_RedshiftStorageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnDataSourcePropsMixin.RelationalFilterConfigurationProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_RelationalFilterConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnDataSourcePropsMixin.SageMakerRunConfigurationInputProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_SageMakerRunConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnDataSourcePropsMixin.ScheduleConfigurationProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_ScheduleConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnDomainMixinProps",
		reflect.TypeOf((*CfnDomainMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnDomainPropsMixin",
		reflect.TypeOf((*CfnDomainPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDomainPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnDomainPropsMixin.SingleSignOnProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_SingleSignOnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnDomainUnitMixinProps",
		reflect.TypeOf((*CfnDomainUnitMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnDomainUnitPropsMixin",
		reflect.TypeOf((*CfnDomainUnitPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDomainUnitPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnEnvironmentActionsMixinProps",
		reflect.TypeOf((*CfnEnvironmentActionsMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnEnvironmentActionsPropsMixin",
		reflect.TypeOf((*CfnEnvironmentActionsPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEnvironmentActionsPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnEnvironmentActionsPropsMixin.AwsConsoleLinkParametersProperty",
		reflect.TypeOf((*CfnEnvironmentActionsPropsMixin_AwsConsoleLinkParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnEnvironmentBlueprintConfigurationMixinProps",
		reflect.TypeOf((*CfnEnvironmentBlueprintConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnEnvironmentBlueprintConfigurationPropsMixin",
		reflect.TypeOf((*CfnEnvironmentBlueprintConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEnvironmentBlueprintConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnEnvironmentBlueprintConfigurationPropsMixin.LakeFormationConfigurationProperty",
		reflect.TypeOf((*CfnEnvironmentBlueprintConfigurationPropsMixin_LakeFormationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnEnvironmentBlueprintConfigurationPropsMixin.ProvisioningConfigurationProperty",
		reflect.TypeOf((*CfnEnvironmentBlueprintConfigurationPropsMixin_ProvisioningConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnEnvironmentBlueprintConfigurationPropsMixin.RegionalParameterProperty",
		reflect.TypeOf((*CfnEnvironmentBlueprintConfigurationPropsMixin_RegionalParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnEnvironmentMixinProps",
		reflect.TypeOf((*CfnEnvironmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnEnvironmentProfileMixinProps",
		reflect.TypeOf((*CfnEnvironmentProfileMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnEnvironmentProfilePropsMixin",
		reflect.TypeOf((*CfnEnvironmentProfilePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEnvironmentProfilePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnEnvironmentProfilePropsMixin.EnvironmentParameterProperty",
		reflect.TypeOf((*CfnEnvironmentProfilePropsMixin_EnvironmentParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnEnvironmentPropsMixin",
		reflect.TypeOf((*CfnEnvironmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEnvironmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnEnvironmentPropsMixin.EnvironmentParameterProperty",
		reflect.TypeOf((*CfnEnvironmentPropsMixin_EnvironmentParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnFormTypeMixinProps",
		reflect.TypeOf((*CfnFormTypeMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnFormTypePropsMixin",
		reflect.TypeOf((*CfnFormTypePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFormTypePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnFormTypePropsMixin.ModelProperty",
		reflect.TypeOf((*CfnFormTypePropsMixin_ModelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnGroupProfileMixinProps",
		reflect.TypeOf((*CfnGroupProfileMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnGroupProfilePropsMixin",
		reflect.TypeOf((*CfnGroupProfilePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGroupProfilePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnOwnerMixinProps",
		reflect.TypeOf((*CfnOwnerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnOwnerPropsMixin",
		reflect.TypeOf((*CfnOwnerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOwnerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnOwnerPropsMixin.OwnerGroupPropertiesProperty",
		reflect.TypeOf((*CfnOwnerPropsMixin_OwnerGroupPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnOwnerPropsMixin.OwnerPropertiesProperty",
		reflect.TypeOf((*CfnOwnerPropsMixin_OwnerPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnOwnerPropsMixin.OwnerUserPropertiesProperty",
		reflect.TypeOf((*CfnOwnerPropsMixin_OwnerUserPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnPolicyGrantMixinProps",
		reflect.TypeOf((*CfnPolicyGrantMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnPolicyGrantPropsMixin",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPolicyGrantPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnPolicyGrantPropsMixin.AddToProjectMemberPoolPolicyGrantDetailProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_AddToProjectMemberPoolPolicyGrantDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnPolicyGrantPropsMixin.CreateAssetTypePolicyGrantDetailProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_CreateAssetTypePolicyGrantDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnPolicyGrantPropsMixin.CreateDomainUnitPolicyGrantDetailProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_CreateDomainUnitPolicyGrantDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnPolicyGrantPropsMixin.CreateEnvironmentProfilePolicyGrantDetailProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_CreateEnvironmentProfilePolicyGrantDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnPolicyGrantPropsMixin.CreateFormTypePolicyGrantDetailProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_CreateFormTypePolicyGrantDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnPolicyGrantPropsMixin.CreateGlossaryPolicyGrantDetailProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_CreateGlossaryPolicyGrantDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnPolicyGrantPropsMixin.CreateProjectFromProjectProfilePolicyGrantDetailProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_CreateProjectFromProjectProfilePolicyGrantDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnPolicyGrantPropsMixin.CreateProjectPolicyGrantDetailProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_CreateProjectPolicyGrantDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnPolicyGrantPropsMixin.DomainUnitFilterForProjectProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_DomainUnitFilterForProjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnPolicyGrantPropsMixin.DomainUnitGrantFilterProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_DomainUnitGrantFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnPolicyGrantPropsMixin.DomainUnitPolicyGrantPrincipalProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_DomainUnitPolicyGrantPrincipalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnPolicyGrantPropsMixin.GroupPolicyGrantPrincipalProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_GroupPolicyGrantPrincipalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnPolicyGrantPropsMixin.OverrideDomainUnitOwnersPolicyGrantDetailProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_OverrideDomainUnitOwnersPolicyGrantDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnPolicyGrantPropsMixin.OverrideProjectOwnersPolicyGrantDetailProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_OverrideProjectOwnersPolicyGrantDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnPolicyGrantPropsMixin.PolicyGrantDetailProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_PolicyGrantDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnPolicyGrantPropsMixin.PolicyGrantPrincipalProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_PolicyGrantPrincipalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnPolicyGrantPropsMixin.ProjectGrantFilterProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_ProjectGrantFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnPolicyGrantPropsMixin.ProjectPolicyGrantPrincipalProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_ProjectPolicyGrantPrincipalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnPolicyGrantPropsMixin.UserPolicyGrantPrincipalProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_UserPolicyGrantPrincipalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnProjectMembershipMixinProps",
		reflect.TypeOf((*CfnProjectMembershipMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnProjectMembershipPropsMixin",
		reflect.TypeOf((*CfnProjectMembershipPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProjectMembershipPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnProjectMembershipPropsMixin.MemberProperty",
		reflect.TypeOf((*CfnProjectMembershipPropsMixin_MemberProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnProjectMixinProps",
		reflect.TypeOf((*CfnProjectMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnProjectProfileMixinProps",
		reflect.TypeOf((*CfnProjectProfileMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnProjectProfilePropsMixin",
		reflect.TypeOf((*CfnProjectProfilePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProjectProfilePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnProjectProfilePropsMixin.AwsAccountProperty",
		reflect.TypeOf((*CfnProjectProfilePropsMixin_AwsAccountProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnProjectProfilePropsMixin.EnvironmentConfigurationParameterProperty",
		reflect.TypeOf((*CfnProjectProfilePropsMixin_EnvironmentConfigurationParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnProjectProfilePropsMixin.EnvironmentConfigurationParametersDetailsProperty",
		reflect.TypeOf((*CfnProjectProfilePropsMixin_EnvironmentConfigurationParametersDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnProjectProfilePropsMixin.EnvironmentConfigurationProperty",
		reflect.TypeOf((*CfnProjectProfilePropsMixin_EnvironmentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnProjectProfilePropsMixin.RegionProperty",
		reflect.TypeOf((*CfnProjectProfilePropsMixin_RegionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnProjectPropsMixin",
		reflect.TypeOf((*CfnProjectPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProjectPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnProjectPropsMixin.EnvironmentConfigurationUserParameterProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_EnvironmentConfigurationUserParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnProjectPropsMixin.EnvironmentParameterProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_EnvironmentParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnSubscriptionTargetMixinProps",
		reflect.TypeOf((*CfnSubscriptionTargetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnSubscriptionTargetPropsMixin",
		reflect.TypeOf((*CfnSubscriptionTargetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSubscriptionTargetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnSubscriptionTargetPropsMixin.SubscriptionTargetFormProperty",
		reflect.TypeOf((*CfnSubscriptionTargetPropsMixin_SubscriptionTargetFormProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnUserProfileMixinProps",
		reflect.TypeOf((*CfnUserProfileMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnUserProfilePropsMixin",
		reflect.TypeOf((*CfnUserProfilePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUserProfilePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnUserProfilePropsMixin.IamUserProfileDetailsProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_IamUserProfileDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnUserProfilePropsMixin.SsoUserProfileDetailsProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_SsoUserProfileDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnUserProfilePropsMixin.UserProfileDetailsProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_UserProfileDetailsProperty)(nil)).Elem(),
	)
}
