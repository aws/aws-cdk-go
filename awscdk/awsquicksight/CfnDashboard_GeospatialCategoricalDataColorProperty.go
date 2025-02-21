package awsquicksight


// The categorical data color for a single category.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geospatialCategoricalDataColorProperty := &GeospatialCategoricalDataColorProperty{
//   	Color: jsii.String("color"),
//   	DataValue: jsii.String("dataValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialcategoricaldatacolor.html
//
type CfnDashboard_GeospatialCategoricalDataColorProperty struct {
	// The color and opacity values for the category data color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialcategoricaldatacolor.html#cfn-quicksight-dashboard-geospatialcategoricaldatacolor-color
	//
	Color *string `field:"required" json:"color" yaml:"color"`
	// The data value for the category data color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialcategoricaldatacolor.html#cfn-quicksight-dashboard-geospatialcategoricaldatacolor-datavalue
	//
	DataValue *string `field:"required" json:"dataValue" yaml:"dataValue"`
}

