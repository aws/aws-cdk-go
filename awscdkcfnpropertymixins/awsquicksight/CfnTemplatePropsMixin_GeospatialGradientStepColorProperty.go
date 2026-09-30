package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   geospatialGradientStepColorProperty := &GeospatialGradientStepColorProperty{
//   	Color: jsii.String("color"),
//   	DataValue: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialgradientstepcolor.html
//
type CfnTemplatePropsMixin_GeospatialGradientStepColorProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialgradientstepcolor.html#cfn-quicksight-template-geospatialgradientstepcolor-color
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialgradientstepcolor.html#cfn-quicksight-template-geospatialgradientstepcolor-datavalue
	//
	// Default: - 0.
	//
	DataValue *float64 `field:"optional" json:"dataValue" yaml:"dataValue"`
}

