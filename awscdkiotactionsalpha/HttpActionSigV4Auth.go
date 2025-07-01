package awscdkiotactionsalpha


// Example:
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
//   	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, year, month, day FROM 'device/+/data'")),
//   })
//
//   topicRule.AddAction(
//   actions.NewHttpsAction(jsii.String("https://example.com/endpoint"), &HttpsActionProps{
//   	ConfirmationUrl: jsii.String("https://example.com"),
//   	Headers: []httpActionHeader{
//   		&httpActionHeader{
//   			Key: jsii.String("key0"),
//   			Value: jsii.String("value0"),
//   		},
//   		&httpActionHeader{
//   			Key: jsii.String("key1"),
//   			Value: jsii.String("value1"),
//   		},
//   	},
//   	Auth: &HttpActionSigV4Auth{
//   		ServiceName: jsii.String("serviceName"),
//   		SigningRegion: jsii.String("us-east-1"),
//   	},
//   }))
//
// Experimental.
type HttpActionSigV4Auth struct {
	// The service name.
	// Experimental.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// The signing region.
	// Experimental.
	SigningRegion *string `field:"required" json:"signingRegion" yaml:"signingRegion"`
}

