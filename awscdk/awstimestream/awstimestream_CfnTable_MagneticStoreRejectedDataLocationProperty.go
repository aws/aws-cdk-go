package awstimestream


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   magneticStoreRejectedDataLocationProperty := &magneticStoreRejectedDataLocationProperty{
//   	s3Configuration: &s3ConfigurationProperty{
//   		bucketName: jsii.String("bucketName"),
//   		encryptionOption: jsii.String("encryptionOption"),
//
//   		// the properties below are optional
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   		objectKeyPrefix: jsii.String("objectKeyPrefix"),
//   	},
//   }
//
type CfnTable_MagneticStoreRejectedDataLocationProperty struct {
	// `CfnTable.MagneticStoreRejectedDataLocationProperty.S3Configuration`.
	S3Configuration interface{} `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
}

