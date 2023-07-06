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
type CfnAnalysis_PivotTableFieldCollapseStateOptionProperty struct {
	// A tagged-union object that sets the collapse state.
	Target interface{} `field:"required" json:"target" yaml:"target"`
	// The state of the field target of a pivot table. Choose one of the following options:.
	//
	// - `COLLAPSED`
	// - `EXPANDED`.
	State *string `field:"optional" json:"state" yaml:"state"`
}

