package awsdynamodb


// Options used to configure global secondary indexes on a replica table.
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
//   	ContributorInsights: jsii.Boolean(true),
//   	Billing: dynamodb.Billing_Provisioned(&ThroughputProps{
//   		ReadCapacity: dynamodb.Capacity_Fixed(jsii.Number(10)),
//   		WriteCapacity: dynamodb.Capacity_Autoscaled(&AutoscaledCapacityOptions{
//   			MaxCapacity: jsii.Number(10),
//   		}),
//   	}),
//   	// each global secondary index will inherit contributor insights as true
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
//   	Replicas: []replicaTableProps{
//   		&replicaTableProps{
//   			Region: jsii.String("us-east-1"),
//   			GlobalSecondaryIndexOptions: map[string]replicaGlobalSecondaryIndexOptions{
//   				"gsi1": &replicaGlobalSecondaryIndexOptions{
//   					"readCapacity": dynamodb.Capacity_*Autoscaled(&AutoscaledCapacityOptions{
//   						"minCapacity": jsii.Number(1),
//   						"maxCapacity": jsii.Number(10),
//   					}),
//   				},
//   			},
//   		},
//   		&replicaTableProps{
//   			Region: jsii.String("us-east-2"),
//   			GlobalSecondaryIndexOptions: map[string]*replicaGlobalSecondaryIndexOptions{
//   				"gsi2": &replicaGlobalSecondaryIndexOptions{
//   					"contributorInsights": jsii.Boolean(false),
//   				},
//   			},
//   		},
//   	},
//   })
//
type ReplicaGlobalSecondaryIndexOptions struct {
	// Whether CloudWatch contributor insights is enabled for a specific global secondary index on a replica table.
	// Default: - inherited from the primary table.
	//
	ContributorInsights *bool `field:"optional" json:"contributorInsights" yaml:"contributorInsights"`
	// The read capacity for a specific global secondary index on a replica table.
	//
	// Note: This can only be configured if primary table billing is provisioned.
	// Default: - inherited from the primary table.
	//
	ReadCapacity Capacity `field:"optional" json:"readCapacity" yaml:"readCapacity"`
}

