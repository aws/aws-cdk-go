package awskinesisfirehose


// Specifies the serializer that you want Firehose to use to convert the format of your data before it writes it to Amazon S3.
//
// This parameter is required if `Enabled` is set to true.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputFormatConfigurationProperty := &OutputFormatConfigurationProperty{
//   	Serializer: &SerializerProperty{
//   		OrcSerDe: &OrcSerDeProperty{
//   			BlockSizeBytes: jsii.Number(123),
//   			BloomFilterColumns: []*string{
//   				jsii.String("bloomFilterColumns"),
//   			},
//   			BloomFilterFalsePositiveProbability: jsii.Number(123),
//   			Compression: jsii.String("compression"),
//   			DictionaryKeyThreshold: jsii.Number(123),
//   			EnablePadding: jsii.Boolean(false),
//   			FormatVersion: jsii.String("formatVersion"),
//   			PaddingTolerance: jsii.Number(123),
//   			RowIndexStride: jsii.Number(123),
//   			StripeSizeBytes: jsii.Number(123),
//   		},
//   		ParquetSerDe: &ParquetSerDeProperty{
//   			BlockSizeBytes: jsii.Number(123),
//   			Compression: jsii.String("compression"),
//   			EnableDictionaryCompression: jsii.Boolean(false),
//   			MaxPaddingBytes: jsii.Number(123),
//   			PageSizeBytes: jsii.Number(123),
//   			WriterVersion: jsii.String("writerVersion"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-outputformatconfiguration.html
//
type CfnDeliveryStream_OutputFormatConfigurationProperty struct {
	// Specifies which serializer to use.
	//
	// You can choose either the ORC SerDe or the Parquet SerDe. If both are non-null, the server rejects the request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-outputformatconfiguration.html#cfn-kinesisfirehose-deliverystream-outputformatconfiguration-serializer
	//
	Serializer interface{} `field:"optional" json:"serializer" yaml:"serializer"`
}

