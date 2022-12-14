// The CDK Construct Library for AWS::Glue
package awscdkgluealpha


// Encryption options for a Table.
//
// Example:
//   var myDatabase database
//
//   glue.NewTable(this, jsii.String("MyTable"), &tableProps{
//   	encryption: glue.tableEncryption_S3_MANAGED,
//   	// ...
//   	database: myDatabase,
//   	columns: []column{
//   		&column{
//   			name: jsii.String("col1"),
//   			type: glue.schema_STRING(),
//   		},
//   	},
//   	dataFormat: glue.dataFormat_JSON(),
//   })
//
// See: https://docs.aws.amazon.com/athena/latest/ug/encryption.html
//
// Experimental.
type TableEncryption string

const (
	// Experimental.
	TableEncryption_UNENCRYPTED TableEncryption = "UNENCRYPTED"
	// Server side encryption (SSE) with an Amazon S3-managed key.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingServerSideEncryption.html
	//
	// Experimental.
	TableEncryption_S3_MANAGED TableEncryption = "S3_MANAGED"
	// Server-side encryption (SSE) with an AWS KMS key managed by the account owner.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html
	//
	// Experimental.
	TableEncryption_KMS TableEncryption = "KMS"
	// Server-side encryption (SSE) with an AWS KMS key managed by the KMS service.
	// Experimental.
	TableEncryption_KMS_MANAGED TableEncryption = "KMS_MANAGED"
	// Client-side encryption (CSE) with an AWS KMS key managed by the account owner.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingClientSideEncryption.html
	//
	// Experimental.
	TableEncryption_CLIENT_SIDE_KMS TableEncryption = "CLIENT_SIDE_KMS"
)

