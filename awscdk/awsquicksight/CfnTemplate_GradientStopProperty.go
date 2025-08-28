package awsquicksight


// Determines the gradient stop configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gradientStopProperty := &GradientStopProperty{
//   	GradientOffset: jsii.Number(123),
//
//   	// the properties below are optional
//   	Color: jsii.String("color"),
//   	DataValue: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gradientstop.html
//
type CfnTemplate_GradientStopProperty struct {
	// Determines gradient offset value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gradientstop.html#cfn-quicksight-template-gradientstop-gradientoffset
	//
	// Default: - 0.
	//
	GradientOffset *float64 `field:"required" json:"gradientOffset" yaml:"gradientOffset"`
	// Determines the color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gradientstop.html#cfn-quicksight-template-gradientstop-color
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
	// Determines the data value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gradientstop.html#cfn-quicksight-template-gradientstop-datavalue
	//
	DataValue *float64 `field:"optional" json:"dataValue" yaml:"dataValue"`
}

