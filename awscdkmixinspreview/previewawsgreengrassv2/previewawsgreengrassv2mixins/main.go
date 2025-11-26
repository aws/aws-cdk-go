package previewawsgreengrassv2mixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnComponentVersionMixinProps",
		reflect.TypeOf((*CfnComponentVersionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnComponentVersionPropsMixin",
		reflect.TypeOf((*CfnComponentVersionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnComponentVersionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnComponentVersionPropsMixin.ComponentDependencyRequirementProperty",
		reflect.TypeOf((*CfnComponentVersionPropsMixin_ComponentDependencyRequirementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnComponentVersionPropsMixin.ComponentPlatformProperty",
		reflect.TypeOf((*CfnComponentVersionPropsMixin_ComponentPlatformProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnComponentVersionPropsMixin.LambdaContainerParamsProperty",
		reflect.TypeOf((*CfnComponentVersionPropsMixin_LambdaContainerParamsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnComponentVersionPropsMixin.LambdaDeviceMountProperty",
		reflect.TypeOf((*CfnComponentVersionPropsMixin_LambdaDeviceMountProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnComponentVersionPropsMixin.LambdaEventSourceProperty",
		reflect.TypeOf((*CfnComponentVersionPropsMixin_LambdaEventSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnComponentVersionPropsMixin.LambdaExecutionParametersProperty",
		reflect.TypeOf((*CfnComponentVersionPropsMixin_LambdaExecutionParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnComponentVersionPropsMixin.LambdaFunctionRecipeSourceProperty",
		reflect.TypeOf((*CfnComponentVersionPropsMixin_LambdaFunctionRecipeSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnComponentVersionPropsMixin.LambdaLinuxProcessParamsProperty",
		reflect.TypeOf((*CfnComponentVersionPropsMixin_LambdaLinuxProcessParamsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnComponentVersionPropsMixin.LambdaVolumeMountProperty",
		reflect.TypeOf((*CfnComponentVersionPropsMixin_LambdaVolumeMountProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnDeploymentMixinProps",
		reflect.TypeOf((*CfnDeploymentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnDeploymentPropsMixin",
		reflect.TypeOf((*CfnDeploymentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDeploymentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnDeploymentPropsMixin.ComponentConfigurationUpdateProperty",
		reflect.TypeOf((*CfnDeploymentPropsMixin_ComponentConfigurationUpdateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnDeploymentPropsMixin.ComponentDeploymentSpecificationProperty",
		reflect.TypeOf((*CfnDeploymentPropsMixin_ComponentDeploymentSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnDeploymentPropsMixin.ComponentRunWithProperty",
		reflect.TypeOf((*CfnDeploymentPropsMixin_ComponentRunWithProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnDeploymentPropsMixin.DeploymentComponentUpdatePolicyProperty",
		reflect.TypeOf((*CfnDeploymentPropsMixin_DeploymentComponentUpdatePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnDeploymentPropsMixin.DeploymentConfigurationValidationPolicyProperty",
		reflect.TypeOf((*CfnDeploymentPropsMixin_DeploymentConfigurationValidationPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnDeploymentPropsMixin.DeploymentIoTJobConfigurationProperty",
		reflect.TypeOf((*CfnDeploymentPropsMixin_DeploymentIoTJobConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnDeploymentPropsMixin.DeploymentPoliciesProperty",
		reflect.TypeOf((*CfnDeploymentPropsMixin_DeploymentPoliciesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnDeploymentPropsMixin.IoTJobAbortConfigProperty",
		reflect.TypeOf((*CfnDeploymentPropsMixin_IoTJobAbortConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnDeploymentPropsMixin.IoTJobAbortCriteriaProperty",
		reflect.TypeOf((*CfnDeploymentPropsMixin_IoTJobAbortCriteriaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnDeploymentPropsMixin.IoTJobExecutionsRolloutConfigProperty",
		reflect.TypeOf((*CfnDeploymentPropsMixin_IoTJobExecutionsRolloutConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnDeploymentPropsMixin.IoTJobExponentialRolloutRateProperty",
		reflect.TypeOf((*CfnDeploymentPropsMixin_IoTJobExponentialRolloutRateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnDeploymentPropsMixin.IoTJobTimeoutConfigProperty",
		reflect.TypeOf((*CfnDeploymentPropsMixin_IoTJobTimeoutConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnDeploymentPropsMixin.SystemResourceLimitsProperty",
		reflect.TypeOf((*CfnDeploymentPropsMixin_SystemResourceLimitsProperty)(nil)).Elem(),
	)
}
