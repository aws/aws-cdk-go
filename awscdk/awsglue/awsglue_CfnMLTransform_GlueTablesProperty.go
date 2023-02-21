package awsglue


// The database and table in the AWS Glue Data Catalog that is used for input or output data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   glueTablesProperty := &glueTablesProperty{
//   	databaseName: jsii.String("databaseName"),
//   	tableName: jsii.String("tableName"),
//
//   	// the properties below are optional
//   	catalogId: jsii.String("catalogId"),
//   	connectionName: jsii.String("connectionName"),
//   }
//
type CfnMLTransform_GlueTablesProperty struct {
	// A database name in the AWS Glue Data Catalog .
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// A table name in the AWS Glue Data Catalog .
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// A unique identifier for the AWS Glue Data Catalog .
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// The name of the connection to the AWS Glue Data Catalog .
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
}

