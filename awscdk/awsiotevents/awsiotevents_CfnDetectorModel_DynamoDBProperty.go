package awsiotevents


// Defines an action to write to the Amazon DynamoDB table that you created.
//
// The standard action payload contains all the information about the detector model instance and the event that triggered the action. You can customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html) . One column of the DynamoDB table receives all attribute-value pairs in the payload that you specify.
//
// You must use expressions for all parameters in `DynamoDBAction` . The expressions accept literals, operators, functions, references, and substitution templates.
//
// **Examples** - For literal values, the expressions must contain single quotes. For example, the value for the `hashKeyType` parameter can be `'STRING'` .
// - For references, you must specify either variables or input values. For example, the value for the `hashKeyField` parameter can be `$input.GreenhouseInput.name` .
// - For a substitution template, you must use `${}` , and the template must be in single quotes. A substitution template can also contain a combination of literals, operators, functions, references, and substitution templates.
//
// In the following example, the value for the `hashKeyValue` parameter uses a substitution template.
//
// `'${$input.GreenhouseInput.temperature * 6 / 5 + 32} in Fahrenheit'`
// - For a string concatenation, you must use `+` . A string concatenation can also contain a combination of literals, operators, functions, references, and substitution templates.
//
// In the following example, the value for the `tableName` parameter uses a string concatenation.
//
// `'GreenhouseTemperatureTable ' + $input.GreenhouseInput.date`
//
// For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *AWS IoT Events Developer Guide* .
//
// If the defined payload type is a string, `DynamoDBAction` writes non-JSON data to the DynamoDB table as binary data. The DynamoDB console displays the data as Base64-encoded text. The value for the `payloadField` parameter is `<payload-field>_raw` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dynamoDBProperty := &dynamoDBProperty{
//   	hashKeyField: jsii.String("hashKeyField"),
//   	hashKeyValue: jsii.String("hashKeyValue"),
//   	tableName: jsii.String("tableName"),
//
//   	// the properties below are optional
//   	hashKeyType: jsii.String("hashKeyType"),
//   	operation: jsii.String("operation"),
//   	payload: &payloadProperty{
//   		contentExpression: jsii.String("contentExpression"),
//   		type: jsii.String("type"),
//   	},
//   	payloadField: jsii.String("payloadField"),
//   	rangeKeyField: jsii.String("rangeKeyField"),
//   	rangeKeyType: jsii.String("rangeKeyType"),
//   	rangeKeyValue: jsii.String("rangeKeyValue"),
//   }
//
type CfnDetectorModel_DynamoDBProperty struct {
	// The name of the hash key (also called the partition key).
	//
	// The `hashKeyField` value must match the partition key of the target DynamoDB table.
	HashKeyField *string `field:"required" json:"hashKeyField" yaml:"hashKeyField"`
	// The value of the hash key (also called the partition key).
	HashKeyValue *string `field:"required" json:"hashKeyValue" yaml:"hashKeyValue"`
	// The name of the DynamoDB table.
	//
	// The `tableName` value must match the table name of the target DynamoDB table.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// The data type for the hash key (also called the partition key). You can specify the following values:.
	//
	// - `'STRING'` - The hash key is a string.
	// - `'NUMBER'` - The hash key is a number.
	//
	// If you don't specify `hashKeyType` , the default value is `'STRING'` .
	HashKeyType *string `field:"optional" json:"hashKeyType" yaml:"hashKeyType"`
	// The type of operation to perform. You can specify the following values:.
	//
	// - `'INSERT'` - Insert data as a new item into the DynamoDB table. This item uses the specified hash key as a partition key. If you specified a range key, the item uses the range key as a sort key.
	// - `'UPDATE'` - Update an existing item of the DynamoDB table with new data. This item's partition key must match the specified hash key. If you specified a range key, the range key must match the item's sort key.
	// - `'DELETE'` - Delete an existing item of the DynamoDB table. This item's partition key must match the specified hash key. If you specified a range key, the range key must match the item's sort key.
	//
	// If you don't specify this parameter, AWS IoT Events triggers the `'INSERT'` operation.
	Operation *string `field:"optional" json:"operation" yaml:"operation"`
	// Information needed to configure the payload.
	//
	// By default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use `contentExpression` .
	Payload interface{} `field:"optional" json:"payload" yaml:"payload"`
	// The name of the DynamoDB column that receives the action payload.
	//
	// If you don't specify this parameter, the name of the DynamoDB column is `payload` .
	PayloadField *string `field:"optional" json:"payloadField" yaml:"payloadField"`
	// The name of the range key (also called the sort key).
	//
	// The `rangeKeyField` value must match the sort key of the target DynamoDB table.
	RangeKeyField *string `field:"optional" json:"rangeKeyField" yaml:"rangeKeyField"`
	// The data type for the range key (also called the sort key), You can specify the following values:.
	//
	// - `'STRING'` - The range key is a string.
	// - `'NUMBER'` - The range key is number.
	//
	// If you don't specify `rangeKeyField` , the default value is `'STRING'` .
	RangeKeyType *string `field:"optional" json:"rangeKeyType" yaml:"rangeKeyType"`
	// The value of the range key (also called the sort key).
	RangeKeyValue *string `field:"optional" json:"rangeKeyValue" yaml:"rangeKeyValue"`
}

