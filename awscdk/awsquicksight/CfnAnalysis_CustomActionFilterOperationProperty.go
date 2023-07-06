package awsquicksight


// The filter operation that filters data included in a visual or in an entire sheet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customActionFilterOperationProperty := &CustomActionFilterOperationProperty{
//   	SelectedFieldsConfiguration: &FilterOperationSelectedFieldsConfigurationProperty{
//   		SelectedColumns: []interface{}{
//   			&ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   		},
//   		SelectedFieldOptions: jsii.String("selectedFieldOptions"),
//   		SelectedFields: []*string{
//   			jsii.String("selectedFields"),
//   		},
//   	},
//   	TargetVisualsConfiguration: &FilterOperationTargetVisualsConfigurationProperty{
//   		SameSheetTargetVisualConfiguration: &SameSheetTargetVisualConfigurationProperty{
//   			TargetVisualOptions: jsii.String("targetVisualOptions"),
//   			TargetVisuals: []*string{
//   				jsii.String("targetVisuals"),
//   			},
//   		},
//   	},
//   }
//
type CfnAnalysis_CustomActionFilterOperationProperty struct {
	// The configuration that chooses the fields to be filtered.
	SelectedFieldsConfiguration interface{} `field:"required" json:"selectedFieldsConfiguration" yaml:"selectedFieldsConfiguration"`
	// The configuration that chooses the target visuals to be filtered.
	TargetVisualsConfiguration interface{} `field:"required" json:"targetVisualsConfiguration" yaml:"targetVisualsConfiguration"`
}

