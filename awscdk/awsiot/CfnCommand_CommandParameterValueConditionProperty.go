package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   commandParameterValueConditionProperty := &CommandParameterValueConditionProperty{
//   	ComparisonOperator: jsii.String("comparisonOperator"),
//   	Operand: &CommandParameterValueComparisonOperandProperty{
//   		Number: jsii.String("number"),
//   		NumberRange: &CommandParameterValueNumberRangeProperty{
//   			Max: jsii.String("max"),
//   			Min: jsii.String("min"),
//   		},
//   		Numbers: []*string{
//   			jsii.String("numbers"),
//   		},
//   		String: jsii.String("string"),
//   		Strings: []*string{
//   			jsii.String("strings"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-commandparametervaluecondition.html
//
type CfnCommand_CommandParameterValueConditionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-commandparametervaluecondition.html#cfn-iot-command-commandparametervaluecondition-comparisonoperator
	//
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-commandparametervaluecondition.html#cfn-iot-command-commandparametervaluecondition-operand
	//
	Operand interface{} `field:"required" json:"operand" yaml:"operand"`
}

