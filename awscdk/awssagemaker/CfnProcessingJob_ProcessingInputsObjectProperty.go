package awssagemaker


// The inputs for a processing job.
//
// The processing input must specify exactly one of either `S3Input` or `DatasetDefinition` types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   processingInputsObjectProperty := &ProcessingInputsObjectProperty{
//   	InputName: jsii.String("inputName"),
//
//   	// the properties below are optional
//   	AppManaged: jsii.Boolean(false),
//   	DatasetDefinition: &DatasetDefinitionProperty{
//   		AthenaDatasetDefinition: &AthenaDatasetDefinitionProperty{
//   			Catalog: jsii.String("catalog"),
//   			Database: jsii.String("database"),
//   			OutputFormat: jsii.String("outputFormat"),
//   			OutputS3Uri: jsii.String("outputS3Uri"),
//   			QueryString: jsii.String("queryString"),
//
//   			// the properties below are optional
//   			KmsKeyId: jsii.String("kmsKeyId"),
//   			OutputCompression: jsii.String("outputCompression"),
//   			WorkGroup: jsii.String("workGroup"),
//   		},
//   		DataDistributionType: jsii.String("dataDistributionType"),
//   		InputMode: jsii.String("inputMode"),
//   		LocalPath: jsii.String("localPath"),
//   		RedshiftDatasetDefinition: &RedshiftDatasetDefinitionProperty{
//   			ClusterId: jsii.String("clusterId"),
//   			ClusterRoleArn: jsii.String("clusterRoleArn"),
//   			Database: jsii.String("database"),
//   			DbUser: jsii.String("dbUser"),
//   			OutputFormat: jsii.String("outputFormat"),
//   			OutputS3Uri: jsii.String("outputS3Uri"),
//   			QueryString: jsii.String("queryString"),
//
//   			// the properties below are optional
//   			KmsKeyId: jsii.String("kmsKeyId"),
//   			OutputCompression: jsii.String("outputCompression"),
//   		},
//   	},
//   	S3Input: &S3InputProperty{
//   		S3DataType: jsii.String("s3DataType"),
//   		S3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		LocalPath: jsii.String("localPath"),
//   		S3CompressionType: jsii.String("s3CompressionType"),
//   		S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   		S3InputMode: jsii.String("s3InputMode"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-processinginputsobject.html
//
type CfnProcessingJob_ProcessingInputsObjectProperty struct {
	// The name for the processing job input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-processinginputsobject.html#cfn-sagemaker-processingjob-processinginputsobject-inputname
	//
	InputName *string `field:"required" json:"inputName" yaml:"inputName"`
	// When `True` , input operations such as data download are managed natively by the processing job application.
	//
	// When `False` (default), input operations are managed by Amazon SageMaker.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-processinginputsobject.html#cfn-sagemaker-processingjob-processinginputsobject-appmanaged
	//
	AppManaged interface{} `field:"optional" json:"appManaged" yaml:"appManaged"`
	// Configuration for Dataset Definition inputs.
	//
	// The Dataset Definition input must specify exactly one of either `AthenaDatasetDefinition` or `RedshiftDatasetDefinition` types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-processinginputsobject.html#cfn-sagemaker-processingjob-processinginputsobject-datasetdefinition
	//
	DatasetDefinition interface{} `field:"optional" json:"datasetDefinition" yaml:"datasetDefinition"`
	// Configuration for downloading input data from Amazon S3 into the processing container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-processinginputsobject.html#cfn-sagemaker-processingjob-processinginputsobject-s3input
	//
	S3Input interface{} `field:"optional" json:"s3Input" yaml:"s3Input"`
}

