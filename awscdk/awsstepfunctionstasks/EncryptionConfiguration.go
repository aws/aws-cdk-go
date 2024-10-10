package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Encryption Configuration of the S3 bucket.
//
// Example:
//   startQueryExecutionJob := tasks.NewAthenaStartQueryExecution(this, jsii.String("Start Athena Query"), &AthenaStartQueryExecutionProps{
//   	QueryString: sfn.JsonPath_StringAt(jsii.String("$.queryString")),
//   	QueryExecutionContext: &QueryExecutionContext{
//   		DatabaseName: jsii.String("mydatabase"),
//   	},
//   	ResultConfiguration: &ResultConfiguration{
//   		EncryptionConfiguration: &EncryptionConfiguration{
//   			EncryptionOption: tasks.EncryptionOption_S3_MANAGED,
//   		},
//   		OutputLocation: &Location{
//   			BucketName: jsii.String("amzn-s3-demo-bucket"),
//   			ObjectKey: jsii.String("folder"),
//   		},
//   	},
//   	ExecutionParameters: []*string{
//   		jsii.String("param1"),
//   		jsii.String("param2"),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/athena/latest/APIReference/API_EncryptionConfiguration.html
//
type EncryptionConfiguration struct {
	// Type of S3 server-side encryption enabled.
	// Default: EncryptionOption.S3_MANAGED
	//
	EncryptionOption EncryptionOption `field:"required" json:"encryptionOption" yaml:"encryptionOption"`
	// KMS key ARN or ID.
	// Default: - No KMS key for Encryption Option SSE_S3 and default master key for Encryption Option SSE_KMS and CSE_KMS.
	//
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
}

