package awsstepfunctionstasks


// A message attribute to add to the SNS message.
//
// Example:
//   topic := sns.NewTopic(this, jsii.String("Topic"))
//
//   // Use a field from the execution data as message.
//   task1 := tasks.NewSnsPublish(this, jsii.String("Publish1"), &SnsPublishProps{
//   	Topic: Topic,
//   	IntegrationPattern: sfn.IntegrationPattern_REQUEST_RESPONSE,
//   	Message: sfn.TaskInput_FromDataAt(jsii.String("$.state.message")),
//   	MessageAttributes: map[string]MessageAttribute{
//   		"place": &MessageAttribute{
//   			"value": sfn.JsonPath_stringAt(jsii.String("$.place")),
//   		},
//   		"pic": &MessageAttribute{
//   			// BINARY must be explicitly set
//   			"dataType": tasks.MessageAttributeDataType_BINARY,
//   			"value": sfn.JsonPath_stringAt(jsii.String("$.pic")),
//   		},
//   		"people": &MessageAttribute{
//   			"value": jsii.Number(4),
//   		},
//   		"handles": &MessageAttribute{
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
//   task2 := tasks.NewSnsPublish(this, jsii.String("Publish2"), &SnsPublishProps{
//   	Topic: Topic,
//   	Message: sfn.TaskInput_FromObject(map[string]interface{}{
//   		"field1": jsii.String("somedata"),
//   		"field2": sfn.JsonPath_stringAt(jsii.String("$.field2")),
//   	}),
//   })
//
// See: https://docs.aws.amazon.com/sns/latest/dg/sns-message-attributes.html
//
type MessageAttribute struct {
	// The value of the attribute.
	Value interface{} `field:"required" json:"value" yaml:"value"`
	// The data type for the attribute.
	// See: https://docs.aws.amazon.com/sns/latest/dg/sns-message-attributes.html#SNSMessageAttributes.DataTypes
	//
	// Default: determined by type inspection if possible, fallback is String.
	//
	DataType MessageAttributeDataType `field:"optional" json:"dataType" yaml:"dataType"`
}

