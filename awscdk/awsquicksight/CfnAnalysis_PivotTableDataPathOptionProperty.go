package awsquicksight


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
type CfnAnalysis_PivotTableDataPathOptionProperty struct {
	// `CfnAnalysis.PivotTableDataPathOptionProperty.DataPathList`.
	DataPathList interface{} `field:"required" json:"dataPathList" yaml:"dataPathList"`
	// `CfnAnalysis.PivotTableDataPathOptionProperty.Width`.
	Width *string `field:"optional" json:"width" yaml:"width"`
}

