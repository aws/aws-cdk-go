package awscodecommit


// Information about the Amazon S3 bucket that contains the code that will be committed to the new repository.
//
// Changes to this property are ignored after initial resource creation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3Property := &s3Property{
//   	bucket: jsii.String("bucket"),
//   	key: jsii.String("key"),
//
//   	// the properties below are optional
//   	objectVersion: jsii.String("objectVersion"),
//   }
//
type CfnRepository_S3Property struct {
	// The name of the Amazon S3 bucket that contains the ZIP file with the content that will be committed to the new repository.
	//
	// This can be specified using the name of the bucket in the AWS account . Changes to this property are ignored after initial resource creation.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The key to use for accessing the Amazon S3 bucket.
	//
	// Changes to this property are ignored after initial resource creation. For more information, see [Creating object key names](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-keys.html) and [Uploading objects](https://docs.aws.amazon.com/AmazonS3/latest/userguide/upload-objects.html) in the Amazon S3 User Guide.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The object version of the ZIP file, if versioning is enabled for the Amazon S3 bucket.
	//
	// Changes to this property are ignored after initial resource creation.
	ObjectVersion *string `field:"optional" json:"objectVersion" yaml:"objectVersion"`
}

