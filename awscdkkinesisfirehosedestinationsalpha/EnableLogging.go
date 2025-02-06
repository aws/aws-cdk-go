package awscdkkinesisfirehosedestinationsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkkinesisfirehosedestinationsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Enables logging for error logs with an optional custom CloudWatch log group.
//
// When this class is used, logging is enabled (`logging: true`) and
// you can optionally provide a CloudWatch log group for storing the error logs.
//
// If no log group is provided, a default one will be created automatically.
//
// Example:
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//   var bucket bucket
//
//
//   logGroup := logs.NewLogGroup(this, jsii.String("Log Group"))
//   destination := destinations.NewS3Bucket(bucket, &S3BucketProps{
//   	LoggingConfig: destinations.NewEnableLogging(logGroup),
//   })
//
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &DeliveryStreamProps{
//   	Destination: destination,
//   })
//
// Deprecated.
type EnableLogging interface {
	ILoggingConfig
	// If true, log errors when data transformation or data delivery fails.
	//
	// `true` when using `EnableLogging`, `false` when using `DisableLogging`.
	// Deprecated.
	Logging() *bool
	// The CloudWatch log group where log streams will be created to hold error logs.
	// Deprecated.
	LogGroup() awslogs.ILogGroup
}

// The jsii proxy struct for EnableLogging
type jsiiProxy_EnableLogging struct {
	jsiiProxy_ILoggingConfig
}

func (j *jsiiProxy_EnableLogging) Logging() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnableLogging) LogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}


// Deprecated.
func NewEnableLogging(logGroup awslogs.ILogGroup) EnableLogging {
	_init_.Initialize()

	j := jsiiProxy_EnableLogging{}

	_jsii_.Create(
		"@aws-cdk/aws-kinesisfirehose-destinations-alpha.EnableLogging",
		[]interface{}{logGroup},
		&j,
	)

	return &j
}

// Deprecated.
func NewEnableLogging_Override(e EnableLogging, logGroup awslogs.ILogGroup) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-kinesisfirehose-destinations-alpha.EnableLogging",
		[]interface{}{logGroup},
		e,
	)
}

