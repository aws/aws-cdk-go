package previewawsiotanalyticsmixins


// Configuration information for coordination with AWS Glue , a fully managed extract, transform and load (ETL) service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   glueConfigurationProperty := &GlueConfigurationProperty{
//   	DatabaseName: jsii.String("databaseName"),
//   	TableName: jsii.String("tableName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-glueconfiguration.html
//
type CfnDatasetPropsMixin_GlueConfigurationProperty struct {
	// The name of the database in your AWS Glue Data Catalog in which the table is located.
	//
	// An AWS Glue Data Catalog database contains metadata tables.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-glueconfiguration.html#cfn-iotanalytics-dataset-glueconfiguration-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The name of the table in your AWS Glue Data Catalog that is used to perform the ETL operations.
	//
	// An AWS Glue Data Catalog table contains partitioned data and descriptions of data sources and targets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-glueconfiguration.html#cfn-iotanalytics-dataset-glueconfiguration-tablename
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

