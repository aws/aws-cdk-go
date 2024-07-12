package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Location to retrieve the input data, prior to calling Bedrock InvokeModel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bedrockInvokeModelInputProps := &BedrockInvokeModelInputProps{
//   	S3Location: &Location{
//   		BucketName: jsii.String("bucketName"),
//   		ObjectKey: jsii.String("objectKey"),
//
//   		// the properties below are optional
//   		ObjectVersion: jsii.String("objectVersion"),
//   	},
//   }
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-bedrock.html
//
type BedrockInvokeModelInputProps struct {
	// S3 object to retrieve the input data from.
	//
	// If the S3 location is not set, then the Body must be set.
	// Default: - Input data is retrieved from the `body` field.
	//
	S3Location *awss3.Location `field:"optional" json:"s3Location" yaml:"s3Location"`
}

