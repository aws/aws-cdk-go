package awsstepfunctionstasks


// Encryption Options of the S3 bucket.
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
// See: https://docs.aws.amazon.com/athena/latest/APIReference/API_EncryptionConfiguration.html#athena-Type-EncryptionConfiguration-EncryptionOption
//
// Experimental.
type EncryptionOption string

const (
	// Server side encryption (SSE) with an Amazon S3-managed key.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingServerSideEncryption.html
	//
	// Experimental.
	EncryptionOption_S3_MANAGED EncryptionOption = "S3_MANAGED"
	// Server-side encryption (SSE) with an AWS KMS key managed by the account owner.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html
	//
	// Experimental.
	EncryptionOption_KMS EncryptionOption = "KMS"
	// Client-side encryption (CSE) with an AWS KMS key managed by the account owner.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingClientSideEncryption.html
	//
	// Experimental.
	EncryptionOption_CLIENT_SIDE_KMS EncryptionOption = "CLIENT_SIDE_KMS"
)

