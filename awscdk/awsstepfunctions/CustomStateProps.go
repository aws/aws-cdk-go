package awsstepfunctions


// Properties for defining a custom state definition.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   // create a table
//   table := dynamodb.NewTable(this, jsii.String("montable"), &TableProps{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("id"),
//   		Type: dynamodb.AttributeType_STRING,
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
//   custom := sfn.NewCustomState(this, jsii.String("my custom task"), &CustomStateProps{
//   	StateJson: StateJson,
//   })
//
//   // catch errors with addCatch
//   errorHandler := sfn.NewPass(this, jsii.String("handle failure"))
//   custom.AddCatch(errorHandler)
//
//   // retry the task if something goes wrong
//   custom.AddRetry(&RetryProps{
//   	Errors: []*string{
//   		sfn.Errors_ALL(),
//   	},
//   	Interval: awscdk.Duration_Seconds(jsii.Number(10)),
//   	MaxAttempts: jsii.Number(5),
//   })
//
//   chain := sfn.Chain_Start(custom).Next(finalStatus)
//
//   sm := sfn.NewStateMachine(this, jsii.String("StateMachine"), &StateMachineProps{
//   	DefinitionBody: sfn.DefinitionBody_FromChainable(chain),
//   	Timeout: awscdk.Duration_*Seconds(jsii.Number(30)),
//   	Comment: jsii.String("a super cool state machine"),
//   })
//
//   // don't forget permissions. You need to assign them
//   table.GrantWriteData(sm)
//
type CustomStateProps struct {
	// Amazon States Language (JSON-based) definition of the state.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html
	//
	StateJson *map[string]interface{} `field:"required" json:"stateJson" yaml:"stateJson"`
}

