package awskinesisfirehosedestinations

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose_destinations.CommonDestinationProps",
		reflect.TypeOf((*CommonDestinationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_kinesisfirehose_destinations.S3Bucket",
		reflect.TypeOf((*S3Bucket)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_S3Bucket{}
			_jsii_.InitJsiiProxy(&j.Type__awskinesisfirehoseIDestination)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose_destinations.S3BucketProps",
		reflect.TypeOf((*S3BucketProps)(nil)).Elem(),
	)
}
