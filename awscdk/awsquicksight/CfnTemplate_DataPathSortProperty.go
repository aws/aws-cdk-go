package awsquicksight


// Allows data paths to be sorted by a specific data value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataPathSortProperty := &DataPathSortProperty{
//   	Direction: jsii.String("direction"),
//   	SortPaths: []interface{}{
//   		&DataPathValueProperty{
//   			DataPathType: &DataPathTypeProperty{
//   				PivotTableDataPathType: jsii.String("pivotTableDataPathType"),
//   			},
//   			FieldId: jsii.String("fieldId"),
//   			FieldValue: jsii.String("fieldValue"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datapathsort.html
//
type CfnTemplate_DataPathSortProperty struct {
	// Determines the sort direction.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datapathsort.html#cfn-quicksight-template-datapathsort-direction
	//
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// The list of data paths that need to be sorted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datapathsort.html#cfn-quicksight-template-datapathsort-sortpaths
	//
	SortPaths interface{} `field:"required" json:"sortPaths" yaml:"sortPaths"`
}

