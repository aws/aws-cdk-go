package awss3

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The key format for the log object.
//
// Example:
//   accessLogsBucket := s3.NewBucket(this, jsii.String("AccessLogsBucket"))
//
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
//   	ServerAccessLogsBucket: accessLogsBucket,
//   	ServerAccessLogsPrefix: jsii.String("logs"),
//   	// You can use a simple prefix with `TargetObjectKeyFormat.simplePrefix()`, but it is the same even if you do not specify `targetObjectKeyFormat` property.
//   	TargetObjectKeyFormat: s3.TargetObjectKeyFormat_SimplePrefix(),
//   })
//
type TargetObjectKeyFormat interface {
}

// The jsii proxy struct for TargetObjectKeyFormat
type jsiiProxy_TargetObjectKeyFormat struct {
	_ byte // padding
}

func NewTargetObjectKeyFormat_Override(t TargetObjectKeyFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.TargetObjectKeyFormat",
		nil, // no parameters
		t,
	)
}

// Use partitioned prefix for log objects. If you do not specify the dateSource argument, the default is EventTime.
//
// The partitioned prefix format as follow:
// [DestinationPrefix][SourceAccountId]/​[SourceRegion]/​[SourceBucket]/​[YYYY]/​[MM]/​[DD]/​[YYYY]-[MM]-[DD]-[hh]-[mm]-[ss]-[UniqueString].
func TargetObjectKeyFormat_PartitionedPrefix(dateSource PartitionDateSource) TargetObjectKeyFormat {
	_init_.Initialize()

	var returns TargetObjectKeyFormat

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.TargetObjectKeyFormat",
		"partitionedPrefix",
		[]interface{}{dateSource},
		&returns,
	)

	return returns
}

// Use the simple prefix for log objects.
//
// The simple prefix format as follow:
// [DestinationPrefix][YYYY]-[MM]-[DD]-[hh]-[mm]-[ss]-[UniqueString].
func TargetObjectKeyFormat_SimplePrefix() TargetObjectKeyFormat {
	_init_.Initialize()

	var returns TargetObjectKeyFormat

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.TargetObjectKeyFormat",
		"simplePrefix",
		nil, // no parameters
		&returns,
	)

	return returns
}

