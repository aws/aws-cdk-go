package mixinsawskinesisanalyticsv2


// Describes the inputs, outputs, and reference data sources for a SQL-based Kinesis Data Analytics application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sqlApplicationConfigurationProperty := &SqlApplicationConfigurationProperty{
//   	Inputs: []interface{}{
//   		&InputProperty{
//   			InputParallelism: &InputParallelismProperty{
//   				Count: jsii.Number(123),
//   			},
//   			InputProcessingConfiguration: &InputProcessingConfigurationProperty{
//   				InputLambdaProcessor: &InputLambdaProcessorProperty{
//   					ResourceArn: jsii.String("resourceArn"),
//   				},
//   			},
//   			InputSchema: &InputSchemaProperty{
//   				RecordColumns: []interface{}{
//   					&RecordColumnProperty{
//   						Mapping: jsii.String("mapping"),
//   						Name: jsii.String("name"),
//   						SqlType: jsii.String("sqlType"),
//   					},
//   				},
//   				RecordEncoding: jsii.String("recordEncoding"),
//   				RecordFormat: &RecordFormatProperty{
//   					MappingParameters: &MappingParametersProperty{
//   						CsvMappingParameters: &CSVMappingParametersProperty{
//   							RecordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   							RecordRowDelimiter: jsii.String("recordRowDelimiter"),
//   						},
//   						JsonMappingParameters: &JSONMappingParametersProperty{
//   							RecordRowPath: jsii.String("recordRowPath"),
//   						},
//   					},
//   					RecordFormatType: jsii.String("recordFormatType"),
//   				},
//   			},
//   			KinesisFirehoseInput: &KinesisFirehoseInputProperty{
//   				ResourceArn: jsii.String("resourceArn"),
//   			},
//   			KinesisStreamsInput: &KinesisStreamsInputProperty{
//   				ResourceArn: jsii.String("resourceArn"),
//   			},
//   			NamePrefix: jsii.String("namePrefix"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-sqlapplicationconfiguration.html
//
type CfnApplicationPropsMixin_SqlApplicationConfigurationProperty struct {
	// The array of [Input](https://docs.aws.amazon.com/managed-flink/latest/apiv2/API_Input.html) objects describing the input streams used by the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-sqlapplicationconfiguration.html#cfn-kinesisanalyticsv2-application-sqlapplicationconfiguration-inputs
	//
	Inputs interface{} `field:"optional" json:"inputs" yaml:"inputs"`
}

