package awsquicksight


// The gradient step color for a single step.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geospatialGradientStepColorProperty := &GeospatialGradientStepColorProperty{
//   	Color: jsii.String("color"),
//   	DataValue: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialgradientstepcolor.html
//
type CfnAnalysis_GeospatialGradientStepColorProperty struct {
	// The color and opacity values for the gradient step color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialgradientstepcolor.html#cfn-quicksight-analysis-geospatialgradientstepcolor-color
	//
	Color *string `field:"required" json:"color" yaml:"color"`
	// The data value for the gradient step color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialgradientstepcolor.html#cfn-quicksight-analysis-geospatialgradientstepcolor-datavalue
	//
	// Default: - 0.
	//
	DataValue *float64 `field:"required" json:"dataValue" yaml:"dataValue"`
}

