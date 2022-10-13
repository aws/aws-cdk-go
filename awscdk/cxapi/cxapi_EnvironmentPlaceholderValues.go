package cxapi


// Return the appropriate values for the environment placeholders.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentPlaceholderValues := &environmentPlaceholderValues{
//   	accountId: jsii.String("accountId"),
//   	partition: jsii.String("partition"),
//   	region: jsii.String("region"),
//   }
//
type EnvironmentPlaceholderValues struct {
	// Return the account.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Return the partition.
	Partition *string `field:"required" json:"partition" yaml:"partition"`
	// Return the region.
	Region *string `field:"required" json:"region" yaml:"region"`
}

