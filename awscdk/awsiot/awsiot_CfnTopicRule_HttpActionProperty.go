package awsiot


// Send data to an HTTPS endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpActionProperty := &httpActionProperty{
//   	url: jsii.String("url"),
//
//   	// the properties below are optional
//   	auth: &httpAuthorizationProperty{
//   		sigv4: &sigV4AuthorizationProperty{
//   			roleArn: jsii.String("roleArn"),
//   			serviceName: jsii.String("serviceName"),
//   			signingRegion: jsii.String("signingRegion"),
//   		},
//   	},
//   	confirmationUrl: jsii.String("confirmationUrl"),
//   	headers: []interface{}{
//   		&httpActionHeaderProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnTopicRule_HttpActionProperty struct {
	// The endpoint URL.
	//
	// If substitution templates are used in the URL, you must also specify a `confirmationUrl` . If this is a new destination, a new `TopicRuleDestination` is created if possible.
	Url *string `field:"required" json:"url" yaml:"url"`
	// The authentication method to use when sending data to an HTTPS endpoint.
	Auth interface{} `field:"optional" json:"auth" yaml:"auth"`
	// The URL to which AWS IoT sends a confirmation message.
	//
	// The value of the confirmation URL must be a prefix of the endpoint URL. If you do not specify a confirmation URL AWS IoT uses the endpoint URL as the confirmation URL. If you use substitution templates in the confirmationUrl, you must create and enable topic rule destinations that match each possible value of the substitution template before traffic is allowed to your endpoint URL.
	ConfirmationUrl *string `field:"optional" json:"confirmationUrl" yaml:"confirmationUrl"`
	// The HTTP headers to send with the message data.
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
}

