package awscleanrooms


// A reference to a table within Snowflake.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snowflakeTableReferenceProperty := &SnowflakeTableReferenceProperty{
//   	AccountIdentifier: jsii.String("accountIdentifier"),
//   	DatabaseName: jsii.String("databaseName"),
//   	SchemaName: jsii.String("schemaName"),
//   	SecretArn: jsii.String("secretArn"),
//   	TableName: jsii.String("tableName"),
//   	TableSchema: &SnowflakeTableSchemaProperty{
//   		V1: []interface{}{
//   			&SnowflakeTableSchemaV1Property{
//   				ColumnName: jsii.String("columnName"),
//   				ColumnType: jsii.String("columnType"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-snowflaketablereference.html
//
type CfnConfiguredTable_SnowflakeTableReferenceProperty struct {
	// The account identifier for the Snowflake table reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-snowflaketablereference.html#cfn-cleanrooms-configuredtable-snowflaketablereference-accountidentifier
	//
	AccountIdentifier *string `field:"required" json:"accountIdentifier" yaml:"accountIdentifier"`
	// The name of the database the Snowflake table belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-snowflaketablereference.html#cfn-cleanrooms-configuredtable-snowflaketablereference-databasename
	//
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The schema name of the Snowflake table reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-snowflaketablereference.html#cfn-cleanrooms-configuredtable-snowflaketablereference-schemaname
	//
	SchemaName *string `field:"required" json:"schemaName" yaml:"schemaName"`
	// The secret ARN of the Snowflake table reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-snowflaketablereference.html#cfn-cleanrooms-configuredtable-snowflaketablereference-secretarn
	//
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// The name of the Snowflake table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-snowflaketablereference.html#cfn-cleanrooms-configuredtable-snowflaketablereference-tablename
	//
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// The schema of the Snowflake table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-snowflaketablereference.html#cfn-cleanrooms-configuredtable-snowflaketablereference-tableschema
	//
	TableSchema interface{} `field:"required" json:"tableSchema" yaml:"tableSchema"`
}

