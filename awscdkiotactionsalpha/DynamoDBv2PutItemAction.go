package awscdkiotactionsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkiotactionsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdkiotactionsalpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdkiotalpha/v2"
)

// The action to put the record from an MQTT message to the DynamoDB table.
//
// Example:
//   import dynamodb "github.com/aws/aws-cdk-go/awscdk"
//
//   var table Table
//
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
//   	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
//   	Actions: []IAction{
//   		actions.NewDynamoDBv2PutItemAction(table),
//   	},
//   })
//
// Experimental.
type DynamoDBv2PutItemAction interface {
	awscdkiotalpha.IAction
}

// The jsii proxy struct for DynamoDBv2PutItemAction
type jsiiProxy_DynamoDBv2PutItemAction struct {
	internal.Type__awscdkiotalphaIAction
}

// Experimental.
func NewDynamoDBv2PutItemAction(table awsdynamodb.ITable, props *DynamoDBv2PutItemActionProps) DynamoDBv2PutItemAction {
	_init_.Initialize()

	if err := validateNewDynamoDBv2PutItemActionParameters(table, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DynamoDBv2PutItemAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.DynamoDBv2PutItemAction",
		[]interface{}{table, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDynamoDBv2PutItemAction_Override(d DynamoDBv2PutItemAction, table awsdynamodb.ITable, props *DynamoDBv2PutItemActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.DynamoDBv2PutItemAction",
		[]interface{}{table, props},
		d,
	)
}

