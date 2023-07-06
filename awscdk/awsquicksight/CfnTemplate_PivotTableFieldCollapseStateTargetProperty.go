package awsquicksight


// The target of a pivot table field collapse state.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pivotTableFieldCollapseStateTargetProperty := &PivotTableFieldCollapseStateTargetProperty{
//   	FieldDataPathValues: []interface{}{
//   		&DataPathValueProperty{
//   			FieldId: jsii.String("fieldId"),
//   			FieldValue: jsii.String("fieldValue"),
//   		},
//   	},
//   	FieldId: jsii.String("fieldId"),
//   }
//
type CfnTemplate_PivotTableFieldCollapseStateTargetProperty struct {
	// The data path of the pivot table's header.
	//
	// Used to set the collapse state.
	FieldDataPathValues interface{} `field:"optional" json:"fieldDataPathValues" yaml:"fieldDataPathValues"`
	// The field ID of the pivot table that the collapse state needs to be set to.
	FieldId *string `field:"optional" json:"fieldId" yaml:"fieldId"`
}

