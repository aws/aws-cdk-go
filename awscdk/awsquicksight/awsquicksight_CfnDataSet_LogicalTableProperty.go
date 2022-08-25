package awsquicksight


// A *logical table* is a unit that joins and that data transformations operate on.
//
// A logical table has a source, which can be either a physical table or result of a join. When a logical table points to a physical table, the logical table acts as a mutable copy of that physical table through transform operations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logicalTableProperty := &logicalTableProperty{
//   	alias: jsii.String("alias"),
//   	source: &logicalTableSourceProperty{
//   		dataSetArn: jsii.String("dataSetArn"),
//   		joinInstruction: &joinInstructionProperty{
//   			leftOperand: jsii.String("leftOperand"),
//   			onClause: jsii.String("onClause"),
//   			rightOperand: jsii.String("rightOperand"),
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			leftJoinKeyProperties: &joinKeyPropertiesProperty{
//   				uniqueKey: jsii.Boolean(false),
//   			},
//   			rightJoinKeyProperties: &joinKeyPropertiesProperty{
//   				uniqueKey: jsii.Boolean(false),
//   			},
//   		},
//   		physicalTableId: jsii.String("physicalTableId"),
//   	},
//
//   	// the properties below are optional
//   	dataTransforms: []interface{}{
//   		&transformOperationProperty{
//   			castColumnTypeOperation: &castColumnTypeOperationProperty{
//   				columnName: jsii.String("columnName"),
//   				newColumnType: jsii.String("newColumnType"),
//
//   				// the properties below are optional
//   				format: jsii.String("format"),
//   			},
//   			createColumnsOperation: &createColumnsOperationProperty{
//   				columns: []interface{}{
//   					&calculatedColumnProperty{
//   						columnId: jsii.String("columnId"),
//   						columnName: jsii.String("columnName"),
//   						expression: jsii.String("expression"),
//   					},
//   				},
//   			},
//   			filterOperation: &filterOperationProperty{
//   				conditionExpression: jsii.String("conditionExpression"),
//   			},
//   			projectOperation: &projectOperationProperty{
//   				projectedColumns: []*string{
//   					jsii.String("projectedColumns"),
//   				},
//   			},
//   			renameColumnOperation: &renameColumnOperationProperty{
//   				columnName: jsii.String("columnName"),
//   				newColumnName: jsii.String("newColumnName"),
//   			},
//   			tagColumnOperation: &tagColumnOperationProperty{
//   				columnName: jsii.String("columnName"),
//   				tags: []columnTagProperty{
//   					&columnTagProperty{
//   						columnDescription: &columnDescriptionProperty{
//   							text: jsii.String("text"),
//   						},
//   						columnGeographicRole: jsii.String("columnGeographicRole"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnDataSet_LogicalTableProperty struct {
	// A display name for the logical table.
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// Source of this logical table.
	Source interface{} `field:"required" json:"source" yaml:"source"`
	// Transform operations that act on this logical table.
	DataTransforms interface{} `field:"optional" json:"dataTransforms" yaml:"dataTransforms"`
}

