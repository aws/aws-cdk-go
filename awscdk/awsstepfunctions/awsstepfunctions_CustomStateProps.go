package awsstepfunctions


// Properties for defining a custom state definition.
//
// Example:
//   import dynamodb "github.com/aws/aws-cdk-go/awscdk"
//
//
//   // create a table
//   table := dynamodb.NewTable(this, jsii.String("montable"), &tableProps{
//   	partitionKey: &attribute{
//   		name: jsii.String("id"),
//   		type: dynamodb.attributeType_STRING,
//   	},
//   })
//
//   finalStatus := sfn.NewPass(this, jsii.String("final step"))
//
//   // States language JSON to put an item into DynamoDB
//   // snippet generated from https://docs.aws.amazon.com/step-functions/latest/dg/tutorial-code-snippet.html#tutorial-code-snippet-1
//   stateJson := map[string]interface{}{
//   	"Type": jsii.String("Task"),
//   	"Resource": jsii.String("arn:aws:states:::dynamodb:putItem"),
//   	"Parameters": map[string]interface{}{
//   		"TableName": table.tableName,
//   		"Item": map[string]map[string]*string{
//   			"id": map[string]*string{
//   				"S": jsii.String("MyEntry"),
//   			},
//   		},
//   	},
//   	"ResultPath": nil,
//   }
//
//   // custom state which represents a task to insert data into DynamoDB
//   custom := sfn.NewCustomState(this, jsii.String("my custom task"), &customStateProps{
//   	stateJson: stateJson,
//   })
//
//   chain := sfn.chain.start(custom).next(finalStatus)
//
//   sm := sfn.NewStateMachine(this, jsii.String("StateMachine"), &stateMachineProps{
//   	definition: chain,
//   	timeout: awscdk.Duration.seconds(jsii.Number(30)),
//   })
//
//   // don't forget permissions. You need to assign them
//   table.grantWriteData(sm)
//
type CustomStateProps struct {
	// Amazon States Language (JSON-based) definition of the state.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html
	//
	StateJson *map[string]interface{} `field:"required" json:"stateJson" yaml:"stateJson"`
}

