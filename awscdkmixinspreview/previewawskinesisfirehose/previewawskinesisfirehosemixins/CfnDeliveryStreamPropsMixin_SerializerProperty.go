package previewawskinesisfirehosemixins


// The serializer that you want Firehose to use to convert data to the target format before writing it to Amazon S3.
//
// Firehose supports two types of serializers: the ORC SerDe and the Parquet SerDe.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   serializerProperty := &SerializerProperty{
//   	OrcSerDe: &OrcSerDeProperty{
//   		BlockSizeBytes: jsii.Number(123),
//   		BloomFilterColumns: []*string{
//   			jsii.String("bloomFilterColumns"),
//   		},
//   		BloomFilterFalsePositiveProbability: jsii.Number(123),
//   		Compression: jsii.String("compression"),
//   		DictionaryKeyThreshold: jsii.Number(123),
//   		EnablePadding: jsii.Boolean(false),
//   		FormatVersion: jsii.String("formatVersion"),
//   		PaddingTolerance: jsii.Number(123),
//   		RowIndexStride: jsii.Number(123),
//   		StripeSizeBytes: jsii.Number(123),
//   	},
//   	ParquetSerDe: &ParquetSerDeProperty{
//   		BlockSizeBytes: jsii.Number(123),
//   		Compression: jsii.String("compression"),
//   		EnableDictionaryCompression: jsii.Boolean(false),
//   		MaxPaddingBytes: jsii.Number(123),
//   		PageSizeBytes: jsii.Number(123),
//   		WriterVersion: jsii.String("writerVersion"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-serializer.html
//
type CfnDeliveryStreamPropsMixin_SerializerProperty struct {
	// A serializer to use for converting data to the ORC format before storing it in Amazon S3.
	//
	// For more information, see [Apache ORC](https://docs.aws.amazon.com/https://orc.apache.org/docs/) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-serializer.html#cfn-kinesisfirehose-deliverystream-serializer-orcserde
	//
	OrcSerDe interface{} `field:"optional" json:"orcSerDe" yaml:"orcSerDe"`
	// A serializer to use for converting data to the Parquet format before storing it in Amazon S3.
	//
	// For more information, see [Apache Parquet](https://docs.aws.amazon.com/https://parquet.apache.org/docs/contribution-guidelines/) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-serializer.html#cfn-kinesisfirehose-deliverystream-serializer-parquetserde
	//
	ParquetSerDe interface{} `field:"optional" json:"parquetSerDe" yaml:"parquetSerDe"`
}

