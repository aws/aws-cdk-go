package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   provisionedConcurrencyConfigProperty := &provisionedConcurrencyConfigProperty{
//   	provisionedConcurrentExecutions: jsii.String("provisionedConcurrentExecutions"),
//   }
//
type CfnFunction_ProvisionedConcurrencyConfigProperty struct {
	// `CfnFunction.ProvisionedConcurrencyConfigProperty.ProvisionedConcurrentExecutions`.
	ProvisionedConcurrentExecutions *string `field:"required" json:"provisionedConcurrentExecutions" yaml:"provisionedConcurrentExecutions"`
}

