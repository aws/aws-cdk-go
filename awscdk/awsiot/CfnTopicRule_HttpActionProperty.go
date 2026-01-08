package awsiot


// Send data to an HTTPS endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpActionProperty := &HttpActionProperty{
//   	Url: jsii.String("url"),
//
//   	// the properties below are optional
//   	Auth: &HttpAuthorizationProperty{
//   		Sigv4: &SigV4AuthorizationProperty{
//   			RoleArn: jsii.String("roleArn"),
//   			ServiceName: jsii.String("serviceName"),
//   			SigningRegion: jsii.String("signingRegion"),
//   		},
//   	},
//   	BatchConfig: &BatchConfigProperty{
//   		MaxBatchOpenMs: jsii.Number(123),
//   		MaxBatchSize: jsii.Number(123),
//   		MaxBatchSizeBytes: jsii.Number(123),
//   	},
//   	ConfirmationUrl: jsii.String("confirmationUrl"),
//   	EnableBatching: jsii.Boolean(false),
//   	Headers: []interface{}{
//   		&HttpActionHeaderProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-httpaction.html
//
type CfnTopicRule_HttpActionProperty struct {
	// The endpoint URL.
	//
	// If substitution templates are used in the URL, you must also specify a `confirmationUrl` . If this is a new destination, a new `TopicRuleDestination` is created if possible.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-httpaction.html#cfn-iot-topicrule-httpaction-url
	//
	Url *string `field:"required" json:"url" yaml:"url"`
	// The authentication method to use when sending data to an HTTPS endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-httpaction.html#cfn-iot-topicrule-httpaction-auth
	//
	Auth interface{} `field:"optional" json:"auth" yaml:"auth"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-httpaction.html#cfn-iot-topicrule-httpaction-batchconfig
	//
	BatchConfig interface{} `field:"optional" json:"batchConfig" yaml:"batchConfig"`
	// The URL to which AWS IoT sends a confirmation message.
	//
	// The value of the confirmation URL must be a prefix of the endpoint URL. If you do not specify a confirmation URL AWS IoT uses the endpoint URL as the confirmation URL. If you use substitution templates in the confirmationUrl, you must create and enable topic rule destinations that match each possible value of the substitution template before traffic is allowed to your endpoint URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-httpaction.html#cfn-iot-topicrule-httpaction-confirmationurl
	//
	ConfirmationUrl *string `field:"optional" json:"confirmationUrl" yaml:"confirmationUrl"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-httpaction.html#cfn-iot-topicrule-httpaction-enablebatching
	//
	EnableBatching interface{} `field:"optional" json:"enableBatching" yaml:"enableBatching"`
	// The HTTP headers to send with the message data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-httpaction.html#cfn-iot-topicrule-httpaction-headers
	//
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
}

