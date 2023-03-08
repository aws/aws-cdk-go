package awsiot


// The input for the DynamoActionVS action that specifies the DynamoDB table to which the message data will be written.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   putItemInputProperty := &putItemInputProperty{
//   	tableName: jsii.String("tableName"),
//   }
//
type CfnTopicRule_PutItemInputProperty struct {
	// The table where the message data will be written.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

