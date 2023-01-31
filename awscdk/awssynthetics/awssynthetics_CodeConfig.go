package awssynthetics

import (
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

// Configuration of the code class.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeConfig := &codeConfig{
//   	inlineCode: jsii.String("inlineCode"),
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
	// Inline code (mutually exclusive with `s3Location`).
	// Experimental.
	InlineCode *string `field:"optional" json:"inlineCode" yaml:"inlineCode"`
	// The location of the code in S3 (mutually exclusive with `inlineCode`).
	// Experimental.
	S3Location *awss3.Location `field:"optional" json:"s3Location" yaml:"s3Location"`
}

