package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

// Location of query result along with S3 bucket configuration.
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

