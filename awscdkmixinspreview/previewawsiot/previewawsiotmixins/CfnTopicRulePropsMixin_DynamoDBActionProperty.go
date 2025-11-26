package previewawsiotmixins


// Describes an action to write to a DynamoDB table.
//
// The `tableName` , `hashKeyField` , and `rangeKeyField` values must match the values used when you created the table.
//
// The `hashKeyValue` and `rangeKeyvalue` fields use a substitution template syntax. These templates provide data at runtime. The syntax is as follows: ${ *sql-expression* }.
//
// You can specify any valid expression in a WHERE or SELECT clause, including JSON properties, comparisons, calculations, and functions. For example, the following field uses the third level of the topic:
//
// `"hashKeyValue": "${topic(3)}"`
//
// The following field uses the timestamp:
//
// `"rangeKeyValue": "${timestamp()}"`
//
// For more information, see [DynamoDBv2 Action](https://docs.aws.amazon.com/iot/latest/developerguide/iot-rule-actions.html) in the *AWS IoT Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dynamoDBActionProperty := &DynamoDBActionProperty{
//   	HashKeyField: jsii.String("hashKeyField"),
//   	HashKeyType: jsii.String("hashKeyType"),
//   	HashKeyValue: jsii.String("hashKeyValue"),
//   	PayloadField: jsii.String("payloadField"),
//   	RangeKeyField: jsii.String("rangeKeyField"),
//   	RangeKeyType: jsii.String("rangeKeyType"),
//   	RangeKeyValue: jsii.String("rangeKeyValue"),
//   	RoleArn: jsii.String("roleArn"),
//   	TableName: jsii.String("tableName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-dynamodbaction.html
//
type CfnTopicRulePropsMixin_DynamoDBActionProperty struct {
	// The hash key name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-dynamodbaction.html#cfn-iot-topicrule-dynamodbaction-hashkeyfield
	//
	HashKeyField *string `field:"optional" json:"hashKeyField" yaml:"hashKeyField"`
	// The hash key type.
	//
	// Valid values are "STRING" or "NUMBER".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-dynamodbaction.html#cfn-iot-topicrule-dynamodbaction-hashkeytype
	//
	HashKeyType *string `field:"optional" json:"hashKeyType" yaml:"hashKeyType"`
	// The hash key value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-dynamodbaction.html#cfn-iot-topicrule-dynamodbaction-hashkeyvalue
	//
	HashKeyValue *string `field:"optional" json:"hashKeyValue" yaml:"hashKeyValue"`
	// The action payload.
	//
	// This name can be customized.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-dynamodbaction.html#cfn-iot-topicrule-dynamodbaction-payloadfield
	//
	PayloadField *string `field:"optional" json:"payloadField" yaml:"payloadField"`
	// The range key name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-dynamodbaction.html#cfn-iot-topicrule-dynamodbaction-rangekeyfield
	//
	RangeKeyField *string `field:"optional" json:"rangeKeyField" yaml:"rangeKeyField"`
	// The range key type.
	//
	// Valid values are "STRING" or "NUMBER".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-dynamodbaction.html#cfn-iot-topicrule-dynamodbaction-rangekeytype
	//
	RangeKeyType *string `field:"optional" json:"rangeKeyType" yaml:"rangeKeyType"`
	// The range key value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-dynamodbaction.html#cfn-iot-topicrule-dynamodbaction-rangekeyvalue
	//
	RangeKeyValue *string `field:"optional" json:"rangeKeyValue" yaml:"rangeKeyValue"`
	// The ARN of the IAM role that grants access to the DynamoDB table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-dynamodbaction.html#cfn-iot-topicrule-dynamodbaction-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The name of the DynamoDB table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-dynamodbaction.html#cfn-iot-topicrule-dynamodbaction-tablename
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

