package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Encryption Configuration of the S3 bucket.
//
// Example:
//   startQueryExecutionJob := tasks.NewAthenaStartQueryExecution(this, jsii.String("Start Athena Query"), &athenaStartQueryExecutionProps{
//   	queryString: sfn.jsonPath.stringAt(jsii.String("$.queryString")),
//   	queryExecutionContext: &queryExecutionContext{
//   		databaseName: jsii.String("mydatabase"),
//   	},
//   	resultConfiguration: &resultConfiguration{
//   		encryptionConfiguration: &encryptionConfiguration{
//   			encryptionOption: tasks.encryptionOption_S3_MANAGED,
//   		},
//   		outputLocation: &location{
//   			bucketName: jsii.String("query-results-bucket"),
//   			objectKey: jsii.String("folder"),
//   		},
//   	},
//   })
//
// See: https://docs.aws.amazon.com/athena/latest/APIReference/API_EncryptionConfiguration.html
//
type EncryptionConfiguration struct {
	// Type of S3 server-side encryption enabled.
	EncryptionOption EncryptionOption `field:"required" json:"encryptionOption" yaml:"encryptionOption"`
	// KMS key ARN or ID.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
}

