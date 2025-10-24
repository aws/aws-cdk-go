package awskinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props for ORC output format for data record format conversion.
//
// Example:
//   outputFormat := firehose.NewOrcOutputFormat(&OrcOutputFormatProps{
//   	FormatVersion: firehose.OrcFormatVersion_V0_11,
//   	BlockSize: awscdk.Size_Mebibytes(jsii.Number(256)),
//   	Compression: firehose.OrcCompression_NONE(),
//   	BloomFilterColumns: []*string{
//   		jsii.String("columnA"),
//   	},
//   	BloomFilterFalsePositiveProbability: jsii.Number(0.1),
//   	DictionaryKeyThreshold: jsii.Number(0.7),
//   	EnablePadding: jsii.Boolean(true),
//   	PaddingTolerance: jsii.Number(0.2),
//   	RowIndexStride: jsii.Number(9000),
//   	StripeSize: awscdk.Size_*Mebibytes(jsii.Number(32)),
//   })
//
type OrcOutputFormatProps struct {
	// The Hadoop Distributed File System (HDFS) block size.
	//
	// This is useful if you intend to copy the data from Amazon S3 to HDFS before querying.
	// Firehose uses this value for padding calculations.
	// Default: `Size.mebibytes(256)`
	//
	BlockSize awscdk.Size `field:"optional" json:"blockSize" yaml:"blockSize"`
	// The column names for which you want Firehose to create bloom filters.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-orcserde.html#cfn-kinesisfirehose-deliverystream-orcserde-bloomfiltercolumns
	//
	// Default: no bloom filters are created.
	//
	BloomFilterColumns *[]*string `field:"optional" json:"bloomFilterColumns" yaml:"bloomFilterColumns"`
	// The Bloom filter false positive probability (FPP).
	//
	// The lower the FPP, the bigger the bloom filter.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-orcserde.html#cfn-kinesisfirehose-deliverystream-orcserde-bloomfilterfalsepositiveprobability
	//
	// Default: `0.05`
	//
	BloomFilterFalsePositiveProbability *float64 `field:"optional" json:"bloomFilterFalsePositiveProbability" yaml:"bloomFilterFalsePositiveProbability"`
	// The compression code to use over data blocks.
	//
	// The possible values are `NONE` , `SNAPPY` , and `ZLIB`.
	// Use `SNAPPY` for higher decompression speed.
	// Use `GZIP` if the compression ratio is more important than speed.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-orcserde.html#cfn-kinesisfirehose-deliverystream-orcserde-compression
	//
	// Default: `SNAPPY`.
	//
	Compression OrcCompression `field:"optional" json:"compression" yaml:"compression"`
	// Determines whether dictionary encoding should be applied to a column.
	//
	// If the number of distinct keys (unique values) in a column exceeds this fraction of the total non-null rows in that column, dictionary encoding will be turned off for that specific column.
	//
	// To turn off dictionary encoding, set this threshold to 0. To always use dictionary encoding, set this threshold to 1.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-orcserde.html#cfn-kinesisfirehose-deliverystream-orcserde-dictionarykeythreshold
	//
	// Default: `0.8`
	//
	DictionaryKeyThreshold *float64 `field:"optional" json:"dictionaryKeyThreshold" yaml:"dictionaryKeyThreshold"`
	// Set this to `true` to indicate that you want stripes to be padded to the HDFS block boundaries.
	//
	// This is useful if you intend to copy the data from Amazon S3 to HDFS before querying.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-orcserde.html#cfn-kinesisfirehose-deliverystream-orcserde-enablepadding
	//
	// Default: `false`.
	//
	EnablePadding *bool `field:"optional" json:"enablePadding" yaml:"enablePadding"`
	// The version of the ORC format to write.
	//
	// The possible values are `V0_11` and `V0_12`.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-orcserde.html#cfn-kinesisfirehose-deliverystream-orcserde-formatversion
	//
	// Default: `V0_12`.
	//
	FormatVersion OrcFormatVersion `field:"optional" json:"formatVersion" yaml:"formatVersion"`
	// A number between 0 and 1 that defines the tolerance for block padding as a decimal fraction of stripe size.
	//
	// The default value is 0.05, which means 5 percent of stripe size.
	//
	// For the default values of 64 MiB ORC stripes and 256 MiB HDFS blocks, the default block padding tolerance of 5 percent reserves a maximum of 3.2 MiB for padding within the 256 MiB block.
	// In such a case, if the available size within the block is more than 3.2 MiB, a new, smaller stripe is inserted to fit within that space.
	// This ensures that no stripe crosses block boundaries and causes remote reads within a node-local task.
	//
	// Kinesis Data Firehose ignores this parameter when `EnablePadding` is `false` .
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-orcserde.html#cfn-kinesisfirehose-deliverystream-orcserde-paddingtolerance
	//
	// Default: `0.05` if `enablePadding` is `true`
	//
	PaddingTolerance *float64 `field:"optional" json:"paddingTolerance" yaml:"paddingTolerance"`
	// The number of rows between index entries.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-orcserde.html#cfn-kinesisfirehose-deliverystream-orcserde-rowindexstride
	//
	// Default: 10000.
	//
	RowIndexStride *float64 `field:"optional" json:"rowIndexStride" yaml:"rowIndexStride"`
	// The number of bytes in each stripe.
	//
	// The default is 64 MiB and the minimum is 8 MiB.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-orcserde.html#cfn-kinesisfirehose-deliverystream-orcserde-stripesizebytes
	//
	// Default: `Size.mebibytes(64)`
	//
	StripeSize awscdk.Size `field:"optional" json:"stripeSize" yaml:"stripeSize"`
}

