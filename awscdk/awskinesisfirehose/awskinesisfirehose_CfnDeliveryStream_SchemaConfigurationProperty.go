package awskinesisfirehose


// Specifies the schema to which you want Kinesis Data Firehose to configure your data before it writes it to Amazon S3.
//
// This parameter is required if `Enabled` is set to true.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   schemaConfigurationProperty := &schemaConfigurationProperty{
//   	catalogId: jsii.String("catalogId"),
//   	databaseName: jsii.String("databaseName"),
//   	region: jsii.String("region"),
//   	roleArn: jsii.String("roleArn"),
//   	tableName: jsii.String("tableName"),
//   	versionId: jsii.String("versionId"),
//   }
//
type CfnDeliveryStream_SchemaConfigurationProperty struct {
	// The ID of the AWS Glue Data Catalog.
	//
	// If you don't supply this, the AWS account ID is used by default.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// Specifies the name of the AWS Glue database that contains the schema for the output data.
	//
	// > If the `SchemaConfiguration` request parameter is used as part of invoking the `CreateDeliveryStream` API, then the `DatabaseName` property is required and its value must be specified.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// If you don't specify an AWS Region, the default is the current Region.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The role that Kinesis Data Firehose can use to access AWS Glue.
	//
	// This role must be in the same account you use for Kinesis Data Firehose. Cross-account roles aren't allowed.
	//
	// > If the `SchemaConfiguration` request parameter is used as part of invoking the `CreateDeliveryStream` API, then the `RoleARN` property is required and its value must be specified.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Specifies the AWS Glue table that contains the column information that constitutes your data schema.
	//
	// > If the `SchemaConfiguration` request parameter is used as part of invoking the `CreateDeliveryStream` API, then the `TableName` property is required and its value must be specified.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// Specifies the table version for the output data schema.
	//
	// If you don't specify this version ID, or if you set it to `LATEST` , Kinesis Data Firehose uses the most recent version. This means that any updates to the table are automatically picked up.
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

