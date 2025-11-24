package mixinsawsiot


// The HTTP action header.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   httpActionHeaderProperty := &HttpActionHeaderProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-httpactionheader.html
//
type CfnTopicRulePropsMixin_HttpActionHeaderProperty struct {
	// The HTTP header key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-httpactionheader.html#cfn-iot-topicrule-httpactionheader-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The HTTP header value.
	//
	// Substitution templates are supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-httpactionheader.html#cfn-iot-topicrule-httpactionheader-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

