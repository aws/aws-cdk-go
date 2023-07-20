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
//   	OverrideDatasetParameterOperation: &OverrideDatasetParameterOperationProperty{
//   		ParameterName: jsii.String("parameterName"),
//
//   		// the properties below are optional
//   		NewDefaultValues: &NewDefaultValuesProperty{
//   			DateTimeStaticValues: []*string{
//   				jsii.String("dateTimeStaticValues"),
//   			},
//   			DecimalStaticValues: []interface{}{
//   				jsii.Number(123),
//   			},
//   			IntegerStaticValues: []interface{}{
//   				jsii.Number(123),
//   			},
//   			StringStaticValues: []*string{
//   				jsii.String("stringStaticValues"),
//   			},
//   		},
//   		NewParameterName: jsii.String("newParameterName"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-transformoperation.html
//
type CfnDataSet_TransformOperationProperty struct {
	// A transform operation that casts a column to a different type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-transformoperation.html#cfn-quicksight-dataset-transformoperation-castcolumntypeoperation
	//
	CastColumnTypeOperation interface{} `field:"optional" json:"castColumnTypeOperation" yaml:"castColumnTypeOperation"`
	// An operation that creates calculated columns.
	//
	// Columns created in one such operation form a lexical closure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-transformoperation.html#cfn-quicksight-dataset-transformoperation-createcolumnsoperation
	//
	CreateColumnsOperation interface{} `field:"optional" json:"createColumnsOperation" yaml:"createColumnsOperation"`
	// An operation that filters rows based on some condition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-transformoperation.html#cfn-quicksight-dataset-transformoperation-filteroperation
	//
	FilterOperation interface{} `field:"optional" json:"filterOperation" yaml:"filterOperation"`
	// <p>A transform operation that overrides the dataset parameter values defined in another dataset.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-transformoperation.html#cfn-quicksight-dataset-transformoperation-overridedatasetparameteroperation
	//
	OverrideDatasetParameterOperation interface{} `field:"optional" json:"overrideDatasetParameterOperation" yaml:"overrideDatasetParameterOperation"`
	// An operation that projects columns.
	//
	// Operations that come after a projection can only refer to projected columns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-transformoperation.html#cfn-quicksight-dataset-transformoperation-projectoperation
	//
	ProjectOperation interface{} `field:"optional" json:"projectOperation" yaml:"projectOperation"`
	// An operation that renames a column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-transformoperation.html#cfn-quicksight-dataset-transformoperation-renamecolumnoperation
	//
	RenameColumnOperation interface{} `field:"optional" json:"renameColumnOperation" yaml:"renameColumnOperation"`
	// An operation that tags a column with additional information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-transformoperation.html#cfn-quicksight-dataset-transformoperation-tagcolumnoperation
	//
	TagColumnOperation interface{} `field:"optional" json:"tagColumnOperation" yaml:"tagColumnOperation"`
}

