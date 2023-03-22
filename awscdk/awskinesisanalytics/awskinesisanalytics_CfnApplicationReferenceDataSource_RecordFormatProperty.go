package awskinesisanalytics


// Describes the record format and relevant mapping information that should be applied to schematize the records on the stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recordFormatProperty := &recordFormatProperty{
//   	recordFormatType: jsii.String("recordFormatType"),
//
//   	// the properties below are optional
//   	mappingParameters: &mappingParametersProperty{
//   		csvMappingParameters: &cSVMappingParametersProperty{
//   			recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   			recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   		},
//   		jsonMappingParameters: &jSONMappingParametersProperty{
//   			recordRowPath: jsii.String("recordRowPath"),
//   		},
//   	},
//   }
//
type CfnApplicationReferenceDataSource_RecordFormatProperty struct {
	// The type of record format.
	RecordFormatType *string `field:"required" json:"recordFormatType" yaml:"recordFormatType"`
	// When configuring application input at the time of creating or updating an application, provides additional mapping information specific to the record format (such as JSON, CSV, or record fields delimited by some delimiter) on the streaming source.
	MappingParameters interface{} `field:"optional" json:"mappingParameters" yaml:"mappingParameters"`
}

