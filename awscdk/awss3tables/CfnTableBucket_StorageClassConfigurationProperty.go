package awss3tables


// Specifies storage class settings for the table bucket.
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
	// The storage class for the table bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-tablebucket-storageclassconfiguration.html#cfn-s3tables-tablebucket-storageclassconfiguration-storageclass
	//
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
}

