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
//   		Columns: []interface{}{
//   			&InputColumnProperty{
//   				Name: jsii.String("name"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		DataSourceArn: jsii.String("dataSourceArn"),
//   		Name: jsii.String("name"),
//   		SqlQuery: jsii.String("sqlQuery"),
//   	},
//   	RelationalTable: &RelationalTableProperty{
//   		DataSourceArn: jsii.String("dataSourceArn"),
//   		InputColumns: []interface{}{
//   			&InputColumnProperty{
//   				Name: jsii.String("name"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		Catalog: jsii.String("catalog"),
//   		Schema: jsii.String("schema"),
//   	},
//   	S3Source: &S3SourceProperty{
//   		DataSourceArn: jsii.String("dataSourceArn"),
//   		InputColumns: []interface{}{
//   			&InputColumnProperty{
//   				Name: jsii.String("name"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//
//   		// the properties below are optional
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
type CfnDataSet_PhysicalTableProperty struct {
	// A physical table type built from the results of the custom SQL query.
	CustomSql interface{} `field:"optional" json:"customSql" yaml:"customSql"`
	// A physical table type for relational data sources.
	RelationalTable interface{} `field:"optional" json:"relationalTable" yaml:"relationalTable"`
	// A physical table type for as S3 data source.
	S3Source interface{} `field:"optional" json:"s3Source" yaml:"s3Source"`
}

