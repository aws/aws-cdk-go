package previewawsiotanalyticsmixins


// The destination to which dataset contents are delivered.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   datasetContentDeliveryRuleDestinationProperty := &DatasetContentDeliveryRuleDestinationProperty{
//   	IotEventsDestinationConfiguration: &IotEventsDestinationConfigurationProperty{
//   		InputName: jsii.String("inputName"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	S3DestinationConfiguration: &S3DestinationConfigurationProperty{
//   		Bucket: jsii.String("bucket"),
//   		GlueConfiguration: &GlueConfigurationProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			TableName: jsii.String("tableName"),
//   		},
//   		Key: jsii.String("key"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-datasetcontentdeliveryruledestination.html
//
type CfnDatasetPropsMixin_DatasetContentDeliveryRuleDestinationProperty struct {
	// Configuration information for delivery of dataset contents to AWS IoT Events .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-datasetcontentdeliveryruledestination.html#cfn-iotanalytics-dataset-datasetcontentdeliveryruledestination-ioteventsdestinationconfiguration
	//
	IotEventsDestinationConfiguration interface{} `field:"optional" json:"iotEventsDestinationConfiguration" yaml:"iotEventsDestinationConfiguration"`
	// Configuration information for delivery of dataset contents to Amazon S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-datasetcontentdeliveryruledestination.html#cfn-iotanalytics-dataset-datasetcontentdeliveryruledestination-s3destinationconfiguration
	//
	S3DestinationConfiguration interface{} `field:"optional" json:"s3DestinationConfiguration" yaml:"s3DestinationConfiguration"`
}

