package awstimestream


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   magneticStoreWritePropertiesProperty := &magneticStoreWritePropertiesProperty{
//   	enableMagneticStoreWrites: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	magneticStoreRejectedDataLocation: &magneticStoreRejectedDataLocationProperty{
//   		s3Configuration: &s3ConfigurationProperty{
//   			bucketName: jsii.String("bucketName"),
//   			encryptionOption: jsii.String("encryptionOption"),
//
//   			// the properties below are optional
//   			kmsKeyId: jsii.String("kmsKeyId"),
//   			objectKeyPrefix: jsii.String("objectKeyPrefix"),
//   		},
//   	},
//   }
//
type CfnTable_MagneticStoreWritePropertiesProperty struct {
	// `CfnTable.MagneticStoreWritePropertiesProperty.EnableMagneticStoreWrites`.
	EnableMagneticStoreWrites interface{} `field:"required" json:"enableMagneticStoreWrites" yaml:"enableMagneticStoreWrites"`
	// `CfnTable.MagneticStoreWritePropertiesProperty.MagneticStoreRejectedDataLocation`.
	MagneticStoreRejectedDataLocation interface{} `field:"optional" json:"magneticStoreRejectedDataLocation" yaml:"magneticStoreRejectedDataLocation"`
}

