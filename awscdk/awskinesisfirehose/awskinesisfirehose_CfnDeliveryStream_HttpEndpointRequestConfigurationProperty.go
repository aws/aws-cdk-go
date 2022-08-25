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
//   httpEndpointRequestConfigurationProperty := &httpEndpointRequestConfigurationProperty{
//   	commonAttributes: []interface{}{
//   		&httpEndpointCommonAttributeProperty{
//   			attributeName: jsii.String("attributeName"),
//   			attributeValue: jsii.String("attributeValue"),
//   		},
//   	},
//   	contentEncoding: jsii.String("contentEncoding"),
//   }
//
type CfnDeliveryStream_HttpEndpointRequestConfigurationProperty struct {
	// Describes the metadata sent to the HTTP endpoint destination.
	CommonAttributes interface{} `field:"optional" json:"commonAttributes" yaml:"commonAttributes"`
	// Kinesis Data Firehose uses the content encoding to compress the body of a request before sending the request to the destination.
	//
	// For more information, see Content-Encoding in MDN Web Docs, the official Mozilla documentation.
	ContentEncoding *string `field:"optional" json:"contentEncoding" yaml:"contentEncoding"`
}

