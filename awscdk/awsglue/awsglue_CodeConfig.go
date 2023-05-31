package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

// Result of binding `Code` into a `Job`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeConfig := &CodeConfig{
//   	S3Location: &Location{
//   		BucketName: jsii.String("bucketName"),
//   		ObjectKey: jsii.String("objectKey"),
//
//   		// the properties below are optional
//   		ObjectVersion: jsii.String("objectVersion"),
//   	},
//   }
//
// Experimental.
type CodeConfig struct {
	// The location of the code in S3.
	// Experimental.
	S3Location *awss3.Location `field:"required" json:"s3Location" yaml:"s3Location"`
}

