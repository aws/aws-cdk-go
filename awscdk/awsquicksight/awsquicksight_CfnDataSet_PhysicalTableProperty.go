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
//   physicalTableProperty := &physicalTableProperty{
//   	customSql: &customSqlProperty{
//   		columns: []interface{}{
//   			&inputColumnProperty{
//   				name: jsii.String("name"),
//   				type: jsii.String("type"),
//   			},
//   		},
//   		dataSourceArn: jsii.String("dataSourceArn"),
//   		name: jsii.String("name"),
//   		sqlQuery: jsii.String("sqlQuery"),
//   	},
//   	relationalTable: &relationalTableProperty{
//   		dataSourceArn: jsii.String("dataSourceArn"),
//   		inputColumns: []interface{}{
//   			&inputColumnProperty{
//   				name: jsii.String("name"),
//   				type: jsii.String("type"),
//   			},
//   		},
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		catalog: jsii.String("catalog"),
//   		schema: jsii.String("schema"),
//   	},
//   	s3Source: &s3SourceProperty{
//   		dataSourceArn: jsii.String("dataSourceArn"),
//   		inputColumns: []interface{}{
//   			&inputColumnProperty{
//   				name: jsii.String("name"),
//   				type: jsii.String("type"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		uploadSettings: &uploadSettingsProperty{
//   			containsHeader: jsii.Boolean(false),
//   			delimiter: jsii.String("delimiter"),
//   			format: jsii.String("format"),
//   			startFromRow: jsii.Number(123),
//   			textQualifier: jsii.String("textQualifier"),
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

