package awssynthetics


// A structure that contains the configuration for canary artifacts, including the encryption-at-rest settings for artifacts that the canary uploads to Amazon S3 .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   artifactConfigProperty := &artifactConfigProperty{
//   	s3Encryption: &s3EncryptionProperty{
//   		encryptionMode: jsii.String("encryptionMode"),
//   		kmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   }
//
type CfnCanary_ArtifactConfigProperty struct {
	// A structure that contains the configuration of the encryption-at-rest settings for artifacts that the canary uploads to Amazon S3 .
	//
	// Artifact encryption functionality is available only for canaries that use Synthetics runtime version syn-nodejs-puppeteer-3.3 or later. For more information, see [Encrypting canary artifacts](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_artifact_encryption.html) .
	S3Encryption interface{} `field:"optional" json:"s3Encryption" yaml:"s3Encryption"`
}

