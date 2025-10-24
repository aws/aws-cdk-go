package awsdynamodb


// Type of compression to use for imported data.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket IBucket
//
//
//   app := cdk.NewApp()
//   stack := cdk.NewStack(app, jsii.String("Stack"))
//
//   dynamodb.NewTable(stack, jsii.String("Table"), &TableProps{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("id"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	ImportSource: &ImportSourceSpecification{
//   		CompressionType: dynamodb.InputCompressionType_GZIP,
//   		InputFormat: dynamodb.InputFormat_DynamoDBJson(),
//   		Bucket: *Bucket,
//   		KeyPrefix: jsii.String("prefix"),
//   	},
//   })
//
type InputCompressionType string

const (
	// GZIP compression.
	InputCompressionType_GZIP InputCompressionType = "GZIP"
	// ZSTD compression.
	InputCompressionType_ZSTD InputCompressionType = "ZSTD"
	// No compression.
	InputCompressionType_NONE InputCompressionType = "NONE"
)

