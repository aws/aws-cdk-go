package awskinesisanalytics


// Describes the inputs, outputs, and reference data sources for a SQL-based Kinesis Data Analytics application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sqlApplicationConfigurationProperty := &SqlApplicationConfigurationProperty{
//   	Inputs: []interface{}{
//   		&InputProperty{
//   			InputSchema: &InputSchemaProperty{
//   				RecordColumns: []interface{}{
//   					&RecordColumnProperty{
//   						Name: jsii.String("name"),
//   						SqlType: jsii.String("sqlType"),
//
//   						// the properties below are optional
//   						Mapping: jsii.String("mapping"),
//   					},
//   				},
//   				RecordFormat: &RecordFormatProperty{
//   					RecordFormatType: jsii.String("recordFormatType"),
//
//   					// the properties below are optional
//   					MappingParameters: &MappingParametersProperty{
//   						CsvMappingParameters: &CSVMappingParametersProperty{
//   							RecordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   							RecordRowDelimiter: jsii.String("recordRowDelimiter"),
//   						},
//   						JsonMappingParameters: &JSONMappingParametersProperty{
//   							RecordRowPath: jsii.String("recordRowPath"),
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				RecordEncoding: jsii.String("recordEncoding"),
//   			},
//   			NamePrefix: jsii.String("namePrefix"),
//
//   			// the properties below are optional
//   			InputParallelism: &InputParallelismProperty{
//   				Count: jsii.Number(123),
//   			},
//   			InputProcessingConfiguration: &InputProcessingConfigurationProperty{
//   				InputLambdaProcessor: &InputLambdaProcessorProperty{
//   					ResourceArn: jsii.String("resourceArn"),
//   				},
//   			},
//   			KinesisFirehoseInput: &KinesisFirehoseInputProperty{
//   				ResourceArn: jsii.String("resourceArn"),
//   			},
//   			KinesisStreamsInput: &KinesisStreamsInputProperty{
//   				ResourceArn: jsii.String("resourceArn"),
//   			},
//   		},
//   	},
//   }
//
type CfnApplicationV2_SqlApplicationConfigurationProperty struct {
	// The array of [Input](https://docs.aws.amazon.com/kinesisanalytics/latest/apiv2/API_Input.html) objects describing the input streams used by the application.
	Inputs interface{} `field:"optional" json:"inputs" yaml:"inputs"`
}

