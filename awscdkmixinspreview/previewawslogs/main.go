package previewawslogs

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_logs.FirehoseLogsDelivery",
		reflect.TypeOf((*FirehoseLogsDelivery)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_FirehoseLogsDelivery{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ILogsDelivery)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/mixins-preview.aws_logs.ILogsDelivery",
		reflect.TypeOf((*ILogsDelivery)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_ILogsDelivery{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/mixins-preview.aws_logs.ILogsDeliveryConfig",
		reflect.TypeOf((*ILogsDeliveryConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "delivery", GoGetter: "Delivery"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryDestination", GoGetter: "DeliveryDestination"},
			_jsii_.MemberProperty{JsiiProperty: "deliverySource", GoGetter: "DeliverySource"},
		},
		func() interface{} {
			return &jsiiProxy_ILogsDeliveryConfig{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_logs.LogGroupLogsDelivery",
		reflect.TypeOf((*LogGroupLogsDelivery)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_LogGroupLogsDelivery{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ILogsDelivery)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_logs.S3LogsDelivery",
		reflect.TypeOf((*S3LogsDelivery)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_S3LogsDelivery{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ILogsDelivery)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_logs.S3LogsDeliveryPermissionsVersion",
		reflect.TypeOf((*S3LogsDeliveryPermissionsVersion)(nil)).Elem(),
		map[string]interface{}{
			"V1": S3LogsDeliveryPermissionsVersion_V1,
			"V2": S3LogsDeliveryPermissionsVersion_V2,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.S3LogsDeliveryProps",
		reflect.TypeOf((*S3LogsDeliveryProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_logs.XRayLogsDelivery",
		reflect.TypeOf((*XRayLogsDelivery)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_XRayLogsDelivery{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ILogsDelivery)
			return &j
		},
	)
}
