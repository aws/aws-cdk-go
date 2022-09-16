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
//   codeConfig := &codeConfig{
//   	s3Location: &location{
//   		bucketName: jsii.String("bucketName"),
//   		objectKey: jsii.String("objectKey"),
//
//   		// the properties below are optional
//   		objectVersion: jsii.String("objectVersion"),
//   	},
//   }
//
// Experimental.
type CodeConfig struct {
	// The location of the code in S3.
	// Experimental.
	S3Location *awss3.Location `field:"required" json:"s3Location" yaml:"s3Location"`
}

