package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   geospatialCategoricalDataColorProperty := &GeospatialCategoricalDataColorProperty{
//   	Color: jsii.String("color"),
//   	DataValue: jsii.String("dataValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialcategoricaldatacolor.html
//
type CfnTemplatePropsMixin_GeospatialCategoricalDataColorProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialcategoricaldatacolor.html#cfn-quicksight-template-geospatialcategoricaldatacolor-color
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialcategoricaldatacolor.html#cfn-quicksight-template-geospatialcategoricaldatacolor-datavalue
	//
	DataValue *string `field:"optional" json:"dataValue" yaml:"dataValue"`
}

