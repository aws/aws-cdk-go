package awsdynamodb


// Options used to configure autoscaled capacity.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   app := cdk.NewApp()
//   stack := cdk.NewStack(app, jsii.String("Stack"), &StackProps{
//   	Env: &Environment{
//   		Region: jsii.String("us-west-2"),
//   	},
//   })
//
//   globalTable := dynamodb.NewTableV2(stack, jsii.String("GlobalTable"), &TablePropsV2{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("pk"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	Billing: dynamodb.Billing_Provisioned(&ThroughputProps{
//   		ReadCapacity: dynamodb.Capacity_Fixed(jsii.Number(10)),
//   		WriteCapacity: dynamodb.Capacity_Autoscaled(&AutoscaledCapacityOptions{
//   			MaxCapacity: jsii.Number(15),
//   		}),
//   	}),
//   	Replicas: []ReplicaTableProps{
//   		&ReplicaTableProps{
//   			Region: jsii.String("us-east-1"),
//   		},
//   		&ReplicaTableProps{
//   			Region: jsii.String("us-east-2"),
//   			ReadCapacity: dynamodb.Capacity_*Autoscaled(&AutoscaledCapacityOptions{
//   				MaxCapacity: jsii.Number(20),
//   				TargetUtilizationPercent: jsii.Number(50),
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
	// If you want to switch a table's billing mode from on-demand to provisioned or from provisioned to on-demand, you must specify a value for this property for each autoscaled resource.
	// Default: no seed capacity.
	//
	SeedCapacity *float64 `field:"optional" json:"seedCapacity" yaml:"seedCapacity"`
	// The ratio of consumed capacity units to provisioned capacity units.
	//
	// Note: Target utilization percent cannot be less than 20 and cannot be greater
	// than 90.
	// Default: 70.
	//
	TargetUtilizationPercent *float64 `field:"optional" json:"targetUtilizationPercent" yaml:"targetUtilizationPercent"`
}

