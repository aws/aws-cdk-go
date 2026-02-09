package previewawsiotmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   commandParameterValueNumberRangeProperty := &CommandParameterValueNumberRangeProperty{
//   	Max: jsii.String("max"),
//   	Min: jsii.String("min"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-commandparametervaluenumberrange.html
//
type CfnCommandPropsMixin_CommandParameterValueNumberRangeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-commandparametervaluenumberrange.html#cfn-iot-command-commandparametervaluenumberrange-max
	//
	Max *string `field:"optional" json:"max" yaml:"max"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-commandparametervaluenumberrange.html#cfn-iot-command-commandparametervaluenumberrange-min
	//
	Min *string `field:"optional" json:"min" yaml:"min"`
}

