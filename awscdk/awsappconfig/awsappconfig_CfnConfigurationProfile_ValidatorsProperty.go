package awsappconfig


// A validator provides a syntactic or semantic check to ensure the configuration that you want to deploy functions as intended.
//
// To validate your application configuration data, you provide a schema or an AWS Lambda function that runs against the configuration. The configuration deployment or update can only proceed when the configuration data is valid.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   validatorsProperty := &validatorsProperty{
//   	content: jsii.String("content"),
//   	type: jsii.String("type"),
//   }
//
type CfnConfigurationProfile_ValidatorsProperty struct {
	// Either the JSON Schema content or the Amazon Resource Name (ARN) of an Lambda function.
	Content *string `field:"optional" json:"content" yaml:"content"`
	// AWS AppConfig supports validators of type `JSON_SCHEMA` and `LAMBDA`.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

