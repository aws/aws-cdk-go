package awsiotanalytics


// Stores data used by AWS IoT SiteWise in an Amazon S3 bucket that you manage.
//
// You can't change the choice of Amazon S3 storage after your data store is created.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iotSiteWiseMultiLayerStorageProperty := &iotSiteWiseMultiLayerStorageProperty{
//   	customerManagedS3Storage: &customerManagedS3StorageProperty{
//   		bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		keyPrefix: jsii.String("keyPrefix"),
//   	},
//   }
//
type CfnDatastore_IotSiteWiseMultiLayerStorageProperty struct {
	// Stores data used by AWS IoT SiteWise in an Amazon S3 bucket that you manage.
	CustomerManagedS3Storage interface{} `field:"optional" json:"customerManagedS3Storage" yaml:"customerManagedS3Storage"`
}

