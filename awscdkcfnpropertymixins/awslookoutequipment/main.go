package awslookoutequipment

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lookoutequipment.CfnInferenceSchedulerMixinProps",
		reflect.TypeOf((*CfnInferenceSchedulerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lookoutequipment.CfnInferenceSchedulerPropsMixin",
		reflect.TypeOf((*CfnInferenceSchedulerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInferenceSchedulerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lookoutequipment.CfnInferenceSchedulerPropsMixin.DataInputConfigurationProperty",
		reflect.TypeOf((*CfnInferenceSchedulerPropsMixin_DataInputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lookoutequipment.CfnInferenceSchedulerPropsMixin.DataOutputConfigurationProperty",
		reflect.TypeOf((*CfnInferenceSchedulerPropsMixin_DataOutputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lookoutequipment.CfnInferenceSchedulerPropsMixin.InputNameConfigurationProperty",
		reflect.TypeOf((*CfnInferenceSchedulerPropsMixin_InputNameConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lookoutequipment.CfnInferenceSchedulerPropsMixin.S3InputConfigurationProperty",
		reflect.TypeOf((*CfnInferenceSchedulerPropsMixin_S3InputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lookoutequipment.CfnInferenceSchedulerPropsMixin.S3OutputConfigurationProperty",
		reflect.TypeOf((*CfnInferenceSchedulerPropsMixin_S3OutputConfigurationProperty)(nil)).Elem(),
	)
}
