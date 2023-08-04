package awscassandra


// Defines an individual column within the clustering key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusteringKeyColumnProperty := &ClusteringKeyColumnProperty{
//   	Column: &ColumnProperty{
//   		ColumnName: jsii.String("columnName"),
//   		ColumnType: jsii.String("columnType"),
//   	},
//
//   	// the properties below are optional
//   	OrderBy: jsii.String("orderBy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-clusteringkeycolumn.html
//
type CfnTable_ClusteringKeyColumnProperty struct {
	// The name and data type of this clustering key column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-clusteringkeycolumn.html#cfn-cassandra-table-clusteringkeycolumn-column
	//
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// The order in which this column's data is stored:.
	//
	// - `ASC` (default) - The column's data is stored in ascending order.
	// - `DESC` - The column's data is stored in descending order.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-clusteringkeycolumn.html#cfn-cassandra-table-clusteringkeycolumn-orderby
	//
	// Default: - "ASC".
	//
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
}

