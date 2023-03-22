package awslambda


// A [provisioned concurrency](https://docs.aws.amazon.com/lambda/latest/dg/configuration-concurrency.html) configuration for a function's version.
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
type CfnVersion_ProvisionedConcurrencyConfigurationProperty struct {
	// The amount of provisioned concurrency to allocate for the version.
	ProvisionedConcurrentExecutions *float64 `field:"required" json:"provisionedConcurrentExecutions" yaml:"provisionedConcurrentExecutions"`
}

