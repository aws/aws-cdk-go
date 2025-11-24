package mixinsawsquicksight


// The collapse state options for the pivot table field options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pivotTableFieldCollapseStateOptionProperty := &PivotTableFieldCollapseStateOptionProperty{
//   	State: jsii.String("state"),
//   	Target: &PivotTableFieldCollapseStateTargetProperty{
//   		FieldDataPathValues: []interface{}{
//   			&DataPathValueProperty{
//   				DataPathType: &DataPathTypeProperty{
//   					PivotTableDataPathType: jsii.String("pivotTableDataPathType"),
//   				},
//   				FieldId: jsii.String("fieldId"),
//   				FieldValue: jsii.String("fieldValue"),
//   			},
//   		},
//   		FieldId: jsii.String("fieldId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottablefieldcollapsestateoption.html
//
type CfnAnalysisPropsMixin_PivotTableFieldCollapseStateOptionProperty struct {
	// The state of the field target of a pivot table. Choose one of the following options:.
	//
	// - `COLLAPSED`
	// - `EXPANDED`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottablefieldcollapsestateoption.html#cfn-quicksight-analysis-pivottablefieldcollapsestateoption-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
	// A tagged-union object that sets the collapse state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottablefieldcollapsestateoption.html#cfn-quicksight-analysis-pivottablefieldcollapsestateoption-target
	//
	Target interface{} `field:"optional" json:"target" yaml:"target"`
}

