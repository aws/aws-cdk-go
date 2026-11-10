package awslambda


// WebFunction code configuration result.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webFunctionCodeConfig := &WebFunctionCodeConfig{
//   	S3Location: &WebFunctionS3Location{
//   		BucketName: jsii.String("bucketName"),
//   		ObjectKey: jsii.String("objectKey"),
//
//   		// the properties below are optional
//   		ObjectVersion: jsii.String("objectVersion"),
//   	},
//   }
//
type WebFunctionCodeConfig struct {
	// The S3 location of the code.
	S3Location *WebFunctionS3Location `field:"required" json:"s3Location" yaml:"s3Location"`
}

