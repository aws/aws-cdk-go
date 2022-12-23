package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   republishActionHeadersProperty := &republishActionHeadersProperty{
//   	contentType: jsii.String("contentType"),
//   	correlationData: jsii.String("correlationData"),
//   	messageExpiry: jsii.String("messageExpiry"),
//   	payloadFormatIndicator: jsii.String("payloadFormatIndicator"),
//   	responseTopic: jsii.String("responseTopic"),
//   	userProperties: []interface{}{
//   		&userPropertyProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnTopicRule_RepublishActionHeadersProperty struct {
	// `CfnTopicRule.RepublishActionHeadersProperty.ContentType`.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// `CfnTopicRule.RepublishActionHeadersProperty.CorrelationData`.
	CorrelationData *string `field:"optional" json:"correlationData" yaml:"correlationData"`
	// `CfnTopicRule.RepublishActionHeadersProperty.MessageExpiry`.
	MessageExpiry *string `field:"optional" json:"messageExpiry" yaml:"messageExpiry"`
	// `CfnTopicRule.RepublishActionHeadersProperty.PayloadFormatIndicator`.
	PayloadFormatIndicator *string `field:"optional" json:"payloadFormatIndicator" yaml:"payloadFormatIndicator"`
	// `CfnTopicRule.RepublishActionHeadersProperty.ResponseTopic`.
	ResponseTopic *string `field:"optional" json:"responseTopic" yaml:"responseTopic"`
	// `CfnTopicRule.RepublishActionHeadersProperty.UserProperties`.
	UserProperties interface{} `field:"optional" json:"userProperties" yaml:"userProperties"`
}

