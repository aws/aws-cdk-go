package awskinesisanalyticsv2


// When you configure a SQL-based Managed Service for Apache Flink application's input at the time of creating or updating an application, provides additional mapping information specific to the record format (such as JSON, CSV, or record fields delimited by some delimiter) on the streaming source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mappingParametersProperty := &MappingParametersProperty{
//   	CsvMappingParameters: &CSVMappingParametersProperty{
//   		RecordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   		RecordRowDelimiter: jsii.String("recordRowDelimiter"),
//   	},
//   	JsonMappingParameters: &JSONMappingParametersProperty{
//   		RecordRowPath: jsii.String("recordRowPath"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-mappingparameters.html
//
type CfnApplication_MappingParametersProperty struct {
	// Provides additional mapping information when the record format uses delimiters (for example, CSV).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-mappingparameters.html#cfn-kinesisanalyticsv2-application-mappingparameters-csvmappingparameters
	//
	CsvMappingParameters interface{} `field:"optional" json:"csvMappingParameters" yaml:"csvMappingParameters"`
	// Provides additional mapping information when JSON is the record format on the streaming source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-mappingparameters.html#cfn-kinesisanalyticsv2-application-mappingparameters-jsonmappingparameters
	//
	JsonMappingParameters interface{} `field:"optional" json:"jsonMappingParameters" yaml:"jsonMappingParameters"`
}

