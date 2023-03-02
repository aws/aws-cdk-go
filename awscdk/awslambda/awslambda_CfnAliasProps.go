package awslambda


// Properties for defining a `CfnAlias`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAliasProps := &cfnAliasProps{
//   	functionName: jsii.String("functionName"),
//   	functionVersion: jsii.String("functionVersion"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	provisionedConcurrencyConfig: &provisionedConcurrencyConfigurationProperty{
//   		provisionedConcurrentExecutions: jsii.Number(123),
//   	},
//   	routingConfig: &aliasRoutingConfigurationProperty{
//   		additionalVersionWeights: []interface{}{
//   			&versionWeightProperty{
//   				functionVersion: jsii.String("functionVersion"),
//   				functionWeight: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
type CfnAliasProps struct {
	// The name of the Lambda function.
	//
	// **Name formats** - *Function name* - `MyFunction` .
	// - *Function ARN* - `arn:aws:lambda:us-west-2:123456789012:function:MyFunction` .
	// - *Partial ARN* - `123456789012:function:MyFunction` .
	//
	// The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// The function version that the alias invokes.
	FunctionVersion *string `field:"required" json:"functionVersion" yaml:"functionVersion"`
	// The name of the alias.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the alias.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies a [provisioned concurrency](https://docs.aws.amazon.com/lambda/latest/dg/configuration-concurrency.html) configuration for a function's alias.
	ProvisionedConcurrencyConfig interface{} `field:"optional" json:"provisionedConcurrencyConfig" yaml:"provisionedConcurrencyConfig"`
	// The [routing configuration](https://docs.aws.amazon.com/lambda/latest/dg/lambda-traffic-shifting-using-aliases.html) of the alias.
	RoutingConfig interface{} `field:"optional" json:"routingConfig" yaml:"routingConfig"`
}

