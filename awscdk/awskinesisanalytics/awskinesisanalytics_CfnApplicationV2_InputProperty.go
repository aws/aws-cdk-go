package awskinesisanalytics


// When you configure the application input for a SQL-based Kinesis Data Analytics application, you specify the streaming source, the in-application stream name that is created, and the mapping between the two.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputProperty := &inputProperty{
//   	inputSchema: &inputSchemaProperty{
//   		recordColumns: []interface{}{
//   			&recordColumnProperty{
//   				name: jsii.String("name"),
//   				sqlType: jsii.String("sqlType"),
//
//   				// the properties below are optional
//   				mapping: jsii.String("mapping"),
//   			},
//   		},
//   		recordFormat: &recordFormatProperty{
//   			recordFormatType: jsii.String("recordFormatType"),
//
//   			// the properties below are optional
//   			mappingParameters: &mappingParametersProperty{
//   				csvMappingParameters: &cSVMappingParametersProperty{
//   					recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   					recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   				},
//   				jsonMappingParameters: &jSONMappingParametersProperty{
//   					recordRowPath: jsii.String("recordRowPath"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		recordEncoding: jsii.String("recordEncoding"),
//   	},
//   	namePrefix: jsii.String("namePrefix"),
//
//   	// the properties below are optional
//   	inputParallelism: &inputParallelismProperty{
//   		count: jsii.Number(123),
//   	},
//   	inputProcessingConfiguration: &inputProcessingConfigurationProperty{
//   		inputLambdaProcessor: &inputLambdaProcessorProperty{
//   			resourceArn: jsii.String("resourceArn"),
//   		},
//   	},
//   	kinesisFirehoseInput: &kinesisFirehoseInputProperty{
//   		resourceArn: jsii.String("resourceArn"),
//   	},
//   	kinesisStreamsInput: &kinesisStreamsInputProperty{
//   		resourceArn: jsii.String("resourceArn"),
//   	},
//   }
//
type CfnApplicationV2_InputProperty struct {
	// Describes the format of the data in the streaming source, and how each data element maps to corresponding columns in the in-application stream that is being created.
	//
	// Also used to describe the format of the reference data source.
	InputSchema interface{} `field:"required" json:"inputSchema" yaml:"inputSchema"`
	// The name prefix to use when creating an in-application stream.
	//
	// Suppose that you specify a prefix " `MyInApplicationStream` ." Kinesis Data Analytics then creates one or more (as per the `InputParallelism` count you specified) in-application streams with the names " `MyInApplicationStream_001` ," " `MyInApplicationStream_002` ," and so on.
	NamePrefix *string `field:"required" json:"namePrefix" yaml:"namePrefix"`
	// Describes the number of in-application streams to create.
	InputParallelism interface{} `field:"optional" json:"inputParallelism" yaml:"inputParallelism"`
	// The [InputProcessingConfiguration](https://docs.aws.amazon.com/kinesisanalytics/latest/apiv2/API_InputProcessingConfiguration.html) for the input. An input processor transforms records as they are received from the stream, before the application's SQL code executes. Currently, the only input processing configuration available is [InputLambdaProcessor](https://docs.aws.amazon.com/kinesisanalytics/latest/apiv2/API_InputLambdaProcessor.html) .
	InputProcessingConfiguration interface{} `field:"optional" json:"inputProcessingConfiguration" yaml:"inputProcessingConfiguration"`
	// If the streaming source is an Amazon Kinesis Data Firehose delivery stream, identifies the delivery stream's ARN.
	KinesisFirehoseInput interface{} `field:"optional" json:"kinesisFirehoseInput" yaml:"kinesisFirehoseInput"`
	// If the streaming source is an Amazon Kinesis data stream, identifies the stream's Amazon Resource Name (ARN).
	KinesisStreamsInput interface{} `field:"optional" json:"kinesisStreamsInput" yaml:"kinesisStreamsInput"`
}

