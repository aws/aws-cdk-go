package awskinesisanalyticsflink

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/aws-cdk-go/awscdk/awss3assets"
)

// Code configuration providing the location to a Flink application JAR file.
//
// Example:
//   import path "github.com/aws-samples/dummy/path"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import cloudwatch "github.com/aws/aws-cdk-go/awscdk"
//
//   app := core.NewApp()
//   stack := core.NewStack(app, jsii.String("FlinkAppTest"))
//
//   flinkApp := flink.NewApplication(stack, jsii.String("App"), &ApplicationProps{
//   	Code: flink.ApplicationCode_FromAsset(path.join(__dirname, jsii.String("code-asset"))),
//   	Runtime: flink.Runtime_FLINK_1_11(),
//   })
//
//   cloudwatch.NewAlarm(stack, jsii.String("Alarm"), &AlarmProps{
//   	Metric: flinkApp.metricFullRestarts(),
//   	EvaluationPeriods: jsii.Number(1),
//   	Threshold: jsii.Number(3),
//   })
//
//   app.Synth()
//
// Experimental.
type ApplicationCode interface {
	// A method to lazily bind asset resources to the parent FlinkApplication.
	// Experimental.
	Bind(scope awscdk.Construct) *ApplicationCodeConfig
}

// The jsii proxy struct for ApplicationCode
type jsiiProxy_ApplicationCode struct {
	_ byte // padding
}

// Experimental.
func NewApplicationCode_Override(a ApplicationCode) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_kinesisanalytics_flink.ApplicationCode",
		nil, // no parameters
		a,
	)
}

// Reference code from a local directory containing a Flink JAR file.
// Experimental.
func ApplicationCode_FromAsset(path *string, options *awss3assets.AssetOptions) ApplicationCode {
	_init_.Initialize()

	if err := validateApplicationCode_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns ApplicationCode

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisanalytics_flink.ApplicationCode",
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

	if err := validateApplicationCode_FromBucketParameters(bucket, fileKey); err != nil {
		panic(err)
	}
	var returns ApplicationCode

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisanalytics_flink.ApplicationCode",
		"fromBucket",
		[]interface{}{bucket, fileKey, objectVersion},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationCode) Bind(scope awscdk.Construct) *ApplicationCodeConfig {
	if err := a.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *ApplicationCodeConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

