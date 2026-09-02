package awsgreengrassv2

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_greengrassv2.CfnComponentVersionMixinProps",
		reflect.TypeOf((*CfnComponentVersionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_greengrassv2.CfnComponentVersionPropsMixin",
		reflect.TypeOf((*CfnComponentVersionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnComponentVersionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_greengrassv2.CfnComponentVersionPropsMixin.ComponentDependencyRequirementProperty",
		reflect.TypeOf((*CfnComponentVersionPropsMixin_ComponentDependencyRequirementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_greengrassv2.CfnComponentVersionPropsMixin.ComponentPlatformProperty",
		reflect.TypeOf((*CfnComponentVersionPropsMixin_ComponentPlatformProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_greengrassv2.CfnComponentVersionPropsMixin.LambdaContainerParamsProperty",
		reflect.TypeOf((*CfnComponentVersionPropsMixin_LambdaContainerParamsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_greengrassv2.CfnComponentVersionPropsMixin.LambdaDeviceMountProperty",
		reflect.TypeOf((*CfnComponentVersionPropsMixin_LambdaDeviceMountProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_greengrassv2.CfnComponentVersionPropsMixin.LambdaEventSourceProperty",
		reflect.TypeOf((*CfnComponentVersionPropsMixin_LambdaEventSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_greengrassv2.CfnComponentVersionPropsMixin.LambdaExecutionParametersProperty",
		reflect.TypeOf((*CfnComponentVersionPropsMixin_LambdaExecutionParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_greengrassv2.CfnComponentVersionPropsMixin.LambdaFunctionRecipeSourceProperty",
		reflect.TypeOf((*CfnComponentVersionPropsMixin_LambdaFunctionRecipeSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_greengrassv2.CfnComponentVersionPropsMixin.LambdaLinuxProcessParamsProperty",
		reflect.TypeOf((*CfnComponentVersionPropsMixin_LambdaLinuxProcessParamsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_greengrassv2.CfnComponentVersionPropsMixin.LambdaVolumeMountProperty",
		reflect.TypeOf((*CfnComponentVersionPropsMixin_LambdaVolumeMountProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_greengrassv2.CfnDeploymentMixinProps",
		reflect.TypeOf((*CfnDeploymentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_greengrassv2.CfnDeploymentPropsMixin",
		reflect.TypeOf((*CfnDeploymentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDeploymentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_greengrassv2.CfnDeploymentPropsMixin.ComponentConfigurationUpdateProperty",
		reflect.TypeOf((*CfnDeploymentPropsMixin_ComponentConfigurationUpdateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_greengrassv2.CfnDeploymentPropsMixin.ComponentDeploymentSpecificationProperty",
		reflect.TypeOf((*CfnDeploymentPropsMixin_ComponentDeploymentSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_greengrassv2.CfnDeploymentPropsMixin.ComponentRunWithProperty",
		reflect.TypeOf((*CfnDeploymentPropsMixin_ComponentRunWithProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_greengrassv2.CfnDeploymentPropsMixin.DeploymentComponentUpdatePolicyProperty",
		reflect.TypeOf((*CfnDeploymentPropsMixin_DeploymentComponentUpdatePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_greengrassv2.CfnDeploymentPropsMixin.DeploymentConfigurationValidationPolicyProperty",
		reflect.TypeOf((*CfnDeploymentPropsMixin_DeploymentConfigurationValidationPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_greengrassv2.CfnDeploymentPropsMixin.DeploymentIoTJobConfigurationProperty",
		reflect.TypeOf((*CfnDeploymentPropsMixin_DeploymentIoTJobConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_greengrassv2.CfnDeploymentPropsMixin.DeploymentPoliciesProperty",
		reflect.TypeOf((*CfnDeploymentPropsMixin_DeploymentPoliciesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_greengrassv2.CfnDeploymentPropsMixin.IoTJobAbortConfigProperty",
		reflect.TypeOf((*CfnDeploymentPropsMixin_IoTJobAbortConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_greengrassv2.CfnDeploymentPropsMixin.IoTJobAbortCriteriaProperty",
		reflect.TypeOf((*CfnDeploymentPropsMixin_IoTJobAbortCriteriaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_greengrassv2.CfnDeploymentPropsMixin.IoTJobExecutionsRolloutConfigProperty",
		reflect.TypeOf((*CfnDeploymentPropsMixin_IoTJobExecutionsRolloutConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_greengrassv2.CfnDeploymentPropsMixin.IoTJobExponentialRolloutRateProperty",
		reflect.TypeOf((*CfnDeploymentPropsMixin_IoTJobExponentialRolloutRateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_greengrassv2.CfnDeploymentPropsMixin.IoTJobTimeoutConfigProperty",
		reflect.TypeOf((*CfnDeploymentPropsMixin_IoTJobTimeoutConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_greengrassv2.CfnDeploymentPropsMixin.SystemResourceLimitsProperty",
		reflect.TypeOf((*CfnDeploymentPropsMixin_SystemResourceLimitsProperty)(nil)).Elem(),
	)
}
