package awscassandra


// Determines the billing mode for the table - on-demand or provisioned.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   billingModeProperty := &BillingModeProperty{
//   	Mode: jsii.String("mode"),
//
//   	// the properties below are optional
//   	ProvisionedThroughput: &ProvisionedThroughputProperty{
//   		ReadCapacityUnits: jsii.Number(123),
//   		WriteCapacityUnits: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-billingmode.html
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-billingmode.html#cfn-cassandra-table-billingmode-mode
	//
	// Default: - "ON_DEMAND".
	//
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// The provisioned read capacity and write capacity for the table.
	//
	// For more information, see [Provisioned throughput capacity mode](https://docs.aws.amazon.com/keyspaces/latest/devguide/ReadWriteCapacityMode.html#ReadWriteCapacityMode.Provisioned) in the *Amazon Keyspaces Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-billingmode.html#cfn-cassandra-table-billingmode-provisionedthroughput
	//
	ProvisionedThroughput interface{} `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
}

