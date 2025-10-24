package awscdkkinesisanalyticsflinkalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkkinesisanalyticsflinkalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Available Flink runtimes for Kinesis Analytics.
//
// Example:
//   var bucket Bucket
//
//   flinkApp := flink.NewApplication(this, jsii.String("Application"), &ApplicationProps{
//   	Code: flink.ApplicationCode_FromBucket(bucket, jsii.String("my-app.jar")),
//   	Runtime: flink.Runtime_FLINK_1_20(),
//   	CheckpointingEnabled: jsii.Boolean(true),
//   	 // default is true
//   	CheckpointInterval: awscdk.Duration_Seconds(jsii.Number(30)),
//   	 // default is 1 minute
//   	MinPauseBetweenCheckpoints: awscdk.Duration_*Seconds(jsii.Number(10)),
//   	 // default is 5 seconds
//   	LogLevel: flink.LogLevel_ERROR,
//   	 // default is INFO
//   	MetricsLevel: flink.MetricsLevel_PARALLELISM,
//   	 // default is APPLICATION
//   	AutoScalingEnabled: jsii.Boolean(false),
//   	 // default is true
//   	Parallelism: jsii.Number(32),
//   	 // default is 1
//   	ParallelismPerKpu: jsii.Number(2),
//   	 // default is 1
//   	SnapshotsEnabled: jsii.Boolean(false),
//   	 // default is true
//   	LogGroup: logs.NewLogGroup(this, jsii.String("LogGroup")),
//   })
//
// Experimental.
type Runtime interface {
	// The Cfn string that represents a version of Flink.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Runtime
type jsiiProxy_Runtime struct {
	_ byte // padding
}

func (j *jsiiProxy_Runtime) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Create a new Runtime with with an arbitrary Flink version string.
// Experimental.
func Runtime_Of(value *string) Runtime {
	_init_.Initialize()

	if err := validateRuntime_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Runtime

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.Runtime",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Runtime_FLINK_1_11() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.Runtime",
		"FLINK_1_11",
		&returns,
	)
	return returns
}

func Runtime_FLINK_1_13() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.Runtime",
		"FLINK_1_13",
		&returns,
	)
	return returns
}

func Runtime_FLINK_1_15() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.Runtime",
		"FLINK_1_15",
		&returns,
	)
	return returns
}

func Runtime_FLINK_1_18() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.Runtime",
		"FLINK_1_18",
		&returns,
	)
	return returns
}

func Runtime_FLINK_1_19() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.Runtime",
		"FLINK_1_19",
		&returns,
	)
	return returns
}

func Runtime_FLINK_1_20() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.Runtime",
		"FLINK_1_20",
		&returns,
	)
	return returns
}

func Runtime_FLINK_1_6() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.Runtime",
		"FLINK_1_6",
		&returns,
	)
	return returns
}

func Runtime_FLINK_1_8() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.Runtime",
		"FLINK_1_8",
		&returns,
	)
	return returns
}

func Runtime_SQL_1_0() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.Runtime",
		"SQL_1_0",
		&returns,
	)
	return returns
}

func Runtime_ZEPPELIN_FLINK_1_0() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.Runtime",
		"ZEPPELIN_FLINK_1_0",
		&returns,
	)
	return returns
}

func Runtime_ZEPPELIN_FLINK_2_0() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.Runtime",
		"ZEPPELIN_FLINK_2_0",
		&returns,
	)
	return returns
}

func Runtime_ZEPPELIN_FLINK_3_0() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.Runtime",
		"ZEPPELIN_FLINK_3_0",
		&returns,
	)
	return returns
}

