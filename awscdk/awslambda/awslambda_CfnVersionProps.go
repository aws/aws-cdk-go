package awslambda


// Properties for defining a `CfnVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVersionProps := &cfnVersionProps{
//   	functionName: jsii.String("functionName"),
//
//   	// the properties below are optional
//   	codeSha256: jsii.String("codeSha256"),
//   	description: jsii.String("description"),
//   	provisionedConcurrencyConfig: &provisionedConcurrencyConfigurationProperty{
//   		provisionedConcurrentExecutions: jsii.Number(123),
//   	},
//   }
//
type CfnVersionProps struct {
	// The name of the Lambda function.
	//
	// **Name formats** - *Function name* - `MyFunction` .
	// - *Function ARN* - `arn:aws:lambda:us-west-2:123456789012:function:MyFunction` .
	// - *Partial ARN* - `123456789012:function:MyFunction` .
	//
	// The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// Only publish a version if the hash value matches the value that's specified.
	//
	// Use this option to avoid publishing a version if the function code has changed since you last updated it. Updates are not supported for this property.
	CodeSha256 *string `field:"optional" json:"codeSha256" yaml:"codeSha256"`
	// A description for the version to override the description in the function configuration.
	//
	// Updates are not supported for this property.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies a provisioned concurrency configuration for a function's version.
	//
	// Updates are not supported for this property.
	ProvisionedConcurrencyConfig interface{} `field:"optional" json:"provisionedConcurrencyConfig" yaml:"provisionedConcurrencyConfig"`
}

