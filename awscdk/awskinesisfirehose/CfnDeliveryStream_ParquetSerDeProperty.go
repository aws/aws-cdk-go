package awskinesisfirehose


// A serializer to use for converting data to the Parquet format before storing it in Amazon S3.
//
// For more information, see [Apache Parquet](https://docs.aws.amazon.com/https://parquet.apache.org/documentation/latest/) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parquetSerDeProperty := &ParquetSerDeProperty{
//   	BlockSizeBytes: jsii.Number(123),
//   	Compression: jsii.String("compression"),
//   	EnableDictionaryCompression: jsii.Boolean(false),
//   	MaxPaddingBytes: jsii.Number(123),
//   	PageSizeBytes: jsii.Number(123),
//   	WriterVersion: jsii.String("writerVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-parquetserde.html
//
type CfnDeliveryStream_ParquetSerDeProperty struct {
	// The Hadoop Distributed File System (HDFS) block size.
	//
	// This is useful if you intend to copy the data from Amazon S3 to HDFS before querying. The default is 256 MiB and the minimum is 64 MiB. Firehose uses this value for padding calculations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-parquetserde.html#cfn-kinesisfirehose-deliverystream-parquetserde-blocksizebytes
	//
	BlockSizeBytes *float64 `field:"optional" json:"blockSizeBytes" yaml:"blockSizeBytes"`
	// The compression code to use over data blocks.
	//
	// The possible values are `UNCOMPRESSED` , `SNAPPY` , and `GZIP` , with the default being `SNAPPY` . Use `SNAPPY` for higher decompression speed. Use `GZIP` if the compression ratio is more important than speed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-parquetserde.html#cfn-kinesisfirehose-deliverystream-parquetserde-compression
	//
	Compression *string `field:"optional" json:"compression" yaml:"compression"`
	// Indicates whether to enable dictionary compression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-parquetserde.html#cfn-kinesisfirehose-deliverystream-parquetserde-enabledictionarycompression
	//
	EnableDictionaryCompression interface{} `field:"optional" json:"enableDictionaryCompression" yaml:"enableDictionaryCompression"`
	// The maximum amount of padding to apply.
	//
	// This is useful if you intend to copy the data from Amazon S3 to HDFS before querying. The default is 0.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-parquetserde.html#cfn-kinesisfirehose-deliverystream-parquetserde-maxpaddingbytes
	//
	MaxPaddingBytes *float64 `field:"optional" json:"maxPaddingBytes" yaml:"maxPaddingBytes"`
	// The Parquet page size.
	//
	// Column chunks are divided into pages. A page is conceptually an indivisible unit (in terms of compression and encoding). The minimum value is 64 KiB and the default is 1 MiB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-parquetserde.html#cfn-kinesisfirehose-deliverystream-parquetserde-pagesizebytes
	//
	PageSizeBytes *float64 `field:"optional" json:"pageSizeBytes" yaml:"pageSizeBytes"`
	// Indicates the version of row format to output.
	//
	// The possible values are `V1` and `V2` . The default is `V1` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-parquetserde.html#cfn-kinesisfirehose-deliverystream-parquetserde-writerversion
	//
	WriterVersion *string `field:"optional" json:"writerVersion" yaml:"writerVersion"`
}

