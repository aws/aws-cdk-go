package previewawsiotmixins


// A key-value pair that you define in the header.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   userPropertyProperty := &UserPropertyProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-userproperty.html
//
type CfnTopicRulePropsMixin_UserPropertyProperty struct {
	// A key to be specified in `UserProperty` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-userproperty.html#cfn-iot-topicrule-userproperty-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A value to be specified in `UserProperty` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-userproperty.html#cfn-iot-topicrule-userproperty-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

