package awscassandra


// Warm throughput configuration for the table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   warmThroughputProperty := &WarmThroughputProperty{
//   	ReadUnitsPerSecond: jsii.Number(123),
//   	WriteUnitsPerSecond: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-warmthroughput.html
//
type CfnTable_WarmThroughputProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-warmthroughput.html#cfn-cassandra-table-warmthroughput-readunitspersecond
	//
	ReadUnitsPerSecond *float64 `field:"optional" json:"readUnitsPerSecond" yaml:"readUnitsPerSecond"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-warmthroughput.html#cfn-cassandra-table-warmthroughput-writeunitspersecond
	//
	WriteUnitsPerSecond *float64 `field:"optional" json:"writeUnitsPerSecond" yaml:"writeUnitsPerSecond"`
}

