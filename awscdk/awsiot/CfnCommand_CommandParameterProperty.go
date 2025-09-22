package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   commandParameterProperty := &CommandParameterProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	DefaultValue: &CommandParameterValueProperty{
//   		B: jsii.Boolean(false),
//   		Bin: jsii.String("bin"),
//   		D: jsii.Number(123),
//   		I: jsii.Number(123),
//   		L: jsii.String("l"),
//   		S: jsii.String("s"),
//   		Ul: jsii.String("ul"),
//   	},
//   	Description: jsii.String("description"),
//   	Value: &CommandParameterValueProperty{
//   		B: jsii.Boolean(false),
//   		Bin: jsii.String("bin"),
//   		D: jsii.Number(123),
//   		I: jsii.Number(123),
//   		L: jsii.String("l"),
//   		S: jsii.String("s"),
//   		Ul: jsii.String("ul"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-commandparameter.html
//
type CfnCommand_CommandParameterProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-commandparameter.html#cfn-iot-command-commandparameter-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-commandparameter.html#cfn-iot-command-commandparameter-defaultvalue
	//
	DefaultValue interface{} `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-commandparameter.html#cfn-iot-command-commandparameter-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-commandparameter.html#cfn-iot-command-commandparameter-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

