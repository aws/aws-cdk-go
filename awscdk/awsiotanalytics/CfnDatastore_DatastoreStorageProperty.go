package awsiotanalytics


// Where data store data is stored.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var serviceManagedS3 interface{}
//
//   datastoreStorageProperty := &DatastoreStorageProperty{
//   	CustomerManagedS3: &CustomerManagedS3Property{
//   		Bucket: jsii.String("bucket"),
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		KeyPrefix: jsii.String("keyPrefix"),
//   	},
//   	IotSiteWiseMultiLayerStorage: &IotSiteWiseMultiLayerStorageProperty{
//   		CustomerManagedS3Storage: &CustomerManagedS3StorageProperty{
//   			Bucket: jsii.String("bucket"),
//
//   			// the properties below are optional
//   			KeyPrefix: jsii.String("keyPrefix"),
//   		},
//   	},
//   	ServiceManagedS3: serviceManagedS3,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-datastorestorage.html
//
type CfnDatastore_DatastoreStorageProperty struct {
	// Use this to store data store data in an S3 bucket that you manage.
	//
	// The choice of service-managed or customer-managed S3 storage cannot be changed after creation of the data store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-datastorestorage.html#cfn-iotanalytics-datastore-datastorestorage-customermanageds3
	//
	CustomerManagedS3 interface{} `field:"optional" json:"customerManagedS3" yaml:"customerManagedS3"`
	// Use this to store data used by AWS IoT SiteWise in an Amazon S3 bucket that you manage.
	//
	// You can't change the choice of Amazon S3 storage after your data store is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-datastorestorage.html#cfn-iotanalytics-datastore-datastorestorage-iotsitewisemultilayerstorage
	//
	IotSiteWiseMultiLayerStorage interface{} `field:"optional" json:"iotSiteWiseMultiLayerStorage" yaml:"iotSiteWiseMultiLayerStorage"`
	// Use this to store data store data in an S3 bucket managed by the AWS IoT Analytics service.
	//
	// The choice of service-managed or customer-managed S3 storage cannot be changed after creation of the data store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-datastorestorage.html#cfn-iotanalytics-datastore-datastorestorage-servicemanageds3
	//
	ServiceManagedS3 interface{} `field:"optional" json:"serviceManagedS3" yaml:"serviceManagedS3"`
}

