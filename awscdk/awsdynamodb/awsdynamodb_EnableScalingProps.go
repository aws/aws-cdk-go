package awsdynamodb


// Properties for enabling DynamoDB capacity scaling.
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
type EnableScalingProps struct {
	// Maximum capacity to scale to.
	MaxCapacity *float64 `field:"required" json:"maxCapacity" yaml:"maxCapacity"`
	// Minimum capacity to scale to.
	MinCapacity *float64 `field:"required" json:"minCapacity" yaml:"minCapacity"`
}

