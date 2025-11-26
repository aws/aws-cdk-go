package previewawsiotmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   commandParameterValueProperty := &CommandParameterValueProperty{
//   	B: jsii.Boolean(false),
//   	Bin: jsii.String("bin"),
//   	D: jsii.Number(123),
//   	I: jsii.Number(123),
//   	L: jsii.String("l"),
//   	S: jsii.String("s"),
//   	Ul: jsii.String("ul"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-commandparametervalue.html
//
type CfnCommandPropsMixin_CommandParameterValueProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-commandparametervalue.html#cfn-iot-command-commandparametervalue-b
	//
	B interface{} `field:"optional" json:"b" yaml:"b"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-commandparametervalue.html#cfn-iot-command-commandparametervalue-bin
	//
	Bin *string `field:"optional" json:"bin" yaml:"bin"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-commandparametervalue.html#cfn-iot-command-commandparametervalue-d
	//
	D *float64 `field:"optional" json:"d" yaml:"d"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-commandparametervalue.html#cfn-iot-command-commandparametervalue-i
	//
	I *float64 `field:"optional" json:"i" yaml:"i"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-commandparametervalue.html#cfn-iot-command-commandparametervalue-l
	//
	L *string `field:"optional" json:"l" yaml:"l"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-commandparametervalue.html#cfn-iot-command-commandparametervalue-s
	//
	S *string `field:"optional" json:"s" yaml:"s"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-commandparametervalue.html#cfn-iot-command-commandparametervalue-ul
	//
	Ul *string `field:"optional" json:"ul" yaml:"ul"`
}

