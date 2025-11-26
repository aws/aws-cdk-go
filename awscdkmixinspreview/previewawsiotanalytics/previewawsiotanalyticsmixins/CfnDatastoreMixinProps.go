package previewawsiotanalyticsmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDatastorePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var jsonConfiguration interface{}
//   var serviceManagedS3 interface{}
//
//   cfnDatastoreMixinProps := &CfnDatastoreMixinProps{
//   	DatastoreName: jsii.String("datastoreName"),
//   	DatastorePartitions: &DatastorePartitionsProperty{
//   		Partitions: []interface{}{
//   			&DatastorePartitionProperty{
//   				Partition: &PartitionProperty{
//   					AttributeName: jsii.String("attributeName"),
//   				},
//   				TimestampPartition: &TimestampPartitionProperty{
//   					AttributeName: jsii.String("attributeName"),
//   					TimestampFormat: jsii.String("timestampFormat"),
//   				},
//   			},
//   		},
//   	},
//   	DatastoreStorage: &DatastoreStorageProperty{
//   		CustomerManagedS3: &CustomerManagedS3Property{
//   			Bucket: jsii.String("bucket"),
//   			KeyPrefix: jsii.String("keyPrefix"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   		IotSiteWiseMultiLayerStorage: &IotSiteWiseMultiLayerStorageProperty{
//   			CustomerManagedS3Storage: &CustomerManagedS3StorageProperty{
//   				Bucket: jsii.String("bucket"),
//   				KeyPrefix: jsii.String("keyPrefix"),
//   			},
//   		},
//   		ServiceManagedS3: serviceManagedS3,
//   	},
//   	FileFormatConfiguration: &FileFormatConfigurationProperty{
//   		JsonConfiguration: jsonConfiguration,
//   		ParquetConfiguration: &ParquetConfigurationProperty{
//   			SchemaDefinition: &SchemaDefinitionProperty{
//   				Columns: []interface{}{
//   					&ColumnProperty{
//   						Name: jsii.String("name"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	RetentionPeriod: &RetentionPeriodProperty{
//   		NumberOfDays: jsii.Number(123),
//   		Unlimited: jsii.Boolean(false),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotanalytics-datastore.html
//
type CfnDatastoreMixinProps struct {
	// The name of the data store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotanalytics-datastore.html#cfn-iotanalytics-datastore-datastorename
	//
	DatastoreName *string `field:"optional" json:"datastoreName" yaml:"datastoreName"`
	// Information about the partition dimensions in a data store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotanalytics-datastore.html#cfn-iotanalytics-datastore-datastorepartitions
	//
	DatastorePartitions interface{} `field:"optional" json:"datastorePartitions" yaml:"datastorePartitions"`
	// Where data store data is stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotanalytics-datastore.html#cfn-iotanalytics-datastore-datastorestorage
	//
	DatastoreStorage interface{} `field:"optional" json:"datastoreStorage" yaml:"datastoreStorage"`
	// Contains the configuration information of file formats. AWS IoT Analytics data stores support JSON and [Parquet](https://docs.aws.amazon.com/https://parquet.apache.org/) .
	//
	// The default file format is JSON. You can specify only one format.
	//
	// You can't change the file format after you create the data store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotanalytics-datastore.html#cfn-iotanalytics-datastore-fileformatconfiguration
	//
	FileFormatConfiguration interface{} `field:"optional" json:"fileFormatConfiguration" yaml:"fileFormatConfiguration"`
	// How long, in days, message data is kept for the data store.
	//
	// When `customerManagedS3` storage is selected, this parameter is ignored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotanalytics-datastore.html#cfn-iotanalytics-datastore-retentionperiod
	//
	RetentionPeriod interface{} `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// Metadata which can be used to manage the data store.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotanalytics-datastore.html#cfn-iotanalytics-datastore-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

