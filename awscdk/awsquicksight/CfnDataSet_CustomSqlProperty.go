package awsquicksight


// A physical table type built from the results of the custom SQL query.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customSqlProperty := &CustomSqlProperty{
//   	Columns: []interface{}{
//   		&InputColumnProperty{
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	DataSourceArn: jsii.String("dataSourceArn"),
//   	Name: jsii.String("name"),
//   	SqlQuery: jsii.String("sqlQuery"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-customsql.html
//
type CfnDataSet_CustomSqlProperty struct {
	// The column schema from the SQL query result set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-customsql.html#cfn-quicksight-dataset-customsql-columns
	//
	Columns interface{} `field:"required" json:"columns" yaml:"columns"`
	// The Amazon Resource Name (ARN) of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-customsql.html#cfn-quicksight-dataset-customsql-datasourcearn
	//
	DataSourceArn *string `field:"required" json:"dataSourceArn" yaml:"dataSourceArn"`
	// A display name for the SQL query result.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-customsql.html#cfn-quicksight-dataset-customsql-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The SQL query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-customsql.html#cfn-quicksight-dataset-customsql-sqlquery
	//
	SqlQuery *string `field:"required" json:"sqlQuery" yaml:"sqlQuery"`
}

