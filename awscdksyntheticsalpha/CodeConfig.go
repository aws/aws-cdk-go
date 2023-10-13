package awscdksyntheticsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Configuration of the code class.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import synthetics_alpha "github.com/aws/aws-cdk-go/awscdksyntheticsalpha"
//
//   codeConfig := &CodeConfig{
//   	InlineCode: jsii.String("inlineCode"),
//   	S3Location: &Location{
//   		BucketName: jsii.String("bucketName"),
//   		ObjectKey: jsii.String("objectKey"),
//
//   		// the properties below are optional
//   		ObjectVersion: jsii.String("objectVersion"),
//   	},
//   }
//
// Deprecated.
type CodeConfig struct {
	// Inline code (mutually exclusive with `s3Location`).
	// Default: - none.
	//
	// Deprecated.
	InlineCode *string `field:"optional" json:"inlineCode" yaml:"inlineCode"`
	// The location of the code in S3 (mutually exclusive with `inlineCode`).
	// Default: - none.
	//
	// Deprecated.
	S3Location *awss3.Location `field:"optional" json:"s3Location" yaml:"s3Location"`
}

