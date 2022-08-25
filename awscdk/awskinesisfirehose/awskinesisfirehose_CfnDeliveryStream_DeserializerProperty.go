package awskinesisfirehose


// The deserializer you want Kinesis Data Firehose to use for converting the input data from JSON.
//
// Kinesis Data Firehose then serializes the data to its final format using the `Serializer` . Kinesis Data Firehose supports two types of deserializers: the [Apache Hive JSON SerDe](https://docs.aws.amazon.com/https://cwiki.apache.org/confluence/display/Hive/LanguageManual+DDL#LanguageManualDDL-JSON) and the [OpenX JSON SerDe](https://docs.aws.amazon.com/https://github.com/rcongiu/Hive-JSON-Serde) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deserializerProperty := &deserializerProperty{
//   	hiveJsonSerDe: &hiveJsonSerDeProperty{
//   		timestampFormats: []*string{
//   			jsii.String("timestampFormats"),
//   		},
//   	},
//   	openXJsonSerDe: &openXJsonSerDeProperty{
//   		caseInsensitive: jsii.Boolean(false),
//   		columnToJsonKeyMappings: map[string]*string{
//   			"columnToJsonKeyMappingsKey": jsii.String("columnToJsonKeyMappings"),
//   		},
//   		convertDotsInJsonKeysToUnderscores: jsii.Boolean(false),
//   	},
//   }
//
type CfnDeliveryStream_DeserializerProperty struct {
	// The native Hive / HCatalog JsonSerDe.
	//
	// Used by Kinesis Data Firehose for deserializing data, which means converting it from the JSON format in preparation for serializing it to the Parquet or ORC format. This is one of two deserializers you can choose, depending on which one offers the functionality you need. The other option is the OpenX SerDe.
	HiveJsonSerDe interface{} `field:"optional" json:"hiveJsonSerDe" yaml:"hiveJsonSerDe"`
	// The OpenX SerDe.
	//
	// Used by Kinesis Data Firehose for deserializing data, which means converting it from the JSON format in preparation for serializing it to the Parquet or ORC format. This is one of two deserializers you can choose, depending on which one offers the functionality you need. The other option is the native Hive / HCatalog JsonSerDe.
	OpenXJsonSerDe interface{} `field:"optional" json:"openXJsonSerDe" yaml:"openXJsonSerDe"`
}

