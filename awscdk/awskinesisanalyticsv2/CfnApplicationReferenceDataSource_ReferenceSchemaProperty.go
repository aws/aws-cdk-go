package awskinesisanalyticsv2


// For a SQL-based Managed Service for Apache Flink application, describes the format of the data in the streaming source, and how each data element maps to corresponding columns created in the in-application stream.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationreferencedatasource-referenceschema.html
//
type CfnApplicationReferenceDataSource_ReferenceSchemaProperty struct {
	// A list of `RecordColumn` objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationreferencedatasource-referenceschema.html#cfn-kinesisanalyticsv2-applicationreferencedatasource-referenceschema-recordcolumns
	//
	RecordColumns interface{} `field:"required" json:"recordColumns" yaml:"recordColumns"`
	// Specifies the format of the records on the streaming source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationreferencedatasource-referenceschema.html#cfn-kinesisanalyticsv2-applicationreferencedatasource-referenceschema-recordformat
	//
	RecordFormat interface{} `field:"required" json:"recordFormat" yaml:"recordFormat"`
	// Specifies the encoding of the records in the streaming source.
	//
	// For example, UTF-8.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationreferencedatasource-referenceschema.html#cfn-kinesisanalyticsv2-applicationreferencedatasource-referenceschema-recordencoding
	//
	RecordEncoding *string `field:"optional" json:"recordEncoding" yaml:"recordEncoding"`
}

