package awsiotanalytics


// When dataset contents are created, they are delivered to destination specified here.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datasetContentDeliveryRuleProperty := &datasetContentDeliveryRuleProperty{
//   	destination: &datasetContentDeliveryRuleDestinationProperty{
//   		iotEventsDestinationConfiguration: &iotEventsDestinationConfigurationProperty{
//   			inputName: jsii.String("inputName"),
//   			roleArn: jsii.String("roleArn"),
//   		},
//   		s3DestinationConfiguration: &s3DestinationConfigurationProperty{
//   			bucket: jsii.String("bucket"),
//   			key: jsii.String("key"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			glueConfiguration: &glueConfigurationProperty{
//   				databaseName: jsii.String("databaseName"),
//   				tableName: jsii.String("tableName"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	entryName: jsii.String("entryName"),
//   }
//
type CfnDataset_DatasetContentDeliveryRuleProperty struct {
	// The destination to which dataset contents are delivered.
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// The name of the dataset content delivery rules entry.
	EntryName *string `field:"optional" json:"entryName" yaml:"entryName"`
}

