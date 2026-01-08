package awss3


// This resource configures your S3 Storage Lens reports to export to read-only S3 table buckets.
//
// With this resource, you can store your Storage Lens metrics in S3 Tables that are created in a read-only S3 table bucket called aws-s3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var sses3 interface{}
//
//   storageLensTableDestinationProperty := &StorageLensTableDestinationProperty{
//   	IsEnabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	Encryption: &EncryptionProperty{
//   		Ssekms: &SSEKMSProperty{
//   			KeyId: jsii.String("keyId"),
//   		},
//   		Sses3: sses3,
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelenstabledestination.html
//
type CfnStorageLens_StorageLensTableDestinationProperty struct {
	// This property indicates whether the export to read-only S3 table buckets is enabled for your S3 Storage Lens configuration.
	//
	// When set to true, Storage Lens reports are automatically exported to tables in addition to other configured destinations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelenstabledestination.html#cfn-s3-storagelens-storagelenstabledestination-isenabled
	//
	IsEnabled interface{} `field:"required" json:"isEnabled" yaml:"isEnabled"`
	// This resource configures your data encryption settings for Storage Lens metrics in read-only S3 table buckets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelenstabledestination.html#cfn-s3-storagelens-storagelenstabledestination-encryption
	//
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
}

