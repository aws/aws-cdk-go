package previewawsiotanalyticsmixins


// Configuration information for delivery of dataset contents to Amazon Simple Storage Service (Amazon S3).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3DestinationConfigurationProperty := &S3DestinationConfigurationProperty{
//   	Bucket: jsii.String("bucket"),
//   	GlueConfiguration: &GlueConfigurationProperty{
//   		DatabaseName: jsii.String("databaseName"),
//   		TableName: jsii.String("tableName"),
//   	},
//   	Key: jsii.String("key"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-s3destinationconfiguration.html
//
type CfnDatasetPropsMixin_S3DestinationConfigurationProperty struct {
	// The name of the S3 bucket to which dataset contents are delivered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-s3destinationconfiguration.html#cfn-iotanalytics-dataset-s3destinationconfiguration-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Configuration information for coordination with AWS Glue , a fully managed extract, transform and load (ETL) service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-s3destinationconfiguration.html#cfn-iotanalytics-dataset-s3destinationconfiguration-glueconfiguration
	//
	GlueConfiguration interface{} `field:"optional" json:"glueConfiguration" yaml:"glueConfiguration"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-s3destinationconfiguration.html#cfn-iotanalytics-dataset-s3destinationconfiguration-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The ARN of the role that grants AWS IoT Analytics permission to interact with your Amazon S3 and AWS Glue resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-s3destinationconfiguration.html#cfn-iotanalytics-dataset-s3destinationconfiguration-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

