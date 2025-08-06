package awslambda


// Properties for defining a `CfnAlias`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAliasProps := &CfnAliasProps{
//   	FunctionName: jsii.String("functionName"),
//   	FunctionVersion: jsii.String("functionVersion"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	ProvisionedConcurrencyConfig: &ProvisionedConcurrencyConfigurationProperty{
//   		ProvisionedConcurrentExecutions: jsii.Number(123),
//   	},
//   	RoutingConfig: &AliasRoutingConfigurationProperty{
//   		AdditionalVersionWeights: []interface{}{
//   			&VersionWeightProperty{
//   				FunctionVersion: jsii.String("functionVersion"),
//   				FunctionWeight: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-alias.html
//
type CfnAliasProps struct {
	// The name or ARN of the Lambda function.
	//
	// **Name formats** - *Function name* - `MyFunction` .
	// - *Function ARN* - `arn:aws:lambda:us-west-2:123456789012:function:MyFunction` .
	// - *Partial ARN* - `123456789012:function:MyFunction` .
	//
	// The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-alias.html#cfn-lambda-alias-functionname
	//
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// The function version that the alias invokes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-alias.html#cfn-lambda-alias-functionversion
	//
	FunctionVersion *string `field:"required" json:"functionVersion" yaml:"functionVersion"`
	// The name of the alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-alias.html#cfn-lambda-alias-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-alias.html#cfn-lambda-alias-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies a [provisioned concurrency](https://docs.aws.amazon.com/lambda/latest/dg/configuration-concurrency.html) configuration for a function's alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-alias.html#cfn-lambda-alias-provisionedconcurrencyconfig
	//
	ProvisionedConcurrencyConfig interface{} `field:"optional" json:"provisionedConcurrencyConfig" yaml:"provisionedConcurrencyConfig"`
	// The [routing configuration](https://docs.aws.amazon.com/lambda/latest/dg/lambda-traffic-shifting-using-aliases.html) of the alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-alias.html#cfn-lambda-alias-routingconfig
	//
	RoutingConfig interface{} `field:"optional" json:"routingConfig" yaml:"routingConfig"`
}

