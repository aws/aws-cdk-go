package awscassandra


// Determines the billing mode for the table - On-demand or provisioned.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   billingModeProperty := &billingModeProperty{
//   	mode: jsii.String("mode"),
//
//   	// the properties below are optional
//   	provisionedThroughput: &provisionedThroughputProperty{
//   		readCapacityUnits: jsii.Number(123),
//   		writeCapacityUnits: jsii.Number(123),
//   	},
//   }
//
type CfnTable_BillingModeProperty struct {
	// The billing mode for the table:.
	//
	// - On-demand mode - `ON_DEMAND`
	// - Provisioned mode - `PROVISIONED`
	//
	// > If you choose `PROVISIONED` mode, then you also need to specify provisioned throughput (read and write capacity) for the table.
	//
	// Valid values: `ON_DEMAND` | `PROVISIONED`.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// The provisioned read capacity and write capacity for the table.
	//
	// For more information, see [Provisioned throughput capacity mode](https://docs.aws.amazon.com/keyspaces/latest/devguide/ReadWriteCapacityMode.html#ReadWriteCapacityMode.Provisioned) in the *Amazon Keyspaces Developer Guide* .
	ProvisionedThroughput interface{} `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
}

