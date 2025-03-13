package awssagemaker


// Input object for the batch transform job.
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
//   	ExcludeFeaturesAttribute: jsii.String("excludeFeaturesAttribute"),
//   	S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   	S3InputMode: jsii.String("s3InputMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-batchtransforminput.html
//
type CfnMonitoringSchedule_BatchTransformInputProperty struct {
	// The Amazon S3 location being used to capture the data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-batchtransforminput.html#cfn-sagemaker-monitoringschedule-batchtransforminput-datacaptureddestinations3uri
	//
	DataCapturedDestinationS3Uri *string `field:"required" json:"dataCapturedDestinationS3Uri" yaml:"dataCapturedDestinationS3Uri"`
	// The dataset format for your batch transform job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-batchtransforminput.html#cfn-sagemaker-monitoringschedule-batchtransforminput-datasetformat
	//
	DatasetFormat interface{} `field:"required" json:"datasetFormat" yaml:"datasetFormat"`
	// Path to the filesystem where the batch transform data is available to the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-batchtransforminput.html#cfn-sagemaker-monitoringschedule-batchtransforminput-localpath
	//
	LocalPath *string `field:"required" json:"localPath" yaml:"localPath"`
	// The attributes of the input data to exclude from the analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-batchtransforminput.html#cfn-sagemaker-monitoringschedule-batchtransforminput-excludefeaturesattribute
	//
	ExcludeFeaturesAttribute *string `field:"optional" json:"excludeFeaturesAttribute" yaml:"excludeFeaturesAttribute"`
	// Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key.
	//
	// Defaults to `FullyReplicated`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-batchtransforminput.html#cfn-sagemaker-monitoringschedule-batchtransforminput-s3datadistributiontype
	//
	S3DataDistributionType *string `field:"optional" json:"s3DataDistributionType" yaml:"s3DataDistributionType"`
	// Whether the `Pipe` or `File` is used as the input mode for transferring data for the monitoring job.
	//
	// `Pipe` mode is recommended for large datasets. `File` mode is useful for small files that fit in memory. Defaults to `File` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-batchtransforminput.html#cfn-sagemaker-monitoringschedule-batchtransforminput-s3inputmode
	//
	S3InputMode *string `field:"optional" json:"s3InputMode" yaml:"s3InputMode"`
}

