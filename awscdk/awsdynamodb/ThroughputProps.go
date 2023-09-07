package awsdynamodb


// Properties used to configure provisioned throughput for a DynamoDB table.
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
type ThroughputProps struct {
	// The read capacity.
	ReadCapacity Capacity `field:"required" json:"readCapacity" yaml:"readCapacity"`
	// The write capacity.
	WriteCapacity Capacity `field:"required" json:"writeCapacity" yaml:"writeCapacity"`
}

