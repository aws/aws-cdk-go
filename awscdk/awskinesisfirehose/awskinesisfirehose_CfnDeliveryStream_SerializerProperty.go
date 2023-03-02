package awskinesisfirehose


// The serializer that you want Kinesis Data Firehose to use to convert data to the target format before writing it to Amazon S3.
//
// Kinesis Data Firehose supports two types of serializers: the [ORC SerDe](https://docs.aws.amazon.com/https://hive.apache.org/javadocs/r1.2.2/api/org/apache/hadoop/hive/ql/io/orc/OrcSerde.html) and the [Parquet SerDe](https://docs.aws.amazon.com/https://hive.apache.org/javadocs/r1.2.2/api/org/apache/hadoop/hive/ql/io/parquet/serde/ParquetHiveSerDe.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serializerProperty := &serializerProperty{
//   	orcSerDe: &orcSerDeProperty{
//   		blockSizeBytes: jsii.Number(123),
//   		bloomFilterColumns: []*string{
//   			jsii.String("bloomFilterColumns"),
//   		},
//   		bloomFilterFalsePositiveProbability: jsii.Number(123),
//   		compression: jsii.String("compression"),
//   		dictionaryKeyThreshold: jsii.Number(123),
//   		enablePadding: jsii.Boolean(false),
//   		formatVersion: jsii.String("formatVersion"),
//   		paddingTolerance: jsii.Number(123),
//   		rowIndexStride: jsii.Number(123),
//   		stripeSizeBytes: jsii.Number(123),
//   	},
//   	parquetSerDe: &parquetSerDeProperty{
//   		blockSizeBytes: jsii.Number(123),
//   		compression: jsii.String("compression"),
//   		enableDictionaryCompression: jsii.Boolean(false),
//   		maxPaddingBytes: jsii.Number(123),
//   		pageSizeBytes: jsii.Number(123),
//   		writerVersion: jsii.String("writerVersion"),
//   	},
//   }
//
type CfnDeliveryStream_SerializerProperty struct {
	// A serializer to use for converting data to the ORC format before storing it in Amazon S3.
	//
	// For more information, see [Apache ORC](https://docs.aws.amazon.com/https://orc.apache.org/docs/) .
	OrcSerDe interface{} `field:"optional" json:"orcSerDe" yaml:"orcSerDe"`
	// A serializer to use for converting data to the Parquet format before storing it in Amazon S3.
	//
	// For more information, see [Apache Parquet](https://docs.aws.amazon.com/https://parquet.apache.org/documentation/latest/) .
	ParquetSerDe interface{} `field:"optional" json:"parquetSerDe" yaml:"parquetSerDe"`
}

