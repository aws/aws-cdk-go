package previewawskinesisanalyticsv2mixins


// When you configure the application input for a SQL-based Kinesis Data Analytics application, you specify the streaming source, the in-application stream name that is created, and the mapping between the two.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   inputProperty := &InputProperty{
//   	InputParallelism: &InputParallelismProperty{
//   		Count: jsii.Number(123),
//   	},
//   	InputProcessingConfiguration: &InputProcessingConfigurationProperty{
//   		InputLambdaProcessor: &InputLambdaProcessorProperty{
//   			ResourceArn: jsii.String("resourceArn"),
//   		},
//   	},
//   	InputSchema: &InputSchemaProperty{
//   		RecordColumns: []interface{}{
//   			&RecordColumnProperty{
//   				Mapping: jsii.String("mapping"),
//   				Name: jsii.String("name"),
//   				SqlType: jsii.String("sqlType"),
//   			},
//   		},
//   		RecordEncoding: jsii.String("recordEncoding"),
//   		RecordFormat: &RecordFormatProperty{
//   			MappingParameters: &MappingParametersProperty{
//   				CsvMappingParameters: &CSVMappingParametersProperty{
//   					RecordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   					RecordRowDelimiter: jsii.String("recordRowDelimiter"),
//   				},
//   				JsonMappingParameters: &JSONMappingParametersProperty{
//   					RecordRowPath: jsii.String("recordRowPath"),
//   				},
//   			},
//   			RecordFormatType: jsii.String("recordFormatType"),
//   		},
//   	},
//   	KinesisFirehoseInput: &KinesisFirehoseInputProperty{
//   		ResourceArn: jsii.String("resourceArn"),
//   	},
//   	KinesisStreamsInput: &KinesisStreamsInputProperty{
//   		ResourceArn: jsii.String("resourceArn"),
//   	},
//   	NamePrefix: jsii.String("namePrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-input.html
//
type CfnApplicationPropsMixin_InputProperty struct {
	// Describes the number of in-application streams to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-input.html#cfn-kinesisanalyticsv2-application-input-inputparallelism
	//
	InputParallelism interface{} `field:"optional" json:"inputParallelism" yaml:"inputParallelism"`
	// The [InputProcessingConfiguration](https://docs.aws.amazon.com/managed-flink/latest/apiv2/API_InputProcessingConfiguration.html) for the input. An input processor transforms records as they are received from the stream, before the application's SQL code executes. Currently, the only input processing configuration available is [InputLambdaProcessor](https://docs.aws.amazon.com/managed-flink/latest/apiv2/API_InputLambdaProcessor.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-input.html#cfn-kinesisanalyticsv2-application-input-inputprocessingconfiguration
	//
	InputProcessingConfiguration interface{} `field:"optional" json:"inputProcessingConfiguration" yaml:"inputProcessingConfiguration"`
	// Describes the format of the data in the streaming source, and how each data element maps to corresponding columns in the in-application stream that is being created.
	//
	// Also used to describe the format of the reference data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-input.html#cfn-kinesisanalyticsv2-application-input-inputschema
	//
	InputSchema interface{} `field:"optional" json:"inputSchema" yaml:"inputSchema"`
	// If the streaming source is an Amazon Kinesis Data Firehose delivery stream, identifies the delivery stream's ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-input.html#cfn-kinesisanalyticsv2-application-input-kinesisfirehoseinput
	//
	KinesisFirehoseInput interface{} `field:"optional" json:"kinesisFirehoseInput" yaml:"kinesisFirehoseInput"`
	// If the streaming source is an Amazon Kinesis data stream, identifies the stream's Amazon Resource Name (ARN).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-input.html#cfn-kinesisanalyticsv2-application-input-kinesisstreamsinput
	//
	KinesisStreamsInput interface{} `field:"optional" json:"kinesisStreamsInput" yaml:"kinesisStreamsInput"`
	// The name prefix to use when creating an in-application stream.
	//
	// Suppose that you specify a prefix " `MyInApplicationStream` ." Kinesis Data Analytics then creates one or more (as per the `InputParallelism` count you specified) in-application streams with the names " `MyInApplicationStream_001` ," " `MyInApplicationStream_002` ," and so on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-input.html#cfn-kinesisanalyticsv2-application-input-nameprefix
	//
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
}

