package awscassandra


// The provisioned throughput for the table, which consists of `ReadCapacityUnits` and `WriteCapacityUnits` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   provisionedThroughputProperty := &ProvisionedThroughputProperty{
//   	ReadCapacityUnits: jsii.Number(123),
//   	WriteCapacityUnits: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-provisionedthroughput.html
//
type CfnTable_ProvisionedThroughputProperty struct {
	// The amount of read capacity that's provisioned for the table.
	//
	// For more information, see [Read/write capacity mode](https://docs.aws.amazon.com/keyspaces/latest/devguide/ReadWriteCapacityMode.html) in the *Amazon Keyspaces Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-provisionedthroughput.html#cfn-cassandra-table-provisionedthroughput-readcapacityunits
	//
	ReadCapacityUnits *float64 `field:"required" json:"readCapacityUnits" yaml:"readCapacityUnits"`
	// The amount of write capacity that's provisioned for the table.
	//
	// For more information, see [Read/write capacity mode](https://docs.aws.amazon.com/keyspaces/latest/devguide/ReadWriteCapacityMode.html) in the *Amazon Keyspaces Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-provisionedthroughput.html#cfn-cassandra-table-provisionedthroughput-writecapacityunits
	//
	WriteCapacityUnits *float64 `field:"required" json:"writeCapacityUnits" yaml:"writeCapacityUnits"`
}

