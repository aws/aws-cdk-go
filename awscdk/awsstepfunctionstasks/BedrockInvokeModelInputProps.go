package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Location to retrieve the input data, prior to calling Bedrock InvokeModel.
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
type BedrockInvokeModelInputProps struct {
	// The source location where the API response is written.
	//
	// This field can be used to specify s3 URI in the form of token.
	// Default: - The API response body is returned in the result.
	//
	S3InputUri *string `field:"optional" json:"s3InputUri" yaml:"s3InputUri"`
	// S3 object to retrieve the input data from.
	//
	// If the S3 location is not set, then the Body must be set.
	// Default: - Input data is retrieved from the `body` field.
	//
	S3Location *awss3.Location `field:"optional" json:"s3Location" yaml:"s3Location"`
}

