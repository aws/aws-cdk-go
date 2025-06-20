package awscleanrooms


// The schema of a Snowflake table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snowflakeTableSchemaProperty := &SnowflakeTableSchemaProperty{
//   	V1: []interface{}{
//   		&SnowflakeTableSchemaV1Property{
//   			ColumnName: jsii.String("columnName"),
//   			ColumnType: jsii.String("columnType"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-snowflaketableschema.html
//
type CfnConfiguredTable_SnowflakeTableSchemaProperty struct {
	// The schema of a Snowflake table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-snowflaketableschema.html#cfn-cleanrooms-configuredtable-snowflaketableschema-v1
	//
	V1 interface{} `field:"required" json:"v1" yaml:"v1"`
}

