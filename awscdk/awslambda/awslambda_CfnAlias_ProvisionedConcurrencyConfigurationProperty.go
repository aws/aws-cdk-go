package awslambda


// A provisioned concurrency configuration for a function's alias.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   provisionedConcurrencyConfigurationProperty := &provisionedConcurrencyConfigurationProperty{
//   	provisionedConcurrentExecutions: jsii.Number(123),
//   }
//
type CfnAlias_ProvisionedConcurrencyConfigurationProperty struct {
	// The amount of provisioned concurrency to allocate for the alias.
	ProvisionedConcurrentExecutions *float64 `field:"required" json:"provisionedConcurrentExecutions" yaml:"provisionedConcurrentExecutions"`
}

