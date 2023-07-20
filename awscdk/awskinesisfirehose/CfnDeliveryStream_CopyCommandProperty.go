package awskinesisfirehose


// The `CopyCommand` property type configures the Amazon Redshift `COPY` command that Amazon Kinesis Data Firehose (Kinesis Data Firehose) uses to load data into an Amazon Redshift cluster from an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   copyCommandProperty := &CopyCommandProperty{
//   	DataTableName: jsii.String("dataTableName"),
//
//   	// the properties below are optional
//   	CopyOptions: jsii.String("copyOptions"),
//   	DataTableColumns: jsii.String("dataTableColumns"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-copycommand.html
//
type CfnDeliveryStream_CopyCommandProperty struct {
	// The name of the target table.
	//
	// The table must already exist in the database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-copycommand.html#cfn-kinesisfirehose-deliverystream-copycommand-datatablename
	//
	DataTableName *string `field:"required" json:"dataTableName" yaml:"dataTableName"`
	// Parameters to use with the Amazon Redshift `COPY` command.
	//
	// For examples, see the `CopyOptions` content for the [CopyCommand](https://docs.aws.amazon.com/firehose/latest/APIReference/API_CopyCommand.html) data type in the *Amazon Kinesis Data Firehose API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-copycommand.html#cfn-kinesisfirehose-deliverystream-copycommand-copyoptions
	//
	CopyOptions *string `field:"optional" json:"copyOptions" yaml:"copyOptions"`
	// A comma-separated list of column names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-copycommand.html#cfn-kinesisfirehose-deliverystream-copycommand-datatablecolumns
	//
	DataTableColumns *string `field:"optional" json:"dataTableColumns" yaml:"dataTableColumns"`
}

