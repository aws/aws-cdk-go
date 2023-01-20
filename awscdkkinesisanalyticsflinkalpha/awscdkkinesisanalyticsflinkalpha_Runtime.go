// A CDK Construct Library for Kinesis Analytics Flink applications
package awscdkkinesisanalyticsflinkalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkkinesisanalyticsflinkalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Available Flink runtimes for Kinesis Analytics.
//
// Example:
//   import path "github.com/aws-samples/dummy/path"
//   import cloudwatch "github.com/aws/aws-cdk-go/awscdk"
//   import core "github.com/aws/aws-cdk-go/awscdk"
//   import flink "github.com/aws/aws-cdk-go/awscdkkinesisanalyticsflinkalpha"
//
//   app := core.NewApp()
//   stack := core.NewStack(app, jsii.String("FlinkAppTest"))
//
//   flinkApp := flink.NewApplication(stack, jsii.String("App"), &applicationProps{
//   	code: flink.applicationCode.fromAsset(path.join(__dirname, jsii.String("code-asset"))),
//   	runtime: flink.runtime_FLINK_1_11(),
//   })
//
//   cloudwatch.NewAlarm(stack, jsii.String("Alarm"), &alarmProps{
//   	metric: flinkApp.metricFullRestarts(),
//   	evaluationPeriods: jsii.Number(1),
//   	threshold: jsii.Number(3),
//   })
//
//   app.synth()
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

