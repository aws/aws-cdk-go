package awskinesisfirehose


// The native Hive / HCatalog JsonSerDe.
//
// Used by Kinesis Data Firehose for deserializing data, which means converting it from the JSON format in preparation for serializing it to the Parquet or ORC format. This is one of two deserializers you can choose, depending on which one offers the functionality you need. The other option is the OpenX SerDe.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hiveJsonSerDeProperty := &hiveJsonSerDeProperty{
//   	timestampFormats: []*string{
//   		jsii.String("timestampFormats"),
//   	},
//   }
//
type CfnDeliveryStream_HiveJsonSerDeProperty struct {
	// Indicates how you want Kinesis Data Firehose to parse the date and timestamps that may be present in your input data JSON.
	//
	// To specify these format strings, follow the pattern syntax of JodaTime's DateTimeFormat format strings. For more information, see [Class DateTimeFormat](https://docs.aws.amazon.com/https://www.joda.org/joda-time/apidocs/org/joda/time/format/DateTimeFormat.html) . You can also use the special value `millis` to parse timestamps in epoch milliseconds. If you don't specify a format, Kinesis Data Firehose uses `java.sql.Timestamp::valueOf` by default.
	TimestampFormats *[]*string `field:"optional" json:"timestampFormats" yaml:"timestampFormats"`
}

