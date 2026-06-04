package awslambda

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnAliasMixinProps",
		reflect.TypeOf((*CfnAliasMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnAliasPropsMixin",
		reflect.TypeOf((*CfnAliasPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAliasPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnAliasPropsMixin.AliasRoutingConfigurationProperty",
		reflect.TypeOf((*CfnAliasPropsMixin_AliasRoutingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnAliasPropsMixin.ProvisionedConcurrencyConfigurationProperty",
		reflect.TypeOf((*CfnAliasPropsMixin_ProvisionedConcurrencyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnAliasPropsMixin.VersionWeightProperty",
		reflect.TypeOf((*CfnAliasPropsMixin_VersionWeightProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnCapacityProviderMixinProps",
		reflect.TypeOf((*CfnCapacityProviderMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnCapacityProviderPropsMixin",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCapacityProviderPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnCapacityProviderPropsMixin.CapacityProviderPermissionsConfigProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_CapacityProviderPermissionsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnCapacityProviderPropsMixin.CapacityProviderScalingConfigProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_CapacityProviderScalingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnCapacityProviderPropsMixin.CapacityProviderVpcConfigProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_CapacityProviderVpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnCapacityProviderPropsMixin.InstanceRequirementsProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_InstanceRequirementsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnCapacityProviderPropsMixin.PropagateTagsConfigProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_PropagateTagsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnCapacityProviderPropsMixin.TargetTrackingScalingPolicyProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_TargetTrackingScalingPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnCodeSigningConfigMixinProps",
		reflect.TypeOf((*CfnCodeSigningConfigMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnCodeSigningConfigPropsMixin",
		reflect.TypeOf((*CfnCodeSigningConfigPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCodeSigningConfigPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnCodeSigningConfigPropsMixin.AllowedPublishersProperty",
		reflect.TypeOf((*CfnCodeSigningConfigPropsMixin_AllowedPublishersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnCodeSigningConfigPropsMixin.CodeSigningPoliciesProperty",
		reflect.TypeOf((*CfnCodeSigningConfigPropsMixin_CodeSigningPoliciesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnEventInvokeConfigMixinProps",
		reflect.TypeOf((*CfnEventInvokeConfigMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnEventInvokeConfigPropsMixin",
		reflect.TypeOf((*CfnEventInvokeConfigPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEventInvokeConfigPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnEventInvokeConfigPropsMixin.DestinationConfigProperty",
		reflect.TypeOf((*CfnEventInvokeConfigPropsMixin_DestinationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnEventInvokeConfigPropsMixin.OnFailureProperty",
		reflect.TypeOf((*CfnEventInvokeConfigPropsMixin_OnFailureProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnEventInvokeConfigPropsMixin.OnSuccessProperty",
		reflect.TypeOf((*CfnEventInvokeConfigPropsMixin_OnSuccessProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnEventSourceMappingMixinProps",
		reflect.TypeOf((*CfnEventSourceMappingMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnEventSourceMappingPropsMixin",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEventSourceMappingPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnEventSourceMappingPropsMixin.AmazonManagedKafkaEventSourceConfigProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_AmazonManagedKafkaEventSourceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnEventSourceMappingPropsMixin.DestinationConfigProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_DestinationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnEventSourceMappingPropsMixin.DocumentDBEventSourceConfigProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_DocumentDBEventSourceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnEventSourceMappingPropsMixin.EndpointsProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_EndpointsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnEventSourceMappingPropsMixin.FilterCriteriaProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_FilterCriteriaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnEventSourceMappingPropsMixin.FilterProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_FilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnEventSourceMappingPropsMixin.LoggingConfigProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_LoggingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnEventSourceMappingPropsMixin.MetricsConfigProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_MetricsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnEventSourceMappingPropsMixin.OnFailureProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_OnFailureProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnEventSourceMappingPropsMixin.ProvisionedPollerConfigProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_ProvisionedPollerConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnEventSourceMappingPropsMixin.ScalingConfigProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_ScalingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnEventSourceMappingPropsMixin.SchemaRegistryAccessConfigProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_SchemaRegistryAccessConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnEventSourceMappingPropsMixin.SchemaRegistryConfigProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_SchemaRegistryConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnEventSourceMappingPropsMixin.SchemaValidationConfigProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_SchemaValidationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnEventSourceMappingPropsMixin.SelfManagedEventSourceProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_SelfManagedEventSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnEventSourceMappingPropsMixin.SelfManagedKafkaEventSourceConfigProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_SelfManagedKafkaEventSourceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnEventSourceMappingPropsMixin.SourceAccessConfigurationProperty",
		reflect.TypeOf((*CfnEventSourceMappingPropsMixin_SourceAccessConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnFunctionMixinProps",
		reflect.TypeOf((*CfnFunctionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnFunctionPropsMixin",
		reflect.TypeOf((*CfnFunctionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFunctionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnFunctionPropsMixin.CapacityProviderConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_CapacityProviderConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnFunctionPropsMixin.CodeProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_CodeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnFunctionPropsMixin.DeadLetterConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_DeadLetterConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnFunctionPropsMixin.DurableConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_DurableConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnFunctionPropsMixin.EnvironmentProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_EnvironmentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnFunctionPropsMixin.EphemeralStorageProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_EphemeralStorageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnFunctionPropsMixin.FileSystemConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_FileSystemConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnFunctionPropsMixin.FunctionScalingConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_FunctionScalingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnFunctionPropsMixin.ImageConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_ImageConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnFunctionPropsMixin.LambdaManagedInstancesCapacityProviderConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_LambdaManagedInstancesCapacityProviderConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnFunctionPropsMixin.LoggingConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_LoggingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnFunctionPropsMixin.RuntimeManagementConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_RuntimeManagementConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnFunctionPropsMixin.SnapStartProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_SnapStartProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnFunctionPropsMixin.SnapStartResponseProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_SnapStartResponseProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnFunctionPropsMixin.TenancyConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_TenancyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnFunctionPropsMixin.TracingConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_TracingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnFunctionPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnLayerVersionMixinProps",
		reflect.TypeOf((*CfnLayerVersionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnLayerVersionPermissionMixinProps",
		reflect.TypeOf((*CfnLayerVersionPermissionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnLayerVersionPermissionPropsMixin",
		reflect.TypeOf((*CfnLayerVersionPermissionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLayerVersionPermissionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnLayerVersionPropsMixin",
		reflect.TypeOf((*CfnLayerVersionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLayerVersionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnLayerVersionPropsMixin.ContentProperty",
		reflect.TypeOf((*CfnLayerVersionPropsMixin_ContentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnPermissionMixinProps",
		reflect.TypeOf((*CfnPermissionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnPermissionPropsMixin",
		reflect.TypeOf((*CfnPermissionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPermissionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnUrlMixinProps",
		reflect.TypeOf((*CfnUrlMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnUrlPropsMixin",
		reflect.TypeOf((*CfnUrlPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUrlPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnUrlPropsMixin.CorsProperty",
		reflect.TypeOf((*CfnUrlPropsMixin_CorsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnVersionMixinProps",
		reflect.TypeOf((*CfnVersionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnVersionPropsMixin",
		reflect.TypeOf((*CfnVersionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVersionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnVersionPropsMixin.FunctionScalingConfigProperty",
		reflect.TypeOf((*CfnVersionPropsMixin_FunctionScalingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnVersionPropsMixin.ProvisionedConcurrencyConfigurationProperty",
		reflect.TypeOf((*CfnVersionPropsMixin_ProvisionedConcurrencyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnVersionPropsMixin.RuntimePolicyProperty",
		reflect.TypeOf((*CfnVersionPropsMixin_RuntimePolicyProperty)(nil)).Elem(),
	)
}
