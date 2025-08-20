package awskinesisanalytics


// Describes the record format and relevant mapping information that should be applied to schematize the records on the stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recordFormatProperty := &RecordFormatProperty{
//   	RecordFormatType: jsii.String("recordFormatType"),
//
//   	// the properties below are optional
//   	MappingParameters: &MappingParametersProperty{
//   		CsvMappingParameters: &CSVMappingParametersProperty{
//   			RecordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   			RecordRowDelimiter: jsii.String("recordRowDelimiter"),
//   		},
//   		JsonMappingParameters: &JSONMappingParametersProperty{
//   			RecordRowPath: jsii.String("recordRowPath"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-recordformat.html
//
type CfnApplication_RecordFormatProperty struct {
	// The type of record format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-recordformat.html#cfn-kinesisanalytics-application-recordformat-recordformattype
	//
	RecordFormatType *string `field:"required" json:"recordFormatType" yaml:"recordFormatType"`
	// When configuring application input at the time of creating or updating an application, provides additional mapping information specific to the record format (such as JSON, CSV, or record fields delimited by some delimiter) on the streaming source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-recordformat.html#cfn-kinesisanalytics-application-recordformat-mappingparameters
	//
	MappingParameters interface{} `field:"optional" json:"mappingParameters" yaml:"mappingParameters"`
}

