package awskinesisanalytics


// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &CfnApplicationProps{
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
//   					RoleArn: jsii.String("roleArn"),
//   				},
//   			},
//   			KinesisFirehoseInput: &KinesisFirehoseInputProperty{
//   				ResourceArn: jsii.String("resourceArn"),
//   				RoleArn: jsii.String("roleArn"),
//   			},
//   			KinesisStreamsInput: &KinesisStreamsInputProperty{
//   				ResourceArn: jsii.String("resourceArn"),
//   				RoleArn: jsii.String("roleArn"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	ApplicationCode: jsii.String("applicationCode"),
//   	ApplicationDescription: jsii.String("applicationDescription"),
//   	ApplicationName: jsii.String("applicationName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-application.html
//
type CfnApplicationProps struct {
	// Use this parameter to configure the application input.
	//
	// You can configure your application to receive input from a single streaming source. In this configuration, you map this streaming source to an in-application stream that is created. Your application code can then query the in-application stream like a table (you can think of it as a constantly updating table).
	//
	// For the streaming source, you provide its Amazon Resource Name (ARN) and format of data on the stream (for example, JSON, CSV, etc.). You also must provide an IAM role that Amazon Kinesis Analytics can assume to read this stream on your behalf.
	//
	// To create the in-application stream, you need to specify a schema to transform your data into a schematized version used in SQL. In the schema, you provide the necessary mapping of the data elements in the streaming source to record columns in the in-app stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-application.html#cfn-kinesisanalytics-application-inputs
	//
	Inputs interface{} `field:"required" json:"inputs" yaml:"inputs"`
	// One or more SQL statements that read input data, transform it, and generate output.
	//
	// For example, you can write a SQL statement that reads data from one in-application stream, generates a running average of the number of advertisement clicks by vendor, and insert resulting rows in another in-application stream using pumps. For more information about the typical pattern, see [Application Code](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-app-code.html) .
	//
	// You can provide such series of SQL statements, where output of one statement can be used as the input for the next statement. You store intermediate results by creating in-application streams and pumps.
	//
	// Note that the application code must create the streams with names specified in the `Outputs` . For example, if your `Outputs` defines output streams named `ExampleOutputStream1` and `ExampleOutputStream2` , then your application code must create these streams.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-application.html#cfn-kinesisanalytics-application-applicationcode
	//
	ApplicationCode *string `field:"optional" json:"applicationCode" yaml:"applicationCode"`
	// Summary description of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-application.html#cfn-kinesisanalytics-application-applicationdescription
	//
	ApplicationDescription *string `field:"optional" json:"applicationDescription" yaml:"applicationDescription"`
	// Name of your Amazon Kinesis Analytics application (for example, `sample-app` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-application.html#cfn-kinesisanalytics-application-applicationname
	//
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
}

