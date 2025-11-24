package mixinsawsarczonalshift

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_arczonalshift.mixins.CfnAutoshiftObserverNotificationStatusMixinProps",
		reflect.TypeOf((*CfnAutoshiftObserverNotificationStatusMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_arczonalshift.mixins.CfnAutoshiftObserverNotificationStatusPropsMixin",
		reflect.TypeOf((*CfnAutoshiftObserverNotificationStatusPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAutoshiftObserverNotificationStatusPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_arczonalshift.mixins.CfnZonalAutoshiftConfigurationMixinProps",
		reflect.TypeOf((*CfnZonalAutoshiftConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_arczonalshift.mixins.CfnZonalAutoshiftConfigurationPropsMixin",
		reflect.TypeOf((*CfnZonalAutoshiftConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnZonalAutoshiftConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_arczonalshift.mixins.CfnZonalAutoshiftConfigurationPropsMixin.ControlConditionProperty",
		reflect.TypeOf((*CfnZonalAutoshiftConfigurationPropsMixin_ControlConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_arczonalshift.mixins.CfnZonalAutoshiftConfigurationPropsMixin.PracticeRunConfigurationProperty",
		reflect.TypeOf((*CfnZonalAutoshiftConfigurationPropsMixin_PracticeRunConfigurationProperty)(nil)).Elem(),
	)
}
