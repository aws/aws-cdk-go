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
//   datastoreStorageProperty := &datastoreStorageProperty{
//   	customerManagedS3: &customerManagedS3Property{
//   		bucket: jsii.String("bucket"),
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		keyPrefix: jsii.String("keyPrefix"),
//   	},
//   	iotSiteWiseMultiLayerStorage: &iotSiteWiseMultiLayerStorageProperty{
//   		customerManagedS3Storage: &customerManagedS3StorageProperty{
//   			bucket: jsii.String("bucket"),
//
//   			// the properties below are optional
//   			keyPrefix: jsii.String("keyPrefix"),
//   		},
//   	},
//   	serviceManagedS3: serviceManagedS3,
//   }
//
type CfnDatastore_DatastoreStorageProperty struct {
	// Use this to store data store data in an S3 bucket that you manage.
	//
	// The choice of service-managed or customer-managed S3 storage cannot be changed after creation of the data store.
	CustomerManagedS3 interface{} `field:"optional" json:"customerManagedS3" yaml:"customerManagedS3"`
	// Use this to store data used by AWS IoT SiteWise in an Amazon S3 bucket that you manage.
	//
	// You can't change the choice of Amazon S3 storage after your data store is created.
	IotSiteWiseMultiLayerStorage interface{} `field:"optional" json:"iotSiteWiseMultiLayerStorage" yaml:"iotSiteWiseMultiLayerStorage"`
	// Use this to store data store data in an S3 bucket managed by the AWS IoT Analytics service.
	//
	// The choice of service-managed or customer-managed S3 storage cannot be changed after creation of the data store.
	ServiceManagedS3 interface{} `field:"optional" json:"serviceManagedS3" yaml:"serviceManagedS3"`
}

