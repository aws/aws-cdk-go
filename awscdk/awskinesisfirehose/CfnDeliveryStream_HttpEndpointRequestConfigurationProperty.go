package awskinesisfirehose


// The configuration of the HTTP endpoint request.
//
// Kinesis Firehose supports any custom HTTP endpoint or HTTP endpoints owned by supported third-party service providers, including Datadog, MongoDB, and New Relic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpEndpointRequestConfigurationProperty := &HttpEndpointRequestConfigurationProperty{
//   	CommonAttributes: []interface{}{
//   		&HttpEndpointCommonAttributeProperty{
//   			AttributeName: jsii.String("attributeName"),
//   			AttributeValue: jsii.String("attributeValue"),
//   		},
//   	},
//   	ContentEncoding: jsii.String("contentEncoding"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointrequestconfiguration.html
//
type CfnDeliveryStream_HttpEndpointRequestConfigurationProperty struct {
	// Describes the metadata sent to the HTTP endpoint destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointrequestconfiguration.html#cfn-kinesisfirehose-deliverystream-httpendpointrequestconfiguration-commonattributes
	//
	CommonAttributes interface{} `field:"optional" json:"commonAttributes" yaml:"commonAttributes"`
	// Kinesis Data Firehose uses the content encoding to compress the body of a request before sending the request to the destination.
	//
	// For more information, see Content-Encoding in MDN Web Docs, the official Mozilla documentation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointrequestconfiguration.html#cfn-kinesisfirehose-deliverystream-httpendpointrequestconfiguration-contentencoding
	//
	ContentEncoding *string `field:"optional" json:"contentEncoding" yaml:"contentEncoding"`
}

