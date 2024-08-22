package awsdynamodb


// Properties used to configure maximum throughput for an on-demand table.
//
// Example:
//   table := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("pk"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	Billing: dynamodb.Billing_OnDemand(&MaxThroughputProps{
//   		MaxReadRequestUnits: jsii.Number(100),
//   		MaxWriteRequestUnits: jsii.Number(115),
//   	}),
//   })
//
type MaxThroughputProps struct {
	// The max read request units.
	// Default: - if table mode is on-demand and this property is undefined,
	// no maximum throughput limit will be put in place for read requests.
	// This property is only applicable for tables using on-demand mode.
	//
	MaxReadRequestUnits *float64 `field:"optional" json:"maxReadRequestUnits" yaml:"maxReadRequestUnits"`
	// The max write request units.
	// Default: - if table mode is on-demand and this property is undefined,
	// no maximum throughput limit will be put in place for write requests.
	// This property is only applicable for tables using on-demand mode.
	//
	MaxWriteRequestUnits *float64 `field:"optional" json:"maxWriteRequestUnits" yaml:"maxWriteRequestUnits"`
}

