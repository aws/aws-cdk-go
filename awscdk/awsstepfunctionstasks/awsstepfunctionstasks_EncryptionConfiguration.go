package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/awskms"
)

// Encryption Configuration of the S3 bucket.
//
// Example:
//   startQueryExecutionJob := tasks.NewAthenaStartQueryExecution(this, jsii.String("Athena Start Query"), &athenaStartQueryExecutionProps{
//   	queryString: sfn.jsonPath.format(jsii.String("select contacts where year={};"), sfn.*jsonPath.stringAt(jsii.String("$.year"))),
//   	queryExecutionContext: &queryExecutionContext{
//   		databaseName: jsii.String("interactions"),
//   	},
//   	resultConfiguration: &resultConfiguration{
//   		encryptionConfiguration: &encryptionConfiguration{
//   			encryptionOption: tasks.encryptionOption_S3_MANAGED,
//   		},
//   		outputLocation: &location{
//   			bucketName: jsii.String("mybucket"),
//   			objectKey: jsii.String("myprefix"),
//   		},
//   	},
//   	integrationPattern: sfn.integrationPattern_RUN_JOB,
//   })
//
// See: https://docs.aws.amazon.com/athena/latest/APIReference/API_EncryptionConfiguration.html
//
// Experimental.
type EncryptionConfiguration struct {
	// Type of S3 server-side encryption enabled.
	// Experimental.
	EncryptionOption EncryptionOption `field:"required" json:"encryptionOption" yaml:"encryptionOption"`
	// KMS key ARN or ID.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
}

