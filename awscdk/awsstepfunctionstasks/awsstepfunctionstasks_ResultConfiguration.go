package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Location of query result along with S3 bucket configuration.
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
// See: https://docs.aws.amazon.com/athena/latest/APIReference/API_ResultConfiguration.html
//
type ResultConfiguration struct {
	// Encryption option used if enabled in S3.
	EncryptionConfiguration *EncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// S3 path of query results.
	//
	// Example value: `s3://query-results-bucket/folder/`.
	OutputLocation *awss3.Location `field:"optional" json:"outputLocation" yaml:"outputLocation"`
}

