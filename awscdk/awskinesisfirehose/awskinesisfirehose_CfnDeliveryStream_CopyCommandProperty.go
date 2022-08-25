package awskinesisfirehose


// The `CopyCommand` property type configures the Amazon Redshift `COPY` command that Amazon Kinesis Data Firehose (Kinesis Data Firehose) uses to load data into an Amazon Redshift cluster from an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   copyCommandProperty := &copyCommandProperty{
//   	dataTableName: jsii.String("dataTableName"),
//
//   	// the properties below are optional
//   	copyOptions: jsii.String("copyOptions"),
//   	dataTableColumns: jsii.String("dataTableColumns"),
//   }
//
type CfnDeliveryStream_CopyCommandProperty struct {
	// The name of the target table.
	//
	// The table must already exist in the database.
	DataTableName *string `field:"required" json:"dataTableName" yaml:"dataTableName"`
	// Parameters to use with the Amazon Redshift `COPY` command.
	//
	// For examples, see the `CopyOptions` content for the [CopyCommand](https://docs.aws.amazon.com/firehose/latest/APIReference/API_CopyCommand.html) data type in the *Amazon Kinesis Data Firehose API Reference* .
	CopyOptions *string `field:"optional" json:"copyOptions" yaml:"copyOptions"`
	// A comma-separated list of column names.
	DataTableColumns *string `field:"optional" json:"dataTableColumns" yaml:"dataTableColumns"`
}

