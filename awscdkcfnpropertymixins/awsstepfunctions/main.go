package awsstepfunctions

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_stepfunctions.CfnActivityMixinProps",
		reflect.TypeOf((*CfnActivityMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_stepfunctions.CfnActivityPropsMixin",
		reflect.TypeOf((*CfnActivityPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnActivityPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_stepfunctions.CfnActivityPropsMixin.EncryptionConfigurationProperty",
		reflect.TypeOf((*CfnActivityPropsMixin_EncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_stepfunctions.CfnActivityPropsMixin.TagsEntryProperty",
		reflect.TypeOf((*CfnActivityPropsMixin_TagsEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_stepfunctions.CfnStateMachineAliasMixinProps",
		reflect.TypeOf((*CfnStateMachineAliasMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_stepfunctions.CfnStateMachineAliasPropsMixin",
		reflect.TypeOf((*CfnStateMachineAliasPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStateMachineAliasPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_stepfunctions.CfnStateMachineAliasPropsMixin.DeploymentPreferenceProperty",
		reflect.TypeOf((*CfnStateMachineAliasPropsMixin_DeploymentPreferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_stepfunctions.CfnStateMachineAliasPropsMixin.RoutingConfigurationVersionProperty",
		reflect.TypeOf((*CfnStateMachineAliasPropsMixin_RoutingConfigurationVersionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_stepfunctions.CfnStateMachineMixinProps",
		reflect.TypeOf((*CfnStateMachineMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_stepfunctions.CfnStateMachinePropsMixin",
		reflect.TypeOf((*CfnStateMachinePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStateMachinePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_stepfunctions.CfnStateMachinePropsMixin.CloudWatchLogsLogGroupProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_CloudWatchLogsLogGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_stepfunctions.CfnStateMachinePropsMixin.EncryptionConfigurationProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_EncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_stepfunctions.CfnStateMachinePropsMixin.LogDestinationProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_LogDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_stepfunctions.CfnStateMachinePropsMixin.LoggingConfigurationProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_LoggingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_stepfunctions.CfnStateMachinePropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_stepfunctions.CfnStateMachinePropsMixin.TagsEntryProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_TagsEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_stepfunctions.CfnStateMachinePropsMixin.TracingConfigurationProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_TracingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_stepfunctions.CfnStateMachineVersionMixinProps",
		reflect.TypeOf((*CfnStateMachineVersionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_stepfunctions.CfnStateMachineVersionPropsMixin",
		reflect.TypeOf((*CfnStateMachineVersionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStateMachineVersionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
