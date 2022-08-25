package awslambda


// A ZIP archive that contains the contents of an [Lambda layer](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contentProperty := &contentProperty{
//   	s3Bucket: jsii.String("s3Bucket"),
//   	s3Key: jsii.String("s3Key"),
//
//   	// the properties below are optional
//   	s3ObjectVersion: jsii.String("s3ObjectVersion"),
//   }
//
type CfnLayerVersion_ContentProperty struct {
	// The Amazon S3 bucket of the layer archive.
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// The Amazon S3 key of the layer archive.
	S3Key *string `field:"required" json:"s3Key" yaml:"s3Key"`
	// For versioned objects, the version of the layer archive object to use.
	S3ObjectVersion *string `field:"optional" json:"s3ObjectVersion" yaml:"s3ObjectVersion"`
}

