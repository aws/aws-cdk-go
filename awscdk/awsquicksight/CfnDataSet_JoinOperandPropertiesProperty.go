package awsquicksight


// Properties that control how columns are handled for a join operand, including column name overrides.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   joinOperandPropertiesProperty := &JoinOperandPropertiesProperty{
//   	OutputColumnNameOverrides: []interface{}{
//   		&OutputColumnNameOverrideProperty{
//   			OutputColumnName: jsii.String("outputColumnName"),
//
//   			// the properties below are optional
//   			SourceColumnName: jsii.String("sourceColumnName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joinoperandproperties.html
//
type CfnDataSet_JoinOperandPropertiesProperty struct {
	// A list of column name overrides to apply to the join operand's output columns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joinoperandproperties.html#cfn-quicksight-dataset-joinoperandproperties-outputcolumnnameoverrides
	//
	OutputColumnNameOverrides interface{} `field:"required" json:"outputColumnNameOverrides" yaml:"outputColumnNameOverrides"`
}

