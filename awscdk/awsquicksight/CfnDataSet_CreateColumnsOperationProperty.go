package awsquicksight


// A transform operation that creates calculated columns.
//
// Columns created in one such operation form a lexical closure.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   createColumnsOperationProperty := &CreateColumnsOperationProperty{
//   	Columns: []interface{}{
//   		&CalculatedColumnProperty{
//   			ColumnId: jsii.String("columnId"),
//   			ColumnName: jsii.String("columnName"),
//   			Expression: jsii.String("expression"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-createcolumnsoperation.html
//
type CfnDataSet_CreateColumnsOperationProperty struct {
	// Calculated columns to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-createcolumnsoperation.html#cfn-quicksight-dataset-createcolumnsoperation-columns
	//
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
}

