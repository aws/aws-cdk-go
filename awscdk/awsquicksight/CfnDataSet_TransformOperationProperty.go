package awsquicksight


// A data transformation on a logical table.
//
// This is a variant type structure. For this structure to be valid, only one of the attributes can be non-null.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transformOperationProperty := &TransformOperationProperty{
//   	CastColumnTypeOperation: &CastColumnTypeOperationProperty{
//   		ColumnName: jsii.String("columnName"),
//   		NewColumnType: jsii.String("newColumnType"),
//
//   		// the properties below are optional
//   		Format: jsii.String("format"),
//   	},
//   	CreateColumnsOperation: &CreateColumnsOperationProperty{
//   		Columns: []interface{}{
//   			&CalculatedColumnProperty{
//   				ColumnId: jsii.String("columnId"),
//   				ColumnName: jsii.String("columnName"),
//   				Expression: jsii.String("expression"),
//   			},
//   		},
//   	},
//   	FilterOperation: &FilterOperationProperty{
//   		ConditionExpression: jsii.String("conditionExpression"),
//   	},
//   	ProjectOperation: &ProjectOperationProperty{
//   		ProjectedColumns: []*string{
//   			jsii.String("projectedColumns"),
//   		},
//   	},
//   	RenameColumnOperation: &RenameColumnOperationProperty{
//   		ColumnName: jsii.String("columnName"),
//   		NewColumnName: jsii.String("newColumnName"),
//   	},
//   	TagColumnOperation: &TagColumnOperationProperty{
//   		ColumnName: jsii.String("columnName"),
//   		Tags: []columnTagProperty{
//   			&columnTagProperty{
//   				ColumnDescription: &ColumnDescriptionProperty{
//   					Text: jsii.String("text"),
//   				},
//   				ColumnGeographicRole: jsii.String("columnGeographicRole"),
//   			},
//   		},
//   	},
//   }
//
type CfnDataSet_TransformOperationProperty struct {
	// A transform operation that casts a column to a different type.
	CastColumnTypeOperation interface{} `field:"optional" json:"castColumnTypeOperation" yaml:"castColumnTypeOperation"`
	// An operation that creates calculated columns.
	//
	// Columns created in one such operation form a lexical closure.
	CreateColumnsOperation interface{} `field:"optional" json:"createColumnsOperation" yaml:"createColumnsOperation"`
	// An operation that filters rows based on some condition.
	FilterOperation interface{} `field:"optional" json:"filterOperation" yaml:"filterOperation"`
	// An operation that projects columns.
	//
	// Operations that come after a projection can only refer to projected columns.
	ProjectOperation interface{} `field:"optional" json:"projectOperation" yaml:"projectOperation"`
	// An operation that renames a column.
	RenameColumnOperation interface{} `field:"optional" json:"renameColumnOperation" yaml:"renameColumnOperation"`
	// An operation that tags a column with additional information.
	TagColumnOperation interface{} `field:"optional" json:"tagColumnOperation" yaml:"tagColumnOperation"`
}

