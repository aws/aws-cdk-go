package awssynthetics


// A structure that contains the configuration of the encryption-at-rest settings for artifacts that the canary uploads to Amazon S3 .
//
// Artifact encryption functionality is available only for canaries that use Synthetics runtime version syn-nodejs-puppeteer-3.3 or later. For more information, see [Encrypting canary artifacts](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_artifact_encryption.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3EncryptionProperty := &s3EncryptionProperty{
//   	encryptionMode: jsii.String("encryptionMode"),
//   	kmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
type CfnCanary_S3EncryptionProperty struct {
	// The encryption method to use for artifacts created by this canary.
	//
	// Specify `SSE_S3` to use server-side encryption (SSE) with an Amazon S3-managed key. Specify `SSE-KMS` to use server-side encryption with a customer-managed AWS KMS key.
	//
	// If you omit this parameter, an AWS -managed AWS KMS key is used.
	EncryptionMode *string `field:"optional" json:"encryptionMode" yaml:"encryptionMode"`
	// The ARN of the customer-managed AWS KMS key to use, if you specify `SSE-KMS` for `EncryptionMode`.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

