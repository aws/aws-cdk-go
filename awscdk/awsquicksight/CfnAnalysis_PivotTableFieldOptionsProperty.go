package awsquicksight


// The field options for a pivot table visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pivotTableFieldOptionsProperty := &PivotTableFieldOptionsProperty{
//   	CollapseStateOptions: []interface{}{
//   		&PivotTableFieldCollapseStateOptionProperty{
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
//
//   			// the properties below are optional
//   			State: jsii.String("state"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottablefieldoptions.html
//
type CfnAnalysis_PivotTableFieldOptionsProperty struct {
	// The collapse state options for the pivot table field options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottablefieldoptions.html#cfn-quicksight-analysis-pivottablefieldoptions-collapsestateoptions
	//
	CollapseStateOptions interface{} `field:"optional" json:"collapseStateOptions" yaml:"collapseStateOptions"`
	// The data path options for the pivot table field options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottablefieldoptions.html#cfn-quicksight-analysis-pivottablefieldoptions-datapathoptions
	//
	DataPathOptions interface{} `field:"optional" json:"dataPathOptions" yaml:"dataPathOptions"`
	// The selected field options for the pivot table field options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottablefieldoptions.html#cfn-quicksight-analysis-pivottablefieldoptions-selectedfieldoptions
	//
	SelectedFieldOptions interface{} `field:"optional" json:"selectedFieldOptions" yaml:"selectedFieldOptions"`
}

