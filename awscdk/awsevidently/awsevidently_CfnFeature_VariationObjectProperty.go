package awsevidently


// This structure contains the name and variation value of one variation of a feature.
//
// It can contain only one of the following parameters: `BooleanValue` , `DoubleValue` , `LongValue` or `StringValue` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   variationObjectProperty := &variationObjectProperty{
//   	variationName: jsii.String("variationName"),
//
//   	// the properties below are optional
//   	booleanValue: jsii.Boolean(false),
//   	doubleValue: jsii.Number(123),
//   	longValue: jsii.Number(123),
//   	stringValue: jsii.String("stringValue"),
//   }
//
type CfnFeature_VariationObjectProperty struct {
	// A name for the variation.
	//
	// It can include up to 127 characters.
	VariationName *string `field:"required" json:"variationName" yaml:"variationName"`
	// The value assigned to this variation, if the variation type is boolean.
	BooleanValue interface{} `field:"optional" json:"booleanValue" yaml:"booleanValue"`
	// The value assigned to this variation, if the variation type is a double.
	DoubleValue *float64 `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// The value assigned to this variation, if the variation type is a long.
	LongValue *float64 `field:"optional" json:"longValue" yaml:"longValue"`
	// The value assigned to this variation, if the variation type is a string.
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

