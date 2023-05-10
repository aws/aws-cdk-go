package awsquicksight


// The data path options for the pivot table field options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pivotTableDataPathOptionProperty := &PivotTableDataPathOptionProperty{
//   	DataPathList: []interface{}{
//   		&DataPathValueProperty{
//   			FieldId: jsii.String("fieldId"),
//   			FieldValue: jsii.String("fieldValue"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Width: jsii.String("width"),
//   }
//
type CfnTemplate_PivotTableDataPathOptionProperty struct {
	// The list of data path values for the data path options.
	DataPathList interface{} `field:"required" json:"dataPathList" yaml:"dataPathList"`
	// The width of the data path option.
	Width *string `field:"optional" json:"width" yaml:"width"`
}

