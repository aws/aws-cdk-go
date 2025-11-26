package previewawsdynamodbmixins


// Provides visibility into the number of read and write operations your table or secondary index can instantaneously support.
//
// The settings can be modified using the `UpdateTable` operation to meet the throughput requirements of an upcoming peak event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   warmThroughputProperty := &WarmThroughputProperty{
//   	ReadUnitsPerSecond: jsii.Number(123),
//   	WriteUnitsPerSecond: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-warmthroughput.html
//
type CfnTablePropsMixin_WarmThroughputProperty struct {
	// Represents the number of read operations your base table can instantaneously support.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-warmthroughput.html#cfn-dynamodb-table-warmthroughput-readunitspersecond
	//
	ReadUnitsPerSecond *float64 `field:"optional" json:"readUnitsPerSecond" yaml:"readUnitsPerSecond"`
	// Represents the number of write operations your base table can instantaneously support.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-warmthroughput.html#cfn-dynamodb-table-warmthroughput-writeunitspersecond
	//
	WriteUnitsPerSecond *float64 `field:"optional" json:"writeUnitsPerSecond" yaml:"writeUnitsPerSecond"`
}

