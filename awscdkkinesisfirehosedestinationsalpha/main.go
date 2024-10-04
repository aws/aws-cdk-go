// CDK Destinations Constructs for AWS Kinesis Firehose
package awscdkkinesisfirehosedestinationsalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-kinesisfirehose-destinations-alpha.BackupMode",
		reflect.TypeOf((*BackupMode)(nil)).Elem(),
		map[string]interface{}{
			"ALL": BackupMode_ALL,
			"FAILED": BackupMode_FAILED,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-kinesisfirehose-destinations-alpha.CommonDestinationProps",
		reflect.TypeOf((*CommonDestinationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-kinesisfirehose-destinations-alpha.CommonDestinationS3Props",
		reflect.TypeOf((*CommonDestinationS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-kinesisfirehose-destinations-alpha.Compression",
		reflect.TypeOf((*Compression)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_Compression{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-kinesisfirehose-destinations-alpha.DestinationS3BackupProps",
		reflect.TypeOf((*DestinationS3BackupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-kinesisfirehose-destinations-alpha.DisableLogging",
		reflect.TypeOf((*DisableLogging)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "logging", GoGetter: "Logging"},
		},
		func() interface{} {
			j := jsiiProxy_DisableLogging{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ILoggingConfig)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-kinesisfirehose-destinations-alpha.EnableLogging",
		reflect.TypeOf((*EnableLogging)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "logging", GoGetter: "Logging"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
		},
		func() interface{} {
			j := jsiiProxy_EnableLogging{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ILoggingConfig)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-kinesisfirehose-destinations-alpha.ILoggingConfig",
		reflect.TypeOf((*ILoggingConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "logging", GoGetter: "Logging"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
		},
		func() interface{} {
			return &jsiiProxy_ILoggingConfig{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-kinesisfirehose-destinations-alpha.S3Bucket",
		reflect.TypeOf((*S3Bucket)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_S3Bucket{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkkinesisfirehosealphaIDestination)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-kinesisfirehose-destinations-alpha.S3BucketProps",
		reflect.TypeOf((*S3BucketProps)(nil)).Elem(),
	)
}
