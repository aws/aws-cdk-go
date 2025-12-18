package awss3tables


// The configuration details for the storage class of tables or table buckets.
//
// This allows you to optimize storage costs by selecting the appropriate storage class based on your access patterns and performance requirements.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   storageClassConfigurationProperty := &StorageClassConfigurationProperty{
//   	StorageClass: jsii.String("storageClass"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-tablebucket-storageclassconfiguration.html
//
type CfnTableBucket_StorageClassConfigurationProperty struct {
	// The storage class for the table or table bucket.
	//
	// Valid values include storage classes optimized for different access patterns and cost profiles.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-tablebucket-storageclassconfiguration.html#cfn-s3tables-tablebucket-storageclassconfiguration-storageclass
	//
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
}

