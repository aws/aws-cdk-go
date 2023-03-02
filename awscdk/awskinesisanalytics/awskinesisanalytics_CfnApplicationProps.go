package awskinesisanalytics


// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &cfnApplicationProps{
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
//   					roleArn: jsii.String("roleArn"),
//   				},
//   			},
//   			kinesisFirehoseInput: &kinesisFirehoseInputProperty{
//   				resourceArn: jsii.String("resourceArn"),
//   				roleArn: jsii.String("roleArn"),
//   			},
//   			kinesisStreamsInput: &kinesisStreamsInputProperty{
//   				resourceArn: jsii.String("resourceArn"),
//   				roleArn: jsii.String("roleArn"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	applicationCode: jsii.String("applicationCode"),
//   	applicationDescription: jsii.String("applicationDescription"),
//   	applicationName: jsii.String("applicationName"),
//   }
//
type CfnApplicationProps struct {
	// Use this parameter to configure the application input.
	//
	// You can configure your application to receive input from a single streaming source. In this configuration, you map this streaming source to an in-application stream that is created. Your application code can then query the in-application stream like a table (you can think of it as a constantly updating table).
	//
	// For the streaming source, you provide its Amazon Resource Name (ARN) and format of data on the stream (for example, JSON, CSV, etc.). You also must provide an IAM role that Amazon Kinesis Analytics can assume to read this stream on your behalf.
	//
	// To create the in-application stream, you need to specify a schema to transform your data into a schematized version used in SQL. In the schema, you provide the necessary mapping of the data elements in the streaming source to record columns in the in-app stream.
	Inputs interface{} `field:"required" json:"inputs" yaml:"inputs"`
	// One or more SQL statements that read input data, transform it, and generate output.
	//
	// For example, you can write a SQL statement that reads data from one in-application stream, generates a running average of the number of advertisement clicks by vendor, and insert resulting rows in another in-application stream using pumps. For more information about the typical pattern, see [Application Code](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-app-code.html) .
	//
	// You can provide such series of SQL statements, where output of one statement can be used as the input for the next statement. You store intermediate results by creating in-application streams and pumps.
	//
	// Note that the application code must create the streams with names specified in the `Outputs` . For example, if your `Outputs` defines output streams named `ExampleOutputStream1` and `ExampleOutputStream2` , then your application code must create these streams.
	ApplicationCode *string `field:"optional" json:"applicationCode" yaml:"applicationCode"`
	// Summary description of the application.
	ApplicationDescription *string `field:"optional" json:"applicationDescription" yaml:"applicationDescription"`
	// Name of your Amazon Kinesis Analytics application (for example, `sample-app` ).
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
}

