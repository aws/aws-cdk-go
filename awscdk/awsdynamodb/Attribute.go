package awsdynamodb


// Represents an attribute for describing the key schema for the table and indexes.
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
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   	DeletionProtection: jsii.Boolean(true),
//   	// only the replica in us-east-1 will be deleted during stack deletion
//   	Replicas: []replicaTableProps{
//   		&replicaTableProps{
//   			Region: jsii.String("us-east-1"),
//   			DeletionProtection: jsii.Boolean(false),
//   		},
//   		&replicaTableProps{
//   			Region: jsii.String("us-east-2"),
//   			DeletionProtection: jsii.Boolean(true),
//   		},
//   	},
//   })
//
type Attribute struct {
	// The name of an attribute.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The data type of an attribute.
	Type AttributeType `field:"required" json:"type" yaml:"type"`
}

