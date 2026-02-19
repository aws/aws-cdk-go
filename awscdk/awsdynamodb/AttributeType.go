package awsdynamodb


// Data types for attributes within a table.
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
// See: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.NamingRulesDataTypes.html#HowItWorks.DataTypes
//
type AttributeType string

const (
	// Up to 400KiB of binary data (which must be encoded as base64 before sending to DynamoDB).
	AttributeType_BINARY AttributeType = "BINARY"
	// Numeric values made of up to 38 digits (positive, negative or zero).
	AttributeType_NUMBER AttributeType = "NUMBER"
	// Up to 400KiB of UTF-8 encoded text.
	AttributeType_STRING AttributeType = "STRING"
)

