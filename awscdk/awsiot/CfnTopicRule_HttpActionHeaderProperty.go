package awsiot


// The HTTP action header.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpActionHeaderProperty := &HttpActionHeaderProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-httpactionheader.html
//
type CfnTopicRule_HttpActionHeaderProperty struct {
	// The HTTP header key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-httpactionheader.html#cfn-iot-topicrule-httpactionheader-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The HTTP header value.
	//
	// Substitution templates are supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-httpactionheader.html#cfn-iot-topicrule-httpactionheader-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

