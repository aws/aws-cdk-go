package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

// Location of query result along with S3 bucket configuration.
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
// See: https://docs.aws.amazon.com/athena/latest/APIReference/API_ResultConfiguration.html
//
// Experimental.
type ResultConfiguration struct {
	// Encryption option used if enabled in S3.
	// Experimental.
	EncryptionConfiguration *EncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// S3 path of query results.
	//
	// Example value: `s3://query-results-bucket/folder/`.
	// Experimental.
	OutputLocation *awss3.Location `field:"optional" json:"outputLocation" yaml:"outputLocation"`
}

