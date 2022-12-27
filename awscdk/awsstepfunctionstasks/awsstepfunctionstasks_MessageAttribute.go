package awsstepfunctionstasks


// A message attribute to add to the SNS message.
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
// See: https://docs.aws.amazon.com/sns/latest/dg/sns-message-attributes.html
//
type MessageAttribute struct {
	// The value of the attribute.
	Value interface{} `field:"required" json:"value" yaml:"value"`
	// The data type for the attribute.
	// See: https://docs.aws.amazon.com/sns/latest/dg/sns-message-attributes.html#SNSMessageAttributes.DataTypes
	//
	DataType MessageAttributeDataType `field:"optional" json:"dataType" yaml:"dataType"`
}

