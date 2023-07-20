package awsquicksight


// The collapse state options for the pivot table field options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pivotTableFieldCollapseStateOptionProperty := &PivotTableFieldCollapseStateOptionProperty{
//   	Target: &PivotTableFieldCollapseStateTargetProperty{
//   		FieldDataPathValues: []interface{}{
//   			&DataPathValueProperty{
//   				FieldId: jsii.String("fieldId"),
//   				FieldValue: jsii.String("fieldValue"),
//   			},
//   		},
//   		FieldId: jsii.String("fieldId"),
//   	},
//
//   	// the properties below are optional
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablefieldcollapsestateoption.html
//
type CfnDashboard_PivotTableFieldCollapseStateOptionProperty struct {
	// A tagged-union object that sets the collapse state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablefieldcollapsestateoption.html#cfn-quicksight-dashboard-pivottablefieldcollapsestateoption-target
	//
	Target interface{} `field:"required" json:"target" yaml:"target"`
	// The state of the field target of a pivot table. Choose one of the following options:.
	//
	// - `COLLAPSED`
	// - `EXPANDED`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablefieldcollapsestateoption.html#cfn-quicksight-dashboard-pivottablefieldcollapsestateoption-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
}

