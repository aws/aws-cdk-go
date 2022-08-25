// A CDK Construct Library for Kinesis Analytics Flink applications
package awscdkkinesisanalyticsflinkalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkkinesisanalyticsflinkalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// Code configuration providing the location to a Flink application JAR file.
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
type ApplicationCode interface {
	// A method to lazily bind asset resources to the parent FlinkApplication.
	// Experimental.
	Bind(scope constructs.Construct) *ApplicationCodeConfig
}

// The jsii proxy struct for ApplicationCode
type jsiiProxy_ApplicationCode struct {
	_ byte // padding
}

// Experimental.
func NewApplicationCode_Override(a ApplicationCode) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.ApplicationCode",
		nil, // no parameters
		a,
	)
}

// Reference code from a local directory containing a Flink JAR file.
// Experimental.
func ApplicationCode_FromAsset(path *string, options *awss3assets.AssetOptions) ApplicationCode {
	_init_.Initialize()

	var returns ApplicationCode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.ApplicationCode",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Reference code from an S3 bucket.
// Experimental.
func ApplicationCode_FromBucket(bucket awss3.IBucket, fileKey *string, objectVersion *string) ApplicationCode {
	_init_.Initialize()

	var returns ApplicationCode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.ApplicationCode",
		"fromBucket",
		[]interface{}{bucket, fileKey, objectVersion},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationCode) Bind(scope constructs.Construct) *ApplicationCodeConfig {
	var returns *ApplicationCodeConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

