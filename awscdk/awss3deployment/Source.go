package awss3deployment

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
)

// Specifies bucket deployment source.
//
// Usage:
//
//     Source.bucket(bucket, key)
//     Source.asset('/local/path/to/directory')
//     Source.asset('/local/path/to/a/file.zip')
//     Source.data('hello/world/file.txt', 'Hello, world!')
//     Source.jsonData('config.json', { baz: topic.topicArn })
//     Source.yamlData('config.yaml', { baz: topic.topicArn })
//
// Example:
//   var websiteBucket bucket
//
//
//   deployment := s3deploy.NewBucketDeployment(this, jsii.String("DeployWebsite"), &BucketDeploymentProps{
//   	Sources: []iSource{
//   		s3deploy.Source_Asset(path.join(__dirname, jsii.String("my-website"))),
//   	},
//   	DestinationBucket: websiteBucket,
//   })
//
//   NewConstructThatReadsFromTheBucket(this, jsii.String("Consumer"), map[string]iBucket{
//   	// Use 'deployment.deployedBucket' instead of 'websiteBucket' here
//   	"bucket": deployment.deployedBucket,
//   })
//
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
func Source_Asset(path *string, options *awss3assets.AssetOptions) ISource {
	_init_.Initialize()

	if err := validateSource_AssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns ISource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.Source",
		"asset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Uses a .zip file stored in an S3 bucket as the source for the destination bucket contents.
//
// Make sure you trust the producer of the archive.
//
// If the `bucket` parameter is an "out-of-app" reference "imported" via static methods such
// as `s3.Bucket.fromBucketName`, be cautious about the bucket's encryption key. In general,
// CDK does not query for additional properties of imported constructs at synthesis time.
// For example, for a bucket created from `s3.Bucket.fromBucketName`, CDK does not know
// its `IBucket.encryptionKey` property, and therefore will NOT give KMS permissions to the
// Lambda execution role of the `BucketDeployment` construct. If you want the
// `kms:Decrypt` and `kms:DescribeKey` permissions on the bucket's encryption key
// to be added automatically, reference the imported bucket via `s3.Bucket.fromBucketAttributes`
// and pass in the `encryptionKey` attribute explicitly.
//
// Example:
//   var destinationBucket bucket
//
//   sourceBucket := s3.bucket_FromBucketAttributes(this, jsii.String("SourceBucket"), &BucketAttributes{
//   	BucketArn: jsii.String("arn:aws:s3:::my-source-bucket-name"),
//   	EncryptionKey: kms.Key_FromKeyArn(this, jsii.String("SourceBucketEncryptionKey"), jsii.String("arn:aws:kms:us-east-1:123456789012:key/<key-id>")),
//   })
//   deployment := s3deploy.NewBucketDeployment(this, jsii.String("DeployFiles"), &BucketDeploymentProps{
//   	Sources: []iSource{
//   		s3deploy.Source_Bucket(sourceBucket, jsii.String("source.zip")),
//   	},
//   	DestinationBucket: DestinationBucket,
//   })
//
func Source_Bucket(bucket awss3.IBucket, zipObjectKey *string) ISource {
	_init_.Initialize()

	if err := validateSource_BucketParameters(bucket, zipObjectKey); err != nil {
		panic(err)
	}
	var returns ISource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.Source",
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
// To store YAML content use `Source.yamlData()`.
func Source_Data(objectKey *string, data *string) ISource {
	_init_.Initialize()

	if err := validateSource_DataParameters(objectKey, data); err != nil {
		panic(err)
	}
	var returns ISource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.Source",
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
func Source_JsonData(objectKey *string, obj interface{}) ISource {
	_init_.Initialize()

	if err := validateSource_JsonDataParameters(objectKey, obj); err != nil {
		panic(err)
	}
	var returns ISource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.Source",
		"jsonData",
		[]interface{}{objectKey, obj},
		&returns,
	)

	return returns
}

// Deploys an object with the specified JSON object formatted as YAML into the bucket.
//
// The object can include deploy-time values (such as `snsTopic.topicArn`) that
// will get resolved only during deployment.
func Source_YamlData(objectKey *string, obj interface{}) ISource {
	_init_.Initialize()

	if err := validateSource_YamlDataParameters(objectKey, obj); err != nil {
		panic(err)
	}
	var returns ISource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.Source",
		"yamlData",
		[]interface{}{objectKey, obj},
		&returns,
	)

	return returns
}

