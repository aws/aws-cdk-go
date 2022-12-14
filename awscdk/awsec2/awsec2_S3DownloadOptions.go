package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Options when downloading files from S3.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var instance instance
//
//
//   asset := awscdk.NewAsset(this, jsii.String("Asset"), &assetProps{
//   	path: jsii.String("./configure.sh"),
//   })
//
//   localPath := instance.userData.addS3DownloadCommand(&s3DownloadOptions{
//   	bucket: asset.bucket,
//   	bucketKey: asset.s3ObjectKey,
//   	region: jsii.String("us-east-1"),
//   })
//   instance.userData.addExecuteFileCommand(&executeFileOptions{
//   	filePath: localPath,
//   	arguments: jsii.String("--verbose -y"),
//   })
//   asset.grantRead(instance.role)
//
type S3DownloadOptions struct {
	// Name of the S3 bucket to download from.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The key of the file to download.
	BucketKey *string `field:"required" json:"bucketKey" yaml:"bucketKey"`
	// The name of the local file.
	LocalFile *string `field:"optional" json:"localFile" yaml:"localFile"`
	// The region of the S3 Bucket (needed for access via VPC Gateway).
	Region *string `field:"optional" json:"region" yaml:"region"`
}

