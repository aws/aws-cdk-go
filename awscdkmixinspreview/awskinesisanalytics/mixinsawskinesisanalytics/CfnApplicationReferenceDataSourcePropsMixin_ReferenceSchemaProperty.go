package mixinsawskinesisanalytics


// The ReferenceSchema property type specifies the format of the data in the reference source for a SQL-based Amazon Kinesis Data Analytics application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   referenceSchemaProperty := &ReferenceSchemaProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-referenceschema.html
//
type CfnApplicationReferenceDataSourcePropsMixin_ReferenceSchemaProperty struct {
	// A list of RecordColumn objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-referenceschema.html#cfn-kinesisanalytics-applicationreferencedatasource-referenceschema-recordcolumns
	//
	RecordColumns interface{} `field:"optional" json:"recordColumns" yaml:"recordColumns"`
	// Specifies the encoding of the records in the reference source.
	//
	// For example, UTF-8.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-referenceschema.html#cfn-kinesisanalytics-applicationreferencedatasource-referenceschema-recordencoding
	//
	RecordEncoding *string `field:"optional" json:"recordEncoding" yaml:"recordEncoding"`
	// Specifies the format of the records on the reference source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-referenceschema.html#cfn-kinesisanalytics-applicationreferencedatasource-referenceschema-recordformat
	//
	RecordFormat interface{} `field:"optional" json:"recordFormat" yaml:"recordFormat"`
}

