package awsivschat

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ivschat.CfnLoggingConfigurationMixinProps",
		reflect.TypeOf((*CfnLoggingConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ivschat.CfnLoggingConfigurationPropsMixin",
		reflect.TypeOf((*CfnLoggingConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLoggingConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ivschat.CfnLoggingConfigurationPropsMixin.CloudWatchLogsDestinationConfigurationProperty",
		reflect.TypeOf((*CfnLoggingConfigurationPropsMixin_CloudWatchLogsDestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ivschat.CfnLoggingConfigurationPropsMixin.DestinationConfigurationProperty",
		reflect.TypeOf((*CfnLoggingConfigurationPropsMixin_DestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ivschat.CfnLoggingConfigurationPropsMixin.FirehoseDestinationConfigurationProperty",
		reflect.TypeOf((*CfnLoggingConfigurationPropsMixin_FirehoseDestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ivschat.CfnLoggingConfigurationPropsMixin.S3DestinationConfigurationProperty",
		reflect.TypeOf((*CfnLoggingConfigurationPropsMixin_S3DestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ivschat.CfnRoomMixinProps",
		reflect.TypeOf((*CfnRoomMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ivschat.CfnRoomPropsMixin",
		reflect.TypeOf((*CfnRoomPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRoomPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ivschat.CfnRoomPropsMixin.MessageReviewHandlerProperty",
		reflect.TypeOf((*CfnRoomPropsMixin_MessageReviewHandlerProperty)(nil)).Elem(),
	)
}
