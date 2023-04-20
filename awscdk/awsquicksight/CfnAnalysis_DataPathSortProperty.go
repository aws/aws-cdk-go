package awsquicksight


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
	// `CfnAnalysis.DataPathSortProperty.Direction`.
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// `CfnAnalysis.DataPathSortProperty.SortPaths`.
	SortPaths interface{} `field:"required" json:"sortPaths" yaml:"sortPaths"`
}

