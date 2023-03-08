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
//   createColumnsOperationProperty := &createColumnsOperationProperty{
//   	columns: []interface{}{
//   		&calculatedColumnProperty{
//   			columnId: jsii.String("columnId"),
//   			columnName: jsii.String("columnName"),
//   			expression: jsii.String("expression"),
//   		},
//   	},
//   }
//
type CfnDataSet_CreateColumnsOperationProperty struct {
	// Calculated columns to create.
	Columns interface{} `field:"required" json:"columns" yaml:"columns"`
}

