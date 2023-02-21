package awsdynamodb


// Throughput for the specified table, which consists of values for `ReadCapacityUnits` and `WriteCapacityUnits` .
//
// For more information about the contents of a provisioned throughput structure, see [Amazon DynamoDB Table ProvisionedThroughput](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-provisionedthroughput.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   provisionedThroughputProperty := &provisionedThroughputProperty{
//   	readCapacityUnits: jsii.Number(123),
//   	writeCapacityUnits: jsii.Number(123),
//   }
//
type CfnTable_ProvisionedThroughputProperty struct {
	// The maximum number of strongly consistent reads consumed per second before DynamoDB returns a `ThrottlingException` .
	//
	// For more information, see [Specifying Read and Write Requirements](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.html#ProvisionedThroughput) in the *Amazon DynamoDB Developer Guide* .
	//
	// If read/write capacity mode is `PAY_PER_REQUEST` the value is set to 0.
	ReadCapacityUnits *float64 `field:"required" json:"readCapacityUnits" yaml:"readCapacityUnits"`
	// The maximum number of writes consumed per second before DynamoDB returns a `ThrottlingException` .
	//
	// For more information, see [Specifying Read and Write Requirements](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.html#ProvisionedThroughput) in the *Amazon DynamoDB Developer Guide* .
	//
	// If read/write capacity mode is `PAY_PER_REQUEST` the value is set to 0.
	WriteCapacityUnits *float64 `field:"required" json:"writeCapacityUnits" yaml:"writeCapacityUnits"`
}

