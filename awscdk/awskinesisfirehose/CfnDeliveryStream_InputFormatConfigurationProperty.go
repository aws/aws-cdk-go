package awskinesisfirehose


// Specifies the deserializer you want to use to convert the format of the input data.
//
// This parameter is required if `Enabled` is set to true.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputFormatConfigurationProperty := &InputFormatConfigurationProperty{
//   	Deserializer: &DeserializerProperty{
//   		HiveJsonSerDe: &HiveJsonSerDeProperty{
//   			TimestampFormats: []*string{
//   				jsii.String("timestampFormats"),
//   			},
//   		},
//   		OpenXJsonSerDe: &OpenXJsonSerDeProperty{
//   			CaseInsensitive: jsii.Boolean(false),
//   			ColumnToJsonKeyMappings: map[string]*string{
//   				"columnToJsonKeyMappingsKey": jsii.String("columnToJsonKeyMappings"),
//   			},
//   			ConvertDotsInJsonKeysToUnderscores: jsii.Boolean(false),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-inputformatconfiguration.html
//
type CfnDeliveryStream_InputFormatConfigurationProperty struct {
	// Specifies which deserializer to use.
	//
	// You can choose either the Apache Hive JSON SerDe or the OpenX JSON SerDe. If both are non-null, the server rejects the request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-inputformatconfiguration.html#cfn-kinesisfirehose-deliverystream-inputformatconfiguration-deserializer
	//
	Deserializer interface{} `field:"optional" json:"deserializer" yaml:"deserializer"`
}

