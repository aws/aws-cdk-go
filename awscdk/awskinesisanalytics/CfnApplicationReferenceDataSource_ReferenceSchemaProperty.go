package awskinesisanalytics


// The ReferenceSchema property type specifies the format of the data in the reference source for a SQL-based Amazon Kinesis Data Analytics application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   referenceSchemaProperty := &ReferenceSchemaProperty{
//   	RecordColumns: []interface{}{
//   		&RecordColumnProperty{
//   			Name: jsii.String("name"),
//   			SqlType: jsii.String("sqlType"),
//
//   			// the properties below are optional
//   			Mapping: jsii.String("mapping"),
//   		},
//   	},
//   	RecordFormat: &RecordFormatProperty{
//   		RecordFormatType: jsii.String("recordFormatType"),
//
//   		// the properties below are optional
//   		MappingParameters: &MappingParametersProperty{
//   			CsvMappingParameters: &CSVMappingParametersProperty{
//   				RecordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   				RecordRowDelimiter: jsii.String("recordRowDelimiter"),
//   			},
//   			JsonMappingParameters: &JSONMappingParametersProperty{
//   				RecordRowPath: jsii.String("recordRowPath"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	RecordEncoding: jsii.String("recordEncoding"),
//   }
//
type CfnApplicationReferenceDataSource_ReferenceSchemaProperty struct {
	// A list of RecordColumn objects.
	RecordColumns interface{} `field:"required" json:"recordColumns" yaml:"recordColumns"`
	// Specifies the format of the records on the reference source.
	RecordFormat interface{} `field:"required" json:"recordFormat" yaml:"recordFormat"`
	// Specifies the encoding of the records in the reference source.
	//
	// For example, UTF-8.
	RecordEncoding *string `field:"optional" json:"recordEncoding" yaml:"recordEncoding"`
}

