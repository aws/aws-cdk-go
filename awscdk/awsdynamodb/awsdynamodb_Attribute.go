package awsdynamodb


// Represents an attribute for describing the key schema for the table and indexes.
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
// Experimental.
type Attribute struct {
	// The name of an attribute.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The data type of an attribute.
	// Experimental.
	Type AttributeType `field:"required" json:"type" yaml:"type"`
}

