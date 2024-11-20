package awsdynamodb


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-warmthroughput.html
//
type CfnGlobalTable_WarmThroughputProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-warmthroughput.html#cfn-dynamodb-globaltable-warmthroughput-readunitspersecond
	//
	ReadUnitsPerSecond *float64 `field:"optional" json:"readUnitsPerSecond" yaml:"readUnitsPerSecond"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-warmthroughput.html#cfn-dynamodb-globaltable-warmthroughput-writeunitspersecond
	//
	WriteUnitsPerSecond *float64 `field:"optional" json:"writeUnitsPerSecond" yaml:"writeUnitsPerSecond"`
}

