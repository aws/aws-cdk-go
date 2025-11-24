package mixinsawskinesisanalyticsv2


// For a SQL-based Kinesis Data Analytics application, describes the format of the data in the streaming source, and how each data element maps to corresponding columns created in the in-application stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   inputSchemaProperty := &InputSchemaProperty{
//   	RecordColumns: []interface{}{
//   		&RecordColumnProperty{
//   			Mapping: jsii.String("mapping"),
//   			Name: jsii.String("name"),
//   			SqlType: jsii.String("sqlType"),
//   		},
//   	},
//   	RecordEncoding: jsii.String("recordEncoding"),
//   	RecordFormat: &RecordFormatProperty{
//   		MappingParameters: &MappingParametersProperty{
//   			CsvMappingParameters: &CSVMappingParametersProperty{
//   				RecordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   				RecordRowDelimiter: jsii.String("recordRowDelimiter"),
//   			},
//   			JsonMappingParameters: &JSONMappingParametersProperty{
//   				RecordRowPath: jsii.String("recordRowPath"),
//   			},
//   		},
//   		RecordFormatType: jsii.String("recordFormatType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-inputschema.html
//
type CfnApplicationPropsMixin_InputSchemaProperty struct {
	// A list of `RecordColumn` objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-inputschema.html#cfn-kinesisanalyticsv2-application-inputschema-recordcolumns
	//
	RecordColumns interface{} `field:"optional" json:"recordColumns" yaml:"recordColumns"`
	// Specifies the encoding of the records in the streaming source.
	//
	// For example, UTF-8.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-inputschema.html#cfn-kinesisanalyticsv2-application-inputschema-recordencoding
	//
	RecordEncoding *string `field:"optional" json:"recordEncoding" yaml:"recordEncoding"`
	// Specifies the format of the records on the streaming source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-inputschema.html#cfn-kinesisanalyticsv2-application-inputschema-recordformat
	//
	RecordFormat interface{} `field:"optional" json:"recordFormat" yaml:"recordFormat"`
}

