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
//   inputFormatConfigurationProperty := &inputFormatConfigurationProperty{
//   	deserializer: &deserializerProperty{
//   		hiveJsonSerDe: &hiveJsonSerDeProperty{
//   			timestampFormats: []*string{
//   				jsii.String("timestampFormats"),
//   			},
//   		},
//   		openXJsonSerDe: &openXJsonSerDeProperty{
//   			caseInsensitive: jsii.Boolean(false),
//   			columnToJsonKeyMappings: map[string]*string{
//   				"columnToJsonKeyMappingsKey": jsii.String("columnToJsonKeyMappings"),
//   			},
//   			convertDotsInJsonKeysToUnderscores: jsii.Boolean(false),
//   		},
//   	},
//   }
//
type CfnDeliveryStream_InputFormatConfigurationProperty struct {
	// Specifies which deserializer to use.
	//
	// You can choose either the Apache Hive JSON SerDe or the OpenX JSON SerDe. If both are non-null, the server rejects the request.
	Deserializer interface{} `field:"optional" json:"deserializer" yaml:"deserializer"`
}

