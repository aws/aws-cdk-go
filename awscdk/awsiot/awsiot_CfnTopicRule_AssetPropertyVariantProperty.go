package awsiot


// Contains an asset property value (of a single type).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetPropertyVariantProperty := &assetPropertyVariantProperty{
//   	booleanValue: jsii.String("booleanValue"),
//   	doubleValue: jsii.String("doubleValue"),
//   	integerValue: jsii.String("integerValue"),
//   	stringValue: jsii.String("stringValue"),
//   }
//
type CfnTopicRule_AssetPropertyVariantProperty struct {
	// Optional.
	//
	// A string that contains the boolean value ( `true` or `false` ) of the value entry. Accepts substitution templates.
	BooleanValue *string `field:"optional" json:"booleanValue" yaml:"booleanValue"`
	// Optional.
	//
	// A string that contains the double value of the value entry. Accepts substitution templates.
	DoubleValue *string `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// Optional.
	//
	// A string that contains the integer value of the value entry. Accepts substitution templates.
	IntegerValue *string `field:"optional" json:"integerValue" yaml:"integerValue"`
	// Optional.
	//
	// The string value of the value entry. Accepts substitution templates.
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

