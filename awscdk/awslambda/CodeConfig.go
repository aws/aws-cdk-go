package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Result of binding `Code` into a `Function`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeConfig := &CodeConfig{
//   	Image: &CodeImageConfig{
//   		ImageUri: jsii.String("imageUri"),
//
//   		// the properties below are optional
//   		Cmd: []*string{
//   			jsii.String("cmd"),
//   		},
//   		Entrypoint: []*string{
//   			jsii.String("entrypoint"),
//   		},
//   		WorkingDirectory: jsii.String("workingDirectory"),
//   	},
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
type CodeConfig struct {
	// Docker image configuration (mutually exclusive with `s3Location` and `inlineCode`).
	// Default: - code is not an ECR container image.
	//
	Image *CodeImageConfig `field:"optional" json:"image" yaml:"image"`
	// Inline code (mutually exclusive with `s3Location` and `image`).
	// Default: - code is not inline code.
	//
	InlineCode *string `field:"optional" json:"inlineCode" yaml:"inlineCode"`
	// The location of the code in S3 (mutually exclusive with `inlineCode` and `image`).
	// Default: - code is not an s3 location.
	//
	S3Location *awss3.Location `field:"optional" json:"s3Location" yaml:"s3Location"`
}

