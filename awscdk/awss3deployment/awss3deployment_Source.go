package awss3deployment

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/aws-cdk-go/awscdk/awss3assets"
)

// Specifies bucket deployment source.
//
// Usage:
//
//      Source.bucket(bucket, key)
//      Source.asset('/local/path/to/directory')
//      Source.asset('/local/path/to/a/file.zip')
//      Source.data('hello/world/file.txt', 'Hello, world!')
//      Source.data('config.json', { baz: topic.topicArn })
//
// Example:
//   var websiteBucket bucket
//
//
//   deployment := s3deploy.NewBucketDeployment(this, jsii.String("DeployWebsite"), &bucketDeploymentProps{
//   	sources: []iSource{
//   		s3deploy.source.asset(path.join(__dirname, jsii.String("my-website"))),
//   	},
//   	destinationBucket: websiteBucket,
//   })
//
//   NewConstructThatReadsFromTheBucket(this, jsii.String("Consumer"), map[string]iBucket{
//   	// Use 'deployment.deployedBucket' instead of 'websiteBucket' here
//   	"bucket": deployment.deployedBucket,
//   })
//
// Experimental.
type Source interface {
}

// The jsii proxy struct for Source
type jsiiProxy_Source struct {
	_ byte // padding
}

// Uses a local asset as the deployment source.
//
// If the local asset is a .zip archive, make sure you trust the
// producer of the archive.
// Experimental.
func Source_Asset(path *string, options *awss3assets.AssetOptions) ISource {
	_init_.Initialize()

	if err := validateSource_AssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns ISource

	_jsii_.StaticInvoke(
		"monocdk.aws_s3_deployment.Source",
		"asset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Uses a .zip file stored in an S3 bucket as the source for the destination bucket contents.
//
// Make sure you trust the producer of the archive.
// Experimental.
func Source_Bucket(bucket awss3.IBucket, zipObjectKey *string) ISource {
	_init_.Initialize()

	if err := validateSource_BucketParameters(bucket, zipObjectKey); err != nil {
		panic(err)
	}
	var returns ISource

	_jsii_.StaticInvoke(
		"monocdk.aws_s3_deployment.Source",
		"bucket",
		[]interface{}{bucket, zipObjectKey},
		&returns,
	)

	return returns
}

// Deploys an object with the specified string contents into the bucket.
//
// The
// content can include deploy-time values (such as `snsTopic.topicArn`) that
// will get resolved only during deployment.
//
// To store a JSON object use `Source.jsonData()`.
// Experimental.
func Source_Data(objectKey *string, data *string) ISource {
	_init_.Initialize()

	if err := validateSource_DataParameters(objectKey, data); err != nil {
		panic(err)
	}
	var returns ISource

	_jsii_.StaticInvoke(
		"monocdk.aws_s3_deployment.Source",
		"data",
		[]interface{}{objectKey, data},
		&returns,
	)

	return returns
}

// Deploys an object with the specified JSON object into the bucket.
//
// The
// object can include deploy-time values (such as `snsTopic.topicArn`) that
// will get resolved only during deployment.
// Experimental.
func Source_JsonData(objectKey *string, obj interface{}) ISource {
	_init_.Initialize()

	if err := validateSource_JsonDataParameters(objectKey, obj); err != nil {
		panic(err)
	}
	var returns ISource

	_jsii_.StaticInvoke(
		"monocdk.aws_s3_deployment.Source",
		"jsonData",
		[]interface{}{objectKey, obj},
		&returns,
	)

	return returns
}

