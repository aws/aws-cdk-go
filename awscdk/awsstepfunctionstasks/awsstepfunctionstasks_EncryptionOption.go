package awsstepfunctionstasks


// Encryption Options of the S3 bucket.
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
// See: https://docs.aws.amazon.com/athena/latest/APIReference/API_EncryptionConfiguration.html#athena-Type-EncryptionConfiguration-EncryptionOption
//
type EncryptionOption string

const (
	// Server side encryption (SSE) with an Amazon S3-managed key.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingServerSideEncryption.html
	//
	EncryptionOption_S3_MANAGED EncryptionOption = "S3_MANAGED"
	// Server-side encryption (SSE) with an AWS KMS key managed by the account owner.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html
	//
	EncryptionOption_KMS EncryptionOption = "KMS"
	// Client-side encryption (CSE) with an AWS KMS key managed by the account owner.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingClientSideEncryption.html
	//
	EncryptionOption_CLIENT_SIDE_KMS EncryptionOption = "CLIENT_SIDE_KMS"
)

