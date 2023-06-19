package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/awskms"
)

// Encryption Configuration of the S3 bucket.
//
// Example:
//   startQueryExecutionJob := tasks.NewAthenaStartQueryExecution(this, jsii.String("Athena Start Query"), &AthenaStartQueryExecutionProps{
//   	QueryString: sfn.JsonPath_Format(jsii.String("select contacts where year={};"), sfn.JsonPath_StringAt(jsii.String("$.year"))),
//   	QueryExecutionContext: &QueryExecutionContext{
//   		DatabaseName: jsii.String("interactions"),
//   	},
//   	ResultConfiguration: &ResultConfiguration{
//   		EncryptionConfiguration: &EncryptionConfiguration{
//   			EncryptionOption: tasks.EncryptionOption_S3_MANAGED,
//   		},
//   		OutputLocation: &Location{
//   			BucketName: jsii.String("mybucket"),
//   			ObjectKey: jsii.String("myprefix"),
//   		},
//   	},
//   	IntegrationPattern: sfn.IntegrationPattern_RUN_JOB,
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

