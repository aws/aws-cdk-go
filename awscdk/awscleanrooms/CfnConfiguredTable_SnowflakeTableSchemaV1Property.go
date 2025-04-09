package awscleanrooms


// The Snowflake table schema.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snowflakeTableSchemaV1Property := &SnowflakeTableSchemaV1Property{
//   	ColumnName: jsii.String("columnName"),
//   	ColumnType: jsii.String("columnType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-snowflaketableschemav1.html
//
type CfnConfiguredTable_SnowflakeTableSchemaV1Property struct {
	// The column name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-snowflaketableschemav1.html#cfn-cleanrooms-configuredtable-snowflaketableschemav1-columnname
	//
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// The column's data type.
	//
	// Supported data types: `ARRAY` , `BIGINT` , `BOOLEAN` , `CHAR` , `DATE` , `DECIMAL` , `DOUBLE` , `DOUBLE PRECISION` , `FLOAT` , `FLOAT4` , `INT` , `INTEGER` , `MAP` , `NUMERIC` , `NUMBER` , `REAL` , `SMALLINT` , `STRING` , `TIMESTAMP` , `TIMESTAMP_LTZ` , `TIMESTAMP_NTZ` , `DATETIME` , `TINYINT` , `VARCHAR` , `TEXT` , `CHARACTER` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-snowflaketableschemav1.html#cfn-cleanrooms-configuredtable-snowflaketableschemav1-columntype
	//
	ColumnType *string `field:"required" json:"columnType" yaml:"columnType"`
}

