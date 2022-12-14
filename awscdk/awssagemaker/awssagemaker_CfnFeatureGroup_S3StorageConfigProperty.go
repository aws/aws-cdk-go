package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3StorageConfigProperty := &s3StorageConfigProperty{
//   	s3Uri: jsii.String("s3Uri"),
//
//   	// the properties below are optional
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   }
//
type CfnFeatureGroup_S3StorageConfigProperty struct {
	// `CfnFeatureGroup.S3StorageConfigProperty.S3Uri`.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// `CfnFeatureGroup.S3StorageConfigProperty.KmsKeyId`.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

