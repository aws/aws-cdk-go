package awsstepfunctionstasks


// The data type set for the SNS message attributes.
//
// Example:
//   topic := sns.NewTopic(this, jsii.String("Topic"))
//
//   // Use a field from the execution data as message.
//   task1 := tasks.NewSnsPublish(this, jsii.String("Publish1"), &snsPublishProps{
//   	topic: topic,
//   	integrationPattern: sfn.integrationPattern_REQUEST_RESPONSE,
//   	message: sfn.taskInput.fromDataAt(jsii.String("$.state.message")),
//   	messageAttributes: map[string]messageAttribute{
//   		"place": &messageAttribute{
//   			"value": sfn.JsonPath.stringAt(jsii.String("$.place")),
//   		},
//   		"pic": &messageAttribute{
//   			// BINARY must be explicitly set
//   			"dataType": tasks.MessageAttributeDataType_BINARY,
//   			"value": sfn.JsonPath.stringAt(jsii.String("$.pic")),
//   		},
//   		"people": &messageAttribute{
//   			"value": jsii.Number(4),
//   		},
//   		"handles": &messageAttribute{
//   			"value": []interface{}{
//   				jsii.String("@kslater"),
//   				jsii.String("@jjf"),
//   				nil,
//   				jsii.String("@mfanning"),
//   			},
//   		},
//   	},
//   })
//
//   // Combine a field from the execution data with
//   // a literal object.
//   task2 := tasks.NewSnsPublish(this, jsii.String("Publish2"), &snsPublishProps{
//   	topic: topic,
//   	message: sfn.*taskInput.fromObject(map[string]interface{}{
//   		"field1": jsii.String("somedata"),
//   		"field2": sfn.JsonPath.stringAt(jsii.String("$.field2")),
//   	}),
//   })
//
// See: https://docs.aws.amazon.com/sns/latest/dg/sns-message-attributes.html#SNSMessageAttributes.DataTypes
//
type MessageAttributeDataType string

const (
	// Strings are Unicode with UTF-8 binary encoding.
	MessageAttributeDataType_STRING MessageAttributeDataType = "STRING"
	// An array, formatted as a string.
	// See: https://docs.aws.amazon.com/sns/latest/dg/sns-message-attributes.html#SNSMessageAttributes.DataTypes
	//
	MessageAttributeDataType_STRING_ARRAY MessageAttributeDataType = "STRING_ARRAY"
	// Numbers are positive or negative integers or floating-point numbers.
	// See: https://docs.aws.amazon.com/sns/latest/dg/sns-message-attributes.html#SNSMessageAttributes.DataTypes
	//
	MessageAttributeDataType_NUMBER MessageAttributeDataType = "NUMBER"
	// Binary type attributes can store any binary data.
	// See: https://docs.aws.amazon.com/sns/latest/dg/sns-message-attributes.html#SNSMessageAttributes.DataTypes
	//
	MessageAttributeDataType_BINARY MessageAttributeDataType = "BINARY"
)

