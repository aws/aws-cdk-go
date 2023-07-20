package awslambda


// A provisioned concurrency configuration for a function's alias.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   provisionedConcurrencyConfigurationProperty := &ProvisionedConcurrencyConfigurationProperty{
//   	ProvisionedConcurrentExecutions: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-alias-provisionedconcurrencyconfiguration.html
//
type CfnAlias_ProvisionedConcurrencyConfigurationProperty struct {
	// The amount of provisioned concurrency to allocate for the alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-alias-provisionedconcurrencyconfiguration.html#cfn-lambda-alias-provisionedconcurrencyconfiguration-provisionedconcurrentexecutions
	//
	ProvisionedConcurrentExecutions *float64 `field:"required" json:"provisionedConcurrentExecutions" yaml:"provisionedConcurrentExecutions"`
}

