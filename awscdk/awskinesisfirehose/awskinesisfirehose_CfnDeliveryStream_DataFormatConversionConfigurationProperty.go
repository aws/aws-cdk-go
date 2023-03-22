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
//   dataFormatConversionConfigurationProperty := &dataFormatConversionConfigurationProperty{
//   	enabled: jsii.Boolean(false),
//   	inputFormatConfiguration: &inputFormatConfigurationProperty{
//   		deserializer: &deserializerProperty{
//   			hiveJsonSerDe: &hiveJsonSerDeProperty{
//   				timestampFormats: []*string{
//   					jsii.String("timestampFormats"),
//   				},
//   			},
//   			openXJsonSerDe: &openXJsonSerDeProperty{
//   				caseInsensitive: jsii.Boolean(false),
//   				columnToJsonKeyMappings: map[string]*string{
//   					"columnToJsonKeyMappingsKey": jsii.String("columnToJsonKeyMappings"),
//   				},
//   				convertDotsInJsonKeysToUnderscores: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	outputFormatConfiguration: &outputFormatConfigurationProperty{
//   		serializer: &serializerProperty{
//   			orcSerDe: &orcSerDeProperty{
//   				blockSizeBytes: jsii.Number(123),
//   				bloomFilterColumns: []*string{
//   					jsii.String("bloomFilterColumns"),
//   				},
//   				bloomFilterFalsePositiveProbability: jsii.Number(123),
//   				compression: jsii.String("compression"),
//   				dictionaryKeyThreshold: jsii.Number(123),
//   				enablePadding: jsii.Boolean(false),
//   				formatVersion: jsii.String("formatVersion"),
//   				paddingTolerance: jsii.Number(123),
//   				rowIndexStride: jsii.Number(123),
//   				stripeSizeBytes: jsii.Number(123),
//   			},
//   			parquetSerDe: &parquetSerDeProperty{
//   				blockSizeBytes: jsii.Number(123),
//   				compression: jsii.String("compression"),
//   				enableDictionaryCompression: jsii.Boolean(false),
//   				maxPaddingBytes: jsii.Number(123),
//   				pageSizeBytes: jsii.Number(123),
//   				writerVersion: jsii.String("writerVersion"),
//   			},
//   		},
//   	},
//   	schemaConfiguration: &schemaConfigurationProperty{
//   		catalogId: jsii.String("catalogId"),
//   		databaseName: jsii.String("databaseName"),
//   		region: jsii.String("region"),
//   		roleArn: jsii.String("roleArn"),
//   		tableName: jsii.String("tableName"),
//   		versionId: jsii.String("versionId"),
//   	},
//   }
//
type CfnDeliveryStream_DataFormatConversionConfigurationProperty struct {
	// Defaults to `true` .
	//
	// Set it to `false` if you want to disable format conversion while preserving the configuration details.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Specifies the deserializer that you want Kinesis Data Firehose to use to convert the format of your data from JSON.
	//
	// This parameter is required if `Enabled` is set to true.
	InputFormatConfiguration interface{} `field:"optional" json:"inputFormatConfiguration" yaml:"inputFormatConfiguration"`
	// Specifies the serializer that you want Kinesis Data Firehose to use to convert the format of your data to the Parquet or ORC format.
	//
	// This parameter is required if `Enabled` is set to true.
	OutputFormatConfiguration interface{} `field:"optional" json:"outputFormatConfiguration" yaml:"outputFormatConfiguration"`
	// Specifies the AWS Glue Data Catalog table that contains the column information.
	//
	// This parameter is required if `Enabled` is set to true.
	SchemaConfiguration interface{} `field:"optional" json:"schemaConfiguration" yaml:"schemaConfiguration"`
}

