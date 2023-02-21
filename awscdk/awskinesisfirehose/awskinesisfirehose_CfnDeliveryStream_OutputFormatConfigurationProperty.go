package awskinesisfirehose


// Specifies the serializer that you want Kinesis Data Firehose to use to convert the format of your data before it writes it to Amazon S3.
//
// This parameter is required if `Enabled` is set to true.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputFormatConfigurationProperty := &outputFormatConfigurationProperty{
//   	serializer: &serializerProperty{
//   		orcSerDe: &orcSerDeProperty{
//   			blockSizeBytes: jsii.Number(123),
//   			bloomFilterColumns: []*string{
//   				jsii.String("bloomFilterColumns"),
//   			},
//   			bloomFilterFalsePositiveProbability: jsii.Number(123),
//   			compression: jsii.String("compression"),
//   			dictionaryKeyThreshold: jsii.Number(123),
//   			enablePadding: jsii.Boolean(false),
//   			formatVersion: jsii.String("formatVersion"),
//   			paddingTolerance: jsii.Number(123),
//   			rowIndexStride: jsii.Number(123),
//   			stripeSizeBytes: jsii.Number(123),
//   		},
//   		parquetSerDe: &parquetSerDeProperty{
//   			blockSizeBytes: jsii.Number(123),
//   			compression: jsii.String("compression"),
//   			enableDictionaryCompression: jsii.Boolean(false),
//   			maxPaddingBytes: jsii.Number(123),
//   			pageSizeBytes: jsii.Number(123),
//   			writerVersion: jsii.String("writerVersion"),
//   		},
//   	},
//   }
//
type CfnDeliveryStream_OutputFormatConfigurationProperty struct {
	// Specifies which serializer to use.
	//
	// You can choose either the ORC SerDe or the Parquet SerDe. If both are non-null, the server rejects the request.
	Serializer interface{} `field:"optional" json:"serializer" yaml:"serializer"`
}

