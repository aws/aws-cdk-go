package awsdynamodb


// The options for imported source files in CSV format.
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
//   		InputFormat: dynamodb.InputFormat_Csv(&CsvOptions{
//   			Delimiter: jsii.String(","),
//   			HeaderList: []*string{
//   				jsii.String("id"),
//   				jsii.String("name"),
//   			},
//   		}),
//   		Bucket: *Bucket,
//   		KeyPrefix: jsii.String("prefix"),
//   	},
//   })
//
type CsvOptions struct {
	// The delimiter used for separating items in the CSV file being imported.
	//
	// Valid delimiters are as follows:
	// - comma (`,`)
	// - tab (`\t`)
	// - colon (`:`)
	// - semicolon (`;`)
	// - pipe (`|`)
	// - space (` `).
	// Default: - use comma as a delimiter.
	//
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// List of the headers used to specify a common header for all source CSV files being imported.
	//
	// **NOTE**: If this field is specified then the first line of each CSV file is treated as data instead of the header.
	// If this field is not specified the the first line of each CSV file is treated as the header.
	// Default: - the first line of the CSV file is treated as the header.
	//
	HeaderList *[]*string `field:"optional" json:"headerList" yaml:"headerList"`
}

