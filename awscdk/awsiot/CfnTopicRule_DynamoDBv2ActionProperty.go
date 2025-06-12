package awsiot


// Describes an action to write to a DynamoDB table.
//
// This DynamoDB action writes each attribute in the message payload into it's own column in the DynamoDB table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dynamoDBv2ActionProperty := &DynamoDBv2ActionProperty{
//   	PutItem: &PutItemInputProperty{
//   		TableName: jsii.String("tableName"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-dynamodbv2action.html
//
type CfnTopicRule_DynamoDBv2ActionProperty struct {
	// Specifies the DynamoDB table to which the message data will be written. For example:.
	//
	// `{ "dynamoDBv2": { "roleArn": "aws:iam:12341251:my-role" "putItem": { "tableName": "my-table" } } }`
	//
	// Each attribute in the message payload will be written to a separate column in the DynamoDB database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-dynamodbv2action.html#cfn-iot-topicrule-dynamodbv2action-putitem
	//
	PutItem interface{} `field:"optional" json:"putItem" yaml:"putItem"`
	// The ARN of the IAM role that grants access to the DynamoDB table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-dynamodbv2action.html#cfn-iot-topicrule-dynamodbv2action-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

