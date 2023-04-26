// The CDK Construct Library for AWS::Glue
package awscdkgluealpha


// Encryption options for a Table.
//
// Example:
//   var myDatabase database
//
//   glue.NewTable(this, jsii.String("MyTable"), &TableProps{
//   	Encryption: glue.TableEncryption_S3_MANAGED,
//   	// ...
//   	Database: myDatabase,
//   	Columns: []column{
//   		&column{
//   			Name: jsii.String("col1"),
//   			Type: glue.Schema_STRING(),
//   		},
//   	},
//   	DataFormat: glue.DataFormat_JSON(),
//   })
//
// See: https://docs.aws.amazon.com/athena/latest/ug/encryption.html
//
// Experimental.
type TableEncryption string

const (
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

