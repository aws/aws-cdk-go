package awscdk


// The location of the published file asset.
//
// This is where the asset
// can be consumed at runtime.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   fileAssetLocation := &FileAssetLocation{
//   	BucketName: jsii.String("bucketName"),
//   	HttpUrl: jsii.String("httpUrl"),
//   	ObjectKey: jsii.String("objectKey"),
//   	S3ObjectUrl: jsii.String("s3ObjectUrl"),
//
//   	// the properties below are optional
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	S3ObjectUrlWithPlaceholders: jsii.String("s3ObjectUrlWithPlaceholders"),
//   }
//
type FileAssetLocation struct {
	// The name of the Amazon S3 bucket.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The HTTP URL of this asset on Amazon S3.
	//
	// This value suitable for inclusion in a CloudFormation template, and
	// may be an encoded token.
	//
	// Example value: `https://s3-us-east-1.amazonaws.com/mybucket/myobject`
	HttpUrl *string `field:"required" json:"httpUrl" yaml:"httpUrl"`
	// The Amazon S3 object key.
	ObjectKey *string `field:"required" json:"objectKey" yaml:"objectKey"`
	// The S3 URL of this asset on Amazon S3.
	//
	// This value suitable for inclusion in a CloudFormation template, and
	// may be an encoded token.
	//
	// Example value: `s3://mybucket/myobject`.
	S3ObjectUrl *string `field:"required" json:"s3ObjectUrl" yaml:"s3ObjectUrl"`
	// The ARN of the KMS key used to encrypt the file asset bucket, if any.
	//
	// The CDK bootstrap stack comes with a key policy that does not require
	// setting this property, so you only need to set this property if you
	// have customized the bootstrap stack to require it.
	// Default: - Asset bucket is not encrypted, or decryption permissions are
	// defined by a Key Policy.
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Like `s3ObjectUrl`, but not suitable for CloudFormation consumption.
	//
	// If there are placeholders in the S3 URL, they will be returned un-replaced
	// and un-evaluated.
	// Default: - This feature cannot be used.
	//
	S3ObjectUrlWithPlaceholders *string `field:"optional" json:"s3ObjectUrlWithPlaceholders" yaml:"s3ObjectUrlWithPlaceholders"`
}

