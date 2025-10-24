package awssynthetics


// Encryption mode for canary artifacts.
//
// Example:
//   import kms "github.com/aws/aws-cdk-go/awscdk"
//
//
//   key := kms.NewKey(this, jsii.String("myKey"))
//
//   canary := synthetics.NewCanary(this, jsii.String("MyCanary"), &CanaryProps{
//   	Schedule: synthetics.Schedule_Rate(awscdk.Duration_Minutes(jsii.Number(5))),
//   	Test: synthetics.Test_Custom(&CustomTestOptions{
//   		Code: synthetics.Code_FromAsset(path.join(__dirname, jsii.String("canary"))),
//   		Handler: jsii.String("index.handler"),
//   	}),
//   	Runtime: synthetics.Runtime_SYNTHETICS_NODEJS_PUPPETEER_7_0(),
//   	ArtifactsBucketLifecycleRules: []LifecycleRule{
//   		&LifecycleRule{
//   			Expiration: awscdk.Duration_Days(jsii.Number(30)),
//   		},
//   	},
//   	ArtifactS3EncryptionMode: synthetics.ArtifactsEncryptionMode_KMS,
//   	ArtifactS3KmsKey: key,
//   })
//
type ArtifactsEncryptionMode string

const (
	// Server-side encryption (SSE) with an Amazon S3-managed key.
	ArtifactsEncryptionMode_S3_MANAGED ArtifactsEncryptionMode = "S3_MANAGED"
	// Server-side encryption (SSE) with an AWS KMS customer managed key.
	ArtifactsEncryptionMode_KMS ArtifactsEncryptionMode = "KMS"
)

