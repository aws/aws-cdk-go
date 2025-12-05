package previewawslambdamixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnAliasMixinProps",
		reflect.TypeOf((*CfnAliasMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnAliasPropsMixin",
		reflect.TypeOf((*CfnAliasPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAliasPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnAliasPropsMixin.AliasRoutingConfigurationProperty",
		reflect.TypeOf((*CfnAliasPropsMixin_AliasRoutingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnAliasPropsMixin.ProvisionedConcurrencyConfigurationProperty",
		reflect.TypeOf((*CfnAliasPropsMixin_ProvisionedConcurrencyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnAliasPropsMixin.VersionWeightProperty",
		reflect.TypeOf((*CfnAliasPropsMixin_VersionWeightProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnCapacityProviderMixinProps",
		reflect.TypeOf((*CfnCapacityProviderMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnCapacityProviderPropsMixin",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCapacityProviderPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnCapacityProviderPropsMixin.CapacityProviderPermissionsConfigProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_CapacityProviderPermissionsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnCapacityProviderPropsMixin.CapacityProviderScalingConfigProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_CapacityProviderScalingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnCapacityProviderPropsMixin.CapacityProviderVpcConfigProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_CapacityProviderVpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnCapacityProviderPropsMixin.InstanceRequirementsProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_InstanceRequirementsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnCapacityProviderPropsMixin.TargetTrackingScalingPolicyProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_TargetTrackingScalingPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnCodeSigningConfigMixinProps",
		reflect.TypeOf((*CfnCodeSigningConfigMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnCodeSigningConfigPropsMixin",
		reflect.TypeOf((*CfnCodeSigningConfigPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCodeSigningConfigPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnCodeSigningConfigPropsMixin.AllowedPublishersProperty",
		reflect.TypeOf((*CfnCodeSigningConfigPropsMixin_AllowedPublishersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnCodeSigningConfigPropsMixin.CodeSigningPoliciesProperty",
		reflect.TypeOf((*CfnCodeSigningConfigPropsMixin_CodeSigningPoliciesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventInvokeConfigMixinProps",
		reflect.TypeOf((*CfnEventInvokeConfigMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventInvokeConfigPropsMixin",
		reflect.TypeOf((*CfnEventInvokeConfigPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEventInvokeConfigPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventInvokeConfigPropsMixin.DestinationConfigProperty",
		reflect.TypeOf((*CfnEventInvokeConfigPropsMixin_DestinationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventInvokeConfigPropsMixin.OnFailureProperty",
		reflect.TypeOf((*CfnEventInvokeConfigPropsMixin_OnFailureProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventInvokeConfigPropsMixin.OnSuccessProperty",
		reflect.TypeOf((*CfnEventInvokeConfigPropsMixin_OnSuccessProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventSourceMappingMixinProps",
		reflect.TypeOf((*CfnEventSourceMappingMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventSourceMappingPropsMixin",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEventSourceMappingPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventSourceMappingPropsMixin.AmazonManagedKafkaEventSourceConfigProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_AmazonManagedKafkaEventSourceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventSourceMappingPropsMixin.DestinationConfigProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_DestinationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventSourceMappingPropsMixin.DocumentDBEventSourceConfigProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_DocumentDBEventSourceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventSourceMappingPropsMixin.EndpointsProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_EndpointsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventSourceMappingPropsMixin.FilterCriteriaProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_FilterCriteriaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventSourceMappingPropsMixin.FilterProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_FilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventSourceMappingPropsMixin.LoggingConfigProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_LoggingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventSourceMappingPropsMixin.MetricsConfigProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_MetricsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventSourceMappingPropsMixin.OnFailureProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_OnFailureProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventSourceMappingPropsMixin.ProvisionedPollerConfigProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_ProvisionedPollerConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventSourceMappingPropsMixin.ScalingConfigProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_ScalingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventSourceMappingPropsMixin.SchemaRegistryAccessConfigProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_SchemaRegistryAccessConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventSourceMappingPropsMixin.SchemaRegistryConfigProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_SchemaRegistryConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventSourceMappingPropsMixin.SchemaValidationConfigProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_SchemaValidationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventSourceMappingPropsMixin.SelfManagedEventSourceProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_SelfManagedEventSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventSourceMappingPropsMixin.SelfManagedKafkaEventSourceConfigProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_SelfManagedKafkaEventSourceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventSourceMappingPropsMixin.SourceAccessConfigurationProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_SourceAccessConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnFunctionMixinProps",
		reflect.TypeOf((*CfnFunctionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnFunctionPropsMixin",
		reflect.TypeOf((*CfnFunctionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFunctionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnFunctionPropsMixin.CapacityProviderConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_CapacityProviderConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnFunctionPropsMixin.CodeProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_CodeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnFunctionPropsMixin.DeadLetterConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_DeadLetterConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnFunctionPropsMixin.DurableConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_DurableConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnFunctionPropsMixin.EnvironmentProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_EnvironmentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnFunctionPropsMixin.EphemeralStorageProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_EphemeralStorageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnFunctionPropsMixin.FileSystemConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_FileSystemConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnFunctionPropsMixin.FunctionScalingConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_FunctionScalingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnFunctionPropsMixin.ImageConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_ImageConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnFunctionPropsMixin.LambdaManagedInstancesCapacityProviderConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_LambdaManagedInstancesCapacityProviderConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnFunctionPropsMixin.LoggingConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_LoggingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnFunctionPropsMixin.RuntimeManagementConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_RuntimeManagementConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnFunctionPropsMixin.SnapStartProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_SnapStartProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnFunctionPropsMixin.SnapStartResponseProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_SnapStartResponseProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnFunctionPropsMixin.TenancyConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_TenancyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnFunctionPropsMixin.TracingConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_TracingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnFunctionPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnLayerVersionMixinProps",
		reflect.TypeOf((*CfnLayerVersionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnLayerVersionPermissionMixinProps",
		reflect.TypeOf((*CfnLayerVersionPermissionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnLayerVersionPermissionPropsMixin",
		reflect.TypeOf((*CfnLayerVersionPermissionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLayerVersionPermissionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnLayerVersionPropsMixin",
		reflect.TypeOf((*CfnLayerVersionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLayerVersionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnLayerVersionPropsMixin.ContentProperty",
		reflect.TypeOf((*CfnLayerVersionPropsMixin_ContentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnPermissionMixinProps",
		reflect.TypeOf((*CfnPermissionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnPermissionPropsMixin",
		reflect.TypeOf((*CfnPermissionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPermissionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnUrlMixinProps",
		reflect.TypeOf((*CfnUrlMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnUrlPropsMixin",
		reflect.TypeOf((*CfnUrlPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUrlPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnUrlPropsMixin.CorsProperty",
		reflect.TypeOf((*CfnUrlPropsMixin_CorsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnVersionMixinProps",
		reflect.TypeOf((*CfnVersionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnVersionPropsMixin",
		reflect.TypeOf((*CfnVersionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVersionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnVersionPropsMixin.FunctionScalingConfigProperty",
		reflect.TypeOf((*CfnVersionPropsMixin_FunctionScalingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnVersionPropsMixin.ProvisionedConcurrencyConfigurationProperty",
		reflect.TypeOf((*CfnVersionPropsMixin_ProvisionedConcurrencyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnVersionPropsMixin.RuntimePolicyProperty",
		reflect.TypeOf((*CfnVersionPropsMixin_RuntimePolicyProperty)(nil)).Elem(),
	)
}
