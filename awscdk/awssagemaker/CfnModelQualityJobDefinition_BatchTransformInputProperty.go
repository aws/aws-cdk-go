package awssagemaker


// The batch transform input for a monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   batchTransformInputProperty := &BatchTransformInputProperty{
//   	DataCapturedDestinationS3Uri: jsii.String("dataCapturedDestinationS3Uri"),
//   	DatasetFormat: &DatasetFormatProperty{
//   		Csv: &CsvProperty{
//   			Header: jsii.Boolean(false),
//   		},
//   		Json: &JsonProperty{
//   			Line: jsii.Boolean(false),
//   		},
//   		Parquet: jsii.Boolean(false),
//   	},
//   	LocalPath: jsii.String("localPath"),
//
//   	// the properties below are optional
//   	EndTimeOffset: jsii.String("endTimeOffset"),
//   	InferenceAttribute: jsii.String("inferenceAttribute"),
//   	ProbabilityAttribute: jsii.String("probabilityAttribute"),
//   	ProbabilityThresholdAttribute: jsii.Number(123),
//   	S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   	S3InputMode: jsii.String("s3InputMode"),
//   	StartTimeOffset: jsii.String("startTimeOffset"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-batchtransforminput.html
//
type CfnModelQualityJobDefinition_BatchTransformInputProperty struct {
	// A URI that identifies the Amazon S3 storage location where Batch Transform Job captures data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-batchtransforminput.html#cfn-sagemaker-modelqualityjobdefinition-batchtransforminput-datacaptureddestinations3uri
	//
	DataCapturedDestinationS3Uri *string `field:"required" json:"dataCapturedDestinationS3Uri" yaml:"dataCapturedDestinationS3Uri"`
	// The dataset format of the data to monitor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-batchtransforminput.html#cfn-sagemaker-modelqualityjobdefinition-batchtransforminput-datasetformat
	//
	DatasetFormat interface{} `field:"required" json:"datasetFormat" yaml:"datasetFormat"`
	// Path to the filesystem where the endpoint data is available to the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-batchtransforminput.html#cfn-sagemaker-modelqualityjobdefinition-batchtransforminput-localpath
	//
	LocalPath *string `field:"required" json:"localPath" yaml:"localPath"`
	// The time offsets in ISO duration format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-batchtransforminput.html#cfn-sagemaker-modelqualityjobdefinition-batchtransforminput-endtimeoffset
	//
	EndTimeOffset *string `field:"optional" json:"endTimeOffset" yaml:"endTimeOffset"`
	// Index or JSONpath to locate predicted label(s).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-batchtransforminput.html#cfn-sagemaker-modelqualityjobdefinition-batchtransforminput-inferenceattribute
	//
	InferenceAttribute *string `field:"optional" json:"inferenceAttribute" yaml:"inferenceAttribute"`
	// Index or JSONpath to locate probabilities.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-batchtransforminput.html#cfn-sagemaker-modelqualityjobdefinition-batchtransforminput-probabilityattribute
	//
	ProbabilityAttribute *string `field:"optional" json:"probabilityAttribute" yaml:"probabilityAttribute"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-batchtransforminput.html#cfn-sagemaker-modelqualityjobdefinition-batchtransforminput-probabilitythresholdattribute
	//
	ProbabilityThresholdAttribute *float64 `field:"optional" json:"probabilityThresholdAttribute" yaml:"probabilityThresholdAttribute"`
	// Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key.
	//
	// Defauts to FullyReplicated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-batchtransforminput.html#cfn-sagemaker-modelqualityjobdefinition-batchtransforminput-s3datadistributiontype
	//
	S3DataDistributionType *string `field:"optional" json:"s3DataDistributionType" yaml:"s3DataDistributionType"`
	// Whether the Pipe or File is used as the input mode for transfering data for the monitoring job.
	//
	// Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-batchtransforminput.html#cfn-sagemaker-modelqualityjobdefinition-batchtransforminput-s3inputmode
	//
	S3InputMode *string `field:"optional" json:"s3InputMode" yaml:"s3InputMode"`
	// The time offsets in ISO duration format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-batchtransforminput.html#cfn-sagemaker-modelqualityjobdefinition-batchtransforminput-starttimeoffset
	//
	StartTimeOffset *string `field:"optional" json:"startTimeOffset" yaml:"startTimeOffset"`
}

