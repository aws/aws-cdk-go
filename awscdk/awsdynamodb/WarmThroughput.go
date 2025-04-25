package awsdynamodb


// Reference to WarmThroughput for a DynamoDB table.
//
// Example:
//   table := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("id"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	WarmThroughput: &WarmThroughput{
//   		ReadUnitsPerSecond: jsii.Number(15000),
//   		WriteUnitsPerSecond: jsii.Number(20000),
//   	},
//   })
//
type WarmThroughput struct {
	// Configures the number of read units per second a table will be able to handle instantly.
	// Default: - no readUnitsPerSecond configured.
	//
	ReadUnitsPerSecond *float64 `field:"optional" json:"readUnitsPerSecond" yaml:"readUnitsPerSecond"`
	// Configures the number of write units per second a table will be able to handle instantly.
	// Default: - no writeUnitsPerSecond configured.
	//
	WriteUnitsPerSecond *float64 `field:"optional" json:"writeUnitsPerSecond" yaml:"writeUnitsPerSecond"`
}

