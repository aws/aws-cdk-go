package previewawsiotmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   commandParameterValueComparisonOperandProperty := &CommandParameterValueComparisonOperandProperty{
//   	Number: jsii.String("number"),
//   	NumberRange: &CommandParameterValueNumberRangeProperty{
//   		Max: jsii.String("max"),
//   		Min: jsii.String("min"),
//   	},
//   	Numbers: []*string{
//   		jsii.String("numbers"),
//   	},
//   	String: jsii.String("string"),
//   	Strings: []*string{
//   		jsii.String("strings"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-commandparametervaluecomparisonoperand.html
//
type CfnCommandPropsMixin_CommandParameterValueComparisonOperandProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-commandparametervaluecomparisonoperand.html#cfn-iot-command-commandparametervaluecomparisonoperand-number
	//
	Number *string `field:"optional" json:"number" yaml:"number"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-commandparametervaluecomparisonoperand.html#cfn-iot-command-commandparametervaluecomparisonoperand-numberrange
	//
	NumberRange interface{} `field:"optional" json:"numberRange" yaml:"numberRange"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-commandparametervaluecomparisonoperand.html#cfn-iot-command-commandparametervaluecomparisonoperand-numbers
	//
	Numbers *[]*string `field:"optional" json:"numbers" yaml:"numbers"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-commandparametervaluecomparisonoperand.html#cfn-iot-command-commandparametervaluecomparisonoperand-string
	//
	String *string `field:"optional" json:"string" yaml:"string"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-commandparametervaluecomparisonoperand.html#cfn-iot-command-commandparametervaluecomparisonoperand-strings
	//
	Strings *[]*string `field:"optional" json:"strings" yaml:"strings"`
}

