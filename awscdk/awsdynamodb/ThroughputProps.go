package awsdynamodb


// Properties used to configure provisioned throughput for a DynamoDB table.
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
type ThroughputProps struct {
	// The read capacity.
	ReadCapacity Capacity `field:"required" json:"readCapacity" yaml:"readCapacity"`
	// The write capacity.
	WriteCapacity Capacity `field:"required" json:"writeCapacity" yaml:"writeCapacity"`
}

