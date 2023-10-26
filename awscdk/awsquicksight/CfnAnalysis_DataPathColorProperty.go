package awsquicksight


// The color map that determines the color options for a particular element.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataPathColorProperty := &DataPathColorProperty{
//   	Color: jsii.String("color"),
//   	Element: &DataPathValueProperty{
//   		FieldId: jsii.String("fieldId"),
//   		FieldValue: jsii.String("fieldValue"),
//
//   		// the properties below are optional
//   		DataPathType: &DataPathTypeProperty{
//   			PivotTableDataPathType: jsii.String("pivotTableDataPathType"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	TimeGranularity: jsii.String("timeGranularity"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datapathcolor.html
//
type CfnAnalysis_DataPathColorProperty struct {
	// The color that needs to be applied to the element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datapathcolor.html#cfn-quicksight-analysis-datapathcolor-color
	//
	Color *string `field:"required" json:"color" yaml:"color"`
	// The element that the color needs to be applied to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datapathcolor.html#cfn-quicksight-analysis-datapathcolor-element
	//
	Element interface{} `field:"required" json:"element" yaml:"element"`
	// The time granularity of the field that the color needs to be applied to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datapathcolor.html#cfn-quicksight-analysis-datapathcolor-timegranularity
	//
	TimeGranularity *string `field:"optional" json:"timeGranularity" yaml:"timeGranularity"`
}

