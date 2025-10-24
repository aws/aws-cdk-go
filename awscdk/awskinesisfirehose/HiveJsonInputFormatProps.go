package awskinesisfirehose


// Props for Hive JSON input format for data record format conversion.
//
// Example:
//   inputFormat := firehose.NewHiveJsonInputFormat(&HiveJsonInputFormatProps{
//   	TimestampParsers: []TimestampParser{
//   		firehose.TimestampParser_FromFormatString(jsii.String("yyyy-MM-dd")),
//   		firehose.TimestampParser_EPOCH_MILLIS(),
//   	},
//   })
//
type HiveJsonInputFormatProps struct {
	// List of TimestampParsers.
	//
	// These are used to parse custom timestamp strings from input JSON into dates.
	//
	// Note: Specifying a parser will override the default timestamp parser. If the default timestamp parser is required,
	//  include `TimestampParser.DEFAULT` in the list of parsers along with the custom parser.
	// Default: the default timestamp parser is used.
	//
	TimestampParsers *[]TimestampParser `field:"optional" json:"timestampParsers" yaml:"timestampParsers"`
}

