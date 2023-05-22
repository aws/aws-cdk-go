package awsdynamodb


// Data types for attributes within a table.
//
// Example:
//   globalTable := dynamodb.NewTable(this, jsii.String("Table"), &TableProps{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("id"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	ReplicationRegions: []*string{
//   		jsii.String("us-east-1"),
//   		jsii.String("us-east-2"),
//   		jsii.String("us-west-2"),
//   	},
//   	BillingMode: dynamodb.BillingMode_PROVISIONED,
//   })
//
//   globalTable.AutoScaleWriteCapacity(&EnableScalingProps{
//   	MinCapacity: jsii.Number(1),
//   	MaxCapacity: jsii.Number(10),
//   }).ScaleOnUtilization(&UtilizationScalingProps{
//   	TargetUtilizationPercent: jsii.Number(75),
//   })
//
// See: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.NamingRulesDataTypes.html#HowItWorks.DataTypes
//
// Experimental.
type AttributeType string

const (
	// Up to 400KiB of binary data (which must be encoded as base64 before sending to DynamoDB).
	// Experimental.
	AttributeType_BINARY AttributeType = "BINARY"
	// Numeric values made of up to 38 digits (positive, negative or zero).
	// Experimental.
	AttributeType_NUMBER AttributeType = "NUMBER"
	// Up to 400KiB of UTF-8 encoded text.
	// Experimental.
	AttributeType_STRING AttributeType = "STRING"
)

