package awslambda


// Properties for defining a `CfnVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVersionProps := &CfnVersionProps{
//   	FunctionName: jsii.String("functionName"),
//
//   	// the properties below are optional
//   	CodeSha256: jsii.String("codeSha256"),
//   	Description: jsii.String("description"),
//   	ProvisionedConcurrencyConfig: &ProvisionedConcurrencyConfigurationProperty{
//   		ProvisionedConcurrentExecutions: jsii.Number(123),
//   	},
//   	RuntimePolicy: &RuntimePolicyProperty{
//   		UpdateRuntimeOn: jsii.String("updateRuntimeOn"),
//
//   		// the properties below are optional
//   		RuntimeVersionArn: jsii.String("runtimeVersionArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-version.html
//
type CfnVersionProps struct {
	// The name or ARN of the Lambda function.
	//
	// **Name formats** - *Function name* - `MyFunction` .
	// - *Function ARN* - `arn:aws:lambda:us-west-2:123456789012:function:MyFunction` .
	// - *Partial ARN* - `123456789012:function:MyFunction` .
	//
	// The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-version.html#cfn-lambda-version-functionname
	//
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// Only publish a version if the hash value matches the value that's specified.
	//
	// Use this option to avoid publishing a version if the function code has changed since you last updated it. Updates are not supported for this property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-version.html#cfn-lambda-version-codesha256
	//
	CodeSha256 *string `field:"optional" json:"codeSha256" yaml:"codeSha256"`
	// A description for the version to override the description in the function configuration.
	//
	// Updates are not supported for this property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-version.html#cfn-lambda-version-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies a provisioned concurrency configuration for a function's version.
	//
	// Updates are not supported for this property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-version.html#cfn-lambda-version-provisionedconcurrencyconfig
	//
	ProvisionedConcurrencyConfig interface{} `field:"optional" json:"provisionedConcurrencyConfig" yaml:"provisionedConcurrencyConfig"`
	// Runtime Management Config of a function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-version.html#cfn-lambda-version-runtimepolicy
	//
	RuntimePolicy interface{} `field:"optional" json:"runtimePolicy" yaml:"runtimePolicy"`
}

