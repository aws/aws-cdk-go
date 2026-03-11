package awsdatazone

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnConnectionMixinProps",
		reflect.TypeOf((*CfnConnectionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnConnectionPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnConnectionPropsMixin.AmazonQPropertiesInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_AmazonQPropertiesInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnConnectionPropsMixin.AthenaPropertiesInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_AthenaPropertiesInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnConnectionPropsMixin.AuthenticationConfigurationInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_AuthenticationConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnConnectionPropsMixin.AuthorizationCodePropertiesProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_AuthorizationCodePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnConnectionPropsMixin.AwsLocationProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_AwsLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnConnectionPropsMixin.BasicAuthenticationCredentialsProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_BasicAuthenticationCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnConnectionPropsMixin.ConnectionPropertiesInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_ConnectionPropertiesInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnConnectionPropsMixin.GlueConnectionInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_GlueConnectionInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnConnectionPropsMixin.GlueOAuth2CredentialsProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_GlueOAuth2CredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnConnectionPropsMixin.GluePropertiesInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_GluePropertiesInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnConnectionPropsMixin.HyperPodPropertiesInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_HyperPodPropertiesInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnConnectionPropsMixin.IamPropertiesInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_IamPropertiesInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnConnectionPropsMixin.LineageSyncScheduleProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_LineageSyncScheduleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnConnectionPropsMixin.MlflowPropertiesInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_MlflowPropertiesInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnConnectionPropsMixin.OAuth2ClientApplicationProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_OAuth2ClientApplicationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnConnectionPropsMixin.OAuth2PropertiesProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_OAuth2PropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnConnectionPropsMixin.PhysicalConnectionRequirementsProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_PhysicalConnectionRequirementsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnConnectionPropsMixin.RedshiftCredentialsProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_RedshiftCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnConnectionPropsMixin.RedshiftLineageSyncConfigurationInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_RedshiftLineageSyncConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnConnectionPropsMixin.RedshiftPropertiesInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_RedshiftPropertiesInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnConnectionPropsMixin.RedshiftStoragePropertiesProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_RedshiftStoragePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnConnectionPropsMixin.S3PropertiesInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_S3PropertiesInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnConnectionPropsMixin.SparkEmrPropertiesInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_SparkEmrPropertiesInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnConnectionPropsMixin.SparkGlueArgsProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_SparkGlueArgsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnConnectionPropsMixin.SparkGluePropertiesInputProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_SparkGluePropertiesInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnConnectionPropsMixin.UsernamePasswordProperty",
		reflect.TypeOf((*CfnConnectionPropsMixin_UsernamePasswordProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnDataSourceMixinProps",
		reflect.TypeOf((*CfnDataSourceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnDataSourcePropsMixin",
		reflect.TypeOf((*CfnDataSourcePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataSourcePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnDataSourcePropsMixin.DataSourceConfigurationInputProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_DataSourceConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnDataSourcePropsMixin.FilterExpressionProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_FilterExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnDataSourcePropsMixin.FormInputProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_FormInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnDataSourcePropsMixin.GlueRunConfigurationInputProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_GlueRunConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnDataSourcePropsMixin.RecommendationConfigurationProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_RecommendationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnDataSourcePropsMixin.RedshiftClusterStorageProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_RedshiftClusterStorageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnDataSourcePropsMixin.RedshiftCredentialConfigurationProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_RedshiftCredentialConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnDataSourcePropsMixin.RedshiftRunConfigurationInputProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_RedshiftRunConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnDataSourcePropsMixin.RedshiftServerlessStorageProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_RedshiftServerlessStorageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnDataSourcePropsMixin.RedshiftStorageProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_RedshiftStorageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnDataSourcePropsMixin.RelationalFilterConfigurationProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_RelationalFilterConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnDataSourcePropsMixin.SageMakerRunConfigurationInputProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_SageMakerRunConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnDataSourcePropsMixin.ScheduleConfigurationProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_ScheduleConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnDomainMixinProps",
		reflect.TypeOf((*CfnDomainMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnDomainPropsMixin",
		reflect.TypeOf((*CfnDomainPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDomainPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnDomainPropsMixin.SingleSignOnProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_SingleSignOnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnDomainUnitMixinProps",
		reflect.TypeOf((*CfnDomainUnitMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnDomainUnitPropsMixin",
		reflect.TypeOf((*CfnDomainUnitPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDomainUnitPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnEnvironmentActionsMixinProps",
		reflect.TypeOf((*CfnEnvironmentActionsMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnEnvironmentActionsPropsMixin",
		reflect.TypeOf((*CfnEnvironmentActionsPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEnvironmentActionsPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnEnvironmentActionsPropsMixin.AwsConsoleLinkParametersProperty",
		reflect.TypeOf((*CfnEnvironmentActionsPropsMixin_AwsConsoleLinkParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnEnvironmentBlueprintConfigurationMixinProps",
		reflect.TypeOf((*CfnEnvironmentBlueprintConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnEnvironmentBlueprintConfigurationPropsMixin",
		reflect.TypeOf((*CfnEnvironmentBlueprintConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEnvironmentBlueprintConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnEnvironmentBlueprintConfigurationPropsMixin.LakeFormationConfigurationProperty",
		reflect.TypeOf((*CfnEnvironmentBlueprintConfigurationPropsMixin_LakeFormationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnEnvironmentBlueprintConfigurationPropsMixin.ProvisioningConfigurationProperty",
		reflect.TypeOf((*CfnEnvironmentBlueprintConfigurationPropsMixin_ProvisioningConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnEnvironmentBlueprintConfigurationPropsMixin.RegionalParameterProperty",
		reflect.TypeOf((*CfnEnvironmentBlueprintConfigurationPropsMixin_RegionalParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnEnvironmentMixinProps",
		reflect.TypeOf((*CfnEnvironmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnEnvironmentProfileMixinProps",
		reflect.TypeOf((*CfnEnvironmentProfileMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnEnvironmentProfilePropsMixin",
		reflect.TypeOf((*CfnEnvironmentProfilePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEnvironmentProfilePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnEnvironmentProfilePropsMixin.EnvironmentParameterProperty",
		reflect.TypeOf((*CfnEnvironmentProfilePropsMixin_EnvironmentParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnEnvironmentPropsMixin",
		reflect.TypeOf((*CfnEnvironmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEnvironmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnEnvironmentPropsMixin.EnvironmentParameterProperty",
		reflect.TypeOf((*CfnEnvironmentPropsMixin_EnvironmentParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnFormTypeMixinProps",
		reflect.TypeOf((*CfnFormTypeMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnFormTypePropsMixin",
		reflect.TypeOf((*CfnFormTypePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFormTypePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnFormTypePropsMixin.ModelProperty",
		reflect.TypeOf((*CfnFormTypePropsMixin_ModelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnGroupProfileMixinProps",
		reflect.TypeOf((*CfnGroupProfileMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnGroupProfilePropsMixin",
		reflect.TypeOf((*CfnGroupProfilePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGroupProfilePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnOwnerMixinProps",
		reflect.TypeOf((*CfnOwnerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnOwnerPropsMixin",
		reflect.TypeOf((*CfnOwnerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOwnerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnOwnerPropsMixin.OwnerGroupPropertiesProperty",
		reflect.TypeOf((*CfnOwnerPropsMixin_OwnerGroupPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnOwnerPropsMixin.OwnerPropertiesProperty",
		reflect.TypeOf((*CfnOwnerPropsMixin_OwnerPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnOwnerPropsMixin.OwnerUserPropertiesProperty",
		reflect.TypeOf((*CfnOwnerPropsMixin_OwnerUserPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnPolicyGrantMixinProps",
		reflect.TypeOf((*CfnPolicyGrantMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnPolicyGrantPropsMixin",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPolicyGrantPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnPolicyGrantPropsMixin.AddToProjectMemberPoolPolicyGrantDetailProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_AddToProjectMemberPoolPolicyGrantDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnPolicyGrantPropsMixin.CreateAssetTypePolicyGrantDetailProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_CreateAssetTypePolicyGrantDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnPolicyGrantPropsMixin.CreateDomainUnitPolicyGrantDetailProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_CreateDomainUnitPolicyGrantDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnPolicyGrantPropsMixin.CreateEnvironmentProfilePolicyGrantDetailProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_CreateEnvironmentProfilePolicyGrantDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnPolicyGrantPropsMixin.CreateFormTypePolicyGrantDetailProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_CreateFormTypePolicyGrantDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnPolicyGrantPropsMixin.CreateGlossaryPolicyGrantDetailProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_CreateGlossaryPolicyGrantDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnPolicyGrantPropsMixin.CreateProjectFromProjectProfilePolicyGrantDetailProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_CreateProjectFromProjectProfilePolicyGrantDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnPolicyGrantPropsMixin.CreateProjectPolicyGrantDetailProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_CreateProjectPolicyGrantDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnPolicyGrantPropsMixin.DomainUnitFilterForProjectProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_DomainUnitFilterForProjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnPolicyGrantPropsMixin.DomainUnitGrantFilterProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_DomainUnitGrantFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnPolicyGrantPropsMixin.DomainUnitPolicyGrantPrincipalProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_DomainUnitPolicyGrantPrincipalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnPolicyGrantPropsMixin.GroupPolicyGrantPrincipalProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_GroupPolicyGrantPrincipalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnPolicyGrantPropsMixin.OverrideDomainUnitOwnersPolicyGrantDetailProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_OverrideDomainUnitOwnersPolicyGrantDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnPolicyGrantPropsMixin.OverrideProjectOwnersPolicyGrantDetailProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_OverrideProjectOwnersPolicyGrantDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnPolicyGrantPropsMixin.PolicyGrantDetailProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_PolicyGrantDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnPolicyGrantPropsMixin.PolicyGrantPrincipalProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_PolicyGrantPrincipalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnPolicyGrantPropsMixin.ProjectGrantFilterProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_ProjectGrantFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnPolicyGrantPropsMixin.ProjectPolicyGrantPrincipalProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_ProjectPolicyGrantPrincipalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnPolicyGrantPropsMixin.UserPolicyGrantPrincipalProperty",
		reflect.TypeOf((*CfnPolicyGrantPropsMixin_UserPolicyGrantPrincipalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnProjectMembershipMixinProps",
		reflect.TypeOf((*CfnProjectMembershipMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnProjectMembershipPropsMixin",
		reflect.TypeOf((*CfnProjectMembershipPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProjectMembershipPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnProjectMembershipPropsMixin.MemberProperty",
		reflect.TypeOf((*CfnProjectMembershipPropsMixin_MemberProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnProjectMixinProps",
		reflect.TypeOf((*CfnProjectMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnProjectProfileMixinProps",
		reflect.TypeOf((*CfnProjectProfileMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnProjectProfilePropsMixin",
		reflect.TypeOf((*CfnProjectProfilePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProjectProfilePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnProjectProfilePropsMixin.AwsAccountProperty",
		reflect.TypeOf((*CfnProjectProfilePropsMixin_AwsAccountProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnProjectProfilePropsMixin.EnvironmentConfigurationParameterProperty",
		reflect.TypeOf((*CfnProjectProfilePropsMixin_EnvironmentConfigurationParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnProjectProfilePropsMixin.EnvironmentConfigurationParametersDetailsProperty",
		reflect.TypeOf((*CfnProjectProfilePropsMixin_EnvironmentConfigurationParametersDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnProjectProfilePropsMixin.EnvironmentConfigurationProperty",
		reflect.TypeOf((*CfnProjectProfilePropsMixin_EnvironmentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnProjectProfilePropsMixin.RegionProperty",
		reflect.TypeOf((*CfnProjectProfilePropsMixin_RegionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnProjectPropsMixin",
		reflect.TypeOf((*CfnProjectPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProjectPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnProjectPropsMixin.EnvironmentConfigurationUserParameterProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_EnvironmentConfigurationUserParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnProjectPropsMixin.EnvironmentParameterProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_EnvironmentParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnSubscriptionTargetMixinProps",
		reflect.TypeOf((*CfnSubscriptionTargetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnSubscriptionTargetPropsMixin",
		reflect.TypeOf((*CfnSubscriptionTargetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSubscriptionTargetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnSubscriptionTargetPropsMixin.SubscriptionTargetFormProperty",
		reflect.TypeOf((*CfnSubscriptionTargetPropsMixin_SubscriptionTargetFormProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnUserProfileMixinProps",
		reflect.TypeOf((*CfnUserProfileMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnUserProfilePropsMixin",
		reflect.TypeOf((*CfnUserProfilePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUserProfilePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnUserProfilePropsMixin.IamUserProfileDetailsProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_IamUserProfileDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnUserProfilePropsMixin.SsoUserProfileDetailsProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_SsoUserProfileDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnUserProfilePropsMixin.UserProfileDetailsProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_UserProfileDetailsProperty)(nil)).Elem(),
	)
}
