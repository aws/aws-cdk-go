package awsquicksight


// A view of a data source that contains information about the shape of the data in the underlying source.
//
// This is a variant type structure. For this structure to be valid, only one of the attributes can be non-null.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   physicalTableProperty := &PhysicalTableProperty{
//   	CustomSql: &CustomSqlProperty{
//   		DataSourceArn: jsii.String("dataSourceArn"),
//   		Name: jsii.String("name"),
//   		SqlQuery: jsii.String("sqlQuery"),
//
//   		// the properties below are optional
//   		Columns: []interface{}{
//   			&InputColumnProperty{
//   				Name: jsii.String("name"),
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				SubType: jsii.String("subType"),
//   			},
//   		},
//   	},
//   	RelationalTable: &RelationalTableProperty{
//   		DataSourceArn: jsii.String("dataSourceArn"),
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		Catalog: jsii.String("catalog"),
//   		InputColumns: []interface{}{
//   			&InputColumnProperty{
//   				Name: jsii.String("name"),
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				SubType: jsii.String("subType"),
//   			},
//   		},
//   		Schema: jsii.String("schema"),
//   	},
//   	S3Source: &S3SourceProperty{
//   		DataSourceArn: jsii.String("dataSourceArn"),
//
//   		// the properties below are optional
//   		InputColumns: []interface{}{
//   			&InputColumnProperty{
//   				Name: jsii.String("name"),
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				SubType: jsii.String("subType"),
//   			},
//   		},
//   		UploadSettings: &UploadSettingsProperty{
//   			ContainsHeader: jsii.Boolean(false),
//   			Delimiter: jsii.String("delimiter"),
//   			Format: jsii.String("format"),
//   			StartFromRow: jsii.Number(123),
//   			TextQualifier: jsii.String("textQualifier"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-physicaltable.html
//
type CfnDataSet_PhysicalTableProperty struct {
	// A physical table type built from the results of the custom SQL query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-physicaltable.html#cfn-quicksight-dataset-physicaltable-customsql
	//
	CustomSql interface{} `field:"optional" json:"customSql" yaml:"customSql"`
	// A physical table type for relational data sources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-physicaltable.html#cfn-quicksight-dataset-physicaltable-relationaltable
	//
	RelationalTable interface{} `field:"optional" json:"relationalTable" yaml:"relationalTable"`
	// A physical table type for as S3 data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-physicaltable.html#cfn-quicksight-dataset-physicaltable-s3source
	//
	S3Source interface{} `field:"optional" json:"s3Source" yaml:"s3Source"`
}

