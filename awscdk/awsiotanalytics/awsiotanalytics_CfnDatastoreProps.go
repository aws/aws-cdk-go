package awsiotanalytics

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDatastore`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var jsonConfiguration interface{}
//   var serviceManagedS3 interface{}
//
//   cfnDatastoreProps := &cfnDatastoreProps{
//   	datastoreName: jsii.String("datastoreName"),
//   	datastorePartitions: &datastorePartitionsProperty{
//   		partitions: []interface{}{
//   			&datastorePartitionProperty{
//   				partition: &partitionProperty{
//   					attributeName: jsii.String("attributeName"),
//   				},
//   				timestampPartition: &timestampPartitionProperty{
//   					attributeName: jsii.String("attributeName"),
//
//   					// the properties below are optional
//   					timestampFormat: jsii.String("timestampFormat"),
//   				},
//   			},
//   		},
//   	},
//   	datastoreStorage: &datastoreStorageProperty{
//   		customerManagedS3: &customerManagedS3Property{
//   			bucket: jsii.String("bucket"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			keyPrefix: jsii.String("keyPrefix"),
//   		},
//   		iotSiteWiseMultiLayerStorage: &iotSiteWiseMultiLayerStorageProperty{
//   			customerManagedS3Storage: &customerManagedS3StorageProperty{
//   				bucket: jsii.String("bucket"),
//
//   				// the properties below are optional
//   				keyPrefix: jsii.String("keyPrefix"),
//   			},
//   		},
//   		serviceManagedS3: serviceManagedS3,
//   	},
//   	fileFormatConfiguration: &fileFormatConfigurationProperty{
//   		jsonConfiguration: jsonConfiguration,
//   		parquetConfiguration: &parquetConfigurationProperty{
//   			schemaDefinition: &schemaDefinitionProperty{
//   				columns: []interface{}{
//   					&columnProperty{
//   						name: jsii.String("name"),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	retentionPeriod: &retentionPeriodProperty{
//   		numberOfDays: jsii.Number(123),
//   		unlimited: jsii.Boolean(false),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDatastoreProps struct {
	// The name of the data store.
	DatastoreName *string `field:"optional" json:"datastoreName" yaml:"datastoreName"`
	// Information about the partition dimensions in a data store.
	DatastorePartitions interface{} `field:"optional" json:"datastorePartitions" yaml:"datastorePartitions"`
	// Where data store data is stored.
	DatastoreStorage interface{} `field:"optional" json:"datastoreStorage" yaml:"datastoreStorage"`
	// Contains the configuration information of file formats. AWS IoT Analytics data stores support JSON and [Parquet](https://docs.aws.amazon.com/https://parquet.apache.org/) .
	//
	// The default file format is JSON. You can specify only one format.
	//
	// You can't change the file format after you create the data store.
	FileFormatConfiguration interface{} `field:"optional" json:"fileFormatConfiguration" yaml:"fileFormatConfiguration"`
	// How long, in days, message data is kept for the data store.
	//
	// When `customerManagedS3` storage is selected, this parameter is ignored.
	RetentionPeriod interface{} `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// Metadata which can be used to manage the data store.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

