package awscdkiotactionsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Configuration properties of an HTTPS action.
//
// Example:
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
//   	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, year, month, day FROM 'device/+/data'")),
//   })
//
//   topicRule.AddAction(
//   actions.NewHttpsAction(jsii.String("https://example.com/endpoint"), &HttpsActionProps{
//   	ConfirmationUrl: jsii.String("https://example.com"),
//   	Headers: []HttpActionHeader{
//   		&HttpActionHeader{
//   			Key: jsii.String("key0"),
//   			Value: jsii.String("value0"),
//   		},
//   		&HttpActionHeader{
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
// See: https://docs.aws.amazon.com/iot/latest/developerguide/https-rule-action.html
//
// Experimental.
type HttpsActionProps struct {
	// The IAM role that allows access to AWS service.
	// Default: a new role will be created.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Use Sigv4 authorization.
	// Experimental.
	Auth *HttpActionSigV4Auth `field:"optional" json:"auth" yaml:"auth"`
	// If specified, AWS IoT uses the confirmation URL to create a matching topic rule destination.
	// Experimental.
	ConfirmationUrl *string `field:"optional" json:"confirmationUrl" yaml:"confirmationUrl"`
	// The headers to include in the HTTPS request to the endpoint.
	// Experimental.
	Headers *[]*HttpActionHeader `field:"optional" json:"headers" yaml:"headers"`
}

