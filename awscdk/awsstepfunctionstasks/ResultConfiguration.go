package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Location of query result along with S3 bucket configuration.
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
//   			BucketName: jsii.String("query-results-bucket"),
//   			ObjectKey: jsii.String("folder"),
//   		},
//   	},
//   	ExecutionParameters: []*string{
//   		jsii.String("param1"),
//   		jsii.String("param2"),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/athena/latest/APIReference/API_ResultConfiguration.html
//
type ResultConfiguration struct {
	// Encryption option used if enabled in S3.
	// Default: - SSE_S3 encrpytion is enabled with default encryption key.
	//
	EncryptionConfiguration *EncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// S3 path of query results.
	//
	// Example value: `s3://query-results-bucket/folder/`.
	// Default: - Query Result Location set in Athena settings for this workgroup.
	//
	OutputLocation *awss3.Location `field:"optional" json:"outputLocation" yaml:"outputLocation"`
}

