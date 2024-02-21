package awskinesisfirehose


// A serializer to use for converting data to the ORC format before storing it in Amazon S3.
//
// For more information, see [Apache ORC](https://docs.aws.amazon.com/https://orc.apache.org/docs/) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   orcSerDeProperty := &OrcSerDeProperty{
//   	BlockSizeBytes: jsii.Number(123),
//   	BloomFilterColumns: []*string{
//   		jsii.String("bloomFilterColumns"),
//   	},
//   	BloomFilterFalsePositiveProbability: jsii.Number(123),
//   	Compression: jsii.String("compression"),
//   	DictionaryKeyThreshold: jsii.Number(123),
//   	EnablePadding: jsii.Boolean(false),
//   	FormatVersion: jsii.String("formatVersion"),
//   	PaddingTolerance: jsii.Number(123),
//   	RowIndexStride: jsii.Number(123),
//   	StripeSizeBytes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-orcserde.html
//
type CfnDeliveryStream_OrcSerDeProperty struct {
	// The Hadoop Distributed File System (HDFS) block size.
	//
	// This is useful if you intend to copy the data from Amazon S3 to HDFS before querying. The default is 256 MiB and the minimum is 64 MiB. Firehose uses this value for padding calculations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-orcserde.html#cfn-kinesisfirehose-deliverystream-orcserde-blocksizebytes
	//
	BlockSizeBytes *float64 `field:"optional" json:"blockSizeBytes" yaml:"blockSizeBytes"`
	// The column names for which you want Firehose to create bloom filters.
	//
	// The default is `null` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-orcserde.html#cfn-kinesisfirehose-deliverystream-orcserde-bloomfiltercolumns
	//
	BloomFilterColumns *[]*string `field:"optional" json:"bloomFilterColumns" yaml:"bloomFilterColumns"`
	// The Bloom filter false positive probability (FPP).
	//
	// The lower the FPP, the bigger the Bloom filter. The default value is 0.05, the minimum is 0, and the maximum is 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-orcserde.html#cfn-kinesisfirehose-deliverystream-orcserde-bloomfilterfalsepositiveprobability
	//
	BloomFilterFalsePositiveProbability *float64 `field:"optional" json:"bloomFilterFalsePositiveProbability" yaml:"bloomFilterFalsePositiveProbability"`
	// The compression code to use over data blocks.
	//
	// The default is `SNAPPY` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-orcserde.html#cfn-kinesisfirehose-deliverystream-orcserde-compression
	//
	Compression *string `field:"optional" json:"compression" yaml:"compression"`
	// Represents the fraction of the total number of non-null rows.
	//
	// To turn off dictionary encoding, set this fraction to a number that is less than the number of distinct keys in a dictionary. To always use dictionary encoding, set this threshold to 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-orcserde.html#cfn-kinesisfirehose-deliverystream-orcserde-dictionarykeythreshold
	//
	DictionaryKeyThreshold *float64 `field:"optional" json:"dictionaryKeyThreshold" yaml:"dictionaryKeyThreshold"`
	// Set this to `true` to indicate that you want stripes to be padded to the HDFS block boundaries.
	//
	// This is useful if you intend to copy the data from Amazon S3 to HDFS before querying. The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-orcserde.html#cfn-kinesisfirehose-deliverystream-orcserde-enablepadding
	//
	EnablePadding interface{} `field:"optional" json:"enablePadding" yaml:"enablePadding"`
	// The version of the file to write.
	//
	// The possible values are `V0_11` and `V0_12` . The default is `V0_12` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-orcserde.html#cfn-kinesisfirehose-deliverystream-orcserde-formatversion
	//
	FormatVersion *string `field:"optional" json:"formatVersion" yaml:"formatVersion"`
	// A number between 0 and 1 that defines the tolerance for block padding as a decimal fraction of stripe size.
	//
	// The default value is 0.05, which means 5 percent of stripe size.
	//
	// For the default values of 64 MiB ORC stripes and 256 MiB HDFS blocks, the default block padding tolerance of 5 percent reserves a maximum of 3.2 MiB for padding within the 256 MiB block. In such a case, if the available size within the block is more than 3.2 MiB, a new, smaller stripe is inserted to fit within that space. This ensures that no stripe crosses block boundaries and causes remote reads within a node-local task.
	//
	// Kinesis Data Firehose ignores this parameter when `EnablePadding` is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-orcserde.html#cfn-kinesisfirehose-deliverystream-orcserde-paddingtolerance
	//
	PaddingTolerance *float64 `field:"optional" json:"paddingTolerance" yaml:"paddingTolerance"`
	// The number of rows between index entries.
	//
	// The default is 10,000 and the minimum is 1,000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-orcserde.html#cfn-kinesisfirehose-deliverystream-orcserde-rowindexstride
	//
	RowIndexStride *float64 `field:"optional" json:"rowIndexStride" yaml:"rowIndexStride"`
	// The number of bytes in each stripe.
	//
	// The default is 64 MiB and the minimum is 8 MiB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-orcserde.html#cfn-kinesisfirehose-deliverystream-orcserde-stripesizebytes
	//
	StripeSizeBytes *float64 `field:"optional" json:"stripeSizeBytes" yaml:"stripeSizeBytes"`
}

