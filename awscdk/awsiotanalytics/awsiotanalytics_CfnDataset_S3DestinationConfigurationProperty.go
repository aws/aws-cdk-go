package awsiotanalytics


// Configuration information for delivery of dataset contents to Amazon Simple Storage Service (Amazon S3).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3DestinationConfigurationProperty := &s3DestinationConfigurationProperty{
//   	bucket: jsii.String("bucket"),
//   	key: jsii.String("key"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	glueConfiguration: &glueConfigurationProperty{
//   		databaseName: jsii.String("databaseName"),
//   		tableName: jsii.String("tableName"),
//   	},
//   }
//
type CfnDataset_S3DestinationConfigurationProperty struct {
	// The name of the S3 bucket to which dataset contents are delivered.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The key of the dataset contents object in an S3 bucket.
	//
	// Each object has a key that is a unique identifier. Each object has exactly one key.
	//
	// You can create a unique key with the following options:
	//
	// - Use `!{iotanalytics:scheduleTime}` to insert the time of a scheduled SQL query run.
	// - Use `!{iotanalytics:versionId}` to insert a unique hash that identifies a dataset content.
	// - Use `!{iotanalytics:creationTime}` to insert the creation time of a dataset content.
	//
	// The following example creates a unique key for a CSV file: `dataset/mydataset/!{iotanalytics:scheduleTime}/!{iotanalytics:versionId}.csv`
	//
	// > If you don't use `!{iotanalytics:versionId}` to specify the key, you might get duplicate keys. For example, you might have two dataset contents with the same `scheduleTime` but different `versionId` s. This means that one dataset content overwrites the other.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The ARN of the role that grants AWS IoT Analytics permission to interact with your Amazon S3 and AWS Glue resources.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Configuration information for coordination with AWS Glue , a fully managed extract, transform and load (ETL) service.
	GlueConfiguration interface{} `field:"optional" json:"glueConfiguration" yaml:"glueConfiguration"`
}

