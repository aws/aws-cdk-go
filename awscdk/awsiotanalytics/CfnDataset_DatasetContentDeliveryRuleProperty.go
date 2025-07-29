package awsiotanalytics


// When dataset contents are created, they are delivered to destination specified here.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datasetContentDeliveryRuleProperty := &DatasetContentDeliveryRuleProperty{
//   	Destination: &DatasetContentDeliveryRuleDestinationProperty{
//   		IotEventsDestinationConfiguration: &IotEventsDestinationConfigurationProperty{
//   			InputName: jsii.String("inputName"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   		S3DestinationConfiguration: &S3DestinationConfigurationProperty{
//   			Bucket: jsii.String("bucket"),
//   			Key: jsii.String("key"),
//   			RoleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			GlueConfiguration: &GlueConfigurationProperty{
//   				DatabaseName: jsii.String("databaseName"),
//   				TableName: jsii.String("tableName"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	EntryName: jsii.String("entryName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-datasetcontentdeliveryrule.html
//
type CfnDataset_DatasetContentDeliveryRuleProperty struct {
	// The destination to which dataset contents are delivered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-datasetcontentdeliveryrule.html#cfn-iotanalytics-dataset-datasetcontentdeliveryrule-destination
	//
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// The name of the dataset content delivery rules entry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-datasetcontentdeliveryrule.html#cfn-iotanalytics-dataset-datasetcontentdeliveryrule-entryname
	//
	EntryName *string `field:"optional" json:"entryName" yaml:"entryName"`
}

