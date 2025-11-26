package previewawsquicksightmixins


// The color map that determines the color options for a particular element.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataPathColorProperty := &DataPathColorProperty{
//   	Color: jsii.String("color"),
//   	Element: &DataPathValueProperty{
//   		DataPathType: &DataPathTypeProperty{
//   			PivotTableDataPathType: jsii.String("pivotTableDataPathType"),
//   		},
//   		FieldId: jsii.String("fieldId"),
//   		FieldValue: jsii.String("fieldValue"),
//   	},
//   	TimeGranularity: jsii.String("timeGranularity"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datapathcolor.html
//
type CfnDashboardPropsMixin_DataPathColorProperty struct {
	// The color that needs to be applied to the element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datapathcolor.html#cfn-quicksight-dashboard-datapathcolor-color
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
	// The element that the color needs to be applied to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datapathcolor.html#cfn-quicksight-dashboard-datapathcolor-element
	//
	Element interface{} `field:"optional" json:"element" yaml:"element"`
	// The time granularity of the field that the color needs to be applied to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datapathcolor.html#cfn-quicksight-dashboard-datapathcolor-timegranularity
	//
	TimeGranularity *string `field:"optional" json:"timeGranularity" yaml:"timeGranularity"`
}

