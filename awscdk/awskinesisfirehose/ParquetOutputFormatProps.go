package awskinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props for Parquet output format for data record format conversion.
//
// Example:
//   outputFormat := firehose.NewParquetOutputFormat(&ParquetOutputFormatProps{
//   	BlockSize: awscdk.Size_Mebibytes(jsii.Number(512)),
//   	Compression: firehose.ParquetCompression_UNCOMPRESSED(),
//   	EnableDictionaryCompression: jsii.Boolean(true),
//   	MaxPadding: awscdk.Size_Bytes(jsii.Number(10)),
//   	PageSize: awscdk.Size_*Mebibytes(jsii.Number(2)),
//   	WriterVersion: firehose.ParquetWriterVersion_V2,
//   })
//
type ParquetOutputFormatProps struct {
	// The Hadoop Distributed File System (HDFS) block size.
	//
	// This is useful if you intend to copy the data from Amazon S3 to HDFS before querying.
	// Firehose uses this value for padding calculations.
	// Default: `Size.mebibytes(256)`
	//
	BlockSize awscdk.Size `field:"optional" json:"blockSize" yaml:"blockSize"`
	// The compression code to use over data blocks.
	//
	// The possible values are `UNCOMPRESSED` , `SNAPPY` , and `GZIP`.
	// Use `SNAPPY` for higher decompression speed.
	// Use `GZIP` if the compression ratio is more important than speed.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-parquetserde.html#cfn-kinesisfirehose-deliverystream-parquetserde-compression
	//
	// Default: `SNAPPY`.
	//
	Compression ParquetCompression `field:"optional" json:"compression" yaml:"compression"`
	// Indicates whether to enable dictionary compression.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-parquetserde.html#cfn-kinesisfirehose-deliverystream-parquetserde-enabledictionarycompression
	//
	// Default: `false`.
	//
	EnableDictionaryCompression *bool `field:"optional" json:"enableDictionaryCompression" yaml:"enableDictionaryCompression"`
	// The maximum amount of padding to apply.
	//
	// This is useful if you intend to copy the data from Amazon S3 to HDFS before querying.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-parquetserde.html#cfn-kinesisfirehose-deliverystream-parquetserde-maxpaddingbytes
	//
	// Default: no padding is applied.
	//
	MaxPadding awscdk.Size `field:"optional" json:"maxPadding" yaml:"maxPadding"`
	// The Parquet page size.
	//
	// Column chunks are divided into pages. A page is conceptually an indivisible unit (in terms of compression and encoding). The minimum value is 64 KiB and the default is 1 MiB.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-parquetserde.html#cfn-kinesisfirehose-deliverystream-parquetserde-pagesizebytes
	//
	// Default: `Size.mebibytes(1)`
	//
	PageSize awscdk.Size `field:"optional" json:"pageSize" yaml:"pageSize"`
	// Indicates the version of Parquet to output.
	//
	// The possible values are `V1` and `V2`.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-parquetserde.html#cfn-kinesisfirehose-deliverystream-parquetserde-writerversion
	//
	// Default: `V1`.
	//
	WriterVersion ParquetWriterVersion `field:"optional" json:"writerVersion" yaml:"writerVersion"`
}

