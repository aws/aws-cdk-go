package awskinesisanalytics


// The ReferenceSchema property type specifies the format of the data in the reference source for a SQL-based Amazon Kinesis Data Analytics application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   referenceSchemaProperty := &referenceSchemaProperty{
//   	recordColumns: []interface{}{
//   		&recordColumnProperty{
//   			name: jsii.String("name"),
//   			sqlType: jsii.String("sqlType"),
//
//   			// the properties below are optional
//   			mapping: jsii.String("mapping"),
//   		},
//   	},
//   	recordFormat: &recordFormatProperty{
//   		recordFormatType: jsii.String("recordFormatType"),
//
//   		// the properties below are optional
//   		mappingParameters: &mappingParametersProperty{
//   			csvMappingParameters: &cSVMappingParametersProperty{
//   				recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   				recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   			},
//   			jsonMappingParameters: &jSONMappingParametersProperty{
//   				recordRowPath: jsii.String("recordRowPath"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	recordEncoding: jsii.String("recordEncoding"),
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

