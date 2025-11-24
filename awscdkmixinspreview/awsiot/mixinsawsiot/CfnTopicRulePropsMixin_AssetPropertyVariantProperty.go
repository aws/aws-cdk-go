package mixinsawsiot


// Contains an asset property value (of a single type).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   assetPropertyVariantProperty := &AssetPropertyVariantProperty{
//   	BooleanValue: jsii.String("booleanValue"),
//   	DoubleValue: jsii.String("doubleValue"),
//   	IntegerValue: jsii.String("integerValue"),
//   	StringValue: jsii.String("stringValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-assetpropertyvariant.html
//
type CfnTopicRulePropsMixin_AssetPropertyVariantProperty struct {
	// Optional.
	//
	// A string that contains the boolean value ( `true` or `false` ) of the value entry. Accepts substitution templates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-assetpropertyvariant.html#cfn-iot-topicrule-assetpropertyvariant-booleanvalue
	//
	BooleanValue *string `field:"optional" json:"booleanValue" yaml:"booleanValue"`
	// Optional.
	//
	// A string that contains the double value of the value entry. Accepts substitution templates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-assetpropertyvariant.html#cfn-iot-topicrule-assetpropertyvariant-doublevalue
	//
	DoubleValue *string `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// Optional.
	//
	// A string that contains the integer value of the value entry. Accepts substitution templates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-assetpropertyvariant.html#cfn-iot-topicrule-assetpropertyvariant-integervalue
	//
	IntegerValue *string `field:"optional" json:"integerValue" yaml:"integerValue"`
	// Optional.
	//
	// The string value of the value entry. Accepts substitution templates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-assetpropertyvariant.html#cfn-iot-topicrule-assetpropertyvariant-stringvalue
	//
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

