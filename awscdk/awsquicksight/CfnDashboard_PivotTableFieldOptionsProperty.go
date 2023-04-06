package awsquicksight


// The field options for a pivot table visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pivotTableFieldOptionsProperty := &PivotTableFieldOptionsProperty{
//   	DataPathOptions: []interface{}{
//   		&PivotTableDataPathOptionProperty{
//   			DataPathList: []interface{}{
//   				&DataPathValueProperty{
//   					FieldId: jsii.String("fieldId"),
//   					FieldValue: jsii.String("fieldValue"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			Width: jsii.String("width"),
//   		},
//   	},
//   	SelectedFieldOptions: []interface{}{
//   		&PivotTableFieldOptionProperty{
//   			FieldId: jsii.String("fieldId"),
//
//   			// the properties below are optional
//   			CustomLabel: jsii.String("customLabel"),
//   			Visibility: jsii.String("visibility"),
//   		},
//   	},
//   }
//
type CfnDashboard_PivotTableFieldOptionsProperty struct {
	// The data path options for the pivot table field options.
	DataPathOptions interface{} `field:"optional" json:"dataPathOptions" yaml:"dataPathOptions"`
	// The selected field options for the pivot table field options.
	SelectedFieldOptions interface{} `field:"optional" json:"selectedFieldOptions" yaml:"selectedFieldOptions"`
}

