package previewawsioteventsmixins


// Information about the variable and its new value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   setVariableProperty := &SetVariableProperty{
//   	Value: jsii.String("value"),
//   	VariableName: jsii.String("variableName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-setvariable.html
//
type CfnDetectorModelPropsMixin_SetVariableProperty struct {
	// The new value of the variable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-setvariable.html#cfn-iotevents-detectormodel-setvariable-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
	// The name of the variable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-setvariable.html#cfn-iotevents-detectormodel-setvariable-variablename
	//
	VariableName *string `field:"optional" json:"variableName" yaml:"variableName"`
}

