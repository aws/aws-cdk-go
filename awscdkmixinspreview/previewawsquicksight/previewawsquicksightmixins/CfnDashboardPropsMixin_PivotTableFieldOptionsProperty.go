package previewawsquicksightmixins


// The field options for a pivot table visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pivotTableFieldOptionsProperty := &PivotTableFieldOptionsProperty{
//   	CollapseStateOptions: []interface{}{
//   		&PivotTableFieldCollapseStateOptionProperty{
//   			State: jsii.String("state"),
//   			Target: &PivotTableFieldCollapseStateTargetProperty{
//   				FieldDataPathValues: []interface{}{
//   					&DataPathValueProperty{
//   						DataPathType: &DataPathTypeProperty{
//   							PivotTableDataPathType: jsii.String("pivotTableDataPathType"),
//   						},
//   						FieldId: jsii.String("fieldId"),
//   						FieldValue: jsii.String("fieldValue"),
//   					},
//   				},
//   				FieldId: jsii.String("fieldId"),
//   			},
//   		},
//   	},
//   	DataPathOptions: []interface{}{
//   		&PivotTableDataPathOptionProperty{
//   			DataPathList: []interface{}{
//   				&DataPathValueProperty{
//   					DataPathType: &DataPathTypeProperty{
//   						PivotTableDataPathType: jsii.String("pivotTableDataPathType"),
//   					},
//   					FieldId: jsii.String("fieldId"),
//   					FieldValue: jsii.String("fieldValue"),
//   				},
//   			},
//   			Width: jsii.String("width"),
//   		},
//   	},
//   	SelectedFieldOptions: []interface{}{
//   		&PivotTableFieldOptionProperty{
//   			CustomLabel: jsii.String("customLabel"),
//   			FieldId: jsii.String("fieldId"),
//   			Visibility: jsii.String("visibility"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablefieldoptions.html
//
type CfnDashboardPropsMixin_PivotTableFieldOptionsProperty struct {
	// The collapse state options for the pivot table field options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablefieldoptions.html#cfn-quicksight-dashboard-pivottablefieldoptions-collapsestateoptions
	//
	CollapseStateOptions interface{} `field:"optional" json:"collapseStateOptions" yaml:"collapseStateOptions"`
	// The data path options for the pivot table field options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablefieldoptions.html#cfn-quicksight-dashboard-pivottablefieldoptions-datapathoptions
	//
	DataPathOptions interface{} `field:"optional" json:"dataPathOptions" yaml:"dataPathOptions"`
	// The selected field options for the pivot table field options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablefieldoptions.html#cfn-quicksight-dashboard-pivottablefieldoptions-selectedfieldoptions
	//
	SelectedFieldOptions interface{} `field:"optional" json:"selectedFieldOptions" yaml:"selectedFieldOptions"`
}

