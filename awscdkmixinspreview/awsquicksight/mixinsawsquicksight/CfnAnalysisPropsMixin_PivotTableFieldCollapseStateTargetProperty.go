package mixinsawsquicksight


// The target of a pivot table field collapse state.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pivotTableFieldCollapseStateTargetProperty := &PivotTableFieldCollapseStateTargetProperty{
//   	FieldDataPathValues: []interface{}{
//   		&DataPathValueProperty{
//   			DataPathType: &DataPathTypeProperty{
//   				PivotTableDataPathType: jsii.String("pivotTableDataPathType"),
//   			},
//   			FieldId: jsii.String("fieldId"),
//   			FieldValue: jsii.String("fieldValue"),
//   		},
//   	},
//   	FieldId: jsii.String("fieldId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottablefieldcollapsestatetarget.html
//
type CfnAnalysisPropsMixin_PivotTableFieldCollapseStateTargetProperty struct {
	// The data path of the pivot table's header.
	//
	// Used to set the collapse state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottablefieldcollapsestatetarget.html#cfn-quicksight-analysis-pivottablefieldcollapsestatetarget-fielddatapathvalues
	//
	FieldDataPathValues interface{} `field:"optional" json:"fieldDataPathValues" yaml:"fieldDataPathValues"`
	// The field ID of the pivot table that the collapse state needs to be set to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottablefieldcollapsestatetarget.html#cfn-quicksight-analysis-pivottablefieldcollapsestatetarget-fieldid
	//
	FieldId *string `field:"optional" json:"fieldId" yaml:"fieldId"`
}

