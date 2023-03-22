package awskinesisanalyticsv2


// Describes the inputs, outputs, and reference data sources for a SQL-based Kinesis Data Analytics application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sqlApplicationConfigurationProperty := &sqlApplicationConfigurationProperty{
//   	inputs: []interface{}{
//   		&inputProperty{
//   			inputSchema: &inputSchemaProperty{
//   				recordColumns: []interface{}{
//   					&recordColumnProperty{
//   						name: jsii.String("name"),
//   						sqlType: jsii.String("sqlType"),
//
//   						// the properties below are optional
//   						mapping: jsii.String("mapping"),
//   					},
//   				},
//   				recordFormat: &recordFormatProperty{
//   					recordFormatType: jsii.String("recordFormatType"),
//
//   					// the properties below are optional
//   					mappingParameters: &mappingParametersProperty{
//   						csvMappingParameters: &cSVMappingParametersProperty{
//   							recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   							recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   						},
//   						jsonMappingParameters: &jSONMappingParametersProperty{
//   							recordRowPath: jsii.String("recordRowPath"),
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				recordEncoding: jsii.String("recordEncoding"),
//   			},
//   			namePrefix: jsii.String("namePrefix"),
//
//   			// the properties below are optional
//   			inputParallelism: &inputParallelismProperty{
//   				count: jsii.Number(123),
//   			},
//   			inputProcessingConfiguration: &inputProcessingConfigurationProperty{
//   				inputLambdaProcessor: &inputLambdaProcessorProperty{
//   					resourceArn: jsii.String("resourceArn"),
//   				},
//   			},
//   			kinesisFirehoseInput: &kinesisFirehoseInputProperty{
//   				resourceArn: jsii.String("resourceArn"),
//   			},
//   			kinesisStreamsInput: &kinesisStreamsInputProperty{
//   				resourceArn: jsii.String("resourceArn"),
//   			},
//   		},
//   	},
//   }
//
type CfnApplication_SqlApplicationConfigurationProperty struct {
	// The array of [Input](https://docs.aws.amazon.com/kinesisanalytics/latest/apiv2/API_Input.html) objects describing the input streams used by the application.
	Inputs interface{} `field:"optional" json:"inputs" yaml:"inputs"`
}

