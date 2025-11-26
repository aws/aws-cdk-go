package previewawstimestreammixins


// The set of properties on a table for configuring magnetic store writes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   magneticStoreWritePropertiesProperty := &MagneticStoreWritePropertiesProperty{
//   	EnableMagneticStoreWrites: jsii.Boolean(false),
//   	MagneticStoreRejectedDataLocation: &MagneticStoreRejectedDataLocationProperty{
//   		S3Configuration: &S3ConfigurationProperty{
//   			BucketName: jsii.String("bucketName"),
//   			EncryptionOption: jsii.String("encryptionOption"),
//   			KmsKeyId: jsii.String("kmsKeyId"),
//   			ObjectKeyPrefix: jsii.String("objectKeyPrefix"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-table-magneticstorewriteproperties.html
//
type CfnTablePropsMixin_MagneticStoreWritePropertiesProperty struct {
	// A flag to enable magnetic store writes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-table-magneticstorewriteproperties.html#cfn-timestream-table-magneticstorewriteproperties-enablemagneticstorewrites
	//
	EnableMagneticStoreWrites interface{} `field:"optional" json:"enableMagneticStoreWrites" yaml:"enableMagneticStoreWrites"`
	// The location to write error reports for records rejected asynchronously during magnetic store writes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-table-magneticstorewriteproperties.html#cfn-timestream-table-magneticstorewriteproperties-magneticstorerejecteddatalocation
	//
	MagneticStoreRejectedDataLocation interface{} `field:"optional" json:"magneticStoreRejectedDataLocation" yaml:"magneticStoreRejectedDataLocation"`
}

