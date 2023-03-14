package awsiotanalytics


// The destination to which dataset contents are delivered.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datasetContentDeliveryRuleDestinationProperty := &DatasetContentDeliveryRuleDestinationProperty{
//   	IotEventsDestinationConfiguration: &IotEventsDestinationConfigurationProperty{
//   		InputName: jsii.String("inputName"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	S3DestinationConfiguration: &S3DestinationConfigurationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		GlueConfiguration: &GlueConfigurationProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			TableName: jsii.String("tableName"),
//   		},
//   	},
//   }
//
type CfnDataset_DatasetContentDeliveryRuleDestinationProperty struct {
	// Configuration information for delivery of dataset contents to AWS IoT Events .
	IotEventsDestinationConfiguration interface{} `field:"optional" json:"iotEventsDestinationConfiguration" yaml:"iotEventsDestinationConfiguration"`
	// Configuration information for delivery of dataset contents to Amazon S3.
	S3DestinationConfiguration interface{} `field:"optional" json:"s3DestinationConfiguration" yaml:"s3DestinationConfiguration"`
}

