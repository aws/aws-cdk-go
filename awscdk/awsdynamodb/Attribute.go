package awsdynamodb


// Represents an attribute for describing the key schema for the table and indexes.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   app := cdk.NewApp()
//
//   // Source table in Account A
//   sourceStack := cdk.NewStack(app, jsii.String("SourceStack"), &StackProps{
//   	Env: &Environment{
//   		Region: jsii.String("us-east-1"),
//   		Account: jsii.String("111111111111"),
//   	},
//   })
//
//   // Region us-west-2
//   sourceTable := dynamodb.NewTableV2(sourceStack, jsii.String("SourceTable"), &TablePropsV2{
//   	TableName: jsii.String("MyMultiAccountTable"),
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("pk"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	GlobalTableSettingsReplicationMode: dynamodb.GlobalTableSettingsReplicationMode_ALL,
//   })
//   // After replica is deployed, update source stack with the ARN
//   sourceTable.Grants.MultiAccountReplicationTo(jsii.String("arn:aws:dynamodb:us-east-1:222222222222:table/MyMultiAccountTable"))
//
type Attribute struct {
	// The name of an attribute.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The data type of an attribute.
	Type AttributeType `field:"required" json:"type" yaml:"type"`
}

