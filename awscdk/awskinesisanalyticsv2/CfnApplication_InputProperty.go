package awskinesisanalyticsv2


// When you configure the application input for a SQL-based Managed Service for Apache Flink application, you specify the streaming source, the in-application stream name that is created, and the mapping between the two.
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
//   		},
//   	},
//   	KinesisFirehoseInput: &KinesisFirehoseInputProperty{
//   		ResourceArn: jsii.String("resourceArn"),
//   	},
//   	KinesisStreamsInput: &KinesisStreamsInputProperty{
//   		ResourceArn: jsii.String("resourceArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-input.html
//
type CfnApplication_InputProperty struct {
	// Describes the format of the data in the streaming source, and how each data element maps to corresponding columns in the in-application stream that is being created.
	//
	// Also used to describe the format of the reference data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-input.html#cfn-kinesisanalyticsv2-application-input-inputschema
	//
	InputSchema interface{} `field:"required" json:"inputSchema" yaml:"inputSchema"`
	// The name prefix to use when creating an in-application stream.
	//
	// Suppose that you specify a prefix " `MyInApplicationStream` ." Managed Service for Apache Flink then creates one or more (as per the `InputParallelism` count you specified) in-application streams with the names " `MyInApplicationStream_001` ," " `MyInApplicationStream_002` ," and so on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-input.html#cfn-kinesisanalyticsv2-application-input-nameprefix
	//
	NamePrefix *string `field:"required" json:"namePrefix" yaml:"namePrefix"`
	// Describes the number of in-application streams to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-input.html#cfn-kinesisanalyticsv2-application-input-inputparallelism
	//
	InputParallelism interface{} `field:"optional" json:"inputParallelism" yaml:"inputParallelism"`
	// The [InputProcessingConfiguration](https://docs.aws.amazon.com/kinesisanalytics/latest/apiv2/API_InputProcessingConfiguration.html) for the input. An input processor transforms records as they are received from the stream, before the application's SQL code executes. Currently, the only input processing configuration available is [InputLambdaProcessor](https://docs.aws.amazon.com/kinesisanalytics/latest/apiv2/API_InputLambdaProcessor.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-input.html#cfn-kinesisanalyticsv2-application-input-inputprocessingconfiguration
	//
	InputProcessingConfiguration interface{} `field:"optional" json:"inputProcessingConfiguration" yaml:"inputProcessingConfiguration"`
	// If the streaming source is an Amazon Kinesis Data Firehose delivery stream, identifies the delivery stream's ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-input.html#cfn-kinesisanalyticsv2-application-input-kinesisfirehoseinput
	//
	KinesisFirehoseInput interface{} `field:"optional" json:"kinesisFirehoseInput" yaml:"kinesisFirehoseInput"`
	// If the streaming source is an Amazon Kinesis data stream, identifies the stream's Amazon Resource Name (ARN).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-input.html#cfn-kinesisanalyticsv2-application-input-kinesisstreamsinput
	//
	KinesisStreamsInput interface{} `field:"optional" json:"kinesisStreamsInput" yaml:"kinesisStreamsInput"`
}

