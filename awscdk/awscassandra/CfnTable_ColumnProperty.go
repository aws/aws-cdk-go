package awscassandra


// The name and data type of an individual column in a table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnProperty := &ColumnProperty{
//   	ColumnName: jsii.String("columnName"),
//   	ColumnType: jsii.String("columnType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-column.html
//
type CfnTable_ColumnProperty struct {
	// The name of the column.
	//
	// For more information, see [Identifiers](https://docs.aws.amazon.com/keyspaces/latest/devguide/cql.elements.html#cql.elements.identifier) in the *Amazon Keyspaces Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-column.html#cfn-cassandra-table-column-columnname
	//
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// The data type of the column.
	//
	// For more information, see [Data types](https://docs.aws.amazon.com/keyspaces/latest/devguide/cql.elements.html#cql.data-types) in the *Amazon Keyspaces Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-column.html#cfn-cassandra-table-column-columntype
	//
	ColumnType *string `field:"required" json:"columnType" yaml:"columnType"`
}

