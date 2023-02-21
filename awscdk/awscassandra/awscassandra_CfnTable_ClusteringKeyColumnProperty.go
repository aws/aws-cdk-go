package awscassandra


// Defines an individual column within the clustering key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusteringKeyColumnProperty := &clusteringKeyColumnProperty{
//   	column: &columnProperty{
//   		columnName: jsii.String("columnName"),
//   		columnType: jsii.String("columnType"),
//   	},
//
//   	// the properties below are optional
//   	orderBy: jsii.String("orderBy"),
//   }
//
type CfnTable_ClusteringKeyColumnProperty struct {
	// The name and data type of this clustering key column.
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// The order in which this column's data is stored:.
	//
	// - `ASC` (default) - The column's data is stored in ascending order.
	// - `DESC` - The column's data is stored in descending order.
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
}

