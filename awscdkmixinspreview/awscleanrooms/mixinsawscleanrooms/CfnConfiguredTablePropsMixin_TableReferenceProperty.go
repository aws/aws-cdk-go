package mixinsawscleanrooms


// A pointer to the dataset that underlies this table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tableReferenceProperty := &TableReferenceProperty{
//   	Athena: &AthenaTableReferenceProperty{
//   		DatabaseName: jsii.String("databaseName"),
//   		OutputLocation: jsii.String("outputLocation"),
//   		Region: jsii.String("region"),
//   		TableName: jsii.String("tableName"),
//   		WorkGroup: jsii.String("workGroup"),
//   	},
//   	Glue: &GlueTableReferenceProperty{
//   		DatabaseName: jsii.String("databaseName"),
//   		Region: jsii.String("region"),
//   		TableName: jsii.String("tableName"),
//   	},
//   	Snowflake: &SnowflakeTableReferenceProperty{
//   		AccountIdentifier: jsii.String("accountIdentifier"),
//   		DatabaseName: jsii.String("databaseName"),
//   		SchemaName: jsii.String("schemaName"),
//   		SecretArn: jsii.String("secretArn"),
//   		TableName: jsii.String("tableName"),
//   		TableSchema: &SnowflakeTableSchemaProperty{
//   			V1: []interface{}{
//   				&SnowflakeTableSchemaV1Property{
//   					ColumnName: jsii.String("columnName"),
//   					ColumnType: jsii.String("columnType"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-tablereference.html
//
type CfnConfiguredTablePropsMixin_TableReferenceProperty struct {
	// If present, a reference to the Athena table referred to by this table reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-tablereference.html#cfn-cleanrooms-configuredtable-tablereference-athena
	//
	Athena interface{} `field:"optional" json:"athena" yaml:"athena"`
	// If present, a reference to the AWS Glue table referred to by this table reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-tablereference.html#cfn-cleanrooms-configuredtable-tablereference-glue
	//
	Glue interface{} `field:"optional" json:"glue" yaml:"glue"`
	// If present, a reference to the Snowflake table referred to by this table reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-tablereference.html#cfn-cleanrooms-configuredtable-tablereference-snowflake
	//
	Snowflake interface{} `field:"optional" json:"snowflake" yaml:"snowflake"`
}

