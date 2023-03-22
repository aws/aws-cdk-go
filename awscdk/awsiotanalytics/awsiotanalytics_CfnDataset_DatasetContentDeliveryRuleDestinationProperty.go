package awsiotanalytics


// The destination to which dataset contents are delivered.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datasetContentDeliveryRuleDestinationProperty := &datasetContentDeliveryRuleDestinationProperty{
//   	iotEventsDestinationConfiguration: &iotEventsDestinationConfigurationProperty{
//   		inputName: jsii.String("inputName"),
//   		roleArn: jsii.String("roleArn"),
//   	},
//   	s3DestinationConfiguration: &s3DestinationConfigurationProperty{
//   		bucket: jsii.String("bucket"),
//   		key: jsii.String("key"),
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		glueConfiguration: &glueConfigurationProperty{
//   			databaseName: jsii.String("databaseName"),
//   			tableName: jsii.String("tableName"),
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

