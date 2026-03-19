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
//   	PropertyGroups: map[string]map[string]*string{
//   		"FlinkApplicationProperties": map[string]*string{
//   			"inputStreamName": jsii.String("my-input-kinesis-stream"),
//   			"outputStreamName": jsii.String("my-output-kinesis-stream"),
//   		},
//   	},
//   	// ...
//   	Runtime: flink.Runtime_FLINK_1_20(),
//   	Code: flink.ApplicationCode_FromBucket(bucket, jsii.String("my-app.jar")),
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

