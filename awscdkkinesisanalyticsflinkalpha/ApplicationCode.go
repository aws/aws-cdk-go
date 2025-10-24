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

	if err := validateApplicationCode_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
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

	if err := validateApplicationCode_FromBucketParameters(bucket, fileKey); err != nil {
		panic(err)
	}
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

