package awslambda


// S3 location for WebFunction code.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webFunctionS3Location := &WebFunctionS3Location{
//   	BucketName: jsii.String("bucketName"),
//   	ObjectKey: jsii.String("objectKey"),
//
//   	// the properties below are optional
//   	ObjectVersion: jsii.String("objectVersion"),
//   }
//
type WebFunctionS3Location struct {
	// The S3 bucket name.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The S3 object key.
	ObjectKey *string `field:"required" json:"objectKey" yaml:"objectKey"`
	// The S3 object version (optional).
	// Default: - the version that is latest when the revision is created.
	//
	ObjectVersion *string `field:"optional" json:"objectVersion" yaml:"objectVersion"`
}

