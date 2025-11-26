package previewawsevidentlymixins


// This structure contains the name and variation value of one variation of a feature.
//
// It can contain only one of the following parameters: `BooleanValue` , `DoubleValue` , `LongValue` or `StringValue` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   variationObjectProperty := &VariationObjectProperty{
//   	BooleanValue: jsii.Boolean(false),
//   	DoubleValue: jsii.Number(123),
//   	LongValue: jsii.Number(123),
//   	StringValue: jsii.String("stringValue"),
//   	VariationName: jsii.String("variationName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-feature-variationobject.html
//
type CfnFeaturePropsMixin_VariationObjectProperty struct {
	// The value assigned to this variation, if the variation type is boolean.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-feature-variationobject.html#cfn-evidently-feature-variationobject-booleanvalue
	//
	BooleanValue interface{} `field:"optional" json:"booleanValue" yaml:"booleanValue"`
	// The value assigned to this variation, if the variation type is a double.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-feature-variationobject.html#cfn-evidently-feature-variationobject-doublevalue
	//
	DoubleValue *float64 `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// The value assigned to this variation, if the variation type is a long.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-feature-variationobject.html#cfn-evidently-feature-variationobject-longvalue
	//
	LongValue *float64 `field:"optional" json:"longValue" yaml:"longValue"`
	// The value assigned to this variation, if the variation type is a string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-feature-variationobject.html#cfn-evidently-feature-variationobject-stringvalue
	//
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
	// A name for the variation.
	//
	// It can include up to 127 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-feature-variationobject.html#cfn-evidently-feature-variationobject-variationname
	//
	VariationName *string `field:"optional" json:"variationName" yaml:"variationName"`
}

