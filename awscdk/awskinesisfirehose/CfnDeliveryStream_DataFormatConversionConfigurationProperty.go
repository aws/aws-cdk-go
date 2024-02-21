package awskinesisfirehose


// Specifies that you want Kinesis Data Firehose to convert data from the JSON format to the Parquet or ORC format before writing it to Amazon S3.
//
// Kinesis Data Firehose uses the serializer and deserializer that you specify, in addition to the column information from the AWS Glue table, to deserialize your input data from JSON and then serialize it to the Parquet or ORC format. For more information, see [Kinesis Data Firehose Record Format Conversion](https://docs.aws.amazon.com/firehose/latest/dev/record-format-conversion.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataFormatConversionConfigurationProperty := &DataFormatConversionConfigurationProperty{
//   	Enabled: jsii.Boolean(false),
//   	InputFormatConfiguration: &InputFormatConfigurationProperty{
//   		Deserializer: &DeserializerProperty{
//   			HiveJsonSerDe: &HiveJsonSerDeProperty{
//   				TimestampFormats: []*string{
//   					jsii.String("timestampFormats"),
//   				},
//   			},
//   			OpenXJsonSerDe: &OpenXJsonSerDeProperty{
//   				CaseInsensitive: jsii.Boolean(false),
//   				ColumnToJsonKeyMappings: map[string]*string{
//   					"columnToJsonKeyMappingsKey": jsii.String("columnToJsonKeyMappings"),
//   				},
//   				ConvertDotsInJsonKeysToUnderscores: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	OutputFormatConfiguration: &OutputFormatConfigurationProperty{
//   		Serializer: &SerializerProperty{
//   			OrcSerDe: &OrcSerDeProperty{
//   				BlockSizeBytes: jsii.Number(123),
//   				BloomFilterColumns: []*string{
//   					jsii.String("bloomFilterColumns"),
//   				},
//   				BloomFilterFalsePositiveProbability: jsii.Number(123),
//   				Compression: jsii.String("compression"),
//   				DictionaryKeyThreshold: jsii.Number(123),
//   				EnablePadding: jsii.Boolean(false),
//   				FormatVersion: jsii.String("formatVersion"),
//   				PaddingTolerance: jsii.Number(123),
//   				RowIndexStride: jsii.Number(123),
//   				StripeSizeBytes: jsii.Number(123),
//   			},
//   			ParquetSerDe: &ParquetSerDeProperty{
//   				BlockSizeBytes: jsii.Number(123),
//   				Compression: jsii.String("compression"),
//   				EnableDictionaryCompression: jsii.Boolean(false),
//   				MaxPaddingBytes: jsii.Number(123),
//   				PageSizeBytes: jsii.Number(123),
//   				WriterVersion: jsii.String("writerVersion"),
//   			},
//   		},
//   	},
//   	SchemaConfiguration: &SchemaConfigurationProperty{
//   		CatalogId: jsii.String("catalogId"),
//   		DatabaseName: jsii.String("databaseName"),
//   		Region: jsii.String("region"),
//   		RoleArn: jsii.String("roleArn"),
//   		TableName: jsii.String("tableName"),
//   		VersionId: jsii.String("versionId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-dataformatconversionconfiguration.html
//
type CfnDeliveryStream_DataFormatConversionConfigurationProperty struct {
	// Defaults to `true` .
	//
	// Set it to `false` if you want to disable format conversion while preserving the configuration details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-dataformatconversionconfiguration.html#cfn-kinesisfirehose-deliverystream-dataformatconversionconfiguration-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Specifies the deserializer that you want Firehose to use to convert the format of your data from JSON.
	//
	// This parameter is required if `Enabled` is set to true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-dataformatconversionconfiguration.html#cfn-kinesisfirehose-deliverystream-dataformatconversionconfiguration-inputformatconfiguration
	//
	InputFormatConfiguration interface{} `field:"optional" json:"inputFormatConfiguration" yaml:"inputFormatConfiguration"`
	// Specifies the serializer that you want Firehose to use to convert the format of your data to the Parquet or ORC format.
	//
	// This parameter is required if `Enabled` is set to true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-dataformatconversionconfiguration.html#cfn-kinesisfirehose-deliverystream-dataformatconversionconfiguration-outputformatconfiguration
	//
	OutputFormatConfiguration interface{} `field:"optional" json:"outputFormatConfiguration" yaml:"outputFormatConfiguration"`
	// Specifies the AWS Glue Data Catalog table that contains the column information.
	//
	// This parameter is required if `Enabled` is set to true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-dataformatconversionconfiguration.html#cfn-kinesisfirehose-deliverystream-dataformatconversionconfiguration-schemaconfiguration
	//
	SchemaConfiguration interface{} `field:"optional" json:"schemaConfiguration" yaml:"schemaConfiguration"`
}

