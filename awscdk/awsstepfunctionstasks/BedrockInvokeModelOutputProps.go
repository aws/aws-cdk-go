package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Location where the Bedrock InvokeModel API response is written.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bedrockInvokeModelOutputProps := &BedrockInvokeModelOutputProps{
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
type BedrockInvokeModelOutputProps struct {
	// S3 object where the Bedrock InvokeModel API response is written.
	//
	// If you specify this field, the API response body is replaced with
	// a reference to the Amazon S3 location of the original output.
	// Default: Response body is returned in the task result.
	//
	S3Location *awss3.Location `field:"optional" json:"s3Location" yaml:"s3Location"`
}

