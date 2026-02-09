package cxapi


// Return the appropriate values for the environment placeholders.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentPlaceholderValues := &EnvironmentPlaceholderValues{
//   	AccountId: jsii.String("accountId"),
//   	Partition: jsii.String("partition"),
//   	Region: jsii.String("region"),
//   }
//
// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
type EnvironmentPlaceholderValues struct {
	// Return the account.
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Return the partition.
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	Partition *string `field:"required" json:"partition" yaml:"partition"`
	// Return the region.
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	Region *string `field:"required" json:"region" yaml:"region"`
}

