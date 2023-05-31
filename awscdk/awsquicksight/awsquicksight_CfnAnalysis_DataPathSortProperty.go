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
//   			FieldId: jsii.String("fieldId"),
//   			FieldValue: jsii.String("fieldValue"),
//   		},
//   	},
//   }
//
type CfnAnalysis_DataPathSortProperty struct {
	// Determines the sort direction.
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// The list of data paths that need to be sorted.
	SortPaths interface{} `field:"required" json:"sortPaths" yaml:"sortPaths"`
}

