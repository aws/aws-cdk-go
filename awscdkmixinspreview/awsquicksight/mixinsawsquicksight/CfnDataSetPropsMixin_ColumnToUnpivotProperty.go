package mixinsawsquicksight


// Specifies a column to be unpivoted, transforming it from a column into rows with associated values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   columnToUnpivotProperty := &ColumnToUnpivotProperty{
//   	ColumnName: jsii.String("columnName"),
//   	NewValue: jsii.String("newValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-columntounpivot.html
//
type CfnDataSetPropsMixin_ColumnToUnpivotProperty struct {
	// The name of the column to unpivot from the source data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-columntounpivot.html#cfn-quicksight-dataset-columntounpivot-columnname
	//
	ColumnName *string `field:"optional" json:"columnName" yaml:"columnName"`
	// The value to assign to this column in the unpivoted result, typically the column name or a descriptive label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-columntounpivot.html#cfn-quicksight-dataset-columntounpivot-newvalue
	//
	NewValue *string `field:"optional" json:"newValue" yaml:"newValue"`
}

