package awsiotanalytics


// Configuration information for coordination with AWS Glue , a fully managed extract, transform and load (ETL) service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   glueConfigurationProperty := &glueConfigurationProperty{
//   	databaseName: jsii.String("databaseName"),
//   	tableName: jsii.String("tableName"),
//   }
//
type CfnDataset_GlueConfigurationProperty struct {
	// The name of the database in your AWS Glue Data Catalog in which the table is located.
	//
	// An AWS Glue Data Catalog database contains metadata tables.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The name of the table in your AWS Glue Data Catalog that is used to perform the ETL operations.
	//
	// An AWS Glue Data Catalog table contains partitioned data and descriptions of data sources and targets.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

