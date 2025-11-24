package mixinsawsstepfunctions

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnActivityMixinProps",
		reflect.TypeOf((*CfnActivityMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnActivityPropsMixin",
		reflect.TypeOf((*CfnActivityPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnActivityPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnActivityPropsMixin.EncryptionConfigurationProperty",
		reflect.TypeOf((*CfnActivityPropsMixin_EncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnActivityPropsMixin.TagsEntryProperty",
		reflect.TypeOf((*CfnActivityPropsMixin_TagsEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineAliasMixinProps",
		reflect.TypeOf((*CfnStateMachineAliasMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineAliasPropsMixin",
		reflect.TypeOf((*CfnStateMachineAliasPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStateMachineAliasPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineAliasPropsMixin.DeploymentPreferenceProperty",
		reflect.TypeOf((*CfnStateMachineAliasPropsMixin_DeploymentPreferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineAliasPropsMixin.RoutingConfigurationVersionProperty",
		reflect.TypeOf((*CfnStateMachineAliasPropsMixin_RoutingConfigurationVersionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineMixinProps",
		reflect.TypeOf((*CfnStateMachineMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachinePropsMixin",
		reflect.TypeOf((*CfnStateMachinePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStateMachinePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachinePropsMixin.CloudWatchLogsLogGroupProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_CloudWatchLogsLogGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachinePropsMixin.EncryptionConfigurationProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_EncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachinePropsMixin.LogDestinationProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_LogDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachinePropsMixin.LoggingConfigurationProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_LoggingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachinePropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachinePropsMixin.TagsEntryProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_TagsEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachinePropsMixin.TracingConfigurationProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_TracingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineVersionMixinProps",
		reflect.TypeOf((*CfnStateMachineVersionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineVersionPropsMixin",
		reflect.TypeOf((*CfnStateMachineVersionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStateMachineVersionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
}
