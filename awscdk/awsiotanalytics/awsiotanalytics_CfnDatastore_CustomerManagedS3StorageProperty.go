package awsiotanalytics


// Amazon S3 -customer-managed;
//
// When you choose customer-managed storage, the `retentionPeriod` parameter is ignored. You can't change the choice of Amazon S3 storage after your data store is created.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customerManagedS3StorageProperty := &customerManagedS3StorageProperty{
//   	bucket: jsii.String("bucket"),
//
//   	// the properties below are optional
//   	keyPrefix: jsii.String("keyPrefix"),
//   }
//
type CfnDatastore_CustomerManagedS3StorageProperty struct {
	// The name of the Amazon S3 bucket where your data is stored.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// (Optional) The prefix used to create the keys of the data store data objects.
	//
	// Each object in an Amazon S3 bucket has a key that is its unique identifier in the bucket. Each object in a bucket has exactly one key. The prefix must end with a forward slash (/).
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
}

