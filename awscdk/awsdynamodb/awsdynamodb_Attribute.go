package awsdynamodb


// Represents an attribute for describing the key schema for the table and indexes.
//
// Example:
//   globalTable := dynamodb.NewTable(this, jsii.String("Table"), &tableProps{
//   	partitionKey: &attribute{
//   		name: jsii.String("id"),
//   		type: dynamodb.attributeType_STRING,
//   	},
//   	replicationRegions: []*string{
//   		jsii.String("us-east-1"),
//   		jsii.String("us-east-2"),
//   		jsii.String("us-west-2"),
//   	},
//   	billingMode: dynamodb.billingMode_PROVISIONED,
//   })
//
//   globalTable.autoScaleWriteCapacity(&enableScalingProps{
//   	minCapacity: jsii.Number(1),
//   	maxCapacity: jsii.Number(10),
//   }).scaleOnUtilization(&utilizationScalingProps{
//   	targetUtilizationPercent: jsii.Number(75),
//   })
//
type Attribute struct {
	// The name of an attribute.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The data type of an attribute.
	Type AttributeType `field:"required" json:"type" yaml:"type"`
}

