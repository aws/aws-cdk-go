package previewawsiotanalyticsmixins


// Stores data used by AWS IoT SiteWise in an Amazon S3 bucket that you manage.
//
// You can't change the choice of Amazon S3 storage after your data store is created.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   iotSiteWiseMultiLayerStorageProperty := &IotSiteWiseMultiLayerStorageProperty{
//   	CustomerManagedS3Storage: &CustomerManagedS3StorageProperty{
//   		Bucket: jsii.String("bucket"),
//   		KeyPrefix: jsii.String("keyPrefix"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-iotsitewisemultilayerstorage.html
//
type CfnDatastorePropsMixin_IotSiteWiseMultiLayerStorageProperty struct {
	// Stores data used by AWS IoT SiteWise in an Amazon S3 bucket that you manage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-iotsitewisemultilayerstorage.html#cfn-iotanalytics-datastore-iotsitewisemultilayerstorage-customermanageds3storage
	//
	CustomerManagedS3Storage interface{} `field:"optional" json:"customerManagedS3Storage" yaml:"customerManagedS3Storage"`
}

