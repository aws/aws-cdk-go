package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Location where the Bedrock InvokeModel API response is written.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   model := bedrock.FoundationModel_FromFoundationModelId(this, jsii.String("Model"), bedrock.FoundationModelIdentifier_AMAZON_TITAN_TEXT_G1_EXPRESS_V1())
//
//   task := tasks.NewBedrockInvokeModel(this, jsii.String("Prompt Model"), &BedrockInvokeModelProps{
//   	Model: Model,
//   	Input: &BedrockInvokeModelInputProps{
//   		S3InputUri: sfn.JsonPath_StringAt(jsii.String("$.prompt")),
//   	},
//   	Output: &BedrockInvokeModelOutputProps{
//   		S3OutputUri: sfn.JsonPath_*StringAt(jsii.String("$.prompt")),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-bedrock.html
//
type BedrockInvokeModelOutputProps struct {
	// S3 object where the Bedrock InvokeModel API response is written.
	//
	// If you specify this field, the API response body is replaced with
	// a reference to the Amazon S3 location of the original output.
	// Default: - Response body is returned in the task result.
	//
	S3Location *awss3.Location `field:"optional" json:"s3Location" yaml:"s3Location"`
	// The destination location where the API response is written.
	//
	// This field can be used to specify s3 URI in the form of token.
	// Default: - The API response body is returned in the result.
	//
	S3OutputUri *string `field:"optional" json:"s3OutputUri" yaml:"s3OutputUri"`
}

