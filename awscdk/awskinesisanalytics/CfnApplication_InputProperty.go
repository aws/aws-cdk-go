package awskinesisanalytics


// When you configure the application input, you specify the streaming source, the in-application stream name that is created, and the mapping between the two.
//
// For more information, see [Configuring Application Input](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-input.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputProperty := &InputProperty{
//   	InputSchema: &InputSchemaProperty{
//   		RecordColumns: []interface{}{
//   			&RecordColumnProperty{
//   				Name: jsii.String("name"),
//   				SqlType: jsii.String("sqlType"),
//
//   				// the properties below are optional
//   				Mapping: jsii.String("mapping"),
//   			},
//   		},
//   		RecordFormat: &RecordFormatProperty{
//   			RecordFormatType: jsii.String("recordFormatType"),
//
//   			// the properties below are optional
//   			MappingParameters: &MappingParametersProperty{
//   				CsvMappingParameters: &CSVMappingParametersProperty{
//   					RecordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   					RecordRowDelimiter: jsii.String("recordRowDelimiter"),
//   				},
//   				JsonMappingParameters: &JSONMappingParametersProperty{
//   					RecordRowPath: jsii.String("recordRowPath"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		RecordEncoding: jsii.String("recordEncoding"),
//   	},
//   	NamePrefix: jsii.String("namePrefix"),
//
//   	// the properties below are optional
//   	InputParallelism: &InputParallelismProperty{
//   		Count: jsii.Number(123),
//   	},
//   	InputProcessingConfiguration: &InputProcessingConfigurationProperty{
//   		InputLambdaProcessor: &InputLambdaProcessorProperty{
//   			ResourceArn: jsii.String("resourceArn"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   	},
//   	KinesisFirehoseInput: &KinesisFirehoseInputProperty{
//   		ResourceArn: jsii.String("resourceArn"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	KinesisStreamsInput: &KinesisStreamsInputProperty{
//   		ResourceArn: jsii.String("resourceArn"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   }
//
type CfnApplication_InputProperty struct {
	// Describes the format of the data in the streaming source, and how each data element maps to corresponding columns in the in-application stream that is being created.
	//
	// Also used to describe the format of the reference data source.
	InputSchema interface{} `field:"required" json:"inputSchema" yaml:"inputSchema"`
	// Name prefix to use when creating an in-application stream.
	//
	// Suppose that you specify a prefix "MyInApplicationStream." Amazon Kinesis Analytics then creates one or more (as per the `InputParallelism` count you specified) in-application streams with names "MyInApplicationStream_001," "MyInApplicationStream_002," and so on.
	NamePrefix *string `field:"required" json:"namePrefix" yaml:"namePrefix"`
	// Describes the number of in-application streams to create.
	//
	// Data from your source is routed to these in-application input streams.
	//
	// See [Configuring Application Input](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-input.html) .
	InputParallelism interface{} `field:"optional" json:"inputParallelism" yaml:"inputParallelism"`
	// The [InputProcessingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-inputprocessingconfiguration.html) for the input. An input processor transforms records as they are received from the stream, before the application's SQL code executes. Currently, the only input processing configuration available is [InputLambdaProcessor](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-inputlambdaprocessor.html) .
	InputProcessingConfiguration interface{} `field:"optional" json:"inputProcessingConfiguration" yaml:"inputProcessingConfiguration"`
	// If the streaming source is an Amazon Kinesis Firehose delivery stream, identifies the delivery stream's ARN and an IAM role that enables Amazon Kinesis Analytics to access the stream on your behalf.
	//
	// Note: Either `KinesisStreamsInput` or `KinesisFirehoseInput` is required.
	KinesisFirehoseInput interface{} `field:"optional" json:"kinesisFirehoseInput" yaml:"kinesisFirehoseInput"`
	// If the streaming source is an Amazon Kinesis stream, identifies the stream's Amazon Resource Name (ARN) and an IAM role that enables Amazon Kinesis Analytics to access the stream on your behalf.
	//
	// Note: Either `KinesisStreamsInput` or `KinesisFirehoseInput` is required.
	KinesisStreamsInput interface{} `field:"optional" json:"kinesisStreamsInput" yaml:"kinesisStreamsInput"`
}

