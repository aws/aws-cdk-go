// The CDK Construct Library for AWS::KinesisFirehose
package awscdkkinesisfirehosealpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-kinesisfirehose-alpha.DataProcessorBindOptions",
		reflect.TypeOf((*DataProcessorBindOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-kinesisfirehose-alpha.DataProcessorConfig",
		reflect.TypeOf((*DataProcessorConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-kinesisfirehose-alpha.DataProcessorIdentifier",
		reflect.TypeOf((*DataProcessorIdentifier)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-kinesisfirehose-alpha.DataProcessorProps",
		reflect.TypeOf((*DataProcessorProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-kinesisfirehose-alpha.DeliveryStream",
		reflect.TypeOf((*DeliveryStream)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamArn", GoGetter: "DeliveryStreamArn"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamName", GoGetter: "DeliveryStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantPutRecords", GoMethod: "GrantPutRecords"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricBackupToS3Bytes", GoMethod: "MetricBackupToS3Bytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricBackupToS3DataFreshness", GoMethod: "MetricBackupToS3DataFreshness"},
			_jsii_.MemberMethod{JsiiMethod: "metricBackupToS3Records", GoMethod: "MetricBackupToS3Records"},
			_jsii_.MemberMethod{JsiiMethod: "metricIncomingBytes", GoMethod: "MetricIncomingBytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricIncomingRecords", GoMethod: "MetricIncomingRecords"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DeliveryStream{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDeliveryStream)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-kinesisfirehose-alpha.DeliveryStreamAttributes",
		reflect.TypeOf((*DeliveryStreamAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-kinesisfirehose-alpha.DeliveryStreamProps",
		reflect.TypeOf((*DeliveryStreamProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-kinesisfirehose-alpha.DestinationBindOptions",
		reflect.TypeOf((*DestinationBindOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-kinesisfirehose-alpha.DestinationConfig",
		reflect.TypeOf((*DestinationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-kinesisfirehose-alpha.IDataProcessor",
		reflect.TypeOf((*IDataProcessor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
		},
		func() interface{} {
			return &jsiiProxy_IDataProcessor{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-kinesisfirehose-alpha.IDeliveryStream",
		reflect.TypeOf((*IDeliveryStream)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamArn", GoGetter: "DeliveryStreamArn"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamName", GoGetter: "DeliveryStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantPutRecords", GoMethod: "GrantPutRecords"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricBackupToS3Bytes", GoMethod: "MetricBackupToS3Bytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricBackupToS3DataFreshness", GoMethod: "MetricBackupToS3DataFreshness"},
			_jsii_.MemberMethod{JsiiMethod: "metricBackupToS3Records", GoMethod: "MetricBackupToS3Records"},
			_jsii_.MemberMethod{JsiiMethod: "metricIncomingBytes", GoMethod: "MetricIncomingBytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricIncomingRecords", GoMethod: "MetricIncomingRecords"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IDeliveryStream{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-kinesisfirehose-alpha.IDestination",
		reflect.TypeOf((*IDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IDestination{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-kinesisfirehose-alpha.LambdaFunctionProcessor",
		reflect.TypeOf((*LambdaFunctionProcessor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaFunctionProcessor{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDataProcessor)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-kinesisfirehose-alpha.StreamEncryption",
		reflect.TypeOf((*StreamEncryption)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			return &jsiiProxy_StreamEncryption{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-kinesisfirehose-alpha.StreamEncryptionType",
		reflect.TypeOf((*StreamEncryptionType)(nil)).Elem(),
		map[string]interface{}{
			"UNENCRYPTED": StreamEncryptionType_UNENCRYPTED,
			"CUSTOMER_MANAGED": StreamEncryptionType_CUSTOMER_MANAGED,
			"AWS_OWNED": StreamEncryptionType_AWS_OWNED,
		},
	)
}
