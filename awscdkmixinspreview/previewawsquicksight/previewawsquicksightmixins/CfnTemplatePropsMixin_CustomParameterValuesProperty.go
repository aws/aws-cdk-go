package previewawsquicksightmixins


// The customized parameter values.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customParameterValuesProperty := &CustomParameterValuesProperty{
//   	DateTimeValues: []*string{
//   		jsii.String("dateTimeValues"),
//   	},
//   	DecimalValues: []interface{}{
//   		jsii.Number(123),
//   	},
//   	IntegerValues: []interface{}{
//   		jsii.Number(123),
//   	},
//   	StringValues: []*string{
//   		jsii.String("stringValues"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-customparametervalues.html
//
type CfnTemplatePropsMixin_CustomParameterValuesProperty struct {
	// A list of datetime-type parameter values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-customparametervalues.html#cfn-quicksight-template-customparametervalues-datetimevalues
	//
	DateTimeValues *[]*string `field:"optional" json:"dateTimeValues" yaml:"dateTimeValues"`
	// A list of decimal-type parameter values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-customparametervalues.html#cfn-quicksight-template-customparametervalues-decimalvalues
	//
	DecimalValues interface{} `field:"optional" json:"decimalValues" yaml:"decimalValues"`
	// A list of integer-type parameter values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-customparametervalues.html#cfn-quicksight-template-customparametervalues-integervalues
	//
	IntegerValues interface{} `field:"optional" json:"integerValues" yaml:"integerValues"`
	// A list of string-type parameter values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-customparametervalues.html#cfn-quicksight-template-customparametervalues-stringvalues
	//
	StringValues *[]*string `field:"optional" json:"stringValues" yaml:"stringValues"`
}

