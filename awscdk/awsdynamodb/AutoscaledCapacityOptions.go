package awsdynamodb


// Options used to configure autoscaled capacity.
//
// Example:
//   table := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("pk"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	Billing: dynamodb.Billing_Provisioned(&ThroughputProps{
//   		ReadCapacity: dynamodb.Capacity_Fixed(jsii.Number(10)),
//   		WriteCapacity: dynamodb.Capacity_Autoscaled(&AutoscaledCapacityOptions{
//   			MaxCapacity: jsii.Number(10),
//   		}),
//   	}),
//   	GlobalSecondaryIndexes: []globalSecondaryIndexPropsV2{
//   		&globalSecondaryIndexPropsV2{
//   			IndexName: jsii.String("gsi1"),
//   			PartitionKey: &Attribute{
//   				Name: jsii.String("pk"),
//   				Type: dynamodb.AttributeType_STRING,
//   			},
//   			ReadCapacity: dynamodb.Capacity_*Fixed(jsii.Number(15)),
//   		},
//   		&globalSecondaryIndexPropsV2{
//   			IndexName: jsii.String("gsi2"),
//   			PartitionKey: &Attribute{
//   				Name: jsii.String("pk"),
//   				Type: dynamodb.AttributeType_STRING,
//   			},
//   			WriteCapacity: dynamodb.Capacity_*Autoscaled(&AutoscaledCapacityOptions{
//   				MinCapacity: jsii.Number(5),
//   				MaxCapacity: jsii.Number(20),
//   			}),
//   		},
//   	},
//   })
//
type AutoscaledCapacityOptions struct {
	// The maximum allowable capacity.
	MaxCapacity *float64 `field:"required" json:"maxCapacity" yaml:"maxCapacity"`
	// The minimum allowable capacity.
	// Default: 1.
	//
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
	// The ratio of consumed capacity units to provisioned capacity units.
	//
	// Note: Target utilization percent cannot be less than 20 and cannot be greater
	// than 90.
	// Default: 70.
	//
	TargetUtilizationPercent *float64 `field:"optional" json:"targetUtilizationPercent" yaml:"targetUtilizationPercent"`
}

